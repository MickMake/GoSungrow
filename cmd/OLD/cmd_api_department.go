package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/department"
	"GoSungro/iSolarCloud/web"
	"fmt"
	"github.com/spf13/cobra"
)

var aliasDepartments = []string{"departments", "dept"}

// ******************************************************************************** //
var cmdCountDepartments = &cobra.Command{
	Use:     web.GetStructName(department.Department{}),
	Aliases: aliasDepartments,
	Short:   fmt.Sprintf("* Count all departments"),
	Long:    fmt.Sprintf("* Count all departments "),
	//Example:               fmt.Sprintf("%s count departments", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCountDepartmentsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdCountDepartmentsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.CountDepartments(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdListDepartments = &cobra.Command{
	Use:                   web.GetStructName(department.Department{}),
	Aliases:               aliasDepartments,
	Short:                 fmt.Sprintf("List all departments"),
	Long:                  fmt.Sprintf("List all departments "),
	Example:               fmt.Sprintf("%s ls departments", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdListDepartmentsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdListDepartmentsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ListDepartments(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

// ******************************************************************************** //
var cmdReadDepartments = &cobra.Command{
	Use:                   web.GetStructName(department.Department{}),
	Aliases:               aliasDepartments,
	Short:                 fmt.Sprintf("Read all departments"),
	Long:                  fmt.Sprintf("Read all departments "),
	Example:               fmt.Sprintf("%s ls departments", DefaultBinaryName),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdReadDepartmentsFunc,
	Args:                  cobra.RangeArgs(0, 1),
}

//goland:noinspection GoUnusedParameter
func cmdReadDepartmentsFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		args = fillArray(1, args)
		Cmd.Error = SunGro.ReadDepartments(args[0])
		if Cmd.Error != nil {
			break
		}
	}
}

//// ******************************************************************************** //
//var cmdUpdateDepartments = &cobra.Command{
//	Use:                   web.GetStructName(department.Department{}),
//	Aliases:               aliasDepartments,
//	Short:                 fmt.Sprintf("Update Google sheet: departments"),
//	Long:                  fmt.Sprintf("Update Google sheet: departments "),
//	Example:               fmt.Sprintf("%s update departments", DefaultBinaryName),
//	DisableFlagParsing:    false,
//	DisableFlagsInUseLine: false,
//	PreRunE:               Cmd.ProcessArgs,
//	Run:                   cmdUpdateDepartmentsFunc,
//	Args:                  cobra.RangeArgs(0, 1),
//}
//
////goland:noinspection GoUnusedParameter
//func cmdUpdateDepartmentsFunc(cmd *cobra.Command, args []string) {
//	for range Only.Once {
//		Cmd.Error = Cmd.GoogleUpdate("department")
//		if Cmd.Error != nil {
//			break
//		}
//	}
//}
