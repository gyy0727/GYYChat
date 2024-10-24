package connect

import (
	"errors"
	"sync"
	"zinx/proto"
	"github.com/sirupsen/logrus"
)

const NoRoom = -1

type Room struct {
	Id          int          // *聊天室id
	OnlineCount int          // *在线人数
	rLock       sync.RWMutex //*读写锁
	drop        bool         // *是否删除
	channel     *Channel     //*指向会话双向链表
}

func NewRoom(roomId int) *Room {
	room := new(Room)
	room.Id = roomId
	room.drop = false
	room.channel = nil
	room.OnlineCount = 0
	return room
}

// *将channel放入房间
func (r *Room) Put(ch *Channel) (err error) {
	r.rLock.Lock()
	defer r.rLock.Unlock()
	if !r.drop {
		if r.channel != nil {
			r.channel.Prev = ch
		}
		ch.Next = r.channel
		ch.Prev = nil
		r.channel = ch //*移动头部指针
		r.OnlineCount++
	} else {
		err = errors.New("room is drop")
	}
	return
}

// *将消息传入所有的会话中
func (r *Room) Push(msg *proto.Msg) {
	r.rLock.RLock()
	//*遍历所有的会话
	for ch := r.channel; ch != nil; ch = ch.Next {
		if err := ch.Push(msg); err != nil {
			logrus.Infof("push msg err:%s", err.Error())
		}
	}
	r.rLock.RUnlock()
	return
}

// *删除会话
func (r *Room) DeleteChannel(ch *Channel) bool {
	r.rLock.RLock()
	if ch.Next != nil {
		ch.Next.Prev = ch.Prev
	}
	if ch.Prev != nil {
		ch.Prev.Next = ch.Next
	} else {
		r.channel = ch.Next
	}
	r.OnlineCount--
	r.drop = false
	if r.OnlineCount <= 0 {
		r.drop = true
	}
	r.rLock.RUnlock()
	return r.drop

}
