package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)

//
// ********************************************************************************
//
var cmdData = &cobra.Command{
	Use:                   "data",
	Aliases:               []string{},
	Short:                 fmt.Sprintf("High-level iSolarCloud functions."),
	Long:                  fmt.Sprintf("High-level iSolarCloud functions."),
	Example:               PrintExamples("data", "get <endpoint>", "put <endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdDataFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdDataFunc(cmd *cobra.Command, args []string) {
	Cmd.Error = cmd.Help()
}

//
// ********************************************************************************
//
var cmdDataList = &cobra.Command{
	Use:                   "ls",
	Aliases:               []string{"list"},
	Short:                 fmt.Sprintf("List iSolarCloud high-level data commands."),
	Long:                  fmt.Sprintf("List iSolarCloud high-level data commands."),
	Example:               PrintExamples("data ls", "", "areas", "endpoints", "<area name>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	// PreRunE:               Cmd.SunGrowArgs,
	Run:  cmdDataListFunc,
	Args: cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdDataListFunc(cmd *cobra.Command, args []string) {
	SunGrow.ListHighLevel()
}

//
// ********************************************************************************
//
var cmdDataLogin = &cobra.Command{
	Use: "login",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Login to iSolarCloud"),
	Long:                  fmt.Sprintf("Login to iSolarCloud"),
	Example:               PrintExamples("data login", ""),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdApiLoginFunc,
	Args:                  cobra.ExactArgs(0),
}

//
// ********************************************************************************
//
var cmdDataGet = &cobra.Command{
	Use: "get",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("GetByJson high-level data from iSolarCloud"),
	Long:                  fmt.Sprintf("GetByJson high-level data from iSolarCloud"),
	Example:               PrintExamples("data get", "<endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdDataGetFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdDataGetFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetHuman()

		args = fillArray(3, args)
		Cmd.Error = SunGrow.GetHighLevel(args[0], args[1:]...)
	}
}

//
// ********************************************************************************
//
var cmdDataRaw = &cobra.Command{
	Use: "raw",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Raw high-level data from iSolarCloud"),
	Long:                  fmt.Sprintf("Raw high-level data from iSolarCloud"),
	Example:               PrintExamples("data raw", "<endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdDataRawFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdDataRawFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetRaw()

		args = fillArray(3, args)
		Cmd.Error = SunGrow.GetHighLevel(args[0], args[1:]...)
	}
}

//
// ********************************************************************************
//
var cmdDataSave = &cobra.Command{
	Use: "save",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Save high-level data from iSolarCloud"),
	Long:                  fmt.Sprintf("Save high-level data from iSolarCloud"),
	Example:               PrintExamples("data save", "<endpoint>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdDataSaveFunc,
	Args:                  cobra.MinimumNArgs(1),
}

//goland:noinspection GoUnusedParameter
func cmdDataSaveFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		SunGrow.OutputType.SetFile()

		args = fillArray(3, args)
		Cmd.Error = SunGrow.GetHighLevel(args[0], args[1:]...)
	}
}

//
// ********************************************************************************
//
var cmdDataPut = &cobra.Command{
	Use:                   "put",
	Aliases:               []string{"write"},
	Short:                 fmt.Sprintf("Set high-level data on iSolarCloud"),
	Long:                  fmt.Sprintf("Set high-level data on iSolarCloud"),
	Example:               PrintExamples("data put", "<endpoint> <value>"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdDataPutFunc,
	Args:                  cobra.ExactArgs(2),
}

//goland:noinspection GoUnusedParameter
func cmdDataPutFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// SunGrow.OutputType.SetFile()
		// args = fillArray(2, args)
		// Cmd.Error = SunGrow.PutHighLevel(args[0], args[1])
	}
}
