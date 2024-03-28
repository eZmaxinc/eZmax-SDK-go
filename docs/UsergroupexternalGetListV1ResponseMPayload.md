# UsergroupexternalGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjUsergroupexternal** | [**[]UsergroupexternalListElement**](UsergroupexternalListElement.md) |  | 

## Methods

### NewUsergroupexternalGetListV1ResponseMPayload

`func NewUsergroupexternalGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjUsergroupexternal []UsergroupexternalListElement, ) *UsergroupexternalGetListV1ResponseMPayload`

NewUsergroupexternalGetListV1ResponseMPayload instantiates a new UsergroupexternalGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupexternalGetListV1ResponseMPayloadWithDefaults

`func NewUsergroupexternalGetListV1ResponseMPayloadWithDefaults() *UsergroupexternalGetListV1ResponseMPayload`

NewUsergroupexternalGetListV1ResponseMPayloadWithDefaults instantiates a new UsergroupexternalGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *UsergroupexternalGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *UsergroupexternalGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjUsergroupexternal

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetAObjUsergroupexternal() []UsergroupexternalListElement`

GetAObjUsergroupexternal returns the AObjUsergroupexternal field if non-nil, zero value otherwise.

### GetAObjUsergroupexternalOk

`func (o *UsergroupexternalGetListV1ResponseMPayload) GetAObjUsergroupexternalOk() (*[]UsergroupexternalListElement, bool)`

GetAObjUsergroupexternalOk returns a tuple with the AObjUsergroupexternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUsergroupexternal

`func (o *UsergroupexternalGetListV1ResponseMPayload) SetAObjUsergroupexternal(v []UsergroupexternalListElement)`

SetAObjUsergroupexternal sets AObjUsergroupexternal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


