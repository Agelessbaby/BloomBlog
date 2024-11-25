package db

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"gorm.io/gorm"
)

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, pid int64) (*Post, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	post := new(Post)

	if err := DB.WithContext(ctx).Model(&user).Association("FavoritePosts").Find(&post, pid); err != nil {
		return nil, err
	}
	return post, nil
}

// Favorite new favorite data.
func Favorite(ctx context.Context, uid int64, post_id int64) error {
	post, err := GetFavoriteRelation(ctx, uid, post_id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if post != nil && post.AuthorID != 0 {
		return errno.ErrRecordAlreadyExists
	}
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		//1. 新增点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		post := new(Post)
		if err := tx.WithContext(ctx).First(post, post_id).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(user).Association("FavoritePosts").Append(post); err != nil {
			return err
		}

		res := tx.Model(post).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
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

// DisFavorite deletes the specified favorite from the database
func DisFavorite(ctx context.Context, uid int64, pid int64) error {
	post, err := GetFavoriteRelation(ctx, uid, pid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	} else if post == nil || errors.Is(err, gorm.ErrRecordNotFound) || post.ID <= 0 {
		return errno.ErrRecordNotFound
	}

	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user := new(User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		post, err := GetFavoriteRelation(ctx, uid, pid)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoritePosts").Delete(post)
		if err != nil {
			return err
		}

		//2.改变 video 表中的 favorite count
		res := tx.Model(post).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
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

// FavoriteList returns a list of Favorite videos.
func FavoriteList(ctx context.Context, uid int64) ([]Post, error) {
	user := new(User)
	if err := DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	posts := []Post{}
	// if err := DB.WithContext(ctx).First(&video, vid).Error; err != nil {
	// 	return nil, err
	// }

	if err := DB.WithContext(ctx).Model(&user).Association("FavoritePosts").Find(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}
