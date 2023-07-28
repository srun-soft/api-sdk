package sdk

import "srunsoft-api-sdk/tools"

// 日志相关

// LogLogin 认证日志
func (APIClient) LogLogin(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(LogsLogin, data, "GET")
}

// LogAuthErr 用户分析-报表
func (APIClient) LogAuthErr(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(LogsAuthErr, data, "GET")
}

// LogAuthErrorTotal 按认证错误原因统计次数
func (APIClient) LogAuthErrorTotal(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(LogsAuthErrorTotal, data, "GET")
}

// LogSys 系统日志
func (APIClient) LogSys(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(LogsSys, data, "GET")
}

// StatisticsDayByte 按天统计当月流量
func (APIClient) StatisticsDayByte(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(StatisticsDayByte, data, "GET")
}

// StatisticsDayTime 按天统计当月时长
func (APIClient) StatisticsDayTime(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(StatisticsDayTime, data, "GET")
}

// StatisticsMonthByte 按月统计流量
func (APIClient) StatisticsMonthByte(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(StatisticsMonthByte, data, "GET")
}

// FlowSearch Flow(流量)日志
func (APIClient) FlowSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FlowsSearch, data, "GET")
}

// AuditSearch 审计日志
func (APIClient) AuditSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(AuditsSearch, data, "GET")
}
