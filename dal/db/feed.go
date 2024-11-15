package db

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Tag []string

// Post Gorm Data Structures
type Post struct {
	gorm.Model
	Author        User   `gorm:"foreignkey:AuthorID"`
	AuthorID      int    `gorm:"index:idx_authorid;not null"`
	ImageUrls     Tag    `gorm:"type:json"`
	CoverUrl      string `gorm:"type:varchar(255)"`
	TextContent   string `gorm:"type:text"`
	FavoriteCount int    `gorm:"default:0"`
	CommentCount  int    `gorm:"default:0"`
	Title         string `gorm:"type:varchar(50);not null"`
}

func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (Post) TableName() string {
	return "post"
}

// MGetPosts multiple get list of posts info
func MGetPosts(ctx context.Context, limit int, latestTime *int64) ([]*Post, error) {
	posts := make([]*Post, 0)

	if latestTime == nil || *latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = &cur_time
	}
	conn := DB.WithContext(ctx)

	if err := conn.Limit(limit).Order("update_time desc").Find(&posts, "update_time < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
