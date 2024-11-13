package db

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	//just a connection to user table
	User     User `gorm:"foreignkey:UserID;"`
	UserID   int  `gorm:"index:idx_userid,unique;not null"`
	ToUser   User `gorm:"foreignkey:ToUserID;"`
	ToUserID int  `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}

func (Relation) TableName() string {
	return "relation"
}

// NewRelation creates a new Relation
func NewRelation(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// use tx as transaction instead of DB
		err := tx.Create(&Relation{UserID: int(uid), ToUserID: int(tid)}).Error
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 2.change following count in user table
		res := tx.Model(new(User)).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		// 3.change follower count in user table
		res = tx.Model(new(User)).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
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

// DisRelation deletes a relation from the database.
func DisRelation(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		relation := new(Relation)
		if err := tx.Where("user_id = ? AND to_user_id=?", uid, tid).First(&relation).Error; err != nil {
			return err
		}

		// 1. delete relation data
		err := tx.Unscoped().Delete(&relation).Error
		if err != nil {
			return err
		}
		// 2.change following count
		var tempuser User
		res := tx.Model(&tempuser).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errno.ErrDatabase
		}

		// 3.change follower count
		res = tx.Model(&tempuser).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
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

// FollowingList returns the Following List.
func FollowingList(ctx context.Context, uid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}

// FollowerList returns the Follower List.
func FollowerList(ctx context.Context, tid int64) ([]*Relation, error) {
	var RelationList []*Relation
	err := DB.WithContext(ctx).Where("to_user_id = ?", tid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}

// GetRelation get relation info
func GetRelation(ctx context.Context, uid int64, tid int64) (*Relation, error) {
	relation := new(Relation)

	if err := DB.WithContext(ctx).First(&relation, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
		return nil, err
	}
	return relation, nil
}
