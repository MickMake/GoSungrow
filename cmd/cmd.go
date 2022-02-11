package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command {
	Use:              DefaultBinaryName,
	Short:            fmt.Sprintf("%s - Manage an SunGro instance", DefaultBinaryName),
	Long:             fmt.Sprintf("%s - Manage an SunGro instance", DefaultBinaryName),
	Run:              gbRootFunc,
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initConfig(cmd)
	},
}
func gbRootFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			_ = cmd.Help()
			break
		}
	}

}
