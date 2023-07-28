package sdk

import "srunsoft-api-sdk/tools"

// 用户组相关

// Groups 查询所有用户组
func (APIClient) Groups(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(GroupsSearchAll, data, http.MethodGet)
}

// GroupSubscribe 查询用户组可订购产品
func (APIClient) GroupSubscribe(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(GroupsSubscribe, data, http.MethodGet)
}

// GroupsCreate 添加用户组
func (APIClient) GroupsCreate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(GroupsCreate, data)
}
