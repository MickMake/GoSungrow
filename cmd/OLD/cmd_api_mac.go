package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/mac"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasMacs = []string{"macs"}

// ******************************************************************************** //
var cmdCountMacs = &cobra.Command{
	Use:                   web.GetStructName(mac.Mac{}),
	Aliases:               aliasMacs,
	Short:                 fmt.Sprintf("Count all Macs"),
	Long:                  fmt.Sprintf("Count all Macs"),
	Example:               fmt.Sprintf("%s count mac", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountMacsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountMacsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountMacs(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListMacs = &cobra.Command{
	Use:                   web.GetStructName(mac.Mac{}),
	Aliases:               aliasMacs,
	Short:                 fmt.Sprintf("List all Macs"),
	Long:                  fmt.Sprintf("List all Macs"),
	Example:               fmt.Sprintf("%s ls mac", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListMacsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListMacsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListMacs(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadMacs = &cobra.Command{
	Use:                   web.GetStructName(mac.Mac{}),
	Aliases:               aliasMacs,
	Short:                 fmt.Sprintf("Read all Macs"),
	Long:                  fmt.Sprintf("Read all Macs"),
	Example:               fmt.Sprintf("%s read mac", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadMacsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadMacsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadMacs(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateMacs = &cobra.Command{
//	Use:                   web.GetStructName(mac.Mac{}),
//	Aliases:               aliasMacs,
//	Short:                 fmt.Sprintf("Update Google sheet: Macs"),
//	Long:                  fmt.Sprintf("Update Google sheet: Macs"),
//	Example:               fmt.Sprintf("%s update ndp", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateMacsFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateMacsFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("macs")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
