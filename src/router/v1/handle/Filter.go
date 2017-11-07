package handle

import "github.com/gin-gonic/gin"

func Filter(c *gin.Context) {
	c.GetPostForm("")
}
