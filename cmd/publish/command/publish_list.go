package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
)

type PublishListService struct {
	ctx context.Context
}

func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

func (s *PublishListService) PublishList(req *publish.BloomblogPublishListRequest, fromid *int64) (posts []*feed.Post, err error) {
	db_posts, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	posts, err = pack.Posts(s.ctx, db_posts, fromid)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
