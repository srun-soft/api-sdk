package version

import "net/http"

type Product interface {
	ProductView() Response
	ProductCreate() Response
	ProductDelete() Response
	ProductUpdate() Response
	ProductIndex() Response
	ProductNextBillingCycle() Response
	ProductTransferProduct() Response
	ProductReservationTransferProducts() Response
	ProductCancelReservationTransferProducts() Response
	ProductProductRecharge() Response
	ProductRecharge() Response
	ProductUseNumber() Response
	ProductSubscribe() Response
	ProductCanSubscribeProducts() Response
	ProductCancel() Response
	ProductRefund() Response
	ProductExpire() Response
	ProductDisableProduct() Response
	ProductEnableProduct() Response
}

func (ac *ApiConfig) ProductView() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = ProductView
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductCreate() Response {
	ac.UrlPath = ProductCreate
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductDelete() Response {
	ac.Method = http.MethodDelete
	ac.UrlPath = ProductDelete
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductUpdate() Response {
	ac.Method = http.MethodPut
	ac.UrlPath = ProductUpdate
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductIndex() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = ProductIndex
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductNextBillingCycle() Response {
	ac.UrlPath = ProductNextBillingCycle
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductTransferProduct() Response {
	ac.UrlPath = ProductTransferProduct
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductReservationTransferProducts() Response {
	ac.UrlPath = ProductReservationTransferProducts
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductCancelReservationTransferProducts() Response {
	ac.UrlPath = ProductCancelReservationTransferProducts
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductProductRecharge() Response {
	ac.UrlPath = ProductProductRecharge
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductRecharge() Response {
	ac.UrlPath = ProductRecharge
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductUseNumber() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = ProductUseNumber
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductSubscribe() Response {
	ac.UrlPath = ProductSubscribe
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductCanSubscribeProducts() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = ProductCanSubscribeProducts
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductCancel() Response {
	ac.UrlPath = ProductCancel
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductRefund() Response {
	ac.UrlPath = ProductRefund
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductExpire() Response {
	ac.Method = http.MethodGet
	ac.UrlPath = ProductExpire
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductDisableProduct() Response {
	ac.UrlPath = ProductDisableProduct
	return ac.ApiCall()
}

func (ac *ApiConfig) ProductEnableProduct() Response {
	ac.UrlPath = ProductEnableProduct
	return ac.ApiCall()
}
