package platform

import "runtime"

func GetPlatform() string {
    return runtime.GOOS
}
