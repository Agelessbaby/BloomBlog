package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

func (s *FavoriteListService) FavoriteList(req *favorite.BloomblogFavoriteListRequest, from_id int64) ([]*feed.Post, error) {
	favoritePosts, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	posts, err := pack.FavoritePosts(s.ctx, favoritePosts, &from_id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
