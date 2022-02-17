package apiReflect


// Old Get method

// func (w *ApiRoot) Get(endpoint api.EndPoint) (api.EndPoint, error) {
// 	// func (w *ApiRoot) Get(u *url.URL, request interface{}) (interface{}, error) {
// 	// func (w *ApiRoot) Get(action interface{}) error {
// 	for range Only.Once {
// 		if w.Url == nil {
// 			w.Error = errors.New("SUNGRO API URL is invalid")
// 			break
// 		}
//
// 		request := endpoint.RequestRef()
// 		w.Error = reflect.VerifyOptionsRequired(request)
// 		if w.Error != nil {
// 			break
// 		}
//
// 		// objectName, actionName := GetNameOld(action)
// 		// httpRequest := FindInStruct(action, "RequestCommon")
// 		// httpResponse := FindInStruct(action, "ResponseCommon")
// 		//
// 		// objectName := GetStructName(object)
// 		// if objectName == "" {
// 		// 	w.Error = errors.New("invalid object name to structure")
// 		// 	break
// 		// }
// 		//
// 		// actionName := GetStructName(action)
// 		// if objectName == "" {
// 		// 	w.Error = errors.New("invalid action name to structure")
// 		// 	break
// 		// }
// 		//
// 		// requestString := Query(httpRequest)
// 		// if objectName == "" {
// 		// 	w.Error = errors.New("invalid httpRequest string for structure")
// 		// 	break
// 		// }
// 		//
// 		// responseString := Query(httpResponse)
// 		// if objectName == "" {
// 		// 	w.Error = errors.New("invalid httpResponse string for structure")
// 		// 	break
// 		// }
// 		//
// 		// u := fmt.Sprintf("%s?format=json&object=%s&action=%s%s",
// 		// 	w.Url.String(),
// 		// 	objectName,
// 		// 	actionName,
// 		// 	queryString,
// 		// )
//
// 		p, _ := json.Marshal(request)
//
// 		// w.httpRequest, w.Error = http.NewRequest("GET", u, nil)
// 		// if w.Error != nil {
// 		// 	break
// 		// }
// 		//
// 		// w.httpRequest.Header.Set("Authorization", w.Auth.GetAuthHeader())
//
// 		postUrl := fmt.Sprintf("%s%s", w.Url.String(), endpoint.GetUrl())
//
// 		for range Only.Twice {
// 			w.httpResponse, w.Error = http.Post(postUrl, "application/json", bytes.NewBuffer(p))
// 			if w.Error != nil {
// 				break
// 			}
//
// 			if strings.Contains(w.httpResponse.Status, "The access token provided is invalid") {
// 				// 401 Unauthorized The access token provided is invalid.
// 				// Will first attempt a refresh of the token OR re-login.
// 				w.Error = w.Auth.Login(&login.SunGroAuth {
// 					Expiry: "",
// 					AppKey:      "",
// 					UserAccount:    "",
// 					UserPassword:    "",
// 				})
// 				if w.Error != nil {
// 					w.Error = errors.New(w.httpResponse.Status)
// 					break
// 				}
// 				//w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
// 				continue
// 			}
//
// 			if w.httpResponse.StatusCode == 401 {
// 				w.Error = errors.New(w.httpResponse.Status)
// 				break
// 			}
//
// 			// All OK.
// 			break
// 		}
// 		//goland:noinspection GoUnhandledErrorResult
// 		defer w.httpResponse.Body.Close()
// 		if w.Error != nil {
// 			break
// 		}
//
// 		if w.httpResponse.StatusCode != 200 {
// 			w.Error = errors.New(fmt.Sprintf("API httpResponse is %s", w.httpResponse.Status))
// 			break
// 		}
//
// 		w.Body, w.Error = ioutil.ReadAll(w.httpResponse.Body)
// 		if w.Error != nil {
// 			break
// 		}
//
// 		if len(w.Body) == 0 {
// 			w.Error = errors.New("empty httpResponse")
// 			break
// 		}
//
// 		endpoint = endpoint.SetResponse(w.Body)
// 		w.Error = endpoint.IsResponseValid()
//
// 		if w.Error != nil {
// 			fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
// 			break
// 		}
// 	}
//
// 	return endpoint, w.Error
// }
