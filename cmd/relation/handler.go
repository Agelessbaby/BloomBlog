package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/relation/command"
	relation "github.com/Agelessbaby/BloomBlog/cmd/relation/kitex_gen/relation"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	env "github.com/Agelessbaby/BloomBlog/util"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.BloomblogRelationActionRequest) (resp *relation.BloomblogRelationActionResponse, err error) {
	// TODO: Your code here...
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildRelationActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.ToUserId < 0 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}

	user_id_val := payload.UserDefined["userid"]
	user_id := int64(user_id_val.(float64))

	if req.UserId == 0 || user_id > 0 {
		req.UserId = user_id
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}
	err = command.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
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
