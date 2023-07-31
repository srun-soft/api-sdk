package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 在线相关

// BaseOnlineDrop 在线设备下线
func (APIClient) BaseOnlineDrop(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesOnlineDrop, data)
}

// BaseBatchOnlineDrop 根据用户名批量下线
func (APIClient) BaseBatchOnlineDrop(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesBatchOnlineDrop, data)
}

// Auth DHCP 认证
func (APIClient) Auth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.Auth, data)
}

// BaseOnlineEquipment 查询在线设备接口
func (APIClient) BaseOnlineEquipment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesOnlineEquipment, data, http.MethodGet)
}

// BaseOnlineData 查询当前在线用户数据
func (APIClient) BaseOnlineData(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesOnlineData, data, http.MethodGet)
}
