package main

import (
	"log"
	"time"

	"github.com/yikuanzz/go-pattern/PublishSubscribe/pubsub"
)

func main() {
	// Create a new PubSub instance
	pubSub := pubsub.NewPubSub(10)

	// Create a publisher
	publisher := pubsub.NewPublisher(pubSub)

	// Create subscribers
	subscriber1 := pubsub.NewSubscriber(10)
	subscriber2 := pubsub.NewSubscriber(10)

	// 订阅主题
	subscriber1.Subscribe("news", pubSub)
	subscriber2.Subscribe("news", pubSub)

	// 发布消息
	publisher.Publish("news", "Hello, World!")
	log.Println("Published message")

	// 接收消息
	go func() {
		for msg := range subscriber1.Receive() {
			log.Println("Subscriber 1 received:", msg)
		}
	}()

	go func() {
		for msg := range subscriber2.Receive() {
			log.Println("Subscriber 2 received:", msg)
		}
	}()

	// 等待消息处理完成
	time.Sleep(1 * time.Second)

	// 取消订阅
	subscriber1.Unsubscribe("news", pubSub)
	subscriber2.Unsubscribe("news", pubSub)
}
