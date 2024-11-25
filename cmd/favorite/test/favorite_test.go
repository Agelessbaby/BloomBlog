package test

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite/favoritesrv"
	"github.com/cloudwego/kitex/client"
	"testing"
)

func TestFavoriteAction(t *testing.T) {
	client, err := favoritesrv.NewClient("BloomBlogFavoriteService", client.WithHostPorts("127.0.0.1:8085"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.FavoriteAction(context.TODO(), &favorite.BloomblogFavoriteActionRequest{
		UserId:     0,
		Token:      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzI1MjE5NTgsIm5iZiI6MCwiZXhwIjoxNzMzMjEzMTU4LCJ1ZCI6eyJ1c2VyaWQiOjF9fQ.4ib5moejerjGtrIB6JjYLp4SrLSmBe7dDvjjxLBCNvQ",
		PostId:     3,
		ActionType: 2,
	})
	fmt.Println(resp, err)
}

func TestFavorite(t *testing.T) {
	client, err := favoritesrv.NewClient("BloomBlogFavoriteService", client.WithHostPorts("127.0.0.1:8085"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.FavoriteList(context.TODO(), &favorite.BloomblogFavoriteListRequest{
		UserId: 0,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzI1MjE5NTgsIm5iZiI6MCwiZXhwIjoxNzMzMjEzMTU4LCJ1ZCI6eyJ1c2VyaWQiOjF9fQ.4ib5moejerjGtrIB6JjYLp4SrLSmBe7dDvjjxLBCNvQ",
	})
	fmt.Println(resp, err)
}
