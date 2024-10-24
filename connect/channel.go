package connect

import (
	"errors"
	"net"
	"zinx/proto"
	"github.com/gorilla/websocket"
)

type Channel struct {
	Room      *Room           //*所在聊天室
	Next      *Channel        //*下一个会话
	Prev      *Channel        //*上一个会话
	broadcast chan *proto.Msg //*广播消息
	userId    int             //*用户id
	conn      *websocket.Conn //*websocket连接
	connTcp   *net.TCPConn    //*tcp连接
}

// *创建一个新的会话
func NewChannel(size int) (c *Channel) {
	c = new(Channel)
	c.broadcast = make(chan *proto.Msg, size)
	c.Next = nil
	c.Prev = nil
	return
}

// *将消息无阻塞的放入广播队列
func (ch *Channel) Push(msg *proto.Msg) (err error) {
	select {
	case ch.broadcast <- msg:
	default:
		err = errors.New("channel is full")
	}
	return
}
