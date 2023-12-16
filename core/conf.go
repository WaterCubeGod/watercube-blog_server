package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"
)

// InitConf 读取yaml的配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	errJ := yaml.Unmarshal(yamlConf, c)
	if errJ != nil {
		log.Fatalf("config Init Unmarshal error: %v", errJ)
	}
	log.Println("config yamlFile load Init success.")
	global.CONFIG = c
}
