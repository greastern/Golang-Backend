package stats

import (
    "strings"
)

func GetPiStats() string {
    sb := strings.Builder{}
    sb.WriteString(GetUptime())
    sb.WriteString(GetCPUStats())
    sb.WriteString(GetMemStats())
    sb.WriteString(GetGPUTemp())
    sb.WriteString(GetNetworkStats())
    sb.WriteString(GetPowerStats())
    sb.WriteString(GetFanStats())

    return sb.String()
}
