package cmd

import (
	"GoSungrow/iSolarCloud/api/GoStruct/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdInfo CmdDefault

func NewCmdInfo() *CmdInfo {
	var ret *CmdInfo

	for range Only.Once {
		ret = &CmdInfo{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *CmdInfo) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "info",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("General iSolarCloud functions."),
			Long:                  fmt.Sprintf("General iSolarCloud functions."),
			Deprecated:            "show",
			Hidden:                true,
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "get <endpoint>", "put <endpoint>")

		// ********************************************************************************
		var cmdInfoGet = &cobra.Command{
			Use:                   "get",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (table)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (table)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdInfoGet)
		cmdInfoGet.Example = cmdHelp.PrintExamples(cmdInfoGet, "[area.]<endpoint>")
		// c.AttachCmdInfo(cmdInfoGet)

		// ********************************************************************************
		var cmdInfoRaw = &cobra.Command{
			Use:                   output.StringTypeRaw,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (raw)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (raw)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdInfoRaw)
		cmdInfoRaw.Example = cmdHelp.PrintExamples(cmdInfoRaw, "[area.]<endpoint>")
		// c.AttachCmdInfo(cmdInfoRaw)

		// ********************************************************************************
		var cmdInfoJson = &cobra.Command{
			Use:                   output.StringTypeJson,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdInfoJson)
		cmdInfoJson.Example = cmdHelp.PrintExamples(cmdInfoJson, "[area.]<endpoint>")
		// c.AttachCmdInfo(cmdInfoJson)

		// ********************************************************************************
		var cmdInfoCsv = &cobra.Command{
			Use:                   output.StringTypeCsv,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdInfoCsv)
		cmdInfoCsv.Example = cmdHelp.PrintExamples(cmdInfoCsv, "[area.]<endpoint>")
		// c.AttachCmdInfo(cmdInfoCsv)

		// ********************************************************************************
		var cmdInfoMarkDown = &cobra.Command{
			Use:                   output.StringTypeMarkDown,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (MarkDown)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (MarkDown)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdInfoMarkDown)
		cmdInfoMarkDown.Example = cmdHelp.PrintExamples(cmdInfoMarkDown, "[area.]<endpoint>")
		// c.AttachCmdInfo(cmdInfoMarkDown)

		// // ********************************************************************************
		// var cmdInfoPut = &cobra.Command{
		// 	Use:                   "put",
		// 	Aliases:               []string{"set", "write"},
		// 	Annotations:           map[string]string{"group": "Info"},
		// 	Short:                 fmt.Sprintf("Set info on iSolarCloud"),
		// 	Long:                  fmt.Sprintf("Set info on iSolarCloud"),
		// 	DisableFlagParsing:    false,
		// 	DisableFlagsInUseLine: false,
		// 	PreRunE:               cmds.SunGrowArgs,
		// 	Run:                   cmds.CmdInfoPut,
		// 	Args:                  cobra.ExactArgs(2),
		// }
		// c.SelfCmd.AddCommand(cmdInfoPut)
		// cmdInfoPut.Example = cmdHelp.PrintExamples(cmdInfoPut, "[area.]<endpoint> <value>")
	}
	return c.SelfCmd
}

// func (c *CmdInfo) AttachCmdInfo(cmd *cobra.Command) *cobra.Command {
// 	for range Only.Once {
// 		if cmd == nil {
// 			break
// 		}
//
// 		// c.AttachCmdInfoStats(cmd)
//
// 		// c.AttachCmdInfoRealTime(cmd)
// 	}
// 	return cmd
// }

// func (c *CmdInfo) AttachCmdInfoStats(cmd *cobra.Command) *cobra.Command {
// 	// ********************************************************************************
// 	var c2 = &cobra.Command{
// 		Use:                   "stats",
// 		Aliases:               []string{},
// 		Annotations:           map[string]string{"group": "Data"},
// 		Short:                 fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
// 		Long:                  fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
// 		DisableFlagParsing:    false,
// 		DisableFlagsInUseLine: false,
// 		PreRunE:               cmds.SunGrowArgs,
// 		RunE:                  func(cmd *cobra.Command, args []string) error {
// 			_ = cmds.SetOutputType(cmd)
// 			return cmds.Api.SunGrow.PrintCurrentStats()
// 		},
// 		Args:                  cobra.ExactArgs(0),
// 	}
// 	cmd.AddCommand(c2)
// 	c2.Example = cmdHelp.PrintExamples(c2, "")
//
// 	return cmd
// }

// func (c *CmdInfo) AttachCmdInfoRealTime(cmd *cobra.Command) *cobra.Command {
// 	// ********************************************************************************
// 	var c2 = &cobra.Command{
// 		Use:                   "real-time",
// 		Aliases:               []string{"realtime"},
// 		Annotations:           map[string]string{"group": "Data"},
// 		Short:                 fmt.Sprintf("Get iSolarCloud real-time data."),
// 		Long:                  fmt.Sprintf("Get iSolarCloud real-time data."),
// 		DisableFlagParsing:    false,
// 		DisableFlagsInUseLine: false,
// 		PreRunE:               cmds.SunGrowArgs,
// 		RunE:                  func(cmd *cobra.Command, args []string) error {
// 			_ = cmds.SetOutputType(cmd)
// 			args = MinimumArraySize(1, args)
// 			return cmds.Api.SunGrow.GetRealTimeData(args[0])
// 		},
// 		Args:                  cobra.RangeArgs(0, 1),
// 	}
// 	cmd.AddCommand(c2)
// 	c2.Example = cmdHelp.PrintExamples(c2, "")
//
// 	return cmd
// }
