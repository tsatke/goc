package goc

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	prefsCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "TODO",
	Long:    "TODO",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for _, k := range Cfg.AllKeys() {
			fmt.Printf("%-10v = %-10v\n", k, Cfg.Get(k))
		}
	},
}
