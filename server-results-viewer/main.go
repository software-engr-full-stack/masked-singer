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

    serveGetVotes()
}

func serveGetVotes() {
    port := 9092
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
}
