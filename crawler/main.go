package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler/scheduler"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}