# Balance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 口座を識別するID  | [default to null]
**AccountTypeCode** | **string** | 科目コード 半角数字 ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用） ・11&#x3D;円定期預金 ・51&#x3D;外貨普通預金 ・81&#x3D;証券コネクト口座  | [default to null]
**AccountTypeName** | **string** | 科目名 全角文字 科目コードの名称  | [default to null]
**Balance** | **string** | 現在残高 半角数字　マイナス含む 基準日時における現在残高 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用） ・11&#x3D;円定期預金 ・81&#x3D;証券コネクト口座  | [optional] [default to null]
**BaseDate** | **string** | 基準日 半角文字 残高および引出可能額を照会した基準日を示します YYYY-MM-DD形式  | [default to null]
**BaseTime** | **string** | 基準時刻 半角文字 残高および引出可能額を照会した基準時刻を示します HH:MM:SS+09:00形式  | [default to null]
**WithdrawableAmount** | **string** | 支払可能残高 半角数字　マイナス含む 応答時点での引出可能額を示します 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**PreviousDayBalance** | **string** | 前日残高 半角数字　マイナス含む 日付が変わった直後は、銀行の締め処理が終わるまでは前々日の残高となります 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**PreviousMonthBalance** | **string** | 前月末残高 半角数字　マイナス含む 月が変わった直後は、銀行の締め処理が終わるまでは前々月の残高となります  科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**CurrencyCode** | **string** | 通貨コード 半角文字 ISO4217に準拠した通貨コード  | [default to null]
**CurrencyName** | **string** | 通貨名 全角文字 ISO4217に準拠した通貨コードの当行での名称  | [default to null]
**FcyTotalBalance** | **string** | 外貨残高 半角数字　少数点および小数部最大3桁　マイナス含む 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51&#x3D;外貨普通預金  | [optional] [default to null]
**Ttb** | **string** | 外貨レート 半角数字　少数点および小数部最大3桁　マイナス含む 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51&#x3D;外貨普通預金  | [optional] [default to null]
**BaseRateDate** | **string** | 外貨レート基準日 半角文字 外貨レートの基準日を示します YYYY-MM-DD形式 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51&#x3D;外貨普通預金  | [optional] [default to null]
**BaseRateTime** | **string** | 外貨レート基準時刻 半角文字 外貨レートの基準時刻を示します HH:MM:SS+09:00形式 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51&#x3D;外貨普通預金  | [optional] [default to null]
**YenEquivalent** | **string** | 外貨円換算額 半角数字　マイナス含む 外貨残高を円に換算した額 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51&#x3D;外貨普通預金  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


