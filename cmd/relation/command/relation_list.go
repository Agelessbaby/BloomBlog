package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
)

type RelationListService struct {
	ctx context.Context
}

func NewRelationListService(ctx context.Context) *RelationListService {
	return &RelationListService{ctx: ctx}
}

func (s *RelationListService) FollowingList(req *relation.BloomblogRelationFollowListRequest, fromid int64) ([]*user.User, error) {
	FollowingUsers, err := db.FollowingList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowingList(s.ctx, FollowingUsers, fromid)
}

func (s *RelationListService) FollowerList(req *relation.BloomblogRelationFollowerListRequest, fromid int64) ([]*user.User, error) {
	FollowerUsers, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowerList(s.ctx, FollowerUsers, fromid)
}
