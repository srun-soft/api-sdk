package version

import "net/http"

type BaseAuth interface {
	BaseListMacAuth() Response
	BaseCreateMacAuth() Response
	BaseDeleteMacAuth() Response
	BaseUpdateMacAuth() Response
	BasePhoneAuth() Response
	BaseMacs() Response
	BaseCreateMac() Response
	BaseDeleteMac() Response
	BaseUpdateMac() Response
	BaseUpdateVlan() Response
	BaseGetIp() Response
	BaseBindIp() Response
	BaseOnlineDrop() Response
	BaseBatchOnlineDrop() Response
	Auth() Response
	BaseOnlineEquipment() Response
	BaseOnlineData() Response
}

func (ac *ApiConfig) BaseListMacAuth() Response {
	ac.UrlPath = BaseListMacAuth
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseCreateMacAuth() Response {
	ac.UrlPath = BaseListMacAuth
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseDeleteMacAuth() Response {
	ac.UrlPath = BaseDeleteMacAuth
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseUpdateMacAuth() Response {
	ac.UrlPath = BaseUpdateMacAuth
	return ac.ApiCall()
}

func (ac *ApiConfig) BasePhoneAuth() Response {
	ac.UrlPath = BasePhoneAuth
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseMacs() Response {
	ac.UrlPath = BaseMacs
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseCreateMac() Response {
	ac.UrlPath = BaseCreateMac
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseDeleteMac() Response {
	ac.UrlPath = BaseDeleteMac
	ac.Method = http.MethodDelete
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseUpdateMac() Response {
	ac.UrlPath = BaseUpdateMac
	ac.Method = http.MethodPut
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseUpdateVlan() Response {
	ac.UrlPath = BaseUpdateVlan
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseGetIp() Response {
	ac.UrlPath = BaseGetIp
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseBindIp() Response {
	ac.UrlPath = BaseBindIp
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseOnlineDrop() Response {
	ac.UrlPath = BaseOnlineDrop
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseBatchOnlineDrop() Response {
	ac.UrlPath = BaseBatchOnlineDrop
	return ac.ApiCall()
}

func (ac *ApiConfig) Auth() Response {
	ac.UrlPath = Auth
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseOnlineEquipment() Response {
	ac.UrlPath = BaseOnlineEquipment
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) BaseOnlineData() Response {
	ac.UrlPath = BaseOnlineData
	ac.Method = http.MethodGet
	return ac.ApiCall()
}
