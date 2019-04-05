package goc

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	prefsCmd.AddCommand(clearCmd)
}

var clearCmd = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"reset", "rm", "cls", "remove"},
	Short:   "TODO",
	Long:    "TODO",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		prefsClear(args...)
	},
}

func prefsClear(args ...string) {
	shouldPrompt := Cfg.GetBool("cmd.prefs.clear.prompt")
	if shouldPrompt {
		if !Prompt("Really? (y/n) ") {
			return
		}
	}

	err := os.Remove(CfgPath)
	if err != nil {
		Printf("Unable to delete config file: %v", err)
	} else {
		Println("Configuration file removed, it will be re-created with default values upon next use.")
	}
}
