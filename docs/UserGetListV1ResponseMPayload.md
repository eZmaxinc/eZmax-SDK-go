# UserGetListV1ResponseMPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IRowReturned** | **int32** | The number of rows returned | 
**IRowFiltered** | **int32** | The number of rows matching your filters (if any) or the total number of rows | 
**AObjUser** | [**[]UserListElement**](UserListElement.md) |  | 

## Methods

### NewUserGetListV1ResponseMPayload

`func NewUserGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjUser []UserListElement, ) *UserGetListV1ResponseMPayload`

NewUserGetListV1ResponseMPayload instantiates a new UserGetListV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetListV1ResponseMPayloadWithDefaults

`func NewUserGetListV1ResponseMPayloadWithDefaults() *UserGetListV1ResponseMPayload`

NewUserGetListV1ResponseMPayloadWithDefaults instantiates a new UserGetListV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIRowReturned

`func (o *UserGetListV1ResponseMPayload) GetIRowReturned() int32`

GetIRowReturned returns the IRowReturned field if non-nil, zero value otherwise.

### GetIRowReturnedOk

`func (o *UserGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool)`

GetIRowReturnedOk returns a tuple with the IRowReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowReturned

`func (o *UserGetListV1ResponseMPayload) SetIRowReturned(v int32)`

SetIRowReturned sets IRowReturned field to given value.


### GetIRowFiltered

`func (o *UserGetListV1ResponseMPayload) GetIRowFiltered() int32`

GetIRowFiltered returns the IRowFiltered field if non-nil, zero value otherwise.

### GetIRowFilteredOk

`func (o *UserGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool)`

GetIRowFilteredOk returns a tuple with the IRowFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIRowFiltered

`func (o *UserGetListV1ResponseMPayload) SetIRowFiltered(v int32)`

SetIRowFiltered sets IRowFiltered field to given value.


### GetAObjUser

`func (o *UserGetListV1ResponseMPayload) GetAObjUser() []UserListElement`

GetAObjUser returns the AObjUser field if non-nil, zero value otherwise.

### GetAObjUserOk

`func (o *UserGetListV1ResponseMPayload) GetAObjUserOk() (*[]UserListElement, bool)`

GetAObjUserOk returns a tuple with the AObjUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjUser

`func (o *UserGetListV1ResponseMPayload) SetAObjUser(v []UserListElement)`

SetAObjUser sets AObjUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


