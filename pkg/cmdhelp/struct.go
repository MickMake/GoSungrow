package cmdhelp

import (
	"fmt"
	"strings"

	"github.com/anicoll/gosungrow/pkg/only"
	"github.com/spf13/cobra"
)

type Help struct {
	Error error

	Command          string
	HelpTemplate     string
	UsageTemplate    string
	FlagHelpTemplate string
	HelpSummary      string
	EnvPrefix        string
	ReadMe           string
	Examples         string

	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

func New() *Help {
	var ret *Help

	for range only.Once {
		ret = &Help{
			Error: nil,

			Command:          "DefaultBinaryName",
			HelpTemplate:     DefaultHelpTemplate,
			UsageTemplate:    DefaultUsageTemplate,
			FlagHelpTemplate: DefaultFlagHelpTemplate,
			HelpSummary:      DefaultHelpSummary,
			EnvPrefix:        "",

			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (h *Help) GetCmd() *cobra.Command {
	return h.SelfCmd
}

func (h *Help) SetCommand(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.Command = text
	}
}

func (h *Help) SetHelpTemplate(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.HelpTemplate = strings.ReplaceAll(text, "DefaultBinaryName", h.Command)
	}
}

func (h *Help) SetUsageTemplate(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.UsageTemplate = strings.ReplaceAll(text, "DefaultBinaryName", h.Command)
	}
}

func (h *Help) SetFlagHelpTemplate(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.FlagHelpTemplate = strings.ReplaceAll(text, "DefaultBinaryName", h.Command)
	}
}

func (h *Help) SetHelpSummary(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.HelpSummary = strings.ReplaceAll(text, "DefaultBinaryName", h.Command)
	}
}

func (h *Help) SetEnvPrefix(text string) {
	for range only.Once {
		if text == "" {
			break
		}

		h.EnvPrefix = strings.ToUpper(strings.ReplaceAll(text, "-", "_"))
	}
}

func (h *Help) ExtendedHelp() {
	fmt.Println(h.HelpSummary)
}

func (h *Help) SetReadMe(text string) {
	h.ReadMe = text
}

func (h *Help) SetExamples(text string) {
	h.Examples = text
}
