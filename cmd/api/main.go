package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/api/handlers"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	hzconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	hz2config "github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	etcd "github.com/hertz-contrib/registry/etcd"
	"time"
)

type Http2 struct {
	Enable           bool          `json:"Enable" yaml:"Enable"`
	DisableKeepalive bool          `json:"DisableKeepalive" yaml:"DisableKeepalive"`
	ReadTimeout      time.Duration `json:"ReadTimeout" yaml:"ReadTimeout"`
}
type HertzCfg struct {
	UseNetpoll bool  `json:"UseNetpoll" yaml:"UseNetpoll"`
	Http2      Http2 `json:"Http2" yaml:"Http2"`
	Tls        Tls   `json:"Tls" yaml:"Tls"`
}

type Tls struct {
	Enable bool `json:"Enable" yaml:"Enable"`
	Cfg    tls.Config
	Cert   string `json:"CertFile" yaml:"CertFile"`
	Key    string `json:"KeyFile" yaml:"KeyFile"`
	ALPN   bool   `json:"ALPN" yaml:"ALPN"`
}

var (
	apiConfig   = config.CreateConfig("apiConfig")
	ServiceName = apiConfig.GetString("Server.Name")
	hertzCfg    HertzCfg
	ServiceAddr = fmt.Sprintf("%s:%d", apiConfig.GetString("Server.Address"), apiConfig.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", apiConfig.GetString("Etcd.Address"), apiConfig.GetInt("Etcd.Port"))
)

// encode the setting into bytes and then decode into struct
func InitHertzCfg() {
	if err := apiConfig.Sub("Hertz").Unmarshal(&hertzCfg); err != nil {
		hlog.Fatalf("Error unmarshalling Hertz config: %v", err)
	}
}

func InitHertz() *server.Hertz {
	InitHertzCfg()

	opts := []hzconfig.Option{server.WithHostPorts(ServiceAddr)}

	// 服务注册
	if apiConfig.GetBool("Etcd.Enable") {
		r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
		if err != nil {
			hlog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r, &registry.Info{
			ServiceName: ServiceName,
			Addr:        utils.NewNetAddr("tcp", ServiceAddr),
			Weight:      10,
			Tags:        nil,
		}))
	}

	//TODO add tracing
	//tracer, tracerCfg := hertztracing.NewServerTracer()
	//opts = append(opts, tracer)

	// net lib
	hertzNet := standard.NewTransporter
	if hertzCfg.UseNetpoll {
		hertzNet = netpoll.NewTransporter
	}
	opts = append(opts, server.WithTransport(hertzNet))

	// TLS & Http2
	tlsEnable := hertzCfg.Tls.Enable
	h2Enable := hertzCfg.Http2.Enable
	hertzCfg.Tls.Cfg = tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	if tlsEnable {
		cert, err := tls.LoadX509KeyPair(hertzCfg.Tls.Cert, hertzCfg.Tls.Key)
		if err != nil {
			hlog.Error(err)
		}
		hertzCfg.Tls.Cfg.Certificates = append(hertzCfg.Tls.Cfg.Certificates, cert)
		opts = append(opts, server.WithTLS(&hertzCfg.Tls.Cfg))

		if alpn := hertzCfg.Tls.ALPN; alpn {
			opts = append(opts, server.WithALPN(alpn))
		}
	} else if h2Enable {
		opts = append(opts, server.WithH2C(h2Enable))
	}

	// Hertz
	h := server.Default(opts...)
	//h.Use(gzip.Gzip(gzip.DefaultCompression),
	//	hertztracing.ServerMiddleware(tracerCfg))
	//TODO add tracing
	h.Use(gzip.Gzip(gzip.DefaultCompression))

	// Protocol
	if h2Enable {
		h.AddProtocol("h2", factory.NewServerFactory(
			hz2config.WithReadTimeout(hertzCfg.Http2.ReadTimeout),
			hz2config.WithDisableKeepAlive(hertzCfg.Http2.DisableKeepalive)))
		if tlsEnable {
			hertzCfg.Tls.Cfg.NextProtos = append(hertzCfg.Tls.Cfg.NextProtos, "h2")
		}
	}

	return h
}

// Register router groups
func registerGroup(h *server.Hertz) {
	bloomblog := h.Group("/bloomblog")

	user := bloomblog.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)
	//TODO add the following groups here
}

// 初始化 API 配置
func init() {
	rpc.InitRPC()
}

func main() {

	h := InitHertz()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	registerGroup(h)
	h.Spin()
}
