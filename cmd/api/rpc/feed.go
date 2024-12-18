package rpc

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed/feedsrv"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	dns "github.com/kitex-contrib/resolver-dns"
	"github.com/spf13/viper"
	"time"
)

var feedClient feedsrv.Client

func initFeedRpc(config *viper.Viper) {
	ServiceName := GetEndPoint(config)

	//TODO Add tracing in future
	//p := provider.NewOpenTelemetryProvider(
	//	provider.WithServiceName(ServiceName),
	//	provider.WithExportEndpoint("localhost:4317"),
	//	provider.WithInsecure(),
	//)
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//
	//go func() {
	//	<-ch
	//	p.Shutdown(context.Background())
	//	os.Exit(0)
	//}()

	c, err := feedsrv.NewClient(
		ServiceName,
		//TODO Add middleware
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(dns.NewDNSResolver()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

func GetFeed(ctx context.Context, req *feed.BloomblogFeedRequest) (*feed.BloomblogFeedResponse, error) {
	resp, err := feedClient.GetFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
