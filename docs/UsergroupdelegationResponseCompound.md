# UsergroupdelegationResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupdelegationID** | **int32** | The unique ID of the Usergroupdelegation | 
**FkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**FkiUserID** | **int32** | The unique ID of the User | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SUserLoginname** | **string** | The login name of the User. | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUsergroupNameX** | **string** | The Name of the Usergroup in the language of the requester | 

## Methods

### NewUsergroupdelegationResponseCompound

`func NewUsergroupdelegationResponseCompound(pkiUsergroupdelegationID int32, fkiUsergroupID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sUsergroupNameX string, ) *UsergroupdelegationResponseCompound`

NewUsergroupdelegationResponseCompound instantiates a new UsergroupdelegationResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupdelegationResponseCompoundWithDefaults

`func NewUsergroupdelegationResponseCompoundWithDefaults() *UsergroupdelegationResponseCompound`

NewUsergroupdelegationResponseCompoundWithDefaults instantiates a new UsergroupdelegationResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupdelegationID

`func (o *UsergroupdelegationResponseCompound) GetPkiUsergroupdelegationID() int32`

GetPkiUsergroupdelegationID returns the PkiUsergroupdelegationID field if non-nil, zero value otherwise.

### GetPkiUsergroupdelegationIDOk

`func (o *UsergroupdelegationResponseCompound) GetPkiUsergroupdelegationIDOk() (*int32, bool)`

GetPkiUsergroupdelegationIDOk returns a tuple with the PkiUsergroupdelegationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupdelegationID

`func (o *UsergroupdelegationResponseCompound) SetPkiUsergroupdelegationID(v int32)`

SetPkiUsergroupdelegationID sets PkiUsergroupdelegationID field to given value.


### GetFkiUsergroupID

`func (o *UsergroupdelegationResponseCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupdelegationResponseCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupdelegationResponseCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupdelegationResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupdelegationResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupdelegationResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetSUserFirstname

`func (o *UsergroupdelegationResponseCompound) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UsergroupdelegationResponseCompound) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UsergroupdelegationResponseCompound) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UsergroupdelegationResponseCompound) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UsergroupdelegationResponseCompound) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UsergroupdelegationResponseCompound) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UsergroupdelegationResponseCompound) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UsergroupdelegationResponseCompound) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UsergroupdelegationResponseCompound) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetSEmailAddress

`func (o *UsergroupdelegationResponseCompound) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UsergroupdelegationResponseCompound) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UsergroupdelegationResponseCompound) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *UsergroupdelegationResponseCompound) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *UsergroupdelegationResponseCompound) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupdelegationResponseCompound) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupdelegationResponseCompound) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


