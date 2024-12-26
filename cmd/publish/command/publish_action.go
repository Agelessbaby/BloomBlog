package command

import (
	"bytes"
	"context"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/oss/S3"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type PublishActionService struct {
	ctx context.Context
}

func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

func (s *PublishActionService) PublishAction(req *publish.BloomblogPublishActionRequest, uid int64) error {
	MinioImageBucketName := oss.MinioVideoBucketName
	images := req.Images
	u3, err := uuid.NewV7()
	if err != nil {
		return err
	}
	coverName := u3.String() + "." + "jpg"
	cover := req.Cover
	coverReader := bytes.NewReader(cover)
	err = oss.UploadFile(MinioImageBucketName, coverName, coverReader, int64(len(cover)))
	if err != nil {
		return err
	}
	Url, err := oss.GetFileUrl(MinioImageBucketName, coverName, 0)
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
		fileName := u2.String() + "." + "jpg"
		err = oss.UploadFile(MinioImageBucketName, fileName, reader, int64(len(image)))
		if err != nil {
			return err
		}
		url, err := oss.GetFileUrl(MinioImageBucketName, fileName, 0)
		playUrl := strings.Split(url.String(), "?")[0]
		if err != nil {
			return err
		}
		imageurls = append(imageurls, playUrl)
	}
	post := &db.Post{
		AuthorID:      int(uid),
		ImageUrls:     imageurls,
		CoverUrl:      coverUrl,
		TextContent:   req.TextContent,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	err = db.CreatePost(s.ctx, post)
	//TODO Add the logic of pushing to user email
	if err != nil {
		return err
	}
	err = s.pushToFollowers(s.ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PublishActionService) pushToFollowers(ctx context.Context, ps *db.Post) error {
	relations, err := db.FollowerList(ctx, int64(ps.AuthorID))
	if err != nil {
		return err
	}
	for _, relation := range relations {
		err := db.InsertIntoTimeline(s.ctx, ps)
		if err != nil {
			klog.Infof("insert into timeline failed:%s %s", relation.UserID, ps.ID)
		}
	}
}
