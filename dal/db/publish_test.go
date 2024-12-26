package db

import (
	"context"
	"testing"
)

func TestInsertIntoTimeline(t *testing.T) {
	post := &Post{
		AuthorID: 1,
		CoverUrl: "http://example.com/cover.jpg",
	}
	post.ID = 1
	InsertIntoTimeline(context.Background(), post)
}
