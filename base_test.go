package goc

import (
	"flag"

	"gitlab.com/TimSatke/gog"
)

var (
	record = flag.Bool("record", false, "Update all golden files")
)

func init() {
	flag.Parse()

	if *record {
		gog.Record()
	} else {
		gog.Replay()
	}
}
