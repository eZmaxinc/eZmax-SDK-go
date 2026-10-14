# AdjustmentGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjAdjustment** | [**[]AdjustmentListElement**](AdjustmentListElement.md) |  | 

## Methods

### NewAdjustmentGetListV1ResponseMPayload

`func NewAdjustmentGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjAdjustment []AdjustmentListElement, ) *AdjustmentGetListV1ResponseMPayload`

NewAdjustmentGetListV1ResponseMPayload instantiates a new AdjustmentGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentGetListV1ResponseMPayloadWithDefaults

`func NewAdjustmentGetListV1ResponseMPayloadWithDefaults() *AdjustmentGetListV1ResponseMPayload`

NewAdjustmentGetListV1ResponseMPayloadWithDefaults instantiates a new AdjustmentGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *AdjustmentGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *AdjustmentGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *AdjustmentGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *AdjustmentGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *AdjustmentGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *AdjustmentGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjAdjustment

`func (o *AdjustmentGetListV1ResponseMPayload) GetAObjAdjustment() []AdjustmentListElement`

GetAObjAdjustment returns the AObjAdjustment field if non-nil, zero value otherwise.

### GetAObjAdjustmentOk

`func (o *AdjustmentGetListV1ResponseMPayload) GetAObjAdjustmentOk() (*[]AdjustmentListElement, bool)`

GetAObjAdjustmentOk returns a tuple with the AObjAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAdjustment

`func (o *AdjustmentGetListV1ResponseMPayload) SetAObjAdjustment(v []AdjustmentListElement)`

SetAObjAdjustment sets AObjAdjustment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


