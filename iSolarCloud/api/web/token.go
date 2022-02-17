package web


// const (
// 	DateTimeFormat       = "2006-01-02T15:04:05"
// 	DefaultAuthTokenFile = "SunGroAuthToken.json"
// )
//
// var (
// 	TokenRequestUrl = "/v1/userService/login"
// )
//
// type SunGroAuth struct {
// 	Expiry string
//
// 	AppKey      string
// 	UserAccount    string
// 	UserPassword    string
// }
//
// type Auth struct {
// 	// Url *url.URL
// 	// RequestCommon  login.RequestCommon
// 	// ResponseCommon login.ResponseCommon
//
// 	api.EndPoint
//
// 	File   string
// 	Expiry time.Time
// 	newToken    bool
// 	retry       int
//
// 	Error error
// }
//
//
// func (t *Auth) Auth(auth *SunGroAuth) error {
// 	for range Only.Once {
// 		_ = t.readTokenFile()
//
// 		t.Error = t.Verify(auth)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		t.Error = t.RetrieveToken()
// 		if t.Error != nil {
// 			break
// 		}
// 	}
//
// 	return t.Error
// }
//
// func (t *Auth) Verify(auth *SunGroAuth) error {
// 	for range Only.Once {
// 		if auth == nil {
// 			// If nil, then assume we haven't set anything.
// 			break
// 		}
//
// 		if auth.AppKey == "" {
// 			t.Error = errors.New("API AppKey")
// 			break
// 		}
// 		if auth.UserAccount == "" {
// 			t.Error = errors.New("empty Client ApiUsername")
// 			break
// 		}
// 		if auth.UserPassword == "" {
// 			t.Error = errors.New("empty Client ApiPassword")
// 			break
// 		}
//
// 		if t.ResponseCommon.ResultData.Auth == "" {
// 			t.newToken = true
// 		}
//
// 		if auth.Expiry == "" {
// 			auth.Expiry = time.Now().Format(DateTimeFormat)
// 		}
// 		t.Expiry, t.Error = time.Parse(DateTimeFormat, auth.Expiry)
// 		if t.Error != nil {
// 			t.newToken = true
// 		}
//
// 		t.RequestCommon = login.RequestCommon {
// 			Appkey:   auth.AppKey,
// 			SysCode:  "900",
// 			UserAccount: auth.UserAccount,
// 			UserPassword: auth.UserPassword,
// 		}
//
// 		t.HasTokenExpired()
// 	}
//
// 	return t.Error
// }
//
// func (t *Auth) RetrieveToken() error {
// 	for range Only.Once {
// 		t.HasTokenExpired()
// 		if !t.newToken {
// 			break
// 		}
//
// 		u := fmt.Sprintf("%s%s",
// 			t.Url.String(),
// 			TokenRequestUrl,
// 		)
// 		//p, _ := json.Marshal(map[string]string {
// 		//	"user_account": t.RequestCommon.UserAccount,
// 		//	"user_password": t.RequestCommon.UserPassword,
// 		//	"appkey": t.RequestCommon.AppKey,
// 		//	"sys_code": "900",
// 		//})
// 		p, _ := json.Marshal(t.RequestCommon)
//
// 		var httpResponse *http.ResponseCommon
// 		httpResponse, t.Error = http.Post(u, "application/json", bytes.NewBuffer(p))
// 		if t.Error != nil {
// 			break
// 		}
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer httpResponse.Body.Close()
// 		if httpResponse.StatusCode != 200 {
// 			t.Error = errors.New(fmt.Sprintf("Status Code is %d", httpResponse.StatusCode))
// 			break
// 		}
//
// 		var body []byte
// 		body, t.Error = ioutil.ReadAll(httpResponse.Body)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		t.Error = json.Unmarshal(body, &t.ResponseCommon)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		t.Expiry = time.Now()
//
// 		t.Error = t.saveToken()
// 		if t.Error != nil {
// 			break
// 		}
// 	}
//
// 	return t.Error
// }
//
// func (t *Auth) HasTokenExpired() bool {
// 	for range Only.Once {
// 		if t.Expiry.Before(time.Now()) {
// 			t.newToken = true
// 			break
// 		}
//
// 		if t.ResponseCommon.ResultData.Auth == "" {
// 			t.newToken = true
// 			break
// 		}
// 	}
//
// 	return t.newToken
// }
//
// func (t *Auth) HasTokenChanged() bool {
// 	ok := t.newToken
// 	if t.newToken {
// 		t.newToken = false
// 	}
// 	return ok
// }
//
// // func (t *Auth) GetAuthHeader() string {
// // 	var ret string
// //
// // 	for range Only.Once {
// // 		//if t.ResponseCommon.TokenType == "" {
// // 		//	break
// // 		//}
// // 		//
// // 		//if t.ResponseCommon.AccessToken == "" {
// // 		//	break
// // 		//}
// // 		//
// // 		//ret = t.ResponseCommon.TokenType + " " + t.ResponseCommon.AccessToken
// // 	}
// //
// // 	return ret
// // }
//
// func (t *Auth) GetToken() string {
// 	//return fmt.Sprintf("%s %s", t.ResponseCommon.TokenType, t.ResponseCommon.AccessToken)
// 	return ""
// }
//
// func (t *Auth) GetTokenExpiry() time.Time {
// 	return t.Expiry
// }
//
// // Retrieves a token from a local file.
// func (t *Auth) readTokenFile() error {
// 	for range Only.Once {
// 		if t.File == "" {
// 			t.File, t.Error = os.UserHomeDir()
// 			if t.Error != nil {
// 				t.File = ""
// 				break
// 			}
// 			t.File = filepath.Join(t.File, ".GoSungro", DefaultAuthTokenFile)
// 		}
//
// 		var f *os.File
// 		f, t.Error = os.Open(t.File)
// 		if t.Error != nil {
// 			break
// 		}
//
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer f.Close()
// 		//ret = &oauth2.token{}
// 		t.Error = json.NewDecoder(f).Decode(&t.ResponseCommon)
// 	}
//
// 	return t.Error
// }
//
// // Saves a token to a file path.
// func (t *Auth) saveToken() error {
// 	for range Only.Once {
// 		if t.File == "" {
// 			t.File, t.Error = os.UserHomeDir()
// 			if t.Error != nil {
// 				t.File = ""
// 				break
// 			}
// 			t.File = filepath.Join(t.File, ".GoSungro", DefaultAuthTokenFile)
// 		}
//
// 		fmt.Printf("Saving token file to: %s\n", t.File)
// 		var f *os.File
// 		f, t.Error = os.OpenFile(t.File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 		if t.Error != nil {
// 			t.Error = errors.New(fmt.Sprintf("Unable to cache SUNGRO oauth token: %v", t.Error))
// 			break
// 		}
//
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer f.Close()
// 		t.Error = json.NewEncoder(f).Encode(t.ResponseCommon.ResultData)
// 	}
//
// 	return t.Error
// }
