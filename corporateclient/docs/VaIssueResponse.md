# VaIssueResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaTypeCode** | **string** | 振込入金口座　種類コード 半角数字 ・1&#x3D;期限型 ・2&#x3D;継続型  | [default to null]
**VaTypeName** | **string** | 振込入金口座　種類名 全角文字 振込入金口座　種類コードの名称  | [default to null]
**ExpireDateTime** | **string** | 有効期限日時 半角文字 YYYY-MM-DDTHH:MM:SS+09:00形式 振込入金口座種別コードが「2&#x3D;継続型」の場合は、項目自体を設定しません  | [optional] [default to null]
**VaHolderNameKana** | **string** | 振込入金口座名義カナ 半角文字  | [default to null]
**VaList** | [**[]Va**](Va.md) | 振込入金口座リスト | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


