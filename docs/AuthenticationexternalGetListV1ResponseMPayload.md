# AuthenticationexternalGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjAuthenticationexternal** | [**[]AuthenticationexternalListElement**](AuthenticationexternalListElement.md) |  | 

## Methods

### NewAuthenticationexternalGetListV1ResponseMPayload

`func NewAuthenticationexternalGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjAuthenticationexternal []AuthenticationexternalListElement, ) *AuthenticationexternalGetListV1ResponseMPayload`

NewAuthenticationexternalGetListV1ResponseMPayload instantiates a new AuthenticationexternalGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationexternalGetListV1ResponseMPayloadWithDefaults

`func NewAuthenticationexternalGetListV1ResponseMPayloadWithDefaults() *AuthenticationexternalGetListV1ResponseMPayload`

NewAuthenticationexternalGetListV1ResponseMPayloadWithDefaults instantiates a new AuthenticationexternalGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *AuthenticationexternalGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *AuthenticationexternalGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjAuthenticationexternal

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetAObjAuthenticationexternal() []AuthenticationexternalListElement`

GetAObjAuthenticationexternal returns the AObjAuthenticationexternal field if non-nil, zero value otherwise.

### GetAObjAuthenticationexternalOk

`func (o *AuthenticationexternalGetListV1ResponseMPayload) GetAObjAuthenticationexternalOk() (*[]AuthenticationexternalListElement, bool)`

GetAObjAuthenticationexternalOk returns a tuple with the AObjAuthenticationexternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAuthenticationexternal

`func (o *AuthenticationexternalGetListV1ResponseMPayload) SetAObjAuthenticationexternal(v []AuthenticationexternalListElement)`

SetAObjAuthenticationexternal sets AObjAuthenticationexternal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


