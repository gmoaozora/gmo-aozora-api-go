# BulkTransferResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 口座を識別するID  | [optional] [default to null]
**RemitterName** | **string** | 振込依頼人名 半角文字  | [optional] [default to null]
**TransferDesignatedDate** | **string** | 振込指定日 半角文字 YYYY-MM-DD形式  | [optional] [default to null]
**TransferDataName** | **string** | 振込データ名 全半角文字 作成した総合振込のデータを区別するためのメモ  | [optional] [default to null]
**TotalCount** | **string** | 合計件数 半角数字  | [optional] [default to null]
**TotalAmount** | **string** | 合計金額 半角数字  | [optional] [default to null]
**BulkTransferInfos** | [**[]BulkTransferInfo**](BulkTransferInfo.md) | 総合振込明細情報 総合振込明細のリスト 明細情報取得フラグが「True：取得する」、かつ明細情報取得結果フラグが「True：取得可」のときのみ設定します それ以外は項目自体を設定しません  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


