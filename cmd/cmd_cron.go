package cmd

import (
	"GoSungro/Only"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)


type CronStruct struct {
	Scheduler *gocron.Scheduler
	Job *gocron.Job
}
var Cron CronStruct

// ******************************************************************************** //
var cmdCron = &cobra.Command{
	Use:                   "cron",
	//Aliases:               []string{},
	Short:                 fmt.Sprintf("Run a command via schedule."),
	Long:                  fmt.Sprintf("Run a command via schedule."),
	Example:               PrintExamples("cron","run 00 18 . . . sync default", "run 42 02 04 . . list all"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCronFunc,
	Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdCronFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		if len(args) == 0 {
			Cmd.Error = cmd.Help()
			break
		}
	}
}


// ******************************************************************************** //
var cmdCronRun = &cobra.Command{
	Use:                   "run <minute> <hour> <month day> <month> <week day>  <command ...>",
	//Aliases:               []string{},
	Short:                 fmt.Sprintf("Run scheduled a command."),
	Long:                  fmt.Sprintf("Run scheduled a command."),
	Example:               PrintExamples("cron run","00 18 . . . sync default", "42 02 04 . . list all", "./1 . . . . list all"),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCronRunFunc,
	Args:                  cobra.MinimumNArgs(6),
}
//goland:noinspection GoUnusedParameter
func cmdCronRunFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		// */1 * * * * /dir/command args args
		cronString := strings.Join(args[0:5], " ")
		cronString = strings.ReplaceAll(cronString, ".", "*")
		ResetArgs(args[5:]...)

		Cron.Scheduler = gocron.NewScheduler(time.UTC)
		Cron.Scheduler = Cron.Scheduler.Cron(cronString)
		Cron.Scheduler = Cron.Scheduler.SingletonMode()

		Cron.Job, Cmd.Error = Cron.Scheduler.Do(ReExecute)
		if Cmd.Error != nil {
			break
		}

		LogPrint("Created job schedule using '%s'\n", cronString)
		LogPrint("Job command '%s'\n", strings.Join(os.Args, " "))

		Cron.Scheduler.StartBlocking()
		if Cmd.Error != nil {
			break
		}
	}
}


// ******************************************************************************** //
var cmdCronAdd = &cobra.Command{
	Use:                   "add",
	//Aliases:               []string{},
	Short:                 fmt.Sprintf("Add scheduled a command."),
	Long:                  fmt.Sprintf("Add scheduled a command."),
	Example:               PrintExamples("cron", "add "),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCronAddFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdCronAddFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")

		//var msg string
		//switch {
		//	case args[0] == "":
		//		fallthrough
		//	case args[0] == "default":
		//		//u, _ := user.Current()
		//		//msg = fmt.Sprintf("Regular sync by %s", u.ApiUsername)
		//	default:
		//		msg = args[0]
		//}
		//
		//args = args[1:]
		//
		////Cmd.Error = Cmd.CronAdd(msg, args...)
		//if Cmd.Error != nil {
		//	break
		//}
	}
}


// ******************************************************************************** //
var cmdCronRemove = &cobra.Command{
	Use:                   "del",
	Aliases:               []string{"remove"},
	Short:                 fmt.Sprintf("Remove a scheduled command."),
	Long:                  fmt.Sprintf("Remove a scheduled command."),
	Example:               PrintExamples("cron", "del "),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCronRemoveFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdCronRemoveFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")

		//var msg string
		//switch {
		//	case args[0] == "":
		//		fallthrough
		//	case args[0] == "default":
		//		//u, _ := user.Current()
		//		//msg = fmt.Sprintf("Regular sync by %s", u.ApiUsername)
		//	default:
		//		msg = args[0]
		//}
		//
		//args = args[1:]
		//
		////Cmd.Error = Cmd.CronAdd(msg, args...)
		//if Cmd.Error != nil {
		//	break
		//}
	}
}


// ******************************************************************************** //
var cmdCronList = &cobra.Command{
	Use:                   "list",
	//Aliases:               []string{},
	Short:                 fmt.Sprintf("List scheduled commands."),
	Long:                  fmt.Sprintf("List scheduled commands."),
	Example:               PrintExamples("cron", "list "),
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: false,
	PreRunE:               Cmd.ProcessArgs,
	Run:                   cmdCronListFunc,
	//Args:                  cobra.MinimumNArgs(1),
}
//goland:noinspection GoUnusedParameter
func cmdCronListFunc(cmd *cobra.Command, args []string) {
	for range Only.Once {
		fmt.Println("Not yet implemented.")

		//var msg string
		//	switch {
		//		case args[0] == "":
		//			fallthrough
		//		case args[0] == "default":
		//			//u, _ := user.Current()
		//			//msg = fmt.Sprintf("Regular sync by %s", u.ApiUsername)
		//		default:
		//			msg = args[0]
		//}
		//
		//args = args[1:]
		//
		//Cmd.Error = Cmd.CronList(msg, args...)
		//if Cmd.Error != nil {
		//	break
		//}
	}
}


func timeStamp() string {
	return time.Now().UTC().Format(time.UnixDate) + " : "
}
func LogPrint(format string, args ...interface{}) {
	format = timeStamp() + format
	fmt.Printf(format, args...)
}

func ReExecute() error {
	for range Only.Once {
		LogPrint("Running scheduled command '%s'\n", strings.Join(os.Args, " "))
		//LogPrint("Last run '%s'\n", Cron.Job.LastRun().Format(time.UnixDate))
		LogPrint("Next run '%s'\n", Cron.Job.ScheduledTime().Format(time.UnixDate))
		LogPrint("Run count '%d'\n", Cron.Job.RunCount())

		Cmd.Error = rootCmd.Execute()
		if Cmd.Error != nil {
			LogPrint("ERROR: %s\n", Cmd.Error)
			break
		}
	}

	return Cmd.Error
}

func ResetArgs(args ...string) {
	for range Only.Once {
		//fmt.Printf("oldArgs: %v\n", os.Args)
		newArgs := []string{os.Args[0]}
		newArgs = append(newArgs, args...)
		os.Args = newArgs
		//fmt.Printf("newArgs: %v\n", os.Args)
	}
}
