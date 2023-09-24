# PermissionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiPermissionID** | **int32** | The unique ID of the Permission | 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiApikeyID** | Pointer to **int32** | The unique ID of the Apikey | [optional] 
**FkiUsergroupID** | Pointer to **int32** | The unique ID of the Usergroup | [optional] 
**FkiCompanyID** | Pointer to **int32** | The unique ID of the Company | [optional] 
**FkiModulesectionID** | **int32** | The unique ID of the Modulesection | 
**SCompanyNameX** | Pointer to **string** | The Name of the Company in the language of the requester | [optional] 

## Methods

### NewPermissionResponse

`func NewPermissionResponse(pkiPermissionID int32, fkiModulesectionID int32, ) *PermissionResponse`

NewPermissionResponse instantiates a new PermissionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionResponseWithDefaults

`func NewPermissionResponseWithDefaults() *PermissionResponse`

NewPermissionResponseWithDefaults instantiates a new PermissionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiPermissionID

`func (o *PermissionResponse) GetPkiPermissionID() int32`

GetPkiPermissionID returns the PkiPermissionID field if non-nil, zero value otherwise.

### GetPkiPermissionIDOk

`func (o *PermissionResponse) GetPkiPermissionIDOk() (*int32, bool)`

GetPkiPermissionIDOk returns a tuple with the PkiPermissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiPermissionID

`func (o *PermissionResponse) SetPkiPermissionID(v int32)`

SetPkiPermissionID sets PkiPermissionID field to given value.


### GetFkiUserID

`func (o *PermissionResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *PermissionResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *PermissionResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *PermissionResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiApikeyID

`func (o *PermissionResponse) GetFkiApikeyID() int32`

GetFkiApikeyID returns the FkiApikeyID field if non-nil, zero value otherwise.

### GetFkiApikeyIDOk

`func (o *PermissionResponse) GetFkiApikeyIDOk() (*int32, bool)`

GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiApikeyID

`func (o *PermissionResponse) SetFkiApikeyID(v int32)`

SetFkiApikeyID sets FkiApikeyID field to given value.

### HasFkiApikeyID

`func (o *PermissionResponse) HasFkiApikeyID() bool`

HasFkiApikeyID returns a boolean if a field has been set.

### GetFkiUsergroupID

`func (o *PermissionResponse) GetFkiUsergroupID() int32`

GetFkiUsergroupID returns the FkiUsergroupID field if non-nil, zero value otherwise.

### GetFkiUsergroupIDOk

`func (o *PermissionResponse) GetFkiUsergroupIDOk() (*int32, bool)`

GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUsergroupID

`func (o *PermissionResponse) SetFkiUsergroupID(v int32)`

SetFkiUsergroupID sets FkiUsergroupID field to given value.

### HasFkiUsergroupID

`func (o *PermissionResponse) HasFkiUsergroupID() bool`

HasFkiUsergroupID returns a boolean if a field has been set.

### GetFkiCompanyID

`func (o *PermissionResponse) GetFkiCompanyID() int32`

GetFkiCompanyID returns the FkiCompanyID field if non-nil, zero value otherwise.

### GetFkiCompanyIDOk

`func (o *PermissionResponse) GetFkiCompanyIDOk() (*int32, bool)`

GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyID

`func (o *PermissionResponse) SetFkiCompanyID(v int32)`

SetFkiCompanyID sets FkiCompanyID field to given value.

### HasFkiCompanyID

`func (o *PermissionResponse) HasFkiCompanyID() bool`

HasFkiCompanyID returns a boolean if a field has been set.

### GetFkiModulesectionID

`func (o *PermissionResponse) GetFkiModulesectionID() int32`

GetFkiModulesectionID returns the FkiModulesectionID field if non-nil, zero value otherwise.

### GetFkiModulesectionIDOk

`func (o *PermissionResponse) GetFkiModulesectionIDOk() (*int32, bool)`

GetFkiModulesectionIDOk returns a tuple with the FkiModulesectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModulesectionID

`func (o *PermissionResponse) SetFkiModulesectionID(v int32)`

SetFkiModulesectionID sets FkiModulesectionID field to given value.


### GetSCompanyNameX

`func (o *PermissionResponse) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *PermissionResponse) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *PermissionResponse) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.

### HasSCompanyNameX

`func (o *PermissionResponse) HasSCompanyNameX() bool`

HasSCompanyNameX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


