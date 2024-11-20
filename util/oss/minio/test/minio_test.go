package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// (TiktokTesst)bucket name  ccontains invalid characters
// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	err := minio.CreateBucket("bloomblogtest")
	if err != nil {
		t.Error(err)
	}
}

func TestUploadLocalFile(t *testing.T) {
	info, err := minio.UploadLocalFile("bloomblogtest", "test.jpg", "./test.jpg", "image/jpg")
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
	err = minio.UploadFile("bloomblogtest", "test2.jpg", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := minio.GetFileUrl("bloomblogtest", "test.jpg", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
