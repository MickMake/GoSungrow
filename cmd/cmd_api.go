package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/sungro/AppService/getPowerDevicePointNames"
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
	Short:                 fmt.Sprintf("List iSolarCloud api endpoints/areas"),
	Long:                  fmt.Sprintf("List iSolarCloud api endpoints/areas"),
	Example:               PrintExamples("api ls", "", "areas", "endpoints", "<area name>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiListFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdApiListFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
			case len(args) == 0:
				fmt.Println("Unknown sub-command.")
				_ = cmd.Help()

			case args[0] == "endpoints":
				Cmd.Error = SunGro.ListEndpoints("")

			case args[0] == "areas":
				SunGro.ListAreas()

			default:
				Cmd.Error = SunGro.ListEndpoints(args[0])
		}
	}
}


// ******************************************************************************** //
var cmdApiGet = &cobra.Command {
	Use:                   "get",
	//Aliases:               []string{""},
	Short:                 fmt.Sprintf("Get details from iSolarCloud"),
	Long:                  fmt.Sprintf("Get details from iSolarCloud"),
	Example:               PrintExamples("api get", "<endpoint> [area]"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiGetFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdApiGetFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		hey1 := SunGro.GetEndpoint(args[0])
		if hey1.IsError() {
			Cmd.Error = hey1.GetError()
			break
		}

		hey1 = hey1.SetRequest(getPowerDevicePointNames.RequestData{ DeviceType: "" })
		if hey1.IsError() {
			fmt.Println(hey1.Help())
			Cmd.Error = hey1.GetError()
			break
		}

		hey1 = hey1.Call()
		if hey1.IsError() {
			hey1.Help()
			Cmd.Error = hey1.GetError()
			break
		}

		fmt.Printf("EndPoint: %v\n", hey1)
		fmt.Printf("EndPoint: %v\n", hey1.GetUrl())
		// fmt.Printf("Request: %s\n", hey1.RequestString())
		// fmt.Printf("EndPoint: %s\n", hey1.ResponseString())

		// hey1 = hey1.Call()
		// if hey1.IsError() {
		// 	Cmd.Error = hey1.GetError()
		// 	break
		// }
		// fmt.Printf("resp: %v\n", hey1)
		//
		//
		// fmt.Printf("HEY:%v\n", hey1)
		// fmt.Printf("HEY:%v\n", hey1.GetError())
		// fmt.Printf("HEY:%s\n", hey1.GetRequestJson())
		// fmt.Printf("HEY:%s\n", hey1.GetData())
		// if Cmd.Error != nil {
		// 	break
		// }
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
