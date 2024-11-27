package pack

import (
	"errors"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
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

// BuildPostResp build PostResp from error
func BuildPostResp(err error) *feed.BloomblogFeedResponse {
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

// BuildFavoriteActionResp build FavoriteActionResp from error
func BuildFavoriteActionResp(err error) *favorite.BloomblogFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return favoriteActionResp(s)
}

func favoriteActionResp(err errno.ErrNo) *favorite.BloomblogFavoriteActionResponse {
	return &favorite.BloomblogFavoriteActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildFavoriteListResp build FavoriteListResp from error
func BuildFavoriteListResp(err error) *favorite.BloomblogFavoriteListResponse {
	if err == nil {
		return favoriteListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return favoriteListResp(s)
}

func favoriteListResp(err errno.ErrNo) *favorite.BloomblogFavoriteListResponse {
	return &favorite.BloomblogFavoriteListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildCommentActionResp build CommentActionResp from error
func BuildCommentActionResp(err error) *comment.BloomblogCommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentActionResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return commentActionResp(s)
}

func commentActionResp(err errno.ErrNo) *comment.BloomblogCommentActionResponse {
	return &comment.BloomblogCommentActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildCommentListResp build CommentListResp from error
func BuildCommentListResp(err error) *comment.BloomblogCommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return commentListResp(s)
}

func commentListResp(err errno.ErrNo) *comment.BloomblogCommentListResponse {
	return &comment.BloomblogCommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func BuildSlCommentListResp(err error) *comment.Bloomblog_SlCommentListResponse {
	if err == nil {
		return slcommentListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return slcommentListResp(e)
	}

	s := errno.ErrUnknown.WithMessage(err.Error())
	return slcommentListResp(s)
}

func slcommentListResp(err errno.ErrNo) *comment.Bloomblog_SlCommentListResponse {
	return &comment.Bloomblog_SlCommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
