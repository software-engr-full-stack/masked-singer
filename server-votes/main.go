package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type RequestType struct {
    CompetitionName string `json:"competition_name"`
    SingerName      string `json:"singer_name"`
}

func main() {
    port := 8082
    http.HandleFunc("/vote", vote)

    server := &http.Server{Addr: fmt.Sprintf(":%d", port)}
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    <-stop

    const delay = 10
    ctx, cancel := context.WithTimeout(context.Background(), delay*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Println(err)
        return
    }
    // TODO: GET /vote/?singer_name=batman&competition_name=contest-1

    competitionName := "contest-1"
    singerName := "michael"

    err := produce(competitionName, singerName)
    if err != nil {
        log.Println(err)
        return
    }
}

func vote(rw http.ResponseWriter, req *http.Request) {
    body, err := io.ReadAll(req.Body)
    if err != nil {
        log.Println(err)
    }

    var request RequestType
    err = json.Unmarshal(body, &request)
    if err != nil {
        log.Println(err)
    }

    err = produce(request.CompetitionName, request.SingerName)
    if err != nil {
        log.Println(err)
    }
}
