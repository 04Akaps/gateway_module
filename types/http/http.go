package http

import (
	_error "github.com/04Akaps/gateway_module/types/error"
)

type HttpMethod string

const (
	GET    = HttpMethod("GET")    // RFC 7231, 4.3.1
	POST   = HttpMethod("POST")   // RFC 7231, 4.3.3
	DELETE = HttpMethod("DELETE") // RFC 7231, 4.3.5
	PUT    = HttpMethod("PUT")    // RFC 7231, 4.3.4
)

func (h HttpMethod) ToString() string {
	return string(h)
}

type GetType string

const (
	QUERY = GetType("query")
	URL   = GetType("url")
)

func (h GetType) CheckType() error {
	switch h {
	case QUERY, URL:
		return nil
	default:
		return _error.CheckType_Error
	}
}

func (h GetType) ToString() string {
	return string(h)
}
