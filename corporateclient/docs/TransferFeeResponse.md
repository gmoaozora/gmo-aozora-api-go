# TransferFeeResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 口座を識別するID  | [default to null]
**BaseDate** | **string** | 基準日 振込の手数料を照会した基準日を示します YYYY-MM-DD形式  | [default to null]
**BaseTime** | **string** | 基準時刻 振込の手数料を照会した基準時刻を示します HH:MM:SS+09:00形式  | [default to null]
**TotalFee** | **string** | 合計振込手数料 半角数字 手数料の合計額を表示  | [default to null]
**TransferFeeDetails** | [**[]TransferFeeDetail**](TransferFeeDetail.md) | 振込手数料明細情報 個別明細を設定  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


