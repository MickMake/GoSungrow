package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/deviceModel"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasModels = []string{"models", "mod"}

// ******************************************************************************** //
var cmdCountModels = &cobra.Command{
	Use:                   web.GetStructName(deviceModel.Model{}),
	Aliases:               aliasModels,
	Short:                 fmt.Sprintf("Count all device models"),
	Long:                  fmt.Sprintf("Count all device models "),
	Example:               fmt.Sprintf("%s count models", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountModelsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountModelsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.CountModels()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListModels = &cobra.Command{
	Use:                   web.GetStructName(deviceModel.Model{}),
	Aliases:               aliasModels,
	Short:                 fmt.Sprintf("List all device models"),
	Long:                  fmt.Sprintf("List all device models "),
	Example:               fmt.Sprintf("%s ls models", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListModelsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListModelsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ListModels()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadModels = &cobra.Command{
	Use:                   web.GetStructName(deviceModel.Model{}),
	Aliases:               aliasModels,
	Short:                 fmt.Sprintf("Read all device profiles"),
	Long:                  fmt.Sprintf("Read all device profiles"),
	Example:               fmt.Sprintf("%s read profiles", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadModelsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadModelsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ReadModels()
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateModels = &cobra.Command{
//	Use:                   web.GetStructName(deviceModel.Model{}),
//	Aliases:               aliasModels,
//	Short:                 fmt.Sprintf("Update Google sheet: device models"),
//	Long:                  fmt.Sprintf("Update Google sheet: device models "),
//	Example:               fmt.Sprintf("%s update models", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateModelsFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateModelsFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("model")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
