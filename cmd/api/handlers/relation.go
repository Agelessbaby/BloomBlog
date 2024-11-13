package handlers

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// RelationAction handles user relation actions (e.g., follow, unfollow).
// @Summary Perform a relation action
// @Description Handles actions like follow, unfollow, or other user relation actions.
// @Tags Relation
// @Accept json
// @Produce json
// @Param token query string true "User authentication token"
// @Param to_user_id query string true "The ID of the user to perform the action on"
// @Param action_type query string true "The type of action to perform" Enums(1, 2) Example(1)
// @Success 200 {object} relation.BloomblogRelationActionResponse "Action completed successfully"
// @Failure 400 {object} relation.BloomblogRelationActionResponse "Invalid input parameters"
// @Failure 500 {object} relation.BloomblogRelationActionResponse "Internal server error"
// @Router /bloomblog/relation/action [post]
func RelationAction(c context.Context, ctx *app.RequestContext) {
	var actionVar RelationActionParam
	token := ctx.Query("token")
	to_user_id := ctx.Query("to_user_id")
	action_type := ctx.Query("action_type")
	tid, err := strconv.Atoi(to_user_id)
	if err != nil {
		SendResponse(ctx, pack.BuildRelationActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		SendResponse(ctx, pack.BuildRelationActionResp(errno.ErrBind))
		return
	}
	actionVar.ActionType = int32(act)
	actionVar.Token = token
	actionVar.ToUserId = int64(tid)

	rpcReq := &relation.BloomblogRelationActionRequest{
		Token:      actionVar.Token,
		ToUserId:   actionVar.ToUserId,
		ActionType: actionVar.ActionType,
	}
	resp, err := rpc.RelationAction(c, rpcReq)
	if err != nil {
		SendResponse(ctx, pack.BuildRelationActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// RelationFollowList retrieves the list of users that a specific user is following.
// @Summary Get following list
// @Description Retrieves the list of users that a specific user is following.
// @Tags Relation
// @Accept json
// @Produce json
// @Param user_id query string true "The ID of the user to retrieve the following list for"
// @Param token query string true "User authentication token"
// @Success 200 {object} relation.BloomblogRelationFollowListResponse "Successfully retrieved following list"
// @Failure 400 {object} relation.BloomblogRelationFollowListResponse "Invalid input parameters"
// @Failure 500 {object} relation.BloomblogRelationFollowListResponse "Internal server error"
// @Router /bloomblog/relation/followlist [get]
func RelationFollowList(c context.Context, ctx *app.RequestContext) {
	var userVar UserParam
	uid, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		SendResponse(ctx, pack.BuildFollowingListResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(uid)
	userVar.Token = ctx.Query("token")
	if len(userVar.Token) == 0 || userVar.UserId <= 0 {
		SendResponse(ctx, pack.BuildFollowingListResp(errno.ErrBind))
		return
	}
	rpcReq := relation.BloomblogRelationFollowListRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	}
	resp, err := rpc.RelationFollowList(c, &rpcReq)
	if err != nil {
		SendResponse(ctx, pack.BuildFollowingListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// RelationFollowerList retrieves the list of followers for a specific user.
// @Summary Get follower list
// @Description Retrieves the list of users that follow the specified user.
// @Tags Relation
// @Accept json
// @Produce json
// @Param user_id query string true "The ID of the user to retrieve the follower list for"
// @Param token query string true "User authentication token"
// @Success 200 {object} relation.BloomblogRelationFollowerListResponse "Successfully retrieved follower list"
// @Failure 400 {object} relation.BloomblogRelationFollowerListResponse "Invalid input parameters"
// @Failure 500 {object} relation.BloomblogRelationFollowerListResponse "Internal server error"
// @Router /bloomblog/relation/followerlist [get]
func RelationFollowerList(c context.Context, ctx *app.RequestContext) {
	var userVar UserParam
	uid, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		SendResponse(ctx, pack.BuildFollowerListResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(uid)
	userVar.Token = ctx.Query("token")
	if len(userVar.Token) == 0 || userVar.UserId <= 0 {
		SendResponse(ctx, pack.BuildFollowerListResp(errno.ErrBind))
		return
	}
	rpcReq := relation.BloomblogRelationFollowerListRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	}
	resp, err := rpc.RelationFollowerList(c, &rpcReq)
	if err != nil {
		SendResponse(ctx, pack.BuildFollowerListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
