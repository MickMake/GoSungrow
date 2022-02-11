package web

import (
	"GoSungro/Only"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"
)


type Web struct {
	Token
	Body  []byte
	Error error
	retry int

	Url      *url.URL
	client   http.Client
	request  *http.Request
	response *http.Response
}

func (w *Web) SetUrl(u string) error {
	w.Url, w.Error = url.Parse(u)
	if w.Error != nil {
		w.Url = nil
	}
	w.Token.Url = w.Url

	return w.Error
}

func (w *Web) GetOLD(s interface{}, url string, args ...interface{}) error {
	for range Only.Once {
		if w.Url == nil {
			w.Error = errors.New("SUNGRO API URL is invalid")
			break
		}

		url = fmt.Sprintf(w.Url.String()+url, args...)

		w.request, w.Error = http.NewRequest("GET", url, nil)
		if w.Error != nil {
			break
		}

		w.request.Header.Set("Authorization", w.Token.GetAuthHeader())

		//w.response, w.Error = http.Get(url)
		w.response, w.Error = w.client.Do(w.request)
		if w.Error != nil {
			break
		}
		//goland:noinspection GoUnhandledErrorResult
		defer w.response.Body.Close()
		if w.response.StatusCode != 200 {
			w.Error = errors.New(fmt.Sprintf("API response is %s", w.response.Status))
			break
		}

		w.Body, w.Error = ioutil.ReadAll(w.response.Body)
		if w.Error != nil {
			break
		}

		if len(w.Body) == 0 {
			w.Error = errors.New("empty response")
			break
		}

		w.Error = json.Unmarshal(w.Body, &s)
		if w.Error != nil {
			fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}
	}

	return w.Error
}

func (w *Web) Get(action interface{}, query interface{}, response interface{}) error {
	for range Only.Once {
		if w.Url == nil {
			w.Error = errors.New("SUNGRO API URL is invalid")
			break
		}

		w.Error = VerifyOptionsRequired(query)
		if w.Error != nil {
			break
		}

		objectName, actionName := GetName(action)
		//FindInStruct(action, "Request")
		//FindInStruct(action, "Response")

		//objectName := GetStructName(object)
		if objectName == "" {
			w.Error = errors.New("invalid object name to structure")
			break
		}

		//actionName := GetStructName(action)
		if objectName == "" {
			w.Error = errors.New("invalid action name to structure")
			break
		}

		queryString := Query(query)
		if objectName == "" {
			w.Error = errors.New("invalid query string for structure")
			break
		}

		u := fmt.Sprintf("%s?format=json&object=%s&action=%s%s",
			w.Url.String(),
			objectName,
			actionName,
			queryString,
		)
		// "?format=json&object=subscriber&action=count"

		//fmt.Printf("Object: %s\n", objectName)
		//fmt.Printf("Action: %s\n", actionName)
		////fmt.Printf("Action: %s\n", actionName)
		//fmt.Printf("Query: %s\n", query)
		//fmt.Printf("ApiUrl: %s\n", url)

		w.request, w.Error = http.NewRequest("GET", u, nil)
		if w.Error != nil {
			break
		}

		w.request.Header.Set("Authorization", w.Token.GetAuthHeader())

		for range Only.Twice {
			//w.response, w.Error = http.Get(url)
			w.response, w.Error = w.client.Do(w.request)
			if w.Error != nil {
				break
			}

			if strings.Contains(w.response.Status, "The access token provided is invalid") {
				// 401 Unauthorized The access token provided is invalid.
				// Will first attempt a refresh of the token OR re-login.
				w.Error = w.Login(nil)
				if w.Error != nil {
					w.Error = errors.New(w.response.Status)
					break
				}
				//w.Error = errors.New(fmt.Sprintf("API response is %s", w.response.Status))
				continue
			}

			if w.response.StatusCode == 401 {
				w.Error = errors.New(w.response.Status)
				break
			}

			// All OK.
			break
		}
		//goland:noinspection GoUnhandledErrorResult
		defer w.response.Body.Close()
		if w.Error != nil {
			break
		}

		if w.response.StatusCode != 200 {
			w.Error = errors.New(fmt.Sprintf("API response is %s", w.response.Status))
			break
		}

		w.Body, w.Error = ioutil.ReadAll(w.response.Body)
		if w.Error != nil {
			break
		}

		if len(w.Body) == 0 {
			w.Error = errors.New("empty response")
			break
		}

		w.Error = json.Unmarshal(w.Body, &response)
		if w.Error != nil {
			fmt.Printf("ERROR: Body is:\n%s\n", w.Body)
			break
		}
	}

	return w.Error
}

func Query(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	for id := 0; id < s.NumField(); id++ {
		value := fmt.Sprintf("%v", s.Field(id).Interface())
		if value == "" {
			continue
		}
		ret += fmt.Sprintf("&%s=%s",
			typeOf.Field(id).Tag.Get("json"),
			value,
		)
		//fmt.Printf("%d: %s %s = %v\n",
		//	i,
		//	typeOfT.Field(i).Name,
		//	s.Field(i).Type(),
		//	s.Field(i).Interface(),
		//)
	}

	return ret
}

func PrintHeader(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
	case "string":
		ret = fmt.Sprintf("%v", s)
	default:
		for id := 0; id < s.NumField(); id++ {
			//value := fmt.Sprintf("%v", s.Field(id).Interface())
			//if value == "" {
			//	continue
			//}
			ret += fmt.Sprintf("%s (%s),",
				typeOf.Field(id).Name,
				typeOf.Field(id).Tag.Get("json"),
			)
			//fmt.Printf("%d: %s %s = %v\n",
			//	i,
			//	typeOfT.Field(i).Name,
			//	s.Field(i).Type(),
			//	s.Field(i).Interface(),
			//)
		}
	}

	return ret
}

func PrintValue(i interface{}) string {
	var ret string

	s := reflect.ValueOf(i) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = fmt.Sprintf("%v", s)
	default:
		for id := 0; id < s.NumField(); id++ {
			ret += fmt.Sprintf("%v,", s.Field(id).Interface())
		}
	}

	return ret
}

func PrintAsJson(ref interface{}) (string, error) {
	var j string
	var err error

	for range Only.Once {
		var ret []byte
		ret, err = json.MarshalIndent(ref, "", "\t")
		if err != nil {
			break
		}

		j = string(ret)
	}

	return j, err
}

func HeaderAsArray(i interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(i) // .Elem()
	typeOf := s.Type()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, "Name")
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%s (%s)",
				typeOf.Field(id).Name,
				typeOf.Field(id).Tag.Get("json"),
			))
		}
	}

	return ret
}
func AsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	switch s.Kind().String() {
	case "string":
		ret = append(ret, fmt.Sprintf("%v", s))
	default:
		for id := 0; id < s.NumField(); id++ {
			ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
		}
	}

	return ret
}

//goland:noinspection GoUnusedFunction
func rowAsArray(ref interface{}) []interface{} {
	var ret []interface{}

	s := reflect.ValueOf(ref) // .Elem()
	for id := 0; id < s.NumField(); id++ {
		ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
	}

	return ret
}

//var headerStyleTable = map[string]json2csv.KeyStyle{
//	"jsonpointer": json2csv.JSONPointerStyle,
//	"slash":       json2csv.SlashStyle,
//	"dot":         json2csv.DotNotationStyle,
//	"dot-bracket": json2csv.DotBracketStyle,
//}
//
//func PrintAsCsv(ref interface{}) (string, error) {
//	var c string
//	var err error
//
//	for range Once.Once {
//		var ret []byte
//		ret, err = json.Marshal(ref)
//		if err != nil {
//			break
//		}
//
//		var results []json2csv.KeyValue
//		results, err = json2csv.JSON2CSV(string(ret))
//		if err != nil {
//			log.Fatal(err)
//		}
//		if len(results) == 0 {
//			break
//		}
//
//		csv := json2csv.NewCSVWriter(os.Stdout)
//		// header style (jsonpointer, slash, dot, dot-bracket
//		csv.HeaderStyle = headerStyleTable["dot"]
//		csv.Transpose = true
//
//		err = csv.WriteCSV(results)
//		if err != nil {
//			break
//		}
//
//	}
//
//	return c, err
//}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func ReflectAsJson(ref interface{}) string {
	var ret string

	for range Only.Once {
		switch reflect.TypeOf(ref).Kind() {
		case reflect.Slice:
		case reflect.Array:
			fmt.Println("The interface is a slice.")
			s := reflect.ValueOf(ref)
			ret += "["
			for i := 0; i < s.Len(); i++ {
				ret += ReflectAsJson(s.Index(i))
			}
			ret += "]"

		case reflect.Struct:
			s := reflect.ValueOf(ref) // .Elem()
			typeOf := s.Type()
			for i := 0; i < s.NumField(); i++ {
				value := fmt.Sprintf("%v", s.Field(i).Interface())
				if value == "" {
					continue
				}
				ret += fmt.Sprintf("%s:%s\n",
					typeOf.Field(i).Tag.Get("json"),
					value,
				)
				//fmt.Printf("%d: %s %s = %v\n",
				//	i,
				//	typeOfT.Field(i).Name,
				//	s.Field(i).Type(),
				//	s.Field(i).Interface(),
				//)
			}
		}
	}

	return ret
}

// FindInStruct Search a given structure for the State object and return its pointer.
//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func FindInStruct(ref interface{}, name string) interface{} {
	var ret interface{}

	for range Only.Once {
		v := reflect.ValueOf(ref)
		var e reflect.Value

		kind := v.Kind()
		// We're doing these checks to ensure ease of future expansion.
		if kind == reflect.Ptr {
			e = v.Elem()
			if e.Kind().String() == "ptr" {
				//PrintflnCyan("POINTER TO POINTER")	@TODO - DEBUG
				ret = FindInStruct(e.Addr().Elem().Interface(), name)
				break
			}
		} else if kind == reflect.Struct {
			// We can't handle a non-pointer, otherwise we get this...
			// reflect.flag.mustBeAssignable using unaddressable value
			e = v
			//Panic(PanicErrorNotGivenAPointer, v.String())
		} else {
			break
		}

		typeOfT := e.Type()
		for i := 0; i < e.NumField(); i++ {
			if typeOfT.Field(i).Name == name {
				ret = typeOfT.Field(i)
				break
			}

			//if typeOfT.Field(i).Name == name {
			//	state = e.Field(i).Interface().(*State)
			//	if state == nil {
			//		e.Field(i).Set(reflect.ValueOf(state))
			//	}
			//	break
			//}
		}
	}

	return ret
}

func GetName(ref interface{}) (string, string) {
	var packageName string
	var structName string

	str := reflect.TypeOf(ref).String()
	str = strings.TrimPrefix(str, "*")
	str = strings.ToLower(str)
	sa := strings.SplitN(str, ".", 2)
	switch len(sa) {
	case 0:
	case 1:
		packageName = sa[0]
	case 2:
		packageName = sa[0]
		structName = sa[1]
	}
	return packageName, structName
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetPackageName(ref interface{}) string {
	str, _ := GetName(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetStructName2(ref interface{}) string {
	str, _ := GetName(ref)
	return str
}

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetCmdName2(ref interface{}) string {
	str := reflect.TypeOf(ref).PkgPath()
	str = filepath.Base(str)
	str = strings.ToLower(str)
	return str
}

func GetStructName(ref interface{}) string {
	var ret string

	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
		ret = strings.ToLower(t.Elem().Name())
	} else {
		ret = strings.ToLower(t.Name())
	}

	return ret
}

// VerifyOptionsRequired Verify fields within the structure that are required.
func VerifyOptionsRequired(ref interface{}) error {
	var err error

	for range Only.Once {
		//required := GetOptionsRequired(ref)

		vo := reflect.ValueOf(ref)
		to := reflect.TypeOf(ref)

		// Iterate over all available fields and read the tag value
		for i := 0; i < vo.NumField(); i++ {
			fieldTo := to.Field(i)
			required := fieldTo.Tag.Get("required")
			if required == "" {
				continue
			}

			fieldVo := vo.Field(i)
			value := fmt.Sprintf("%v", fieldVo.Interface())
			if value == "" {
				err = errors.New(fmt.Sprintf("option '%s' is empty", fieldTo.Name))
				break
			}
		}
	}

	return err
}

type Required []string

//goland:noinspection GoUnusedFunction,GoUnusedExportedFunction
func GetOptionsRequired(ref interface{}) Required {
	var ret []string

	for range Only.Once {
		t := reflect.TypeOf(ref)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			required := field.Tag.Get("required")
			if required == "" {
				break
			}

			ret = append(ret, field.Name)
		}
	}

	return ret
}

func (r *Required) IsRequired(field string) bool {
	var ok bool
	for _, f := range *r {
		if f == field {
			ok = true
		}
	}
	return ok
}
func (r *Required) IsNotRequired(field string) bool {
	return !r.IsRequired(field)
}

// ******************************************************************************** //

func (w *Web) GetDomain() string {
	return w.Response.ResultMsg
}
func (w *Web) VerifyDomain(domain string) string {
	if domain == "" {
		domain = w.Response.ResultMsg
	}
	if domain == "." {
		domain = w.Response.ResultMsg
	}
	return domain
}

func (w *Web) GetUser() string {
	return w.Response.ResultMsg
}
func (w *Web) VerifyUser(user string) string {
	if user == "" {
		user = w.Response.ResultData.UserID
	}
	if user == "." {
		user = w.Response.ResultData.UserID
	}
	return user
}

func (w *Web) GetUserMac() string {
	return "."
}
func (w *Web) VerifyUserMac(user string) string {
	// @TODO - Check MAC is sane.
	return user
}

func (w *Web) GetUsername() string {
	return w.Response.ResultData.UserName
}
func (w *Web) VerifyUsername(name string) string {
	if name == "" {
		name = w.Response.ResultData.UserName
	}
	if name == "." {
		name = w.Response.ResultData.UserName
	}
	return name
}

func (w *Web) GetUserEmail() string {
	return w.Response.ResultData.Email
}
func (w *Web) VerifyUserEmail(email string) string {
	if email == "" {
		email = w.Response.ResultData.Email
	}
	if email == "." {
		email = w.Response.ResultData.Email
	}
	return email
}

func (w *Web) GetDisplayName() string {
	return w.Response.ResultData.UserAccount
}
func (w *Web) VerifyDisplayName(name string) string {
	if name == "" {
		name = w.Response.ResultData.UserAccount
	}
	if name == "." {
		name = w.Response.ResultData.UserAccount
	}
	return name
}
