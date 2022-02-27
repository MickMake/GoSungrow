package cmd

import (
	"GoSungrow/iSolarCloud/api"
	"fmt"
	"github.com/spf13/cobra"
)


func AttachCmdDataTemplatePoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var cmdDataTemplatePoints = &cobra.Command{
		Use:                   "template-points <template_id>",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("List data points used in report template."),
		Long:                  fmt.Sprintf("List data points used in report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = SwitchOutput(cmd)
			return Cmd.SunGrow.GetTemplatePoints(args[0])
		},
		Args:                  cobra.ExactArgs(1),
	}
	cmd.AddCommand(cmdDataTemplatePoints)
	cmdDataTemplatePoints.Example = PrintExamples(cmdDataTemplatePoints, "8042", "8040")

	return cmd
}

func AttachCmdDataTemplate(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var cmdDataGetTemplate = &cobra.Command{
		Use:                   "template <template_id> <date> [filter]",
		Short:                 fmt.Sprintf("Get data from report template."),
		Long:                  fmt.Sprintf("Get data from report template."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = SwitchOutput(cmd)
			return Cmd.SunGrow.GetTemplateData(args[0], args[1], args[2])
		},
		Args:                  cobra.MinimumNArgs(2),
	}
	cmd.AddCommand(cmdDataGetTemplate)
	cmdDataGetTemplate.Example = PrintExamples(cmdDataGetTemplate, "8042 20220212", "8042 20220212 '{\"search_string\":\"p83106\",\"min_left_axis\":-6000,\"max_left_axis\":12000}'")

	return cmd
}

func AttachCmdDataStats(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var cmdDataStats = &cobra.Command{
		Use:                   "stats",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
		Long:                  fmt.Sprintf("Get current inverter stats, (last 5 minutes)."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = SwitchOutput(cmd)
			return Cmd.SunGrow.GetCurrentStats()
		},
		Args:                  cobra.ExactArgs(0),
	}
	cmd.AddCommand(cmdDataStats)
	cmdDataStats.Example = PrintExamples(cmdDataStats, "")

	return cmd
}

func AttachCmdDataPoints(cmd *cobra.Command) *cobra.Command {
	// ********************************************************************************
	var cmdDataStats = &cobra.Command{
		Use:                   "points <date> <device_id.point_id> ...",
		Aliases:               []string{""},
		Short:                 fmt.Sprintf("Get points data for a specific date."),
		Long:                  fmt.Sprintf("Get points data for a specific date."),
		DisableFlagParsing:    false,
		DisableFlagsInUseLine: false,
		PreRunE:               Cmd.SunGrowArgs,
		RunE:                  func(cmd *cobra.Command, args []string) error {
			_ = SwitchOutput(cmd)
			return Cmd.SunGrow.GetPointData(args[0], api.CreatePoints(args[1:]))
		},
		Args:                  cobra.MinimumNArgs(2),
	}
	cmd.AddCommand(cmdDataStats)
	cmdDataStats.Example = PrintExamples(cmdDataStats, "20220202 1129147.p13019 1129147.p83106")

	return cmd
}
