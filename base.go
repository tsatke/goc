package goc

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goc",
	Short: "Cmd enables you to define or undefine custom commands.",
	Long:  `Cmd provides sub-commands such as define or undefine, that let you create or delete custom commands on the fly.`,
}

// Execute is the entry point for command execution.
// See spf13/cobra documentation for more information.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Println(err)
		os.Exit(1)
	}
}
