package monitor

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type FanStats struct {
    Speed int `json:"speed"`
}

func (m *Monitor) getFanStats() FanStats {
    return FanStats{
        Speed: m.getFanSpeed(),
    }
}

func (m *Monitor) getFanSpeed() int {
    // This path might need to be adjusted based on your specific setup
    contents, err := ioutil.ReadFile("/sys/class/thermal/cooling_device0/cur_state")
    if err != nil {
        return 0
    }
    speed, err := strconv.Atoi(strings.TrimSpace(string(contents)))
    if err != nil {
        return 0
    }
    return speed
}