package main

import (
	"flag"
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
)

// loglevel could be "info", "debug", "fatal", "error"
var (
	loglevel = flag.String("loglevel", "info", "log level")
)

// ./output/bin/user -loglevel=debug
func main() {
	klog.SetLevel(env.Loglevelmap[*loglevel])

	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
