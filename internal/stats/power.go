package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
	"strings"
)

func GetPowerStats() string {
    temps, err := host.SensorsTemperatures()
    if err != nil {
        return "Error retrieving power stats"
    }

    powerInfo := ""
    for _, t := range temps {
        if strings.Contains(t.SensorKey, "power") {
            powerInfo += fmt.Sprintf("Power Temp: %.2f°C\n", t.Temperature)
        }
    }
    return powerInfo
}
