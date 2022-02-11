package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
)


// ******************************************************************************** //
var cmdMqtt = &cobra.Command{
	Use:                   "mqtt",
	//Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("All things MQTT related."),
	Long:                  fmt.Sprintf("All things MQTT related."),
	Example:               PrintExamples("mqtt", "sync", "sync all"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdMqttFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdMqttFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				Cmd.Error = cmd.Help()

			case args[0] == "all":
				// Cmd.Error = Cmd.GoogleUpdate()

			default:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()
		}
	}
}


// ******************************************************************************** //
var cmdMqttSync = &cobra.Command{
	Use:                   "update",
	//Aliases:               []string{"refresh"},
	Short:                 fmt.Sprintf("Sync to an MQTT broker."),
	Long:                  fmt.Sprintf("Sync to an MQTT broker."),
	Example:               PrintExamples("mqtt sync", "", "all"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdMqttSyncFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdMqttSyncFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				Cmd.Error = cmd.Help()

			case args[0] == "all":
				// Cmd.Error = Cmd.GoogleUpdate()

			default:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()
		}
	}
}
