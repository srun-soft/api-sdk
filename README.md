# 北向接口SDK
#### `version v2.0.0`

go调用北向接口SDK
声明全局SDK指针结构体后,调用对应的接口即可。
接口常量按照北向接口`Api fox`的接口地址驼峰命名 eg: 

**接口文档地址** : /api/v2/auth/check-modify-password

**sdk对应常量** : AuthCheckModifyPassword
## Installation
To install the SDK package, use the following command:
```bash
go get -u github.com/srun-soft/api-apisdk
```

## Usage
import the SDK package in your Go code:
```go
import (
	"github.com/srun-soft/api-apisdk"
)
```

initialize the SDK package and create an API client:
```go
// 声明全局sdk变量
var SDK *apisdk.ApiConfig

func main() {
    SDK = &apisdk.ApiConfig{
        AppID       : "srunsoft",
        AppSecret   : "ba62f23e6212790052f387ee7af2943e4cfcece0",
        Scheme      : "https",
        Host        : "192.168.0.191",
        Port        : 8001,
        Version     : 2,
		// 使用redis缓存access_token
        Cache: redis.NewClient(&redis.Options{
            Addr        : fmt.Sprintf("%s:%s", "192.168.0.191", "16384"),
            Password    : "srun_3000@redis",
            DB          : 0,
            DialTimeout : 10 * time.Second,
            ReadTimeout : 30 * time.Second,
            WriteTimeout: 30 * time.Second,
            PoolSize    : 100,
            PoolTimeout : 30 * time.Second,
        }),
    }
}
```

## Examples
You can find example usage in the `example` directory of this repository;
```go
// 调用查询接口
// 接口参数使用 url.Values{}, 赋值给 sdk.Params
// 随后调用响应的方法, sdk.{MethodName}
// 响应为 response{code:int,message:string,data:interface{}}
// 详细调用示例查看example
func version2() {
    v := url.Values{}
    v.Add("user_name", "yuantong")
    SDK.Params = v
    res := SDK.UserView()
    fmt.Println(res.Code)
    fmt.Println(res.Message)
    if res.Code == 0 {
    m := res.Data.(map[string]interface{})
    fmt.Printf("%v", m["user_name"])
}

```

