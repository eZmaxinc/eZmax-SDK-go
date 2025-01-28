# CreditcardmerchantGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjCreditcardmerchant** | [**[]CreditcardmerchantListElement**](CreditcardmerchantListElement.md) |  | 

## Methods

### NewCreditcardmerchantGetListV1ResponseMPayload

`func NewCreditcardmerchantGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjCreditcardmerchant []CreditcardmerchantListElement, ) *CreditcardmerchantGetListV1ResponseMPayload`

NewCreditcardmerchantGetListV1ResponseMPayload instantiates a new CreditcardmerchantGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardmerchantGetListV1ResponseMPayloadWithDefaults

`func NewCreditcardmerchantGetListV1ResponseMPayloadWithDefaults() *CreditcardmerchantGetListV1ResponseMPayload`

NewCreditcardmerchantGetListV1ResponseMPayloadWithDefaults instantiates a new CreditcardmerchantGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *CreditcardmerchantGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *CreditcardmerchantGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjCreditcardmerchant

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetAObjCreditcardmerchant() []CreditcardmerchantListElement`

GetAObjCreditcardmerchant returns the AObjCreditcardmerchant field if non-nil, zero value otherwise.

### GetAObjCreditcardmerchantOk

`func (o *CreditcardmerchantGetListV1ResponseMPayload) GetAObjCreditcardmerchantOk() (*[]CreditcardmerchantListElement, bool)`

GetAObjCreditcardmerchantOk returns a tuple with the AObjCreditcardmerchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCreditcardmerchant

`func (o *CreditcardmerchantGetListV1ResponseMPayload) SetAObjCreditcardmerchant(v []CreditcardmerchantListElement)`

SetAObjCreditcardmerchant sets AObjCreditcardmerchant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


