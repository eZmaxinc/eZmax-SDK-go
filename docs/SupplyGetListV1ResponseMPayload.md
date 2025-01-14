# SupplyGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjSupply** | [**[]SupplyListElement**](SupplyListElement.md) |  | 

## Methods

### NewSupplyGetListV1ResponseMPayload

`func NewSupplyGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjSupply []SupplyListElement, ) *SupplyGetListV1ResponseMPayload`

NewSupplyGetListV1ResponseMPayload instantiates a new SupplyGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyGetListV1ResponseMPayloadWithDefaults

`func NewSupplyGetListV1ResponseMPayloadWithDefaults() *SupplyGetListV1ResponseMPayload`

NewSupplyGetListV1ResponseMPayloadWithDefaults instantiates a new SupplyGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *SupplyGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *SupplyGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *SupplyGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *SupplyGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *SupplyGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *SupplyGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjSupply

`func (o *SupplyGetListV1ResponseMPayload) GetAObjSupply() []SupplyListElement`

GetAObjSupply returns the AObjSupply field if non-nil, zero value otherwise.

### GetAObjSupplyOk

`func (o *SupplyGetListV1ResponseMPayload) GetAObjSupplyOk() (*[]SupplyListElement, bool)`

GetAObjSupplyOk returns a tuple with the AObjSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjSupply

`func (o *SupplyGetListV1ResponseMPayload) SetAObjSupply(v []SupplyListElement)`

SetAObjSupply sets AObjSupply field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


