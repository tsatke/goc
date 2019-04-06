package goc

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	prefsCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Prints a preference by key.",
	Long:  "Given a preference key, this command prints the key and its current value.",
	Args:  cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		return prefsGet(args...)
	},
}

func prefsGet(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Need exactly one argument, got %v", len(args))
	}

	Printf("%-10v = %-10v\n", args[0], Cfg.Get(args[0]))
	return nil
}
