package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler // 调度器
	WorkerCount int       // 进程数
}

// 定义scheduler接口
type Scheduler interface {
	Submit(Request) // 提交任务到调度器，
	ConfigChan(chan Request)
}

func (e ConcurrentEngine) Run(seed ...Request) {
	for _, v := range seed {
		e.Scheduler.Submit(v)
	}

	in := make(chan Request) // 发送任务
	e.Scheduler.ConfigChan(in)
	out := make(chan ParseResult) // 接收任务的返回结果
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Item {
			fmt.Println(item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request) // 提交请求到in管道
		}
	}
}

func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for { // 保证循环取数据
			Request := <-in
			rlt, _ := e.Worker(Request)
			out <- rlt
		}
	}()
}

func (e ConcurrentEngine) Worker(request Request) (ParseResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("fetch 出错，msg %s", err)
		return ParseResult{}, err
	}
	rlt := request.ParserFunc(body)

	return rlt, nil
}
