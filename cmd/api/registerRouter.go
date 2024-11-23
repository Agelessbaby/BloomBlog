package main

import (
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/handlers"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

// Register router groups
func registerGroup(h *server.Hertz) {
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	bloomblog := h.Group("/bloomblog")

	// swagger
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	// User service
	user := bloomblog.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	// TODO Now the method is POST, which should be changed into GET
	user.GET("/getuserbyid", handlers.GetUserById)

	// Relation service
	relation := bloomblog.Group("/relation")
	relation.POST("/action", handlers.RelationAction)
	relation.GET("/followlist", handlers.RelationFollowList)
	relation.GET("/followerlist", handlers.RelationFollowerList)

	// Publish service
	publish := bloomblog.Group("/publish")
	publish.POST("/action", handlers.PublishAction)
	//TODO Add Publish service here
}
