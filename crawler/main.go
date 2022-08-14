package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
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
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	e.Run(seeds...)
}
