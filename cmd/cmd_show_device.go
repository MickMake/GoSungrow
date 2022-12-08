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
			Short:                 fmt.Sprintf("Sungrow devices."),
			Long:                  fmt.Sprintf("Sungrow devices."),
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
		c.AttachDevicePoints(self)
		c.AttachDeviceModels(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachDeviceList(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("List all device types."),
		Long:                  fmt.Sprintf("List all device types."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDeviceTypeList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}
func (c *CmdShow) funcDeviceTypeList(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		ret := cmds.Api.SunGrow.DeviceTypeList()
		fmt.Println(ret)
	}
	return c.Error
}

func (c *CmdShow) AttachDevicePoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "points <device_type ...>",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("List data points used by a device."),
		Long:                  fmt.Sprintf("List data points used by a device."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDevicePoints,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2,
		"1",
		"11",
	)

	return cmd
}
func (c *CmdShow) funcDevicePoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		ret := cmds.Api.SunGrow.DevicePoints(args...)
		fmt.Println(ret)
	}
	return c.Error
}

func (c *CmdShow) AttachDeviceData(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "data <device_type> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDeviceData,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"11 20221201 20221202 30",
		"11 20221201 20221202 5",
		"11 20221201 20221202",
		"11 20221201",
		"11",
	)

	return cmd
}
func (c *CmdShow) funcDeviceData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.DeviceData(args[0], args[1], args[2], args[3])
	}
	return c.Error
}

func (c *CmdShow) AttachDeviceDataGraph(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var self = &cobra.Command{
		Use:                   "data-graph <device_type> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Device"},
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDeviceDataGraph,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"11 20221201 20221202 30",
		"11 20221201 20221202 5",
		"11 20221201 20221202",
		"11 20221201",
		"11",
	)

	return cmd
}
func (c *CmdShow) funcDeviceDataGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetGraph()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.DeviceData(args[0], args[1], args[2], args[3])
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
		RunE:                  c.funcDeviceModels,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}
func (c *CmdShow) funcDeviceModels(_ *cobra.Command, _ []string) error {
	for range Only.Once {
		// _ = cmds.SetOutputType(cmd)
		c.Error = cmds.Api.SunGrow.DeviceModelInfoList()
	}
	return c.Error
}
