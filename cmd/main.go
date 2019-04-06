package main

import "gitlab.com/TimSatke/goc"

var (
	// Version information that can be set during the build with
	// the build option
	//
	//	-ldflags "-X main.version=1.0.0"
	version string
)

func main() {
	if version != "" {
		goc.Version = version
	}
	goc.Execute()
}
