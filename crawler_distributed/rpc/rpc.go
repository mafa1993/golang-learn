package rpcc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 
func RpcServer(service interface{}, host string) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err %s", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

// 创建jsonrpc 的 client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	client := jsonrpc.NewClient(conn)
	return client, nil
}
