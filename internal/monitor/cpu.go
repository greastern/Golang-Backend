package monitor

import (
    "io/ioutil"
    "strconv"
    "strings"
    "time"
)

type CPUStats struct {
    Usage       float64 `json:"usage"`
    Temperature float64 `json:"temperature"`
}

func (m *Monitor) getCPUStats() CPUStats {
    return CPUStats{
        Usage:       m.getCPUUsage(),
        Temperature: m.getCPUTemperature(),
    }
}

func (m *Monitor) getCPUUsage() float64 {
    idle0, total0 := getCPUSample()
    time.Sleep(3 * time.Second)
    idle1, total1 := getCPUSample()

    idleTicks := float64(idle1 - idle0)
    totalTicks := float64(total1 - total0)
    cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

    return cpuUsage
}

func getCPUSample() (idle, total uint64) {
    contents, err := ioutil.ReadFile("/proc/stat")
    if err != nil {
        return
    }
    lines := strings.Split(string(contents), "\n")
    for _, line := range lines {
        fields := strings.Fields(line)
        if fields[0] == "cpu" {
            numFields := len(fields)
            for i := 1; i < numFields; i++ {
                val, _ := strconv.ParseUint(fields[i], 10, 64)
                total += val
                if i == 4 { // idle is the 5th field in the cpu line
                    idle = val
                }
            }
            return
        }
    }
    return
}

func (m *Monitor) getCPUTemperature() float64 {
    contents, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
    if err != nil {
        return 0
    }
    temp, err := strconv.ParseFloat(strings.TrimSpace(string(contents)), 64)
    if err != nil {
        return 0
    }
    return temp / 1000 // Convert millidegrees to degrees
}