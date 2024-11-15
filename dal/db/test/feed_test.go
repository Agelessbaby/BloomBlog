package test

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"testing"
)

func TestFeed(t *testing.T) {
	user := db.User{
		UserName:       "asd",
		Password:       "asd",
		FollowingCount: 0,
		FollowerCount:  0,
	}
	fmt.Println(user)
}
