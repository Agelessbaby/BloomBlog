package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/errno"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *comment.BloomblogCommentActionRequest) error {
	if req.ActionType == 1 {
		var par_id, rep_id *int
		if par := int(req.GetParentId()); par == 0 {
			par_id = nil
		} else {
			par_id = &par
		}
		if rep := int(req.GetReplyId()); rep == 0 {
			rep_id = nil
		} else {
			rep_id = &rep
		}
		return db.NewComment(s.ctx, &db.Comment{
			PostID:   int(req.PostId),
			UserID:   int(req.UserId),
			Content:  *req.CommentText,
			ParentID: par_id,
			ReplyID:  rep_id,
		})
	}
	if req.ActionType == 2 {
		return db.DeleteComment(s.ctx, *req.CommentId, req.PostId, int(req.UserId))
	}
	return errno.ErrBind
}
