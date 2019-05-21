# TransferError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | エラーコード 半角英数文字  | [default to null]
**ErrorMessage** | **string** | エラーメッセージ 全半角英数記号文字  | [default to null]
**ErrorDetails** | [**[]ErrorDetail**](ErrorDetail.md) | エラー詳細情報 該当する情報が無い場合は空のリストを返却  | [optional] [default to null]
**TransferErrorDetails** | [**[]TransferErrorDetail**](TransferErrorDetail.md) | 振込明細エラー情報 個別明細がエラーの場合のみ応答 該当する情報が無い場合は空のリストを返却  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


