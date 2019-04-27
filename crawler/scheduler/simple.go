package scheduler

import (
	"go-crawler/crawler/types"
)

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan types.Request) {
}

func (s *SimpleScheduler) Submit(r types.Request) {
	// 将 request 送进 worker chan
	// s.workerChan <- r
	// 需要变为 goroutine ，让主线程 engine 消费 worker 的操作不被卡住
	go func() { s.workerChan <- r }()
}


func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}

