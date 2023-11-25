package request


type InternalRequest struct {
	Path string
	Body map[string]string
	Headers map[string]string
	Method string
}