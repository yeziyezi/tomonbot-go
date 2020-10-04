package main

import (
	"golang.org/x/net/websocket"
	"time"
	"tomonbot-go/src/one/yezii/tomon/util"
)

var logger = util.NewYLogger("tomonMain")
var ctx = util.NewTomonContext()
var inChan = make(chan *messageFromServer, 10)

type messageFromServer struct {
	Op int                    `json:"op"`
	D  map[string]interface{} `json:"d"`
	E  string                 `json:"e"`
}
type heartbeatMessageFromServer struct {
	Op int `json:"op"`
	D  struct {
		HeartbeatInterval int    `json:"heartbeat_interval"`
		ServerTs          string `json:"server_ts"`
		SessionId         string `json:"session_id"`
	} `json:"d"`
}

func initConn() error {
	conn, err := websocket.Dial("wss://gateway.tomon.co", "", "https://beta.tomon.co")
	ctx.WsConn = conn
	return logger.ErrOrNil(err)
}

//解析心跳间隔和session id
//todo 如果超过一定时间没有接收到或接收到的数据无法正确解析则可能是认证失败，加log
func resolveHeartbeatInterval() error {
	var hm *heartbeatMessageFromServer
	err := ctx.Receive(&hm)
	if err != nil {
		return err
	}
	ctx.SessionId = hm.D.SessionId
	ctx.HeartbeatInterval = hm.D.HeartbeatInterval
	return nil
}
func main() {
	//获取配置
	conf, err := util.ReadConfig()
	if err != nil {
		return
	}
	ctx.Conf = conf
	//建立websocket连接
	err = initConn()
	if err != nil {
		return
	}
	//鉴权
	err = ctx.DoAuth()
	if err != nil {
		return
	}
	//解析出心跳间隔时长
	err = resolveHeartbeatInterval()
	if err != nil {
		return
	}
	//接收消息推送到chan中
	go startListeningMessage()
	//发送心跳
	go startSendHeartbeat()
	//处理接收到的消息
	go startDealingMessage()
	blockingWaiting()
}

func blockingWaiting() {
	for {
		time.Sleep(1 * time.Hour)
	}
}
func startListeningMessage() {
	for {
		var msg *messageFromServer
		err := ctx.Receive(&msg)
		if err != nil {
			continue
		}
		inChan <- msg
	}
}
func startSendHeartbeat() {
	for {
		//如果发送心跳失败可能是连接断开，尝试重新连接
		if ctx.SendHeartbeat() != nil {
			logger.Warn("try reconnect to server")
			if ctx.DoAuth() != nil {
				_ = resolveHeartbeatInterval()
			}
		}
		time.Sleep(time.Duration(ctx.HeartbeatInterval/2) * time.Millisecond)
	}
}
func startDealingMessage() {
	for msg := range inChan {
		//在心跳发送成功后server端会返回一个Op为4的message
		if msg.Op != 4 {
			//todo 修改为实际的业务逻辑
			logger.Info(msg)
		}
	}
}
