# TokenApi

All URIs are relative to *https://stg-api.gmo-aozora.com/ganb/api/auth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuthGetToken**](TokenApi.md#tokenUsingPOST) | **POST** /token | アクセストークン発行(OAuth2.0)
[**OpenIDGetToken**](TokenApi.md#tokenUsingPOST) | **POST** /token | アクセストークン発行(OpenID Content)


# **tokenUsingPOST**
> func (g Ganb) OAuthGetToken(redirectURI string, code string, authMethod string) (Token, error)
> func (g Ganb) OpenIDGetToken(redirectURI string, code string, authMethod string) (Token, error)

### アクセストークン発行

認可エンドポイントで取得した認可コードを用いたアクセストークンの取得及びリフレッシュトークンを用いたアクセストークンの更 新を行うためのエンドポイント

### Example

#### get new token
```go
package main

import (
    "fmt"
    
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

    ganb, err := ganblib.New(clientId, clientSecret, nonceSave, nonceCheckAndDelete)
    if err != nil {
        fmt.Println("failed to create ganb")
    }

    redirectURI := "redirectURI_example"
    code := "code_example"
    authMethod := "basic"

    token, err := ganb.OAuthGetToken(redirectURI, code, authMethod)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

#### refresh token
```go
package main

import (
    "fmt"
    
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

    ganb, err := ganblib.New(clientId, clientSecret, nonceSave, nonceCheckAndDelete)
    if err != nil {
        fmt.Println("failed to create ganb")
    }

    refreshToken := "refreshToken_example"
    
    new_token, err := ganb.RefreshTokens(refreshToken)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string**| クライアントID (当社が事前に発行する貴社向けのID)  minLength: 1 maxLength: 128  |
 **clientSecret** | **string**| クライアントシークレット（貴社認証用の項目）。（当社が事前に発行する貴社向けのシークレット）  minLength: 1 maxLength: 128  |
 **redirectUri** | **string**| 貴社が指定する認可コードをリダイレクトするためのURI  minLength: 1 maxLength: 256  |
 **code** | **string** | 新規発行時のみ必須 認可エンドポイントにて当社から返却した認可コード minLength: 1 maxLength: 128 |
 **authMethod** | **string** | 事前に登録したクライアント認証方式。ベーシック認証: 「basic」、クライアントシークレット認証: 「post」 |
 **nonceSave** | **function** | 貴社にて実装。nonce値を保存するための関数。 |
 **nonceCheckAndDelete** | **function** | 貴社にて実装。返却されたnonce値を検証し、その後削除するための関数。 | 
 **nonce** | **string**| Client セッションと ID Token を紐づける文字列。設定された場合はそのままの値をID Tokenに含めて返却する。 リプレイアタック攻撃を防止するため、リクエスト毎にランダム値(十分な不規則性となる値)を設定し ID Tokenに含まれるnonce値が認可エンドポイントリクエスト時と同一であることを一度だけ検証するために利用する項目。  minLength: 1 maxLength: 128  | [optional]
 **refreshToken** | **string** | 新しいアクセストークンを取得する際に使用するリフレッシュトークン。 |

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8
