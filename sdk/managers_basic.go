package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 管理员相关

// UserValidateManager 校验管理员
func (APIClient) UserValidateManager(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersValidateManager, data)
}

// UserAddManager 添加管理员
func (APIClient) UserAddManager(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersAddManager, data)
}

// UserResetPasswordManager 修改管理员密码接口（通过原密码）
func (APIClient) UserResetPasswordManager(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersResetPasswordManager, data, http.MethodPut)
}

// UserSuperResetPasswordManager 重置管理员密码
func (APIClient) UserSuperResetPasswordManager(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersSuperResetPasswordManager, data, http.MethodPut)
}
