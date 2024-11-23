package handlers

import (
	"bytes"
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/api/rpc"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/dal/pack"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	"github.com/Agelessbaby/BloomBlog/util/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

// PublishAction handles the publication of a blog post with images.
//
//	@Summary		Publish a blog post
//	@Description	This endpoint allows users to publish a blog post with images, title, and content.
//
//	The images are uploaded as a multipart form, and the first image is used as the cover.
//
//	@Tags			publish
//	@Accept			multipart/form-data
//	@Produce		json
//	@Security		BearerAuth
//	@Param			title	formData	string									true	"The title of the blog post"
//	@Param			content	formData	string									true	"The content of the blog post"
//	@Param			images	formData	file									true	"Multiple image files (repeatable), upload all images as part of 'images' field"
//	@Success		200		{object}	publish.BloomblogPublishActionResponse	"The response object"
//	@Failure		400		{object}	publish.BloomblogPublishActionResponse	"Invalid input data"
//	@Failure		500		{object}	publish.BloomblogPublishActionResponse	"Internal server error"
//	@Router			/bloomblog/publish/action [post]
func PublishAction(c context.Context, ctx *app.RequestContext) {
	var publishVar PublishActionParam
	token := jwt.TrimPrefix(string(ctx.GetHeader("Authorization")))
	publishVar.Token = token
	publishVar.Title = ctx.PostForm("title")
	publishVar.Content = ctx.PostForm("content")
	if len(publishVar.Title) == 0 || len(publishVar.Content) == 0 || len(publishVar.Token) == 0 {
		SendResponse(ctx, pack.BuildPublishResp(errno.ErrBind))
		return
	}
	form, err := ctx.Request.MultipartForm()
	if err != nil {
		SendResponse(ctx, pack.BuildPublishResp(errno.ErrBind))
		return
	}
	files := form.File["images"]
	images := make([][]byte, 0)
	for _, fileheader := range files {
		file, err := fileheader.Open()
		if err != nil {
			SendResponse(ctx, pack.BuildPublishResp(errno.ErrBind))
			return
		}
		buf := bytes.NewBuffer(nil)
		_, err = buf.ReadFrom(file)
		defer file.Close()
		if err != nil {
			SendResponse(ctx, pack.BuildPublishResp(errno.ErrBind))
			return
		}
		images = append(images, buf.Bytes())
	}
	resp, err := rpc.PublishAction(c, &publish.BloomblogPublishActionRequest{
		Token:       publishVar.Token,
		Images:      images,
		TextContent: publishVar.Content,
		Title:       publishVar.Title,
		Cover:       images[0],
	})
	if err != nil {
		SendResponse(ctx, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(ctx, resp)
}
