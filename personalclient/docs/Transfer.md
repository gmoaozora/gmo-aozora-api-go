# Transfer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | 明細番号 半角数字 重複/0はエラー　1～99の範囲で１からの連番とします 1件のみの場合は省略可（項目自体の設定が不要です）  | [optional] [default to null]
**TransferAmount** | **string** | 振込金額 半角数字 1以上、整数のみ  | [default to null]
**EdiInfo** | **string** | EDI情報 半角文字 振込許容文字を指定可 ・個人事業主のみ対象の項目のため、個人は対象外となります ・個人の場合は設定しても無効となり、処理に利用しません  | [optional] [default to null]
**BeneficiaryBankCode** | **string** | 被仕向金融機関番号 半角数字  | [default to null]
**BeneficiaryBankName** | **string** | 被仕向金融機関名カナ 半角文字 参考値、処理に利用しません  | [optional] [default to null]
**BeneficiaryBranchCode** | **string** | 被仕向支店番号 半角数字  | [default to null]
**BeneficiaryBranchName** | **string** | 被仕向支店名カナ 半角文字 参考値、処理に利用しません  | [optional] [default to null]
**AccountTypeCode** | **string** | 科目コード（預金種別コード） 半角数字 1：普通、2：当座、4：貯蓄、9：その他  | [default to null]
**AccountNumber** | **string** | 口座番号 半角数字 7桁未満の番号は右詰で、前ゼロで埋めること  | [default to null]
**BeneficiaryName** | **string** | 受取人名 半角文字 振込許容文字を指定可 ただし、一部の非許容文字は、許容文字に変換を行います  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


