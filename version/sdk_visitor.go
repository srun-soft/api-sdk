package version

import "net/http"

type Visitor interface {
	UserVisitors() Response
	UserTokenVisitors() Response
	UserEventVisitors() Response
	VisitorViewEventVisitor() Response
	VisitorCreateEventVisitor() Response
	VisitorUpdateEventVisitor() Response
	VisitorDeleteEventVisitor() Response
	UserViewInvite() Response
	UserInviteVisitors() Response
	UserUseInvite() Response
	UserCreateInvite() Response
	UserDisabledInvite() Response
	VisitorCreateVisitorWhite() Response
	VisitorDeleteVisitorWhite() Response
}

func (ac *ApiConfig) UserVisitors() Response {
	ac.UrlPath = UserVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserTokenVisitors() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserEventVisitors() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorViewEventVisitor() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorCreateEventVisitor() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorUpdateEventVisitor() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorDeleteEventVisitor() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserViewInvite() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserInviteVisitors() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserUseInvite() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserCreateInvite() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) UserDisabledInvite() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorCreateVisitorWhite() Response {
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}

func (ac *ApiConfig) VisitorDeleteVisitorWhite() Response {
	ac.Method = http.MethodDelete
	ac.UrlPath = UserTokenVisitors
	return ac.ApiCall()
}
