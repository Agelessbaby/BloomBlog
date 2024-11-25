package handlers

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	_ "github.com/Agelessbaby/BloomBlog/docs"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// Login handles user login
//
//	@Summary		User Login
//	@Description	Authenticate user with username and password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			loginParam	body		UserRegisterParam	true	"User login data"
//	@Success		200			{object}	user.BloomBlogUserRegisterResponse
//	@Failure		400			{object}	errno.ErrNo
//	@Router			/bloomblog/user/login [post]
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

// Register handles user registration
//
//	@Summary		User Registration
//	@Description	Register a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			registerParam	body		UserRegisterParam	true	"User registration data"
//	@Success		200				{object}	user.BloomBlogUserRegisterResponse
//	@Failure		400				{object}	errno.ErrNo
//	@Router			/bloomblog/user/register [post]
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
	fmt.Println(err)
	if err != nil {
		SendResponse(ctx, pack.BuildUserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}

// GetUserById retrieves a user by ID
//
//	@Summary		Get User by ID
//	@Description	Get user information by ID and token
//	@Tags			User
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID"
//	@Success		200		{object}	user.BloomBlogUserResponse
//	@Failure		400		{object}	errno.ErrNo
//	@Router			/bloomblog/user/getuserbyid [GET]
func GetUserById(c context.Context, ctx *app.RequestContext) {
	var userVar UserParam
	uid, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		fmt.Println(err)
		SendResponse(ctx, pack.BuildUserUserResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(uid)
	token := string(ctx.GetHeader("Authorization"))
	token = jwt.TrimPrefix(token)
	userVar.Token = token
	fmt.Println(userVar)
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
