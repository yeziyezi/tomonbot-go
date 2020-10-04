package util

import (
	"golang.org/x/net/websocket"
	"tomonbot-go/src/one/yezii/tomon/util/ws"
)

type TomonContext struct {
	SessionId         string
	WsConn            *websocket.Conn
	HeartbeatInterval int
	Conf              *Config
	logger            *YLogger
}

func NewTomonContext() TomonContext {
	var ctx TomonContext
	ctx.logger = NewYLogger("tomonContext")
	return ctx
}
func (it *TomonContext) Send(v interface{}) error {
	return it.logger.ErrOrNil(websocket.JSON.Send(it.WsConn, v))
}
func (it *TomonContext) DoAuth() error {
	return it.Send(ws.AuthMessageForSend(it.Conf.Token))
}
func (it *TomonContext) SendHeartbeat() error {
	return it.Send(ws.HeartbeatMessageForSend())
}
func (it *TomonContext) Receive(v interface{}) error {
	return it.logger.ErrOrNil(websocket.JSON.Receive(it.WsConn, v))
}
