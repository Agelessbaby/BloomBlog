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

}

func Register(c context.Context, ctx *app.RequestContext) {
	var registerParam UserRegisterParam
	if err := ctx.Bind(&registerParam); err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ErrBind))
	}

	if len(registerParam.UserName) == 0 || len(registerParam.PassWord) == 0 {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ErrBind))
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

func GetUserById(c context.Context, ctx *app.RequestContext) {}
