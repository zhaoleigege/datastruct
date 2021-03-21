package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"strings"
)

const natsUrl = "nats://nats-server:4222"
const topic = "nats-test"

func main() {
	con, err := nats.Connect(natsUrl)
	if err != nil {
		panic(err)
	}
	defer con.Close()

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

		if err = con.Publish(topic, []byte(data)); err != nil {
			_ = fmt.Errorf("获取输入错误err: %s", err.Error())
			continue
		}
	}
}
