package command

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/mq"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
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
		commentParam := &db.Comment{
			PostID:   int(req.PostId),
			UserID:   int(req.UserId),
			Content:  *req.CommentText,
			ParentID: par_id,
			ReplyID:  rep_id,
		}
		return s.comment_action(commentParam)
	}
	if req.ActionType == 2 {
		delparam := &DelParam{
			CommentId: *req.CommentId,
			PostId:    req.PostId,
			UserId:    int(req.UserId),
		}
		return s.del_comment_action(delparam)
	}
	return errno.ErrBind
}

func (s *CommentActionService) comment_action(comment2 *db.Comment) error {
	bs, err := sonic.Marshal(*comment2)
	if err != nil {
		klog.Errorf("marshal comment err: %v", err)
	}
	mq.ProduceDirect(bs, "bloomblog-comment")
	return nil
}

type DelParam struct {
	CommentId int64
	PostId    int64
	UserId    int
}

func (s *CommentActionService) del_comment_action(delparam *DelParam) error {
	bs, err := sonic.Marshal(*delparam)
	if err != nil {
		klog.Errorf("marshal delcomment err: %v", err)
	}
	mq.ProduceDirect(bs, "bloomblog-delcomment")
	return nil
}

func CommentInMq(bs []byte) error {
	var commentParam db.Comment
	err := sonic.Unmarshal(bs, &commentParam)
	if err != nil {
		klog.Errorf("Unmarshal err: %v", err)
		return err
	}
	err = db.NewComment(&commentParam)
	if err != nil {
		klog.Errorf("NewComment err: %v", err)
		return err
	}
	return nil
}

func DelCommentInMq(bs []byte) error {
	var delCommentParam DelParam
	if err := sonic.Unmarshal(bs, &delCommentParam); err != nil {
		klog.Errorf("Unmarshal err: %v", err)
		return err
	}
	err := db.DeleteComment(delCommentParam.CommentId, delCommentParam.PostId, delCommentParam.UserId)
	if err != nil {
		klog.Errorf("DeleteComment err: %v", err)
		return err
	}
	return nil
}
