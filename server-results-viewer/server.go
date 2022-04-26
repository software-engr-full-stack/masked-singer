package main

import (
    "encoding/json"
    "net/http"
    "net/url"
    "io"
    "strings"
    "log"
    "os"
    "syscall"
    "fmt"

    "github.com/gorilla/websocket"
)

type RequestType struct {
    Action string `json:"action"` // "vote" or "get-votes"
    CompetitionName string `json:"competition_name"`
    SingerName string `json:"singer_name"`
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

func getVotes(rw http.ResponseWriter, req *http.Request) {
    query, err := url.ParseQuery(req.URL.RawQuery)
    if err != nil {
        log.Println(err)
    }

    competitionName := strings.TrimSpace(query["competition_name"][0])

    data := make(chan ConsumeType)
    closeConsumer := make(chan os.Signal)
    go func() {
        err = consume(competitionName, data, closeConsumer)
        if err != nil {
            log.Println(err)
        }
    }()

    var upgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }

    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(rw, req, nil)
    if err != nil {
        log.Println(err)
    }

    log.Println("Client Connected")

    for item := range data {
        marsh, err := json.Marshal(item)
        if err != nil {
            log.Println(err)
        }

        err = ws.WriteMessage(1, marsh)
        if err != nil {
            log.Println(err)
            // ws.Close()
            closeConsumer <- syscall.SIGTERM
            break
        }

        // // DEBUG: for curl test, put line breaks between responses
        // err = ws.WriteMessage(1, []byte("\n"))
        // if err != nil {
        //     log.Println(err)
        // }
    }

    // wsReceiver(ws)
}

func wsReceiver(conn *websocket.Conn) {
    for {
        // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        // print out that message for clarity
        fmt.Println("DEBUG, received from Websocket:", string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}
