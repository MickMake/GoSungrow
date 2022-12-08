package cmd

import (
	"GoSungrow/iSolarCloud"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"strings"
)


func (c *CmdShow) AttachPsId(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "ps",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "PsId"},
			Short:                 fmt.Sprintf("General iSolarCloud functions."),
			Long:                  fmt.Sprintf("General iSolarCloud functions."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               cmds.SunGrowArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(self)
		self.Example = cmdHelp.PrintExamples(self, "")

		c.AttachPsIdList(self)
		c.AttachPsTree(self)
		c.AttachPsPoints(self)
	}
	return c.SelfCmd
}

func (c *CmdShow) AttachPsIdList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("Show all devices on account."),
		Long:                  fmt.Sprintf("Show all devices on account."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsIdListList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsIdListList(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var devices string
		devices, c.Error = cmds.Api.SunGrow.Devices(args...)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", devices)
	}
	return c.Error
}

func (c *CmdShow) AttachPsIdList2(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all available PS."),
		Long:                  fmt.Sprintf("Show all available PS."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsIdList2,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsIdList2(_ *cobra.Command, args []string) error {
	for range Only.Once {
		pids := cmds.Api.SunGrow.SetPsIds(args...)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", pids)
	}
	return c.Error
}

func (c *CmdShow) AttachPsTree(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "tree",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show the PS tree."),
		Long:                  fmt.Sprintf("Show the PS tree."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsTree,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsTree(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var pids iSolarCloud.PsTree
		pids, c.Error = cmds.Api.SunGrow.PsTreeMenu(args...)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", pids)
	}
	return c.Error
}

func (c *CmdShow) AttachPsPoints(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "points <ps_ids | .> [device_type]",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show available PS points."),
		Long:                  fmt.Sprintf("Show available PS points."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsPoints,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsPoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(2, args)
		points := cmds.Api.SunGrow.PointNames(strings.Split(args[0], ","), args[1])
		if cmds.Api.SunGrow.Error != nil {
			c.Error = cmds.Api.SunGrow.Error
			break
		}
		fmt.Printf("%s\n", points)
	}
	return c.Error
}
