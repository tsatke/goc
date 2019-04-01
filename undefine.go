package goc

import (
	"fmt"
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
	Long:  `TODO`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		Prompt := Cfg.GetBool("cmd.undefine.prompt")
		if Prompt {
			panic("TODO")
		}

		CmdOutputDirectory := Cfg.GetString("cmd.output.directory")

		err := os.Remove(path.Join(CmdOutputDirectory, args[0]))
		if err != nil {
			return fmt.Errorf("undefine: %v", err)
		}

		fmt.Printf("Done, the command %v has been removed.\n", args[0])
		return nil
	},
}
