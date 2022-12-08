package cmd

import (
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	"github.com/MickMake/GoUnify/cmdHelp"
	"github.com/spf13/cobra"
)


type CmdShow CmdDefault

func NewCmdShow() *CmdShow {
	return &CmdShow {
			Error:   nil,
			cmd:     nil,
			SelfCmd: nil,
	}
}


func (c *CmdShow) AttachCommand(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		c.cmd = cmd

		// ******************************************************************************** //
		c.SelfCmd = &cobra.Command{
			Use:                   "show",
			Aliases:               []string{},
			Annotations:           map[string]string{"group": "Show"},
			Short:                 fmt.Sprintf("Helpful Sungrow functions."),
			Long:                  fmt.Sprintf("Helpful Sungrow functions."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               func(cmd *cobra.Command, args []string) error {
				cmds.Error = cmds.ProcessArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				cmds.Error = cmds.SunGrowArgs(cmd, args)
				if cmds.Error != nil {
					return cmds.Error
				}
				return nil
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return cmd.Help()
			},
			Args: cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(c.SelfCmd)
		c.SelfCmd.Example = cmdHelp.PrintExamples(c.SelfCmd, "")

		c.AttachPsId(c.SelfCmd)
		c.AttachDevice(c.SelfCmd)
		c.AttachPoint(c.SelfCmd)
		c.AttachTemplate(c.SelfCmd)
	}
	return c.SelfCmd
}
