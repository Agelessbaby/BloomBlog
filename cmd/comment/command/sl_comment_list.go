package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
)

type SlCommentListService struct {
	ctx context.Context
}

func NewSlCommentListService(ctx context.Context) *SlCommentListService {
	return &SlCommentListService{ctx: ctx}
}

func (s *SlCommentListService) SlCommentListService(req *comment.Bloomblog_SlCommentListRequest, fromID int64) ([]*comment.Comment, error) {
	comments, err := db.GetSecondLevelComments(s.ctx, req.ParentId)
	if err != nil {
		return nil, err
	}
	cs, err := pack.Comments(s.ctx, comments, fromID)
	if err != nil {
		return nil, err
	}
	return cs, nil
}
