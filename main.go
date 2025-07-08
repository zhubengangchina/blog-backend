package main

import (
	"blog-backend/config"
	"blog-backend/middleware"
	"blog-backend/routes"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	//初始化数据库连接
	config.InitEnv()
	config.InitLogger()
	config.InitDB()
	utils.InitJWT()

	//注册路由
	routes.RegisterAuthRoutes(r)
	routes.RegisterPostRoutes(r)
	routes.RegisterCommentRoutes(r)
	r.Run(":8080")
}
