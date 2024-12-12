package main

import (
	"flag"
	"fmt"
	feed "github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed/feedsrv"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

var (
	loglevel       = flag.String("loglevel", "info", "log level")
	relationConfig = config.CreateConfig("feedConfig")
	ServiceName    = relationConfig.GetString("Server.Name")
	ServiceAddr    = fmt.Sprintf("%s:%d", relationConfig.GetString("Server.Address"), relationConfig.GetInt("Server.Port"))
)

func main() {
	klog.SetLevel(klog.LevelDebug)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	svr := feed.NewServer(new(FeedSrvImpl),
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
