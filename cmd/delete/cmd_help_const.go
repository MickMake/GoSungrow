package cmd


// const DefaultHelpTemplate = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}
//
// {{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
//
// const DefaultUsageTemplate = `Usage:{{if .Runnable}}
//   {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
//   {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}
//
// Aliases:
//   {{.NameAndAliases}}{{end}}{{if .HasExample}}
//
// Examples:
// {{.Example}}{{end}}{{if .HasAvailableSubCommands}}
//
// Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
//   {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
//
// Flags: Use "{{.Root.CommandPath}} help-all" for more info.
//
// Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand }}
//   {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
//
// Use "{{.CommandPath}} help [command]" for more information about a command.{{end}}
// `
//
// const DefaultVersionTemplate = `
// {{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
// `
//
// //goland:noinspection GoUnusedConst
// const DefaultFlagHelpTemplate = `{{if .HasAvailableInheritedFlags}}Flags available for all commands:
// {{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
// `
//
// const ExtendedHelpText = `
// DefaultBinaryName - SunGrow to Gitlab syncing tool.
//
// This tool does several things:
// 1. Pull a Gitlab repo that holds SunGrow data.
// 2. Fetch all data available from the SunGrow.
// 3. Save this data as a JSON file.
// 4. Commit changes to the Gitlab repo.
// 5. Push changes to remote.
//
// It is intended to provide full revision history for any changes made to the SunGrow.
//
// Use case example:
// # Record changes made to user data on SunGrow. (Will clone if not existing.)
// 	% DefaultBinaryName sync 'Updated users' users
//
// # Record changes made to all SunGrow manually.
// 	% DefaultBinaryName sync 'Updated all zones'
//
// # Record changes made to the SunGrow with default commit message.
// 	% DefaultBinaryName sync default
//
// # Record changes made to the SunGrow via every 30 minutes.
// 	% DefaultBinaryName cron run ./30 . . . . sync default
//
// # Show changes made to a zone JSON file.
// 	% DefaultBinaryName git diff devices.json
//
// # Other available Gitlab commands.
// 	Clone repo.
// 	% DefaultBinaryName git clone
//
// 	Pull repo.
// 	% DefaultBinaryName git pull
//
// 	Add files to repo.
// 	% DefaultBinaryName git add .
//
// 	Push repo.
// 	% DefaultBinaryName git push
//
// 	Commit changes to repo.
// 	% DefaultBinaryName git commit 'this is a commit message'
//
// # Config file.
// 	Show current config.
// 	% DefaultBinaryName config read
//
// 	Change diff command used in compares.
// 	% DefaultBinaryName --diff-cmd='sdiff' config write
//
// 	Change Git repo directory.
// 	% DefaultBinaryName --git-dir=/some/other/directory config write
//
// 	Change Git repo url.
// 	% DefaultBinaryName --git-url=https://github.com/MickMake/iSolarData config write
//
// 	Change SunGrow API token.
// 	% DefaultBinaryName --cf-token='this is a token string' config write
//
// `