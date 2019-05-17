package main

import (
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/types"
	"go-crawler/crawler/persist"
	"go-crawler/crawler/jb51/parser"
)

const (
	ZhenAiWang = "http://www.zhenai.com/zhenghun"
	JiaoBenZhiJiaGolang = "https://www.jb51.net/list/list_246_1.htm"
	JiaoBenZhiJiaPython = "https://www.jb51.net/list/list_97_1.htm"

	ZhenDataIndex = "dating_profile"
	KbsDataIndex = "knowledge_base"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	itemSaver, err := persist.ItemSaver(KbsDataIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: itemSaver,
	}

	url := JiaoBenZhiJiaGolang
	e.Run(types.Request{
		Url:        url,
		ParserFunc: parser.ParseActicleList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

