package main

import (
	"context"
	"fmt"
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
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildRelationActionResp(errno.ErrTokenInvalid)
		return resp, nil
	}

	if req.ToUserId < 0 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}

	user_id := jwt.GetUserIdFromPayload(payload)

	if req.UserId == 0 && user_id > 0 {
		req.UserId = user_id
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildRelationActionResp(errno.ErrBind)
		return resp, nil
	}
	err = command.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		fmt.Println(1)
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.BloomblogRelationFollowListRequest) (resp *relation.BloomblogRelationFollowListResponse, err error) {
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildFollowingListResp(errno.ErrTokenInvalid)
	}
	current_user_id := jwt.GetUserIdFromPayload(payload)

	//Only when UserId not put in
	//current_user_id is the user who launch this request, req.UserId is the user whose following list is being queried.
	if req.UserId == 0 && current_user_id > 0 {
		req.UserId = current_user_id
	}
	followingUsers, err := command.NewRelationListService(ctx).FollowingList(req, current_user_id)
	if err != nil {
		resp = pack.BuildFollowingListResp(err)
		return resp, nil
	}
	resp = pack.BuildFollowingListResp(errno.Success)
	resp.UserList = followingUsers
	return resp, nil
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.BloomblogRelationFollowerListRequest) (resp *relation.BloomblogRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	_, payload, err := jwt.VerifyJwt(req.Token, env.JWT_SECRET)
	if err != nil {
		resp = pack.BuildFollowerListResp(errno.ErrTokenInvalid)
		return resp, nil
	}
	current_user_id := jwt.GetUserIdFromPayload(payload)

	//Only when UserId not put in
	//current_user_id is the user who launch this request, req.UserId is the user whose follower list is being queried.
	if req.UserId == 0 && current_user_id > 0 {
		req.UserId = current_user_id
	}

	followers, err := command.NewRelationListService(ctx).FollowerList(req, current_user_id)
	if err != nil {
		resp = pack.BuildFollowerListResp(err)
		return resp, nil
	}
	resp = pack.BuildFollowerListResp(errno.Success)
	resp.UserList = followers
	return resp, nil
}
