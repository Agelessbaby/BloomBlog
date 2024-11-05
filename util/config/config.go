package config

import (
	"fmt"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/spf13/viper"
)

// Viper可以解析JSON、TOML、YAML、HCL、INI、ENV等格式的配置文件。甚至可以监听配置文件的变化(WatchConfig)，不需要重启程序就可以读到最新的值。
func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := env.ProjectRootPath + "config/"
	config.AddConfigPath(configPath) // 文件所在目录
	config.SetConfigName(file)       // 文件名
	config.SetConfigType("yaml")     // 文件类型
	configFile := configPath + file + ".yaml"

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件:%s", configFile)) //系统初始化阶段发生任何错误，直接结束进程
		} else {
			panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
		}
	}

	return config
}
