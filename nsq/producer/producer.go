package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

const nsqAddr = "172.17.0.4:4150"
const topic = "nsq-topic"

func InitConfig() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(nsqAddr, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func main() {
	producer, err := InitConfig()
	if err != nil {
		panic(err)
	}
	defer producer.Stop()

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			_ = fmt.Errorf("获取输入错误err: %s", err.Error())
			continue
		}

		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" {
			break
		}

		if data == "" {
			continue
		}

		if err = producer.Publish(topic, []byte(data)); err != nil {
			_ = fmt.Errorf("获取输入错误err: %s", err.Error())
			continue
		}
	}
}
