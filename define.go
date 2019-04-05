package goc

import (
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(defineCmd)
}

var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "Define opens an editor to let you create your custom command.",
	Long:  `Define attempts to open vim to let you create your command. Vim must be installed.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return define(args...)
	},
}

func define(args ...string) error {
	CmdOutputDirectory := Cfg.GetString("cmd.output.directory")
	if _, err := os.Stat(CmdOutputDirectory); os.IsNotExist(err) {
		return err
	}

	command := exec.Command(Cfg.GetString("cmd.define.editor"), path.Join(CmdOutputDirectory, args[0]))
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		return err
	}

	command = exec.Command("chmod", "-R", "a+x", CmdOutputDirectory)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	err = command.Run()
	if err != nil {
		return err
	}

	Printf("Done, you can now use the command %v.\n", args[0])
	return nil
}
