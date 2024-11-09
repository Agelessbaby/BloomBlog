package handlers

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func Login(c context.Context, ctx *app.RequestContext) {
	var loginParam UserRegisterParam
	if err := ctx.Bind(&loginParam); err != nil {
		SendResponse(ctx, pack.BuildUserLoginResponse(errno.ErrBind))
		return
	}
	if len(loginParam.UserName) == 0 || len(loginParam.PassWord) == 0 {
		SendResponse(ctx, pack.BuildUserLoginResponse(errno.ErrBind))
		return
	}
	resp, err := rpc.Login(c, &user.BloomBlogUserRegisterRequest{
		Username: loginParam.UserName,
		Password: loginParam.PassWord,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildUserLoginResponse(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

func Register(c context.Context, ctx *app.RequestContext) {
	var registerParam UserRegisterParam
	if err := ctx.Bind(&registerParam); err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(c, &user.BloomBlogUserRegisterRequest{
		Username: registerParam.UserName,
		Password: registerParam.PassWord,
	})

	if err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

func GetUserById(c context.Context, ctx *app.RequestContext) {
	var userVar UserParam
	if err := ctx.Bind(&userVar); err != nil {
		SendResponse(ctx, pack.BuildUserUserResp(errno.ErrBind))
	}
	if len(userVar.Token) == 0 || userVar.UserId < 0 {
		SendResponse(ctx, pack.BuildUserUserResp(errno.ErrBind))
		return
	}
	resp, err := rpc.GetUserById(c, &user.BloomBlogUserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildUserUserResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
