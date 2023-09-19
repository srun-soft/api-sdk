package version

import "net/http"

type Logs interface {
	LogLogin() Response
	LogAuthErr() Response
	LogAuthErrorTotal() Response
	LogSys() Response
	StatisticsDayByte() Response
	StatisticsDayTime() Response
	StatisticsMonthByte() Response
	FlowSearch() Response
	AuditSearch() Response
}

func (ac *ApiConfig) LogLogin() Response {
	ac.UrlPath = LogLogin
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) LogAuthErr() Response {
	ac.UrlPath = LogAuthErr
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) LogAuthErrorTotal() Response {
	ac.UrlPath = LogAuthErrorTotal
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) LogSys() Response {
	ac.UrlPath = LogSys
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) StatisticsDayByte() Response {
	ac.UrlPath = StatisticsDayByte
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) StatisticsDayTime() Response {
	ac.UrlPath = StatisticsDayTime
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) StatisticsMonthByte() Response {
	ac.UrlPath = StatisticsMonthByte
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FlowSearch() Response {
	ac.UrlPath = FlowSearch
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) AuditSearch() Response {
	ac.UrlPath = AuditSearch
	ac.Method = http.MethodGet
	return ac.ApiCall()
}
