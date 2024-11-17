package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/command"
	publish "github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.BloomblogPublishActionRequest) (resp *publish.BloomblogPublishActionResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildPublishResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	user_id := jwt.GetUserIdFromPayload(payload)
	if user_id <= 0 {
		resp = pack.BuildPublishResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if len(req.Title) == 0 || len(req.Images) == 0 || len(req.Cover) == 0 {
		resp = pack.BuildPublishResp(errno.ErrBind)
		return resp, nil
	}

	err = command.NewPublishActionService(ctx).PublishAction(req, user_id)
	if err != nil {
		resp = pack.BuildPublishResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.BloomblogPublishListRequest) (resp *publish.BloomblogPublishListResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildPublishListResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	userid := jwt.GetUserIdFromPayload(payload)
	if userid <= 0 {
		resp = pack.BuildPublishListResp(errno.ErrTokenInvalid)
	}
	if req.UserId == 0 {
		req.UserId = userid
	}
	posts, err := command.NewPublishListService(ctx).PublishList(req, &userid)
	if err != nil {
		resp = pack.BuildPublishListResp(errno.ConvertErr(err))
		return resp, nil
	}
	resp = pack.BuildPublishListResp(errno.Success)
	resp.PostList = posts
	return resp, nil
}
