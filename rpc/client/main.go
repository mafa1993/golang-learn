package main

import (
	"fmt"
	rpcc "learn/rpc/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1100")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcc.Args{A: 11, B: 2}, &result)

	fmt.Printf("%s,%f", err, result)
}
