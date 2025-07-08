package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
| 功能     | 说明                |
| ------ | ----------------- |
| 创建文章   | 登录用户可创建文章（JWT 鉴权） |
| 查询文章列表 | 公开接口，分页可选         |
| 查询文章详情 | 公开接口，通过 `id` 获取   |
| 更新文章   | 登录用户，且只能修改自己发布的文章 |
| 删除文章   | 登录用户，且只能删除自己发布的文章 |

*/

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建文章（需登录）
func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID.(uint),
	}
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// 获取所有文章
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	var total int64

	//读取查询参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	keyword := c.DefaultQuery("keyword", "")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	//构建查询
	query := config.DB.Model(&models.Post{}).Preload("User")
	if keyword != "" {
		query.Where("title LIKE ?", "%"+keyword+"%")
	}
	//获取总数
	query.Count(&total)

	//查询分页数据
	if err := query.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":      posts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// 获取单篇文章
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// 更新文章（只能作者自己更新）
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	userID, _ := c.Get("userID")
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改此文章"})
		return
	}

	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = input.Title
	post.Content = input.Content

	config.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// 删除文章（只能作者自己删除）
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}
	userID, _ := c.Get("userID")
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除此文章"})
		return
	}

	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
