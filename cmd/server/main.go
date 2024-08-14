package main

import (
    "log"
    "net/http"

    "github.com/greastern/Golang-Backend/internal/handlers"
)

func main() {
    http.HandleFunc("/ws", handlers.HandleConnections)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
