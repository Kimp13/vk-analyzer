package network

import "net/url"

type PermissionResponse struct {
	Response uint32
}

var DefaultURL = url.URL{}
var DefaultQuery = DefaultURL.Query()
