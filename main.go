package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_service API文档
// @version 1.0
// @description
// @host 127.0.0.1:8080
// @BasePath /

func main() {
	global.ARTICLE_INDEX = "article"
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.LOG = core.InitLogger()
	// 连接数据库mysql和redis
	global.DB = core.InitGorm()
	global.RDB = core.ConnectRedis()
	// 连接es
	global.ES = core.InitEs()
	//命令行参数绑定
	option := flag.Parse()
	if !flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	// 初始化路由
	router := routers.InitRouter()
	addr := global.CONFIG.System.Addr()
	global.LOG.Infof("gvb_server运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.LOG.Fatalf("路由启动失败 %v", err)
		return
	}
}
