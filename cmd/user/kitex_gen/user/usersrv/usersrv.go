// Code generated by Kitex v0.11.3. DO NOT EDIT.

package usersrv

import (
	"context"
	"errors"
	user "github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserById": kitex.NewMethodInfo(
		getUserByIdHandler,
		newGetUserByIdArgs,
		newGetUserByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userSrvServiceInfo                = NewServiceInfo()
	userSrvServiceInfoForClient       = NewServiceInfoForClient()
	userSrvServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userSrvServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userSrvServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userSrvServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserSrv"
	handlerType := (*user.UserSrv)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.BloomBlogUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserSrv).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.BloomBlogUserRegisterRequest
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.BloomBlogUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.BloomBlogUserRegisterRequest

func (p *RegisterArgs) GetReq() *user.BloomBlogUserRegisterRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.BloomBlogUserRegisterResponse
}

var RegisterResult_Success_DEFAULT *user.BloomBlogUserRegisterResponse

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.BloomBlogUserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.BloomBlogUserRegisterResponse {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.BloomBlogUserRegisterResponse)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.BloomBlogUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserSrv).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.BloomBlogUserRegisterRequest
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.BloomBlogUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.BloomBlogUserRegisterRequest

func (p *LoginArgs) GetReq() *user.BloomBlogUserRegisterRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.BloomBlogUserRegisterResponse
}

var LoginResult_Success_DEFAULT *user.BloomBlogUserRegisterResponse

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.BloomBlogUserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.BloomBlogUserRegisterResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.BloomBlogUserRegisterResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.BloomBlogUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).GetUserById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserByIdArgs:
		success, err := handler.(user.UserSrv).GetUserById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserByIdArgs() interface{} {
	return &GetUserByIdArgs{}
}

func newGetUserByIdResult() interface{} {
	return &GetUserByIdResult{}
}

type GetUserByIdArgs struct {
	Req *user.BloomBlogUserRequest
}

func (p *GetUserByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.BloomBlogUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserByIdArgs) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserByIdArgs_Req_DEFAULT *user.BloomBlogUserRequest

func (p *GetUserByIdArgs) GetReq() *user.BloomBlogUserRequest {
	if !p.IsSetReq() {
		return GetUserByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserByIdResult struct {
	Success *user.BloomBlogUserResponse
}

var GetUserByIdResult_Success_DEFAULT *user.BloomBlogUserResponse

func (p *GetUserByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.BloomBlogUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserByIdResult) Unmarshal(in []byte) error {
	msg := new(user.BloomBlogUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserByIdResult) GetSuccess() *user.BloomBlogUserResponse {
	if !p.IsSetSuccess() {
		return GetUserByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.BloomBlogUserResponse)
}

func (p *GetUserByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserByIdResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *user.BloomBlogUserRegisterRequest) (r *user.BloomBlogUserRegisterResponse, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.BloomBlogUserRegisterRequest) (r *user.BloomBlogUserRegisterResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, Req *user.BloomBlogUserRequest) (r *user.BloomBlogUserResponse, err error) {
	var _args GetUserByIdArgs
	_args.Req = Req
	var _result GetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
