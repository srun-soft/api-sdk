package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 运营商相关

// UserSearchMobilePhone 查询用户订购的产品是否绑定运营商账号
func (APIClient) UserSearchMobilePhone(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersSearchMobilePhone, data, http.MethodGet)
}

// ProductCheckOperators 验证运营上账号密码
func (APIClient) ProductCheckOperators(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsCheckOperators, data)
}

// ProductOperators 绑定/解绑运营商账号
func (APIClient) ProductOperators(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.ProductsOperators, data)
}
