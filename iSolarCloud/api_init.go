package iSolarCloud

import "GoSungro/iSolarCloud/api"


func Initialize(contentHandlers api.ContentHandlers) {
	for _, ch := range contentHandlers {
		resourceHandler, ok := ch.(ResourceHandler)
		if !ok {
			continue
		}
		rt := NewResourceType(ch)
		rt = resourceHandler.InitializeResourceType(rt)
		resourceTypes[rt.GetSlug()] = rt
	}

}

type CommonAttributes struct {
	Appkey       string `json:"appkey"`
	SysCode      string `json:"sys_code"`
	UserAccount  string `json:"user_account"`
	UserPassword string `json:"user_password"`
}
