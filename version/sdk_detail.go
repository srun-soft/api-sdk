package version

import "net/http"

type Detail interface {
	Details() Response
	DetailTop() Response
}

func (ac *ApiConfig) Details() Response {
	ac.UrlPath = Details
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) DetailTop() Response {
	ac.UrlPath = DetailTop
	ac.Method = http.MethodGet
	return ac.ApiCall()
}
