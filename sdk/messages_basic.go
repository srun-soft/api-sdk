package sdk

import "srunsoft-api-sdk/tools"

// 通知公告相关

// MessageNewMessage 查询新版本消息
func (APIClient) MessageNewMessage(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(MessagesNew, data, "GET")
}

// Message 查询通知公告
func (APIClient) Message(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(Messages, data, "GET")
}

// MessageNotice 发送通知消息接口
func (APIClient) MessageNotice(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(MessagesNotice, data)
}
