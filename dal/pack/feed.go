package pack

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"gorm.io/gorm"
)

func Post(ctx context.Context, p *db.Post, fromID int64) (*feed.Post, error) {
	if p == nil {
		return nil, nil
	}
	user, err := db.GetUserByID(ctx, int64(p.AuthorID))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	author, err := User(ctx, user, fromID)
	if err != nil {
		return nil, err
	}
	return &feed.Post{
		Id:           int64(p.ID),
		Author:       author,
		TextContent:  p.TextContent,
		ImageUrls:    p.ImageUrls,
		CoverUrl:     p.CoverUrl,
		LikeCount:    int64(p.FavoriteCount),
		CommentCount: int64(p.CommentCount),
		Title:        p.Title,
		ModifiedAt:   p.UpdatedAt.Unix(),
	}, nil
}

func Posts(ctx context.Context, ps []*db.Post, fromID *int64) ([]*feed.Post, error) {
	Ps := make([]*feed.Post, 0)
	for _, p := range ps {
		P, err := Post(ctx, p, *fromID)
		if err != nil {
			return nil, err
		}
		if P != nil {
			flag := false
			results, err := db.GetFavoriteRelation(ctx, *fromID, int64(p.ID))
			if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				flag = false
			} else if results != nil && results.AuthorID != 0 {
				flag = true
			}
			P.IsLiked = flag
			Ps = append(Ps, P)
		}
	}
	return Ps, nil
}
