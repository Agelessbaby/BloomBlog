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
