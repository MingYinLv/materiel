package handle

import (
	"github.com/gin-gonic/gin"
	//"time"
	"fmt"
	"time"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "x-requested-with, content-type, authorization")
	fmt.Println("cors")
	// 所有请求暂停1秒
	time.Sleep(time.Second * 1)

}
