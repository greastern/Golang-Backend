package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/host"
	"strings"
)

func GetCPUStats() string {
    stats, err := cpu.Info()
    if err != nil {
        return "Error retrieving CPU stats"
    }
    
    temps, err := host.SensorsTemperatures()
    if err != nil {
        return "Error retrieving CPU temperature"
    }

    cpuInfo := fmt.Sprintf("CPU: %v %v Cores\n", stats[0].ModelName, stats[0].Cores)
    for _, t := range temps {
        if strings.Contains(t.SensorKey, "cpu") {
            cpuInfo += fmt.Sprintf("CPU Temp: %.2f°C\n", t.Temperature)
        }
    }
    return cpuInfo
}
