package main

import (
	rpcc "learn/rpc/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 测试方法，使用telnet11000 发送数据 {"method":"DemoService.Div","params":[{"A":1,"B":2}],"id":1}
func main() {
	rpc.Register(rpcc.DemoService{})

	listener, err := net.Listen("tcp", ":1100")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err %s", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
