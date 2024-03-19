package main

import (
    "log"
    "net/http"
    "github.com/filmons/chiapi/handlers"
    "github.com/filmons/chiapi/db"
)

func main() {
    // Initialize the database

    if err := db.InitDB("root:dbfilmon@tcp(localhost:3306)/flask"); err != nil {

        log.Fatal(err)
    }

    r := handlers.Routes()

    log.Println("Starting server on :3000")
    http.ListenAndServe(":3000", r)
}
