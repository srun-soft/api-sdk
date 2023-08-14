package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
)

// 产品转移

// ProductNextBillingCycle 产品转移(下月生效)
func (APIClient) ProductNextBillingCycle(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsNextBillingCycle, data)
}

// ProductTransferProduct 产品转移(立即生效)
func (APIClient) ProductTransferProduct(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsTransferProduct, data)
}

// ProductReservationTransferProducts 产品转移(预约转移)
func (APIClient) ProductReservationTransferProducts(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsReservationTransferProducts, data)
}

// ProductCancelReservationTransferProducts 产品转移(取消预约转移)
func (APIClient) ProductCancelReservationTransferProducts(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsCancelReservationTransferProducts, data)
}
