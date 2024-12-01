package test

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/dal/db"
	"github.com/Agelessbaby/BloomBlog/util/mq"
	"github.com/bytedance/sonic"
	"testing"
	"time"
)

var c = make(chan struct{})

func TestMqInit(t *testing.T) {
	mq.InitMqConn()
	go func() {
		for {
			time.Sleep(time.Second * 1)
			param := db.Comment{
				PostID:   1,
				UserID:   1,
				Content:  "123dfsafd",
				ParentID: nil,
				Parent:   nil,
				ReplyID:  nil,
				Reply:    nil,
			}
			bs, err := sonic.Marshal(param)
			if err != nil {
				t.Error(err)
			}
			mq.ProduceDirect(bs, "test")
		}

	}()
	mq.SubscribeByKey(mq.Conn, "bloomblog-exg", func(bs []byte) error {
		fmt.Println(bs)
		return nil
	}, "test")
}
