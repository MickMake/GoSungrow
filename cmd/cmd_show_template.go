package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdConfig"
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
		RunE:                  c.funcAttachTemplateList,
		Args:                  cobra.ExactArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachTemplateList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
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
		RunE:                  c.funcAttachTemplatePoints,
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachTemplatePoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
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
		RunE:                  c.funcAttachTemplateData,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "8042 20220212", "8042 20220212 '{\"search_string\":\"p83106\",\"min_left_axis\":-6000,\"max_left_axis\":12000}'")

	return cmd
}
func (c *CmdShow) funcAttachTemplateData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
		args = cmdConfig.FillArray(4, args)
		c.Error = cmds.Api.SunGrow.TemplateData(args[0], args[1], args[2], args[3])
	}
	return c.Error
}

