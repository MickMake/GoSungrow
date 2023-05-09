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
			Short:                 fmt.Sprintf("Template related Sungrow commands."),
			Long:                  fmt.Sprintf("Template related Sungrow commands."),
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
		c.AttachTemplateGraph(self)
		c.AttachTemplateSave(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachTemplateList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
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
	var self = &cobra.Command{
		Use:                   "points <template id>",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("List points used for a given template_id."),
		Long:                  fmt.Sprintf("List points used for a given template_id."),
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
	var self = &cobra.Command{
		Use:                   "data <template_id> " + ArgsDateInterval,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Generate points table for a given template_id."),
		Long:                  fmt.Sprintf("Generate points table for a given template_id."),
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
		args = MinimumArraySize(4, args)
		cmds.Api.SunGrow.OutputType.SetTable()
		c.Error = cmds.Api.SunGrow.TemplateData(args[0], args[1], args[2], args[3])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachTemplateGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "graph <template_id> " + ArgsDateInterval,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given template_id."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given template_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplateGraph,
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
func (c *CmdShow) funcTemplateGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(4, args)
		cmds.Api.SunGrow.OutputType.SetGraph()
		c.Error = cmds.Api.SunGrow.TemplateData(args[0], args[1], args[2], args[3])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachTemplateSave(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "save <template_id> " + ArgsDateInterval,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Template"},
		Short:                 fmt.Sprintf("Generate and save data points for a given template_id."),
		Long:                  fmt.Sprintf("Generate and save data points for a given template_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplateSave,
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
func (c *CmdShow) funcTemplateSave(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.TemplateDataSave(args[0], args[1], args[2], args[3])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}
