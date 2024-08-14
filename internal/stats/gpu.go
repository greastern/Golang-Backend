package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
	"strings"
)

func GetGPUTemp() string {
    temps, err := host.SensorsTemperatures()
    if err != nil {
        return "Error retrieving GPU temperature"
    }

    gpuInfo := ""
    for _, t := range temps {
        if strings.Contains(t.SensorKey, "gpu") {
            gpuInfo += fmt.Sprintf("GPU Temp: %.2f°C\n", t.Temperature)
        }
    }
    return gpuInfo
}
