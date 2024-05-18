package rudrya

// Consumer is responsible for receiving messages from a broker.
type Consumer struct {
    brokers []*Broker
    topic  []string
    ch     chan string
}

// NewConsumer initializes a new Consumer instance and registers it with the broker.
func NewConsumer(brokers []*Broker, topics []string, bufferSize int) *Consumer {
	ch := make(chan string, bufferSize)
    consumer := &Consumer{
		brokers: brokers,
		topic:   topics,
		ch:      ch,
	}

	for _, broker := range brokers {
		for _, topic := range topics {
			broker.RegisterConsumer(topic, ch)
		}
	}

	return consumer
}

// Start begins listening for messages on the consumer's channel.
// The handleFunc parameter is a callback function that processes each message.
func (c *Consumer) Start(handleFunc func(string)) {
	go func() {
        for msg := range c.ch {
            handleFunc(msg)
        }
    }()
}
