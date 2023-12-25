package utilities

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ResponseBadRequest is ... params[0] = status message, params[1] = status code
func ResponseBadRequest(c *gin.Context, params ...string) {
	//WriteResponseHeader(c, -1, "Bad Request")
	message := "Bad Request"
	code := 400
	if len(params) > 0 {
		message = params[0]
	}
	if len(params) > 1 {
		cs, err := strconv.Atoi(params[1])
		if err != nil {
			code = cs
		}
	}
	c.JSON(http.StatusBadRequest, &gin.H{
		"status":      message,
		"message":     message,
		"code_return": code,
	})
}
