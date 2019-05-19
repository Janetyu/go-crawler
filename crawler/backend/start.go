package main

import (
	"net/http"
	"time"
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/gin-gonic/gin"
	"go-crawler/crawler/backend/config"
	"go-crawler/crawler/backend/model"
	"go-crawler/crawler/backend/router"
	"go-crawler/crawler/backend/router/middleware"
)

var (
	cfg = pflag.StringP("config", "c", "", "kbs_admin config file path.")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// 设置gin模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建引擎
	g := gin.New()

	// 路由加载
	router.Load(
		// gin 核心引擎
		g,
		// 中间件列表加载
		middleware.Logging(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 60 second.")
		time.Sleep(60 * time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
