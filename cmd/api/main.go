package main

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/util/config"
)

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

	registerGroup(h)
	h.Spin()
}
