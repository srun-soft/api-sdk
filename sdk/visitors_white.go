package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
	"net/http"
)

// 访客白名单

// VisitorCreateVisitorWhite 添加访客白名单
func (APIClient) VisitorCreateVisitorWhite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.VisitorsWhiteCreate, data)
}

// VisitorDeleteVisitorWhite 删除访客白名单
func (APIClient) VisitorDeleteVisitorWhite(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.VisitorsWhiteDelete, data, http.MethodDelete)
}
