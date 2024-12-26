package db

import (
	"context"
	"encoding/json"
	"fmt"
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

type Timeline struct {
	PID      uint   `json:"pid"`
	UID      int    `json:"uid"`
	CoverUrl string `json:"coverUrl"`
}

// InsertIntoTimeline
func InsertIntoTimeline(ctx context.Context, post *Post) error {
	tl := &Timeline{
		PID:      post.ID,
		UID:      post.AuthorID,
		CoverUrl: post.CoverUrl,
	}

	timelineJSON, err := json.Marshal(tl)
	if err != nil {
		return fmt.Errorf("failed to serialize timeline: %w", err)
	}

	// 使用用户 ID 作为时间线的 Key，将 Timeline 数据存储到 Redis 的列表中
	key := fmt.Sprintf("%d", post.ID)

	// 将序列化的 Timeline 插入到列表中
	if err := Redis_client.LPush(ctx, key, timelineJSON).Err(); err != nil {
		return fmt.Errorf("failed to insert into timeline: %w", err)
	}

	return nil
}
