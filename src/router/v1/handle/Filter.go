package handle

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	c.Header("Access-Control-Allow-Headers", "x-requested-with, content-type, authorization")
	// 所有请求暂停2秒
	time.Sleep(time.Second * 2)

}
