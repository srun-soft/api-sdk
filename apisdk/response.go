package apisdk

type Response struct {
	Code    int
	Message string
	Version string
	Data    interface{}
}

func NewResponse() *Response {
	return &Response{}
}
