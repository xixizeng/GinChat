package service

import (
	"github.com/gin-gonic/gin"
	"text/template"
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
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(g.Writer, "index")
	//g.JSON(http.StatusOK, "helloworld")
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/main.shtml")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
}
