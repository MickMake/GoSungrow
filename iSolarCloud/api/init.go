package api

// func Initialize(contentHandlers ContentHandlers) {
// 	for _, ch := range contentHandlers {
// 		contentHandler, ok := ch.(ContentHandler)
// 		if !ok {
// 			LogError(ErrorArgs{
// 				Instance:  ch,
// 				ErrorType: DoesNotImplement,
// 				Interface: (ContentHandler)(nil),
// 			})
// 			continue
// 		}
// 		ct := NewContentType(ch)
// 		ct = contentHandler.InitializeContentType(ct)
// 		contentTypes[ct.GetSlug()] = ct
// 	}
//
// }
