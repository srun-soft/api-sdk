package sdk

import (
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

type APIClient struct {
}

func ApiCall(urlPath string, data map[string]interface{}, method ...string) (tools.SrunResponse, error) {
	var (
		res tools.SrunResponse
		err error
	)
	v := tools.Map2Values(data)

	// 设置默认请求为"POST"
	httpMethod := "POST"
	if len(method) > 0 {
		httpMethod = method[0]
	}
	switch httpMethod {
	case http.MethodGet:
		res, err = tools.HandleGet(urlPath, v)
	case "POST":
		res, err = tools.HandlePost(urlPath, v)
	case http.MethodPut:
		res, err = tools.HandlePut(urlPath, v)
	case http.MethodDelete:
		res, err = tools.HandleDelete(urlPath, v)
	default:
		res, err = tools.HandlePost(urlPath, v)
	}
	if err != nil {
		configs.Log.WithField("API call failed", err).Error()
	}
	return res, err
}
