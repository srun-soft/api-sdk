package version

import (
	"fmt"
	"github.com/srun-soft/api-sdk/tools/cache"
	"net/http"
)

type SDK interface {
	GetAccessToken() string
	ApiCall() Response
}

type ApiConfig struct {
	AppID       string
	AppSecret   string
	Scheme      string
	Host        string
	Version     int
	AccessToken string
	*Request
}

// GetAccessToken 获取令牌 /**
func (ac *ApiConfig) GetAccessToken() string {
	cacheComponent := cache.Component()
	token := cacheComponent.GetCache(KeyApiSdkAccessToken)
	if token == "" {
		if ac.Version == 1 {
			ac.UrlPath = AuthGetAccessToken
			ac.Method = http.MethodGet
		} else {
			ac.UrlPath = AuthGetAccessToken
			ac.Method = http.MethodPost
			ac.Params.Set("appId", ac.AppID)
			ac.Params.Set("appSecret", ac.AppSecret)
		}
		res := ac.ApiCall()
		if res.Code != 0 {
			fmt.Println("获取access_token 失败", res.Message)
			return ""
		}

		data := res.Data.(map[string]interface{})
		token = data["access_token"].(string)
		if lt, ok := data["lifetime"]; ok {
			cacheComponent.SetCache(KeyApiSdkAccessToken, token, lt.(float64))
		} else {
			cacheComponent.SetCache(KeyApiSdkAccessToken, token, float64(3600))
		}
	}
	return token
}

// ApiCall API调用
func (ac *ApiConfig) ApiCall() Response {
	client := newClient()
	if ac.Method == "" {
		ac.Method = http.MethodPost
	}
	res, err := client.DoRequest(ac)
	if err != nil {
		fmt.Println("API Call failed:", err)
	}
	return res
}
