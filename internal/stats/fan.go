package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/host"
	"strings"
)

func GetFanStats() string {
    temps, err := host.SensorsTemperatures()
    if err != nil {
        return "Error retrieving fan stats"
    }

    fanInfo := ""
    for _, t := range temps {
        if strings.Contains(t.SensorKey, "fan") {
            fanInfo += fmt.Sprintf("Fan Speed: %.2f RPM\n", t.Temperature)
        }
    }
    return fanInfo
}
