package sdk

import (
	"net/http"
	"srunsoft-api-sdk/tools"
)

// 上网明细相关

// Details 查询上网明细
func (APIClient) Details(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(Details, data, http.MethodGet)
}

// DetailTop 查询上网流量TOP排名
func (APIClient) DetailTop(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(DetailsTop, data, http.MethodGet)
}
