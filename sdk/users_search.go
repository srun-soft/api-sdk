package sdk

import (
	"net/http"
	"srunsoft-api-sdk/tools"
)

// 查询用户相关接口

// UserView 查询用户详情
func (APIClient) UserView(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersView, data, http.MethodGet)
}

// UserSuperSearch 搜索用户高级接口
func (APIClient) UserSuperSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersSuperSearch, data, http.MethodGet)
}

// UserSearch 搜索用户接口
func (APIClient) UserSearch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersSearch, data, http.MethodGet)
}
