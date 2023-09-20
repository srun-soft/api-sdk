package apisdk

import "net/http"

type Financial interface {
	FinancialRechargeWallet() Response
	FinancialPayType() Response
	FinancialCreatePayment() Response
	FinancialDeletePayment() Response
	FinancialUpdatePayment() Response
	FinancialExtraPay() Response
	FinancialTransfer() Response
	FinancialRechargeCards() Response
	FinancialPaymentDataSync() Response
	CheckoutListDetail() Response
	FinancialPaymentRecords() Response
	FinancialPaymentOverview() Response
	FinancialRefund() Response
	FinancialPackageRechargeRecord() Response
}

func (ac *ApiConfig) FinancialRechargeWallet() Response {
	ac.UrlPath = FinancialRechargeWallet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialPayType() Response {
	ac.UrlPath = FinancialPayType
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialCreatePayment() Response {
	ac.UrlPath = FinancialCreatePayment
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialDeletePayment() Response {
	ac.UrlPath = FinancialDeletePayment
	ac.Method = http.MethodDelete
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialUpdatePayment() Response {
	ac.UrlPath = FinancialUpdatePayment
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialExtraPay() Response {
	ac.UrlPath = FinancialExtraPay
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialTransfer() Response {
	ac.UrlPath = FinancialTransfer
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialRechargeCards() Response {
	ac.UrlPath = FinancialRechargeCards
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialPaymentDataSync() Response {
	ac.UrlPath = FinancialPaymentDataSync
	return ac.ApiCall()
}

func (ac *ApiConfig) CheckoutListDetail() Response {
	ac.UrlPath = CheckoutListDetail
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialPaymentRecords() Response {
	ac.UrlPath = FinancialPaymentRecords
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialPaymentOverview() Response {
	ac.UrlPath = FinancialPaymentOverview
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialRefund() Response {
	ac.UrlPath = FinancialRefund
	ac.Method = http.MethodGet
	return ac.ApiCall()
}

func (ac *ApiConfig) FinancialPackageRechargeRecord() Response {
	ac.UrlPath = FinancialPackageRechargeRecord
	ac.Method = http.MethodGet
	return ac.ApiCall()
}
