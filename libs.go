package goc

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(libCmd)
}

var libCmd = &cobra.Command{
	Use:     "libs",
	Aliases: []string{"lib"},
	Short:   "List third-party dependencies.",
	Long:    "Lists all used third-party libraries together with their respective license.",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		libs(args...)
	},
}

func libs(args ...string) {
	type Lib struct {
		Name    string
		Version string
		ScmLink string
		License string
	}

	const (
		LicenseApache20 = "Apache License Version 2.0"
		LicenseMIT      = "The MIT License (MIT)"
	)

	libraries := []Lib{
		{
			Name:    "spf13/cobra",
			Version: "v0.0.3",
			ScmLink: "https://github.com/spf13/cobra",
			License: LicenseApache20,
		},
		{
			Name:    "spf13/viper",
			Version: "v1.3.2",
			ScmLink: "https://github.com/spf13/viper",
			License: LicenseMIT,
		},
		{
			Name:    "stretchr/testify",
			Version: "v1.2.2",
			ScmLink: "https://github.com/stretchr/testify",
			License: LicenseMIT,
		},
		{
			Name:    "TimSatke/abc",
			Version: "v0.0.0",
			ScmLink: "https://gitlab.com/TimSatke/abc",
			License: LicenseMIT,
		},
		{
			Name:    "TimSatke/gog",
			Version: "v0.0.0",
			ScmLink: "https://gitlab.com/TimSatke/gog",
			License: "-", // TODO: gog does not have a license yet, update as soon as one is available
		},
	}

	for i, library := range libraries {
		format := `(%2d) %v %v (%v)
     License: %v

`
		Printf(format, i+1, library.Name, library.Version, library.ScmLink, library.License)
	}
}
