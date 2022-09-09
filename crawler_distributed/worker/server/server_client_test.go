package main

import (
	"crawler_dis/config"
	"fmt"
	"rpcc"
	"testing"
	"time"
	"crawler_dis/worker"
)

func TestCrawService(t *testing.T) {
	host := ":9003"

	go func() {
		rpcc.RpcServer(host, worker.CrawlService{})
	}()
	time.Sleep(time.Second)
	client, err := rpcc.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1077868794",
		Parser: worker.SerializedParser{
			Name: "ProfileParser",
			Args: "冰之泪",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
