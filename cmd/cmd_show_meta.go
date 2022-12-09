package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


func (c *CmdShow) AttachMeta(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "meta",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Meta"},
			Short:                 fmt.Sprintf("Meta related Sungrow commands."),
			Long:                  fmt.Sprintf("Meta related Sungrow commands."),
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

		c.AttachMetaList(self)
		// c.AttachMetaTree(self)
		// c.AttachMetaPoints(self)
		// c.AttachMetaData(self)
		// c.AttachMetaGraph(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachMetaList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "unit-list",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Meta"},
		Short:                 fmt.Sprintf("Show all unit lists."),
		Long:                  fmt.Sprintf("Show all unit lists."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcMetaList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcMetaList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		c.Error = cmds.Api.SunGrow.MetaUnitList()
		if c.Error != nil {
			break
		}
	}
	return c.Error
}
