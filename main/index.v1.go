package main

import (
	"github.com/gin-gonic/gin"
	"materiel/src/router/v1/handle"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	v1 := router.Group("/v1", handle.Filter)
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
		materiel.POST("", handle.AddMateriel)
		// 修改物料
		materiel.PUT("/:id", handle.UpdateMateriel)
		// 获取物料信息
		materiel.GET("/:id", handle.GetMaterielById)

	}
	router.Run(":3333")
}
