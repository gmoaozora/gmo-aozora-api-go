# VAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaId** | **string** | 振込入金口座ID 半角数字 振込入金口座を識別するID  | [default to null]
**VaBranchCode** | **string** | 支店コード 半角数字  | [default to null]
**VaBranchName** | **string** | 支店名 全角文字  | [default to null]
**VaBranchNameKana** | **string** | 支店名カナ 半角文字  | [default to null]
**VaAccountNumber** | **string** | 口座番号 半角数字  | [default to null]
**VaHolderNameKana** | **string** | 口座名義カナ 半角文字  | [default to null]
**VaTypeCode** | **string** | 振込入金口座　種類コード 半角数字 ・1&#x3D;期限型 ・2&#x3D;継続型  | [default to null]
**VaTypeName** | **string** | 振込入金口座　種類名 全角文字 振込入金口座　種類コードの名称  | [default to null]
**VaStatusCode** | **string** | 振込入金口座　状態コード 半角数字 ・1&#x3D;利用可能 ・2&#x3D;停止中 ・3&#x3D;削除済  | [default to null]
**VaStatusName** | **string** | 振込入金口座　状態名 全角文字 振込入金口座　状態コードの名称  | [default to null]
**VaStatusChangePossibleCode** | **string** | 振込入金口座　状態変更可能コード 半角数字 現在の口座状態から変更可能な口座状態を表したコード値 ・1&#x3D;停止・削除 ・2&#x3D;再開・削除 ・3&#x3D;該当なし  | [default to null]
**VaIssueDateTime** | **string** | 発行日時 半角文字 YYYY-MM-DDTHH:MM:SS+09:00形式  | [default to null]
**VaChangeStatusDateTime** | **string** | 状態変更日時 半角文字 YYYY-MM-DDTHH:MM:SS+09:00形式 状態が変更されている場合のみ設定されます 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**VaCloseDateTime** | **string** | 削除日時 半角文字 YYYY-MM-DDTHH:MM:SS+09:00形式 状態コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・3&#x3D;削除済  | [optional] [default to null]
**ExpireDateTime** | **string** | 有効期限日時 半角文字 YYYY-MM-DDTHH:MM:SS+09:00形式 種類コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・2&#x3D;期限型  | [optional] [default to null]
**LatestDepositDate** | **string** | 最終入金日 半角文字 YYYY-MM-DD形式 入金がない場合は項目自体を設定しません  | [optional] [default to null]
**LatestDepositAmount** | **string** | 最終入金金額 半角数字 入金がない場合は項目自体を設定しません  | [optional] [default to null]
**LatestDepositCount** | **string** | 最終入金日入金回数 半角数字 入金がない場合は項目自体を設定しません  | [optional] [default to null]
**RaId** | **string** | 入金口座ID 半角数字 入金先の口座を識別するID  | [default to null]
**RaBranchCode** | **string** | 入金口座　支店コード 半角数字  | [default to null]
**RaBranchName** | **string** | 入金口座　支店名 全角文字  | [default to null]
**RaAccountNumber** | **string** | 入金口座　口座番号 半角数字  | [default to null]
**RaHolderName** | **string** | 入金口座名義 全角文字  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


