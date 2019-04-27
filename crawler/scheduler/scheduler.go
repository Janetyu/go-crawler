package scheduler

import (
	"go-crawler/crawler/types"
)

type Scheduler interface {
	ReadyNotifier
	Submit(types.Request)
	WorkerChan() chan types.Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan types.Request)
}
