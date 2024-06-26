// Package version provides honeycombo application version information.
package version

import (
	"runtime/debug"
)

// Version value is set by ldflags
var Version string //nolint:gochecknoglobals

// GetVersion return command version.
// Version global variable is set by ldflags.
func GetVersion() string {
	version := "unknown"
	if Version != "" {
		version = Version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		version = buildInfo.Main.Version
	}
	return version
}
