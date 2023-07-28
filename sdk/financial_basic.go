package sdk

import "srunsoft-api-sdk/tools"

// 缴费相关

// FinancialRechargeWallet 电子钱包缴费
func (APIClient) FinancialRechargeWallet(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsRechargeWallet, data)
}

// FinancialPayType 缴费方式查询
func (APIClient) FinancialPayType(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsPayType, data, "GET")
}

// FinancialCreatePayment 添加缴费方式
func (APIClient) FinancialCreatePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsCreatePayment, data)
}

// FinancialDeletePayment 删除缴费方式
func (APIClient) FinancialDeletePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsDeletePayment, data, "DELETE")
}

// FinancialUpdatePayment 修改缴费方式
func (APIClient) FinancialUpdatePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsUpdatePayment, data)
}

// FinancialExtraPay 交开户费/附加费用
func (APIClient) FinancialExtraPay(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsExtraPay, data)
}

// FinancialTransfer 电子钱包内余额转账
func (APIClient) FinancialTransfer(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsTransfer, data)
}

// FinancialRechargeCards 充值卡查询接口
func (APIClient) FinancialRechargeCards(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsRechargeCards, data, "GET")
}
