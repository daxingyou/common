package rpcService

import (
	"casino_common/proto/ddproto"
	"casino_common/proto/ddproto/mjproto"
)

//Greeter
var GreeterPool Pool
func GetGreeter() ddproto.GreeterClient {
	return ddproto.NewGreeterClient(GreeterPool.Get())
}

//转转红中麻将
var ZzHzPool Pool
func GetZzHz() mjproto.ZzHzClient {
	return mjproto.NewZzHzClient(ZzHzPool.Get())
}
