package config

import (
	"fmt"
	"strconv"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` //高级配置，例如 charset
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级，debug就是输出全部sql， dev、release
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",
		m.User, m.Password, m.Host, strconv.Itoa(m.Port), m.DB, m.Config)
}
