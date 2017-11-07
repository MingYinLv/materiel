package main

import (
	"github.com/gin-gonic/gin"
	"materiel/src/router/v1/handle"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("/v1", handle.Filter)
	{
		user := v1.Group("/user", handle.CheckLogin)
		user.GET("/info", handle.Info)
		v1.POST("/login", handle.UserLogin)
	}
	router.Run(":3333")
}
