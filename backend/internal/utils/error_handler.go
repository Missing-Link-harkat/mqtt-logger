package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleError(c *gin.Context, err error) {

	var statusCode int
	var errorMessage string

	if err == gorm.ErrRecordNotFound {
		statusCode = http.StatusNotFound
		errorMessage = "Not found"
	} else {
		statusCode = http.StatusInternalServerError
		errorMessage = fmt.Sprintf("Error: %v", err)
	}
	c.JSON(statusCode, errorMessage)
}
