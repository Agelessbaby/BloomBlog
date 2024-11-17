package main

import (
	"context"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish"
	"github.com/Agelessbaby/BloomBlog/cmd/publish/kitex_gen/publish/publishsrv"
	"github.com/cloudwego/kitex/client"
)

func main() {
	//filePath := "/Users/liuwenjie/GolandProjects/BloomBlog/util/minio/minio/test/test.jpg"
	//// 读取文件内容
	//imageData, err := os.ReadFile(filePath)
	//if err != nil {
	//	fmt.Printf("Error reading file: %v\n", err)
	//	return
	//}
	client, err := publishsrv.NewClient("BloomBlogPublishServer", client.WithHostPorts("0.0.0.0:8084"))
	if err != nil {
		panic(err)
	}
	resp, err := client.PublishList(context.TODO(), &publish.BloomblogPublishListRequest{
		UserId: 2,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIiLCJpc3MiOiJCbG9vbUJMb2ciLCJhdWQiOiIiLCJzdWIiOiIiLCJpYXQiOjE3MzE2NzQ1ODgsIm5iZiI6MCwiZXhwIjoxNzMyMzY1Nzg4LCJ1ZCI6eyJ1c2VyaWQiOjJ9fQ.YHJ_b2ytwROGvAI7e1jvVZcIY5jtixws5BclVxjVrSw",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resp: %v\n", resp)
}
