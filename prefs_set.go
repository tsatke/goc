package goc

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	prefsCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a preference key to a given value.",
	Long:  "Sets a preference to a given value. \"prefs set <key> <value>\" will set <key> to <value>. Changes instantly apply.",
	Args:  cobra.ExactArgs(2),
	RunE: func(_ *cobra.Command, args []string) error {
		return prefsSet(args...)
	},
}

func prefsSet(args ...string) error {
	Cfg.Set(args[0], args[1])
	err := Cfg.WriteConfig()
	if err != nil {
		return fmt.Errorf("set: %v", err)
	}

	Printf("Setting %v = %v\n", args[0], args[1])

	return nil
}
