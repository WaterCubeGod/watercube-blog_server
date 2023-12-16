package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.LOG = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化路由
	router := routers.InitRouter()
	addr := global.CONFIG.System.Addr()
	global.LOG.Infof("gvb_server运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.LOG.Errorln("路由启动失败")
		return
	}
}
