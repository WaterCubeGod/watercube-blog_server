package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string // -u admin
	ES   string // -es
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "创建es索引")
	//解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止Web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return false
	}
	if option.User != "" {
		return false
	}
	if option.ES == "es" {
		return false
	}
	return true
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	if option.ES == "es" {
		createES()
		return
	}
	sys_flag.Usage()
}
