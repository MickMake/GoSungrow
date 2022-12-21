package cmd

import (
	"GoSungrow/iSolarCloud/api/GoStruct/output"
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

		c.AttachMetaUnitList(self)
		c.AttachMetaMqtt(self)
		c.AttachMetaRealTime(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachMetaUnitList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "unit-list",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Meta"},
		Short:                 fmt.Sprintf("Show all unit lists."),
		Long:                  fmt.Sprintf("Show all unit lists."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcMetaUnitList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcMetaUnitList(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		c.Error = cmds.Api.SunGrow.MetaUnitList()
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachMetaMqtt(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "mqtt",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Meta"},
		Short:                 fmt.Sprintf("Show iSolarCloud mqtt info."),
		Long:                  fmt.Sprintf("Show iSolarCloud mqtt info."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcMetaMqtt,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcMetaMqtt(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.SetOutputType(output.StringTypeTable)
		args = MinimumArraySize(1, args)
		c.Error = cmds.Api.SunGrow.GetIsolarcloudMqtt(args[0])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachMetaRealTime(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "real-time",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Meta"},
		Short:                 fmt.Sprintf("Show iSolarCloud real-time info."),
		Long:                  fmt.Sprintf("Show iSolarCloud real-time info."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcMetaRealTime,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcMetaRealTime(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.SetOutputType(output.StringTypeTable)
		args = MinimumArraySize(1, args)
		c.Error = cmds.Api.SunGrow.GetRealTimeData(args[0])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}
