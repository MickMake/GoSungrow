package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
)


// ******************************************************************************** //
var cmdGoogle = &cobra.Command{
	Use:                   "google",
	//Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("Update and view Google sheets."),
	Long:                  fmt.Sprintf("Update and view Google sheets."),
	Example:               PrintExamples("google", "update all", "update users"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGoogleFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdGoogleFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
		case len(args) == 0:
			Cmd.Error = cmd.Help()

		case args[0] == "all":
			Cmd.Error = Cmd.GoogleUpdate()

		default:
			fmt.Println("Unknown sub-command.")
			_ = cmd.Help()
		}
	}
}

// ******************************************************************************** //
var cmdGoogleUpdate = &cobra.Command{
	Use:                   "update",
	//Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("Update Google sheets."),
	Long:                  fmt.Sprintf("Update Google sheets."),
	Example:               PrintExamples("google update", "all", "presence", "user"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGoogleUpdateFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdGoogleUpdateFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
		case len(args) == 0:
			Cmd.Error = cmd.Help()

		case args[0] == "all":
			Cmd.Error = Cmd.GoogleUpdate()

		default:
			fmt.Println("Unknown sub-command.")
			_ = cmd.Help()
		}
	}
}
