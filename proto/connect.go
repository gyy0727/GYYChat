package proto

type Msg struct {
	Ver       int    `json:"ver"`       //*版本
	Operation int    `json:"operation"` //*操作
	SeqId     string `json:"seq"`       //*序列号
	Body      string `json:"body"`      //*消息体
}

// *user
type PushMsgRequest struct {
	UserId int
	Msg    Msg
}

// *消息
type PushRoomMsgRequest struct {
	RoomId int
	Msg    Msg
}

// *聊天室人数
type PushRoomCountRequest struct {
	RoomId int
	Count  int
}
