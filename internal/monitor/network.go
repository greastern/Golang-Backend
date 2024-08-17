package monitor

import (
    "io/ioutil"
    "strconv"
    "strings"
)

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

func (m *Monitor) getNetworkStats() NetworkStats {
    stats := getNetworkData()
    rxSpeedMbps := float64(stats.RxBytes-m.lastNetworkStats.RxBytes) * 8 / 1000000 // Mbps
    txSpeedMbps := float64(stats.TxBytes-m.lastNetworkStats.TxBytes) * 8 / 1000000 // Mbps

    stats.RxBytesHR = formatBytes(stats.RxBytes)
    stats.TxBytesHR = formatBytes(stats.TxBytes)
    stats.RxSpeedMbps = rxSpeedMbps
    stats.TxSpeedMbps = txSpeedMbps

    m.lastNetworkStats = stats
    return stats
}

func getNetworkData() NetworkStats {
    contents, err := ioutil.ReadFile("/proc/net/dev")
    if err != nil {
        return NetworkStats{}
    }

    lines := strings.Split(string(contents), "\n")
    var stats NetworkStats

    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) < 17 || (!strings.HasPrefix(fields[0], "eth") && !strings.HasPrefix(fields[0], "wlan")) {
            continue
        }
        rx, _ := strconv.ParseUint(fields[1], 10, 64)
        tx, _ := strconv.ParseUint(fields[9], 10, 64)
        rxPackets, _ := strconv.ParseUint(fields[2], 10, 64)
        txPackets, _ := strconv.ParseUint(fields[10], 10, 64)
        stats.RxBytes += rx
        stats.TxBytes += tx
        stats.RxPackets += rxPackets
        stats.TxPackets += txPackets
    }

    return stats
}

func formatBytes(bytes uint64) string {
    const unit = 1024
    if bytes < unit {
        return strconv.FormatUint(bytes, 10) + " B"
    }
    div, exp := uint64(unit), 0
    for n := bytes / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return strconv.FormatFloat(float64(bytes)/float64(div), 'f', 1, 64) + " " + []string{"KB", "MB", "GB", "TB"}[exp]
}