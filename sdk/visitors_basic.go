package sdk

import (
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 访客相关

// UserVisitors 访客开户
func (APIClient) UserVisitors(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersVisitors, data)
}

// UserTokenVisitors 二维码访客开户
func (APIClient) UserTokenVisitors(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersTokenVisitors, data)
}
