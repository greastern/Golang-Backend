package handlers

import (
    "log"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
    "github.com/greastern/Golang-Backend/internal/platform"
    "github.com/greastern/Golang-Backend/internal/stats"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    platformType := platform.GetPlatform()
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    var prevStats string

    for {
        select {
        case <-ticker.C:
            var statsData string

            switch platformType {
            case "linux":
                statsData = stats.GetPiStats()
            case "darwin":
                statsData = "Running on macOS"
            case "ios":
                statsData = "Running on iOS"
            default:
                statsData = "Unknown platform"
            }

            if statsData != prevStats {
                err := ws.WriteMessage(websocket.TextMessage, []byte(statsData))
                if err != nil {
                    log.Println("write:", err)
                    return
                }
                prevStats = statsData
            }

        default:
            _, _, err := ws.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
        }
    }
}
