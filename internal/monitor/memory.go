package monitor

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type MemoryStats struct {
    Total     uint64  `json:"total"`
    Used      uint64  `json:"used"`
    Free      uint64  `json:"free"`
    UsagePerc float64 `json:"usagePerc"`
}

func (m *Monitor) getMemoryStats() MemoryStats {
    contents, err := ioutil.ReadFile("/proc/meminfo")
    if err != nil {
        return MemoryStats{}
    }

    lines := strings.Split(string(contents), "\n")
    memInfo := make(map[string]uint64)

    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) < 2 {
            continue
        }
        key := strings.TrimSuffix(fields[0], ":")
        value, err := strconv.ParseUint(fields[1], 10, 64)
        if err != nil {
            continue
        }
        memInfo[key] = value
    }

    total := memInfo["MemTotal"]
    free := memInfo["MemFree"]
    buffers := memInfo["Buffers"]
    cached := memInfo["Cached"]

    used := total - free - buffers - cached
    usagePerc := float64(used) / float64(total) * 100

    return MemoryStats{
        Total:     total * 1024, // Convert to bytes
        Used:      used * 1024,  // Convert to bytes
        Free:      free * 1024,  // Convert to bytes
        UsagePerc: usagePerc,
    }
}