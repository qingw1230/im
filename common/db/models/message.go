package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/qingw1230/im/common/utils"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendID     int64  // 发送者
	ReceiveID  int64  // 接受者
	Type       int    // 发送类型  1私聊  2群聊  3心跳
	Media      int    // 消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string // 消息内容
	CreateTime uint64 // 创建时间
	ReadTime   uint64 // 读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int // 其他数字统计
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)
var rw sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("sendID")
	sendID, _ := strconv.ParseInt(id, 10, 64)
	// ReceiveID := query.Get("receiveID")
	// msgType := query.Get("type")
	// content := query.Get("content")
	isValid := true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	rw.Lock()
	clientMap[sendID] = node
	rw.Unlock()

	go sendProc(node)
	go recvProc(node)
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
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
		fmt.Println("[ws] <<<<<", data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
	fmt.Println("init goroutine ")
}

func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(172, 23, 137, 168),
		Port: viper.GetInt("port.udp"),
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpsendChan:
			fmt.Println("udpSendProc  data :", string(data))
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: viper.GetInt("port.udp"),
	})
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc  data :", string(buf[0:n]))
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	msg := Message{}
	msg.CreateTime = uint64(time.Now().Unix())
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		fmt.Println("dispatch  data :", string(data))
		sendMsg(msg.ReceiveID, data)
	}
}

func sendMsg(userId int64, msg []byte) {
	rw.RLock()
	node, ok := clientMap[userId]
	rw.RUnlock()
	jsonMsg := Message{}
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.SendID))
	jsonMsg.CreateTime = uint64(time.Now().Unix())
	r, err := utils.RedisDB.Get(ctx, "online_"+userIdStr).Result()
	if err != nil {
		fmt.Println(err)
	}
	if r != "" {
		if ok {
			fmt.Println("sendMsg >>> userID: ", userId, "  msg:", string(msg))
			node.DataQueue <- msg
		}
	}
	var key string
	if userId > jsonMsg.SendID {
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}
	res, err := utils.RedisDB.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	score := float64(cap(res)) + 1
	ress, e := utils.RedisDB.ZAdd(ctx, key, redis.Z{Score: score, Member: msg}).Result()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ress)
}
