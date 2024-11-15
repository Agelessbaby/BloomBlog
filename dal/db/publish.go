package db

import (
	"context"
	"gorm.io/gorm"
)

// CreateVideo creates a new video
func CreatePost(ctx context.Context, post *Post) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(post).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, authorId int64) ([]*Post, error) {
	var pubList []*Post
	err := DB.WithContext(ctx).Model(&Post{}).Where(&Post{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}
