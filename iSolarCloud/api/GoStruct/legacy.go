package GoStruct


// func FindStart(fieldName string, Parent *Reflect, Current *Reflect, name EndPointPath) *Reflect {
// 	var ret Reflect
//
// 	for range Only.Once {
// 		if Current.Kind == reflect.Pointer {
// 			// Special case:
// 			// We're going to change the pointer to a proper object reference.
// 			if Current.IsNil {
// 				break
// 			}
// 			ref2 := Current.ValueOf.Elem().Interface()
// 			if valueTypes.IsNil(ref2) {
// 				break
// 			}
// 			Current.Init(Current, ref2, Current.DataStructure.Endpoint)
// 			if Current.IsNil {
// 				break
// 			}
// 			// DO NOT BREAK!
// 			// KEEP FIRST!
// 		}
//
// 		if Current.Kind == reflect.Struct {
// 			// Iterate over all available fields and read the tag value
// 			for index := 0; index < Current.Length; index++ {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, reflect.Value{})
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Struct {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		if Current.Kind == reflect.Slice {
// 			// Iterate over all available fields and read the tag value
// 			for index := 0; index < Current.Length; index++ {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, reflect.Value{})
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Slice {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		if Current.Kind == reflect.Map {
// 			// Iterate over all available fields and read the tag value
// 			for index, key := range Current.FieldVo.MapKeys() {
// 				var Child Reflect
// 				Child.SetByIndex(Parent, Current, index, key)
//
// 				if Child.FieldName == fieldName {
// 					ret = Child
// 					break
// 				}
//
// 				if Child.Kind != reflect.Map {
// 					continue
// 				}
//
// 				if Child.IsKnown() {
// 					continue
// 				}
//
// 				Child = *FindStart(fieldName, Current, &Child, name)
// 			}
// 			break
// 		}
//
// 		_, _ = fmt.Fprintf(os.Stderr,"ERROR: Field '%s' type not supported: Type %s\n", Current.FieldName, Current.Kind.String())
// 	}
//
// 	return &ret
// }
