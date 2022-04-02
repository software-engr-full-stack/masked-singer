package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/google/uuid"
)

func main() {
    fmt.Printf("... %#v\n", os.Args)
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
            os.Args[0])
        os.Exit(1)
    }
    config := NewConfig()

    rand.Seed(time.Now().UnixNano())

    fmt.Printf("... %#v %#v\n", config, rand.Float32())

    topic := config.User["topic_name"].(string)
    p, err := kafka.NewProducer(&config.Kafka)

    if err != nil {
        fmt.Printf("Failed to create producer: %s", err)
        os.Exit(1)
    }

    // Go-routine to handle message delivery reports and
    // possibly other event types (errors, stats, etc)
    go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
                } else {
                    fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
                        *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
                }
            }
        }
    }()

    singerName := os.Args[1]
    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key:            []byte(singerName),
        Value:          []byte(uuid.New().String()),
    }, nil)

    // Wait for all messages to be delivered
    p.Flush(15 * 1000)
    p.Close()
}
