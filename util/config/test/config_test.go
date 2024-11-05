package test

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"testing"
)

func TestGetDbConfig(t *testing.T) {
	config_viper := config.CreateConfig("DataBase")
	DBName := config_viper.GetString("MySQL.DBName")
	fmt.Println(DBName)
}
