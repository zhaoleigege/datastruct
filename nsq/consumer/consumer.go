package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const nsqAddr = "172.17.0.4:4150"
const nsqLookupAddr = "172.17.0.3:4161"
const topic = "nsq-topic"
const channel = "nsq-channel"

type ConsumerHandler struct {
}

func (handler *ConsumerHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("address: %s, msg: %s\n", msg.NSQDAddress, string(msg.Body))

	return nil
}

func InitConfig() (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	handler := &ConsumerHandler{}
	consumer.AddHandler(handler)

	// 使用docker配置NSQLookupd时会报错，因为连接的是docker内部的nsqd的地址，
	if err = consumer.ConnectToNSQLookupd(nsqLookupAddr); err != nil {
		return nil, err
	}

	return consumer, nil
}
// 参考 https://product.reverb.com/how-to-write-an-nsq-consumer-in-go-96ed8bde29ef
func main() {
	consumer, err := InitConfig()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)

	for {
		select {
		case <-consumer.StopChan:
			fmt.Printf("服务端关闭\n")
			return
		case <-c:
			fmt.Printf("客户端关闭\n")
			consumer.Stop()
		}
	}
}
