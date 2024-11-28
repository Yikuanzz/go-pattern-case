package pubsub

type Publisher struct {
	pubsub *PubSub
}

func NewPublisher(pubsub *PubSub) *Publisher {
	return &Publisher{pubsub: pubsub}
}

func (p *Publisher) Publish(topic string, message string) error {
	return p.pubsub.Publish(topic, message)
}
