package handlers

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// FavoriteAction handles the favorite action for a post.
//
//	@Summary		Perform favorite action on a post
//	@Description	This endpoint allows a user to favorite or unfavorite a post. The user must provide a valid JWT token in the Authorization header.
//	@Tags			favorite
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			FavoriteAction	body		FavoriteActionParam							true	"Favorite Action Parameters"
//	@Success		200				{object}	favorite.BloomblogFavoriteActionResponse	"Action performed successfully"
//	@Failure		400				{object}	favorite.BloomblogFavoriteActionResponse	"Invalid request"
//	@Failure		401				{object}	favorite.BloomblogFavoriteActionResponse	"Unauthorized, invalid token"
//	@Failure		500				{object}	favorite.BloomblogFavoriteActionResponse	"Internal server error"
//	@Router			/bloomblog/favorite/action [post]
func FavoriteAction(c context.Context, ctx *app.RequestContext) {
	var Param FavoriteActionParam
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	if len(token) == 0 {
		SendResponse(ctx, pack.BuildFavoriteActionResp(errno.ErrTokenInvalid))
		return
	}
	if ctx.Bind(&Param) != nil {
		SendResponse(ctx, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	Param.Token = token
	resp, err := rpc.Action(c, &favorite.BloomblogFavoriteActionRequest{
		UserId:     0,
		Token:      Param.Token,
		PostId:     Param.PostId,
		ActionType: Param.ActionType,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// FavoriteList  handles the list action for a post.
//
//	@Summary		retrieve the favorite list of a user
//	@Description	retrieve the favorite list of a user
//	@Tags			favorite
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			user_id	query		int64									true	"userid"
//	@Success		200		{object}	favorite.BloomblogFavoriteListResponse	"Action performed successfully"
//	@Failure		400		{object}	favorite.BloomblogFavoriteListResponse	"Invalid request"
//	@Failure		401		{object}	favorite.BloomblogFavoriteListResponse	"Unauthorized, invalid token"
//	@Failure		500		{object}	favorite.BloomblogFavoriteListResponse	"Internal server error"
//	@Router			/bloomblog/favorite/list [get]
func FavoriteList(c context.Context, ctx *app.RequestContext) {
	var Param FavoriteListParam
	uid, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		SendResponse(ctx, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	if uid <= 0 {
		SendResponse(ctx, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	Param.UserId = int64(uid)
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	if len(token) == 0 {
		SendResponse(ctx, pack.BuildFavoriteListResp(errno.ErrTokenInvalid))
		return
	}
	Param.Token = token
	resp, err := rpc.List(c, &favorite.BloomblogFavoriteListRequest{
		UserId: Param.UserId,
		Token:  Param.Token,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
