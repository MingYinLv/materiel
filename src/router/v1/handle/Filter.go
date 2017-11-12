package handle

import "github.com/gin-gonic/gin"

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	c.Header("Access-Control-Allow-Headers", "x-requested-with, content-type, authorization")
}
