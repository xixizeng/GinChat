package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath

// GetIndex PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags testswag
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/index [get]
func GetIndex(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
