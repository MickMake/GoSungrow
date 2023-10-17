package cmdhelp

const DefaultHelpTemplate = `
{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

const DefaultUsageTemplate = `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{ with (index .Annotations "group") }}{{ . }}	- {{ end }}{{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags: Use "{{.Root.CommandPath}} help flags" for more info.

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand }}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} help [command]" for more information about a command.{{end}}
`

// const DefaultVersionTemplate = `
// {{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
// `

//goland:noinspection GoUnusedConst
const DefaultFlagHelpTemplate = `{{if .HasAvailableInheritedFlags}}Flags available for all commands:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`

const DefaultHelpSummary = `
DefaultBinaryName - 
`
