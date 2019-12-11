# BulkTransferApi

All URIs are relative to *https://stg-api.gmo-aozora.com/ganb/api/corporation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkTransferCancelUsingPOST**](BulkTransferApi.md#BulkTransferCancelUsingPOST) | **Post** /bulktransfer/cancel | 総合振込取消依頼
[**BulkTransferFeeUsingPOST**](BulkTransferApi.md#BulkTransferFeeUsingPOST) | **Post** /bulktransfer/transferfee | 総合振込手数料事前照会
[**BulkTransferRequestResultUsingGET**](BulkTransferApi.md#BulkTransferRequestResultUsingGET) | **Get** /bulktransfer/request-result | 総合振込依頼結果照会
[**BulkTransferRequestUsingPOST**](BulkTransferApi.md#BulkTransferRequestUsingPOST) | **Post** /bulktransfer/request | 総合振込依頼
[**BulkTransferStatusUsingGET**](BulkTransferApi.md#BulkTransferStatusUsingGET) | **Get** /bulktransfer/status | 総合振込状況照会


# **BulkTransferCancelUsingPOST**
> func (a *BulkTransferApiService) BulkTransferCancelUsingPOST(ctx context.Context, body TransferCancelRequest, xAccessToken string) (TransferCancelResponse, *http.Response, error)

### 総合振込取消依頼

総合振込の取消を行うための依頼をします

### 詳細説明

##### 取消対象ステータス
* 申請中以降のステータスで取消が可能です
* 依頼中、作成中の状態は取消対象外です

#### 取消対象キー区分
* 取消対象の取引の内容について指定して下さい
* 取消対象キー区分と、取消対象の振込申請番号の状態がマッチしない場合は、「400 Bad Request」を返却します
* 総合振込の対象は3または4のみとなります
* 1、2は指定不可となります
* ビジネスID管理未利用の場合は、4を指定してください。その他は指定不可となります
* ビジネスID管理利用中かつ、申請者による申請中ステータスの「取下」を行いたい場合は、3を指定してください
* ビジネスID管理利用中かつ、承認可能者による予約中ステータスの「承認取消」を行いたい場合は、4を指定してください

#### 重複した依頼
* 同一の受付番号（振込申請番号）への重複した依頼はできません
* なお、前回の総合振込取消依頼が期限切れとなれば依頼は可能となります

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/corporateclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.BulkTransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    body := swagger.TransferCancelRequest{AccountId: "AccountId_example", CancelTargetKeyClass: "4", ApplyNo: "ApplyNo_example"}

    res, _, err := apiAccount.BulkTransferCancelUsingPOST(ctx, body, xAccessToken)
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

# **BulkTransferFeeUsingPOST**
> func (a *BulkTransferApiService) BulkTransferFeeUsingPOST(ctx context.Context, body BulkTransferRequest, xAccessToken string) (TransferFeeResponse, *http.Response, error)

### 総合振込手数料事前照会

総合振込・総合振込予約を行うための依頼内容の事前チェックと手数料を照会します

### 詳細説明
* 合計振込手数料および個別振込手数料は、振込実行時までに手数料の改定や消費税の変更等が行われた場合は、当APIで返却した手数料とは異なる手数料が適用されることがあります
* ポイントについては、算出から振込実行までの間に変動する可能性があるため、手数料算出時の計算対象から除外して返却されます

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/corporateclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.BulkTransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    bulkTransfer := swagger.BulkTransfer{ItemId: "1", TransferAmount: "100", BeneficiaryBankCode: "BeneficiaryBankCode_example", BeneficiaryBranchCode: "BeneficiaryBranchCode_example", AccountTypeCode: "1", AccountNumber: "AccountNumber_example", BeneficiaryName: "BeneficiaryName_example"}
    bulkTransfers := []swagger.BulkTransfer{bulkTransfer}

    body := swagger.BulkTransferRequest{AccountId: "AccountId_example", TransferDesignatedDate: "TransferDesignatedDate_example", TotalCount: "1", TotalAmount: "100", BulkTransfers: bulkTransfers}

    res, _, err := apiAccount.BulkTransferFeeUsingPOST(ctx, body, xAccessToken)
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
  **body** | [**BulkTransferRequest**](BulkTransferRequest.md)| HTTPリクエストボディ | 
  **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**TransferFeeResponse**](TransferFeeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkTransferRequestResultUsingGET**
> func (a *BulkTransferApiService) BulkTransferRequestResultUsingGET(ctx context.Context, accountId string, applyNo string, xAccessToken string) (TransferRequestResultResponse, *http.Response, error)

### 総合振込依頼結果照会

総合振込依頼、総合振込取消依頼の処理状態を照会します
* 総合振込取消依頼をした場合は、最後の依頼の結果コードが照会対象となります

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/corporateclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.BulkTransferApi

    ctx := context.Background()

    accountId := "accountId_example"
    applyNo := "applyNo_example"
    xAccessToken := "xAccessToken_example"

    res, _, err := apiAccount.BulkTransferRequestResultUsingGET(ctx, accountId, applyNo, xAccessToken)
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

# **BulkTransferRequestUsingPOST**
> func (a *BulkTransferApiService) BulkTransferRequestUsingPOST(ctx context.Context, body BulkTransferRequest, xAccessToken string) (BulkTransferRequestResponse, *http.Response, error)

### 総合振込依頼

総合振込・総合振込予約を行うための依頼をします

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/gmoaozora/gmo-aozora-api-go/corporateclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.BulkTransferApi

    ctx := context.Background()

    xAccessToken := "xAccessToken_example"

    bulkTransfer := swagger.BulkTransfer{ItemId: "1", TransferAmount: "100", BeneficiaryBankCode: "BeneficiaryBankCode_example", BeneficiaryBranchCode: "BeneficiaryBranchCode_example", AccountTypeCode: "1", AccountNumber: "AccountNumber_example", BeneficiaryName: "BeneficiaryName_example"}
    bulkTransfers := []swagger.BulkTransfer{bulkTransfer}

    body := swagger.BulkTransferRequest{AccountId: "AccountId_example", TransferDesignatedDate: "TransferDesignatedDate_example", TotalCount: "1", TotalAmount: "100", BulkTransfers: bulkTransfers}

    res, _, err := apiAccount.BulkTransferRequestUsingPOST(ctx, body, xAccessToken)
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
 **body** | [**BulkTransferRequest**](BulkTransferRequest.md)| HTTPリクエストボディ | 
 **xAccessToken** | **string**| アクセストークン  minLength: 1 maxLength: 128  | 

### Return type

[**BulkTransferRequestResponse**](BulkTransferRequestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkTransferStatusUsingGET**
> func (a *BulkTransferApiService) BulkTransferStatusUsingGET(ctx context.Context, accountId string, queryKeyClass string, xAccessToken string, localVarOptionals *BulkTransferStatusUsingGETOpts) (BulkTransferStatusResponse, *http.Response, error)

### 総合振込状況照会

仕向の総合振込状況および履歴を照会します

### 詳細説明

#### 取得上限件数
* 総合振込照会明細情報: 500件（明細情報取得フラグが「False:取得しない」の場合）
* 総合振込明細情報: 500件（明細情報取得フラグが「True：取得する」の場合）
* 1回のリクエストで照会できる明細は、総合振込照会明細情報または総合振込明細情報のいずれか最大500件です
* 取得できる明細が存在しない場合は「200：OK」とし明細を返却しません
* 取得できる明細数が500に満たないときは取得できる明細のみを返却します

#### ページング
* 2ページ目以降を照会する際は、初回と同じリクエスト内容に、初回レスポンスの次明細キーまたは総合振込明細情報取得対象キーを追加してリクエストしてください

#### ソート順
##### 明細情報取得フラグが「False：取得しない」のとき

振込照会対象期間区分の指定により下記となります

  * 1：振込申請受付日　第1ソート：振込申請日昇順　第２ソート：振込申請番号昇順
  * 2：振込指定日　　　第1ソート：振込指定日昇順　第２ソート：振込申請番号昇順

※振込照会対象期間区分の指定がない場合は、1：振込申請受付日と同じとします

##### 明細情報取得フラグが、「True：取得する」のとき
* 明細番号の昇順

#### 対象期間

日本語名     | &#9312; | &#9313; | &#9314; | &#9315;
:----|:----:|:----:|:----:|:----:
対象期間From | × | ○ | × | ○
対象期間To   | × | × | ○ | ○
* &#9312;の場合: 当日分の明細を返却
* &#9313;の場合: 対象期間From ～ 当日までの明細を返却
* &#9314;の場合: 取引初回 ～ 対象期間Toまでの明細を返却
* &#9315;の場合: 対象期間From ～ 対象期間Toまでの明細を返却

※現在契約中の総合振込サービスのみ照会可能となります

#### 照会対象ステータス
* 申請中以降のステータスで照会が可能となります
* 依頼中、作成中の状態は照会対象外です

#### 照会対象となる明細
* 総合振込・およびその予約の履歴と状況が照会対象となります（APIからの依頼結果のみではありません）

### Example
```go
package main

import (
    "context"
    "fmt"

    "github.com/antihax/optional"
    "github.com/gmoaozora/gmo-aozora-api-go/corporateclient"
    "github.com/k0kubun/pp"
)

func main() {
    configuration := swagger.NewConfiguration()
    apiClinent := swagger.NewAPIClient(configuration)
    apiAccount := apiClinent.BulkTransferApi

    ctx := context.Background()

    accountId := "accountId_example"
    queryKeyClass := "queryKeyClass_example"
    xAccessToken := "xAccessToken_example"

    opts := swagger.BulkTransferStatusUsingGETOpts{DateFrom: optional.NewString("DateFrom_example")}

    res, _, err := apiAccount.BulkTransferStatusUsingGET(ctx, accountId, queryKeyClass, xAccessToken, &opts)
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
 **optional** | ***BulkTransferStatusUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BulkTransferStatusUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detailInfoNecessity** | **optional.Bool**| 明細情報取得フラグ 総合振込明細情報の取得要否 照会対象キー区分が、1のときは「True：取得する」を指定可 それ以外で「True：取得する」が設定されている場合は、「400 Bad Request」を返却 True：取得する、False:取得しない　省略/NULLは　false扱い  | 
 **bulktransferItemKey** | **optional.String**| 総合振込明細情報取得対象キー 半角数字 明細情報取得フラグが、「True：取得する」のとき指定可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 総合振込明細情報を取得するときに取得を開始する番号 明細情報取得フラグが、「True：取得する」のときの省略/NULLは1扱い 1500明細を取得する場合、は以下のように設定 1電文目：1 2電文目：501 3電文目：1001  minLength: 1 maxLength: 6  | 
 **applyNo** | **optional.String**| 受付番号（振込申請番号） 半角数字 照会対象の番号を設定 照会対象キー区分が、1のときは必須 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）  minLength: 16 maxLength: 16  | 
 **dateFrom** | **optional.String**| 対象期間From 半角文字 YYYY-MM-DD形式 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）              minLength: 10 maxLength: 10  | 
 **dateTo** | **optional.String**| 対象期間To 半角文字 YYYY-MM-DD形式 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 対象期間Fromと対象期間Toを指定する場合は、対象期間From≦対象期間Toとし、それ以外は「400 Bad Request」を返却  minLength: 10 maxLength: 10  | 
 **nextItemKey** | **optional.String**| 次明細キー 半角数字 照会対象キー区分が、2のときは入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）              minLength: 1 maxLength: 24  | 
 **requestTransferStatus** | [**optional.Interface of []string**](string.md)| 照会対象ステータス  半角数字  2:申請中、3:差戻、4:取下げ、5:期限切れ、8:承認取消/予約取消、  11:予約中、12:手続中、13:リトライ中、  20:手続済、30:不能・組戻あり、40:手続不成立   照会対象キー区分が、2のときは設定可  それ以外は設定しません（値が設定されている場合は、「400 Bad Request」を返却）  配列のため、複数設定した場合は対象のステータスをOR条件で検索します  省略した場合は全てを設定したものとみなします  minLength: 1 maxLength: 3  | 
 **requestTransferClass** | **optional.String**| 振込照会対象取得区分 半角数字 1：ALL、2：振込申請のみ、3：振込受付情報のみ 照会対象キー区分が2かつ、ビジネスID管理利用中の場合のときのみ入力可 それ以外はNULLを設定(値が設定されている場合は、「400 Bad Request」を返却) 照会対象キー区分が、2のときに指定しない場合は1と扱います minLength: 1 maxLength: 1  | 
 **requestTransferTerm** | **optional.String**| 振込照会対象期間区分 半角数字 対象期間Fromと対象期間Toで指定する日付の区分 1：振込申請受付日　2：振込指定日 照会対象キー区分が2のときのみ入力可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 照会対象キー区分が、2のときに指定しない場合は1と扱います  minLength: 1 maxLength: 1  | 

### Return type

[**BulkTransferStatusResponse**](BulkTransferStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
