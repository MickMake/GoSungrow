package reflection

// func GetStringFrom(ref interface{}, name string) string {
// 	var ret string
// 	for range Only.Once {
// 		vo := reflect.ValueOf(ref)
//
// 		switch vo.Kind() {
// 		case reflect.Struct:
// 			// Iterate over all available fields, looking for the field name.
// 			for i := 0; i < vo.NumField(); i++ {
// 				if vo.Type().Field(i).Name == name {
// 					ret = valueTypes.AnyToValueString(vo.Field(i).Interface(), 0, "")
// 					break
// 				}
// 			}
//
// 		case reflect.Array:
// 			fallthrough
// 		case reflect.Slice:
// 			// Iterate over all available entities, looking for the field name under struct or map.
// 			for i := 0; i < vo.Len(); i++ {
// 				ivo := reflect.ValueOf(vo.Index(i).Interface())
// 				foo := ivo.Interface()
// 				fmt.Sprintf("%s", foo)
// 				foo2 := vo.Index(i).Interface()
// 				fmt.Sprintf("%s", foo2)
// 				switch ivo.Kind() {
// 				case reflect.Struct:
// 					// Iterate over all available fields, looking for the field name.
// 					for ii := 0; ii < ivo.NumField(); ii++ {
// 						if ivo.Type().Field(ii).Name == name {
// 							ret = valueTypes.AnyToValueString(ivo.Field(i).Interface(), 0, "")
// 							break
// 						}
// 					}
// 				case reflect.Map:
// 					// Iterate over map, looking for the key name.
// 					for _, key := range ivo.MapKeys() {
// 						if key.String() == name {
// 							ret = valueTypes.AnyToValueString(ivo.MapIndex(key).Interface(), 0, "")
// 							break
// 						}
// 					}
// 				}
// 				// ret = GetStringFrom(vo.Index(i).Interface(), name)
// 				// if ret != "" {
// 				// 	break
// 				// }
// 			}
//
// 		case reflect.Map:
// 			// Iterate over map, looking for the key name.
// 			for _, key := range vo.MapKeys() {
// 				if key.String() == name {
// 					ret = valueTypes.AnyToValueString(vo.MapIndex(key).Interface(), 0, "")
// 					break
// 				}
// 			}
// 		}
// 	}
//
// 	return ret
// }
//
// type StructKey struct {
// 	Name      string
// 	JsonName  string
// 	JsonValue string
// 	Required  string
// 	Value     string
// 	Type      reflect.Type
// 	Interface interface{}
// }
// type StructKeys map[string]StructKey
// func GetStructKeys(ref interface{}, keys ...string) StructKeys {
// 	ret := make(StructKeys)
//
// 	s := New(ref)
// 	if len(keys) == 0 {
// 		keys = s.Names()
// 	}
//
// 	keyMap := make(map[string]bool)
// 	for _, k := range keys {
// 		keyMap[k] = true
// 	}
//
// 	for _, k := range New(ref).Fields() {
// 		if _, ok := keyMap[k.Name()]; !ok {
// 			continue
// 		}
//
// 		jValue := ""
// 		value := ""
// 		switch k.value.Type().Name() {
// 			case "string":
// 				value = k.Value().(string)
// 			case "int":
// 				v := k.Value().(int)
// 				value = strconv.FormatInt(int64(v), 10)
// 			case "int64":
// 				value = strconv.FormatInt(k.Value().(int64), 10)
// 			case "float64":
// 				value = strconv.FormatFloat(k.Value().(float64), 'f', -1, 64)
// 			default:
// 				j, e := json.Marshal(k.Value())
// 				if e == nil {
// 					jValue = string(j)
// 				} else {
// 					jValue = fmt.Sprintf("%v", k.Value())
// 				}
// 		}
//
// 		ret[k.Name()] = StructKey {
// 			Name:      k.Name(),
// 			JsonName:  k.Tag("json"),
// 			JsonValue: jValue,
// 			Value:     value,
// 			Required:  k.Tag("required"),
// 			Type:      k.value.Type(),
// 			Interface: k.Value(),
// 		}
// 	}
//
// 	return ret
// }
//
//
// func GetCallerPackage(skip int) string {
// 	var ret string
// 	if pc, _, _, ok := runtime.Caller(skip); ok {
// 		funcName := runtime.FuncForPC(pc).Name()
// 		slash := strings.LastIndexByte(funcName, '/')
// 		if slash < 0 {
// 			slash = 0
// 		}
// 		dot := strings.IndexByte(funcName, '.')
// 		ret = funcName[slash+1:dot]
// 	}
// 	return ret
// }
//
// func hash(s string) uint32 {
// 	h := fnv.New32a()
// 	_, _ = h.Write([]byte(s))
// 	return h.Sum32()
// }
//
// func PackageName(trim string, v interface{}) string {
// 	var ret string
// 	for range Only.Once {
// 		if v == nil {
// 			break
// 		}
//
// 		val := reflect.ValueOf(v)
// 		if val.Kind() == reflect.Ptr {
// 			ret = val.Elem().Type().PkgPath()
// 			break
// 		}
//
// 		ret2 := val.Type().Name()
// 		ret3 := val.Type().String()
// 		ret = val.Type().PkgPath()
// 		ret = strings.TrimPrefix(ret, trim)
//
// 		ret = fmt.Sprintf("%s\t%s\t%s\n", ret, ret2, ret3)
// 	}
// 	return ret
// }
//
// func Query(i interface{}) string {
// 	var ret string
//
// 	s := reflect.ValueOf(i) // .Elem()
// 	typeOf := s.Type()
// 	for id := 0; id < s.NumField(); id++ {
// 		value := fmt.Sprintf("%v", s.Field(id).Interface())
// 		if value == "" {
// 			continue
// 		}
// 		ret += fmt.Sprintf("&%s=%s",
// 			typeOf.Field(id).Tag.Get("json"),
// 			value,
// 		)
// 		// fmt.Printf("%d: %s %s = %v\n",
// 		//	i,
// 		//	typeOfT.Field(i).Name,
// 		//	s.Field(i).Type(),
// 		//	s.Field(i).Interface(),
// 		// )
// 	}
//
// 	return ret
// }
//
// func PrintHeader(i interface{}) string {
// 	var ret string
//
// 	s := reflect.ValueOf(i) // .Elem()
// 	typeOf := s.Type()
// 	switch s.Kind().String() {
// 		case "string":
// 			ret = fmt.Sprintf("%v", s)
// 		default:
// 			for id := 0; id < s.NumField(); id++ {
// 				// value := fmt.Sprintf("%v", s.Field(id).Interface())
// 				// if value == "" {
// 				//	continue
// 				// }
// 				ret += fmt.Sprintf("%s (%s),",
// 					typeOf.Field(id).Name,
// 					typeOf.Field(id).Tag.Get("json"),
// 				)
// 				// fmt.Printf("%d: %s %s = %v\n",
// 				//	i,
// 				//	typeOfT.Field(i).Name,
// 				//	s.Field(i).Type(),
// 				//	s.Field(i).Interface(),
// 				// )
// 			}
// 	}
//
// 	return ret
// }
//
// func PrintValue(i interface{}) string {
// 	var ret string
//
// 	s := reflect.ValueOf(i) // .Elem()
// 	switch s.Kind().String() {
// 	case "string":
// 		ret = fmt.Sprintf("%v", s)
// 	default:
// 		for id := 0; id < s.NumField(); id++ {
// 			ret += fmt.Sprintf("%v,", s.Field(id).Interface())
// 		}
// 	}
//
// 	return ret
// }
//
// func PrintAsJson(ref interface{}) (string, error) {
// 	var j string
// 	var err error
//
// 	for range Only.Once {
// 		var ret []byte
// 		ret, err = json.MarshalIndent(ref, "", "\t")
// 		if err != nil {
// 			break
// 		}
//
// 		j = string(ret)
// 	}
//
// 	return j, err
// }
//
// func HeaderAsArray(i interface{}) []interface{} {
// 	var ret []interface{}
//
// 	s := reflect.ValueOf(i) // .Elem()
// 	typeOf := s.Type()
// 	switch s.Kind().String() {
// 		case "string":
// 			ret = append(ret, "Name")
// 		default:
// 			for id := 0; id < s.NumField(); id++ {
// 				ret = append(ret, fmt.Sprintf("%s (%s)",
// 					typeOf.Field(id).Name,
// 					typeOf.Field(id).Tag.Get("json"),
// 				))
// 			}
// 	}
//
// 	return ret
// }
//
// func AsArray(ref interface{}) []interface{} {
// 	var ret []interface{}
//
// 	s := reflect.ValueOf(ref) // .Elem()
// 	switch s.Kind().String() {
// 		case "string":
// 			ret = append(ret, fmt.Sprintf("%v", s))
// 		default:
// 			for id := 0; id < s.NumField(); id++ {
// 				ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
// 			}
// 	}
//
// 	return ret
// }
//
// func rowAsArray(ref interface{}) []interface{} {
// 	var ret []interface{}
//
// 	s := reflect.ValueOf(ref) // .Elem()
// 	for id := 0; id < s.NumField(); id++ {
// 		ret = append(ret, fmt.Sprintf("%v", s.Field(id).Interface()))
// 	}
//
// 	return ret
// }
//
// var headerStyleTable = map[string]json2csv.KeyStyle{
//	"jsonpointer": json2csv.JSONPointerStyle,
//	"slash":       json2csv.SlashStyle,
//	"dot":         json2csv.DotNotationStyle,
//	"dot-bracket": json2csv.DotBracketStyle,
// }
//
// func PrintAsCsv(ref interface{}) (string, error) {
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
// }
//
// func ReflectAsJson(ref interface{}) string {
// 	var ret string
//
// 	for range Only.Once {
// 		switch reflect.TypeOf(ref).Kind() {
// 			case reflect.Slice:
// 				fallthrough
// 			case reflect.Array:
// 				fmt.Println("The interface is a slice.")
// 				s := reflect.ValueOf(ref)
// 				ret += "["
// 				for i := 0; i < s.Len(); i++ {
// 					ret += ReflectAsJson(s.Index(i))
// 				}
// 				ret += "]"
//
// 			case reflect.Struct:
// 				s := reflect.ValueOf(ref) // .Elem()
// 				typeOf := s.Type()
// 				for i := 0; i < s.NumField(); i++ {
// 					value := fmt.Sprintf("%v", s.Field(i).Interface())
// 					if value == "" {
// 						continue
// 					}
// 					ret += fmt.Sprintf("%s:%s\n",
// 						typeOf.Field(i).Tag.Get("json"),
// 						value,
// 					)
// 					// fmt.Printf("%d: %s %s = %v\n",
// 					//	i,
// 					//	typeOfT.Field(i).Name,
// 					//	s.Field(i).Type(),
// 					//	s.Field(i).Interface(),
// 					// )
// 			}
// 		}
// 	}
//
// 	return ret
// }
//
// // FindInStruct Search a given structure for the State object and return its pointer.
// func FindInStruct(ref interface{}, name string) interface{} {
// 	var ret interface{}
//
// 	for range Only.Once {
// 		v := reflect.ValueOf(ref)
// 		var e reflect.Value
//
// 		kind := v.Kind()
// 		// We're doing these checks to ensure ease of future expansion.
// 		if kind == reflect.Ptr {
// 			e = v.Elem()
// 			if e.Kind().String() == "ptr" {
// 				// PrintflnCyan("POINTER TO POINTER")	@TODO - DEBUG
// 				ret = FindInStruct(e.Addr().Elem().Interface(), name)
// 				break
// 			}
// 		} else if kind == reflect.Struct {
// 			// We can't handle a non-pointer, otherwise we get this...
// 			// reflect.flag.mustBeAssignable using unaddressable value
// 			e = v
// 			// Panic(PanicErrorNotGivenAPointer, v.String())
// 		} else {
// 			break
// 		}
//
// 		typeOfT := e.Type()
// 		for i := 0; i < e.NumField(); i++ {
// 			if typeOfT.Field(i).Name == name {
// 				ret = typeOfT.Field(i)
// 				break
// 			}
//
// 			// if typeOfT.Field(i).Name == name {
// 			//	state = e.Field(i).Interface().(*State)
// 			//	if state == nil {
// 			//		e.Field(i).Set(reflect.ValueOf(state))
// 			//	}
// 			//	break
// 			// }
// 		}
// 	}
//
// 	return ret
// }
//
// func GetNameOld(ref interface{}) (string, string) {
// 	var packageName string
// 	var structName string
//
// 	str := reflect.TypeOf(ref).String()
// 	str = strings.TrimPrefix(str, "*")
// 	str = strings.ToLower(str)
// 	sa := strings.SplitN(str, ".", 2)
// 	switch len(sa) {
// 		case 0:
// 		case 1:
// 			packageName = sa[0]
// 		case 2:
// 			packageName = sa[0]
// 			structName = sa[1]
// 	}
// 	return packageName, structName
// }
//
// func GetPackageName(ref interface{}) string {
// 	str, _ := GetNameOld(ref)
// 	return str
// }
//
// func GetStructName2(ref interface{}) string {
// 	str, _ := GetNameOld(ref)
// 	return str
// }
//
// func GetCmdName2(ref interface{}) string {
// 	str := reflect.TypeOf(ref).PkgPath()
// 	str = filepath.Base(str)
// 	str = strings.ToLower(str)
// 	return str
// }
//
// func GetStructName(ref interface{}) string {
// 	var ret string
//
// 	if t := reflect.TypeOf(ref); t.Kind() == reflect.Ptr {
// 		ret = strings.ToLower(t.Elem().Name())
// 	} else {
// 		ret = strings.ToLower(t.Name())
// 	}
//
// 	return ret
// }
