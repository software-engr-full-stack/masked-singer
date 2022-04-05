package main

import (
    "net/http"
    "os"
    "log"
    "fmt"
)

func main() {
    if len(os.Args) < 3 {
        panic("must pass at least one arg, 'produce' or 'consume'")
    }

    switch action := os.Args[1]; action {
    case "serve":
        port := os.Args[2]
        http.HandleFunc("/vote", vote)
        http.HandleFunc("/get-votes", getVotes)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

    case "vote":
        competitionName := os.Args[2]
        singerName := os.Args[3]

        err := produce(competitionName, singerName)
        if err != nil {
            panic(err)
        }

    case "get-votes":
        competitionName := os.Args[2]

        data := make(chan ConsumeType)
        err := consume(competitionName, data)
        if err != nil {
            panic(err)
        }
    default:
        panic(fmt.Errorf("invalid action %#v", action))
    }
}
