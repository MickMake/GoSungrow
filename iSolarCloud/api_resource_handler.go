package iSolarCloud

type ResourceHandlers []ResourceHandler
type ResourceHandler interface {
	InitializeResourceType(*ResourceType) *ResourceType
}
