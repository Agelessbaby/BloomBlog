package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer = <-chan amqp.Delivery

func SubscribeByKey(conn *amqp.Connection, exchange string, operationFunc operation, keys ...string) {
	//创建Channel
	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("open channel failed: %s", err)
	}
	defer ch.Close()
	//声明队列
	q, err := ch.QueueDeclare(
		"",    //队列名为空时Server指定一个随机（且唯一）的队列名
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panicf("declare queue failed: %s", err)
	}

	//队列和Exchange建立绑定关系。
	//队列默认绑定到Name为""的Exchange，该Exchange不需要创建默认已存在，且类型为direct。
	for _, key := range keys {
		err = ch.QueueBind(
			q.Name, //Queue Name
			key,    //routing key。匹配上这个key的消息会被发送到这个队列
			exchange,
			false, //noWait
			nil,   //arguments
		)
		if err != nil {
			klog.Fatalf("bind queue failed: %s", err)
		}
	}

	//一旦开始消费，就不要再修改绑定关系了
	consumer := createConsumer(ch, q.Name)
	consume(consumer, operationFunc)
}

// 一个amqp.Channel上可以创建多个consumer
func createConsumer(ch *amqp.Channel, qName string) Consumer {
	deliveryCh, err := ch.Consume(
		qName, //queue
		"",    //consumer
		false, //auto-ack。autoAck其实就是noAck，只要server把消息传给consumer，本消息就会被标记为ack，而不管它有没有被consumer成功消费。
		false, //exclusive
		false, //no-local
		false, //no-wait
		nil,   //args
	)
	if err != nil {
		log.Panicf("regist consumer failed: %s", err)
	}
	return deliveryCh
}

type operation func([]byte) error

func consume(deliveryCh Consumer, operationFunc operation) {
	for delivery := range deliveryCh {
		klog.Infof("receive message [%s][%s]", delivery.RoutingKey, delivery.Body)
		err := operationFunc(delivery.Body)
		if err != nil {
			klog.Errorf("operation comment failed: %s", err)
		}
		delivery.Ack(false) //通知Server此消息已成功消费。Ack参数为true时，此channel里之前未ack的消息会一并被ack（相当于批量ack）。如果没有ack，则下一次启动时还消费到此消息（除非超时30分钟，因为delivery在30分钟后会被强制ack）,因为channel close时，它里没有ack的消息会再次被放入队列的尾部。
	}
}
