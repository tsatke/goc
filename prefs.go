package goc

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Cfg is the globally used configuration.
	Cfg     *viper.Viper
	CfgPath string
)

func init() {
	rootCmd.AddCommand(prefsCmd)

	Cfg = loadCfg()
}

var prefsCmd = &cobra.Command{
	Use:     "prefs",
	Aliases: []string{"preferences", "settings", "opts", "options"},
	Short:   "The base command for preference commands.",
}

func loadCfg() *viper.Viper {
	cfg := viper.New()

	cfg.SetDefault("cmd.output.directory", os.Getenv("HOME")+"/Development/tools")
	cfg.SetDefault("cmd.define.editor", "vim")
	cfg.SetDefault("cmd.undefine.prompt", true)
	cfg.SetDefault("cmd.prefs.clear.prompt", true)

	wd, err := os.Executable()
	if err != nil {
		panic(err)
	}

	CfgPath = path.Join(path.Dir(wd), "define.yaml")

	cfg.SetConfigFile(CfgPath)

	err = cfg.ReadInConfig()
	if err != nil {
		Println("It seems like there is no configuration file yet. It will be created.")
	}

	err = cfg.WriteConfigAs(CfgPath)
	if err != nil {
		panic(err)
	}

	return cfg
}
