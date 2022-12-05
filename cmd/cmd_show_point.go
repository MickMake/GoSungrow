package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdConfig"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


func (c *CmdShow) AttachPoint(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "point",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Point"},
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

		c.AttachPointNames(self)
		c.AttachPointScan(self)
		c.AttachPointData(self)
		c.AttachPointTemplate(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachPointNames(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "names [ps_id]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Get point names for a given ps_id."),
		Long:                  fmt.Sprintf("Get point names for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPointNames,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPointNames(_ *cobra.Command, args []string) error {
	for range Only.Once {
		points := cmds.Api.SunGrow.PointNames(args...)
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", points)
	}
	return c.Error
}

func (c *CmdShow) AttachPointScan(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "scan [min] [max]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Scan full list of points."),
		Long:                  fmt.Sprintf("Scan full list of points."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPointScan,
		Args:                  cobra.MinimumNArgs(2),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPointScan(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(2, args)
		points := cmds.Api.SunGrow.PointScan(args[0], args[1])
		if c.Error != nil {
			break
		}
		fmt.Printf("%s\n", points)
	}
	return c.Error
}

func (c *CmdShow) AttachPointData(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "data <start:YYYYmmdd[HHMMSS]> <end:YYYYmmdd[HHMMSS]> <interval minutes> <points ...>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Get data points."),
		Long:                  fmt.Sprintf("Get data points."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPointData,
		Args:                  cobra.MinimumNArgs(4),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachPointData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = cmdConfig.FillArray(4, args)
		cmds.Api.SunGrow.PointData(args[0], args[1], args[2], args[3:]...)
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPointTemplate(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "template <template_id>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("List data points used in report template."),
		Long:                  fmt.Sprintf("List data points used in report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachPointTemplate,
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "8042", "8040")

	return cmd
}
func (c *CmdShow) funcAttachPointTemplate(_ *cobra.Command, args []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
		c.Error = cmds.Api.SunGrow.TemplatePoints(args[0])
	}
	return c.Error
}
