package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 测试index
// @Description 用来测试swagger
// @Tags
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Fialure 500 {string} notwelcom
// @Router /user/getUserList [get]

func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"data": data,
	})
}
