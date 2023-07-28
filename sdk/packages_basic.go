package sdk

import "srunsoft-api-sdk/tools"

// 套餐相关

// PackageUsersPackages 查询已订购的产品/套餐【统计时长/流量包可用总和】
func (APIClient) PackageUsersPackages(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(PackagesUsersPackages, data, "GET")
}

// PackageBuy 购买套餐
func (APIClient) PackageBuy(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(PackagesBuy, data)
}

// PackageBuySuper 购买套餐 - 赠送
func (APIClient) PackageBuySuper(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(PackagesBuySuper, data)
}

// Packages 查询可购买的套餐
func (APIClient) Packages(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(Packages, data, "GET")
}

// PackageBuys 购买套餐二
func (APIClient) PackageBuys(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(PackagesBuys, data)
}

// PackageBatch 购买套餐批处理
func (APIClient) PackageBatch(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(PackagesBatch, data)
}
