package main

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/hertz-contrib/cors"
	"time"
)

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

var (
	apiConfig   = config.CreateConfig("apiConfig")
	ServiceName = apiConfig.GetString("Server.Name")
	hertzCfg    HertzCfg
	ServiceAddr = fmt.Sprintf("%s:%d", apiConfig.GetString("Server.Address"), apiConfig.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", apiConfig.GetString("Etcd.Address"), apiConfig.GetInt("Etcd.Port"))
)

// init Rpc Client
func init() {
	rpc.InitRPC()
}

func main() {
	h := InitHertz()
	h.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	registerGroup(h)
	h.Spin()
}
