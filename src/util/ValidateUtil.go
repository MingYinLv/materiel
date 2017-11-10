package util

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RequiredValidate(c *gin.Context, fields []string) bool {
	for _, k := range fields {
		if v, b := c.GetPostForm(k); !b || strings.TrimSpace(v) == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  fmt.Sprintf("%s is required", k),
				"detail": []map[string]string{{"message": "required", "title": k}},
			})
			return false
		}
	}
	return true
}

func IntValidate(c *gin.Context, fields []string) bool {
	for _, k := range fields {
		if v, b := c.GetPostForm(k); !b || !govalidator.IsInt(v) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  fmt.Sprintf("%s is invalid", k),
				"detail": []map[string]string{{"message": "must be int", "title": k}},
			})
			return false
		}
	}
	return true
}
