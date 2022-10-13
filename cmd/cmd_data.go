package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/cmdConfig"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"strings"
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
			Use:                   "get",
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
		cmdDataGet.Example = cmdHelp.PrintExamples(cmdDataGet, "[area.]<endpoint>")

		// ********************************************************************************
		var cmdDataRaw = &cobra.Command{
			Use:                   output.StringTypeRaw,
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
			Use:                   output.StringTypeJson,
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
		cmdDataJson.Example = cmdHelp.PrintExamples(cmdDataJson, "[area.]<endpoint>")

		// ********************************************************************************
		var cmdDataCsv = &cobra.Command{
			Use:                   output.StringTypeCsv,
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
		cmdDataCsv.Example = cmdHelp.PrintExamples(cmdDataCsv, "[area.]<endpoint>")

		// ********************************************************************************
		var cmdDataGraph = &cobra.Command{
			Use:                   output.StringTypeGraph,
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
		cmdDataGraph.Example = cmdHelp.PrintExamples(cmdDataGraph, "[area.]<endpoint> ''")
	}
	return c.SelfCmd
}


func SplitArg(arg string) []string {
	var ret []string
	for range Only.Once {
		ret = []string{arg}
		for _, s := range []string{ ",", "/", " "} {
			if strings.Contains(arg, s) {
				ret = strings.Split(arg, s)
				break
			}
		}
	}
	return ret
}

func (c *CmdData) GetEndpoints(cmd *cobra.Command, args []string) error {
	// endpoints string, psIds string, date string
	for range Only.Once {
		cmds.Api.SunGrow.SetOutputType(cmd.Use)
		args = cmdConfig.FillArray(3, args)

		var e []string
		if args[0] != "" {
			e = SplitArg(args[0])
		}

		var p []api.Integer
		for _, psId := range SplitArg(args[1]) {
			if psId == "" {
				continue
			}
			p = append(p, api.SetIntegerString(psId))
		}

		d := api.SetDateTimeString(args[2])

		c.Error = cmds.Api.SunGrow.GetEndpoints(e, p, *d)
	}

	return c.Error
}
