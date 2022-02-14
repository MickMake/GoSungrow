package iSolarCloud

//
//
// var resourceTypes ResourceTypes
//
// func init() {
// 	resourceTypes = make(ResourceTypes, 0)
// }
//
// type ResourceTypeSlug string
// type ResourceTypes map[ResourceTypeSlug]*ResourceType
//
// func GetResourceTypes() ResourceTypes {
// 	return resourceTypes
// }
//
// type ResourceType struct {
// 	slug         ResourceTypeSlug
// 	ResultData   interface{}
// 	ReqSerialNum string `json:"req_serial_num"`
// 	ResultCode   string `json:"result_code"`
// 	ResultMsg string `json:"result_msg"`
// 	//HasImages         bool
// 	//HasGeoProfile     bool
// 	//HasCountries      bool
// 	//HasAudience       bool
// 	//HasGenre          bool
// 	//HasPlatforms      bool
// 	//HasLinks          bool
// 	//HasTopics         bool
// 	//HasVideos         bool
// 	//HasCaptions       bool
// 	//HasAvailabilities bool
// }
//
// func NewResourceType(instanceType interface{}) *ResourceType {
// 	return &ResourceType{
// 		ResultData: instanceType,
// 	}
// }
//
// func (me *ResourceType) SetSlug(slug ResourceTypeSlug) {
// 	me.slug = slug
// }
//
// func (me *ResourceType) GetSlug() ResourceTypeSlug {
// 	for range Only.Once {
// 		if me.slug != "" {
// 			break
// 		}
// 		s := strings.Split(reflect.TypeOf(me.ResultData).String(), ".")[1]
// 		me.slug = ResourceTypeSlug(api.ToSnakeCase(s))
// 	}
// 	return me.slug
// }