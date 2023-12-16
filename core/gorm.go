package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"time"
)

func InitGorm() *gorm.DB {
	if global.CONFIG.Mysql.Host == "" {
		global.LOG.Warnln("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.CONFIG.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		//开发环境，显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) //只打印错误的logger
	}
	//global.MySqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.LOG.Fatalf(fmt.Sprintf("[%s] mysql连接失败", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过Mysql的wait_time
	return db
}
