package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/presence"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasPresence = []string{"state", "online"}

// ******************************************************************************** //
var cmdCountPresence = &cobra.Command{
	Use:     web.GetStructName(presence.Presence{}),
	Aliases: aliasPresence,
	Short:   fmt.Sprintf("* Count all presence"),
	Long:    fmt.Sprintf("* Count all presence "),
	//Example:               fmt.Sprintf("%s count presence", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountPresenceFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountPresenceFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountPresence(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListPresence = &cobra.Command{
	Use:                   web.GetStructName(presence.Presence{}),
	Aliases:               aliasPresence,
	Short:                 fmt.Sprintf("List all presence"),
	Long:                  fmt.Sprintf("List all presence "),
	Example:               fmt.Sprintf("%s ls presence", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListPresenceFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListPresenceFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListPresence(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadPresence = &cobra.Command{
	Use:     web.GetStructName(presence.Presence{}),
	Aliases: aliasPresence,
	Short:   fmt.Sprintf("* Read all presence"),
	Long:    fmt.Sprintf("* Read all presence "),
	//Example:               fmt.Sprintf("%s ls presence", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadPresenceFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadPresenceFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadPresence(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdatePresence = &cobra.Command{
//	Use:                   web.GetStructName(presence.Presence{}),
//	Aliases:               aliasPresence,
//	Short:                 fmt.Sprintf("Update Google sheet: presence"),
//	Long:                  fmt.Sprintf("Update Google sheet: presence "),
//	Example:               fmt.Sprintf("%s update presence", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdatePresenceFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdatePresenceFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("presence")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
