package pubsub

import (
	"errors"
	"log"
	"sync"
	"time"
)

var ErrTopicNotExists = errors.New("topic does not exist")

type PubSub struct {
	mu         sync.Mutex
	topics     map[string][]chan interface{}
	workerPool chan func()
}

func NewPubSub(workerCount int) *PubSub {
	ps := &PubSub{
		topics:     make(map[string][]chan interface{}),
		workerPool: make(chan func(), workerCount),
	}
	for i := 0; i < workerCount; i++ {
		go func() {
			for task := range ps.workerPool {
				task()
			}
		}()
	}
	return ps
}

func (ps *PubSub) Publish(topic string, msg interface{}) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if subscribers, ok := ps.topics[topic]; ok {
		for _, subscriber := range subscribers {
			ps.workerPool <- func() {
				select {
				case subscriber <- msg:
				case <-time.After(5 * time.Second):
					log.Printf("Failed to send message to subscriber: timeout")
				}
			}
		}
		return nil
	}
	log.Printf("Topic %s does not exist", topic)
	return ErrTopicNotExists
}

func (ps *PubSub) Subscribe(topic string, ch chan interface{}) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.topics[topic]; !ok {
		ps.topics[topic] = []chan interface{}{}
	}
	ps.topics[topic] = append(ps.topics[topic], ch)
	return nil
}

func (ps *PubSub) Unsubscribe(topic string, ch chan interface{}) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if subscribers, ok := ps.topics[topic]; ok {
		for i, subscriber := range subscribers {
			if subscriber == ch {
				ps.topics[topic] = append(subscribers[:i], subscribers[i+1:]...)
				break
			}
		}
	}
	return nil
}
