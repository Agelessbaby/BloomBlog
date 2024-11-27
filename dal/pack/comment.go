package pack

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"gorm.io/gorm"
)

func Comments(ctx context.Context, comments []*db.Comment, fromID int64) ([]*comment.Comment, error) {
	Cs := make([]*comment.Comment, 0)
	for _, c := range comments {
		user, err := db.GetUserByID(ctx, int64(c.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		packUser, err := User(ctx, user, fromID)
		if err != nil {
			return nil, err
		}
		Cs = append(Cs, &comment.Comment{
			Id:         int64(c.ID),
			User:       packUser,
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format("01-02"),
		})
	}
	return Cs, nil
}
