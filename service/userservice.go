package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /user

// GetUserList PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserList
// @Router /getuserlist [get]
func GetUserList(g *gin.Context) {
	data := models.GetUserList()
	g.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
