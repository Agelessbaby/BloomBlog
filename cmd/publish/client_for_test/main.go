package main

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish/publishsrv"
	"github.com/cloudwego/kitex/client"
	"os"
)

func main() {
	filePath := "/Users/liuwenjie/GolandProjects/BloomBlog/util/minio/test/test.jpg"
	// 读取文件内容
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("Image as binary: %v\n", imageData) // 打印二进制数据（可选择存储或传输）
	client, err := publishsrv.NewClient("BloomBlogPublishServer", client.WithHostPorts("0.0.0.0:8084"))
	if err != nil {
		panic(err)
	}
	resp, err := client.PublishAction(context.TODO(), &publish.BloomblogPublishActionRequest{
		Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzE2NzQ1ODgsIm5iZiI6MCwiZXhwIjoxNzMyMzY1Nzg4LCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.YHJ_b2ytwROGvAI7e1jvVZcIY5jtixws5BclVxjVrSw",
		Images:      [][]byte{imageData},
		TextContent: "A test",
		Title:       "Title",
		Cover:       imageData,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resp: %v\n", resp)
}
