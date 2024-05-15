package rudrya

// Consumer is responsible for receiving messages from a broker.
type Consumer struct {
    broker *Broker
    topic  string
    ch     chan string
}

// NewConsumer initializes a new Consumer instance and registers it with the broker.
func NewConsumer(broker *Broker, topic string) *Consumer {
    ch := make(chan string)
    broker.RegisterConsumer(topic, ch)
    return &Consumer{
        broker: broker,
        topic:  topic,
        ch:     ch,
    }
}

// Start begins listening for messages on the consumer's channel.
// The handleFunc parameter is a callback function that processes each message.
func (c *Consumer) Start(handleFunc func(string)) {
    for message := range c.ch {
        handleFunc(message)
    }
}
