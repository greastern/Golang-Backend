// internal/monitor/monitor.go
package monitor

import (
    "io/ioutil"
    "strconv"
    "strings"
    "time"

    "github.com/greastern/Golang-Backend/internal/websocket"
)

type Monitor struct {
    hub              *websocket.Hub
    lastNetworkStats NetworkStats
}

type SystemStats struct {
    CPU     CPUStats    `json:"cpu"`
    Memory  MemoryStats `json:"memory"`
    GPU     GPUStats    `json:"gpu"`
    Network NetworkStats `json:"network"`
    Power   PowerStats  `json:"power"`
    Fan     FanStats    `json:"fan"`
    Uptime  string      `json:"uptime"`
}

func NewMonitor(hub *websocket.Hub) *Monitor {
    return &Monitor{
        hub: hub,
    }
}

func (m *Monitor) Start() {
    for {
        stats := m.collectStats()
        m.hub.BroadcastJSON(stats)
        time.Sleep(time.Second) // Add a small delay to prevent excessive CPU usage
    }
}

func (m *Monitor) collectStats() SystemStats {
    return SystemStats{
        CPU:     m.getCPUStats(),
        Memory:  m.getMemoryStats(),
        GPU:     m.getGPUStats(),
        Network: m.getNetworkStats(),
        Power:   m.getPowerStats(),
        Fan:     m.getFanStats(),
        Uptime:  m.getUptime(),
    }
}

func (m *Monitor) getUptime() string {
    contents, err := ioutil.ReadFile("/proc/uptime")
    if err != nil {
        return ""
    }
    fields := strings.Fields(string(contents))
    if len(fields) < 1 {
        return ""
    }
    uptime, err := strconv.ParseFloat(fields[0], 64)
    if err != nil {
        return ""
    }
    duration := time.Duration(uptime) * time.Second
    return duration.String()
}