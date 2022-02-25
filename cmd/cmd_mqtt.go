package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)


func AttachCmdMqtt(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdMqtt = &cobra.Command{
		Use:                   "mqtt",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("All things MQTT related."),
		Long:                  fmt.Sprintf("All things MQTT related."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdMqttFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmd.AddCommand(cmdMqtt)
	cmdMqtt.Example = PrintExamples(cmdMqtt, "sync", "sync all")


	// ******************************************************************************** //
	var cmdMqttSync = &cobra.Command{
		Use:                   "update",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Sync to an MQTT broker."),
		Long:                  fmt.Sprintf("Sync to an MQTT broker."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdMqttSyncFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdMqtt.AddCommand(cmdMqttSync)
	cmdMqttSync.Example = PrintExamples(cmdMqttSync, "", "all")

	return cmdMqtt
}


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
