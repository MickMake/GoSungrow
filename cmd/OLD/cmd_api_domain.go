package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/domain"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasDomains = []string{"domains"}

// ******************************************************************************** //
var cmdCountDomain = &cobra.Command{
	Use:                   web.GetStructName(domain.Domain{}),
	Aliases:               aliasDomains,
	Short:                 fmt.Sprintf("Count info on domain."),
	Long:                  fmt.Sprintf("Count info on domain. "),
	Example:               fmt.Sprintf("%s count domain", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountDomainFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountDomainFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountDomain(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListDomain = &cobra.Command{
	Use:     web.GetStructName(domain.Domain{}),
	Aliases: aliasDomains,
	Short:   fmt.Sprintf("* List info on domain."),
	Long:    fmt.Sprintf("* List info on domain. "),
	//Example:               fmt.Sprintf("%s ls domain", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListDomainFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListDomainFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListDomain(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadDomain = &cobra.Command{
	Use:                   web.GetStructName(domain.Domain{}),
	Aliases:               aliasDomains,
	Short:                 fmt.Sprintf("Read all Domains"),
	Long:                  fmt.Sprintf("Read all Domains"),
	Example:               fmt.Sprintf("%s read contact", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadDomainsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadDomainsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadDomain(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateDomain = &cobra.Command{
//	Use:                   web.GetStructName(domain.ApiAppKey{}),
//	Aliases:               aliasDomains,
//	Short:                 fmt.Sprintf("Update Google sheet: domain."),
//	Long:                  fmt.Sprintf("Update Google sheet: domain. "),
//	Example:               fmt.Sprintf("%s update domain", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateDomainFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateDomainFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("domain")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
