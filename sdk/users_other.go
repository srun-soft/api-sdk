package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 其他辅助接口

// UserStatusControl 修改用户状态
func (APIClient) UserStatusControl(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersStatusControl, data)
}

// UserValidate 校验用户密码接口
func (APIClient) UserValidate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersValidate, data)
}

// UserBalance 返回电子钱包余额接口
func (APIClient) UserBalance(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersBalance, data, http.MethodGet)
}

// UserSendCode 绑定/更换 手机号码接口 step1
func (APIClient) UserSendCode(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersSendCode, data, http.MethodGet)
}

// UserMaxOnlineNum 修改最大在线数接口
func (APIClient) UserMaxOnlineNum(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersMaxOnlineNum, data)
}

// UserBindingPhone 绑定/更换 手机号码接口 step2
func (APIClient) UserBindingPhone(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersBindingPhone, data, http.MethodPut)
}
