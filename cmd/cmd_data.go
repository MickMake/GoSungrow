package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/cmdConfig"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdData CmdDefault

func NewCmdData() *CmdData {
	var ret *CmdData

	for range Only.Once {
		ret = &CmdData{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *CmdData) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "data",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("High-level iSolarCloud functions."),
			Long:                  fmt.Sprintf("High-level iSolarCloud functions."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "get <endpoint>", "put <endpoint>")
		
		// ********************************************************************************
		var cmdDataGet = &cobra.Command{
			Use:                   "get",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (table)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (table)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdDataGet)
		cmdDataGet.Example = cmdHelp.PrintExamples(cmdDataGet, "[area.]<endpoint>")
		c.AttachCmdDataStats(cmdDataGet)
		c.AttachCmdDataTemplate(cmdDataGet)
		c.AttachCmdDataPoints(cmdDataGet)
		c.AttachCmdDataRealTime(cmdDataGet)
		c.AttachCmdDataPsDetails(cmdDataGet)
		c.AttachCmdDataAll(cmdDataGet)

		// ********************************************************************************
		var cmdDataRaw = &cobra.Command{
			Use:                   output.StringTypeRaw,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (raw)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (raw)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdDataRaw)
		cmdDataRaw.Example = cmdHelp.PrintExamples(cmdDataRaw, "[area.]<endpoint>")
		c.AttachCmdDataStats(cmdDataRaw)
		c.AttachCmdDataTemplate(cmdDataRaw)
		c.AttachCmdDataPoints(cmdDataRaw)
		c.AttachCmdDataRealTime(cmdDataRaw)
		c.AttachCmdDataPsDetails(cmdDataRaw)
		c.AttachCmdDataAll(cmdDataRaw)

		// ********************************************************************************
		var cmdDataJson = &cobra.Command{
			Use:                   output.StringTypeJson,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdDataJson)
		cmdDataJson.Example = cmdHelp.PrintExamples(cmdDataJson, "[area.]<endpoint>")
		c.AttachCmdDataStats(cmdDataJson)
		c.AttachCmdDataTemplate(cmdDataJson)
		c.AttachCmdDataPoints(cmdDataJson)
		c.AttachCmdDataRealTime(cmdDataJson)
		c.AttachCmdDataPsDetails(cmdDataJson)
		c.AttachCmdDataAll(cmdDataJson)

		// ********************************************************************************
		var cmdDataCsv = &cobra.Command{
			Use:                   output.StringTypeCsv,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (json)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdDataCsv)
		cmdDataCsv.Example = cmdHelp.PrintExamples(cmdDataCsv, "[area.]<endpoint>")
		c.AttachCmdDataStats(cmdDataCsv)
		c.AttachCmdDataTemplate(cmdDataCsv)
		c.AttachCmdDataPoints(cmdDataCsv)
		c.AttachCmdDataRealTime(cmdDataCsv)
		c.AttachCmdDataPsDetails(cmdDataCsv)
		c.AttachCmdDataAll(cmdDataCsv)

		// ********************************************************************************
		var cmdDataGraph = &cobra.Command{
			Use:                   output.StringTypeGraph,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Get high-level data from iSolarCloud (graph)"),
			Long:                  fmt.Sprintf("Get high-level data from iSolarCloud (graph)"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		c.SelfCmd.AddCommand(cmdDataGraph)
		cmdDataGraph.Example = cmdHelp.PrintExamples(cmdDataGraph, "[area.]<endpoint> ''")
		c.AttachCmdDataStats(cmdDataGraph)
		c.AttachCmdDataTemplate(cmdDataGraph)
		c.AttachCmdDataPoints(cmdDataGraph)
		c.AttachCmdDataRealTime(cmdDataGraph)
		// c.AttachCmdDataAll(cmdDataGraph)

		// ********************************************************************************
		var cmdDataPut = &cobra.Command{
			Use:                   "put",
			Aliases:               []string{"set", "write"},
			Annotations:           map[string]string{"group": "Data"},
			Short:                 fmt.Sprintf("Set high-level data on iSolarCloud"),
			Long:                  fmt.Sprintf("Set high-level data on iSolarCloud"),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			Run:                   cmds.CmdDataPut,
			Args:                  cobra.ExactArgs(2),
		}
		c.SelfCmd.AddCommand(cmdDataPut)
		cmdDataPut.Example = cmdHelp.PrintExamples(cmdDataPut, "[area.]<endpoint> <value>")
	}
	return c.SelfCmd
}


func (ca *Cmds) CmdDataPut(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// ca.Api.SunGrow.OutputType.SetFile()
		// args = cmdConfig.FillArray(2, args)
		// c.Error = SunGrow.PutHighLevel(args[0], args[1])
	}
}


func (c *CmdData) AttachCmdDataTemplate(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "template <template_id> <date> [filter]",
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			args = cmdConfig.FillArray(3, args)
			return cmds.Api.SunGrow.GetTemplateData(args[0], args[1], args[2])
		},
		Args:                  cobra.RangeArgs(2, 3),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "8042 20220212", "8042 20220212 '{\"search_string\":\"p83106\",\"min_left_axis\":-6000,\"max_left_axis\":12000}'")

	return cmd
}

func (c *CmdData) AttachCmdDataStats(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "stats",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
		Long:                  fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			return cmds.Api.SunGrow.PrintCurrentStats()
		},
		Args:                  cobra.ExactArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdData) AttachCmdDataPoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var cmdDataPoints = &cobra.Command{
		Use:                   "points <date> <device_id.point_id> ...",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get points data for a specific date."),
		Long:                  fmt.Sprintf("Get points data for a specific date."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			return cmds.Api.SunGrow.GetPointData(args[0], api.CreatePoints(args[1:]))
		},
		Args:                  cobra.MinimumNArgs(2),
	}
	cmd.AddCommand(cmdDataPoints)
	cmdDataPoints.Example = cmdHelp.PrintExamples(cmdDataPoints, "20220202 1129147.p13019 1129147.p83106")

	return cmd
}

func (c *CmdData) AttachCmdDataRealTime(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "real-time",
		Aliases:               []string{"realtime"},
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get iSolarCloud real-time data."),
		Long:                  fmt.Sprintf("Get iSolarCloud real-time data."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			args = cmdConfig.FillArray(1, args)
			return cmds.Api.SunGrow.GetRealTimeData(args[0])
		},
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdData) AttachCmdDataPsDetails(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "psdetails",
		Aliases:               []string{"ps-details"},
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get iSolarCloud ps details."),
		Long:                  fmt.Sprintf("Get iSolarCloud ps details."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			pids, err := cmds.Api.SunGrow.StringToPids(args...)
			if err != nil {
				return err
			}
			return cmds.Api.SunGrow.CmdDataPsDetail(pids...)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdData) AttachCmdDataAll(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "all",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Data"},
		Short:                 fmt.Sprintf("Get all iSolarCloud ps details."),
		Long:                  fmt.Sprintf("Get all iSolarCloud ps details."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               func(cmd *cobra.Command, args []string) error {
			cmds.Error = cmds.ProcessArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			cmds.Error = cmds.SunGrowArgs(cmd, args)
			if cmds.Error != nil {
				return cmds.Error
			}
			return nil
		},
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = cmds.SetOutputType(cmd)
			// pids, err := cmds.Api.SunGrow.StringToPids(args...)
			// if err != nil {
			// 	return err
			// }
			return cmds.Api.SunGrow.GetEndpoints(args)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}
