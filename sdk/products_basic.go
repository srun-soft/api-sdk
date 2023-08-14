package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 产品相关

// ProductView 产品详情
func (APIClient) ProductView(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsView, data, http.MethodGet)
}

// ProductCreate 增加产品
func (APIClient) ProductCreate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsCreate, data)
}

// ProductDelete 删除产品
func (APIClient) ProductDelete(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsDelete, data, http.MethodDelete)
}

// ProductUpdate 修改产品
func (APIClient) ProductUpdate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsUpdate, data, http.MethodPut)
}

// ProductIndex 查询计费策略列表
func (APIClient) ProductIndex(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsIndex, data, http.MethodGet)
}
