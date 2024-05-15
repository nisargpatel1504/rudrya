package rudrya

// Producer is responsible for sending messages to a broker.
type Producer struct {
    broker *Broker
}

// NewProducer initializes a new Producer instance.
func NewProducer(broker *Broker) *Producer {
    return &Producer{broker: broker}
}

// SendMessage sends a message to a specified topic.
func (p *Producer) SendMessage(topic, message string) {
    p.broker.SendMessage(topic, message)
}
