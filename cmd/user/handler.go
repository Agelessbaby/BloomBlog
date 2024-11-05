package main

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/user/command"
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
)

// UserSrvImpl implements the last service interface defined in the IDL.
type UserSrvImpl struct{}

// Register implements the UserSrvImpl interface.
func (s *UserSrvImpl) Register(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	// TODO: Your code here...
	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errors.New("username or password is empty"))
		return resp, nil
	}
	if err := command.NewCreateUserService(ctx).CreateUser(req); err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return resp, nil
	}
	return pack.BuildUserRegisterResp(nil), nil
}

// Login implements the UserSrvImpl interface.
func (s *UserSrvImpl) Login(ctx context.Context, req *user.BloomBlogUserRegisterRequest) (resp *user.BloomBlogUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserSrvImpl interface.
func (s *UserSrvImpl) GetUserById(ctx context.Context, req *user.BloomBlogUserRequest) (resp *user.BloomBlogUserResponse, err error) {
	// TODO: Your code here...
	return
}
