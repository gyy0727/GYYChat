package connect

import "zinx/proto"

// *连接和断开连接的操作
type Operator interface {
	Connect(conn *proto.ConnectRequest) (int, error)
	DisConnect(disConn *proto.DisConnectRequest) (err error)
}

type DefaultOperator struct {}


func(o* DefaultOperator)Connect(conn *proto.ConnectRequest) (uid int, err error) {
	// rpcConnect := new(RpcConnect)
	return 
}
