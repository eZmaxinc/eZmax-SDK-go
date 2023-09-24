# UsergroupdelegationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupdelegationID** | Pointer to **int32** | The unique ID of the Usergroupdelegation | [optional] 
**FkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**FkiUserID** | **int32** | The unique ID of the User | 

## Methods

### NewUsergroupdelegationRequest

`func NewUsergroupdelegationRequest(fkiUsergroupID int32, fkiUserID int32, ) *UsergroupdelegationRequest`

NewUsergroupdelegationRequest instantiates a new UsergroupdelegationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupdelegationRequestWithDefaults

`func NewUsergroupdelegationRequestWithDefaults() *UsergroupdelegationRequest`

NewUsergroupdelegationRequestWithDefaults instantiates a new UsergroupdelegationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupdelegationID

`func (o *UsergroupdelegationRequest) GetPkiUsergroupdelegationID() int32`

GetPkiUsergroupdelegationID returns the PkiUsergroupdelegationID field if non-nil, zero value otherwise.

### GetPkiUsergroupdelegationIDOk

`func (o *UsergroupdelegationRequest) GetPkiUsergroupdelegationIDOk() (*int32, bool)`

GetPkiUsergroupdelegationIDOk returns a tuple with the PkiUsergroupdelegationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupdelegationID

`func (o *UsergroupdelegationRequest) SetPkiUsergroupdelegationID(v int32)`

SetPkiUsergroupdelegationID sets PkiUsergroupdelegationID field to given value.

### HasPkiUsergroupdelegationID

`func (o *UsergroupdelegationRequest) HasPkiUsergroupdelegationID() bool`

HasPkiUsergroupdelegationID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *UsergroupdelegationRequest) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupdelegationRequest) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupdelegationRequest) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupdelegationRequest) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupdelegationRequest) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupdelegationRequest) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


