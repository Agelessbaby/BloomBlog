package mq

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

// produce messages in direct mode
func ProduceDirect(msg []byte, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := Ch.PublishWithContext(
		ctx,
		Exg,   //exchange。""为默认的Exchange（direct类型），这种Exchange会把消息传递给routing key指定的Queue。
		key,   //routing key。Exchange为""时，routing key就是QueueName
		false, //mandatory
		false, //immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,    //消息如果想持久化进磁盘，即确保RabbitMQ Server（或称broker）重启后消息不丢失，需同时满足2个条件：队列需要是durable，消息需要是Persistent。Transient显然意味着更高的吞吐。另外即使设置了Persistent，消息也不是立即会写入磁盘，中间有缓冲，如果broker突然挂掉，缓冲里的数据会丢失。
			ContentType:  "application/json", //MIME content type
			Body:         msg,
		},
	)
	if err != nil {
		klog.Errorf("publish message failed: %s", err)
	}
}
