package mq

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	mqConfig      = config.CreateConfig("mqConfig.yaml")
	host          = mqConfig.GetString("rabbitmq.host")
	port          = mqConfig.GetString("rabbitmq.port")
	username      = mqConfig.GetString("rabbitmq.username")
	password      = mqConfig.GetString("rabbitmq.password")
	Exg           = mqConfig.GetString("rabbitmq.exg")
	Ch            *amqp.Channel
	Conn          *amqp.Connection
	rabbitmq_once sync.Once
)

func InitMqConn() {
	rabbitmq_once.Do(func() {
		var err error
		Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port))
		if err != nil {
			klog.Fatalf("connect to RabbitMQ failed: %s", err)
		}
		Ch, err = Conn.Channel()
		if err != nil {
			klog.Fatalf("open channel failed: %s", err)
		}
		declareExchange()
		go func() {
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
			sig := <-signalChan
			klog.Info("Received signal:", sig)
			Ch.Close()
			Conn.Close()
		}()
	})
}

func declareExchange() {
	//声明Exchange。如果Exchange不存在会创建它；如果Exchange已存在，Server会检查声明的参数和Exchange的真实参数是否一致。
	err := Ch.ExchangeDeclare(
		Exg,
		"direct", //type
		true,     //durable
		false,    //auto delete
		false,    //internal
		false,    //no-wait
		nil,      //arguments
	)
	if err != nil {
		klog.Fatalf("declare exchange failed: %s", err)
	}
}
