package engine_test

import (
	"testing"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/types"
)

func TestConcurrentEngine_simpleScheduler(t *testing.T) {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

func TestConcurrentEngine_queuedScheduler(t *testing.T) {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
