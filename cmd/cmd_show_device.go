package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


func (c *CmdShow) AttachDevice(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "device",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Device"},
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

		c.AttachDeviceList(self)
		c.AttachDeviceModels(self)
	}
	return c.SelfCmd
}

func (c *CmdShow) AttachDeviceList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("Show all devices on account."),
		Long:                  fmt.Sprintf("Show all devices on account."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachDeviceList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcAttachDeviceList(_ *cobra.Command, args []string) error {
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


func (c *CmdShow) AttachDeviceModels(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "models",
		Aliases:               []string{"model"},
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("Get ALL Sungrow models (large list)."),
		Long:                  fmt.Sprintf("Get ALL Sungrow models (large list)."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcAttachDeviceModels,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}
func (c *CmdShow) funcAttachDeviceModels(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
		c.Error = cmds.Api.SunGrow.GetDeviceModelInfoList()
	}
	return c.Error
}
