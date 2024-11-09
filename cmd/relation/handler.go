package main

import (
	"context"
	relation "github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.BloomblogRelationActionRequest) (resp *relation.BloomblogRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.BloomblogRelationFollowListRequest) (resp *relation.BloomblogRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.BloomblogRelationFollowerListRequest) (resp *relation.BloomblogRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
