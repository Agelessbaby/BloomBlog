package command

import (
	"bytes"
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/oss/minio"
	"github.com/google/uuid"
	"strings"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *publish.BloomblogPublishActionRequest, uid int64) error {
	MinioImageBucketName := minio.MinioVideoBucketName
	images := req.Images
	u3, err := uuid.NewV7()
	if err != nil {
		return err
	}
	coverName := u3.String() + "." + "jpg"
	cover := req.Cover
	coverReader := bytes.NewReader(cover)
	err = minio.UploadFile(MinioImageBucketName, coverName, coverReader, int64(len(cover)))
	if err != nil {
		return err
	}
	Url, err := minio.GetFileUrl(MinioImageBucketName, coverName, 0)
	if err != nil {
		return err
	}
	coverUrl := strings.Split(Url.String(), "?")[0]
	var imageurls []string
	for _, image := range images {
		reader := bytes.NewReader(image)
		//no need to consider uuid is replica
		u2, err := uuid.NewV7()
		if err != nil {
			return err
		}
		fileName := u2.String() + "." + "JPG"
		err = minio.UploadFile(MinioImageBucketName, fileName, reader, int64(len(image)))
		if err != nil {
			return err
		}
		url, err := minio.GetFileUrl(MinioImageBucketName, fileName, 0)
		playUrl := strings.Split(url.String(), "?")[0]
		if err != nil {
			return err
		}
		imageurls = append(imageurls, playUrl)
	}
	return db.CreatePost(s.ctx, &db.Post{
		AuthorID:      int(uid),
		ImageUrls:     imageurls,
		CoverUrl:      coverUrl,
		TextContent:   req.TextContent,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	})
}
