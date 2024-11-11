package pack

import (
	"context"
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"gorm.io/gorm"
)

// FollowingList pack lists of following info.
func FollowingList(ctx context.Context, vs []*db.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.GetUserByID(ctx, int64(v.ToUserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}

// FollowerList pack lists of follower info.
func FollowerList(ctx context.Context, vs []*db.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*db.User, 0)
	for _, v := range vs {
		user2, err := db.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}

func relationActionResp(err errno.ErrNo) *relation.BloomblogRelationActionResponse {
	return &relation.BloomblogRelationActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildRelationActionResp build RelationActionResp from error
func BuildRelationActionResp(err error) *relation.BloomblogRelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return relationActionResp(s)
}

// BuildFollowingListResp build FollowingListResp from error
func BuildFollowingListResp(err error) *relation.BloomblogRelationFollowListResponse {
	if err == nil {
		return followingListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return followingListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return followingListResp(s)
}

func followingListResp(err errno.ErrNo) *relation.BloomblogRelationFollowListResponse {
	return &relation.BloomblogRelationFollowListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}

// BuildFollowerListResp build FollowerListResp from error
func BuildFollowerListResp(err error) *relation.BloomblogRelationFollowerListResponse {
	if err == nil {
		return followerListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return followerListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return followerListResp(s)
}

func followerListResp(err errno.ErrNo) *relation.BloomblogRelationFollowerListResponse {
	return &relation.BloomblogRelationFollowerListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
