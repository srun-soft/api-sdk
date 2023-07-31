package sdk

import (
	"fmt"
	"net/http"
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

	// 设置默认请求为http.MethodPost
	httpMethod := http.MethodPost
	if len(method) > 0 {
		httpMethod = method[0]
	}
	switch httpMethod {
	case http.MethodGet:
		res, err = tools.HandleGet(urlPath, v)
	case http.MethodPost:
		res, err = tools.HandlePost(urlPath, v)
	case http.MethodPut:
		res, err = tools.HandlePut(urlPath, v)
	case http.MethodDelete:
		res, err = tools.HandleDelete(urlPath, v)
	default:
		res, err = tools.HandlePost(urlPath, v)
	}
	if err != nil {
		fmt.Println("API call failed", err)
		return res, err
	}
	return res, err
}
