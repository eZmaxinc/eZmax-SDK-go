# BillingentityinternalGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjBillingentityinternal** | [**[]BillingentityinternalListElement**](BillingentityinternalListElement.md) |  | 

## Methods

### NewBillingentityinternalGetListV1ResponseMPayload

`func NewBillingentityinternalGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjBillingentityinternal []BillingentityinternalListElement, ) *BillingentityinternalGetListV1ResponseMPayload`

NewBillingentityinternalGetListV1ResponseMPayload instantiates a new BillingentityinternalGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingentityinternalGetListV1ResponseMPayloadWithDefaults

`func NewBillingentityinternalGetListV1ResponseMPayloadWithDefaults() *BillingentityinternalGetListV1ResponseMPayload`

NewBillingentityinternalGetListV1ResponseMPayloadWithDefaults instantiates a new BillingentityinternalGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *BillingentityinternalGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *BillingentityinternalGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjBillingentityinternal

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetAObjBillingentityinternal() []BillingentityinternalListElement`

GetAObjBillingentityinternal returns the AObjBillingentityinternal field if non-nil, zero value otherwise.

### GetAObjBillingentityinternalOk

`func (o *BillingentityinternalGetListV1ResponseMPayload) GetAObjBillingentityinternalOk() (*[]BillingentityinternalListElement, bool)`

GetAObjBillingentityinternalOk returns a tuple with the AObjBillingentityinternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjBillingentityinternal

`func (o *BillingentityinternalGetListV1ResponseMPayload) SetAObjBillingentityinternal(v []BillingentityinternalListElement)`

SetAObjBillingentityinternal sets AObjBillingentityinternal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


