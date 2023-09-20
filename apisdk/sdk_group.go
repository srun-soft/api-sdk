package apisdk

import "net/http"

type Group interface {
	Groups() Response
	GroupSubscribe() Response
	GroupsCreate() Response
}

func (ac *ApiConfig) Groups() Response {
	ac.UrlPath = Groups
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) GroupSubscribe() Response {
	ac.UrlPath = GroupSubscribe
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) GroupsCreate() Response {
	ac.UrlPath = Groups
	return ac.ApiCall()
}
