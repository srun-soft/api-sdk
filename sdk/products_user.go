package sdk

import "srunsoft-api-sdk/tools"

// 用户上网策略相关

// 产品相关

// ProductProductRecharge 产品缴费
func (APIClient) ProductProductRecharge(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsRecharge, data)
}

// ProductRecharge 产品缴费高级接口
func (APIClient) ProductRecharge(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsRechargeAdvanced, data)
}

// ProductUseNumber 产品使用人数
func (APIClient) ProductUseNumber(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsUseNumber, data, http.MethodGet)
}

// ProductSubscribe 订购产品
func (APIClient) ProductSubscribe(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsSubscribe, data)
}

// ProductCanSubscribeProducts 查询用户可订购的产品列表
func (APIClient) ProductCanSubscribeProducts(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsCanSubscribeProducts, data, http.MethodGet)
}

// ProductCancel 产品取消
func (APIClient) ProductCancel(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsCancel, data)
}

// ProductRefund 产品退费
func (APIClient) ProductRefund(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsRefund, data)
}

// ProductExpire 查询产品到期用户
func (APIClient) ProductExpire(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsExpire, data, http.MethodGet)
}

// ProductDisable 禁用产品
func (APIClient) ProductDisable(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsDisable, data)
}

// ProductEnable 启用产品
func (APIClient) ProductEnable(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(ProductsEnable, data)
}
