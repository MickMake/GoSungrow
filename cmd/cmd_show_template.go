package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


func (c *CmdShow) AttachTemplate(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "template",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Template"},
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

		c.AttachTemplateList(self)
		c.AttachTemplatePoints(self)
		c.AttachTemplateData(self)
		c.AttachTemplateDataGraph(self)
	}
	return c.SelfCmd
}

func (c *CmdShow) AttachTemplateList(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Get all defined templates."),
		Long:                  fmt.Sprintf("Get all defined templates."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplateList,
		Args:                  cobra.ExactArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcTemplateList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		c.Error = cmds.Api.SunGrow.TemplateList()
	}
	return c.Error
}

func (c *CmdShow) AttachTemplatePoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "points <template id>",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Get all points defined within a template."),
		Long:                  fmt.Sprintf("Get all points defined within a template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplatePoints,
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "8092")

	return cmd
}
func (c *CmdShow) funcTemplatePoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		c.Error = cmds.Api.SunGrow.TemplatePoints(args[0])
	}
	return c.Error
}

func (c *CmdShow) AttachTemplateData(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "data <template_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplateData,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"8092 20221201 20221202 30",
		"8092 20221201 20221202 5",
		"8092 20221201 20221202",
		"8092 20221201",
		"8092",
		)

	return cmd
}
func (c *CmdShow) funcTemplateData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.TemplateData(args[0], args[1], args[2], args[3])
	}
	return c.Error
}

func (c *CmdShow) AttachTemplateDataGraph(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "data-graph <template_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplateDataGraph,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"8092 20221201 20221202 30",
		"8092 20221201 20221202 5",
		"8092 20221201 20221202",
		"8092 20221201",
		"8092",
	)

	return cmd
}
func (c *CmdShow) funcTemplateDataGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetGraph()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.TemplateData(args[0], args[1], args[2], args[3])
	}
	return c.Error
}
