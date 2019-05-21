# VaTransaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaId** | **string** | 振込入金口座ID 半角数字 振込入金口座を識別するID  | [default to null]
**TransactionDate** | **string** | 取引日 半角文字 YYYY-MM-DD形式  | [default to null]
**ValueDate** | **string** | 起算日 半角文字 YYYY-MM-DD形式  | [default to null]
**VaBranchCode** | **string** | 支店コード 半角数字  | [default to null]
**VaBranchNameKana** | **string** | 支店名カナ 半角文字  | [default to null]
**VaAccountNumber** | **string** | 口座番号 半角数字  | [default to null]
**VaAccountNameKana** | **string** | 口座名義カナ 半角文字  | [default to null]
**DepositAmount** | **string** | 入金金額 半角数字  | [default to null]
**RemitterNameKana** | **string** | 振込依頼人名カナ 半角文字  | [default to null]
**PaymentBankName** | **string** | 仕向金融機関名カナ 半角文字  | [default to null]
**PaymentBranchName** | **string** | 仕向支店名カナ 半角文字  | [default to null]
**PartnerName** | **string** | サービス企業名 全角文字 振込入金口座契約サービス企業名  | [default to null]
**Remarks** | **string** | 摘要内容 全半角文字 該当データがない場合は項目自体を設定しません  | [optional] [default to null]
**ItemKey** | **string** | 明細キー 半角数字 口座ID毎に設定される明細キー（明細データtimestamp（μs）） | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


