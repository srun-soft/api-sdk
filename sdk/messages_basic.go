package sdk

import (
	"net/http"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

// 通知公告相关

// MessageNewMessage 查询新版本消息
func (APIClient) MessageNewMessage(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.MessagesNew, data, http.MethodGet)
}

// Message 查询通知公告
func (APIClient) Message(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.Messages, data, http.MethodGet)
}

// MessageNotice 发送通知消息接口
func (APIClient) MessageNotice(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.MessagesNotice, data)
}
