package cmd

import (
	"GoSungrow/iSolarCloud"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
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
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all available PS."),
		Long:                  fmt.Sprintf("Show all available PS."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPsIdList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPsIdList(_ *cobra.Command, args []string) error {
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
		Use:                   "points",
		Aliases:               []string{},
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
		points := cmds.Api.SunGrow.PointNames(args...)
		if cmds.Api.SunGrow.Error != nil {
			c.Error = cmds.Api.SunGrow.Error
			break
		}
		fmt.Printf("%s\n", points)
	}
	return c.Error
}
