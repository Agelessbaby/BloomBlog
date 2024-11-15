package main

import (
	"context"
	publish "github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.BloomblogPublishActionRequest) (resp *publish.BloomblogPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.BloomblogPublishListRequest) (resp *publish.BloomblogPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
