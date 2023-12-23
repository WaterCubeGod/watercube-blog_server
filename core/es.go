package core

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"log"
	"os"
)

func InitEs() *elastic.Client {
	url := fmt.Sprintf("http://%s:%d", global.CONFIG.Es.Ip, global.CONFIG.Es.Port)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		elastic.SetSniff(false),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, global.CONFIG.Es.LogErr, log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		global.LOG.Error("Failed to create elastic client")
	}
	return client
}
