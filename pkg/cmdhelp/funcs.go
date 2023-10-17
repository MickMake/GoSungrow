package cmdhelp

import (
	"fmt"
	"os"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// func (h *Help) PrintConfig(cmd *cobra.Command) {
// 	for range only.Once {
// 		// fmt.Printf("Config file '%s':\n", Cmd.ConfigFile)	// @TODO - fixup.
//
// 		table := tablewriter.NewWriter(os.Stdout)
// 		table.SetHeader([]string{"Flag", "Short flag", "Environment", "Description", "Value"})
// 		table.SetBorder(true)
//
// 		cmd.PersistentFlags().SortFlags = false
// 		cmd.Flags().SortFlags = false
// 		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
// 			if flag.Hidden {
// 				return
// 			}
//
// 			sh := ""
// 			if flag.Shorthand != "" {
// 				sh = "-" + flag.Shorthand
// 			}
//
// 			// fmt.Printf("key: %s => %v (%s)\n", flag.Name, flag.Value, flag.Value.String())
// 			table.Append([]string{
// 				"--" + flag.Name,
// 				sh,
// 				PrintFlagEnv(h.EnvPrefix, flag.Name),
// 				flag.Usage,
// 				flag.Value.String(),
// 			})
// 		})
//
// 		table.Render()
// 	}
// }

func PrintFlagEnv(prefix string, flag string) string {
	fenv := strings.ReplaceAll(flag, "-", "_")
	fenv = strings.ToUpper(fenv)

	// ret := fmt.Sprintf("--%s\t%s_%s\n", flag, EnvPrefix, fenv)
	// ret := fmt.Sprintf("%s_%s", defaults.EnvPrefix, fenv)
	// ret := fmt.Sprintf("%s_%s", cmdVersion.GetEnvPrefix(), fenv)
	ret := fmt.Sprintf("%s_%s", prefix, fenv)
	return ret
}

func (h *Help) PrintConfig(cmd *cobra.Command) {
	PrintConfig(cmd, h.EnvPrefix)
}

func PrintConfig(cmd *cobra.Command, prefix string) {
	for range only.Once {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Flag", "Short flag", "Environment", "Description", "Value (* = default)"})
		table.SetBorder(true)

		cmd.PersistentFlags().SortFlags = false
		cmd.Flags().SortFlags = false
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if flag.Hidden {
				return
			}

			sh := ""
			if flag.Shorthand != "" {
				sh = "-" + flag.Shorthand
			}

			var value string
			foo := flag.Value.Type()
			switch foo {
			case "stringArray":
				va, _ := cmd.Flags().GetStringArray(flag.Name)
				value = va[0]
			default:
				value = flag.Value.String()
			}
			if value == flag.DefValue {
				value += " *"
			}
			table.Append([]string{
				"--" + flag.Name,
				sh,
				PrintFlagEnv(prefix, flag.Name),
				flag.Usage,
				value,
				// flag.Value.String(),
				// flag.DefValue,
			})
		})

		table.Render()
	}
}

func PrintExamples(cmd *cobra.Command, examples ...string) string {
	var ret string

	c := BuildCmd(cmd)
	for _, example := range examples {
		ret += fmt.Sprintf("\t%s %s\n", c, example)
	}

	return ret
}

func BuildCmd(cmd *cobra.Command) string {
	var ret string
	if cmd.HasParent() {
		ret += BuildCmd(cmd.Parent())
	}
	ret += cmd.Name() + " "
	return ret
}
