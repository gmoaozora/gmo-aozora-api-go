# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 口座を識別するID  | [default to null]
**BranchCode** | **string** | 支店コード 半角数字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**BranchName** | **string** | 支店名 全角文字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**AccountTypeCode** | **string** | 科目コード 半角数字 ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用） ・11&#x3D;円定期預金 ・51&#x3D;外貨普通預金 ・81&#x3D;証券コネクト口座  | [default to null]
**AccountTypeName** | **string** | 科目名 全角文字 科目コードの名称  | [default to null]
**AccountNumber** | **string** | 口座番号 半角数字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**PrimaryAccountCode** | **string** | 代表口座コード 半角数字 ・1&#x3D;代表口座 ・2&#x3D;追加口座 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**PrimaryAccountCodeName** | **string** | 代表口座コード名 全角文字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**AccountName** | **string** | 口座名義 全角文字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**AccountNameKana** | **string** | 口座名義カナ 半角文字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**CurrencyCode** | **string** | 通貨コード 半角文字 ISO4217に準拠した通貨コード  | [default to null]
**CurrencyName** | **string** | 通貨名 全角文字 ISO4217に準拠した通貨コードの当行での名称  | [default to null]
**TransferLimitAmount** | **string** | 振込限度額 半角数字 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


