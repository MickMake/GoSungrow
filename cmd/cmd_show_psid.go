package cmd

import (
	"GoSungrow/iSolarCloud"
	"GoSungrow/iSolarCloud/AppService/getPsDetail"
	"GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
	"strings"
)


func (c *CmdShow) AttachPs(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		var self = &cobra.Command{
			Use:                   "ps",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "PsId"},
			Short:                 fmt.Sprintf("Ps related Sungrow commands."),
			Long:                  fmt.Sprintf("Ps related Sungrow commands."),
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

		c.AttachPsList(self)
		c.AttachPsTree(self)
		c.AttachPsDetail(self)
		c.AttachPsPoints(self)
		c.AttachPsData(self)
		c.AttachPsGraph(self)
		c.AttachPsSave(self)
	}
	return c.SelfCmd
}


func (c *CmdShow) AttachPsList(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all devices on account."),
		Long:                  fmt.Sprintf("Show all devices on account."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsList,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsList(_ *cobra.Command, args []string) error {
	for range Only.Once {
		var devices string
		devices, c.Error = cmds.Api.SunGrow.PsList(args...)
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", devices)
	}
	return c.Error
}

func (c *CmdShow) AttachPsIdList2(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "list2",
		Aliases:               []string{"ls"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show all available PS."),
		Long:                  fmt.Sprintf("Show all available PS."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsIdList2,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsIdList2(_ *cobra.Command, args []string) error {
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
		RunE:                  c.funcPsTree,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsTree(_ *cobra.Command, args []string) error {
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
		Use:                   "points [ps_ids | .] [device_type]",
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("List points used for a given ps_id."),
		Long:                  fmt.Sprintf("List points used for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsPoints,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsPoints(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(2, args)
		var points string
		points, c.Error = cmds.Api.SunGrow.PsPoints(strings.Split(args[0], ","), args[1])
		if c.Error != nil {
			break
		}

		fmt.Printf("%s\n", points)
	}
	return c.Error
}

func (c *CmdShow) AttachPsData(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "data <ps_ids | .> [device_type | .] " + ArgsDateInterval,
		Aliases:               []string{"point"},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Generate points table for a given ps_id."),
		Long:                  fmt.Sprintf("Generate points table for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsData,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsData(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetTable()
		args = MinimumArraySize(5, args)
		c.Error = cmds.Api.SunGrow.PsPointsData(strings.Split(args[0], ","), args[1], args[2], args[3], args[4])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPsGraph(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "graph <ps_ids | .> [device_type | .] " + ArgsDateInterval,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Generate graphs of points for a given ps_id."),
		Long:                  fmt.Sprintf("Generate graphs of points for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsGraph,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsGraph(_ *cobra.Command, args []string) error {
	for range Only.Once {
		cmds.Api.SunGrow.OutputType.SetGraph()
		args = MinimumArraySize(5, args)
		c.Error = cmds.Api.SunGrow.PsPointsData(strings.Split(args[0], ","), args[1], args[2], args[3], args[4])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPsSave(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "save <ps_ids | .> [device_type | .] " + ArgsDateInterval,
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Generate and save data points for a given ps_id."),
		Long:                  fmt.Sprintf("Generate and save data points for a given ps_id."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsSave,
		Args:                  cobra.MinimumNArgs(1),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsSave(_ *cobra.Command, args []string) error {
	for range Only.Once {
		args = MinimumArraySize(5, args)
		c.Error = cmds.Api.SunGrow.PsPointsDataSave(strings.Split(args[0], ","), args[1], args[2], args[3], args[4])
		if c.Error != nil {
			break
		}
	}
	return c.Error
}

func (c *CmdShow) AttachPsDetail(cmd *cobra.Command) *cobra.Command {
	var self = &cobra.Command{
		Use:                   "detail [ps_id ...]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "PsId"},
		Short:                 fmt.Sprintf("Show detailed info on PS."),
		Long:                  fmt.Sprintf("Show detailed info on PS."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               cmds.SunGrowArgs,
		RunE:                  c.funcPsDetail,
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(self)
	self.Example = cmdHelp.PrintExamples(self, "")

	return cmd
}
func (c *CmdShow) funcPsDetail(_ *cobra.Command, args []string) error {
	for range Only.Once {
		pids := cmds.Api.SunGrow.SetPsIds(args...)
		if c.Error != nil {
			break
		}

		data := cmds.Api.SunGrow.NewSunGrowData()
		data.SetPsIds(pids.Strings()...)
		data.SetEndpoints(getPsDetail.EndPointName, getPsDetailWithPsType.EndPointName)

		c.Error = data.GetData()
		if c.Error != nil {
			break
		}

		c.Error = data.Output()
		if c.Error != nil {
			break
		}
	}
	return c.Error
}
