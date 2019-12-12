# AuthorizationApi

All URIs are relative to *https://stg-api.gmo-aozora.com/ganb/api/auth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuthAuthorization**](AuthorizationApi.md#authorization) | **GET** /authorization | 認可(OAuth2.0)
[**OpenIDAuthorization**](AuthorizationApi.md#authorization) | **GET** /authorization | 認可(OpenID Connect)

# **authorization**
> func (g Ganb) OAuthAuthorization(sessionID string, scope string, redirectURI string) string\
> func (g Ganb) OpenIDAuthorization(sessionID string, scope string, redirectURI string) (string, error)

### 認可

クライアントがユーザーの認証・認可を得るためのエンドポイント

### Example
```go
package main

import (
    "errors"
    "fmt"
    "io"
    "os"
    
    "github.com/gmoaozora/gmo-aozora-api-go"
)

func nonceCheckAndDelete(nonce string) error {
    // implement your function
}

func nonceSave(nonce string) error {
    // implement your function
}

func main() {
    clientId := "clientId_example"
    clientSecret := "clientSecret_example"
    sessionID := "sessionID_example"
    scope := "personal:transfer personal:bulk-transfer"
    redirectURI := "redirectURI_example"

    ganb, err := ganblib.New(clientId, clientSecret, nonceSave, nonceCheckAndDelete)
    if err != nil {
        fmt.Println("failed to create ganb")
    }

    url := ganb.OAuthAuthorization(sessionID, scope, redirectURI)

    fmt.Println(url)

}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| クライアントID (当社が事前に発行する貴社向けのID)  minLength: 1 maxLength: 128  |
 **clientSecret** | **string**| クライアントシークレット（貴社認証用の項目）。（当社が事前に発行する貴社向けのシークレット）  minLength: 1 maxLength: 128  |
 **sessionID** | **string** | 貴社にて指定。conf.jsonで定めたSALT値と組み合わせてstateパラメータ(ハッシュ値)を作成するための項目 |
 **scope** | **string**| 要求されるアクセス権限を示すスコープID。 複数設定する場合は半角スペース区切りにて連結。 リフレッシュトークンを発行する場合は offline_access scope値が必要。  minLength: 1 maxLength: 256  |
 **redirectUri** | **string**| 貴社が指定する認可コードをリダイレクトするためのURI  minLength: 1 maxLength: 256  |
 **nonceSave** | **function** | 貴社にて実装。nonce値を保存するための関数。 |
 **nonceCheckAndDelete** | **function** | 貴社にて実装。返却されたnonce値を検証し、その後削除するための関数。 | 
 **nonce** | **string**| Client セッションと ID Token を紐づける文字列。設定された場合はそのままの値をID Tokenに含めて返却する。 リプレイアタック攻撃を防止するため、リクエスト毎にランダム値(十分な不規則性となる値)を設定し ID Tokenに含まれるnonce値が認可エンドポイントリクエスト時と同一であることを一度だけ検証するために利用する項目。  minLength: 1 maxLength: 128  | [optional]

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8
