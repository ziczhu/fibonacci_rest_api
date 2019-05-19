package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ziczhu/fibonacci_rest_api/pkg/config"
	"github.com/ziczhu/fibonacci_rest_api/pkg/fibonacci"
)

// GetFibonacciSequence parses the number in path and calculate the first N fibonacci sequence
// The N start from 0
func GetFibonacciSequence(conf *config.Config) gin.HandlerFunc {
	fib := fibonacci.New(conf.InitFibCacheSize, conf.MaxFibCacheSize)
	return func(c *gin.Context) {
		numberStr := c.Param("n")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			RespondWithError(c, http.StatusBadRequest, InvalidNumberErrorCode,
				fmt.Sprintf("invalid number '%s'", numberStr))
			return
		}

		if number < 0 {
			RespondWithError(c, http.StatusBadRequest, NegativeNumberErrorCode,
				fmt.Sprintf("negative number '%s', only accepts number >= 0", numberStr))
			return
		}

		if number > conf.MaxFibInput {
			RespondWithError(c, http.StatusBadRequest, OverLimitNumberErrorCode,
				fmt.Sprintf("number '%s' is too big, only accepts number <= %d", numberStr, conf.MaxFibInput))
			return
		}

		list := fib.GetSequence(number)
		RespondWithStatus(c, http.StatusOK, list)
	}
}
