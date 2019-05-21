# SpAccountTransferResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptDatetime** | **string** | 更新受付日時 更新を受付した銀行側の日時  | [optional] [default to null]
**DepositSpAccountId** | **string** | つかいわけ口座_入金口座ID 半角数字 口座を識別するID　入金先となる口座情報  | [default to null]
**DebitSpAccountId** | **string** | つかいわけ口座_出金口座ID 半角数字 口座を識別するID　引き落とし元となる口座情報  | [default to null]
**CurrencyCode** | **string** | 通貨コード 半角文字 ISO4217準拠した通貨コード  | [optional] [default to null]
**CurrencyName** | **string** | 通貨名 全角文字 ISO4217準拠した通貨コードの当行での名称  | [optional] [default to null]
**PaymentAmount** | **string** | 支払金額 数値のみでカンマ、マイナス不可 円貨(JPY)の場合１以上であること 外貨の場合当該通貨の補助通貨単位以上であること  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


