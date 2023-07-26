package sdk

import (
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
)

type SDK interface {
	GetToken() (string, error)
}

type Users interface {
	CreateUser(map[string]interface{}) (tools.SrunResponse, error)
	UpdateUser(map[string]interface{}) (tools.SrunResponse, error)
	DeleteUser(map[string]interface{}) (tools.SrunResponse, error)
}

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
	case "GET":
		res, err = tools.HandleGet(urlPath, v)
	case "POST":
		res, err = tools.HandlePost(urlPath, v)
	case "PUT":
		res, err = tools.HandlePut(urlPath, v)
	case "DELETE":
		res, err = tools.HandleDelete(urlPath, v)
	default:
		res, err = tools.HandlePost(urlPath, v)
	}
	if err != nil {
		configs.Log.WithField("API call failed", err).Error()
	}
	return res, err
}
