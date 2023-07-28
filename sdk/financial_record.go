package sdk

import "srunsoft-api-sdk/tools"

// 财务清单

// CheckoutlistDetail 结算清单
func (APIClient) CheckoutlistDetail(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(CheckoutListDetail, data, "GET")
}

// FinancialPaymentRecords 缴费记录
func (APIClient) FinancialPaymentRecords(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsPaymentRecords, data, "GET")
}

// FinancialPaymentOverview 财务缴费对账
func (APIClient) FinancialPaymentOverview(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsPaymentOverview, data, "GET")
}

// FinancialRefund 退费记录
func (APIClient) FinancialRefund(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsRefund, data, "GET")
}

// FinancialPackageRechargeRecord 套餐充值流量包记录
func (APIClient) FinancialPackageRechargeRecord(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(FinancialsPackageRechargeRecord, data, "GET")
}
