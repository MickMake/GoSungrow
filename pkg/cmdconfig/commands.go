package cmdconfig

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/anicoll/gosungrow/pkg/cmdhelp"
	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const Group = "Config"

func (c *Config) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "config",
			Short:                 "Create, update or show config file.",
			Long:                  "Create, update or show config file.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdConfig,
			Args:                  cobra.RangeArgs(0, 1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdhelp.PrintExamples(c.SelfCmd, "read", "write", "write --timeout=60s")
		c.SelfCmd.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		cmdConfigWrite := &cobra.Command{
			Use:                   "write",
			Short:                 "Update config file from CLI args.",
			Long:                  "Update config file from CLI args.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdWrite,
			Args:                  cobra.RangeArgs(0, 1),
		}
		c.SelfCmd.AddCommand(cmdConfigWrite)
		cmdConfigWrite.Example = cmdhelp.PrintExamples(cmdConfigWrite, "", "write --timeout=60s", "--debug=true")
		cmdConfigWrite.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		cmdConfigRead := &cobra.Command{
			Use:                   "read",
			Short:                 "Read config file.",
			Long:                  "Read config file.",
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               c.InitArgs,
			RunE:                  c.CmdRead,
			Args:                  cobra.RangeArgs(0, 1),
		}
		c.SelfCmd.AddCommand(cmdConfigRead)
		cmdConfigRead.Example = cmdhelp.PrintExamples(cmdConfigRead, "")
		cmdConfigRead.Annotations = map[string]string{"group": Group}
	}

	return c.SelfCmd
}

func (c *Config) InitArgs(_ *cobra.Command, _ []string) error {
	var err error
	for range only.Once {
		//
	}
	return err
}

func (c *Config) CmdConfig(cmd *cobra.Command, args []string) error {
	for range only.Once {
		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())
		if len(args) == 0 {
			_ = cmd.Help()
		}
	}

	return c.Error
}

func (c *Config) CmdWrite(_ *cobra.Command, args []string) error {
	for range only.Once {
		if len(args) == 1 {
			c.File = args[0]
			c.SetFile(c.File)
			c.Error = c.Open()
			if c.Error != nil {
				break
			}
		}

		c.Error = c.UpdateFlags()
		if c.Error != nil {
			break
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())
		fmt.Println("New config:")
		fmt.Print(c.PrintConfig())

		c.Error = c.Write()
		if c.Error != nil {
			break
		}
	}

	return c.Error
}

func (c *Config) PrintConfig() string {
	var ret string
	for range only.Once {
		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"Flag", "Short flag", "Environment", "Description", "Value (* = default)"})
		table.SetBorder(true)

		c.cmd.PersistentFlags().SortFlags = false
		c.cmd.Flags().SortFlags = false
		c.cmd.Flags().VisitAll(func(flag *pflag.Flag) {
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
				va, _ := c.cmd.Flags().GetStringArray(flag.Name)
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
				PrintFlagEnv(c.EnvPrefix, flag.Name),
				flag.Usage,
				value,
				// flag.Value.String(),
				// flag.DefValue,
			})
		})

		table.Render()
		ret = buf.String()
	}
	return ret
}

func PrintFlagEnv(prefix string, flag string) string {
	fenv := strings.ReplaceAll(flag, "-", "_")
	fenv = strings.ToUpper(fenv)

	// ret := fmt.Sprintf("--%s\t%s_%s\n", flag, EnvPrefix, fenv)
	// ret := fmt.Sprintf("%s_%s", defaults.EnvPrefix, fenv)
	// ret := fmt.Sprintf("%s_%s", cmdVersion.GetEnvPrefix(), fenv)
	ret := fmt.Sprintf("%s_%s", prefix, fenv)
	return ret
}

func (c *Config) CmdRead(_ *cobra.Command, args []string) error {
	for range only.Once {
		if len(args) == 1 {
			c.File = args[0]
			c.SetFile(c.File)

			c.Error = c.Open()
			if c.Error != nil {
				break
			}
		}

		c.Error = c.UpdateFlags()
		if c.Error != nil {
			break
		}

		_, _ = fmt.Fprintf(os.Stderr, "Using config file '%s'\n", c.viper.ConfigFileUsed())

		cmdhelp.PrintConfig(c.cmd, c.EnvPrefix) // rootCmd
	}

	return c.Error
}
