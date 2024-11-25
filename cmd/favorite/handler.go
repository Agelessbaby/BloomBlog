package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/command"
	favorite "github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.BloomblogFavoriteActionRequest) (resp *favorite.BloomblogFavoriteActionResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	uid := jwt.GetUserIdFromPayload(payload)
	if uid <= 0 || req.PostId <= 0 || (req.ActionType != 1 && req.ActionType != 2) {
		resp = pack.BuildFavoriteActionResp(errno.ErrBind)
		return resp, nil
	}
	req.UserId = uid
	err = command.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.BloomblogFavoriteListRequest) (resp *favorite.BloomblogFavoriteListResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildFavoriteListResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	from_id := jwt.GetUserIdFromPayload(payload)
	if from_id <= 0 || req.UserId <= 0 {
		resp = pack.BuildFavoriteListResp(errno.ErrBind)
		return resp, nil
	}
	favoritePosts, err := command.NewFavoriteListService(ctx).FavoriteList(req, from_id)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.PostList = favoritePosts
	return resp, nil
}
