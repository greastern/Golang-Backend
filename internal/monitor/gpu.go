package monitor

import (
    "os/exec"
    "strconv"
    "strings"
)

type GPUStats struct {
    Usage       float64 `json:"usage"`
    Temperature float64 `json:"temperature"`
}

func (m *Monitor) getGPUStats() GPUStats {
    return GPUStats{
        Usage:       m.getGPUUsage(),
        Temperature: m.getGPUTemperature(),
    }
}

func (m *Monitor) getGPUUsage() float64 {
    cmd := exec.Command("vcgencmd", "get_mem", "gpu")
    output, err := cmd.Output()
    if err != nil {
        return 0
    }
    parts := strings.Split(string(output), "=")
    if len(parts) != 2 {
        return 0
    }
    usage, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
    if err != nil {
        return 0
    }
    return usage
}

func (m *Monitor) getGPUTemperature() float64 {
    cmd := exec.Command("vcgencmd", "measure_temp")
    output, err := cmd.Output()
    if err != nil {
        return 0
    }
    parts := strings.Split(string(output), "=")
    if len(parts) != 2 {
        return 0
    }
    temp, err := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(parts[1]), "'C"), 64)
    if err != nil {
        return 0
    }
    return temp
}