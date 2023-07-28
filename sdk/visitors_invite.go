package sdk

import "srunsoft-api-sdk/tools"

// 邀请码访客

// UserViewInvite 查询访客邀请码
func (APIClient) UserViewInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersViewInvite, data, "GET")
}

// UserInviteVisitors 邀请码访客开户
func (APIClient) UserInviteVisitors(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersInviteVisitor, data)
}

// UserUseInvite 使用访客邀请码
func (APIClient) UserUseInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersUseInvite, data)
}

// UserCreateInvite 创建邀请码
func (APIClient) UserCreateInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersInviteCreate, data)
}

// UserDisabledInvite 禁用邀请码
func (APIClient) UserDisabledInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(UsersInviteDisable, data)
}
