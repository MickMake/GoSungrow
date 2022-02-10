package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/contact"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasContacts = []string{"contacts"}

// ******************************************************************************** //
var cmdCountContacts = &cobra.Command{
	Use:                   web.GetStructName(contact.Contact{}),
	Aliases:               aliasContacts,
	Short:                 fmt.Sprintf("Count all Contacts"),
	Long:                  fmt.Sprintf("Count all Contacts"),
	Example:               fmt.Sprintf("%s count contact", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountContactsFunc,
	Args:                  cobra.RangeArgs(0, 2),
}

//goland:noinspection GoUnusedParameter
func cmdCountContactsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(2, args)
		Cmd.Error = SunGro.CountContacts(args[0], args[1])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListContacts = &cobra.Command{
	Use:     web.GetStructName(contact.Contact{}),
	Aliases: aliasContacts,
	Short:   fmt.Sprintf("* List all Contacts"),
	Long:    fmt.Sprintf("* List all Contacts"),
	//Example:               fmt.Sprintf("%s ls contact", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListContactsFunc,
	Args:                  cobra.RangeArgs(0, 2),
}

//goland:noinspection GoUnusedParameter
func cmdListContactsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(2, args)
		Cmd.Error = SunGro.ListContacts(args[0], args[1])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadContacts = &cobra.Command{
	Use:                   web.GetStructName(contact.Contact{}),
	Aliases:               aliasContacts,
	Short:                 fmt.Sprintf("Read all Contacts"),
	Long:                  fmt.Sprintf("Read all Contacts"),
	Example:               fmt.Sprintf("%s ls contact", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadContactsFunc,
	Args:                  cobra.RangeArgs(0, 2),
}

//goland:noinspection GoUnusedParameter
func cmdReadContactsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(2, args)
		Cmd.Error = SunGro.ReadContacts(args[0], args[1])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateContacts = &cobra.Command{
//	Use:                   web.GetStructName(contact.Contact{}),
//	Aliases:               aliasContacts,
//	Short:                 fmt.Sprintf("Update Google sheet: Contacts"),
//	Long:                  fmt.Sprintf("Update Google sheet: Contacts"),
//	Example:               fmt.Sprintf("%s update contact", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateContactsFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateContactsFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("contact")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
