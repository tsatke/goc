package goc

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undefineCmd)
}

var undefineCmd = &cobra.Command{
	Use:   "undefine",
	Short: "Undefine removes a created command.",
	Long:  "Undefine deletes the command file in the defined output path, and will prompt the operation with respect to the setting \"cmd.undefine.prompt\".",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		undefine(args...)
	},
}

func undefine(args ...string) {
	CmdOutputDirectory := Cfg.GetString("cmd.output.directory")

	path := path.Join(CmdOutputDirectory, args[0])
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Printf("undefine: %v\n", err)
		return
	}

	shouldPrompt := Cfg.GetBool("cmd.undefine.prompt")
	if shouldPrompt && !Prompt("Really? (y/n) ") {
		return
	}

	err := os.Remove(path)
	if err != nil {
		Printf("undefine: %v\n", err)
		return
	}

	Printf("Done, the command %v has been removed.\n", args[0])
}
