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
	resp, err := client.RelationFollowerList(context.TODO(), &relation.BloomblogRelationFollowerListRequest{
		UserId: 1,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzEzMTExNzAsIm5iZiI6MCwiZXhwIjoxNzMyMDAyMzcwLCJ1ZCI6eyJ1c2VyaWQiOjN9fQ.2mwVT7xI46NS_dvTZZOvcslbQ4X6z82oMWh07UQ9_tY",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
