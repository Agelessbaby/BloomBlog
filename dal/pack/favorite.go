package pack

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/db"
)

// FavoritePosts pack favoriteVideos info.
func FavoritePosts(ctx context.Context, vs []db.Post, uid *int64) ([]*feed.Post, error) {
	posts := make([]*db.Post, 0)
	for _, p := range vs {
		posts = append(posts, &p)
	}

	packPosts, err := Posts(ctx, posts, uid)
	if err != nil {
		return nil, err
	}

	return packPosts, nil
}
