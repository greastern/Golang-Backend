package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/url"
    "os"
    "os/signal"

    "github.com/gorilla/websocket"
)

type SystemStats struct {
    CPU     CPUStats    `json:"cpu"`
    Memory  MemoryStats `json:"memory"`
    GPU     GPUStats    `json:"gpu"`
    Network NetworkStats `json:"network"`
    Power   PowerStats  `json:"power"`
    Fan     FanStats    `json:"fan"`
    Uptime  string      `json:"uptime"`
}

type CPUStats struct {
    Usage       float64 `json:"usage"`
    Temperature float64 `json:"temperature"`
}

type MemoryStats struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Free      uint64  `json:"free"`
    UsagePerc float64 `json:"usagePerc"`
}

type GPUStats struct {
    Usage       float64 `json:"usage"`
    Temperature float64 `json:"temperature"`
}

type NetworkStats struct {
    RxBytes     uint64  `json:"rxBytes"`
    TxBytes     uint64  `json:"txBytes"`
    RxPackets   uint64  `json:"rxPackets"`
    TxPackets   uint64  `json:"txPackets"`
    RxBytesHR   string  `json:"rxBytesHR"`
    TxBytesHR   string  `json:"txBytesHR"`
    RxSpeedMbps float64 `json:"rxSpeedMbps"`
    TxSpeedMbps float64 `json:"txSpeedMbps"`
}

type PowerStats struct {
    Voltage     float64 `json:"voltage"`
    CurrentDraw float64 `json:"currentDraw"`
}

type FanStats struct {
    Speed int `json:"speed"`
}

func main() {
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
    log.Printf("Connecting to %s", u.String())

    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            var stats SystemStats
            err = json.Unmarshal(message, &stats)
            if err != nil {
                log.Println("unmarshal:", err)
                continue
            }
            prettyJSON, _ := json.MarshalIndent(stats, "", "  ")
            fmt.Println(string(prettyJSON))
        }
    }()

    for {
        select {
        case <-done:
            return
        case <-interrupt:
            log.Println("interrupt")
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("write close:", err)
                return
            }
            select {
            case <-done:
            }
            return
        }
    }
}