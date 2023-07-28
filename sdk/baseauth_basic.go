package sdk

import (
	"net/http"
	"srunsoft-api-sdk/tools"
)

// 上网设备相关

// 无感知相关

// BaseListMacAuth 查询mac认证设备
func (APIClient) BaseListMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesListMacAuth, data, http.MethodGet)
}

// BaseCreateMacAuth 添加mac认证地址
func (APIClient) BaseCreateMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesCreateMacAuth, data)
}

// BaseDeleteMacAuth 删除mac认证地址
func (APIClient) BaseDeleteMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesDeleteMacAuth, data)
}

// BaseUpdateMacAuth 更新mac认证地址
func (APIClient) BaseUpdateMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesUpdateMacAuth, data)
}

// BasePhoneAuth 手机号无感知
func (APIClient) BasePhoneAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(BasesPhoneAuth, data)
}
