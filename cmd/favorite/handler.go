package main

import (
	"context"
	favorite "github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.BloomblogFavoriteActionRequest) (resp *favorite.BloomblogFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.BloomblogFavoriteListRequest) (resp *favorite.BloomblogFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
