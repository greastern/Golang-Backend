package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
)

func GetUptime() string {
    uptime, err := host.Uptime()
    if err != nil {
        return "Error retrieving uptime"
    }
    return fmt.Sprintf("Uptime: %v seconds\n", uptime)
}
