package cmd

import (
	"GoSungrow/Only"
	"fmt"
	"github.com/spf13/cobra"
)


// ********************************************************************************
var cmdDataStats = &cobra.Command{
	Use: "stats",
	// Aliases:               []string{""},
	Short:                 fmt.Sprintf("Get live data from iSolarCloud"),
	Long:                  fmt.Sprintf("Get live data from iSolarCloud"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.SunGrowArgs,
	Run:                   cmdDataStatsFunc,
	Args:                  cobra.MinimumNArgs(1),
}

func cmdDataStatsFunc(_ *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.SunGrow.OutputType.SetFile()

		Cmd.Error = Cmd.SunGrow.GetCurrentStats()
	}
}

