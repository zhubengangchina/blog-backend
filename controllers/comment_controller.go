package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCommentInput struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// 添加评论
func CreateComment(c *gin.Context) {
	var input CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	userID, _ := c.Get("userID")
	//检查文章是否存在
	var post models.Post
	if err := config.DB.First(&post, input.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	var comment = models.Comment{
		Content: input.Content,
		PostID:  post.ID,
		UserID:  userID.(uint),
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// 获取某篇文章下的所有评论（公开接口）
func GetCommentByPost(c *gin.Context) {
	postID := c.Param("postID")
	var comments []models.Comment
	if err := config.DB.
		Model(&models.Comment{}).
		Preload("User").
		Order("created_at asc").
		Where("post_id = ?", postID).
		Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}
