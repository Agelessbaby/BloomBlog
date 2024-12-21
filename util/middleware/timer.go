package middleware

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/comment/kitex_gen/comment"
	"github.com/Agelessbaby/BloomBlog/cmd/favorite/kitex_gen/favorite"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/cmd/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"time"
)

func TimerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		begin := time.Now()
		err = next(ctx, req, resp)

		// caller is the upstream invoker
		caller, _ := kitexutil.GetCaller(ctx)
		addr, _ := kitexutil.GetCallerAddr(ctx)
		method, _ := kitexutil.GetMethod(ctx)

		klog.Infof("caller:%s,addr:%s,method:%s,time:%dms", caller, addr, method, time.Since(begin).Milliseconds())
		return err
	}
}

func RecordArgMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response any) error {
		if arg, ok := request.(utils.KitexArgs); ok {
			switch req := arg.GetFirstArgument().(type) {
			default:
				klog.Debug("Request %v", req)
			}
		}
		err := next(ctx, request, response)
		ri := rpcinfo.GetRPCInfo(ctx)
		if stats := ri.Stats(); stats != nil {
			//打印panic信息
			if panicHappened, panicInfo := stats.Panicked(); panicHappened { //panicInfo就是recover()的返回值
				klog.Errorf("panic info %s", panicInfo)
				return fmt.Errorf("panic inside the service") //hide the original panic info for the client
			} else {
				if stats.Error() == nil {
					//打印响应参数
					if result, ok := response.(utils.KitexResult); ok {
						switch resp := result.GetResult().(type) {
						case *user.BloomBlogUserRegisterResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("AddResponse Sum %d", resp)
							}
						case *user.BloomBlogUserResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d", resp)
							}
						case *relation.BloomblogRelationFollowerListRequest:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *relation.BloomblogRelationActionResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *relation.BloomblogRelationFollowListResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *comment.Bloomblog_SlCommentListResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *comment.BloomblogCommentActionResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *comment.BloomblogCommentListResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *feed.BloomblogFeedResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *publish.BloomblogPublishActionResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *publish.BloomblogPublishListResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *favorite.BloomblogFavoriteListResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						case *favorite.BloomblogFavoriteActionResponse:
							if resp == nil {
								klog.Error("kitex result is nil")
							} else {
								klog.Debugf("SubResponse Diff %d")
							}
						default:
							klog.Debug("Response %v", resp)
						}

					}
				} else {
					klog.Errorf("biz error: %s", stats.Error().Error())
				}
			}
		}
		return err
	}
}
