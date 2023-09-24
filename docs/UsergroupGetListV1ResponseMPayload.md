# UsergroupGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjUsergroup** | [**[]UsergroupListElement**](UsergroupListElement.md) |  | 

## Methods

### NewUsergroupGetListV1ResponseMPayload

`func NewUsergroupGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjUsergroup []UsergroupListElement, ) *UsergroupGetListV1ResponseMPayload`

NewUsergroupGetListV1ResponseMPayload instantiates a new UsergroupGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupGetListV1ResponseMPayloadWithDefaults

`func NewUsergroupGetListV1ResponseMPayloadWithDefaults() *UsergroupGetListV1ResponseMPayload`

NewUsergroupGetListV1ResponseMPayloadWithDefaults instantiates a new UsergroupGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *UsergroupGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *UsergroupGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *UsergroupGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *UsergroupGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *UsergroupGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *UsergroupGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjUsergroup

`func (o *UsergroupGetListV1ResponseMPayload) GetAObjUsergroup() []UsergroupListElement`

GetAObjUsergroup returns the AObjUsergroup field if non-nil, zero value otherwise.

### GetAObjUsergroupOk

`func (o *UsergroupGetListV1ResponseMPayload) GetAObjUsergroupOk() (*[]UsergroupListElement, bool)`

GetAObjUsergroupOk returns a tuple with the AObjUsergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroup

`func (o *UsergroupGetListV1ResponseMPayload) SetAObjUsergroup(v []UsergroupListElement)`

SetAObjUsergroup sets AObjUsergroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


