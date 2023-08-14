package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 上网明细相关

// Details 查询上网明细
func (APIClient) Details(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.Details, data, http.MethodGet)
}

// DetailTop 查询上网流量TOP排名
func (APIClient) DetailTop(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.DetailsTop, data, http.MethodGet)
}
