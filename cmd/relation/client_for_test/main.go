package main

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation/relationsrv"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := relationsrv.NewClient("github.com/Agelessbaby/BloomBlog.relation", client.WithHostPorts("0.0.0.0:8087"))
	if err != nil {
		panic(err)
	}
	followresp, err := client.RelationAction(context.TODO(), &relation.BloomblogRelationActionRequest{
		UserId:     0,
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzEyMTA0MzMsIm5iZiI6MCwiZXhwIjoxNzMxOTAxNjMzLCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.b_9UvMPofEBUOiWXeXEmG-JVG4zwURbEE_FyX51XkYY",
		ToUserId:   1,
		ActionType: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(followresp)
}
