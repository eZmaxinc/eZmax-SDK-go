# ApikeyGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjApikey** | [**[]ApikeyListElement**](ApikeyListElement.md) |  | 

## Methods

### NewApikeyGetListV1ResponseMPayload

`func NewApikeyGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjApikey []ApikeyListElement, ) *ApikeyGetListV1ResponseMPayload`

NewApikeyGetListV1ResponseMPayload instantiates a new ApikeyGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeyGetListV1ResponseMPayloadWithDefaults

`func NewApikeyGetListV1ResponseMPayloadWithDefaults() *ApikeyGetListV1ResponseMPayload`

NewApikeyGetListV1ResponseMPayloadWithDefaults instantiates a new ApikeyGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *ApikeyGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *ApikeyGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *ApikeyGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *ApikeyGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *ApikeyGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *ApikeyGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjApikey

`func (o *ApikeyGetListV1ResponseMPayload) GetAObjApikey() []ApikeyListElement`

GetAObjApikey returns the AObjApikey field if non-nil, zero value otherwise.

### GetAObjApikeyOk

`func (o *ApikeyGetListV1ResponseMPayload) GetAObjApikeyOk() (*[]ApikeyListElement, bool)`

GetAObjApikeyOk returns a tuple with the AObjApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjApikey

`func (o *ApikeyGetListV1ResponseMPayload) SetAObjApikey(v []ApikeyListElement)`

SetAObjApikey sets AObjApikey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


