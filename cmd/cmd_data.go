package cmd

import (
	"fmt"

	"github.com/MickMake/GoSungrow/iSolarCloud"
	"github.com/MickMake/GoSungrow/iSolarCloud/api/GoStruct/output"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdData CmdDefault

func NewCmdData() *CmdData {
	var ret *CmdData

	for range Only.Once {
		ret = &CmdData{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *CmdData) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "data",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Mid-level access to the Sungrow api."),
			Long:                  fmt.Sprintf("Mid-level access to the Sungrow api."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "get <endpoint>", "put <endpoint>")

		// ******************************************************************************** //
		var cmdDataGet = &cobra.Command{
			Use:                   output.StringTypeList + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{"get"},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (list)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (list)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataGet)
		cmdDataGet.Example = cmdHelp.PrintExamples(cmdDataGet, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdDataTable = &cobra.Command{
			Use:                   output.StringTypeTable + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (table)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (table)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataTable)
		cmdDataGet.Example = cmdHelp.PrintExamples(cmdDataTable, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdDataRaw = &cobra.Command{
			Use:                   output.StringTypeRaw + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (raw)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (raw)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataRaw)
		cmdDataRaw.Example = cmdHelp.PrintExamples(cmdDataRaw, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdDataJson = &cobra.Command{
			Use:                   output.StringTypeJson + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (json)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataJson)
		cmdDataJson.Example = cmdHelp.PrintExamples(cmdDataJson, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdDataCsv = &cobra.Command{
			Use:                   output.StringTypeCsv + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (csv)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (csv)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataCsv)
		cmdDataCsv.Example = cmdHelp.PrintExamples(cmdDataCsv, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdDataGraph = &cobra.Command{
			Use:                   output.StringTypeGraph + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (graph)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (graph)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataGraph)
		cmdDataGraph.Example = cmdHelp.PrintExamples(cmdDataGraph, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		var cmdApiXML = &cobra.Command{
			Use:                   output.StringTypeXML + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (xml)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (xml)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdApiXML)
		cmdApiXML.Example = cmdHelp.PrintExamples(cmdApiXML, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		var cmdApiXLSX = &cobra.Command{
			Use:                   output.StringTypeXLSX + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (XLSX)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (XLSX)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdApiXLSX)
		cmdApiXLSX.Example = cmdHelp.PrintExamples(cmdApiXLSX, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		var cmdApiMarkDown = &cobra.Command{
			Use:                   output.StringTypeMarkDown + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get data from the Sungrow api (MarkDown)"),
			Long:                  fmt.Sprintf("Get data from the Sungrow api (MarkDown)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdApiMarkDown)
		cmdApiMarkDown.Example = cmdHelp.PrintExamples(cmdApiMarkDown, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		var cmdApiStruct = &cobra.Command{
			Use:                   output.StringTypeStruct + " <[area.]endpoint> [endpoint args ...]",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Show response as Go structure (debug)"),
			Long:                  fmt.Sprintf("Show response as Go structure (debug)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdApiStruct)
		cmdApiStruct.Example = cmdHelp.PrintExamples(cmdApiStruct, "queryDeviceList", "WebAppService.showPSView")
	}
	return c.SelfCmd
}

func (c *CmdData) GetEndpoints(cmd *cobra.Command, args []string) error {
	// endpoints string, psIds string, date string
	for range Only.Once {
		cmds.Api.SunGrow.SetOutputType(cmd.Name())
		if cmd.Name() == output.StringTypeXLSX {
			cmds.Api.SaveFile = true
		}
		if cmd.Name() == output.StringTypeGraph {
			cmds.Api.SaveFile = true
		}

		args = MinimumArraySize(2, args)
		eps := iSolarCloud.SplitArg(args[0])

		data := cmds.Api.SunGrow.NewSunGrowData()
		data.SetOutput(cmd.Name())
		data.SetSaveAsFile(cmds.Api.SaveFile)
		data.SetEndpoints(eps...)
		data.SetArgs(args[1:]...)

		c.Error = data.GetData()
		if c.Error != nil {
			break
		}

		c.Error = data.Output()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}
