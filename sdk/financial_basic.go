package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 缴费相关

// FinancialRechargeWallet 电子钱包缴费
func (APIClient) FinancialRechargeWallet(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsRechargeWallet, data)
}

// FinancialPayType 缴费方式查询
func (APIClient) FinancialPayType(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsPayType, data, http.MethodGet)
}

// FinancialCreatePayment 添加缴费方式
func (APIClient) FinancialCreatePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsCreatePayment, data)
}

// FinancialDeletePayment 删除缴费方式
func (APIClient) FinancialDeletePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsDeletePayment, data, http.MethodDelete)
}

// FinancialUpdatePayment 修改缴费方式
func (APIClient) FinancialUpdatePayment(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsUpdatePayment, data)
}

// FinancialExtraPay 交开户费/附加费用
func (APIClient) FinancialExtraPay(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsExtraPay, data)
}

// FinancialTransfer 电子钱包内余额转账
func (APIClient) FinancialTransfer(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsTransfer, data)
}

// FinancialRechargeCards 充值卡查询接口
func (APIClient) FinancialRechargeCards(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsRechargeCards, data, http.MethodGet)
}

// FinancialPaymentDataSync 缴费数据同步
func (APIClient) FinancialPaymentDataSync(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.FinancialsPaymentDataSync, data)
}
