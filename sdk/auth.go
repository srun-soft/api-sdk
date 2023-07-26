package sdk

import (
	"net/url"
	"srunsoft-api-sdk/configs"
	"srunsoft-api-sdk/tools"
	"time"
)

func (c APIClient) GetToken() (string, error) {
	var (
		token string
		err   error
		res   tools.SrunResponse
	)
	cache := configs.Cache
	if cache == nil {
		configs.Log.WithField("Cache", "连接失败").Error()
	}
	if token, err = cache.Get(cache.Context(), KeyApiSdkAccessToken).Result(); err != nil && token != "" {
		return token, err
	}
	v := url.Values{}
	v.Set("appId", configs.Config.AppId)
	v.Set("appSecret", configs.Config.AppSecret)

	res, err = tools.HandlePost(GetAccessToken, v)
	if err != nil {
		configs.Log.WithField("获取access_token失败", err).Error()
	}
	if res.Code != 0 {
		configs.Log.WithField("获取access_token失败", res.Message).Error()
	}
	token = res.Data["access_token"].(string)
	err = cache.Set(cache.Context(), KeyApiSdkAccessToken, token, time.Duration(res.Data["lifetime"].(float64))).Err()
	if err != nil {
		configs.Log.WithField("缓存access_token失败", err).Error()
		return token, err
	}
	return token, err
}
