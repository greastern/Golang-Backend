package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/net"
)

func GetNetworkStats() string {
    interfaces, err := net.IOCounters(true)
    if err != nil {
        return "Error retrieving network stats"
    }

    var netInfo string
    for _, iface := range interfaces {
        netInfo += fmt.Sprintf("Interface: %v\nBytes Sent: %v\nBytes Received: %v\n", iface.Name, iface.BytesSent, iface.BytesRecv)
    }
    return netInfo
}
