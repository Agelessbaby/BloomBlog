// Code generated by Kitex v0.11.3. DO NOT EDIT.

package feedsrv

import (
	"context"
	"errors"
	feed "github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetFeed": kitex.NewMethodInfo(
		getFeedHandler,
		newGetFeedArgs,
		newGetFeedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetPostById": kitex.NewMethodInfo(
		getPostByIdHandler,
		newGetPostByIdArgs,
		newGetPostByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	feedSrvServiceInfo                = NewServiceInfo()
	feedSrvServiceInfoForClient       = NewServiceInfoForClient()
	feedSrvServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return feedSrvServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return feedSrvServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return feedSrvServiceInfoForClient
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
	serviceName := "FeedSrv"
	handlerType := (*feed.FeedSrv)(nil)
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
		"PackageName": "feed",
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

func getFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(feed.BloomblogFeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(feed.FeedSrv).GetFeed(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetFeedArgs:
		success, err := handler.(feed.FeedSrv).GetFeed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFeedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetFeedArgs() interface{} {
	return &GetFeedArgs{}
}

func newGetFeedResult() interface{} {
	return &GetFeedResult{}
}

type GetFeedArgs struct {
	Req *feed.BloomblogFeedRequest
}

func (p *GetFeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(feed.BloomblogFeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetFeedArgs) Unmarshal(in []byte) error {
	msg := new(feed.BloomblogFeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFeedArgs_Req_DEFAULT *feed.BloomblogFeedRequest

func (p *GetFeedArgs) GetReq() *feed.BloomblogFeedRequest {
	if !p.IsSetReq() {
		return GetFeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFeedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFeedResult struct {
	Success *feed.BloomblogFeedResponse
}

var GetFeedResult_Success_DEFAULT *feed.BloomblogFeedResponse

func (p *GetFeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(feed.BloomblogFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetFeedResult) Unmarshal(in []byte) error {
	msg := new(feed.BloomblogFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFeedResult) GetSuccess() *feed.BloomblogFeedResponse {
	if !p.IsSetSuccess() {
		return GetFeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*feed.BloomblogFeedResponse)
}

func (p *GetFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFeedResult) GetResult() interface{} {
	return p.Success
}

func getPostByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(feed.PostIdRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(feed.FeedSrv).GetPostById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetPostByIdArgs:
		success, err := handler.(feed.FeedSrv).GetPostById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPostByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetPostByIdArgs() interface{} {
	return &GetPostByIdArgs{}
}

func newGetPostByIdResult() interface{} {
	return &GetPostByIdResult{}
}

type GetPostByIdArgs struct {
	Req *feed.PostIdRequest
}

func (p *GetPostByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(feed.PostIdRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPostByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPostByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPostByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPostByIdArgs) Unmarshal(in []byte) error {
	msg := new(feed.PostIdRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPostByIdArgs_Req_DEFAULT *feed.PostIdRequest

func (p *GetPostByIdArgs) GetReq() *feed.PostIdRequest {
	if !p.IsSetReq() {
		return GetPostByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPostByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPostByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPostByIdResult struct {
	Success *feed.Post
}

var GetPostByIdResult_Success_DEFAULT *feed.Post

func (p *GetPostByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(feed.Post)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPostByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPostByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPostByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPostByIdResult) Unmarshal(in []byte) error {
	msg := new(feed.Post)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPostByIdResult) GetSuccess() *feed.Post {
	if !p.IsSetSuccess() {
		return GetPostByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPostByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*feed.Post)
}

func (p *GetPostByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPostByIdResult) GetResult() interface{} {
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

func (p *kClient) GetFeed(ctx context.Context, Req *feed.BloomblogFeedRequest) (r *feed.BloomblogFeedResponse, err error) {
	var _args GetFeedArgs
	_args.Req = Req
	var _result GetFeedResult
	if err = p.c.Call(ctx, "GetFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPostById(ctx context.Context, Req *feed.PostIdRequest) (r *feed.Post, err error) {
	var _args GetPostByIdArgs
	_args.Req = Req
	var _result GetPostByIdResult
	if err = p.c.Call(ctx, "GetPostById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
