package main

import (
    "net/http"
    "os"
    "log"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        panic("must pass at least one arg, 'produce' or 'consume'")
    }

    switch action := os.Args[1]; action {
    case "serve":
        http.HandleFunc("/vote", serve)
        log.Fatal(http.ListenAndServe(":8082", nil))

    case "vote":
        competitionName := os.Args[2]
        singerName := os.Args[3]

        err := produce(competitionName, singerName)
        if err != nil {
            panic(err)
        }

    case "get-votes":
        competitionName := os.Args[2]

        err := consume(competitionName)
        if err != nil {
            panic(err)
        }
    default:
        panic(fmt.Errorf("invalid action %#v", action))
    }
}
