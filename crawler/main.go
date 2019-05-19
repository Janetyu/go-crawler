package main

import (
	"go-crawler/crawler/types"
	"go-crawler/crawler/jb51/parser"
	"go-crawler/crawler/engine"
	"go-crawler/crawler/scheduler"
	"go-crawler/crawler/persist"
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

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个 gin 路由：
	// 日志与恢复中间件（无崩溃）。

	//go func() {
	//	router := gin.Default()
	//
	//	//router.GET("/sd/monitor", AllCheck)
	//	v := router.Group("/sd")
	//	{
	//		v.GET("/all", monitor.AllCheck)
	//		v.GET("/health", monitor.HealthCheck)
	//		v.GET("/disk", monitor.DiskCheck)
	//		v.GET("/cpu", monitor.CPUCheck)
	//		v.GET("/ram", monitor.RAMCheck)
	//	}
	//	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	//	//router.Run(":8989") 硬编码端口
	//	log.Info(http.ListenAndServe(":8989", router).Error())
	//}()


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

	url := JiaoBenZhiJiaPython
	e.Run(types.Request{
		Url:        url,
		ParserFunc: parser.ParseActicleList,
	})
}

