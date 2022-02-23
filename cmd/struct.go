package cmd

import (
	"GoSungrow/Only"
	"GoSungrow/iSolarCloud"
	"GoSungrow/iSolarCloud/AppService/login"
	"GoSungrow/lsgo"
	"GoSungrow/mmGit"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

//goland:noinspection SpellCheckingInspection
const (
	DefaultBinaryName = "GoSungrow"
	EnvPrefix         = "SunGrow"
	defaultConfigFile = "config.json"
	defaultTokenFile  = "AuthToken.json"

	flagConfigFile = "config"
	flagDebug      = "debug"
	flagQuiet      = "quiet"

	flagApiUrl        = "host"
	flagApiTimeout    = "timeout"
	flagApiUsername   = "user"
	flagApiPassword   = "password"
	flagApiAppKey     = "appkey"
	flagApiLastLogin  = "token-expiry"
	flagApiOutputType = "out"

	flagGoogleSheet       = "google-sheet"
	flagGoogleSheetUpdate = "update"

	flagGitUsername = "git-username"
	flagGitPassword = "git-password"
	flagGitKeyFile  = "git-sshkey"
	flagGitToken    = "git-token"
	flagGitRepo     = "git-repo"
	flagGitRepoDir  = "git-dir"
	flagGitDiffCmd  = "diff-cmd"

	defaultHost      = "https://augateway.isolarcloud.com"
	defaultUsername  = "harry@potter.net"
	defaultPassword  = "hogwarts"
	defaultApiAppKey = "93D72E60331ABDCDC7B39ADC2D1F32B3"

	defaultTimeout = time.Duration(time.Second * 30)
)

var DefaultAreas = []string{"all"}

type CommandArgs struct {
	ConfigDir   string
	ConfigFile  string
	WriteConfig bool
	Quiet       bool
	Debug       bool
	OutputType  string
	OutputFile  string

	// iSolarCloud api
	ApiTimeout    time.Duration
	ApiUrl        string
	ApiUsername   string
	ApiPassword   string
	ApiAppKey     string
	ApiLastLogin  string
	ApiToken      string
	ApiTokenFile  string
	ApiOutputType string

	// Google sheets
	GoogleSheet       string
	GoogleSheetUpdate bool

	// GitHub api
	GitRepo     string
	GitRepoDir  string
	GitUsername string
	GitPassword string
	GitKeyFile  string
	GitToken    string
	GitDiffCmd  string

	Args []string

	Valid bool
	Error error
}

func (ca *CommandArgs) IsValid() error {
	for range Only.Once {
		if !ca.Valid {
			ca.Error = errors.New("args are not valid")
			break
		}
	}

	return ca.Error
}

func (ca *CommandArgs) ProcessArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		ca.Args = args

		SunGrow = iSolarCloud.NewSunGro(ca.ApiUrl)
		if SunGrow.Error != nil {
			break
		}

		Cmd.Error = SunGrow.Init()
		if Cmd.Error != nil {
			break
		}

		switch Cmd.ApiOutputType {
		case "json":
			SunGrow.OutputType.SetJson()
		case "raw":
			SunGrow.OutputType.SetRaw()
		case "file":
			SunGrow.OutputType.SetFile()
		default:
			SunGrow.OutputType.SetJson()
		}

		if ca.ApiAppKey == "" {
			ca.ApiAppKey = defaultApiAppKey
		}

		// if Cmd.GoogleSheetUpdate {
		// 	SunGrow.OutputType = iSolarCloud.TypeGoogle
		// }

		// Git.Error = Cmd.GitSet()
		// if Cmd.Error != nil {
		//	break
		// }

		ca.Valid = true
	}

	return ca.Error
}

func (ca *CommandArgs) SunGrowArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		Cmd.Error = Cmd.ProcessArgs(cmd, args)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = SunGrow.Login(login.SunGrowAuth{
			AppKey:       ca.ApiAppKey,
			UserAccount:  ca.ApiUsername,
			UserPassword: ca.ApiPassword,
			TokenFile:    ca.ApiTokenFile,
			Force:        false,
		})
		if Cmd.Error != nil {
			break
		}

		if Cmd.Debug {
			SunGrow.Auth.Print()
		}

		if SunGrow.HasTokenChanged() {
			ca.ApiLastLogin = SunGrow.GetLastLogin()
			ca.ApiToken = SunGrow.GetToken()
			ca.Error = writeConfig()
		}

		// if Cmd.GoogleSheetUpdate {
		// 	SunGrow.OutputType = iSolarCloud.TypeGoogle
		// }

		// Git.Error = Cmd.GitSet()
		// if Cmd.Error != nil {
		//	break
		// }

		ca.Valid = true
	}

	return ca.Error
}

func (ca *CommandArgs) GitSet() error {
	for range Only.Once {
		if Git != nil {
			break
		}

		Git = mmGit.New()
		if Git.Error != nil {
			ca.Error = Git.Error
			break
		}

		// Cmd.Error = Git.SetAuth(ca.GitUsername, ca.GitPassword)
		// if Cmd.Error != nil {
		//	break
		// }

		Cmd.Error = Git.SetKeyFile(ca.GitKeyFile)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.SetToken(ca.GitToken)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.SetRepo(ca.GitRepo)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.SetDir(ca.GitRepoDir)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.SetDiffCmd(ca.GitDiffCmd)
		if Cmd.Error != nil {
			break
		}
	}

	return Cmd.Error
}

func (ca *CommandArgs) GitLs(options ...string) error {
	for range Only.Once {
		os.Args = []string{"GitLs"}
		os.Args = append(os.Args, options...)
		ca.Error = os.Chdir(ca.GitRepoDir)
		if ca.Error != nil {
			break
		}

		// ls-go is a standalone GoLang executable,
		// but I've modified it to be a package and so directly callable.
		ca.Error = lsgo.LsGo()
		if ca.Error != nil {
			break
		}
	}

	return Cmd.Error
}

func (ca *CommandArgs) GitSync(msg string, entities ...string) error {
	for range Only.Once {
		Cmd.Error = Git.Pull()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = ca.GitSave(entities...)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Add(".")
		if Cmd.Error != nil {
			break
		}

		if msg == "" {
			msg = fmt.Sprintf("Updated %d files.", len(entities))
		}
		Cmd.Error = Git.Commit(msg)
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = Git.Push()
		if Cmd.Error != nil {
			break
		}
	}

	return Cmd.Error
}

func (ca *CommandArgs) GitSave(entities ...string) error {
	for range Only.Once {
		if len(entities) == 0 {
			entities = DefaultAreas
		}
		fmt.Printf("Saving %d entities from the SunGrow to Git...\n", len(entities))

		// SunGrow.OutputType = iSolarCloud.StringTypeJson
		// SunGrow.OutputType = iSolarCloud.TypeJson

		for _, entity := range entities {
			// Remove plurals.
			entity = strings.TrimSuffix(entity, "s")
			// SunGrow.OutputString = ""

			switch entity {
			case "domain":
				SunGrow.Error = SunGrow.Init()
			}
			if SunGrow.Error != nil {
				break
			}

			jf := AddJsonSuffix(entity)
			Cmd.Error = Git.SaveFile(jf, []byte(ca.OutputFile))
			if Cmd.Error != nil {
				break
			}
		}

		fmt.Printf("Saved %d files.", len(entities))
	}

	return Cmd.Error
}

func (ca *CommandArgs) GoogleUpdate(entities ...string) error {

	for range Only.Once {
		// SunGrow.OutputType = iSolarCloud.TypeGoogle

		if len(entities) == 0 {
			entities = DefaultAreas
		}
		fmt.Printf("Saving %d entities from the SunGrow to Google Docs...\n", len(entities))

		for _, entity := range entities {
			switch entity {
			case "domain":
				ca.Error = SunGrow.Init()
				if ca.Error != nil {
					break
				}
			}

			// sheet := google.Sheet {
			// 	Id:          "",
			// 	Credentials: nil,
			// 	SheetName:   entity,
			// 	Range:       "",
			// 	Data:        ca.OutputFile,
			// }
			// sheet.Set(sheet)
			// ca.Error = sheet.WriteSheet()
			if ca.Error != nil {
				break
			}
		}
	}

	return ca.Error
}

// func (ca *CommandArgs) UpdateGoogleSheet(name string, data [][]interface{}) error {
//
//	for range Only.Once {
//		sheet := google.Sheet{
//			Id:          "",
//			Credentials: nil,
//			SheetName:   name,
//			Range:       "",
//			Data:        data,
//		}
//		sheet.Set(sheet)
//		ca.Error = sheet.WriteSheet()
//		if ca.Error != nil {
//			break
//		}
//	}
//
//	return p.Error
// }
