package google

import (
	"GoSungro/Only"
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"net/http"
	"os"
	"path/filepath"
)

const DefaultId = "1R4OmGyQJZJWBZVgEazNS3UZE5YVkXBa_XFulut3pohQ"

type Sheet struct {
	Id          string
	SheetName   string
	Range       string
	Credentials []byte
	TokenFile   string
	token       *oauth2.Token
	oAuthConfig *oauth2.Config

	Data SheetData
}
type SheetData [][]interface{}

func (s *Sheet) Set(cfg Sheet) {
	for range Only.Once {
		s.Id = cfg.Id
		s.SheetName = cfg.SheetName
		s.Range = cfg.Range

		s.Credentials = cfg.Credentials
		s.TokenFile = cfg.TokenFile
		s.token = cfg.token

		s.Verify()
	}
}

func (s *Sheet) Verify() bool {
	var ok bool
	for range Only.Once {
		if s.Id == "" {
			s.Id = DefaultId
		}
		if s.SheetName == "" {
			s.SheetName = ""
		}
		if s.Range == "" {
			s.Range = "A1"
		}
		if len(s.Credentials) == 0 {
			s.Credentials = []byte(DefaultGoogleCredentials)
		}

		// 	token  *oauth2.token
		//	oAuthConfig *oauth2.Config

		if s.oAuthConfig == nil {
			s.oAuthConfig = &oauth2.Config{}
		}
		if s.token == nil {
			s.token = &oauth2.Token{}
		}

		if len(s.TokenFile) == 0 {
			var err error
			s.TokenFile, err = os.UserHomeDir()
			if err != nil {
				s.TokenFile = ""
				break
			}
			s.TokenFile = filepath.Join(s.TokenFile, ".GoSungro", DefaultAuthTokenFile)
		}

		ok = true
	}

	return ok
}

func (s *Sheet) ReadSheet() (SheetData, error) {
	var err error
	var csv SheetData

	for range Only.Once {
		if !s.Verify() {
			break
		}

		// If modifying these scopes, delete your previously saved token.json.
		s.oAuthConfig, err = google.ConfigFromJSON(s.Credentials, sheets.SpreadsheetsReadonlyScope)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
			break
		}

		var client *http.Client
		client, err = s.getClient()

		ctx := context.Background()
		var srv *sheets.Service
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to retrieve Sheets client: %v", err))
			break
		}

		readRange := fmt.Sprintf("%s!%s", s.SheetName, s.Range)
		var resp *sheets.ValueRange
		resp, err = srv.Spreadsheets.Values.Get(s.Id, readRange).Do()
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to retrieve data from sheet: %v", err))
			break
		}

		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
			break
		}

		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%v\n", row)
		}
	}

	return csv, err
}

func (s *Sheet) WriteSheet() error {
	var err error

	for range Only.Once {
		if !s.Verify() {
			break
		}

		// If modifying these scopes, delete your previously saved token.json.
		s.oAuthConfig, err = google.ConfigFromJSON(s.Credentials, sheets.SpreadsheetsScope)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to parse client secret file to config: %v", err))
			break
		}

		var client *http.Client
		client, err = s.getClient()
		//client = getClient(s.oAuthConfig)

		ctx := context.Background()
		var srv *sheets.Service
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to retrieve Sheets client: %v", err))
			break
		}

		//var itt = srv.Spreadsheets.getSheetByName('_EmailList');
		//if (!itt) {
		//	ss.insertSheet('_EmailList');
		//}

		vr := sheets.ValueRange{
			MajorDimension:  "",
			Range:           s.SheetName + "!" + s.Range,
			Values:          s.Data,
			ServerResponse:  googleapi.ServerResponse{},
			ForceSendFields: nil,
			NullFields:      nil,
		}

		writeRange := fmt.Sprintf("%s!%s", s.SheetName, s.Range)
		fmt.Printf("Updating Google sheet '%s'.\n\tWorksheet: %s\n\tStarting from: %s\n",
			s.Id,
			s.SheetName,
			s.Range,
		)
		_, err = srv.Spreadsheets.Values.Update(s.Id, writeRange, &vr).ValueInputOption("USER_ENTERED").Do() // or "RAW"
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to retrieve data from sheet. %v", err))
			break
		}
	}

	return err
}
