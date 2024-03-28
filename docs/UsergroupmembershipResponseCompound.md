# UsergroupmembershipResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupmembershipID** | **int32** | The unique ID of the Usergroupmembership | 
**FkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiUsergroupexternalID** | Pointer to **int32** | The unique ID of the Usergroupexternal | [optional] 
**SUserFirstname** | Pointer to **string** | The first name of the user | [optional] 
**SUserLastname** | Pointer to **string** | The last name of the user | [optional] 
**SUserLoginname** | Pointer to **string** | The login name of the User. | [optional] 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUsergroupNameX** | **string** | The Name of the Usergroup in the language of the requester | 
**SUsergroupexternalName** | Pointer to **string** | The name of the Usergroupexternal | [optional] 

## Methods

### NewUsergroupmembershipResponseCompound

`func NewUsergroupmembershipResponseCompound(pkiUsergroupmembershipID int32, fkiUsergroupID int32, sUsergroupNameX string, ) *UsergroupmembershipResponseCompound`

NewUsergroupmembershipResponseCompound instantiates a new UsergroupmembershipResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupmembershipResponseCompoundWithDefaults

`func NewUsergroupmembershipResponseCompoundWithDefaults() *UsergroupmembershipResponseCompound`

NewUsergroupmembershipResponseCompoundWithDefaults instantiates a new UsergroupmembershipResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupmembershipID

`func (o *UsergroupmembershipResponseCompound) GetPkiUsergroupmembershipID() int32`

GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field if non-nil, zero value otherwise.

### GetPkiUsergroupmembershipIDOk

`func (o *UsergroupmembershipResponseCompound) GetPkiUsergroupmembershipIDOk() (*int32, bool)`

GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupmembershipID

`func (o *UsergroupmembershipResponseCompound) SetPkiUsergroupmembershipID(v int32)`

SetPkiUsergroupmembershipID sets PkiUsergroupmembershipID field to given value.


### GetFkiUsergroupID

`func (o *UsergroupmembershipResponseCompound) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupmembershipResponseCompound) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupmembershipResponseCompound) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupmembershipResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupmembershipResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupmembershipResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *UsergroupmembershipResponseCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiUsergroupexternalID

`func (o *UsergroupmembershipResponseCompound) GetFkiUsergroupexternalID() int32`

GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field if non-nil, zero value otherwise.

### GetFkiUsergroupexternalIDOk

`func (o *UsergroupmembershipResponseCompound) GetFkiUsergroupexternalIDOk() (*int32, bool)`

GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupexternalID

`func (o *UsergroupmembershipResponseCompound) SetFkiUsergroupexternalID(v int32)`

SetFkiUsergroupexternalID sets FkiUsergroupexternalID field to given value.

### HasFkiUsergroupexternalID

`func (o *UsergroupmembershipResponseCompound) HasFkiUsergroupexternalID() bool`

HasFkiUsergroupexternalID returns a boolean if a field has been set.

### GetSUserFirstname

`func (o *UsergroupmembershipResponseCompound) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UsergroupmembershipResponseCompound) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UsergroupmembershipResponseCompound) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.

### HasSUserFirstname

`func (o *UsergroupmembershipResponseCompound) HasSUserFirstname() bool`

HasSUserFirstname returns a boolean if a field has been set.

### GetSUserLastname

`func (o *UsergroupmembershipResponseCompound) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UsergroupmembershipResponseCompound) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UsergroupmembershipResponseCompound) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.

### HasSUserLastname

`func (o *UsergroupmembershipResponseCompound) HasSUserLastname() bool`

HasSUserLastname returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *UsergroupmembershipResponseCompound) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UsergroupmembershipResponseCompound) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UsergroupmembershipResponseCompound) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *UsergroupmembershipResponseCompound) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.

### GetSEmailAddress

`func (o *UsergroupmembershipResponseCompound) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UsergroupmembershipResponseCompound) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UsergroupmembershipResponseCompound) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *UsergroupmembershipResponseCompound) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *UsergroupmembershipResponseCompound) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupmembershipResponseCompound) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupmembershipResponseCompound) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.


### GetSUsergroupexternalName

`func (o *UsergroupmembershipResponseCompound) GetSUsergroupexternalName() string`

GetSUsergroupexternalName returns the SUsergroupexternalName field if non-nil, zero value otherwise.

### GetSUsergroupexternalNameOk

`func (o *UsergroupmembershipResponseCompound) GetSUsergroupexternalNameOk() (*string, bool)`

GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupexternalName

`func (o *UsergroupmembershipResponseCompound) SetSUsergroupexternalName(v string)`

SetSUsergroupexternalName sets SUsergroupexternalName field to given value.

### HasSUsergroupexternalName

`func (o *UsergroupmembershipResponseCompound) HasSUsergroupexternalName() bool`

HasSUsergroupexternalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


