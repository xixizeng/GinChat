package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} welcom
// @Fialure 500 {string} notwelcom
// @Router /user/getUserListt [get]

func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"data": data,
	})
}
