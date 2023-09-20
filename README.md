# 北向接口 SDK
## 版本：v2.0.0

这是一个用于调用北向接口的 Go SDK。通过声明全局 SDK 指针结构体，你可以轻松地调用北向接口的各种功能。当使用本 SDK 时，你可以轻松地适配 v1 和 v2 版本的北向接口。要切换版本，只需指定 version 即可。这个 SDK 的设计旨在使版本切换变得简单而直观，帮助你无缝集成北向接口到你的应用中。接口常量的命名方式遵循北向接口的地址驼峰命名规则，例如：

- 接口文档地址：`/api/v2/auth/check-modify-password`
- SDK 对应常量：`AuthCheckModifyPassword`

## 安装

要安装 SDK 包，使用以下命令：

```bash
go get -u github.com/srun-soft/api-apisdk
```

## 使用
在你的 Go 代码中导入 SDK 包:
```go
import (
    "github.com/srun-soft/api-apisdk"
)
```

初始化 SDK 包并创建一个 API 客户端:
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
        // 版本可在初始化时统一指定，也可在调用接口时灵活修改版本
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

## 示例
你可以在本仓库的 example 目录中找到示例用法。例如：
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

