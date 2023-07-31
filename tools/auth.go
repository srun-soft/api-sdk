package tools

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools/cache"
)

const (
	KeyApiSdkAccessToken = "key:api:sdk:access_token"
	GetAccessToken       = "/api/v2/auth/get-access-token"
)

// GetToken 获取 access_token
func GetToken(client HTTPClient, cache cache.Cache) (string, error) {
	var (
		token string
		err   error
		res   SrunResponse
	)

	if cache == nil {
		return "", errors.New(fmt.Sprintf("cache connection failed:%v", err))
	}

	token = cache.GetCache(KeyApiSdkAccessToken)
	if token == "" {
		v := url.Values{}
		v.Set("appId", configs.Config.AppId)
		v.Set("appSecret", configs.Config.AppSecret)

		res, err = client.DoRequest(http.MethodPost, GetAccessToken, v)
		if err != nil {
			fmt.Println("获取access_token失败", err)
			return "", err
		}
		if res.Code != 0 {
			fmt.Println("获取access_token失败", res.Message)
			return "", fmt.Errorf("获取access_token失败: %s", res.Message)
		}

		data := res.Data.(map[string]interface{})
		token = data["access_token"].(string)
		cache.SetCache(KeyApiSdkAccessToken, token, data["lifetime"].(float64))
	}

	return token, nil
}
