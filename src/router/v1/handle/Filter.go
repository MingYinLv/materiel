package handle

import (
	"github.com/gin-gonic/gin"
	//"time"
	"fmt"
)

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	c.Header("Access-Control-Allow-Headers", "x-requested-with, content-type, authorization")
	fmt.Println("cors")
	// 所有请求暂停2秒
	//time.Sleep(time.Second * 2)

}
