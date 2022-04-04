package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
    "strings"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/google/uuid"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s <name-of-competition> <name-of-singer>\n",
            os.Args[0])
        os.Exit(1)
    }
    competitionName := os.Args[1]
    config := NewConfig(competitionName)

    rand.Seed(time.Now().UnixNano())

    p, err := kafka.NewProducer(&config.Kafka)

    if err != nil {
        panic(fmt.Errorf("Failed to create producer: %s", err))
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

    singerName := os.Args[2]
    singerNameTrimmed := strings.TrimSpace(singerName)
    if singerNameTrimmed == "" {
        panic("singer name must not be blank")
    }

    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{
            Topic: &config.User.CompetitionName, Partition: kafka.PartitionAny,
        },
        Key:            []byte(singerNameTrimmed),
        Value:          []byte(uuid.New().String()),
    }, nil)

    // Wait for all messages to be delivered
    p.Flush(15 * 1000)
    p.Close()
}
