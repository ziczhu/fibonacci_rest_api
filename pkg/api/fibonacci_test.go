package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ziczhu/fibonacci_rest_api/pkg/config"
)

const (
	fibonacciRouterPath = "/api/v1/fibonacci/:n"
	fibonacciTestPath   = "/api/v1/fibonacci/%s"
)

func setupFibRouter() *gin.Engine {
	r := gin.Default()
	conf := config.GetConfig()
	r.GET(fibonacciRouterPath, GetFibonacciSequence(conf))
	return r
}

func testErrorBody(t *testing.T, body []byte,
	expectedStatus int, expectedErrorCode ErrorCode, expectedErrorMsg string) {
	var resp ErrorResponse
	err := json.Unmarshal(body, &resp)
	require.NoError(t, err, "unmarshal error body should not fail")
	assert.Equal(t, expectedStatus, resp.Status)
	assert.Equal(t, expectedErrorCode, resp.ErrorCode)
	assert.Equal(t, expectedErrorMsg, resp.ErrorMsg)
}

func testSuccessBody(t *testing.T, body []byte, expectedStatus int, expectedData interface{}) {
	var resp struct {
		Status int   `json:"status"`
		Data   []int `json:"data"`
	}
	err := json.Unmarshal(body, &resp)
	require.NoError(t, err, "unmarshal success body should not fail")
	assert.Equal(t, expectedStatus, resp.Status)
	assert.Equal(t, expectedData, resp.Data)
}

func TestGetFibonacciSequence(t *testing.T) {
	router := setupFibRouter()

	t.Run("It should return appropriate error when the input is not a number", func(t *testing.T) {
		w := httptest.NewRecorder()
		num := "notNumber"
		req, _ := http.NewRequest("GET", fmt.Sprintf(fibonacciTestPath, num), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		testErrorBody(t, w.Body.Bytes(), http.StatusBadRequest, InvalidNumberErrorCode,
			fmt.Sprintf("invalid number '%s'", num))
	})

	t.Run("It should return appropriate error when the input is a negative number", func(t *testing.T) {
		w := httptest.NewRecorder()
		num := "-100"
		req, _ := http.NewRequest("GET", fmt.Sprintf(fibonacciTestPath, num), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		testErrorBody(t, w.Body.Bytes(), http.StatusBadRequest, NegativeNumberErrorCode,
			fmt.Sprintf("negative number '%s', only accepts number >= 0", num))
	})

	t.Run("It should return appropriate error when the input is larger than limit", func(t *testing.T) {
		w := httptest.NewRecorder()
		num := "100000"
		req, _ := http.NewRequest("GET", fmt.Sprintf(fibonacciTestPath, "100000"), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		conf := config.GetConfig()
		testErrorBody(t, w.Body.Bytes(), http.StatusBadRequest, OverLimitNumberErrorCode,
			fmt.Sprintf("number '%s' is too big, only accepts number <= %d", num, conf.MaxFibInput))
	})

	t.Run("It should return ok when the input is a positive number", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf(fibonacciTestPath, "5"), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		testSuccessBody(t, w.Body.Bytes(), http.StatusOK, []int{0, 1, 1, 2, 3})
	})
}
