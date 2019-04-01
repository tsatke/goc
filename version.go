package goc

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	Version    = "99.9.9-developer-preview"
	versionFmt = `go-command (goc) v%v
(c) 2019-today Tim-Philipp Satke
`
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Short:   "Version prints the version information.",
	Long:    "Version prints the version information about the build that you are using.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionFmt, Version)
	},
}
