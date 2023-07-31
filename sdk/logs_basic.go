package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 日志相关

// LogLogin 认证日志
func (APIClient) LogLogin(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.LogsLogin, data, http.MethodGet)
}

// LogAuthErr 用户分析-报表
func (APIClient) LogAuthErr(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.LogsAuthErr, data, http.MethodGet)
}

// LogAuthErrorTotal 按认证错误原因统计次数
func (APIClient) LogAuthErrorTotal(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.LogsAuthErrorTotal, data, http.MethodGet)
}

// LogSys 系统日志
func (APIClient) LogSys(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.LogsSys, data, http.MethodGet)
}

// StatisticsDayByte 按天统计当月流量
func (APIClient) StatisticsDayByte(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.StatisticsDayByte, data, http.MethodGet)
}

// StatisticsDayTime 按天统计当月时长
func (APIClient) StatisticsDayTime(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.StatisticsDayTime, data, http.MethodGet)
}

// StatisticsMonthByte 按月统计流量
func (APIClient) StatisticsMonthByte(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.StatisticsMonthByte, data, http.MethodGet)
}

// FlowSearch Flow(流量)日志
func (APIClient) FlowSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FlowsSearch, data, http.MethodGet)
}

// AuditSearch 审计日志
func (APIClient) AuditSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.AuditsSearch, data, http.MethodGet)
}
