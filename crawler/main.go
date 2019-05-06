package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/zhenai/parser"
	"go-crawler/crawler/types"
	"go-crawler/crawler/persist"
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
		ItemChan: persist.ItemSaver(),
	}

	e.Run(types.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}