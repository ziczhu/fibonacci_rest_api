package api

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status    int       `json:"status"`
	ErrorCode ErrorCode `json:"error_code"`
	ErrorMsg  string    `json:"error_msg"`
}

// RespondWithStatus defines the structure of success response
func RespondWithStatus(c *gin.Context, status int, data interface{}) {
	respBody := SuccessResponse{
		Status: status,
		Data:   data,
	}
	c.JSON(status, respBody)
}

// RespondWithError defines the structure of error response
func RespondWithError(c *gin.Context, status int, errorCode ErrorCode, errorMsg string) {
	respBody := ErrorResponse{
		Status:    status,
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
	}
	c.JSON(status, respBody)
}
