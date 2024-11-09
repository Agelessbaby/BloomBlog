package pack

import (
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
