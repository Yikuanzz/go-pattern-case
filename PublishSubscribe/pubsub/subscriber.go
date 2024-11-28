package pubsub

type subsciber struct {
	ch chan interface{}
}

func NewSubscriber(buffer int) *subsciber {
	return &subsciber{
		ch: make(chan interface{}, buffer),
	}
}

func (s *subsciber) Subscribe(topic string, pubsub *PubSub) error {
	return pubsub.Subscribe(topic, s.ch)
}

func (s *subsciber) Unsubscribe(topic string, pubsub *PubSub) error {
	return pubsub.Unsubscribe(topic, s.ch)
}

func (s *subsciber) Receive() <-chan interface{} {
	return s.ch
}
