package rudrya

import (
	"sync"
)

// Broker manages the topics and their subscribers.
type Broker struct {
	topics map[string][]chan string
	mu     sync.RWMutex
}

// NewBroker initializes a new Broker instance.
func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string][]chan string),
	}
}

// CreateTopic initializes a new topic with a buffered channel.
func (b *Broker) CreateTopic(topic string, bufferSize int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.topics[topic]; !exists {
		b.topics[topic] = []chan string{}
	}
}

func (b *Broker) RegisterConsumer(topic string, consumer chan string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.topics[topic]; exists {
		b.topics[topic] = append(b.topics[topic], consumer)
	}
}
// SendMessage sends a message to a specific topic's channel.
func (b *Broker) SendMessage(topic, message string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if consumers, exists := b.topics[topic]; exists {
		for _, consumer := range consumers {
			consumer <- message
		}
	}
}

// GetChannel returns the channel for a specific topic.
// func (b *Broker) GetChannel(topic string) (chan string, bool) {
// 	b.mu.RLock()
// 	defer b.mu.RUnlock()

// 	ch, exists := b.topics[topic]
// 	return ch, exists
// }
