package goc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Cfg *viper.Viper
)

func init() {
	rootCmd.AddCommand(prefsCmd)

	Cfg = loadCfg()
}

var prefsCmd = &cobra.Command{
	Use:     "prefs",
	Aliases: []string{"preferences", "settings", "opts", "options"},
	Short:   "Version prints the version information.",
	Long:    "Version prints the version information about the build that you are using.",
}

func loadCfg() *viper.Viper {
	cfg := viper.New()

	cfg.SetDefault("cmd.output.directory", os.Getenv("HOME")+"/Development/tools")
	cfg.SetDefault("cmd.define.editor", "vim")
	cfg.SetDefault("cmd.undefine.prompt", true)

	cfg.SetConfigFile("define.yaml")

	err := cfg.ReadInConfig()
	if err != nil {
		fmt.Printf("Error while reading: %v\n", err)
	}

	err = cfg.WriteConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
