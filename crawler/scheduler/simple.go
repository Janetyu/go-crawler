package scheduler

import "go-crawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 将 request 送进 worker chan
	// s.workerChan <- r
	// 需要变为 goroutine ，让主线程 engine 消费 worker 的操作不被卡住
	go func() { s.workerChan <- r }()
}


func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

