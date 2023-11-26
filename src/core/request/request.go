package request

import "net/http"


type InternalRequest struct {
	Path string
	Body map[string]string
	Headers map[string]string
	Method string
}


func GenerateInternalRequest(req *http.Request) InternalRequest {
	return InternalRequest{}
}