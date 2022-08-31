package engine

import (
	"crawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler // 调度器
	WorkerCount int       // 进程数
	ItemChan    chan Item // 用于接收item信息，进行分发存储
}

// 定义scheduler接口
type Scheduler interface {
	Submit(Request) // 提交任务到调度器，
	ConfigChan() chan Request
	Run()
	ReadyNotifier
}

// 将ready单独出来，在只需调用ready的地方，不需要传整个scheduler，传ReadyNotifier即可
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e ConcurrentEngine) Run(seed ...Request) {

	e.Scheduler.Run()
	for _, v := range seed {
		e.Scheduler.Submit(v)
	}
	out := make(chan ParseResult) // 接收任务的返回结果
	for i := 0; i < e.WorkerCount; i++ {
		// 多调度器适配，是走队列模式的scheduler还是抢占式的scheduler去询问scheduler
		// 队列模式的，每个woker有自己的chan
		// 抢占式的 共用一个chan
		e.createWorker(e.Scheduler.ConfigChan(), out, e.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Item {
			// log.Printf("item 为%s", item)
			go func(item Item) { e.ItemChan <- item }(item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request) // 提交请求到in管道
		}
	}
}

// 这里用Readynotifier只接受Scheduler的一部分
func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s ReadyNotifier) {

	go func() {
		for { // 保证循环取数据
			// 先告诉scheduler worker空出来了，再去收数据
			s.WorkerReady(in)
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
