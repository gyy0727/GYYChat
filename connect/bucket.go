package connect

import (
	"sync"
	"sync/atomic"
	"zinx/proto"
)

type Bucket struct {
	cLock         sync.RWMutex     //*读写锁
	chs           map[int]*Channel //*所有的会话
	bucketOptions BucketOptions    //*配置选项
	rooms         map[int]*Room    //*房间列表
	routines      []chan *proto.PushRoomMsgRequest
	routinesNum   uint64
	broadcast     chan []byte
}

type BucketOptions struct {
	ChannelSize   int    //*会话的数量
	RoomSize      int    //*房间的数量
	RoutineAmount uint64 //*协程的数量
	RoutineSize   int    //*每个协程的数量
}

// *获得指定id对应的房间
func (b *Bucket) Room(rid int) (room *Room) {
	b.cLock.RLock()
	room, _ = b.rooms[rid]
	b.cLock.RUnlock()
	return
}

// *将新用户对应的会话放入篮子和房间中
func (b *Bucket) Put(userId int, roomId int, ch *Channel) (err error) {
	var (
		room *Room
		ok   bool
	)
	b.cLock.Lock()
	if roomId != NoRoom {
		if room, ok = b.rooms[roomId]; !ok {
			room = NewRoom(roomId)
			b.rooms[roomId] = room
		}
		ch.Room = room
	}
	ch.userId = userId
	b.chs[userId] = ch
	b.cLock.Unlock()
	if room != nil {
		err = room.Put(ch)
	}
	return
}

func (b *Bucket) PushRoom(ch chan *proto.PushRoomMsgRequest) {
	for {
		var (
			arg  *proto.PushRoomMsgRequest
			room *Room
		)
		arg = <-ch
		if room = b.Room(arg.RoomId); room != nil {
			room.Push(&arg.Msg)
		}
	}
}

// *取出userid对应的会话
func (b *Bucket) Channel(userId int) (ch *Channel) {
	b.cLock.RLock()
	ch, _ = b.chs[userId]
	b.cLock.RUnlock()
	return
}

func (b *Bucket) BroadcastRoom(pushRoomMsgReq *proto.PushRoomMsgRequest) {
	num := atomic.AddUint64(&b.routinesNum, 1) % b.bucketOptions.RoutineAmount
	b.routines[num] <- pushRoomMsgReq
}

// *删除会话
func (b *Bucket) DeleteChannel(ch *Channel) {
	var (
		ok   bool
		room *Room
	)
	b.cLock.RLock()
	//*先判断是否存在
	if ch, ok = b.chs[ch.userId]; ok {
		room = b.chs[ch.userId].Room
		//*从哈希表删除
		delete(b.chs, ch.userId)
	}
	if room != nil && room.DeleteChannel(ch) {
		//*如果删除的是最后一个会话,代表房间已经没人了
		if room.drop == true {
			delete(b.rooms, room.Id)
		}
	}
	b.cLock.RUnlock()
}

// *新建一个篮子
func NewBucket(options BucketOptions) (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[int]*Channel, options.ChannelSize)
	b.bucketOptions = options
	b.routines = make([]chan *proto.PushRoomMsgRequest, options.RoutineAmount)
	b.rooms = make(map[int]*Room, options.RoomSize)
	for i := uint64(0); i < b.bucketOptions.RoutineAmount; i++ {
		c := make(chan *proto.PushRoomMsgRequest, options.RoutineSize)
		b.routines[i] = c
		go b.PushRoom(c)
	}
	return
}
