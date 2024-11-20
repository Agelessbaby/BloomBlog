package test

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed/feedsrv"
	"github.com/cloudwego/kitex/client"
	"testing"
)

func TestGetFeed(t *testing.T) {
	client, err := feedsrv.NewClient("BloomBlogPublishServer", client.WithHostPorts("0.0.0.0:8083"))
	if err != nil {
		t.Fatal(err)
	}
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzIwOTYxNzgsIm5iZiI6MCwiZXhwIjoxNzMyNzg3Mzc4LCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.Jz9bFRqvet7mE48axLQK-vcDqCcjFbTxaF7y7AFO-fY"
	resp, err := client.GetFeed(context.TODO(), &feed.BloomblogFeedRequest{
		LatestTime: nil,
		Token:      &str,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
