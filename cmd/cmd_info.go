package cmd

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/output"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdConfig"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


//goland:noinspection GoNameStartsWithPackageName
type CmdInfo CmdDefault

func NewCmdInfo() *CmdInfo {
	var ret *CmdInfo

	for range Only.Once {
		ret = &CmdInfo{
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *CmdInfo) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "info",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("General iSolarCloud functions."),
			Long:                  fmt.Sprintf("General iSolarCloud functions."),
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
		var cmdInfoGet = &cobra.Command{
			Use:                   "get",
			Aliases:               []string{output.StringTypeTable},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (table)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (table)"),
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
		c.SelfCmd.AddCommand(cmdInfoGet)
		cmdInfoGet.Example = cmdHelp.PrintExamples(cmdInfoGet, "[area.]<endpoint>")
		c.AttachCmdInfo(cmdInfoGet)

		// ********************************************************************************
		var cmdInfoRaw = &cobra.Command{
			Use:                   output.StringTypeRaw,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (raw)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (raw)"),
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
		c.SelfCmd.AddCommand(cmdInfoRaw)
		cmdInfoRaw.Example = cmdHelp.PrintExamples(cmdInfoRaw, "[area.]<endpoint>")
		c.AttachCmdInfo(cmdInfoRaw)

		// ********************************************************************************
		var cmdInfoJson = &cobra.Command{
			Use:                   output.StringTypeJson,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (json)"),
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
		c.SelfCmd.AddCommand(cmdInfoJson)
		cmdInfoJson.Example = cmdHelp.PrintExamples(cmdInfoJson, "[area.]<endpoint>")
		c.AttachCmdInfo(cmdInfoJson)

		// ********************************************************************************
		var cmdInfoCsv = &cobra.Command{
			Use:                   output.StringTypeCsv,
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Get info from iSolarCloud (json)"),
			Long:                  fmt.Sprintf("Get info from iSolarCloud (json)"),
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
		c.SelfCmd.AddCommand(cmdInfoCsv)
		cmdInfoCsv.Example = cmdHelp.PrintExamples(cmdInfoCsv, "[area.]<endpoint>")
		c.AttachCmdInfo(cmdInfoCsv)

		// ********************************************************************************
		var cmdInfoPut = &cobra.Command{
			Use:                   "put",
			Aliases:               []string{"set", "write"},
			Annotations:           map[string]string{"group": "Info"},
			Short:                 fmt.Sprintf("Set info on iSolarCloud"),
			Long:                  fmt.Sprintf("Set info on iSolarCloud"),
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
			Run:                   cmds.CmdInfoPut,
			Args:                  cobra.ExactArgs(2),
		}
		c.SelfCmd.AddCommand(cmdInfoPut)
		cmdInfoPut.Example = cmdHelp.PrintExamples(cmdInfoPut, "[area.]<endpoint> <value>")
	}
	return c.SelfCmd
}

func (c *CmdInfo) AttachCmdInfo(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}

		c.AttachCmdInfoPointNames(cmd)
		c.AttachCmdInfoMqtt(cmd)
		c.AttachCmdInfoSearchPointInfo(cmd)
		c.AttachCmdInfoDevices(cmd)
		c.AttachCmdInfoDeviceModels(cmd)
		c.AttachCmdInfoTemplates(cmd)
		c.AttachCmdInfoTemplatePoints(cmd)
		c.AttachCmdInfoGetDevicePoints(cmd)

		c.AttachCmdInfoStats(cmd)
		c.AttachCmdInfoTemplate(cmd)
		c.AttachCmdInfoPoints(cmd)
		c.AttachCmdInfoRealTime(cmd)
		c.AttachCmdInfoPsDetails(cmd)

	}
	return cmd
}

func (ca *Cmds) CmdInfoPut(_ *cobra.Command, _ []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")
		// ca.Api.SunGrow.OutputType.SetFile()
		// args = cmdConfig.FillArray(2, args)
		// c.Error = SunGrow.PutHighLevel(args[0], args[1])
	}
}


func (c *CmdInfo) AttachCmdInfoMqtt(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "mqtt",
		Aliases:               []string{"mqtt-server"},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get iSolarCloud MQTT service login details."),
		Long:                  fmt.Sprintf("Get iSolarCloud MQTT service login details."),
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
			return cmds.Api.SunGrow.GetIsolarcloudMqtt(args[0])
		},
		Args:                  cobra.RangeArgs(0, 1),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}


func (c *CmdInfo) AttachCmdInfoPointNames(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "point-names",
		Aliases:               []string{"names"},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get iSolarCloud point names."),
		Long:                  fmt.Sprintf("Get iSolarCloud point names."),
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
			return cmds.Api.SunGrow.GetPointNames(args...)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdInfo) AttachCmdInfoSearchPointInfo(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "search-point-names",
		Aliases:               []string{"names"},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get iSolarCloud search point names."),
		Long:                  fmt.Sprintf("Get iSolarCloud search point names."),
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
			return cmds.Api.SunGrow.SearchPointNames(args...)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}


func (c *CmdInfo) AttachCmdInfoDevices(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "devices",
		Aliases:               []string{"device"},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get iSolarCloud devices."),
		Long:                  fmt.Sprintf("Get iSolarCloud devices."),
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
			return cmds.Api.SunGrow.GetDeviceList(pids...)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdInfo) AttachCmdInfoDeviceModels(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "models",
		Aliases:               []string{"model"},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get ALL iSolarCloud models."),
		Long:                  fmt.Sprintf("Get ALL iSolarCloud models."),
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
			return cmds.Api.SunGrow.GetDeviceModelInfoList()
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}


func (c *CmdInfo) AttachCmdInfoTemplates(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "templates",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("Get all defined templates."),
		Long:                  fmt.Sprintf("Get all defined templates."),
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
			return cmds.Api.SunGrow.GetTemplates()
		},
		Args:                  cobra.ExactArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "")

	return cmd
}

func (c *CmdInfo) AttachCmdInfoTemplatePoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "template-points <template_id>",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("List data points used in report template."),
		Long:                  fmt.Sprintf("List data points used in report template."),
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
			return cmds.Api.SunGrow.GetTemplatePoints(args[0])
		},
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "8042", "8040")

	return cmd
}


func (c *CmdInfo) AttachCmdInfoGetDevicePoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var c2 = &cobra.Command{
		Use:                   "device-points [ps_id]",
		Aliases:               []string{},
		Annotations:           map[string]string{"group": "Info"},
		Short:                 fmt.Sprintf("List all available device data points."),
		Long:                  fmt.Sprintf("List all available device data points."),
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
			return cmds.Api.SunGrow.GetDevicePoints(pids...)
		},
		Args:                  cobra.MinimumNArgs(0),
	}
	cmd.AddCommand(c2)
	c2.Example = cmdHelp.PrintExamples(c2, "", "1129147")

	return cmd
}


func (c *CmdInfo) AttachCmdInfoTemplate(cmd *cobra.Command) *cobra.Command {
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

func (c *CmdInfo) AttachCmdInfoStats(cmd *cobra.Command) *cobra.Command {
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

func (c *CmdInfo) AttachCmdInfoPoints(cmd *cobra.Command) *cobra.Command {
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

func (c *CmdInfo) AttachCmdInfoRealTime(cmd *cobra.Command) *cobra.Command {
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

func (c *CmdInfo) AttachCmdInfoPsDetails(cmd *cobra.Command) *cobra.Command {
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
