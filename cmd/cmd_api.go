package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/AppService/login"
	"fmt"
	"github.com/spf13/cobra"
)


func AttachCmdApi(cmd *cobra.Command) *cobra.Command {
	// ******************************************************************************** //
	var cmdApi = &cobra.Command{
		Use:                   "api",
		Aliases:               []string{},
		Short:                 fmt.Sprintf("Interact with the SunGrow api."),
		Long:                  fmt.Sprintf("Interact with the SunGrow api."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdApiFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(cmdApi)
	cmdApi.Example = PrintExamples(cmdApi, "get <endpoint>", "put <endpoint>")

	cmdApi.PersistentFlags().StringVarP(&Cmd.ApiOutputType, flagApiOutputType, "o", "", fmt.Sprintf("Output type: 'json', 'raw', 'file'"))
	_ = cmdApi.PersistentFlags().MarkHidden(flagApiOutputType)


	// ******************************************************************************** //
	var cmdApiList = &cobra.Command{
		Use:                   "ls",
		Aliases:               []string{"list"},
		Short:                 fmt.Sprintf("List iSolarCloud api endpoints/areas"),
		Long:                  fmt.Sprintf("List iSolarCloud api endpoints/areas"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdApiListFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdApi.AddCommand(cmdApiList)
	cmdApiList.Example = PrintExamples(cmdApiList, "", "areas", "endpoints", "<area name>")

	// ******************************************************************************** //
	var cmdApiLogin = &cobra.Command{
		Use: "login",
		// Aliases:               []string{""},
		Short:                 fmt.Sprintf("Login to iSolarCloud"),
		Long:                  fmt.Sprintf("Login to iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdApiLoginFunc,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmdApi.AddCommand(cmdApiLogin)
	cmdApiLogin.Example = PrintExamples(cmdApiLogin, "")

	// ******************************************************************************** //
	var cmdApiGet = &cobra.Command{
		Use: "get",
		// Aliases:               []string{""},
		Short:                 fmt.Sprintf("Get details from iSolarCloud"),
		Long:                  fmt.Sprintf("Get details from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdApiGetFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdApi.AddCommand(cmdApiGet)
	cmdApiGet.Example = PrintExamples(cmdApiGet, "[area].<endpoint>")

	// ******************************************************************************** //
	var cmdApiRaw = &cobra.Command{
		Use: "raw",
		// Aliases:               []string{""},
		Short:                 fmt.Sprintf("Raw details from iSolarCloud"),
		Long:                  fmt.Sprintf("Raw details from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdApiRawFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdApi.AddCommand(cmdApiRaw)
	cmdApiRaw.Example = PrintExamples(cmdApiRaw, "[area].<endpoint>")

	// ******************************************************************************** //
	var cmdApiSave = &cobra.Command{
		Use: "save",
		// Aliases:               []string{""},
		Short:                 fmt.Sprintf("Save details from iSolarCloud"),
		Long:                  fmt.Sprintf("Save details from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdApiSaveFunc,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdApi.AddCommand(cmdApiSave)
	cmdApiSave.Example = PrintExamples(cmdApiSave, "[area].<endpoint>")

	// ******************************************************************************** //
	var cmdApiPut = &cobra.Command{
		Use:                   "put",
		Aliases:               []string{"write"},
		Short:                 fmt.Sprintf("Put details onto iSolarCloud"),
		Long:                  fmt.Sprintf("Put details onto iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdApiPutFunc,
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmdApi.AddCommand(cmdApiPut)
	cmdApiPut.Example = PrintExamples(cmdApiPut, "[area].<endpoint> <value>")

	return cmdApi
}


func cmdApiFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			Cmd.Error = cmd.Help()
			break
		}
	}
}

func cmdApiListFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		switch {
		case len(args) == 0:
			fmt.Println("Unknown sub-command.")
			_ = cmd.Help()

		case args[0] == "endpoints":
			Cmd.Error = Cmd.SunGrow.ListEndpoints("")

		case args[0] == "areas":
			Cmd.SunGrow.ListAreas()

		default:
			Cmd.Error = Cmd.SunGrow.ListEndpoints(args[0])
		}
	}
}

func cmdApiLoginFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		Cmd.Error = Cmd.SunGrow.Login(login.SunGrowAuth{
			AppKey:       Cmd.ApiAppKey,
			UserAccount:  Cmd.ApiUsername,
			UserPassword: Cmd.ApiPassword,
			TokenFile:    Cmd.ApiTokenFile,
			Force:        true,
		})
		if Cmd.Error != nil {
			break
		}

		Cmd.SunGrow.Auth.Print()

		if Cmd.SunGrow.HasTokenChanged() {
			Cmd.ApiLastLogin = Cmd.SunGrow.GetLastLogin()
			Cmd.ApiToken = Cmd.SunGrow.GetToken()
			Cmd.Error = writeConfig()
		}
	}
}

func cmdApiGetFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.SunGrow.OutputType.SetJson()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = Cmd.SunGrow.AllCritical()
			break
		}

		ep := Cmd.SunGrow.GetByJson(args[0], args[1])
		if Cmd.SunGrow.Error != nil {
			Cmd.Error = Cmd.SunGrow.Error
			break
		}
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = ep.GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdApiRawFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.SunGrow.OutputType.SetRaw()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = Cmd.SunGrow.AllCritical()
			break
		}

		ep := Cmd.SunGrow.GetByJson(args[0], args[1])
		if Cmd.SunGrow.Error != nil {
			Cmd.Error = Cmd.SunGrow.Error
			break
		}
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = ep.GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdApiSaveFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.SunGrow.OutputType.SetFile()

		args = fillArray(2, args)
		if args[0] == "all" {
			Cmd.Error = Cmd.SunGrow.AllCritical()
			break
		}

		ep := Cmd.SunGrow.GetByJson(args[0], args[1])
		if Cmd.SunGrow.Error != nil {
			Cmd.Error = Cmd.SunGrow.Error
			break
		}
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = ep.GetError()
		if Cmd.Error != nil {
			break
		}
	}
}

func cmdApiPutFunc(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// args = fillArray(1, args)
		// Cmd.Error = SunGrow.Init()
		// if Cmd.Error != nil {
		// 	break
		// }
	}
}
