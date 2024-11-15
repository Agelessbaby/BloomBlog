package minio

import (
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient          *minio.Client
	minioConfig          = config.CreateConfig("minioConfig")
	MinioEndpoint        = minioConfig.GetString("minio.Endpoint")
	MinioAccessKeyId     = minioConfig.GetString("minio.AccessKeyId")
	MinioSecretAccessKey = minioConfig.GetString("minio.SecretAccessKey")
	MinioUseSSL          = minioConfig.GetBool("minio.UseSSL")
	MinioVideoBucketName = minioConfig.GetString("minio.ImageBucketName")
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	// fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
}
