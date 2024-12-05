package oss

import (
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

var (
	minioClient          *minio.Client
	minioConfig          = config.CreateConfig("ossConfig")
	MinioEndpoint        = minioConfig.GetString("S3.Endpoint")
	MinioAccessKeyId     = os.Getenv("AWS_ACCESS_KEY_ID")
	MinioSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	MinioUseSSL          = minioConfig.GetBool("S3.UseSSL")
	MinioVideoBucketName = minioConfig.GetString("S3.ImageBucketName")
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		klog.Fatalf("test client init failed: %v", err)
	}
	klog.Debug("test client init successfully")
	minioClient = client
	//if err := CreateBucket(MinioVideoBucketName); err != nil {
	//	klog.Errorf("test client init failed: %v", err)
	//}
}
