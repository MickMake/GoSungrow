package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/cmdConfig"
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
			Short:                 fmt.Sprintf("High-level iSolarCloud functions."),
			Long:                  fmt.Sprintf("High-level iSolarCloud functions."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args:                  cobra.MinimumNArgs(0),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "get <endpoint>", "put <endpoint>")

		// ********************************************************************************
		var cmdDataGet = &cobra.Command{
			Use:                   "get < [area.]<endpoint> | <command> >",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (table)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (table)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataGet)
		cmdDataGet.Example = cmdHelp.PrintExamples(cmdDataGet, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ********************************************************************************
		var cmdDataRaw = &cobra.Command{
			Use:                   output.StringTypeRaw + " < [area.]<endpoint> | <command> >",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (raw)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (raw)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataRaw)
		cmdDataRaw.Example = cmdHelp.PrintExamples(cmdDataRaw, "[area.]<endpoint>")

		// ********************************************************************************
		var cmdDataJson = &cobra.Command{
			Use:                   output.StringTypeJson + " < [area.]<endpoint> | <command> >",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataJson)
		cmdDataJson.Example = cmdHelp.PrintExamples(cmdDataJson, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ********************************************************************************
		var cmdDataCsv = &cobra.Command{
			Use:                   output.StringTypeCsv + " < [area.]<endpoint> | <command> >",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataCsv)
		cmdDataCsv.Example = cmdHelp.PrintExamples(cmdDataCsv, "queryDeviceList", "WebAppService.showPSView", "stats")

		// ********************************************************************************
		var cmdDataGraph = &cobra.Command{
			Use:                   output.StringTypeGraph + " < [area.]<endpoint> | <command> >",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (graph)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (graph)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE:                  c.GetEndpoints,
			Args:                  cobra.MinimumNArgs(0),
		}
		c.SelfCmd.AddCommand(cmdDataGraph)
		cmdDataGraph.Example = cmdHelp.PrintExamples(cmdDataGraph, "queryDeviceList", "WebAppService.showPSView", "stats")
	}
	return c.SelfCmd
}


func (c *CmdData) GetEndpoints(cmd *cobra.Command, args []string) error {
	// endpoints string, psIds string, date string
	for range Only.Once {
		cmds.Api.SunGrow.SetOutputType(cmd.Name())
		args = cmdConfig.FillArray(2, args)
		c.Error = cmds.Api.SunGrow.GetEndpoints(args[0], args[1:]...)
	}

	return c.Error
}
