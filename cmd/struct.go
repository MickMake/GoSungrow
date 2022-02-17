package cmd

import (
	"GoSungro/Only"
	"GoSungro/iSolarCloud/sungro"
	"GoSungro/iSolarCloud/sungro/AppService/login"
	"GoSungro/lsgo"
	"GoSungro/mmGit"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)


//goland:noinspection SpellCheckingInspection
const (
	DefaultBinaryName = "GoSungro"
	EnvPrefix         = "SunGro"
	defaultConfigFile = "config.json"
	defaultTokenFile  = "AuthToken.json"

	flagConfigFile     = "config"
	flagDebug          = "debug"
	flagQuiet          = "quiet"

	flagApiUrl         = "host"
	flagApiTimeout     = "timeout"
	flagApiUsername    = "user"
	flagApiPassword    = "password"
	flagApiAppKey      = "appkey"
	flagApiTokenExpiry = "token-expiry"

	flagGoogleSheet       = "google-sheet"
	flagGoogleSheetUpdate = "update"

	flagGitUsername = "git-username"
	flagGitPassword = "git-password"
	flagGitKeyFile  = "git-sshkey"
	flagGitToken    = "git-token"
	flagGitRepo     = "git-repo"
	flagGitRepoDir  = "git-dir"
	flagGitDiffCmd  = "diff-cmd"

	defaultHost     = "https://augateway.isolarcloud.com"
	defaultUsername = "harry@potter.net"
	defaultPassword = "hogwarts"

	defaultTimeout = time.Duration(time.Second * 30)
)

var DefaultAreas = []string{"all"}


type CommandArgs struct {
	ConfigDir   string
	ConfigFile  string
	WriteConfig bool
	Quiet   bool
	Debug   bool
	OutputType string
	OutputFile string

	// iSolarCloud api
	ApiTimeout     time.Duration
	ApiUrl         string
	ApiUsername    string
	ApiPassword    string
	ApiAppKey      string
	ApiTokenExpiry string
	ApiToken       string
	ApiTokenFile   string

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

//goland:noinspection GoUnusedParameter
func (ca *CommandArgs) ProcessArgs(cmd *cobra.Command, args []string) error {
	for range Only.Once {
		ca.Args = args

		SunGro = sungro.NewSunGro(ca.ApiUrl)
		if SunGro.Error != nil {
			break
		}

		Cmd.Error = SunGro.Init()
		if Cmd.Error != nil {
			break
		}

		Cmd.Error = SunGro.Login(login.SunGroAuth {
			AppKey:       ca.ApiAppKey,
			UserAccount:  ca.ApiUsername,
			UserPassword: ca.ApiPassword,
			TokenFile:    ca.ApiTokenFile,
		})
		if Cmd.Error != nil {
			break
		}

		if SunGro.HasTokenChanged() {
			ca.ApiTokenExpiry = SunGro.GetTokenExpiry()
			ca.ApiToken = SunGro.GetToken()
			ca.Error = writeConfig()
		}

		// if Cmd.GoogleSheetUpdate {
		// 	SunGro.OutputType = iSolarCloud.TypeGoogle
		// }

		//Git.Error = Cmd.GitSet()
		//if Cmd.Error != nil {
		//	break
		//}

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

		//Cmd.Error = Git.SetAuth(ca.GitUsername, ca.GitPassword)
		//if Cmd.Error != nil {
		//	break
		//}

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
		fmt.Printf("Saving %d entities from the SunGro to Git...\n", len(entities))

		//SunGro.OutputType = iSolarCloud.StringTypeJson
		// SunGro.OutputType = iSolarCloud.TypeJson

		for _, entity := range entities {
			// Remove plurals.
			entity = strings.TrimSuffix(entity, "s")
			// SunGro.OutputString = ""

			switch entity {
				case "domain":
					SunGro.Error = SunGro.Init()
			}
			if SunGro.Error != nil {
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
		// SunGro.OutputType = iSolarCloud.TypeGoogle

		if len(entities) == 0 {
			entities = DefaultAreas
		}
		fmt.Printf("Saving %d entities from the SunGro to Google Docs...\n", len(entities))

		for _, entity := range entities {
			switch entity {
				case "domain":
					ca.Error = SunGro.Init()
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

//func (ca *CommandArgs) UpdateGoogleSheet(name string, data [][]interface{}) error {
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
//}
