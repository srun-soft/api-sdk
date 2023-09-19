package version

import "net/http"

type Messages interface {
	MessageNew() Response
	Message() Response
	MessageNotice() Response
}

func (ac *ApiConfig) MessageNew() Response {
	ac.UrlPath = MessageNew
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) Message() Response {
	ac.UrlPath = Message
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) MessageNotice() Response {
	ac.UrlPath = MessageNotice
	return ac.ApiCall()
}
