package sdk

import "srunsoft-api-sdk/tools"

// 产品相关

// ProductView 产品详情
func (APIClient) ProductView(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsView, data, "GET")
}

// ProductCreate 增加产品
func (APIClient) ProductCreate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsCreate, data)
}

// ProductDelete 删除产品
func (APIClient) ProductDelete(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsDelete, data, "DELETE")
}

// ProductUpdate 修改产品
func (APIClient) ProductUpdate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsUpdate, data, "PUT")
}

// ProductIndex 查询计费策略列表
func (APIClient) ProductIndex(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsIndex, data, "GET")
}
