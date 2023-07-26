package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"srunsoft-api-sdk/configs"
)

// SrunResponse 北向接口响应结构
type SrunResponse struct {
	Code    int
	Message string
	Version string
	Data    map[string]interface{}
}

func HandlePost(urlPath string, data url.Values) (SrunResponse, error) {
	var (
		resp *http.Response
		sr   SrunResponse
		err  error
	)

	resp, err = http.PostForm(fmt.Sprintf("%s%s%s", configs.Config.Scheme, configs.Config.InterfaceIP, urlPath), data)
	if err != nil {
		configs.Log.Error(err)
		return sr, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			configs.Log.Error(err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		return sr, err
	}

	return sr, nil
}
