package stats

import (
    "fmt"
    "github.com/shirou/gopsutil/mem"
)

func GetMemStats() string {
    v, err := mem.VirtualMemory()
    if err != nil {
        return "Error retrieving memory stats"
    }
    
    return fmt.Sprintf("Memory: Total: %v, Used: %v, Free: %v\n", v.Total, v.Used, v.Free)
}
