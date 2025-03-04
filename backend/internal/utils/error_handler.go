package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleError(c *gin.Context, err error) {
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve resource"})
	}
}
