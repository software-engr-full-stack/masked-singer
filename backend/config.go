package main

import (
    "strings"
    _ "embed"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Config struct {
    Kafka kafka.ConfigMap
    User UserConfigType
}

type UserConfigType struct {
    CompetitionName string
}

//go:embed kafka.properties
var kafkaProperties []byte

func NewConfig(competitionName string) Config {
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

    competitionNameTrimmed := strings.TrimSpace(competitionName)
    if competitionNameTrimmed == "" {
        panic("competition name must not be blank")
    }

    return Config{
        Kafka: kconfig,
        User: UserConfigType{
            CompetitionName: competitionNameTrimmed,
        },
    }
}
