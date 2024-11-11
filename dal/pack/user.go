package pack

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"gorm.io/gorm"
)

// User pack user info
// db.User is for the connection with database
// user.User is for rpc
// Also judge the relation between fromID and u.id
func User(ctx context.Context, u *db.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "Account deleted",
		}, nil
	}

	follow_count := int64(u.FollowingCount)
	follower_count := int64(u.FollowerCount)

	// true->fromID follows u.ID，false-fromID doesn't follow u.ID
	isFollow := false
	relation, err := db.GetRelation(ctx, fromID, int64(u.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if relation != nil {
		isFollow = true
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   &follow_count,
		FollowerCount: &follower_count,
		IsFollow:      isFollow,
	}, nil
}

func Users(ctx context.Context, us []*db.User, fromID int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
