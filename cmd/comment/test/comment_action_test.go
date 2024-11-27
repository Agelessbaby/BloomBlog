package test

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment/commentsrv"
	client2 "github.com/cloudwego/kitex/client"
	"testing"
)

func TestCommentAction(t *testing.T) {
	client, err := commentsrv.NewClient("BloomBlogCommentServer", client2.WithHostPorts("127.0.0.1:8086"))
	if err != nil {
		t.Error(err)
	}
	var text = "asas"
	var par_id = int64(2)
	var rep_id = int64(10)
	var comment_id = int64(20)
	resp, err := client.CommentAction(context.TODO(), &comment.BloomblogCommentActionRequest{
		Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzI2OTg5NDIsIm5iZiI6MCwiZXhwIjoxNzMzMzkwMTQyLCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.Ag-oIpcalax8YtdwjKS52zqHJxL8Mlg94HplTy62Cxc",
		PostId:      1,
		ActionType:  2,
		CommentId:   &comment_id,
		CommentText: &text,
		ParentId:    &par_id,
		ReplyId:     &rep_id,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestSlCommentAction(t *testing.T) {
	client, err := commentsrv.NewClient("BloomBlogCommentServer", client2.WithHostPorts("127.0.0.1:8086"))
	if err != nil {
		t.Error(err)
	}
	resp, err := client.SlCommentList(context.TODO(), &comment.Bloomblog_SlCommentListRequest{
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzI2OTg5NDIsIm5iZiI6MCwiZXhwIjoxNzMzMzkwMTQyLCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.Ag-oIpcalax8YtdwjKS52zqHJxL8Mlg94HplTy62Cxc",
		ParentId: 2,
	})
	if err != nil {
		t.Error(err)
	}
	if resp != nil {
		for _, c := range resp.CommentList {
			fmt.Println(c.Id, c.Content)
		}
	}
}
