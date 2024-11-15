package main

import (
	"context"
	feed "github.com/Agelessbaby/BloomBlog/cmd/feed/kitex_gen/feed"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetFeed(ctx context.Context, req *feed.BloomblogFeedRequest) (resp *feed.BloomblogFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPostById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetPostById(ctx context.Context, req *feed.PostIdRequest) (resp *feed.Post, err error) {
	// TODO: Your code here...
	return
}
