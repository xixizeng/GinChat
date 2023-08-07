package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// @BasePath

// GetUserList PingExample godoc
// @Summary 拉取所有用户
// @Schemes
// @Description do ping
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} models.UserBasic "用户信息"
// @Router /user/getUserList [get]
func GetUserList(g *gin.Context) {
	data := models.GetUserList()
	g.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// CreateUser PingExample godoc
// @Summary 新增用户
// @Schemes
// @Description do ping
// @Tags 用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Accept json
// @Produce json
// @Success 200 {object} models.UserBasic "用户信息"
// @Router /user/createUser [put]
func CreateUser(c *gin.Context) {

	Name := c.Request.FormValue("name")
	data := models.FindUserByName(Name)
	if data.Name != "" {
		c.JSON(200, gin.H{
			"message": "该用户已注册",
		})
		return
	}
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println()
	if password == "" || Name == "" {
		c.JSON(200, gin.H{
			"message": "用户名或密码不能为空",
		})
		return
	}

	if password != repassword {
		c.JSON(200, gin.H{
			"message": "两次密码输入不一致",
		})
		return
	}
	user := models.UserBasic{}
	user.Name = Name
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, user.Salt)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功!",
	})
}

// DeleteUser PingExample godoc
// @Summary 删除用户
// @Schemes
// @Description do ping
// @Tags 用户
// @param id query string false "ID"
// @Accept json
// @Produce json
// @Success 200 {object} models.UserBasic "用户信息"
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		panic("ID Atoi failed")
	}
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除用户成功!",
	})
}

// UpdateUser PingExample godoc
// @Summary 修改用户
// @Schemes
// @Description do ping
// @Tags 用户
// @Param id formData string true "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Param phone formData string false "phone"
// @Param email formData string false "email"
// @Success 200 {string} {"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.PostForm("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "ID格式有误!",
			"error":   fmt.Sprintf("%s", err),
		})
		return
	}

	user.ID = uint(id)
	data := models.FindUserByID(user.ID)
	if data.Name == "" {
		c.JSON(501, gin.H{
			"code":    500,
			"message": "该ID不存在",
		})
		return
	}
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "修改参数不匹配",
		})
		return
	}

	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "修改用户成功!",
	})
}

// FindUserByNameAndPassword	PingExample godoc
// @Summary 寻找用户
// @Schemes
// @Description do ping
// @Tags 用户
// @Param name query string false "name"
// @Param password query string false "password"
// @Success 200 {object} models.UserBasic
// @Router /user/findUserByNameAndPassword [post]
func FindUserByNameAndPassword(g *gin.Context) {
	Name := g.PostForm("name")
	PassWord := g.PostForm("password")
	fmt.Println(Name, PassWord)
	user := models.FindUserByName(Name)
	if user.Name == "" {
		g.JSON(http.StatusOK, gin.H{
			"code":    501,
			"message": "用户名或密码不正确1",
		})
		return
	}

	_, flag := utils.ValidPassword(PassWord, user.Salt, user.PassWord)
	if !flag {
		g.JSON(http.StatusOK, gin.H{
			"code":    501,
			"message": "用户名或密码不正确2",
		})
		return
	}

	user = models.FindUserByNameAndPassword(Name, user.PassWord)
	if user.Name == "" {
		g.JSON(http.StatusOK, gin.H{
			"code":    501,
			"message": "用户名或密码不正确3",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    user,
	})
}

// 防止跨域站点的伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func SenMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			panic(err)
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		panic(err)
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s]:%s ", tm, msg)
	ws.WriteMessage(1, []byte(m))
	if err != nil {
		panic(err)
	}
}

func SenUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

func SearchFriends(c *gin.Context) {
	id, err := strconv.ParseUint(c.PostForm("userId"), 10, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "用户ID有误",
		})
		return
	}
	//c.JSON(500, gin.H{
	//	"code":    -1,
	//	"message": "获取用户关联消息成功",
	//	"data":    models.SearchFriends(uint(id)),
	//})
	users := models.SearchFriends(uint(id))

	utils.RespOkList(c.Writer, users, len(users))
}
