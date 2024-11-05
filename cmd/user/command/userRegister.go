package command

import (
	"context"
	_ "crypto/md5"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// TODO use crypto algorithm for security
func (s *CreateUserService) CreateUser(req *user.BloomBlogUserRegisterRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errors.New("This username already exists")
	}
	return db.CreateUser(s.ctx, []*db.User{
		&db.User{
			UserName: req.Username,
			Password: req.Password,
		},
	})
}
