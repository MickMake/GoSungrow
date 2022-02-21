package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/defaults"
	"fmt"
	"github.com/spf13/cobra"
)


// ******************************************************************************** //
var cmdVersion = &cobra.Command{
	Use:                   "version",
	Short:                 "Version info.",
	Long:                  "Version info.",
	Example:               PrintExamples("version", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdVersionFunc,
	Args:                  cobra.RangeArgs(0, 0),
}
//goland:noinspection GoUnusedParameter
func cmdVersionFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Printf("%s v%s\n", DefaultBinaryName, defaults.Version)
	}
}
