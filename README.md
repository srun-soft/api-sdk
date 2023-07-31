# 北向接口SDK
#### `version v1.0.0`

go调用北向接口SDK

## Installation
To install the SDK package, use the following command:
```bash
go get -u github.com/srun-soft/api-sdk
```

## Usage
import the SDK package in your Go code:
```go
import (
	"github.com/srun-soft/api-sdk"
)
```

initialize the SDK package and create an API client:
```go
func main() {
	configs.Config = &configs.APIConfig{
        Scheme:      "http://", // 接口协议
        InterfaceIP: "127.0.0.1:8001", // 接口地址及端口
        AppId:       "srunsoft", // appid
        AppSecret:   "ba62f23e6212790052f387ee7af2943e4cfcece0", // appSecret
        LogDir:      "runtime/", // 日志目录
    }
	
	// ... 初始化其他配置 ...
	
	// 创建 API 客户端
	API := &sdk.APIClient{}
}
```

## Examples
You can find example usage in the `example` directory of this repository;
```go
    // 调用开户接口
    data := make(map[string]interface{})
    data["user_name"] = "yt"
    data["group_id"] = 1
    data["user_password"] = "12345678"
    data["products_id"] = 1
    API := &sdk.APIClient{}
    response, err := API.CreateUser(data)
    if err != nil {
        return
    }
    configs.Log.Info(response)
```

