package handlers

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// CommentAction handles comment actions such as adding, editing, or deleting a comment.
//
//	@Summary		Perform actions on comments
//	@Description	This endpoint allows users to perform actions on comments, including adding a new comment, replying to a comment, or deleting a comment.
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			comment	body		CommentActionParam						true	"Comment action parameters"
//	@Success		200		{object}	comment.BloomblogCommentActionResponse	"Response with the result of the comment action"
//	@Failure		400		{object}	comment.BloomblogCommentActionResponse	"Bad request or validation error"
//	@Failure		401		{object}	comment.BloomblogCommentActionResponse	"Unauthorized"
//	@Failure		500		{object}	comment.BloomblogCommentActionResponse	"Internal server error"
//
//	@Router			/bloomblog/comment/action [post]
func CommentAction(c context.Context, ctx *app.RequestContext) {
	var ActionVar CommentActionParam
	if err := ctx.Bind(&ActionVar); err != nil {
		SendResponse(ctx, pack.BuildCommentActionResp(errno.ErrBind))
		return
	}
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	ActionVar.Token = token

	resp, err := rpc.CommentAction(c, &comment.BloomblogCommentActionRequest{
		Token:       ActionVar.Token,
		PostId:      ActionVar.PostId,
		ActionType:  ActionVar.ActionType,
		CommentText: ActionVar.Content,
		CommentId:   ActionVar.CommentId,
		ParentId:    ActionVar.ParentId,
		ReplyId:     ActionVar.ReplyId,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// CommentList retrieves the list of comments for a specific post.
//
//	@Summary		Get comments for a post
//	@Description	This endpoint retrieves the list of comments associated with a specific post. Users must provide a valid token and the post ID.
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			post_id	query		int										true	"ID of the post to fetch comments for"
//	@Success		200		{object}	comment.BloomblogCommentListResponse	"Response containing the list of comments"
//	@Failure		400		{object}	comment.BloomblogCommentListResponse	"Bad request or validation error"
//	@Failure		401		{object}	comment.BloomblogCommentListResponse	"Unauthorized"
//	@Failure		500		{object}	comment.BloomblogCommentListResponse	"Internal server error"
//	@Router			/bloomblog/comment/list [get]
func CommentList(c context.Context, ctx *app.RequestContext) {
	var ListVar CommentListParam
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	ListVar.Token = token
	post_id, err := (strconv.Atoi(ctx.Query("post_id")))
	if err != nil {
		SendResponse(ctx, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	if len(token) == 0 || post_id <= 0 {
		SendResponse(ctx, pack.BuildSlCommentListResp(errno.ErrBind))
		return
	}
	resp, err := rpc.CommentList(c, &comment.BloomblogCommentListRequest{
		Token:  ListVar.Token,
		PostId: int64(post_id),
	})
	if err != nil {
		SendResponse(ctx, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// SlCommentList retrieves the list of sub-comments for a specific parent comment.
//
//	@Summary		Get sub-comments for a comment
//	@Description	This endpoint retrieves the list of sub-comments associated with a specific parent comment. Users must provide a valid token and the parent comment ID.
//	@Tags			comment
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			parent_id	query		int										true	"ID of the parent comment to fetch sub-comments for"
//	@Success		200			{object}	comment.Bloomblog_SlCommentListResponse	"Response containing the list of sub-comments"
//	@Failure		400			{object}	comment.Bloomblog_SlCommentListResponse	"Bad request or validation error"
//	@Failure		401			{object}	comment.Bloomblog_SlCommentListResponse	"Unauthorized"
//	@Failure		500			{object}	comment.Bloomblog_SlCommentListResponse	"Internal server error"
//	@Router			/bloomblog/comment/sl-list [get]
func SlCommentList(c context.Context, ctx *app.RequestContext) {
	var ListVar SlCommentListParam
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	ListVar.Token = token
	parent_id, err := (strconv.Atoi(ctx.Query("parent_id")))
	if err != nil {
		SendResponse(ctx, pack.BuildSlCommentListResp(errno.ConvertErr(err)))
		return
	}
	if parent_id == 0 || len(token) == 0 {
		SendResponse(ctx, pack.BuildSlCommentListResp(errno.ErrBind))
		return
	}
	resp, err := rpc.SlCommentList(c, &comment.Bloomblog_SlCommentListRequest{
		Token:    ListVar.Token,
		ParentId: int64(parent_id),
	})
	if err != nil {
		SendResponse(ctx, pack.BuildSlCommentListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
