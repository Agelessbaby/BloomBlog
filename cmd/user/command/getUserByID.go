package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"gorm.io/gorm"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *user.BloomBlogUserRequest, fromID int64) (*user.User, error) {
	modeluser, err := db.GetUserByID(s.ctx, req.UserId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	User, err := pack.User(s.ctx, modeluser, fromID)
	if err != nil {
		return nil, err
	}
	return User, nil
}
