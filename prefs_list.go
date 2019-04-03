package goc

import (
	"sort"

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
		prefsList(args...)
	},
}

func prefsList(args ...string) {
	Println("Listing all properties")

	allKeys := Cfg.AllKeys()
	sort.Strings(allKeys)

	for _, k := range allKeys {
		Printf("%-10v = %-10v\n", k, Cfg.Get(k))
	}
}
