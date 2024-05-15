package rudrya

import "sync"

type Broker struct{
	topics map[string][]chan string
	mu sync.RWMutex
}

func NewBroker() *Broker{
	return &Broker{
		topics: make(map[string][]chan string),
	}
}

func (b *Broker) RegisterConsumer(topic string,consumer chan string ){
	b.mu.Lock();
	defer b.mu.Unlock();

	if _,exists := b.topics[topic];!exists {
        b.topics[topic] = []chan string{}
    }
	b.topics[topic] = append(b.topics[topic], consumer)
}

func (b *Broker) SendMessage(topic, message string) {
    b.mu.RLock()
    defer b.mu.RUnlock()

    if consumers, exists := b.topics[topic]; exists {
        for _, consumer := range consumers {
            consumer <- message
        }
    }
}