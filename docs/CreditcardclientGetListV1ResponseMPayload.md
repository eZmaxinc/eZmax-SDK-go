# CreditcardclientGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjCreditcardclient** | [**[]CreditcardclientListElement**](CreditcardclientListElement.md) |  | 

## Methods

### NewCreditcardclientGetListV1ResponseMPayload

`func NewCreditcardclientGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjCreditcardclient []CreditcardclientListElement, ) *CreditcardclientGetListV1ResponseMPayload`

NewCreditcardclientGetListV1ResponseMPayload instantiates a new CreditcardclientGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditcardclientGetListV1ResponseMPayloadWithDefaults

`func NewCreditcardclientGetListV1ResponseMPayloadWithDefaults() *CreditcardclientGetListV1ResponseMPayload`

NewCreditcardclientGetListV1ResponseMPayloadWithDefaults instantiates a new CreditcardclientGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *CreditcardclientGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *CreditcardclientGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *CreditcardclientGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *CreditcardclientGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *CreditcardclientGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *CreditcardclientGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjCreditcardclient

`func (o *CreditcardclientGetListV1ResponseMPayload) GetAObjCreditcardclient() []CreditcardclientListElement`

GetAObjCreditcardclient returns the AObjCreditcardclient field if non-nil, zero value otherwise.

### GetAObjCreditcardclientOk

`func (o *CreditcardclientGetListV1ResponseMPayload) GetAObjCreditcardclientOk() (*[]CreditcardclientListElement, bool)`

GetAObjCreditcardclientOk returns a tuple with the AObjCreditcardclient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjCreditcardclient

`func (o *CreditcardclientGetListV1ResponseMPayload) SetAObjCreditcardclient(v []CreditcardclientListElement)`

SetAObjCreditcardclient sets AObjCreditcardclient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


