package apisdk

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HTTPClient 定义HTTP请求接口
type HTTPClient interface {
	DoRequest(method, urlPath string, data url.Values) (Response, error)
}

type HttpClient struct {
	client *http.Client
}

func newClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 10 * time.Second,
		}}
}

// DoRequest 发起请求
func (c *HttpClient) DoRequest(ac *ApiConfig) (Response, error) {
	var (
		res Response
		err error
	)
	if ac.Version == 1 {
		ac.UrlPath = strings.ReplaceAll(ac.UrlPath, "v2", "v1")
	}
	reqURL := fmt.Sprintf("%s://%s:%d%s", ac.Scheme, ac.Host, ac.Port, ac.UrlPath)
	if ac.Params != nil {
		if ac.UrlPath != AuthGetAccessToken && ac.UrlPath != strings.ReplaceAll(AuthGetAccessToken, "v2", "v1") {
			ac.Params.Set("access_token", ac.GetAccessToken())
		}
		if ac.Method == http.MethodGet {
			reqURL = fmt.Sprintf("%s?%s", reqURL, ac.Params.Encode())
		}
	}

	req, err := http.NewRequest(ac.Method, reqURL, nil)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if ac.Method != http.MethodGet && ac.Params != nil {
		req.Body = io.NopCloser(strings.NewReader(ac.Params.Encode()))
		req.Header.Set("Content-Length", strconv.Itoa(len(ac.Params.Encode())))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println("client do request failed:", err)
		return res, nil
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			_ = fmt.Errorf("body close failed:%v", err)
		}
	}(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("json decode to response failed:", err)
		return res, err
	}

	return res, nil
}
