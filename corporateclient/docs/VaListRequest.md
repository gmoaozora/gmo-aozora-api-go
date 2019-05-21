# VaListRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VaContractAuthKey** | **string** | 振込入金口座API認証情報 半角英数字 銀行契約の方はNULLを設定 提携企業の方が、契約された顧客の一覧照会を依頼される場合は必須 提携企業以外の方が値を設定されている場合は「400 Bad Request」を返却  | [optional] [default to null]
**VaTypeCode** | **string** | 振込入金口座　種類コード 半角数字 ・1&#x3D;期限型 ・2&#x3D;継続型  | [optional] [default to null]
**DepositAmountExistCode** | **string** | 入金有無コード 半角数字 ・1&#x3D;入金あり ・2&#x3D;入金なし  | [optional] [default to null]
**VaHolderNameKana** | **string** | 振込入金口座名義カナ 半角文字 振込許容文字を指定可 ただし、一部の非許容文字は、許容文字に変換を行います  | [optional] [default to null]
**VaStatusCodeList** | [**[]VaStatusCode**](VaStatusCode.md) | 振込入金口座IDリスト 照会したい振込入金口座状態コードのリスト 上限3件まで設定可能  | [optional] [default to null]
**LatestDepositDateFrom** | **string** | 最終入金日From 半角文字 YYYY-MM-DD形式 入金有無コードが未設定もしくは、「1&#x3D;入金あり」が設定されている場合は指定可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）  | [optional] [default to null]
**LatestDepositDateTo** | **string** | 最終入金日To 半角文字 YYYY-MM-DD形式 入金有無コードが未設定もしくは、「1&#x3D;入金あり」が設定されている場合は指定可 それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 最終入金日Fromと最終入金日Toを指定する場合は、最終入金日From≦最終入金日Toとし、それ以外は「400 Bad Request」を返却  | [optional] [default to null]
**VaIssueDateFrom** | **string** | 振込入金口座発行日From 半角文字 YYYY-MM-DD形式  | [optional] [default to null]
**VaIssueDateTo** | **string** | 振込入金口座発行日To 半角文字 YYYY-MM-DD形式 振込入金口座発行Fromと振込入金口座発行Toを指定する場合は、振込入金口座発行From≦振込入金口座発行Toとし、それ以外は「400 Bad Request」を返却  | [optional] [default to null]
**ExpireDateFrom** | **string** | 有効期限From 半角文字 YYYY-MM-DD形式 振込入金口座 種類コードが未設定もしくは、「1＝期限型」が設定されている場合は指定可　それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却）  | [optional] [default to null]
**ExpireDateTo** | **string** | 有効期限To 半角文字 YYYY-MM-DD形式 振込入金口座 種類コードが未設定もしくは、「1＝期限型」が設定されている場合は指定可　それ以外はNULLを設定（値が設定されている場合は、「400 Bad Request」を返却） 有効期限Fromと有効期限Toを指定する場合は、有効期限From≦有効期限Toとし、それ以外は「400 Bad Request」を返却  | [optional] [default to null]
**RaId** | **string** | 入金口座ID 半角数字 口座一覧照会APIで取得できる口座IDを設定 科目コードが以下の口座IDのみ受け付けます ・01&#x3D;普通預金（有利息） ・02&#x3D;普通預金（決済用）  | [optional] [default to null]
**NextItemKey** | **string** | 次一覧キー 半角英数字  | [optional] [default to null]
**SortItemCode** | **string** | ソート項目コード 半角数字 ・1&#x3D;有効期限日時 ・2&#x3D;最終入金日 ・3&#x3D;発行日時 ・4&#x3D;最終入金金額 有効期限日時、最終入金日、最終入金金額など、レスポンスで返却されない場合がある項目をソート項目に指定した場合、ソートは効きません  | [optional] [default to null]
**SortOrderCode** | **string** | ソート順コード 半角数字 ソート項目コードのソート順を指定するコード値 ・1&#x3D;昇順 ・2&#x3D;降順  | [optional] [default to null]
**VaIdList** | [**[]VaId**](VaId.md) | 振込入金口座IDリスト 照会したい振込入金口座IDのリスト 上限500件まで設定可能  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


