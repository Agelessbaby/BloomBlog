package rpc

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	dns "github.com/kitex-contrib/resolver-dns"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var userClient usersrv.Client

func GetEndPoint(config *viper.Viper) string {
	ServiceName := config.GetString("Server.Name")
	PortNum := config.GetInt("Server.Port")
	Port := strconv.Itoa(PortNum)
	return ServiceName + ":" + Port
}

func initUserRpc(config *viper.Viper) {
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

	c, err := usersrv.NewClient(
		ServiceName,
		//TODO Add middleware
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(tracing.NewClientSuite()),        // tracer
		client.WithResolver(dns.NewDNSResolver()), // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// using rpc client to register,return the resp of rpc server
func Register(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	resp, err = userClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	resp, err = userClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

func GetUserById(ctx context.Context, req *user.BloomBlogUserRequest) (resp *user.BloomBlogUserResponse, err error) {
	resp, err = userClient.GetUserById(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
