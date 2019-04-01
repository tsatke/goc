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
	Short: "TODO",
	Long:  "TODO",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Need exactly two arguments, got %v", len(args))
		}

		Cfg.Set(args[0], args[1])
		err := Cfg.WriteConfig()
		if err != nil {
			return fmt.Errorf("set: %v", err)
		}

		return nil
	},
}
