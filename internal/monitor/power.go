package monitor

import (
    "os/exec"
    "strconv"
    "strings"
)

type PowerStats struct {
    Voltage     float64 `json:"voltage"`
    CurrentDraw float64 `json:"currentDraw"`
}

func (m *Monitor) getPowerStats() PowerStats {
    return PowerStats{
        Voltage:     m.getVoltage(),
        CurrentDraw: m.getCurrentDraw(),
    }
}

func (m *Monitor) getVoltage() float64 {
    cmd := exec.Command("vcgencmd", "measure_volts", "core")
    output, err := cmd.Output()
    if err != nil {
        return 0
    }
    parts := strings.Split(string(output), "=")
    if len(parts) != 2 {
        return 0
    }
    voltage, err := strconv.ParseFloat(strings.TrimSuffix(strings.TrimSpace(parts[1]), "V"), 64)
    if err != nil {
        return 0
    }
    return voltage
}

func (m *Monitor) getCurrentDraw() float64 {
    // Note: Raspberry Pi 5 might not provide direct current draw information.
    // This is a placeholder and may need to be updated based on available methods.
    return 0
}