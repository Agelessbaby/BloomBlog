package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/command"
	feed "github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetFeed(ctx context.Context, req *feed.BloomblogFeedRequest) (resp *feed.BloomblogFeedResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.GetToken(), env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildPostResp(errno.ErrTokenInvalid)
	}
	uid := jwt.GetUserIdFromPayload(payload)
	if uid <= 0 {
		resp = pack.BuildPostResp(errno.ErrTokenInvalid)
	}
	posts, nexTime, err := command.NewGetFeedService(ctx).GetFeed(req, uid)
	if err != nil {
		resp = pack.BuildPostResp(err)
		return resp, nil
	}
	resp = pack.BuildPostResp(errno.Success)
	resp.PostList = posts
	resp.NextTime = &nexTime
	return resp, nil
}

// GetPostById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetPostById(ctx context.Context, req *feed.PostIdRequest) (resp *feed.Post, err error) {
	// TODO: Your code here...
	return
}
