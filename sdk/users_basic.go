package sdk

import (
	"net/http"
	"srunsoft-api-sdk/tools"
)

// 用户基本操作

// CreateUser 添加用户
func (APIClient) CreateUser(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersCreate, data)
}

// UpdateUser 编辑用户
func (APIClient) UpdateUser(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersUpdate, data, http.MethodPut)
}

// DeleteUser 删除用户
func (APIClient) DeleteUser(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersDelete, data, http.MethodDelete)
}
