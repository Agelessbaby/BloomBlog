package test

import (
	"fmt"
	oss "github.com/Agelessbaby/BloomBlog/util/oss/minio"
	"os"
	"strings"
	"testing"
)

// (TiktokTesst)bucket name  ccontains invalid characters
// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	err := oss.CreateBucket("bloomblog-images")
	if err != nil {
		t.Error(err)
	}
}

func TestUploadLocalFile(t *testing.T) {
	info, err := oss.UploadLocalFile("bloomblogtest", "test.jpg", "./test.jpg", "image/jpg")
	fmt.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, err := os.Open("./test.jpg")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer file.Close()
	fi, _ := os.Stat("./test.jpg")
	err = oss.UploadFile("bloomblogtest", "test2.jpg", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := oss.GetFileUrl("bloomblogtest", "test.jpg", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
