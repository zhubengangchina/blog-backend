package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"blog-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//密码加密
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		config.Logger.Error("密码加密失败", zap.Error(err))
		utils.Fail(c, http.StatusInternalServerError, "密码加密失败")
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	user := models.User{
		UserName: input.Username,
		Password: hash,
		Email:    input.Email,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或邮箱已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	//验证账号是否存在
	if err := config.DB.Where("user_name = ?", input.Username).Find(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名不存在"})
		return
	}

	//验证密码
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误"})
		return
	}

	//生成JWT
	token, err := utils.GenerateJWT(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
