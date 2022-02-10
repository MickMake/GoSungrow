package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/deviceProfile"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasProfiles = []string{"profiles", "prof"}

// ******************************************************************************** //
var cmdCountProfiles = &cobra.Command{
	Use:                   web.GetStructName(deviceProfile.Profile{}),
	Aliases:               aliasProfiles,
	Short:                 fmt.Sprintf("Count all device profiles"),
	Long:                  fmt.Sprintf("Count all device profiles"),
	Example:               fmt.Sprintf("%s count profiles", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountProfilesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountProfilesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.CountProfiles()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListProfiles = &cobra.Command{
	Use:     web.GetStructName(deviceProfile.Profile{}),
	Aliases: aliasProfiles,
	Short:   fmt.Sprintf("* List all device profiles"),
	Long:    fmt.Sprintf("* List all device profiles"),
	//Example:               fmt.Sprintf("%s ls profiles", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListProfilesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListProfilesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ListProfiles()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadProfiles = &cobra.Command{
	Use:                   web.GetStructName(deviceProfile.Profile{}),
	Aliases:               aliasProfiles,
	Short:                 fmt.Sprintf("Read all device profiles"),
	Long:                  fmt.Sprintf("Read all device profiles"),
	Example:               fmt.Sprintf("%s ls profiles", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadProfilesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadProfilesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ReadProfiles()
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateProfiles = &cobra.Command{
//	Use:                   web.GetStructName(deviceProfile.Profile{}),
//	Aliases:               aliasProfiles,
//	Short:                 fmt.Sprintf("Update Google sheet: device profiles"),
//	Long:                  fmt.Sprintf("Update Google sheet: device profiles"),
//	Example:               fmt.Sprintf("%s update profiles", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateProfilesFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateProfilesFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("profile")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
