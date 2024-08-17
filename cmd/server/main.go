// cmd/server/main.go
package main

import (
    "log"
    "net/http"
    "github.com/greastern/Golang-Backend/internal/monitor"
    "github.com/greastern/Golang-Backend/internal/websocket"
)

func main() {
    hub := websocket.NewHub()
    go hub.Run()

    monitor := monitor.NewMonitor(hub)
    go monitor.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        websocket.ServeWs(hub, w, r)
    })

    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}