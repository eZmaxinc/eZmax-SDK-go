# UserResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserID** | **int32** | The unique ID of the User | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EUserType** | [**FieldEUserType**](FieldEUserType.md) |  | 
**SUserFirstname** | **string** | The First name of the user | 
**SUserLastname** | **string** | The Last name of the user | 
**SUserLoginname** | **string** | The Login name of the User. | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewUserResponseAllOf

`func NewUserResponseAllOf(pkiUserID int32, fkiLanguageID int32, eUserType FieldEUserType, sUserFirstname string, sUserLastname string, sUserLoginname string, objAudit CommonAudit, ) *UserResponseAllOf`

NewUserResponseAllOf instantiates a new UserResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseAllOfWithDefaults

`func NewUserResponseAllOfWithDefaults() *UserResponseAllOf`

NewUserResponseAllOfWithDefaults instantiates a new UserResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserID

`func (o *UserResponseAllOf) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *UserResponseAllOf) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *UserResponseAllOf) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.


### GetFkiLanguageID

`func (o *UserResponseAllOf) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *UserResponseAllOf) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *UserResponseAllOf) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEUserType

`func (o *UserResponseAllOf) GetEUserType() FieldEUserType`

GetEUserType returns the EUserType field if non-nil, zero value otherwise.

### GetEUserTypeOk

`func (o *UserResponseAllOf) GetEUserTypeOk() (*FieldEUserType, bool)`

GetEUserTypeOk returns a tuple with the EUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserType

`func (o *UserResponseAllOf) SetEUserType(v FieldEUserType)`

SetEUserType sets EUserType field to given value.


### GetSUserFirstname

`func (o *UserResponseAllOf) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UserResponseAllOf) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UserResponseAllOf) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UserResponseAllOf) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UserResponseAllOf) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UserResponseAllOf) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UserResponseAllOf) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UserResponseAllOf) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UserResponseAllOf) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetObjAudit

`func (o *UserResponseAllOf) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *UserResponseAllOf) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *UserResponseAllOf) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


