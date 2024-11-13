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
	loginResp, err := client.Login(context.TODO(), &user.BloomBlogUserRegisterRequest{
		Username: "Jeff",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(loginResp)
	getUserResp, err := client.GetUserById(context.TODO(), &user.BloomBlogUserRequest{
		UserId: 1,
		Token:  loginResp.Token,
	})
	fmt.Println(getUserResp)
}
