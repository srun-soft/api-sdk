package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 套餐相关

// PackageUsersPackages 查询已订购的产品/套餐【统计时长/流量包可用总和】
func (APIClient) PackageUsersPackages(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.PackagesUsersPackages, data, http.MethodGet)
}

// PackageBuy 购买套餐
func (APIClient) PackageBuy(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.PackagesBuy, data)
}

// PackageBuySuper 购买套餐 - 赠送
func (APIClient) PackageBuySuper(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.PackagesBuySuper, data)
}

// Packages 查询可购买的套餐
func (APIClient) Packages(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.Packages, data, http.MethodGet)
}

// PackageBuys 购买套餐二
func (APIClient) PackageBuys(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.PackagesBuys, data)
}

// PackageBatch 购买套餐批处理
func (APIClient) PackageBatch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.PackagesBatch, data)
}
