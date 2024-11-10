package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/errno"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.BloomblogRelationActionRequest) error {
	//follow
	if req.ActionType == 1 {
		return db.NewRelation(s.ctx, req.UserId, req.ToUserId)
	}
	//unfollow
	if req.ActionType == 2 {
		return db.DisRelation(s.ctx, req.UserId, req.ToUserId)
	}

	return errno.ErrBind
}
