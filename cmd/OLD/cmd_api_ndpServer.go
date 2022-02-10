package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/ndpServer"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasNdpServers = []string{"ndpservers", "ndp"}

// ******************************************************************************** //
var cmdCountNdpServers = &cobra.Command{
	Use:                   web.GetStructName(ndpServer.Ndp{}),
	Aliases:               aliasNdpServers,
	Short:                 fmt.Sprintf("Count all NdpServers"),
	Long:                  fmt.Sprintf("Count all NdpServers"),
	Example:               fmt.Sprintf("%s count ndp", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountNdpServersFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountNdpServersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountNdpServers()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListNdpServers = &cobra.Command{
	Use:                   web.GetStructName(ndpServer.Ndp{}),
	Aliases:               aliasNdpServers,
	Short:                 fmt.Sprintf("List all NdpServers"),
	Long:                  fmt.Sprintf("List all NdpServers"),
	Example:               fmt.Sprintf("%s ls ndp", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListNdpServersFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListNdpServersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ListNdpServer()
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadNdpServers = &cobra.Command{
	Use:                   web.GetStructName(ndpServer.Ndp{}),
	Aliases:               aliasNdpServers,
	Short:                 fmt.Sprintf("Read all NdpServers"),
	Long:                  fmt.Sprintf("Read all NdpServers"),
	Example:               fmt.Sprintf("%s read profiles", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadNdpServersFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadNdpServersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		Cmd.Error = SunGro.ReadNdpServer()
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateNdpServers = &cobra.Command{
//	Use:                   web.GetStructName(ndpServer.Ndp{}),
//	Aliases:               aliasNdpServers,
//	Short:                 fmt.Sprintf("Update Google sheet: NdpServers"),
//	Long:                  fmt.Sprintf("Update Google sheet: NdpServers"),
//	Example:               fmt.Sprintf("%s update ndp", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateNdpServersFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateNdpServersFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("ndpserver")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
