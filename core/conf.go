package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml的配置
func InitConf() {
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

func SetYaml() error {
	byteData, err := yaml.Marshal(global.CONFIG)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.LOG.Infof("系统配置修改成功")
	return nil
}
