# SpAccountBalance

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | 口座ID 半角英数字 つかいわけ口座を識別するID  | [default to null]
**OdBalance** | **string** | 円普通預金残高 半角数字　マイナス含む 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**TdTotalBalance** | **string** | 円定期預金（総額） 半角数字　マイナス含む つかいわけ口座に紐付いた、円定期預金の総残高 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**FodTotalBalanceYenEquivalent** | **string** | 外貨普通預金（総評価額） 半角数字　マイナス含む つかいわけ口座に紐付いた、外貨普通預金の総残高を円に換算した額 該当しない場合は項目自体を設定しません  | [optional] [default to null]
**SpAccountFcyBalances** | [**[]SpAccountFcyBalance**](SpAccountFcyBalance.md) | つかわけ口座外貨残高情報リスト 該当する情報が無い場合は、空のリストを返却します  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


