# TransferDetailResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeneficiaryBankNameKanji** | **string** | 被仕向金融機関名漢字 全角文字 被仕向金融機関番号に該当する金融機関名漢字 該当する情報が無い場合は項目自体を設定しません  | [optional] [default to null]
**BeneficiaryBranchNameKanji** | **string** | 被仕向支店名漢字 全角文字 被仕向支店番号に該当する支店名漢字 該当する情報が無い場合は項目自体を設定しません  | [optional] [default to null]
**UsedPoint** | **string** | 利用ポイント 半角数字 実際に利用したポイント 該当する情報が無い場合は項目自体を設定しません  | [optional] [default to null]
**IsFeeFreeUsed** | **bool** | 振込無料回数利用有無 振込無料回数の実際の利用有無 総合振込では無料回数は利用できないため、常にfalse  | [optional] [default to null]
**TransferFee** | **string** | 手数料 半角数字 個別明細単位の手数料 該当する情報が無い場合は項目自体を設定しません  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


