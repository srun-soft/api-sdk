package apisdk

import "net/http"

type Operator interface {
	UserSearchMobilePhone() Response
	ProductCheckOperators() Response
	ProductOperators() Response
}

func (ac *ApiConfig) UserSearchMobilePhone() Response {
	ac.UrlPath = UserSearchMobilePhone
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductCheckOperators() Response {
	ac.UrlPath = ProductCheckOperators
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductOperators() Response {
	ac.UrlPath = ProductOperators
	return ac.ApiCall()
}
