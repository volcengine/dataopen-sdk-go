# openapi 使用说明

## Client 参数说明

| 字段       | 类型           | 默认值                                | 含义                                                |
| ---------- | -------------- | ------------------------------------- | --------------------------------------------------- |
| app_id     | string         | 无                                    | 应用的唯一标识符                                    |
| app_secret | string         | 无                                    | 用于应用的安全认证的密钥                            |
| url        | string or null | "https://analytics.volcengineapi.com" | 服务器的 URL 地址                                   |
| env        | string or null | "dataopen"                            | 环境设置，可选值为 "dataopen" 或 "dataopen_staging" |
| expiration | string or null | "1800"                                | 过期时间，单位是秒                                  |

## client.Request 参数说明

| 字段        | 类型                   | 默认值 | 含义                                            |
| ----------- | ---------------------- | ------ | ----------------------------------------------- |
| service_url | string                 | 无     | 请求的服务 URL 地址                             |
| method      | string                 | 无     | 请求的 HTTP 方法，例如 "GET", "POST" 等         |
| headers     | map[string]string      | {}     | 请求头，包含的信息如认证凭据，内容类型等        |
| params      | map[string]interface{} | {}     | URL 参数，用于 GET 请求                         |
| body        | map[string]interface{} | {}     | 请求体，通常在 POST 或 PUT 请求中包含发送的数据 |

## 获取与安装

go get -u github.com/volcengine/dataopen-sdk-go

## 举例

### 1、Get 方法

```go

import (
  "github.com/volcengine/dataopen-openapi-sdk/client"
)

app_id := ""
app_secret := ""

client := Client(app_id, app_secret, "", "", "")

headers := make(map[string]string)

params := make(map[string]ParamsValueType)
params["app"] = "46"
params["page_size"] = "1"
params["page"] = "1"

body := make(map[string]interface{})

res, err := client.Request("/xxx/openapi/v1/open/flight-list", "GET", headers, params, body)
if err != nil {
  t.Fatal(err)
}

fmt.Println("Output: ", res)
```

### 2、Post 方法

```go
import (
  "github.com/volcengine/dataopen-openapi-sdk/client"
)

app_id := ""
app_secret := ""

client := Client(app_id, app_secret, "", "", "")

headers := make(map[string]string)

params := make(map[string]ParamsValueType)

body := make(map[string]interface{})
body["uid_list"] = []string{"1111111110000"}

res, err := client.Request(
  "/xxx/openapi/v1/open/flight/version/6290880/add-test-user",
  "POST",
  headers,
  params,
  body,
)
if err != nil {
  t.Fatal(err)
}

fmt.Println("Output: ", res)
```
