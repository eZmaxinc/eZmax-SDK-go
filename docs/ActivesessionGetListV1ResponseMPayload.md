# ActivesessionGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjActivesession** | [**[]ActivesessionListElement**](ActivesessionListElement.md) |  | 

## Methods

### NewActivesessionGetListV1ResponseMPayload

`func NewActivesessionGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjActivesession []ActivesessionListElement, ) *ActivesessionGetListV1ResponseMPayload`

NewActivesessionGetListV1ResponseMPayload instantiates a new ActivesessionGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionGetListV1ResponseMPayloadWithDefaults

`func NewActivesessionGetListV1ResponseMPayloadWithDefaults() *ActivesessionGetListV1ResponseMPayload`

NewActivesessionGetListV1ResponseMPayloadWithDefaults instantiates a new ActivesessionGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *ActivesessionGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *ActivesessionGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *ActivesessionGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *ActivesessionGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *ActivesessionGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *ActivesessionGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjActivesession

`func (o *ActivesessionGetListV1ResponseMPayload) GetAObjActivesession() []ActivesessionListElement`

GetAObjActivesession returns the AObjActivesession field if non-nil, zero value otherwise.

### GetAObjActivesessionOk

`func (o *ActivesessionGetListV1ResponseMPayload) GetAObjActivesessionOk() (*[]ActivesessionListElement, bool)`

GetAObjActivesessionOk returns a tuple with the AObjActivesession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjActivesession

`func (o *ActivesessionGetListV1ResponseMPayload) SetAObjActivesession(v []ActivesessionListElement)`

SetAObjActivesession sets AObjActivesession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


