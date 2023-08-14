package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 密码相关

// UserCode 忘记密码重置接口【step1.发送短信验证码接口】
func (APIClient) UserCode(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersCode, data, http.MethodGet)
}

// UserForgetResetPassword  忘记密码重置接口【step2.重置密码】
func (APIClient) UserForgetResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersForgetResetPassword, data)
}

// AuthCheckModifyPassword 验证是否需要修改密码
func (APIClient) AuthCheckModifyPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.AuthCheckModifyPassword, data)
}

// UserGetPassword 查询用户密码 base64
func (APIClient) UserGetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersGetPassword, data, http.MethodGet)
}

// UserSuperResetPassword 修改密码高级接口
func (APIClient) UserSuperResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersSuperResetPassword, data, http.MethodPut)
}

// UserResetPassword 修改密码接口【根据老密码修改新密码】
func (APIClient) UserResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersResetPassword, data)
}
