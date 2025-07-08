package routes

import (
	"blog-backend/controllers"
	"blog-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(c *gin.Engine) {
	comments := c.Group("/comments")

	//获取文章的评论（公开接口）
	comments.GET("/post/:postID", controllers.GetCommentByPost)

	//需要登录才能操作
	comments.Use(middleware.JWTAuthMiddleware())
	{
		comments.POST("", controllers.CreateComment)
	}
}
