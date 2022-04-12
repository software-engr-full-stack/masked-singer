package main

import (
    "fmt"
    "os"
    "time"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumeType struct {
    Topic string `json:"topic"`
    Key string `json:"key"`
    Value string `json:"value"`
}

func consume(competitionName string, data chan<- ConsumeType, closeConsumer chan os.Signal) error {
    config, err := NewConfig(competitionName)
    if err != nil {
        return err
    }

    kconfig := config.Kafka
    kconfig["group.id"] = config.User.GroupID
    kconfig["auto.offset.reset"] = "earliest"

    c, err := kafka.NewConsumer(&kconfig)
    if err != nil {
        return fmt.Errorf("failed to create consumer: %s", err)
    }

    err = c.SubscribeTopics([]string{config.User.CompetitionName}, nil)
    if err != nil {
        return err
    }

    // Process messages
    run := true
    const readDelay = 100
    for run {
        select {
        case sig := <-closeConsumer:
            fmt.Printf("Caught signal %#v: terminating\n", sig)
            run = false
        default:
            ev, err := c.ReadMessage(readDelay * time.Millisecond)
            if err != nil {
                // Errors are informational and automatically handled by the consumer
                continue
            }

            data <- ConsumeType{
                Topic: *ev.TopicPartition.Topic,
                Key: string(ev.Key),
                Value: string(ev.Value),
            }
        }
    }
    c.Close()
    return nil
}
