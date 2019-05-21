# SpAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 つかいわけ口座を識別するID  | [default to null]
**SpAccountTypeCode** | **string** | つかいわけ口座種別コード 半角数字 ・1&#x3D;親口座 ・2&#x3D;子口座  | [default to null]
**SpAccountTypeCodeName** | **string** | つかいわけ口座種別コード名称 全角文字 つかいわけ口座種別コードの名称  | [default to null]
**SpAccountName** | **string** | つかいわけ口座名 全半角英数記号文字  | [default to null]
**SpAccountBranchCode** | **string** | 支店コード 半角数字 つかいわけ口座種別コードが「2&#x3D;子口座」の場合のみ設定 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**SpAccountBranchName** | **string** | 支店名 全角文字 つかいわけ口座種別コードが「2&#x3D;子口座」の場合のみ設定 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**SpAccountNumber** | **string** | 口座番号 半角数字 つかいわけ口座種別コードが「2&#x3D;子口座」の場合のみ設定 該当しない場合は項目自体を設定しません  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


