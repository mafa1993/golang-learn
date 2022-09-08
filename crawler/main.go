package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	//disp "crawler_dis/persist"
)

func main() {
	var seeds []engine.Request

	seeds = []engine.Request{
		{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.CityListParser,
		},
	}

	//engine.Run(seeds...)
	//engine.SimpleEngine{}.Run(seeds...)  // 单任务执行
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 10,
	// }
	itemchan, err := persist.ItemSave(":1234")
	//itemchan, err := disp.ItemSave(":1234")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemchan, // 相当于create Worker，会生成消费者，阻塞等待item传入
	}
	e.Run(seeds...)
}
