package apisdk

import "net/http"

type User interface {
	Users() Response
	UserUpdate() Response
	UserDelete() Response
	UserUserStatusControl() Response
	UserValidateUsers() Response
	UserBalance() Response
	UserSendCode() Response
	UserMaxOnlineNum() Response
	UserBindingPhone() Response
	UserCode() Response
	UserForgetResetPassword() Response
	AuthCheckModifyPassword() Response
	UserGetPassword() Response
	UserSuperResetPassword() Response
	UserResetPassword() Response
	UserView() Response
	UserSuperSearch() Response
	UserSearch() Response
}

func (ac *ApiConfig) Users() Response {
	ac.UrlPath = Users
	return ac.ApiCall()
}

func (ac *ApiConfig) UserUpdate() Response {
	ac.UrlPath = UserUpdate
	ac.Method = http.MethodPut
	return ac.ApiCall()
}

func (ac *ApiConfig) UserDelete() Response {
	ac.UrlPath = UserDelete
	ac.Method = http.MethodDelete
	return ac.ApiCall()
}

func (ac *ApiConfig) UserUserStatusControl() Response {
	ac.UrlPath = UserUserStatusControl
	return ac.ApiCall()
}

func (ac *ApiConfig) UserValidateUsers() Response {
	ac.UrlPath = UserValidateUsers
	return ac.ApiCall()
}

func (ac *ApiConfig) UserBalance() Response {
	ac.UrlPath = UserBalance
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserSendCode() Response {
	ac.UrlPath = UserSendCode
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserMaxOnlineNum() Response {
	ac.UrlPath = UserMaxOnlineNum
	return ac.ApiCall()
}

func (ac *ApiConfig) UserBindingPhone() Response {
	ac.UrlPath = UserBindingPhone
	return ac.ApiCall()
}

func (ac *ApiConfig) UserCode() Response {
	ac.UrlPath = UserCode
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserForgetResetPassword() Response {
	ac.UrlPath = UserForgetResetPassword
	return ac.ApiCall()
}

func (ac *ApiConfig) AuthCheckModifyPassword() Response {
	ac.UrlPath = AuthCheckModifyPassword
	return ac.ApiCall()
}

func (ac *ApiConfig) UserGetPassword() Response {
	ac.UrlPath = UserGetPassword
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserSuperResetPassword() Response {
	ac.UrlPath = UserSuperResetPassword
	ac.Method = http.MethodPut
	return ac.ApiCall()
}

func (ac *ApiConfig) UserResetPassword() Response {
	ac.UrlPath = UserResetPassword
	return ac.ApiCall()
}

func (ac *ApiConfig) UserView() Response {
	ac.UrlPath = UserView
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserSuperSearch() Response {
	ac.UrlPath = UserSuperSearch
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) UserSearch() Response {
	ac.UrlPath = UserSearch
	ac.Method = http.MethodGet
	return ac.ApiCall()
}
