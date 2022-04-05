package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
    "strings"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func consume(competitionName string) error {
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
    // Set up a channel for handling Ctrl-C, etc
    sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

    // Process messages
    run := true
    const readDelay = 100
    for run {
        select {
        case sig := <-sigchan:
            fmt.Printf("Caught signal %v: terminating\n", sig)
            run = false
        default:
            ev, err := c.ReadMessage(readDelay * time.Millisecond)
            if err != nil {
                // Errors are informational and automatically handled by the consumer
                continue
            }
            fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
                *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
        }
    }

    c.Close()

    return nil
}
