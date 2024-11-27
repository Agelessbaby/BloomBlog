package db

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"gorm.io/gorm"
)

// `gorm:"constraint:OnDelete:CASCADE;"`
type Comment struct {
	gorm.Model
	Post     Post     `gorm:"foreignkey:PostID"`
	PostID   int      `gorm:"index:idx_postid;not null"`
	User     User     `gorm:"foreignkey:UserID"`
	UserID   int      `gorm:"index:idx_userid;not null"`
	Content  string   `gorm:"type:varchar(255);not null"`
	ParentID *int     `gorm:"index:idx_parentid"`
	Parent   *Comment `gorm:"foreignKey:ParentID;references:ID"` // 显式定义外键
	ReplyID  *int     `gorm:"index:idx_replyid"`                 // 外键指向回复的评论
	Reply    *Comment `gorm:"foreignKey:ReplyID;references:ID"`
}

func (Comment) TableName() string {
	return "comment"
}

func NewComment(ctx context.Context, comment *Comment) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := DB.Create(comment).Error
		if err != nil {
			return err
		}
		res := tx.Model(&Post{}).Where("ID = ?", comment.PostID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}
		return nil
	})
	return err
}

func DeleteComment(ctx context.Context, commentID int64, pid int64) error {
	comment := new(Comment)
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.First(comment, commentID).Error; err != nil {
			return err
		}
		res := tx.Unscoped().Delete(&comment)
		if err := res.Error; err != nil {
			return err
		}
		//TODO The affected_comment is always 1, this needs to be corrected
		var affected_comment int64
		affected_comment = res.RowsAffected
		res = tx.Model(&Post{}).Where("ID=?", pid).Update("comment_count", gorm.Expr("comment_count - ?", affected_comment))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != affected_comment {
			return errno.ErrDatabase
		}
		return nil
	})
	return err
}

// 得到帖子的一级评论
func GetPostComments(ctx context.Context, pid int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where("post_id = ? AND parent_id IS NULL", pid).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func GetSecondLevelComments(ctx context.Context, commentID int64) ([]*Comment, error) {
	var cid int
	cid = int(commentID)
	var comments []*Comment
	res := DB.WithContext(ctx).Where(Comment{ParentID: &cid}).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
