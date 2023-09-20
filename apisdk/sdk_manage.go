package apisdk

import "net/http"

type Manage interface {
	UserValidateManager() Response
	UserAddManager() Response
	UserResetPasswordManager() Response
	UserSuperResetPasswordManager() Response
}

func (ac *ApiConfig) UserValidateManager() Response {
	ac.UrlPath = UserValidateManager
	return ac.ApiCall()
}

func (ac *ApiConfig) UserAddManager() Response {
	ac.UrlPath = UserAddManager
	return ac.ApiCall()
}

func (ac *ApiConfig) UserResetPasswordManager() Response {
	ac.UrlPath = UserResetPasswordManager
	ac.Method = http.MethodPut
	return ac.ApiCall()
}

func (ac *ApiConfig) UserSuperResetPasswordManager() Response {
	ac.UrlPath = UserSuperResetPasswordManager
	ac.Method = http.MethodPut
	return ac.ApiCall()
}
