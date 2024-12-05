package main

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user/usersrv"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := usersrv.NewClient("github.com/Agelessbaby/BloomBlog.user", client.WithHostPorts("0.0.0.0:8081"))
	if err != nil {
		panic(err)
	}
	resp, err := client.Login(context.TODO(), &user.BloomBlogUserRegisterRequest{
		Username: "Jeff",
		Password: "123456",
	})
	fmt.Println(resp, err)
}
