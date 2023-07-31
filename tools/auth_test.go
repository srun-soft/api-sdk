package tools

import (
	"net/url"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools/cache"
	"testing"
)

// mockHTTPClient 实现 HTTPClient 接口，用于在测试中模拟 HTTP 请求和响应
type mockHTTPClient struct {
	response SrunResponse
	err      error
}

// DoRequest 模拟 HTTP 请求并返回预先定义的响应
func (c *mockHTTPClient) DoRequest(method, urlPath string, data url.Values) (SrunResponse, error) {
	return c.response, c.err
}

func TestGetToken(t *testing.T) {
	configs.Config = &configs.APIConfig{
		Scheme:      "http",
		InterfaceIP: "127.0.0.1",
		AppId:       "srunsoft",
		AppSecret:   "ba62f23e6212790052f387ee7af2943e4cfcece0",
		LogDir:      "runtime",
	}
	// 模拟一个HTTP客户端，返回一个成功的响应
	mockResponse := SrunResponse{
		Code:    0,
		Message: "Success",
		Data: map[string]interface{}{
			"access_token": "mock_access_token",
			"lifetime":     float64(3600), // 设置token生存时间为1小时
		},
	}
	mockClient := &mockHTTPClient{
		response: mockResponse,
	}

	// 调用 GetToken 函数进行测试，传入 mockClient 的指针
	token, err := GetToken(mockClient, cache.Component())

	// 断言是否有错误发生
	if err != nil {
		t.Errorf("GetToken returned an unexpected error: %v", err)
	}

	// 断言是否得到了正确的token
	expectedToken := "mock_access_token"
	if token != expectedToken {
		t.Errorf("GetToken returned an unexpected token. Expected: %s, Got: %s", expectedToken, token)
	}
}
