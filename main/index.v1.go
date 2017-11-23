package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"materiel/src/config"
	"materiel/src/router/v1/handle"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var listenPort int64
	config.Initial(&listenPort)
	router := gin.Default()
	v1 := router.Group("/v1", handle.Cors)
	//router.OPTIONS("/:path(/:path1(/:path2))", handle.Cors)
	router.OPTIONS("/:path", handle.Cors)
	router.OPTIONS("/:path/:path1", handle.Cors)
	router.OPTIONS("/:path/:path1/:path2", handle.Cors)
	// 登录
	v1.POST("/login", handle.UserLogin)
	{
		// 验证登录
		user := v1.Group("/user", handle.CheckLogin)
		// 获取用户信息
		user.GET("", handle.Info)
		// 更新token
		user.POST("/refresh_token", handle.RefreshToken)
	}
	{
		materiel := v1.Group("/materiels")
		// 查询列表
		materiel.GET("", handle.GetMaterielList)
		// 添加物料
		materiel.POST("", handle.CheckLogin, handle.AddMateriel)
		// 修改物料
		materiel.PUT("/:id", handle.CheckLogin, handle.UpdateMateriel)
		// 获取物料信息
		materiel.GET("/:id", handle.GetMaterielById)
	}
	{
		log := v1.Group("/logs")
		log.GET("/:id", handle.GetLogById)
		log.GET("", handle.GetLogList)
	}
	router.Static("/static", "../static")
	router.Run(fmt.Sprintf(":%d", listenPort))
}
