# ActivesessionResponseCompoundUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserID** | **int32** | The unique ID of the User | 
**FkiTimezoneID** | **int32** | The unique ID of the Timezone | 
**SAvatarUrl** | Pointer to **string** | The url of the picture used as avatar | [optional] 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**EUserEzsignsendreminderfrequency** | [**FieldEUserEzsignsendreminderfrequency**](FieldEUserEzsignsendreminderfrequency.md) |  | 
**IUserInterfacecolor** | **int32** | The int32 representation of the interface color. For example, RGB color #39435B would be 3752795 | 
**BUserInterfacedark** | **bool** | Whether to use a dark mode interface | 
**IUserListresult** | **int32** | The number of rows to return by default in lists | 

## Methods

### NewActivesessionResponseCompoundUser

`func NewActivesessionResponseCompoundUser(pkiUserID int32, fkiTimezoneID int32, sUserFirstname string, sUserLastname string, eUserEzsignsendreminderfrequency FieldEUserEzsignsendreminderfrequency, iUserInterfacecolor int32, bUserInterfacedark bool, iUserListresult int32, ) *ActivesessionResponseCompoundUser`

NewActivesessionResponseCompoundUser instantiates a new ActivesessionResponseCompoundUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionResponseCompoundUserWithDefaults

`func NewActivesessionResponseCompoundUserWithDefaults() *ActivesessionResponseCompoundUser`

NewActivesessionResponseCompoundUserWithDefaults instantiates a new ActivesessionResponseCompoundUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserID

`func (o *ActivesessionResponseCompoundUser) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *ActivesessionResponseCompoundUser) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *ActivesessionResponseCompoundUser) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.


### GetFkiTimezoneID

`func (o *ActivesessionResponseCompoundUser) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *ActivesessionResponseCompoundUser) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *ActivesessionResponseCompoundUser) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.


### GetSAvatarUrl

`func (o *ActivesessionResponseCompoundUser) GetSAvatarUrl() string`

GetSAvatarUrl returns the SAvatarUrl field if non-nil, zero value otherwise.

### GetSAvatarUrlOk

`func (o *ActivesessionResponseCompoundUser) GetSAvatarUrlOk() (*string, bool)`

GetSAvatarUrlOk returns a tuple with the SAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAvatarUrl

`func (o *ActivesessionResponseCompoundUser) SetSAvatarUrl(v string)`

SetSAvatarUrl sets SAvatarUrl field to given value.

### HasSAvatarUrl

`func (o *ActivesessionResponseCompoundUser) HasSAvatarUrl() bool`

HasSAvatarUrl returns a boolean if a field has been set.

### GetSUserFirstname

`func (o *ActivesessionResponseCompoundUser) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *ActivesessionResponseCompoundUser) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *ActivesessionResponseCompoundUser) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *ActivesessionResponseCompoundUser) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *ActivesessionResponseCompoundUser) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *ActivesessionResponseCompoundUser) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSEmailAddress

`func (o *ActivesessionResponseCompoundUser) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *ActivesessionResponseCompoundUser) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *ActivesessionResponseCompoundUser) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *ActivesessionResponseCompoundUser) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetEUserEzsignsendreminderfrequency

`func (o *ActivesessionResponseCompoundUser) GetEUserEzsignsendreminderfrequency() FieldEUserEzsignsendreminderfrequency`

GetEUserEzsignsendreminderfrequency returns the EUserEzsignsendreminderfrequency field if non-nil, zero value otherwise.

### GetEUserEzsignsendreminderfrequencyOk

`func (o *ActivesessionResponseCompoundUser) GetEUserEzsignsendreminderfrequencyOk() (*FieldEUserEzsignsendreminderfrequency, bool)`

GetEUserEzsignsendreminderfrequencyOk returns a tuple with the EUserEzsignsendreminderfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignsendreminderfrequency

`func (o *ActivesessionResponseCompoundUser) SetEUserEzsignsendreminderfrequency(v FieldEUserEzsignsendreminderfrequency)`

SetEUserEzsignsendreminderfrequency sets EUserEzsignsendreminderfrequency field to given value.


### GetIUserInterfacecolor

`func (o *ActivesessionResponseCompoundUser) GetIUserInterfacecolor() int32`

GetIUserInterfacecolor returns the IUserInterfacecolor field if non-nil, zero value otherwise.

### GetIUserInterfacecolorOk

`func (o *ActivesessionResponseCompoundUser) GetIUserInterfacecolorOk() (*int32, bool)`

GetIUserInterfacecolorOk returns a tuple with the IUserInterfacecolor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIUserInterfacecolor

`func (o *ActivesessionResponseCompoundUser) SetIUserInterfacecolor(v int32)`

SetIUserInterfacecolor sets IUserInterfacecolor field to given value.


### GetBUserInterfacedark

`func (o *ActivesessionResponseCompoundUser) GetBUserInterfacedark() bool`

GetBUserInterfacedark returns the BUserInterfacedark field if non-nil, zero value otherwise.

### GetBUserInterfacedarkOk

`func (o *ActivesessionResponseCompoundUser) GetBUserInterfacedarkOk() (*bool, bool)`

GetBUserInterfacedarkOk returns a tuple with the BUserInterfacedark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserInterfacedark

`func (o *ActivesessionResponseCompoundUser) SetBUserInterfacedark(v bool)`

SetBUserInterfacedark sets BUserInterfacedark field to given value.


### GetIUserListresult

`func (o *ActivesessionResponseCompoundUser) GetIUserListresult() int32`

GetIUserListresult returns the IUserListresult field if non-nil, zero value otherwise.

### GetIUserListresultOk

`func (o *ActivesessionResponseCompoundUser) GetIUserListresultOk() (*int32, bool)`

GetIUserListresultOk returns a tuple with the IUserListresult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIUserListresult

`func (o *ActivesessionResponseCompoundUser) SetIUserListresult(v int32)`

SetIUserListresult sets IUserListresult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


