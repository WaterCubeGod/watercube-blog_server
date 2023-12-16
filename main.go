package main

import (
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.LOG = core.InitLogger()
	global.LOG.Warnln("------------")
	global.LOG.Infoln("------------")
	global.LOG.Errorln("------------")
	logrus.Warnln("------------")
	logrus.Infoln("------------")
	logrus.Errorln("------------")
	// 连接数据库
	global.DB = core.InitGorm()
}
