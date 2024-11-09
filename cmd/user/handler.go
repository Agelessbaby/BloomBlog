package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/user/command"
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	// if username or password is null
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errors.New("username or password is empty"))
		return resp, nil
	}

	if err := command.NewCreateUserService(ctx).CreateUser(req); err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return resp, nil
	}
	return pack.BuildUserRegisterResp(errno.Success), nil
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	// if username or password is null
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserLoginResponse(errors.New("username or password is empty"))
		return resp, nil
	}
	userId, err := command.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildUserLoginResponse(err)
		return resp, nil
	}
	token, err := jwt.GenJWT(userId)
	if err != nil {
		resp = pack.BuildUserLoginResponse(errors.New("creating signature failed"))
		return resp, nil
	}
	resp = pack.BuildUserLoginResponse(errno.Success)
	resp.UserId = userId
	resp.Token = token
	return resp, nil
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.BloomBlogUserRequest) (resp *user.BloomBlogUserResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildUserUserResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	var from_id int64
	fromid_val := payload.UserDefined["userid"]
	switch v := fromid_val.(type) {
	case int64:
		from_id = v
	case float64:
		from_id = int64(v)
	}
	if from_id <= 0 {
		resp = pack.BuildUserUserResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	if req.UserId <= 0 {
		resp = pack.BuildUserUserResp(errno.ErrBind)
		return resp, nil
	}
	user, err := command.NewGetUserService(ctx).GetUser(req, from_id)
	if err != nil {
		resp = pack.BuildUserUserResp(err)
		return resp, nil
	}

	resp = pack.BuildUserUserResp(errno.Success)
	resp.User = user
	fmt.Println(resp)
	return resp, nil
}
