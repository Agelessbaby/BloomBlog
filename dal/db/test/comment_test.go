package test

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"testing"
)

func TestNewComment(t *testing.T) {
	var parent_id int
	parent_id = 2
	var reply_id int
	reply_id = 2
	err := db.NewComment(context.TODO(), &db.Comment{
		PostID:   1,
		UserID:   1,
		Content:  "okokoask",
		ParentID: &parent_id,
		ReplyID:  &reply_id,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteComment(t *testing.T) {
	err := db.DeleteComment(context.TODO(), 3, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetComments(t *testing.T) {
	comments, err := db.GetPostComments(context.TODO(), 1)
	if err != nil {
		t.Error(err)
	}
	for _, comment := range comments {
		fmt.Println(comment.Content)
	}
}

func TestGetSecondLevelComments(t *testing.T) {
	comments, err := db.GetSecondLevelComments(context.TODO(), 2)
	if err != nil {
		t.Error(err)
	}
	for _, comment := range comments {
		fmt.Println(comment.Content, comment.ID)
	}
}
