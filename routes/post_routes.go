package routes

import (
	"blog-backend/controllers"
	"blog-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {

	posts := r.Group("/posts")

	//公开接口
	posts.GET("", controllers.GetAllPosts)
	posts.GET("/:id", controllers.GetPostByID)

	posts.Use(middleware.JWTAuthMiddleware())
	{
		// 登录后才可访问
		posts.POST("", controllers.CreatePost)
		posts.PUT("/:id", controllers.UpdatePost)
		posts.DELETE("/:id", controllers.DeletePost)
	}

}
