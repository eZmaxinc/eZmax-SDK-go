# UsergroupexternalmembershipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUsergroupexternalmembershipID** | **int32** | The unique ID of the Usergroupexternalmembership | 
**FkiUsergroupexternalID** | **int32** | The unique ID of the Usergroupexternal | 
**FkiUserID** | **int32** | The unique ID of the User | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SUserLoginname** | **string** | The login name of the User. | 
**SEmailAddress** | **string** | The email address. | 
**SUsergroupexternalName** | **string** | The name of the Usergroupexternal | 

## Methods

### NewUsergroupexternalmembershipResponse

`func NewUsergroupexternalmembershipResponse(pkiUsergroupexternalmembershipID int32, fkiUsergroupexternalID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sEmailAddress string, sUsergroupexternalName string, ) *UsergroupexternalmembershipResponse`

NewUsergroupexternalmembershipResponse instantiates a new UsergroupexternalmembershipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsergroupexternalmembershipResponseWithDefaults

`func NewUsergroupexternalmembershipResponseWithDefaults() *UsergroupexternalmembershipResponse`

NewUsergroupexternalmembershipResponseWithDefaults instantiates a new UsergroupexternalmembershipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUsergroupexternalmembershipID

`func (o *UsergroupexternalmembershipResponse) GetPkiUsergroupexternalmembershipID() int32`

GetPkiUsergroupexternalmembershipID returns the PkiUsergroupexternalmembershipID field if non-nil, zero value otherwise.

### GetPkiUsergroupexternalmembershipIDOk

`func (o *UsergroupexternalmembershipResponse) GetPkiUsergroupexternalmembershipIDOk() (*int32, bool)`

GetPkiUsergroupexternalmembershipIDOk returns a tuple with the PkiUsergroupexternalmembershipID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUsergroupexternalmembershipID

`func (o *UsergroupexternalmembershipResponse) SetPkiUsergroupexternalmembershipID(v int32)`

SetPkiUsergroupexternalmembershipID sets PkiUsergroupexternalmembershipID field to given value.


### GetFkiUsergroupexternalID

`func (o *UsergroupexternalmembershipResponse) GetFkiUsergroupexternalID() int32`

GetFkiUsergroupexternalID returns the FkiUsergroupexternalID field if non-nil, zero value otherwise.

### GetFkiUsergroupexternalIDOk

`func (o *UsergroupexternalmembershipResponse) GetFkiUsergroupexternalIDOk() (*int32, bool)`

GetFkiUsergroupexternalIDOk returns a tuple with the FkiUsergroupexternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupexternalID

`func (o *UsergroupexternalmembershipResponse) SetFkiUsergroupexternalID(v int32)`

SetFkiUsergroupexternalID sets FkiUsergroupexternalID field to given value.


### GetFkiUserID

`func (o *UsergroupexternalmembershipResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *UsergroupexternalmembershipResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *UsergroupexternalmembershipResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetSUserFirstname

`func (o *UsergroupexternalmembershipResponse) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UsergroupexternalmembershipResponse) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UsergroupexternalmembershipResponse) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UsergroupexternalmembershipResponse) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UsergroupexternalmembershipResponse) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UsergroupexternalmembershipResponse) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UsergroupexternalmembershipResponse) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UsergroupexternalmembershipResponse) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UsergroupexternalmembershipResponse) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetSEmailAddress

`func (o *UsergroupexternalmembershipResponse) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *UsergroupexternalmembershipResponse) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *UsergroupexternalmembershipResponse) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.


### GetSUsergroupexternalName

`func (o *UsergroupexternalmembershipResponse) GetSUsergroupexternalName() string`

GetSUsergroupexternalName returns the SUsergroupexternalName field if non-nil, zero value otherwise.

### GetSUsergroupexternalNameOk

`func (o *UsergroupexternalmembershipResponse) GetSUsergroupexternalNameOk() (*string, bool)`

GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUsergroupexternalName

`func (o *UsergroupexternalmembershipResponse) SetSUsergroupexternalName(v string)`

SetSUsergroupexternalName sets SUsergroupexternalName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


