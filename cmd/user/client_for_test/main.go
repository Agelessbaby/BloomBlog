package main

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := usersrv.NewClient("github.com/Agelessbaby/BloomBlog.user", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	resp, err := client.Login(context.Background(), &user.BloomBlogUserRegisterRequest{
		Username: "Tom",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
