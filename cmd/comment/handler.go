package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/command"
	comment "github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, req *comment.BloomblogCommentActionRequest) (resp *comment.BloomblogCommentActionResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildCommentActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	user_id := jwt.GetUserIdFromPayload(payload)
	if user_id <= 0 || req.ActionType != 1 && req.ActionType != 2 || req.PostId <= 0 {
		resp = pack.BuildCommentActionResp(errno.ErrBind)
		return resp, nil
	}
	req.UserId = user_id
	err = command.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	resp = pack.BuildCommentActionResp(errno.Success)
	return resp, nil
}

// CommentList implements the CommentSrvImpl interface.
// retrieves the first level comment list
// 得到一级评论
func (s *CommentSrvImpl) CommentList(ctx context.Context, req *comment.BloomblogCommentListRequest) (resp *comment.BloomblogCommentListResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildCommentListResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	user_id := jwt.GetUserIdFromPayload(payload)
	if user_id <= 0 || req.PostId <= 0 {
		resp = pack.BuildCommentListResp(errno.ErrBind)
		return resp, nil
	}
	comments, err := command.NewCommentListService(ctx).CommentList(req, user_id)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}
	resp = pack.BuildCommentListResp(errno.Success)
	resp.CommentList = comments
	return resp, nil
}

// SlCommentList implements the CommentSrvImpl interface.
// retrieves the second level comment list
// 得到二级评论
func (s *CommentSrvImpl) SlCommentList(ctx context.Context, req *comment.Bloomblog_SlCommentListRequest) (resp *comment.Bloomblog_SlCommentListResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildSlCommentListResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	user_id := jwt.GetUserIdFromPayload(payload)
	if user_id <= 0 || req.ParentId <= 0 {
		resp = pack.BuildSlCommentListResp(errno.ErrBind)
		return resp, nil
	}
	comments, err := command.NewSlCommentListService(ctx).SlCommentListService(req, user_id)
	if err != nil {
		resp = pack.BuildSlCommentListResp(err)
		return resp, nil
	}
	resp = pack.BuildSlCommentListResp(errno.Success)
	resp.CommentList = comments
	return resp, nil
}
