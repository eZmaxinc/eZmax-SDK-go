# UsergroupmembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupmembershipID** | **int32** | The unique ID of the Usergroupmembership | 
**FkiUsergroupID** | **int32** | The unique ID of the Usergroup | 
**FkiUserID** | **int32** | The unique ID of the User | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SUserLoginname** | **string** | The login name of the User. | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUsergroupNameX** | **string** | The Name of the Usergroup in the language of the requester | 

## Methods

### NewUsergroupmembershipResponse

`func NewUsergroupmembershipResponse(pkiUsergroupmembershipID int32, fkiUsergroupID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sUsergroupNameX string, ) *UsergroupmembershipResponse`

NewUsergroupmembershipResponse instantiates a new UsergroupmembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupmembershipResponseWithDefaults

`func NewUsergroupmembershipResponseWithDefaults() *UsergroupmembershipResponse`

NewUsergroupmembershipResponseWithDefaults instantiates a new UsergroupmembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupmembershipID

`func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipID() int32`

GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field if non-nil, zero value otherwise.

### GetPkiUsergroupmembershipIDOk

`func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipIDOk() (*int32, bool)`

GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupmembershipID

`func (o *UsergroupmembershipResponse) SetPkiUsergroupmembershipID(v int32)`

SetPkiUsergroupmembershipID sets PkiUsergroupmembershipID field to given value.


### GetFkiUsergroupID

`func (o *UsergroupmembershipResponse) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupmembershipResponse) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupmembershipResponse) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupmembershipResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupmembershipResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupmembershipResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetSUserFirstname

`func (o *UsergroupmembershipResponse) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UsergroupmembershipResponse) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UsergroupmembershipResponse) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UsergroupmembershipResponse) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UsergroupmembershipResponse) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UsergroupmembershipResponse) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UsergroupmembershipResponse) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UsergroupmembershipResponse) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UsergroupmembershipResponse) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetSEmailAddress

`func (o *UsergroupmembershipResponse) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UsergroupmembershipResponse) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UsergroupmembershipResponse) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *UsergroupmembershipResponse) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *UsergroupmembershipResponse) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupmembershipResponse) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupmembershipResponse) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


