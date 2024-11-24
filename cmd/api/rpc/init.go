package rpc

import (
	"github.com/Agelessbaby/BloomBlog/util/config"
)

// InitRPC init rpc client
func InitRPC() {
	UserConfig := config.CreateConfig("userConfig")
	RelationConfig := config.CreateConfig("relationConfig")
	PublishConfig := config.CreateConfig("publishConfig")
	FeedConfig := config.CreateConfig("feedConfig")
	initUserRpc(UserConfig)
	initRelationRpc(RelationConfig)
	initPublishRpc(PublishConfig)
	initFeedRpc(FeedConfig)
}
