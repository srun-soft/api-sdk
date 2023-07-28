package sdk

import "srunsoft-api-sdk/tools"

// 访客白名单

// VisitorCreateVisitorWhite 添加访客白名单
func (APIClient) VisitorCreateVisitorWhite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(VisitorsWhiteCreate, data)
}

// VisitorDeleteVisitorWhite 删除访客白名单
func (APIClient) VisitorDeleteVisitorWhite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(VisitorsWhiteDelete, data, "DELETE")
}
