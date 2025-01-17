# UsergroupmembershipRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupmembershipID** | Pointer to **int32** | The unique ID of the Usergroupmembership | [optional] 
**FkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupexternalID** | Pointer to **int32** | The unique ID of the Usergroupexternal | [optional] 

## Methods

### NewUsergroupmembershipRequestCompound

`func NewUsergroupmembershipRequestCompound(fkiUsergroupID int32, ) *UsergroupmembershipRequestCompound`

NewUsergroupmembershipRequestCompound instantiates a new UsergroupmembershipRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupmembershipRequestCompoundWithDefaults

`func NewUsergroupmembershipRequestCompoundWithDefaults() *UsergroupmembershipRequestCompound`

NewUsergroupmembershipRequestCompoundWithDefaults instantiates a new UsergroupmembershipRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupmembershipID

`func (o *UsergroupmembershipRequestCompound) GetPkiUsergroupmembershipID() int32`

GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field if non-nil, zero value otherwise.

### GetPkiUsergroupmembershipIDOk

`func (o *UsergroupmembershipRequestCompound) GetPkiUsergroupmembershipIDOk() (*int32, bool)`

GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupmembershipID

`func (o *UsergroupmembershipRequestCompound) SetPkiUsergroupmembershipID(v int32)`

SetPkiUsergroupmembershipID sets PkiUsergroupmembershipID field to given value.

### HasPkiUsergroupmembershipID

`func (o *UsergroupmembershipRequestCompound) HasPkiUsergroupmembershipID() bool`

HasPkiUsergroupmembershipID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupmembershipRequestCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupmembershipRequestCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupmembershipRequestCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupmembershipRequestCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *UsergroupmembershipRequestCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupexternalID

`func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupexternalID() int32`

GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field if non-nil, zero value otherwise.

### GetFkiUsergroupexternalIDOk

`func (o *UsergroupmembershipRequestCompound) GetFkiUsergroupexternalIDOk() (*int32, bool)`

GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupexternalID

`func (o *UsergroupmembershipRequestCompound) SetFkiUsergroupexternalID(v int32)`

SetFkiUsergroupexternalID sets FkiUsergroupexternalID field to given value.

### HasFkiUsergroupexternalID

`func (o *UsergroupmembershipRequestCompound) HasFkiUsergroupexternalID() bool`

HasFkiUsergroupexternalID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


