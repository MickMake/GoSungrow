package cmd

import (
	"fmt"
	"sync"

	"github.com/anicoll/gosungrow/iSolarCloud"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/spf13/cobra"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdData CmdDefault

func NewCmdData() *CmdData {
	var ret *CmdData

	var once sync.Once
	once.Do(func() {
		ret = &CmdData{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	})

	return ret
}

func (c *CmdData) AttachCommand(cmd *cobra.Command) *cobra.Command {
	var once sync.Once
	once.Do(func() {
		if cmd == nil {
			return
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
		c.SelfCmd.Example = cmdhelp.PrintExamples(c.SelfCmd, "get <endpoint>", "put <endpoint>")

		// ******************************************************************************** //
		cmdDataGet := &cobra.Command{
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
		cmdDataGet.Example = cmdhelp.PrintExamples(cmdDataGet, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdDataTable := &cobra.Command{
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
		cmdDataGet.Example = cmdhelp.PrintExamples(cmdDataTable, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdDataRaw := &cobra.Command{
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
		cmdDataRaw.Example = cmdhelp.PrintExamples(cmdDataRaw, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdDataJson := &cobra.Command{
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
		cmdDataJson.Example = cmdhelp.PrintExamples(cmdDataJson, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdDataCsv := &cobra.Command{
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
		cmdDataCsv.Example = cmdhelp.PrintExamples(cmdDataCsv, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdDataGraph := &cobra.Command{
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
		cmdDataGraph.Example = cmdhelp.PrintExamples(cmdDataGraph, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ******************************************************************************** //
		cmdApiXML := &cobra.Command{
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
		cmdApiXML.Example = cmdhelp.PrintExamples(cmdApiXML, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		cmdApiXLSX := &cobra.Command{
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
		cmdApiXLSX.Example = cmdhelp.PrintExamples(cmdApiXLSX, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		cmdApiMarkDown := &cobra.Command{
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
		cmdApiMarkDown.Example = cmdhelp.PrintExamples(cmdApiMarkDown, "queryDeviceList", "WebAppService.showPSView")

		// ******************************************************************************** //
		cmdApiStruct := &cobra.Command{
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
		cmdApiStruct.Example = cmdhelp.PrintExamples(cmdApiStruct, "queryDeviceList", "WebAppService.showPSView")
	})
	return c.SelfCmd
}

func (c *CmdData) GetEndpoints(cmd *cobra.Command, args []string) error {
	// endpoints string, psIds string, date string
	var once sync.Once
	once.Do(func() {
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
			return
		}

		c.Error = data.Output()
		if c.Error != nil {
			return
		}
	})

	return c.Error
}
