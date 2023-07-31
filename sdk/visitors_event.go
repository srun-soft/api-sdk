package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 事件访客相关

// UserEventVisitors 访客开户
func (APIClient) UserEventVisitors(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersEventVisitors, data)
}

// VisitorViewEventVisitor 查询访客详情
func (APIClient) VisitorViewEventVisitor(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.EventVisitorView, data, http.MethodGet)
}

// VisitorCreateEventVisitor 添加事件访客
func (APIClient) VisitorCreateEventVisitor(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.EventVisitorCreate, data)
}

// VisitorUpdateEventVisitor 修改事件访客
func (APIClient) VisitorUpdateEventVisitor(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.EventVisitorUpdate, data)
}

// VisitorDeleteEventVisitor 删除事件访客
func (APIClient) VisitorDeleteEventVisitor(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.EventVisitorDelete, data)
}
