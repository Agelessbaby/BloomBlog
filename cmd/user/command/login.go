package command

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *user.BloomBlogUserRegisterRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errors.New("user not found")
	}
	loginUser := users[0]
	match, err := AuthPassword(req.Password, loginUser.Password)
	if err != nil {
		return 0, err
	}
	if !match {
		return 0, errors.New("password incorrect")
	}
	return int64(loginUser.ID), nil
}

func AuthPassword(reqPassword, userPassword string) (bool, error) {
	if userPassword != reqPassword {
		return false, nil
	}
	return true, nil
}
