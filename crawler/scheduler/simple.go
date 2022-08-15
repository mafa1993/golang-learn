package scheduler

import "crawler/engine"

// 定义scheduler的一种实现，可以定义多种, 在使用时看是定义成那个类型
type SimpleScheduler struct {
	in chan engine.Request
}

func (sim *SimpleScheduler) Submit(r engine.Request) {
	go func() { // 防止没有足够的woker接收时产生阻塞
		sim.in <- r
	}()
}

func (sim *SimpleScheduler) ConfigChan() chan engine.Request {
	return sim.in
}

func (sim *SimpleScheduler) Run() {
	sim.in = make(chan engine.Request)
}

func (sim *SimpleScheduler) WorkerReady() {}
