package main

import (
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
    "context"
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

        server := &http.Server{Addr: fmt.Sprintf(":%s", port)}

        go func() {
            if err := server.ListenAndServe(); err != http.ErrServerClosed {
                log.Fatal(err)
            }
        }()

        stop := make(chan os.Signal, 1)
        signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

        <-stop

        const delay = 10
        ctx, cancel := context.WithTimeout(context.Background(), delay * time.Second)
        defer cancel()
        if err := server.Shutdown(ctx); err != nil {
            log.Println(err)
            return
        }

    case "vote":
        competitionName := os.Args[2]
        singerName := os.Args[3]

        err := produce(competitionName, singerName)
        if err != nil {
            log.Println(err)
            return
        }

    case "get-votes":
        competitionName := os.Args[2]

        data := make(chan ConsumeType)

        closeConsumer := make(chan os.Signal, 1)
        signal.Notify(closeConsumer, syscall.SIGINT, syscall.SIGTERM)

        go func() {
            err := consume(competitionName, data, closeConsumer)
            if err != nil {
                log.Println(err)
                return
            }
            close(data)
        }()

        for item := range data {
            log.Println("DEBUG:", item)
        }
    default:
        log.Printf("invalid action %#v", action)
        return
    }
}
