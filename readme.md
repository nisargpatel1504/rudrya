# Messaging System Library

A simple messaging system library in Go for implementing topic-based messaging with producers and consumers. This library allows you to create brokers to manage topics, send messages to topics using producers, and receive messages from topics using consumers.

## Features

- **Topic-Based Messaging**: Organize messages by topics.
- **Producers**: Send messages to specific topics.
- **Consumers**: Register to receive messages from specific topics.
- **Concurrency**: Efficient handling of multiple producers and consumers concurrently.
- **Thread-Safe**: Safe access to shared data using mutexes.

## Installation

To use this library in your Go project, you can either import it from a remote repository or use it locally.

### Import from Remote Repository

1. Initialize a Go module for your project if you haven't already:

    ```bash
    go mod init github.com/yourusername/yourproject
    ```

2. Import the library in your project:

    ```go
    import "github.com/yourusername/messaging"
    ```

3. Run `go mod tidy` to install the library:

    ```bash
    go mod tidy
    ```

### Use Locally

1. Clone the library repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/messaging.git
    ```

2. Use the `replace` directive in your `go.mod` file to reference the local path:

    ```go
    module github.com/yourusername/yourproject

    go 1.18

    replace github.com/yourusername/messaging => ../messaging

    require github.com/yourusername/messaging v0.0.0
    ```

3. Run `go mod tidy` to install the library:

    ```bash
    go mod tidy
    ```

## Usage

Here's an example of how to use the messaging library in your project.

### Example

Create a new file `main.go` in your project directory:

```go
package main

import (
    "fmt"
    "time"
    "github.com/yourusername/messaging"
)

func main() {
    // Initialize a new broker
    broker := messaging.NewBroker()

    // Create a producer associated with the broker
    producer := messaging.NewProducer(broker)

    // Create two consumers associated with the broker and topic "topic1"
    consumer1 := messaging.NewConsumer(broker, "topic1")
    consumer2 := messaging.NewConsumer(broker, "topic1")

    // Start the consumers in separate goroutines
    go consumer1.Start(func(message string) {
        fmt.Printf("Consumer 1 received message: %s\n", message)
    })

    go consumer2.Start(func(message string) {
        fmt.Printf("Consumer 2 received message: %s\n", message)
    })

    // Send messages to "topic1" through the producer
    producer.SendMessage("topic1", "Hello, World!")
    producer.SendMessage("topic1", "Another message")

    // Give the consumers some time to process the messages
    time.Sleep(2 * time.Second)
}
