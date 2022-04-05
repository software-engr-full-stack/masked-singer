package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "strings"
)

type RequestType struct {
    Action string `json:"action"` // "vote" or "get-votes"
    CompetitionName string `json:"competition_name"`
    SingerName string `json:"singer_name"`
}

func serve(rw http.ResponseWriter, req *http.Request) {
    body, err := io.ReadAll(req.Body)
    if err != nil {
        panic(err)
    }

    var request RequestType
    err = json.Unmarshal(body, &request)
    if err != nil {
        panic(err)
    }

    switch action := strings.TrimSpace(request.Action); action {
    case "vote":
        err := produce(request.CompetitionName, request.SingerName)
        if err != nil {
            panic(err)
        }

    case "get-votes":
        err := consume(request.CompetitionName)
        if err != nil {
            panic(err)
        }
    default:
        panic(fmt.Errorf("invalid action %#v", action))
    }
}
