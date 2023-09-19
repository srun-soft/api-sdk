package version

import (
	"net/url"
)

type Request struct {
	Params  url.Values
	Method  string
	UrlPath string
}

func NewRequest() *Request {
	return &Request{
		Params: url.Values{},
	}
}

func test() {
	sdk := &ApiConfig{}
	sdk.Version = 1
	sdk.Request.Params = url.Values{}
	sdk.Users()
}
