package handlers

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// GetFeed retrieves a feed based on the latest time and user token.
//
//	@Summary		Retrieve feed
//	@Description	This endpoint allows users to fetch feed data using the latest timestamp and an optional authentication token.
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			latest_time	query		int64						false	"The timestamp for the latest feed item (optional). Defaults to now if not provided."
//	@Success		200			{object}	feed.BloomblogFeedResponse	"Successful operation, returns the feed data."
//	@Failure		400			{object}	feed.BloomblogFeedResponse	"Invalid input data, such as malformed timestamp."
//	@Failure		500			{object}	feed.BloomblogFeedResponse	"Internal server error."
//	@Router			/bloomblog/feed/getfeed [get]
func GetFeed(c context.Context, ctx *app.RequestContext) {
	var feedVar FeedParam
	latest_time := ctx.Query("latest_time")
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	var lt int64
	if len(latest_time) != 0 {
		if latestTime, err := strconv.Atoi(latest_time); err != nil {
			SendResponse(ctx, pack.BuildPostResp(errno.ErrDecodingFailed))
			return
		} else {
			lt = int64(latestTime)
		}
	}
	feedVar.LatestTime = &lt
	feedVar.Token = &token
	resp, err := rpc.GetFeed(c, &feed.BloomblogFeedRequest{
		LatestTime: feedVar.LatestTime,
		Token:      feedVar.Token,
	})
	if err != nil {
		SendResponse(ctx, pack.BuildPostResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
