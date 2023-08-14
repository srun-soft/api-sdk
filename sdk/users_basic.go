package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 用户基本操作

// Users  添加用户
func (APIClient) Users(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersCreate, data)
}

// UserUpdate 编辑用户
func (APIClient) UserUpdate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersUpdate, data, http.MethodPut)
}

// UserDelete 删除用户
func (APIClient) UserDelete(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersDelete, data, http.MethodDelete)
}
