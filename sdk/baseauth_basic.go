package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 上网设备相关

// 无感知相关

// BaseListMacAuth 查询mac认证设备
func (APIClient) BaseListMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesListMacAuth, data, http.MethodGet)
}

// BaseCreateMacAuth 添加mac认证地址
func (APIClient) BaseCreateMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesCreateMacAuth, data)
}

// BaseDeleteMacAuth 删除mac认证地址
func (APIClient) BaseDeleteMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesDeleteMacAuth, data)
}

// BaseUpdateMacAuth 更新mac认证地址
func (APIClient) BaseUpdateMacAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesUpdateMacAuth, data)
}

// BasePhoneAuth 手机号无感知
func (APIClient) BasePhoneAuth(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.BasesPhoneAuth, data)
}
