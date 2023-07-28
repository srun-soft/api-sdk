package sdk

import "srunsoft-api-sdk/tools"

// 设备绑定相关

// BaseMacs 获取用户已绑定的设备列表
func (APIClient) BaseMacs(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesMacs, data, "GET")
}

// BaseCreateMac 增加用户设备（mac地址）
func (APIClient) BaseCreateMac(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesCreateMac, data)
}

// BaseDeleteMac 删除用户设置（mac地址）
func (APIClient) BaseDeleteMac(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesDeleteMac, data, "DELETE")
}

// BaseUpdateMac 更新用户设置（mac地址）
func (APIClient) BaseUpdateMac(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesUpdateMac, data, "PUT")
}

// BaseUpdateVlan 绑定/更换用户的vlan地址
func (APIClient) BaseUpdateVlan(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesUpdateVlan, data)
}

// BaseGetIp 根据用户名搜索可用/已用IP
func (APIClient) BaseGetIp(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesGetIp, data, "GET")
}

// BaseBindIp 绑定IP地址
func (APIClient) BaseBindIp(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesBindIp, data)
}
