package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)


func AttachCmdData(cmd *cobra.Command) *cobra.Command {
	var cmdData = &cobra.Command{
		Use:                   "data",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("High-level iSolarCloud functions."),
		Long:                  fmt.Sprintf("High-level iSolarCloud functions."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(cmdData)
	cmdData.Example = PrintExamples(cmdData, "get <endpoint>", "put <endpoint>")


	// // ********************************************************************************
	// var cmdDataList = &cobra.Command{
	// 	Use:                   "ls",
	// 	Aliases:               []string{"list"},
	// 	Short:                 fmt.Sprintf("List iSolarCloud high-level data commands."),
	// 	Long:                  fmt.Sprintf("List iSolarCloud high-level data commands."),
	// 	DisableFlagParsing:    false,
	// 	DisableFlagsInUseLine: false,
	// 	// PreRunE:               Cmd.SunGrowArgs,
	// 	Run:  cmdDataListFunc,
	// 	Args: cobra.RangeArgs(0, 1),
	// }
	// cmdData.AddCommand(cmdDataList)
	// cmdDataList.Example = PrintExamples(cmdDataList, "", "areas", "endpoints", "<area name>")

	// ********************************************************************************
	var cmdDataLogin = &cobra.Command{
		Use:                   "login",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Login to iSolarCloud"),
		Long:                  fmt.Sprintf("Login to iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.ProcessArgs,
		Run:                   cmdApiLoginFunc,
		Args:                  cobra.ExactArgs(0),
	}
	cmdData.AddCommand(cmdDataLogin)
	cmdDataLogin.Example = PrintExamples(cmdDataLogin, "")


	// ********************************************************************************
	var cmdDataGet = &cobra.Command{
		Use:                   "get",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Get high-level data from iSolarCloud"),
		Long:                  fmt.Sprintf("Get high-level data from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdData.AddCommand(cmdDataGet)
	cmdDataGet.Example = PrintExamples(cmdDataGet, "[area.]<endpoint>")
	AttachCmdDataStats(cmdDataGet)
	AttachCmdDataTemplate(cmdDataGet)
	AttachCmdDataTemplatePoints(cmdDataGet)
	AttachCmdDataPoints(cmdDataGet)
	AttachCmdDataPointNames(cmdDataGet)

	// ********************************************************************************
	var cmdDataRaw = &cobra.Command{
		Use:                   "raw",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Raw high-level data from iSolarCloud"),
		Long:                  fmt.Sprintf("Raw high-level data from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdData.AddCommand(cmdDataRaw)
	cmdDataRaw.Example = PrintExamples(cmdDataRaw, "[area.]<endpoint>")
	AttachCmdDataStats(cmdDataRaw)
	AttachCmdDataTemplate(cmdDataRaw)
	AttachCmdDataPoints(cmdDataRaw)
	AttachCmdDataPointNames(cmdDataRaw)

	// ********************************************************************************
	var cmdDataSave = &cobra.Command{
		Use:                   "save",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Save high-level data from iSolarCloud"),
		Long:                  fmt.Sprintf("Save high-level data from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdData.AddCommand(cmdDataSave)
	cmdDataSave.Example = PrintExamples(cmdDataSave, "[area.]<endpoint>")
	AttachCmdDataStats(cmdDataSave)
	AttachCmdDataTemplate(cmdDataSave)
	AttachCmdDataPoints(cmdDataSave)
	AttachCmdDataPointNames(cmdDataSave)

	// ********************************************************************************
	var cmdDataGraph = &cobra.Command{
		Use:                   "graph",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Graph high-level data from iSolarCloud"),
		Long:                  fmt.Sprintf("Graph high-level data from iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Args:                  cobra.MinimumNArgs(1),
	}
	cmdData.AddCommand(cmdDataGraph)
	cmdDataGraph.Example = PrintExamples(cmdDataGraph, "[area.]<endpoint> ''")
	AttachCmdDataStats(cmdDataGraph)
	AttachCmdDataTemplate(cmdDataGraph)
	AttachCmdDataPoints(cmdDataGraph)

	// ********************************************************************************
	var cmdDataPut = &cobra.Command{
		Use:                   "put",
		Aliases:               []string{"set", "write"},
		Short:                 fmt.Sprintf("Set high-level data on iSolarCloud"),
		Long:                  fmt.Sprintf("Set high-level data on iSolarCloud"),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		Run:                   cmdDataPutFunc,
		Args:                  cobra.ExactArgs(2),
	}
	cmdData.AddCommand(cmdDataPut)
	cmdDataPut.Example = PrintExamples(cmdDataPut, "[area.]<endpoint> <value>")

	return cmdData
}


// func cmdDataFunc(cmd *cobra.Command, _ []string) {
// 	Cmd.Error = cmd.Help()
// }
//
// func cmdDataListFunc(_ *cobra.Command, _ []string) {
// 	Cmd.SunGrow.ListHighLevel()
// }
//
// func cmdDataGetFunc(_ *cobra.Command, args []string) {
// 	for range Only.Once {
// 		Cmd.SunGrow.OutputType.SetHuman()
//
// 		args = fillArray(3, args)
// 		Cmd.Error = Cmd.SunGrow.GetHighLevel(args[0], args[1:]...)
// 	}
// }
//
// func cmdDataRawFunc(_ *cobra.Command, args []string) {
// 	for range Only.Once {
// 		Cmd.SunGrow.OutputType.SetRaw()
//
// 		args = fillArray(3, args)
// 		Cmd.Error = Cmd.SunGrow.GetHighLevel(args[0], args[1:]...)
// 	}
// }
//
// func cmdDataSaveFunc(_ *cobra.Command, args []string) {
// 	for range Only.Once {
// 		Cmd.SunGrow.OutputType.SetFile()
//
// 		args = fillArray(3, args)
// 		Cmd.Error = Cmd.SunGrow.GetHighLevel(args[0], args[1:]...)
// 	}
// }
//
// func cmdDataGraphFunc(_ *cobra.Command, args []string) {
// 	for range Only.Once {
// 		Cmd.SunGrow.OutputType.SetGraph()
//
// 		args = fillArray(4, args)
// 		Cmd.Error = Cmd.SunGrow.GetHighLevel(args[0], args[1:]...)
// 	}
// }

func cmdDataPutFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		Cmd.SunGrow.OutputType.SetFile()
		args = fillArray(2, args)
		// Cmd.Error = SunGrow.PutHighLevel(args[0], args[1])
	}
}


func SwitchOutput(cmd *cobra.Command) error {
	var err error
	for range Only.Once {
		foo := cmd.Parent()
		switch foo.Use {
			case "get":
				Cmd.SunGrow.OutputType.SetHuman()
			case "raw":
				Cmd.SunGrow.OutputType.SetRaw()
			case "save":
				Cmd.SunGrow.OutputType.SetFile()
			case "graph":
				Cmd.SunGrow.OutputType.SetGraph()
			default:
				Cmd.SunGrow.OutputType.SetHuman()
		}
	}

	return err
}
