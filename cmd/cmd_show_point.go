package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


func (c *CmdShow) AttachPoint(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "point",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Point"},
			Short:                 fmt.Sprintf("Point related Sungrow commands."),
			Long:                  fmt.Sprintf("Point related Sungrow commands."),
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

		c.AttachPointPs(self)
		c.AttachPointPsTable(self)
		c.AttachPointPsGraph(self)

		c.AttachPointDevice(self)
		c.AttachPointDeviceTable(self)
		c.AttachPointDeviceGraph(self)

		c.AttachPointTemplate(self)
		c.AttachPointTemplateTable(self)
		c.AttachPointTemplateGraph(self)

		c.AttachPointData(self)
		c.AttachPointGraph(self)

		c.AttachPointScan(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachPointPs(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "ps <ps_ids | .> [device_type]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("List data points used by a given ps_id."),
		Long:                  fmt.Sprintf("List data points used by a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsPoints,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "", "1171348", "1171348 14")

	return cmd
}

func (c *CmdShow) AttachPointPsTable(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "ps-data <ps_ids | .> <device_type | .> (start)<YYYYmmdd[HHMMSS]> (end)<YYYYmmdd[HHMMSS]> <interval_minutes | .>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate points table for a given ps_id."),
		Long:                  fmt.Sprintf("Generate points table for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsData,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"1171348 43 20221201 20221202 10",
		"1171348 43 20221201 20221202",
		"1171348 43 20221201",
		". . 20221201 20221202 60",
		)

	return cmd
}

func (c *CmdShow) AttachPointPsGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "ps-graph <ps_ids | .> <device_type | .> (start)<YYYYmmdd[HHMMSS]> (end)<YYYYmmdd[HHMMSS]> <interval_minutes | .>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given ps_id."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsGraph,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"1171348 43 20221201 20221202 10",
		"1171348 43 20221201 20221202",
		"1171348 43 20221201",
		". . 20221201 20221202 60",
	)

	return cmd
}


func (c *CmdShow) AttachPointDevice(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "device <device_type ...>",
		Aliases:               []string{"devices"},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("List data points used by a device."),
		Long:                  fmt.Sprintf("List data points used by a device."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDevicePoints,
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"1",
		"11",
	)

	return cmd
}

func (c *CmdShow) AttachPointDeviceTable(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "device-data <device_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate points table for a given device."),
		Long:                  fmt.Sprintf("Generate points table for a given device."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDeviceData,
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

func (c *CmdShow) AttachPointDeviceGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "device-graph <template_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given device."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given device."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcDeviceData,
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


func (c *CmdShow) AttachPointTemplate(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "template <template_id>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("List data points used by a report template."),
		Long:                  fmt.Sprintf("List data points used by a report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcTemplatePoints,
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"8042",
		"8040",
	)

	return cmd
}

func (c *CmdShow) AttachPointTemplateTable(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "template-data <template_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate points table for a given report template."),
		Long:                  fmt.Sprintf("Generate points table for a given report template."),
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

func (c *CmdShow) AttachPointTemplateGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "template-graph <template_id> [start date] [end date] [interval]",
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given report template."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given report template."),
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


func (c *CmdShow) AttachPointData(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "data (start)<YYYYmmdd[HHMMSS]> (end)<YYYYmmdd[HHMMSS]> <interval_minutes | .> <points ...>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Get data points."),
		Long:                  fmt.Sprintf("Get data points."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPointData,
		Args:                  cobra.MinimumNArgs(4),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"20221101 20221104 5 1171348_11_0_0.p83033,1171348_11_0_0.p83072,1171348_11_0_0.p83128",
		"20221101 20221104 . 1171348_11_0_0.p83033 1171348_11_0_0.p83072 1171348_11_0_0.p83128",
	)

	return cmd
}
func (c *CmdShow) funcPointData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.PointData(args[0], args[1], args[2], args[3:]...)
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPointGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "graph (start)<YYYYmmdd[HHMMSS]> (end)<YYYYmmdd[HHMMSS]> <interval_minutes | .> <points ...>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Point"},
		Short:                 fmt.Sprintf("Graph data points."),
		Long:                  fmt.Sprintf("Graph data points."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPointGraph,
		Args:                  cobra.MinimumNArgs(4),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"20221101 20221104 5 1171348_11_0_0.p83033,1171348_11_0_0.p83072,1171348_11_0_0.p83128",
		"20221101 20221104 . 1171348_11_0_0.p83033 1171348_11_0_0.p83072 1171348_11_0_0.p83128",
		)

	return cmd
}
func (c *CmdShow) funcPointGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetGraph()
		args = MinimumArraySize(4, args)
		c.Error = cmds.Api.SunGrow.PointData(args[0], args[1], args[2], args[3:]...)
		if c.Error != nil {
			break
		}
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
		RunE:                  c.funcPointScan,
		Args:                  cobra.MinimumNArgs(2),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self,
		"1 1000",
	)

	return cmd
}
func (c *CmdShow) funcPointScan(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(2, args)
		var points string
		points, c.Error = cmds.Api.SunGrow.PointScan(args[0], args[1])
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", points)
	}
	return c.Error
}
