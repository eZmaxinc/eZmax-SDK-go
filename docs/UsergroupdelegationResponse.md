# UsergroupdelegationResponse

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

### NewUsergroupdelegationResponse

`func NewUsergroupdelegationResponse(pkiUsergroupdelegationID int32, fkiUsergroupID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sUsergroupNameX string, ) *UsergroupdelegationResponse`

NewUsergroupdelegationResponse instantiates a new UsergroupdelegationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupdelegationResponseWithDefaults

`func NewUsergroupdelegationResponseWithDefaults() *UsergroupdelegationResponse`

NewUsergroupdelegationResponseWithDefaults instantiates a new UsergroupdelegationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupdelegationID

`func (o *UsergroupdelegationResponse) GetPkiUsergroupdelegationID() int32`

GetPkiUsergroupdelegationID returns the PkiUsergroupdelegationID field if non-nil, zero value otherwise.

### GetPkiUsergroupdelegationIDOk

`func (o *UsergroupdelegationResponse) GetPkiUsergroupdelegationIDOk() (*int32, bool)`

GetPkiUsergroupdelegationIDOk returns a tuple with the PkiUsergroupdelegationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupdelegationID

`func (o *UsergroupdelegationResponse) SetPkiUsergroupdelegationID(v int32)`

SetPkiUsergroupdelegationID sets PkiUsergroupdelegationID field to given value.


### GetFkiUsergroupID

`func (o *UsergroupdelegationResponse) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *UsergroupdelegationResponse) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *UsergroupdelegationResponse) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.


### GetFkiUserID

`func (o *UsergroupdelegationResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupdelegationResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupdelegationResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetSUserFirstname

`func (o *UsergroupdelegationResponse) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UsergroupdelegationResponse) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UsergroupdelegationResponse) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UsergroupdelegationResponse) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UsergroupdelegationResponse) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UsergroupdelegationResponse) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UsergroupdelegationResponse) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UsergroupdelegationResponse) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UsergroupdelegationResponse) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetSEmailAddress

`func (o *UsergroupdelegationResponse) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UsergroupdelegationResponse) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UsergroupdelegationResponse) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *UsergroupdelegationResponse) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUsergroupNameX

`func (o *UsergroupdelegationResponse) GetSUsergroupNameX() string`

GetSUsergroupNameX returns the SUsergroupNameX field if non-nil, zero value otherwise.

### GetSUsergroupNameXOk

`func (o *UsergroupdelegationResponse) GetSUsergroupNameXOk() (*string, bool)`

GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupNameX

`func (o *UsergroupdelegationResponse) SetSUsergroupNameX(v string)`

SetSUsergroupNameX sets SUsergroupNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


