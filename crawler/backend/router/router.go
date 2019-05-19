package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-crawler/crawler/backend/router/middleware"
	"go-crawler/crawler/backend/handler/sd"
	"go-crawler/crawler/backend/handler/admin"
	"go-crawler/crawler/backend/handler/data"
)

// 加载 中间件，路由器，处理器等
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	g.MaxMultipartMemory = 2 << 20 // 2 MB

	g.Static("/view", "./crawler/backend/view")
	g.StaticFS("/static", http.Dir("view"))
	//g.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 中间件函数
	g.Use(gin.Recovery())     // 用于panic时恢复API服务器
	g.Use(middleware.NoCache) // 强制浏览器不使用缓存
	g.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)

	// 404 的处理器函数
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	global := g.Group("/global")
	{
		global.POST("/adminlogin", admin.AdminLogin)
	}

	admin := g.Group("/admin")
	{
		admin.GET("/data", data.DataList)
		admin.POST("/data", data.SaveOrUpdateData)
		admin.DELETE("/data", data.DeleteData)
	}

	// 健康检查处理器的路由组
	svcd := g.Group("/sd")
	{
		svcd.GET("/all", sd.AllCheck)
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

