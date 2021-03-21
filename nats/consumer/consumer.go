package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const natsUrl = "nats://nats-server:4222"
const topic = "nats-test"

func main() {
	con, err := nats.Connect(natsUrl, func(options *nats.Options) error {
		options.SubChanLen = 1
		return nil
	})
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	con.Subscribe(topic, func(msg *nats.Msg) {
		wg.Add(1)
		defer wg.Done()

		time.Sleep(5 * time.Second)
		fmt.Println(string(msg.Data))
	})

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("客户端关闭\n")
			wg.Wait()
			return
		case <-c:
			fmt.Printf("强制退出\n")
			cancel()
			con.Close()
		}
	}
}
