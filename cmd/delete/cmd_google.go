package cmd


// func AttachCmdGoogle(cmd *cobra.Command) *cobra.Command {
// 	// ******************************************************************************** //
// 	var cmdGoogle = &cobra.Command{
// 		Use: "google",
// 		Aliases:               []string{""},
// 		Short:                 fmt.Sprintf("Update and view Google sheets."),
// 		Long:                  fmt.Sprintf("Update and view Google sheets."),
// 		DisableFlagParsing:    false,
// 		DisableFlagsInUseLine: false,
// 		PreRunE:               Cmd.ProcessArgs,
// 		Run:                   cmdGoogleFunc,
// 		Args:                  cobra.RangeArgs(0, 1),
// 	}
// 	cmd.AddCommand(cmdGoogle)
// 	cmdGoogle.Example = PrintExamples(cmdGoogle, "update all", "update devices")
//
//
// 	// ******************************************************************************** //
// 	var cmdGoogleSync = &cobra.Command{
// 		Use: "update",
// 		Aliases:               []string{""},
// 		Short:                 fmt.Sprintf("Update Google sheets."),
// 		Long:                  fmt.Sprintf("Update Google sheets."),
// 		DisableFlagParsing:    false,
// 		DisableFlagsInUseLine: false,
// 		PreRunE:               Cmd.ProcessArgs,
// 		Run:                   cmdGoogleSyncFunc,
// 		Args:                  cobra.RangeArgs(0, 1),
// 	}
// 	cmdGoogle.AddCommand(cmdGoogleSync)
// 	cmdGoogleSync.Example = PrintExamples(cmdGoogleSync, "all", "device", "unit")
//
// 	return cmdGoogle
// }
//
//
// func cmdGoogleFunc(_ *cobra.Command, _ []string) {
// 	for range Only.Once {
// 		fmt.Println("Not yet implemented.")
// 		return
//
// 		// switch {
// 		// 	case len(args) == 0:
// 		// 		Cmd.Error = cmd.Help()
// 		//
// 		// 	case args[0] == "all":
// 		// 		Cmd.Error = Cmd.GoogleUpdate()
// 		//
// 		// 	default:
// 		// 		fmt.Println("Unknown sub-command.")
// 		// 		_ = cmd.Help()
// 		// }
// 	}
// }
//
// func cmdGoogleSyncFunc(_ *cobra.Command, _ []string) {
// 	for range Only.Once {
// 		fmt.Println("Not yet implemented.")
// 		return
//
// 		// switch {
// 		// 	case len(args) == 0:
// 		// 		Cmd.Error = cmd.Help()
// 		//
// 		// 	case args[0] == "all":
// 		// 		Cmd.Error = Cmd.GoogleUpdate()
// 		//
// 		// 	default:
// 		// 		fmt.Println("Unknown sub-command.")
// 		// 		_ = cmd.Help()
// 		// }
// 	}
// }
