package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        panic("must pass at least one arg, 'produce' or 'consume'")
    }

    competitionName := os.Args[2]
    singerName := os.Args[3]

    switch op := os.Args[1]; op {
    case "vote":
        err := produce(competitionName, singerName)
        if err != nil {
            panic(err)
        }

    case "get-votes":
        err := consume(competitionName)
        if err != nil {
            panic(err)
        }
    default:
        panic(fmt.Errorf("invalid op %#v", op))
    }
}
