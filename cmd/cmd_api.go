package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/spf13/cobra"
)


// ******************************************************************************** //
var cmdApi = &cobra.Command {
	Use:                   "api",
	Aliases:               []string{},
	Short:                 fmt.Sprintf("List SunGro areas."),
	Long:                  fmt.Sprintf("List SunGro areas."),
	Example:               PrintExamples("api", "get <endpoint>", "put <endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdApiFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			Cmd.Error = cmd.Help()
			break
		}
	}
}

// ******************************************************************************** //
var cmdApiList = &cobra.Command {
	Use:                   "ls",
	Aliases:               []string{"list"},
	Short:                 fmt.Sprintf("List api endpoints on iSolarCloud"),
	Long:                  fmt.Sprintf("List api endpoints on iSolarCloud"),
	Example:               PrintExamples("api ls", "", "<endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiListFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdApiListFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) > 0 {
			fmt.Println("Unknown sub-command.")
		}
		_ = cmd.Help()
	}
}

// ******************************************************************************** //
var cmdApiGet = &cobra.Command {
	Use:                   "get",
	//Aliases:               []string{""},
	Short:                 fmt.Sprintf("Get details from iSolarCloud"),
	Long:                  fmt.Sprintf("Get details from iSolarCloud"),
	Example:               PrintExamples("api get", "<endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiGetFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdApiGetFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		hey := SunGro.GetEndpoint("getPsList")
		fmt.Printf("HEY:%v\n", hey)
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdApiPut = &cobra.Command {
	Use:                   "put",
	Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Put details onto iSolarCloud"),
	Long:                  fmt.Sprintf("Put details onto iSolarCloud"),
	Example:               PrintExamples("api put", "<endpoint> <value>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiPutFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdApiPutFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.Init()
		if Cmd.Error != nil {
			break
		}
	}
}
