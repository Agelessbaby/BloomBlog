package main

import (
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
)

func main() {
	klog.SetLevel(klog.LevelDebug)
	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
