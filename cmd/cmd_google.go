package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)

// ******************************************************************************** //
var cmdGoogle = &cobra.Command{
	Use: "google",
	// Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("Update and view Google sheets."),
	Long:                  fmt.Sprintf("Update and view Google sheets."),
	Example:               PrintExamples("google", "update all", "update devices"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGoogleFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdGoogleFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		return

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
var cmdGoogleSync = &cobra.Command{
	Use: "update",
	// Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("Update Google sheets."),
	Long:                  fmt.Sprintf("Update Google sheets."),
	Example:               PrintExamples("google update", "all", "device", "unit"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdGoogleSyncFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdGoogleSyncFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		return

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
