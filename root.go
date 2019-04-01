package goc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goc",
	Short: "Cmd enables you to define or undefine custom commands.",
	Long:  `Cmd procides sub-commands such as define or undefine, that let you create or delete custom commands on the fly.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
