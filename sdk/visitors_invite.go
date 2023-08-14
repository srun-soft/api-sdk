package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 邀请码访客

// UserViewInvite 查询访客邀请码
func (APIClient) UserViewInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersViewInvite, data, http.MethodGet)
}

// UserInviteVisitors 邀请码访客开户
func (APIClient) UserInviteVisitors(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersInviteVisitor, data)
}

// UserUseInvite 使用访客邀请码
func (APIClient) UserUseInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersUseInvite, data)
}

// UserCreateInvite 创建邀请码
func (APIClient) UserCreateInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersInviteCreate, data)
}

// UserDisabledInvite 禁用邀请码
func (APIClient) UserDisabledInvite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.UsersInviteDisable, data)
}
