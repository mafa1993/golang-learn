package main

import (
	"crawler_dis/persist"
	rpcc "crawler_dis/rpc"

	"gopkg.in/olivere/elastic.v5"
)

func main() {

	err:=serverRpc(":1234")
	if err != nil {
		panic(err)
	}

}

// 服务端启动
func serverRpc(host string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	// 开启服务
	return rpcc.RpcServer(persist.ItemSaveServer{
		Client: client,
	}, host)
}
