# UserstagedGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjUserstaged** | [**[]UserstagedListElement**](UserstagedListElement.md) |  | 

## Methods

### NewUserstagedGetListV1ResponseMPayload

`func NewUserstagedGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjUserstaged []UserstagedListElement, ) *UserstagedGetListV1ResponseMPayload`

NewUserstagedGetListV1ResponseMPayload instantiates a new UserstagedGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserstagedGetListV1ResponseMPayloadWithDefaults

`func NewUserstagedGetListV1ResponseMPayloadWithDefaults() *UserstagedGetListV1ResponseMPayload`

NewUserstagedGetListV1ResponseMPayloadWithDefaults instantiates a new UserstagedGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *UserstagedGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *UserstagedGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *UserstagedGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *UserstagedGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *UserstagedGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *UserstagedGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjUserstaged

`func (o *UserstagedGetListV1ResponseMPayload) GetAObjUserstaged() []UserstagedListElement`

GetAObjUserstaged returns the AObjUserstaged field if non-nil, zero value otherwise.

### GetAObjUserstagedOk

`func (o *UserstagedGetListV1ResponseMPayload) GetAObjUserstagedOk() (*[]UserstagedListElement, bool)`

GetAObjUserstagedOk returns a tuple with the AObjUserstaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUserstaged

`func (o *UserstagedGetListV1ResponseMPayload) SetAObjUserstaged(v []UserstagedListElement)`

SetAObjUserstaged sets AObjUserstaged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


