package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/device"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasDevices = []string{"devices", "dev"}


// ******************************************************************************** //
var cmdCountDevices = &cobra.Command{
	Use:                   web.GetStructName(device.Device{}),
	Aliases:               aliasDevices,
	Short:                 fmt.Sprintf("Count all devices"),
	Long:                  fmt.Sprintf("Count all devices "),
	Example:               fmt.Sprintf("%s count devices", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountDevicesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdCountDevicesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountDevices(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdListDevices = &cobra.Command{
	Use:                   web.GetStructName(device.Device{}),
	Aliases:               aliasDevices,
	Short:                 fmt.Sprintf("List all devices"),
	Long:                  fmt.Sprintf("List all devices "),
	Example:               fmt.Sprintf("%s ls devices", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListDevicesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdListDevicesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListDevices(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdReadDevices = &cobra.Command{
	Use:                   web.GetStructName(device.Device{}),
	Aliases:               aliasDevices,
	Short:                 fmt.Sprintf("Read all device models"),
	Long:                  fmt.Sprintf("Read all device models "),
	Example:               fmt.Sprintf("%s read models", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadDevicesFunc,
	Args:                  cobra.RangeArgs(0, 1),
}
//goland:noinspection GoUnusedParameter
func cmdReadDevicesFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadDevices(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}


//// ******************************************************************************** //
//var cmdUpdateDevices = &cobra.Command{
//	Use:                   web.GetStructName(device.Device{}),
//	Aliases:               aliasDevices,
//	Short:                 fmt.Sprintf("Update Google sheet: devices"),
//	Long:                  fmt.Sprintf("Update Google sheet: devices "),
//	Example:               fmt.Sprintf("%s update devices", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:cmdUpdateDevicesFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
////goland:noinspection GoUnusedParameter
//func cmdUpdateDevicesFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("device")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
