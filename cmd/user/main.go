package main

import (
	"flag"
	"fmt"
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
)

// loglevel could be "info", "debug", "fatal", "error"
var (
	loglevel    = flag.String("loglevel", "info", "log level")
	userConfig  = config.CreateConfig("userConfig")
	ServiceName = userConfig.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", userConfig.GetString("Server.Address"), userConfig.GetInt("Server.Port"))
)

// ./output/bin/user -loglevel=debug
func main() {
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	svr := user.NewServer(new(UserSrvImpl),
		server.WithServiceAddr(addr), // address
		//server.WithMiddleware(middleware.CommonMiddleware),                 // middleware
		//server.WithMiddleware(middleware.ServerMiddleware),                 // middleware
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(), // Multiplex
		//TODO add tracing
		//server.WithSuite(tracing.NewServerSuite()), // trace
		// Please keep the same as provider.WithServiceName
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}))

	if err := svr.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}
