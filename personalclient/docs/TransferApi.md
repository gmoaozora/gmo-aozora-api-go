# TransferApi

All URIs are relative to *https://stg-api.gmo-aozora.com/ganb/api/personal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpAccountTransferUsingPOST**](TransferApi.md#SpAccountTransferUsingPOST) | **Post** /transfer/spaccounts-transfer | つかいわけ口座間振替
[**TransferCancelUsingPOST**](TransferApi.md#TransferCancelUsingPOST) | **Post** /transfer/cancel | 振込取消依頼
[**TransferFeeUsingPOST**](TransferApi.md#TransferFeeUsingPOST) | **Post** /transfer/transferfee | 振込手数料事前照会
[**TransferRequestResultUsingGET**](TransferApi.md#TransferRequestResultUsingGET) | **Get** /transfer/request-result | 振込依頼結果照会
[**TransferRequestUsingPOST**](TransferApi.md#TransferRequestUsingPOST) | **Post** /transfer/request | 振込依頼
[**TransferStatusUsingGET**](TransferApi.md#TransferStatusUsingGET) | **Get** /transfer/status | 振込状況照会


# **SpAccountTransferUsingPOST**
> func (a *TransferApiService) SpAccountTransferUsingPOST(ctx context.Context, body SpAccountTransferRequest, xAccessToken string) (SpAccountTransferResponse, *http.Response, error)

### つかいわけ口座間振替

つかいわけ口座間の振替を実行します 振替の実行は即時となります つかいわけ口座間の明細移動は当APIの対象外です 

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    body := swagger.SpAccountTransferRequest{DepositSpAccountId: "DepositSpAccountId_example", DebitSpAccountId: "DebitSpAccountId_example", PaymentAmount: "100"}

    res, _, err := apiAccount.SpAccountTransferUsingPOST(ctx, body, xAccessToken)
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
 **body** | [**SpAccountTransferRequest**](SpAccountTransferRequest.md)| HTTPリクエストボディ | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**SpAccountTransferResponse**](SpAccountTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferCancelUsingPOST**
> func (a *TransferApiService) TransferCancelUsingPOST(ctx context.Context, body TransferCancelRequest, xAccessToken string) (TransferCancelResponse, *http.Response, error)

### 振込取消依頼

振込・振込予約の取消を行うための依頼をします

### 詳細説明

#### 取消対象ステータス
* 申請中以降のステータスで取消が可能です
* 依頼中、作成中の状態は取消対象外です

#### 取消対象キー区分
* 取消対象の取引の内容について指定して下さい
* 取消対象キー区分と、取消対象の振込申請番号の状態がマッチしない場合は、「400 Bad Request」を返却します
* 振込・振替/一括振込の対象は1または2のみとなります
* 3、4は指定不可となります
* ビジネスID管理未利用の場合は、2を指定してください。その他は指定不可となります
* ビジネスID管理利用中かつ、申請者による申請中ステータスの「取下」を行いたい場合は、1を指定してください
* ビジネスID管理利用中かつ、承認可能者による予約中ステータスの「承認取消」を行いたい場合は、2を指定してください

#### 重複した依頼
* 同一の受付番号（振込申請番号）への重複した依頼はできません
* なお、前回の振込取消依頼が期限切れとなれば依頼は可能となります

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    body := swagger.TransferCancelRequest{AccountId: "AccountId_example", CancelTargetKeyClass: "CancelTargetKeyClass_example", ApplyNo: "ApplyNo_example"}

    res, _, err := apiAccount.TransferCancelUsingPOST(ctx, body, xAccessToken)
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
 **body** | [**TransferCancelRequest**](TransferCancelRequest.md)| HTTPリクエストボディ | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**TransferCancelResponse**](TransferCancelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferFeeUsingPOST**
> func (a *TransferApiService) TransferFeeUsingPOST(ctx context.Context, body TransferRequest, xAccessToken string) (TransferFeeResponse, *http.Response, error)

### 振込手数料事前照会

振込・振込予約を行うための依頼内容の事前チェックと手数料を照会します

### 詳細説明

* 最大99件まで登録可能
* 1件の場合通常の振込扱いとなり、2件以上で一括振込扱いとなります
* 合計振込手数料および個別振込手数料は、振込実行時までに手数料の改定や消費税の変更等が行われた場合は、当APIで返却した手数料とは異なる手数料が適用されることがあります
* 振込無料回数とポイントについては、算出から振込実行までの間に変動する可能性があるため、手数料算出時の計算対象から除外して返却されます

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    transfer := swagger.Transfer{TransferAmount: "100", BeneficiaryBankCode: "BeneficiaryBankCode_example", BeneficiaryBranchCode: "BeneficiaryBranchCode_example", AccountTypeCode: "1", AccountNumber: "AccountNumber_example", BeneficiaryName: "BeneficiaryName_example"}
    transfers := []swagger.Transfer{transfer}

    body := swagger.TransferRequest{AccountId: "AccountId_example", TransferDesignatedDate: "TransferDesignatedDate_example", Transfers: transfers}

    res, _, err := apiAccount.TransferFeeUsingPOST(ctx, body, xAccessToken)
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
 **body** | [**TransferRequest**](TransferRequest.md)| HTTPリクエストボディ | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**TransferFeeResponse**](TransferFeeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferRequestResultUsingGET**
> func (a *TransferApiService) TransferRequestResultUsingGET(ctx context.Context, accountId string, applyNo string, xAccessToken string) (TransferRequestResultResponse, *http.Response, error)

### 振込依頼結果照会

振込依頼、振込取消依頼の処理状態を照会します
* 振込取消依頼をした場合は、最後の依頼の結果コードが照会対象となります 

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    accountId := "accountId_example"
    applyNo := "applyNo_example"
    xAccessToken := "xAccessToken_example"

    res, _, err := apiAccount.TransferRequestResultUsingGET(ctx, accountId, applyNo, xAccessToken)
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
 **accountId** | **string**| 口座ID 半角数字 口座を識別するID  minLength: 12 maxLength: 29  | 
 **applyNo** | **string**| 受付番号（振込申請番号） 半角数字 すべての振込・総合振込で採番される、照会の基本単位となる番号  minLength: 16 maxLength: 16  | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**TransferRequestResultResponse**](TransferRequestResultResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferRequestUsingPOST**
> func (a *TransferApiService) TransferRequestUsingPOST(ctx context.Context, body TransferRequest, xAccessToken string) (TransferRequestResponse, *http.Response, error)

### 振込依頼

振込・振込予約を行うための依頼をします

### 詳細説明

* 最大99件まで登録可能
* 1件の場合通常の振込扱いとなり、2件以上で一括振込扱いとなります

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    accountId := "accountId_example"
    applyNo := "applyNo_example"
    xAccessToken := "xAccessToken_example"

    res, _, err := apiAccount.TransferRequestUsingPOST(ctx, accountId, applyNo, xAccessToken)
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
 **body** | [**TransferRequest**](TransferRequest.md)| HTTPリクエストボディ | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**TransferRequestResponse**](TransferRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferStatusUsingGET**
> func (a *TransferApiService) TransferStatusUsingGET(ctx context.Context, accountId string, queryKeyClass string, xAccessToken string, opts *TransferStatusUsingGETOpts) (TransferStatusResponse, *http.Response, error)

### 振込状況照会

仕向の振込状況および履歴を照会します

### 詳細説明

#### 取得上限件数
* 500件
* 取得できる明細数が500に満たないときは取得できる明細のみを返却します
* 取得できる明細が存在しない場合は「200：OK」とし明細を返却しません

#### ページング
* 2ページ目以降を照会する際は、初回と同じリクエスト内容に、初回レスポンスの次明細キーを追加してリクエストしてください

#### ソート順

* 振込照会対象期間区分の指定により下記となります
  * 1：振込申請受付日　第1ソート：振込申請日昇順　第２ソート：振込申請番号昇順
  * 2：振込指定日　　　第1ソート：振込指定日昇順　第２ソート：振込申請番号昇順

※振込照会対象期間区分の指定がない場合は、1：振込申請受付日と同じとします

#### 対象期間

日本語名     | &#9312; | &#9313; | &#9314; | &#9315;
:----|:----:|:----:|:----:|:----:
対象期間From | × | ○ | × | ○
対象期間To   | × | × | ○ | ○
* &#9312;の場合: 当日分の明細を返却
* &#9313;の場合: 対象期間From ～ 当日までの明細を返却
* &#9314;の場合: 取引初回 ～ 対象期間Toまでの明細を返却
* &#9315;の場合: 対象期間From ～ 対象期間Toまでの明細を返却

#### 照会対象ステータス
* 申請中以降のステータスで照会が可能となります
* 依頼中、作成中の状態は照会対象外です

#### 照会対象となる明細
* 振込・振替・およびその予約の履歴と状況が照会対象となります（APIからの依頼結果のみではありません）
* 定額自動振込契約によって自動登録された振込は照会対象となります
* 定額自動振込契約の申請状態と状況は対象外となります

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/antihax/optional"
    "github.com/gmoaozora/gmo-aozora-api-go/personalclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.TransferApi

    ctx := context.Background()

    accountId := "accountId_example"
    queryKeyClass := "queryKeyClass_example"
    xAccessToken := "xAccessToken_example"

    opts := swagger.TransferStatusUsingGETOpts{DateFrom: optional.NewString("DateFrom_example")}

    res, _, err := apiAccount.TransferStatusUsingGET(ctx, accountId, queryKeyClass, xAccessToken, &opts)
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
 **accountId** | **string**| 口座ID 半角数字 口座を識別するID  minLength: 12 maxLength: 29  | 
 **queryKeyClass** | **string**| 照会対象キー区分 半角数字 照会対象のキー 1：振込申請照会対象指定、2：振込一括照会対象指定  minLength: 1 maxLength: 1  | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 
 **opts** | ***TransferStatusUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransferStatusUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applyNo** | **optional.String**| 受付番号（振込申請番号） 半角数字 照会対象の番号を設定 照会対象キー区分が、1のときは必須 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）  minLength: 16 maxLength: 16  | 
 **dateFrom** | **optional.String**| 対象期間From 半角文字 YYYY-MM-DD形式 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）  minLength: 10 maxLength: 10  | 
 **dateTo** | **optional.String**| 対象期間To 半角文字 YYYY-MM-DD形式 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 対象期間Fromと対象期間Toを指定する場合は、対象期間From≦対象期間Toとし、それ以外は「400 Bad Request」を返却  minLength: 10 maxLength: 10  | 
 **nextItemKey** | **optional.String**| 次明細キー 半角数字 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）              minLength: 1 maxLength: 24  | 
 **requestTransferStatus** | [**optional.Interface of []string**](string.md)| 照会対象ステータス  半角数字  2:申請中、3:差戻、4:取下げ、5:期限切れ、8:承認取消/予約取消、  11:予約中、12:手続中、13:リトライ中、  20:手続済、22:資金返却、24:組戻手続中、25:組戻済、26:組戻不成立、  40:手続不成立  照会対象キー区分が、2のときは設定可  それ以外は設定しません（値が設定されている場合は、「400 Bad Request」を返却）  配列のため、複数設定した場合は対象のステータスをOR条件で検索します  省略した場合は全てを設定したものとみなします  minLength: 1 maxLength: 3  | 
 **requestTransferClass** | **optional.String**| 振込照会対象取得区分 半角数字 1：ALL、2：振込申請のみ、3：振込受付情報のみ NULLを設定 （値が設定されている場合は、「400 Bad Request」を返却） 照会対象キー区分が、2のときに指定しない場合は1と扱います minLength: 1 maxLength: 1  | 
 **requestTransferTerm** | **optional.String**| 振込照会対象期間区分 半角数字 対象期間Fromと対象期間Toで指定する日付の区分 1：振込申請受付日　2：振込指定日 照会対象キー区分が2のときのみ入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 照会対象キー区分が、2のときに指定しない場合は1と扱います  minLength: 1 maxLength: 1  | 

### Return type

[**TransferStatusResponse**](TransferStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
