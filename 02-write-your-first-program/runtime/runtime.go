package runtime

import "runtime"

// GoVersion returns the Go tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "go1.3".
func GoVersion() string {
	return runtime.Version()
}
