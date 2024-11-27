package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse pack response
func SendResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}

// The input parameter for user register
type UserRegisterParam struct {
	UserName string `json:"username"` // 用户名
	PassWord string `json:"password"` // 用户密码
}

// The unput parameter for getting user by id
type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}

// The input parameter for follow or unfollow
type RelationActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id,omitempty"`  // 对方用户id
	ActionType int32  `json:"action_type,omitempty"` // 1-关注，2-取消关注
}

// The input parameter for publish posts
type PublishActionParam struct {
	Token   string   `json:"token,omitempty"`
	Images  [][]byte `json:"images,omitempty"`
	Content string   `json:"content,omitempty"`
	Title   string   `json:"title,omitempty"`
}

// The input parameter for getting user feed
type FeedParam struct {
	LatestTime *int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty"`       // 可选参数，登录用户设置
}

type FavoriteActionParam struct {
	UserId     int64  `json:"user_id,omitempty"`     // 用户id
	Token      string `json:"token,omitempty"`       // 用户鉴权token
	PostId     int64  `json:"post_id,omitempty"`     // 视频id
	ActionType int32  `json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}

type FavoriteListParam struct {
	Token  string `json:"token,omitempty"` // 用户鉴权token
	UserId int64  `json:"user_id,omitempty"`
}

type CommentActionParam struct {
	UserId     int64   `json:"user_id,omitempty"`
	Token      string  `json:"token,omitempty"`
	ActionType int32   `json:"action_type,omitempty"`
	Content    *string `json:"content,omitempty"`
	PostId     int64   `json:"post_id,omitempty"`
	CommentId  *int64  `json:"comment_id,omitempty"`
	ParentId   *int64  `json:"parent_id,omitempty"`
	ReplyId    *int64  `json:"reply_id,omitempty"`
}

type CommentListParam struct {
	Token  string `json:"token,omitempty"`
	PostId int64  `json:"post_id,omitempty"`
}

type SlCommentListParam struct {
	ParentId int64  `json:"parent_id,omitempty"`
	Token    string `json:"token,omitempty"`
}
