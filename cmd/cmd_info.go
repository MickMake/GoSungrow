package cmd

import (
	"fmt"
	"sync"

	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/anicoll/gosungrow/iSolarCloud/api/GoStruct/output"
	"github.com/spf13/cobra"
)

//goland:noinspection GoNameStartsWithPackageName
type CmdInfo CmdDefault

func NewCmdInfo() *CmdInfo {
	var ret *CmdInfo

	var once sync.Once
	once.Do(func() {
		ret = &CmdInfo{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	})

	return ret
}

func (c *CmdInfo) AttachCommand(cmd *cobra.Command) *cobra.Command {
	var once sync.Once
	once.Do(func() {
		if cmd == nil {
			return
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
		cmdInfoGet := &cobra.Command{
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
		cmdInfoRaw := &cobra.Command{
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
		cmdInfoJson := &cobra.Command{
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
		cmdInfoCsv := &cobra.Command{
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
		cmdInfoMarkDown := &cobra.Command{
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
	})
	return c.SelfCmd
}
