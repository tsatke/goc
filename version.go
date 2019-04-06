package goc

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	// Version information that can be set during the build with
	// the build option
	//
	//	-ldflags "-X goc.Version=1.0.0"
	Version    = "99.9.9-developer-preview"
	versionFmt = `(c) 2019-today Tim-Philipp Satke
go-command (goc) v%v
`
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Short:   "Version prints the version information.",
	Long:    "Version prints the version information about the build that you are using.",
	Args:    cobra.ExactArgs(0),
	Run: func(*cobra.Command, []string) {
		version()
	},
}

func version() {
	Printf(versionFmt, Version)
}
