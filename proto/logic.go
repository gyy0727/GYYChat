package proto

// *登录请求
type LoginRequest struct {
	Name     string
	Password string
}

// *登录响应
type LoginResponse struct {
	Code      int
	AuthToken string
}

// *获取用户消息请求
type GetUserInfoRequest struct {
	UserId int
}

// *获取用户消息的响应
type GetUserInfoResponse struct {
	Code     int
	UserId   int
	UserName string
}

// *注册请求
type RegisterRequest struct {
	Name     string
	Password string
}

// *注册响应
type RegisterReply struct {
	Code      int
	AuthToken string
}

// *登出请求
type LogoutRequest struct {
	AuthToken string
}

// *登出响应
type LogoutResponse struct {
	Code int
}

// *检查token是否有效请求
type CheckAuthRequest struct {
	AuthToken string
}

// * 检查token是否有效响应
type CheckAuthResponse struct {
	Code     int
	UserId   int
	UserName string
}

// *连接请求
type ConnectRequest struct {
	AuthToken string `json:"authToken"`
	RoomId    int    `json:"roomId"`
	ServerId  string `json:"serverId"`
}

// * 连接响应
type ConnectReply struct {
	UserId int
}

// *断开连接请求
type DisConnectRequest struct {
	RoomId int
	UserId int
}

// *断开连接响应
type DisConnectReply struct {
	Has bool
}

// *发送消息请求
type Send struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	FromUserId   int    `json:"fromUserId"`
	FromUserName string `json:"fromUserName"`
	ToUserId     int    `json:"toUserId"`
	ToUserName   string `json:"toUserName"`
	RoomId       int    `json:"roomId"`
	Op           int    `json:"op"`
	CreateTime   string `json:"createTime"`
}

// *tcp方式发送
type SendTcp struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	FromUserId   int    `json:"fromUserId"`
	FromUserName string `json:"fromUserName"`
	ToUserId     int    `json:"toUserId"`
	ToUserName   string `json:"toUserName"`
	RoomId       int    `json:"roomId"`
	Op           int    `json:"op"`
	CreateTime   string `json:"createTime"`
	AuthToken    string `json:"authToken"` //仅tcp时使用，发送msg时带上
}
