package stats

import (
    "os/exec"
)

func RebootSystem() error {
    return exec.Command("sudo", "reboot").Run()
}

func ShutdownSystem() error {
    return exec.Command("sudo", "shutdown", "now").Run()
}
