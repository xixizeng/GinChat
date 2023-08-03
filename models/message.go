package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Message struct {
	gorm.Model
	FromID   int64  //发送者
	TargetId int64  //接收者
	Type     int    //消息类型 :群聊 私聊 广播
	Media    int    //消息类型 :文字 图片 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 需要：发送者ID，接收者ID,消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request) {
	//1】 获取参数并校验合法性 校验token 等合法性
	//token := query.Get("token")
	query := r.URL.Query()
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalida := true //checkToken()
	conn, err := (&websocket.Upgrader{

		//token校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	//2获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//3、用户关系

	//4、userid跟 node绑定 并加锁
	rwLocker.Lock() //上锁
	clientMap[userId] = node
	rwLocker.Unlock() //开锁
	//5.完成发送逻辑
	go sendProc(node)
	//6.完成接收逻辑
	go recvProc(node)
	sendMsg(userId, []byte("欢迎进入测试"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				panic(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws]	<<<<", data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp发送
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 2, 1),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			_, err := conn.Write(data)
			if err != nil {
				panic(err)
				return
			}
		}
	}
}

// 完成udp数据接收协议
func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			panic(err)
		}
		dispatch(buf[0:n])
	}
}

// 后端调度处理逻辑
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
	}
	switch msg.Type {
	case 1: //1私信
		sendMsg(msg.TargetId, data)
	case 2: //群发
		sendGroupMsg()
	case 3: //广播
		sendAllMsg()
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

func sendGroupMsg() {

}

func sendAllMsg() {
}
