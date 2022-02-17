package google

import (
	"GoSungro/Only"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

// Retrieve a token, saves the token, then returns the generated client.
func (s *Sheet) getClient() (*http.Client, error) {
	var ret *http.Client
	var err error

	for range Only.Once {
		// The file token.json stores the user's access and refresh tokens, and is
		// created automatically when the authorization flow completes for the first
		// time.

		var err error
		err = s.tokenFromFile()
		if err != nil {
			err = s.getTokenFromWeb()
			if err != nil {
				break
			}
			err = s.saveToken()
			if err != nil {
				break
			}
		}

		ret = s.oAuthConfig.Client(context.Background(), s.token)
	}

	return ret, err
}

// RequestCommon a token from the web, then returns the retrieved token.
func (s *Sheet) getTokenFromWeb() error {
	var err error

	for range Only.Once {
		authURL := s.oAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Printf("Allow access to this application by going to this URL:\n%v\n", authURL)
		fmt.Printf("\nThen copy-paste the authorization code here: ")

		var authCode string
		_, err = fmt.Scan(&authCode)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to read authorization code: %v", err))
			break
		}

		s.token, err = s.oAuthConfig.Exchange(context.TODO(), authCode)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to retrieve token from web: %v", err))
			break
		}
	}

	return err
}

// Retrieves a token from a local file.
func (s *Sheet) tokenFromFile() error {
	var err error

	for range Only.Once {
		if s.TokenFile == "" {
			err = errors.New("Empty token filename")
			break
		}

		var f *os.File
		f, err = os.Open(s.TokenFile)
		if err != nil {
			break
		}

		defer f.Close()
		//ret = &oauth2.token{}
		err = json.NewDecoder(f).Decode(s.token)
	}

	return err
}

// Saves a token to a file path.
func (s *Sheet) saveToken() error {
	var err error

	for range Only.Once {
		fmt.Printf("Saving token file to: %s\n", s.TokenFile)
		var f *os.File
		f, err = os.OpenFile(s.TokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to cache oauth token: %v", err))
			break
		}

		defer f.Close()
		err = json.NewEncoder(f).Encode(s.token)
	}

	return err
}
