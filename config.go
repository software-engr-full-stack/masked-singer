package main

import (
    "strings"
    _ "embed"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Config struct {
    Kafka kafka.ConfigMap
    User map[string]interface{}
}

//go:embed kafka.properties
var kafkaProperties []byte

//go:embed config.sh
var config []byte

func NewConfig() Config {
    kconfig := make(map[string]kafka.ConfigValue)

    lines := strings.Split(string(kafkaProperties), "\n")
    for _, line := range lines {
        if !strings.HasPrefix(line, "#") && strings.TrimSpace(line) != "" {
            kv := strings.Split(line, "=")
            parameter := strings.TrimSpace(kv[0])
            value := strings.TrimSpace(kv[1])
            kconfig[parameter] = value
        }
    }

    lines = strings.Split(string(config), "\n")
    user := make(map[string]interface{})
    for _, line := range lines {
        if !strings.HasPrefix(line, "#") && strings.TrimSpace(line) != "" {
            kv := strings.Split(line, "=")
            parameter := strings.TrimSpace(kv[0])
            value := strings.TrimSpace(kv[1])
            user[parameter] = value
        }
    }

    return Config{
        Kafka: kconfig,
        User: user,
    }
}
