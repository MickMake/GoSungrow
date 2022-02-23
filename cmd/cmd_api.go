package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/login"
	"fmt"
	"github.com/spf13/cobra"
)

//
// ********************************************************************************
//
var cmdApi = &cobra.Command{
	Use:                   "api",
	Aliases:               []string{},
	Short:                 fmt.Sprintf("Interact with the SunGrow api."),
	Long:                  fmt.Sprintf("Interact with the SunGrow api."),
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

//
// ********************************************************************************
//
var cmdApiList = &cobra.Command{
	Use:                   "ls",
	Aliases:               []string{"list"},
	Short:                 fmt.Sprintf("List iSolarCloud api endpoints/areas"),
	Long:                  fmt.Sprintf("List iSolarCloud api endpoints/areas"),
	Example:               PrintExamples("api ls", "", "areas", "endpoints", "<area name>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
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
			Cmd.Error = SunGrow.ListEndpoints("")

		case args[0] == "areas":
			SunGrow.ListAreas()

		default:
			Cmd.Error = SunGrow.ListEndpoints(args[0])
		}
	}
}

//
// ********************************************************************************
//
var cmdApiLogin = &cobra.Command{
	Use: "login",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Login to iSolarCloud"),
	Long:                  fmt.Sprintf("Login to iSolarCloud"),
	Example:               PrintExamples("api login", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiLoginFunc,
	Args:                  cobra.MinimumNArgs(0),
}

//goland:noinspection GoUnusedParameter
func cmdApiLoginFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGrow.Login(login.SunGrowAuth{
			AppKey:       Cmd.ApiAppKey,
			UserAccount:  Cmd.ApiUsername,
			UserPassword: Cmd.ApiPassword,
			TokenFile:    Cmd.ApiTokenFile,
			Force:        true,
		})
		if Cmd.Error != nil {
			break
		}

		SunGrow.Auth.Print()

		if SunGrow.HasTokenChanged() {
			Cmd.ApiLastLogin = SunGrow.GetLastLogin()
			Cmd.ApiToken = SunGrow.GetToken()
			Cmd.Error = writeConfig()
		}
	}
}

//
// ********************************************************************************
//
var cmdApiGet = &cobra.Command{
	Use: "get",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Get details from iSolarCloud"),
	Long:                  fmt.Sprintf("Get details from iSolarCloud"),
	Example:               PrintExamples("api get", "<endpoint> [area]"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdApiGetFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdApiGetFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetJson()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = SunGrow.AllCritical()
			break
		}

		Cmd.Error = SunGrow.Get(args[0], args[1]).GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

//
// ********************************************************************************
//
var cmdApiRaw = &cobra.Command{
	Use: "raw",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Raw details from iSolarCloud"),
	Long:                  fmt.Sprintf("Raw details from iSolarCloud"),
	Example:               PrintExamples("api get", "<endpoint> [area]"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdApiRawFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdApiRawFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetRaw()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = SunGrow.AllCritical()
			break
		}

		Cmd.Error = SunGrow.Get(args[0], args[1]).GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

//
// ********************************************************************************
//
var cmdApiSave = &cobra.Command{
	Use: "save",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Save details from iSolarCloud"),
	Long:                  fmt.Sprintf("Save details from iSolarCloud"),
	Example:               PrintExamples("api save", "<endpoint> [area]"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdApiSaveFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdApiSaveFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetFile()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = SunGrow.AllCritical()
			break
		}

		Cmd.Error = SunGrow.Get(args[0], args[1]).GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

//
// ********************************************************************************
//
var cmdApiPut = &cobra.Command{
	Use:                   "put",
	Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Put details onto iSolarCloud"),
	Long:                  fmt.Sprintf("Put details onto iSolarCloud"),
	Example:               PrintExamples("api put", "<endpoint> <value>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdApiPutFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdApiPutFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// args = fillArray(1, args)
		// Cmd.Error = SunGrow.Init()
		// if Cmd.Error != nil {
		// 	break
		// }
	}
}
