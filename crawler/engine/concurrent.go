package engine

import (
	"log"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/types"
)

type ConcurrentEngine struct {
	Scheduler scheduler.Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine)Run(seeds ...types.Request) {
	//in := make(chan Request)
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i ++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		// 避免重复爬取 url
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		// 这里是 engine 消费 worker
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

// URL deduplicate
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

func createWorker(in chan types.Request, out chan types.ParseResult, ready scheduler.ReadyNotifier) {
	go func() {
		for {
			// 这里是 worker 消费 scheduler
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

