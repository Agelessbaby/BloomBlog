package main

import (
	"flag"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/command"
	comment "github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment/commentsrv"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/Agelessbaby/BloomBlog/util/middleware"
	"github.com/Agelessbaby/BloomBlog/util/mq"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/monitor-prometheus"
	"log"
	"net"
)

var (
	loglevel       = flag.String("loglevel", "info", "log level")
	relationConfig = config.CreateConfig("commentConfig")
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

	mq.InitMqConn()
	go func() { mq.SubscribeByKey(mq.Conn, mq.Exg, command.CommentInMq, "bloomblog-comment") }()
	go func() { mq.SubscribeByKey(mq.Conn, mq.Exg, command.DelCommentInMq, "bloomblog-delcomment") }()
	svr := comment.NewServer(new(CommentSrvImpl),
		server.WithServiceAddr(addr), // address
		//server.WithMiddleware(middleware.CommonMiddleware),                 // middleware
		//server.WithMiddleware(middleware.ServerMiddleware),                 // middleware
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(), // Multiplex
		server.WithMiddleware(middleware.TimerMW),
		server.WithMiddleware(middleware.RecordArgMW),
		server.WithTracer(prometheus.NewServerTracer(":9091", "/bloomblog-metrics")),
		//TODO add tracing
		//server.WithSuite(tracing.NewServerSuite()), // trace
		// Please keep the same as provider.WithServiceName
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}))

	if err := svr.Run(); err != nil {
		klog.Fatalf("%s stopped with error:", ServiceName, err)
	}
}
