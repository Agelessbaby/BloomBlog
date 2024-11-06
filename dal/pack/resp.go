package pack

import "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"

func BuildUserRegisterResp(err error) *user.BloomBlogUserRegisterResponse {
	if err == nil {
		msg := "success"
		return &user.BloomBlogUserRegisterResponse{StatusCode: int32(0), StatusMsg: &msg}
	}
	msg := err.Error()
	return &user.BloomBlogUserRegisterResponse{StatusCode: int32(1), StatusMsg: &msg}
}

func BuildUserLoginResponse(err error) *user.BloomBlogUserRegisterResponse {
	if err == nil {
		msg := "success"
		return &user.BloomBlogUserRegisterResponse{StatusCode: int32(0), StatusMsg: &msg}
	}
	msg := err.Error()
	return &user.BloomBlogUserRegisterResponse{StatusCode: int32(1), StatusMsg: &msg}
}
