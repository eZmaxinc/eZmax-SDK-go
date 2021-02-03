# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserID** | **int32** | The unique ID of the User | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EUserType** | [**FieldEUserType**](Field-eUserType.md) |  | 
**SUserFirstname** | **string** | The First name of the user | 
**SUserLastname** | **string** | The Last name of the user | 
**SUserLoginname** | **string** | The Login name of the User. | 
**ObjAudit** | [**CommonAudit**](Common-Audit.md) |  | 

## Methods

### NewUserResponse

`func NewUserResponse(pkiUserID int32, fkiLanguageID int32, eUserType FieldEUserType, sUserFirstname string, sUserLastname string, sUserLoginname string, objAudit CommonAudit, ) *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserID

`func (o *UserResponse) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *UserResponse) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *UserResponse) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.


### GetFkiLanguageID

`func (o *UserResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *UserResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *UserResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEUserType

`func (o *UserResponse) GetEUserType() FieldEUserType`

GetEUserType returns the EUserType field if non-nil, zero value otherwise.

### GetEUserTypeOk

`func (o *UserResponse) GetEUserTypeOk() (*FieldEUserType, bool)`

GetEUserTypeOk returns a tuple with the EUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserType

`func (o *UserResponse) SetEUserType(v FieldEUserType)`

SetEUserType sets EUserType field to given value.


### GetSUserFirstname

`func (o *UserResponse) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UserResponse) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UserResponse) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UserResponse) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UserResponse) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UserResponse) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UserResponse) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UserResponse) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UserResponse) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetObjAudit

`func (o *UserResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *UserResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *UserResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


