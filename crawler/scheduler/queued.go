package scheduler

import (
	"go-crawler/crawler/types"
)


// Request队列和Worker队列
type QueuedScheduler struct {
	requestChan chan types.Request
	workerChan chan chan types.Request // 每一个 worker 对外接口的是 chan Request，每一个 worker 都是 chan
}

func (s *QueuedScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan types.Request) {
	s.workerChan <- w
}


func (s *QueuedScheduler) Submit(r types.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan types.Request)
	s.workerChan = make(chan chan types.Request)
	go func() {
		var requestQ []types.Request
		var workerQ []chan types.Request
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQ) > 0 && len(workerQ) >0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- s.requestChan :
				requestQ = append(requestQ, r)
			case w := <- s.workerChan :
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}


