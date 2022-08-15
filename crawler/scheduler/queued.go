package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request      // 接收request
	workerChan  chan chan engine.Request // 用来管理worker，应该传递的为worker，worker暴露的在外的 为chan engine.Request  每个worker建立自己的channel
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (sim *QueuedScheduler) ConfigChan() chan engine.Request {
	return make(chan engine.Request)
}

// 告知scheduler Woker已经准备好
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	// 将可以使用的worker放入worker队列
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	// 整体调度，
	go func() {
		var (
			requestQ []engine.Request      // 保存rquest的队列
			workerQ  []chan engine.Request // 保存worker的队列
		)
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 如果request队列和worker队列都有数据，那么就说明可以往worker里发送数据执行
			if len(requestQ) > 0 && len(workerQ) > 0 {
				// 为了避免阻塞，这里不直接进行 activeWorker <- activeReuqest  在select中做
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan: // 取出任务
				// 把收到的request发送给worker，
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 取出可用的worker
				// 把request发送给w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				// 发送成功 在队列中去除
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]

			}

		}
	}()
}
