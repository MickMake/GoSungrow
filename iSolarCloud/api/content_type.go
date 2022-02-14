package api


// var contentTypes ContentTypes
//
// func init() {
// 	contentTypes = make(ContentTypes, 0)
// }
//
// type ContentTable string
// type ContentTypeSlug string
// type ContentTypes map[ContentTypeSlug]*ContentType
//
// func GetContentTypes() ContentTypes {
// 	return contentTypes
// }
//
// type DbFields []DbField
// type DbField string
//
// //
// // Nil instances are created by using (*<TypeName>)(nil)
// // or (<TypeName>)(nil), e.g. (*ContentType)(nil)
// //
// type ContentType struct {
// 	//
// 	// Lowercase version of the Struct name w/o the package name
// 	//   e.g. syncer.ContentType => 'contenttype'
// 	//
// 	slug ContentTypeSlug
//
// 	//
// 	// Table name in MySQL database to store content for this content type.
// 	//
// 	table ContentTable
//
// 	DbFields DbFields
//
// 	DbOnDupKey DbFields
//
// 	//
// 	// Nil instance of the type to contain the actual content
// 	// after unmarshalling JSON from API response
// 	//
// 	InstanceType interface{}
//
// 	//
// 	// Nil instance of the type that can be used to unmarshal JSON
// 	// for a API response containing a single "EndPoint."
// 	//
// 	ResourceType interface{}
//
// 	//
// 	// Nil instance of the type that can be used to unmarshal JSON
// 	// for a API response containing a multiples "resources."
// 	//
// 	CollectionType interface{}
//
// 	//
// 	// The MIME type returned by API responses. Typically "application/json"
// 	// but we envision supporting "application/xml"
// 	//
// 	MimeType MimeType
//
// 	//
// 	// Content with Api Entry Points have direct URLs that do
// 	// not need to be discovered in other API responses.
// 	//
// 	IsApiEntryPoint bool
//
// 	//
// 	// Content with a Download Priority can be downloaded directly from the API.
// 	// Content w/o Download Priority is found embedded in downloaded content.
// 	//
// 	DownloadPriority int
// }
//
// func (me *ContentType) SetSlug(slug ContentTypeSlug) {
// 	me.slug = slug
// }
//
// func (me *ContentType) GetSlug() ContentTypeSlug {
// 	if me.slug == "" {
// 		s := strings.Split(reflect.TypeOf(me.InstanceType).String(), ".")[1]
// 		me.slug = ContentTypeSlug(ToSnakeCase(s))
// 	}
// 	return me.slug
// }
//
// func (me *ContentType) SetTable(table ContentTable) {
// 	me.table = table
// }
//
// func (me *ContentType) GetTable() ContentTable {
// 	if me.table == "" {
// 		me.table = ContentTable(me.GetSlug())
// 	}
// 	return me.table
// }
//
// func NewContentType(instanceType interface{}) *ContentType {
// 	return &ContentType{
// 		InstanceType: instanceType,
// 		MimeType:     ApplicationJson,
// 	}
// }
