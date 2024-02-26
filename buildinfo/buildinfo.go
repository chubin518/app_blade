package buildinfo

import (
	"fmt"
	"runtime"
)

var (
	// Version of app blade
	Version string = "0.0.1"
	// Name of app blade
	Name string = "app_blade"
	// Build time of app blade
	BuildTime string = "2021-05-21 14:30:00"
	// Environment of app blade
	Env string = "dev"
)

func String() string {
	return fmt.Sprintf("Name=%s Env=%s Version=%s BuildTime=%s runtime=%s %s/%s", Name, Env, Version, BuildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
