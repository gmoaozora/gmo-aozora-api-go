# WebhooksApi

All URIs are relative to *https://stg-api.gmo-aozora.com/ganb/api/webhooks/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsUsing**](WebhooksApi.md#AccountsUsing) | **Post** /subscribe | 通知配信制御
[**AccountsUsingGET**](WebhooksApi.md#AccountsUsingGET) | **Get** /unsentlist/va-deposit-transaction | 振込入金口座未送信明細取得


# **AccountsUsing**
> func (a *WebhooksApiService) AccountsUsing(ctx context.Context, authorization string, body SubscribeRequestBody) (*http.Response, error)

### 通知配信制御

指定したイベント通知に対応するイベント通知（WebHook）の配信開始、配信停止をコントロールします

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/webhookclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.WebhooksApi

    ctx := context.Background()

    authorization := "authorization_example"
    eventType := swagger.EventType{EventType: "va-deposit-transaction"}
    eventTypes := []swagger.EventType{eventType}

    body := swagger.SubscribeRequestBody{SubscribeStatus: "1", EventTypes: eventTypes}

    res, err := apiAccount.AccountsUsing(ctx, authorization, body)
    if err != nil {
        fmt.Println(err)
        return
    }
    pp.Print(res)
}
```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **authorization** | **string**| 認証情報 銀行システムが配信先システムに発行した、クライアントIDとクライアントシーレットを:(コロン)で連結し、Base64エンコードした値を設定 minLength: ‐ maxLength: ‐  | 
 **body** | [**SubscribeRequestBody**](SubscribeRequestBody.md)| HTTPリクエストボディ | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsUsingGET**
> func (a *WebhooksApiService) AccountsUsingGET(ctx context.Context, authorization string) (VaDepositTransactionUnsentResponse, *http.Response, error)

### 振込入金口座未送信明細取得

配信停止状態となっている場合、本APIを利用することで未送信または送信エラーとなっている、振込入金口座の入金明細を一括で取得することができます
* 通常、未送信または送信エラーとなっている明細は配信再開後に通知されますが、本APIで取得された明細は配信済みとなるため、配信再開後には通知されません
* 未送信または送信エラーとなっている明細が無い場合は404 Not Foundを返却します

※法人口座および個人事業主口座のみ対象となり、個人口座は対象外となります

### 詳細説明

#### 取得上限件数
* 500件
* 取得できる明細数が500に満たないときは取得できる明細のみを返却します
* 取得できた明細数が500件の場合、まだ取得できる明細が残っている可能性がありますので、「404：Not Found」が返却されるまで、リクエストを繰り返してください。

### Example
```go
package main

import (
    "context"

    "github.com/gmoaozora/gmo-aozora-api-go/webhookclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.WebhooksApi

    ctx := context.Background()

    authorization := "authorization_example"

    _, httpRes, err := apiAccount.AccountsUsingGET(ctx, authorization)
    if err != nil {
        pp.Print(httpRes)
        pp.Print(err)
        return
    }
    pp.Print(httpRes)
}
```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **authorization** | **string**| 認証情報 銀行システムが配信先システムに発行した、クライアントIDとクライアントシーレットを:(コロン)で連結し、Base64エンコードした値を設定 minLength: ‐ maxLength: ‐  | 

### Return type

[**VaDepositTransactionUnsentResponse**](vaDepositTransactionUnsentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
