package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/subscriber"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasSubscribers = []string{"users", "user", "subscribers", "sub"}

// ******************************************************************************** //
var cmdCountUsers = &cobra.Command{
	Use:                   web.GetStructName(subscriber.Subscriber{}),
	Aliases:               aliasSubscribers,
	Short:                 fmt.Sprintf("Count all users"),
	Long:                  fmt.Sprintf("Count all users "),
	Example:               fmt.Sprintf("%s count users", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountUsersFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountUsersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountUsers(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListUsers = &cobra.Command{
	Use:                   web.GetStructName(subscriber.Subscriber{}),
	Aliases:               aliasSubscribers,
	Short:                 fmt.Sprintf("List all users"),
	Long:                  fmt.Sprintf("List all users "),
	Example:               fmt.Sprintf("%s ls users", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListUsersFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListUsersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListUsers(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadUsers = &cobra.Command{
	Use:                   web.GetStructName(subscriber.Subscriber{}),
	Aliases:               aliasSubscribers,
	Short:                 fmt.Sprintf("Read all users"),
	Long:                  fmt.Sprintf("Read all users "),
	Example:               fmt.Sprintf("%s ls users", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadUsersFunc,
	Args:                  cobra.RangeArgs(0, 2),
}

//goland:noinspection GoUnusedParameter
func cmdReadUsersFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(2, args)
		Cmd.Error = SunGro.ReadUsers(args[0], args[1])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateUsers = &cobra.Command{
//	Use:                   web.GetStructName(subscriber.Subscriber{}),
//	Aliases:               aliasSubscribers,
//	Short:                 fmt.Sprintf("Update Google sheet: users"),
//	Long:                  fmt.Sprintf("Update Google sheet: users "),
//	Example:               fmt.Sprintf("%s update users", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateUsersFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateUsersFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("user")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
