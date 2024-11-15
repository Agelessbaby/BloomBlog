package pack

import (
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/Agelessbaby/BloomBlog/util/errno"
)

func BuildUserRegisterResp(err error) *user.BloomBlogUserRegisterResponse {
	if err == nil || err == errno.Success {
		msg := "success"
		return &user.BloomBlogUserRegisterResponse{StatusCode: int32(0), StatusMsg: &msg}
	}
	msg := err.Error()
	return &user.BloomBlogUserRegisterResponse{StatusCode: int32(1), StatusMsg: &msg}
}

func BuildUserLoginResponse(err error) *user.BloomBlogUserRegisterResponse {
	if err == nil || err == errno.Success {
		msg := "success"
		return &user.BloomBlogUserRegisterResponse{StatusCode: int32(0), StatusMsg: &msg}
	}
	msg := err.Error()
	return &user.BloomBlogUserRegisterResponse{StatusCode: int32(1), StatusMsg: &msg}
}

func BuildUserUserResp(err error) *user.BloomBlogUserResponse {
	if err == nil || err == errno.Success {
		msg := "success"
		return &user.BloomBlogUserResponse{StatusCode: int32(0), StatusMsg: &msg}
	}
	msg := err.Error()
	return &user.BloomBlogUserResponse{StatusCode: int32(1), StatusMsg: &msg}
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

// BuildVideoResp build VideoResp from error
func BuildVideoResp(err error) *feed.BloomblogFeedResponse {
	if err == nil {
		return PostResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return PostResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return PostResp(s)
}

func PostResp(err errno.ErrNo) *feed.BloomblogFeedResponse {
	return &feed.BloomblogFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildPublishResp build PublishResp from error
func BuildPublishResp(err error) *publish.BloomblogPublishActionResponse {
	if err == nil {
		return publishResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return publishResp(s)
}

func publishResp(err errno.ErrNo) *publish.BloomblogPublishActionResponse {
	return &publish.BloomblogPublishActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildPublishResp build PublishResp from error
func BuildPublishListResp(err error) *publish.BloomblogPublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return publishListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return publishListResp(s)
}

func publishListResp(err errno.ErrNo) *publish.BloomblogPublishListResponse {
	return &publish.BloomblogPublishListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
