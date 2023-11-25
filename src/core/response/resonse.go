package response


type InternalResponse struct {
	Body string
	Headers map[string]string
	Status int
}


func CreateFakeResponse(ipAddress string) InternalResponse {
	return InternalResponse{
		Body: ipAddress,
		Headers: make(map[string]string),
		Status: 200,
	}
}