package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"strconv"
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
	ind, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
	)
	if err != nil {
		panic(err)
	}

	userId, _ := strconv.ParseUint(c.Query("userId"), 10, 32)
	token := c.Query("token")
	user := models.UserBasic{}

	user.ID = uint(userId)
	user.Identity = token
	ind.Execute(c.Writer, user)
	c.JSON(0, gin.H{
		"code": 0,
		"data": user,
	})
}
