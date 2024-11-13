package rpc

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation/relationsrv"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"time"
)

var relationClient relationsrv.Client

func initRelationRpc(config *viper.Viper) {
	EtcdAddress := fmt.Sprintf("%s:%d", config.GetString("Etcd.Address"), config.GetInt("Etcd.Port"))
	// 服务发现
	r, err := etcd.NewEtcdResolver([]string{EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := config.GetString("Server.Name")

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

	c, err := relationsrv.NewClient(
		ServiceName,
		//TODO Add middleware
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(r), // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.BloomblogRelationActionRequest) (resp *relation.BloomblogRelationActionResponse, err error) {
	resp, err = relationClient.RelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
