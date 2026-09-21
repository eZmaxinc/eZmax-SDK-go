# PurchaseGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjPurchase** | [**[]PurchaseListElement**](PurchaseListElement.md) |  | 

## Methods

### NewPurchaseGetListV1ResponseMPayload

`func NewPurchaseGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjPurchase []PurchaseListElement, ) *PurchaseGetListV1ResponseMPayload`

NewPurchaseGetListV1ResponseMPayload instantiates a new PurchaseGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseGetListV1ResponseMPayloadWithDefaults

`func NewPurchaseGetListV1ResponseMPayloadWithDefaults() *PurchaseGetListV1ResponseMPayload`

NewPurchaseGetListV1ResponseMPayloadWithDefaults instantiates a new PurchaseGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *PurchaseGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *PurchaseGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *PurchaseGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *PurchaseGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *PurchaseGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *PurchaseGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjPurchase

`func (o *PurchaseGetListV1ResponseMPayload) GetAObjPurchase() []PurchaseListElement`

GetAObjPurchase returns the AObjPurchase field if non-nil, zero value otherwise.

### GetAObjPurchaseOk

`func (o *PurchaseGetListV1ResponseMPayload) GetAObjPurchaseOk() (*[]PurchaseListElement, bool)`

GetAObjPurchaseOk returns a tuple with the AObjPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPurchase

`func (o *PurchaseGetListV1ResponseMPayload) SetAObjPurchase(v []PurchaseListElement)`

SetAObjPurchase sets AObjPurchase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


