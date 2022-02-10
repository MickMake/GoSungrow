package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/site"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasSites = []string{"sites"}

// ******************************************************************************** //
var cmdCountSites = &cobra.Command{
	Use:                   web.GetStructName(site.Site{}),
	Aliases:               aliasSites,
	Short:                 fmt.Sprintf("Count all sites"),
	Long:                  fmt.Sprintf("Count all sites "),
	Example:               fmt.Sprintf("%s count sites", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountSitesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountSitesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountSites(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListSites = &cobra.Command{
	Use:                   web.GetStructName(site.Site{}),
	Aliases:               aliasSites,
	Short:                 fmt.Sprintf("List all sites"),
	Long:                  fmt.Sprintf("List all sites "),
	Example:               fmt.Sprintf("%s ls sites", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListSitesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListSitesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListSites(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadSites = &cobra.Command{
	Use:                   web.GetStructName(site.Site{}),
	Aliases:               aliasSites,
	Short:                 fmt.Sprintf("Read all sites"),
	Long:                  fmt.Sprintf("Read all sites "),
	Example:               fmt.Sprintf("%s read sites", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadSitesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadSitesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadSites(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateSites = &cobra.Command{
//	Use:                   web.GetStructName(site.Site{}),
//	Aliases:               aliasSites,
//	Short:                 fmt.Sprintf("Update Google sheet: sites"),
//	Long:                  fmt.Sprintf("Update Google sheet: sites "),
//	Example:               fmt.Sprintf("%s update sites", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateSitesFunc,
//	Args:                  cobra.RangeArgs(0, 2),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateSitesFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("site")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
