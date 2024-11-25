package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.BloomblogFavoriteActionRequest) error {
	if req.ActionType == 1 {
		return db.Favorite(s.ctx, req.UserId, req.PostId)
	}
	if req.ActionType == 2 {
		return db.DisFavorite(s.ctx, req.UserId, req.PostId)
	}
	return errno.ErrBind
}
