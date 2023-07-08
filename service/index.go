package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Summary 测试index
// @Description 用来测试swagger
// @Tags
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Fialure 500 {string} notwelcom
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
