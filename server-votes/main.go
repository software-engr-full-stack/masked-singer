package main

import (
    "os"
    "log"
)

func main() {

    if len(os.Args) < 2 {
        panic("must pass two arg, 'competition name' and 'singer name'")
    }

    competitionName := os.Args[1]
    singerName := os.Args[2]

    // action := os.Args[1]

    err := produce(competitionName, singerName)
    if err != nil {
        log.Println(err)
        return
    }
}
