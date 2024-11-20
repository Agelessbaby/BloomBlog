package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"time"
)

const (
	LIMIT = 30
)

type GetFeedService struct {
	ctx context.Context
}

func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{ctx: ctx}
}

func (s *GetFeedService) GetFeed(req *feed.BloomblogFeedRequest, fromID int64) (ps []*feed.Post, nextTime int64, err error) {
	posts, err := db.MGetPosts(s.ctx, LIMIT, req.LatestTime)
	if err != nil {
		return ps, nextTime, err
	}
	if len(posts) == 0 {
		nextTime = time.Now().UnixMilli()
		return ps, nextTime, nil
	} else {
		nextTime = posts[len(posts)-1].UpdatedAt.UnixMilli()
	}
	if ps, err = pack.Posts(s.ctx, posts, &fromID); err != nil {
		nextTime = time.Now().UnixMilli()
		return ps, nextTime, err
	}
	return ps, nextTime, nil
}
