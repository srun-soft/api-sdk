package sdk

import "srunsoft-api-sdk/tools"

// 密码相关

// SendCode 忘记密码重置接口【step1.发送短信验证码接口】
func (APIClient) SendCode(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersCode, data, "GET")
}

// ForgetResetPassword 忘记密码重置接口【step2.重置密码】
func (APIClient) ForgetResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersForgetResetPassword, data)
}

// CheckModifyPassword 验证是否需要修改密码
func (APIClient) CheckModifyPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(AuthCheckModifyPassword, data)
}

// GetPassword 查询用户密码 base64
func (APIClient) GetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersGetPassword, data, "GET")
}

// SuperResetPassword 修改密码高级接口
func (APIClient) SuperResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersSuperResetPassword, data, "PUT")
}

// ResetPassword 修改密码接口【根据老密码修改新密码】
func (APIClient) ResetPassword(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersResetPassword, data)
}
