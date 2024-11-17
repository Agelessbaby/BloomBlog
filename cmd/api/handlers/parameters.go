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

// The inputn parameter for publish posts
type PublishActionParam struct {
	Token   string   `json:"token,omitempty"`
	Images  [][]byte `json:"images,omitempty"`
	Content string   `json:"content,omitempty"`
	Title   string   `json:"title,omitempty"`
}
