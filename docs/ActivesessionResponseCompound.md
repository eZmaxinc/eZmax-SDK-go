# ActivesessionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiEzsignuserID** | Pointer to **int32** | The unique ID of the Ezsignuser | [optional] 
**BSystemconfigurationEzsignpaidbyoffice** | Pointer to **bool** | Whether if Ezsign is paid by the company or not | [optional] 
**ESystemconfigurationEzsignofficeplan** | Pointer to [**FieldESystemconfigurationEzsignofficeplan**](FieldESystemconfigurationEzsignofficeplan.md) |  | [optional] 
**EUserEzsignaccess** | [**FieldEUserEzsignaccess**](FieldEUserEzsignaccess.md) |  | 
**EUserEzsignprepaid** | Pointer to [**FieldEUserEzsignprepaid**](FieldEUserEzsignprepaid.md) |  | [optional] 
**BUserEzsigntrial** | Pointer to **bool** | Whether the User&#39;s eZsign subscription is a trial | [optional] 
**DtUserEzsignprepaidexpiration** | Pointer to **string** | The eZsign prepaid expiration date | [optional] 
**APkiPermissionID** | **[]int32** | An array of permissions granted to the user or api key | 
**ObjUserReal** | [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | 
**ObjUserCloned** | Pointer to [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | [optional] 
**ObjApikey** | Pointer to [**ActivesessionResponseCompoundApikey**](ActivesessionResponseCompoundApikey.md) |  | [optional] 
**AEModuleInternalname** | **[]string** | An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key. | 

## Methods

### NewActivesessionResponseCompound

`func NewActivesessionResponseCompound(eUserEzsignaccess FieldEUserEzsignaccess, aPkiPermissionID []int32, objUserReal ActivesessionResponseCompoundUser, aEModuleInternalname []string, ) *ActivesessionResponseCompound`

NewActivesessionResponseCompound instantiates a new ActivesessionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionResponseCompoundWithDefaults

`func NewActivesessionResponseCompoundWithDefaults() *ActivesessionResponseCompound`

NewActivesessionResponseCompoundWithDefaults instantiates a new ActivesessionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiEzsignuserID

`func (o *ActivesessionResponseCompound) GetFkiEzsignuserID() int32`

GetFkiEzsignuserID returns the FkiEzsignuserID field if non-nil, zero value otherwise.

### GetFkiEzsignuserIDOk

`func (o *ActivesessionResponseCompound) GetFkiEzsignuserIDOk() (*int32, bool)`

GetFkiEzsignuserIDOk returns a tuple with the FkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignuserID

`func (o *ActivesessionResponseCompound) SetFkiEzsignuserID(v int32)`

SetFkiEzsignuserID sets FkiEzsignuserID field to given value.

### HasFkiEzsignuserID

`func (o *ActivesessionResponseCompound) HasFkiEzsignuserID() bool`

HasFkiEzsignuserID returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionResponseCompound) GetBSystemconfigurationEzsignpaidbyoffice() bool`

GetBSystemconfigurationEzsignpaidbyoffice returns the BSystemconfigurationEzsignpaidbyoffice field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpaidbyofficeOk

`func (o *ActivesessionResponseCompound) GetBSystemconfigurationEzsignpaidbyofficeOk() (*bool, bool)`

GetBSystemconfigurationEzsignpaidbyofficeOk returns a tuple with the BSystemconfigurationEzsignpaidbyoffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionResponseCompound) SetBSystemconfigurationEzsignpaidbyoffice(v bool)`

SetBSystemconfigurationEzsignpaidbyoffice sets BSystemconfigurationEzsignpaidbyoffice field to given value.

### HasBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionResponseCompound) HasBSystemconfigurationEzsignpaidbyoffice() bool`

HasBSystemconfigurationEzsignpaidbyoffice returns a boolean if a field has been set.

### GetESystemconfigurationEzsignofficeplan

`func (o *ActivesessionResponseCompound) GetESystemconfigurationEzsignofficeplan() FieldESystemconfigurationEzsignofficeplan`

GetESystemconfigurationEzsignofficeplan returns the ESystemconfigurationEzsignofficeplan field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignofficeplanOk

`func (o *ActivesessionResponseCompound) GetESystemconfigurationEzsignofficeplanOk() (*FieldESystemconfigurationEzsignofficeplan, bool)`

GetESystemconfigurationEzsignofficeplanOk returns a tuple with the ESystemconfigurationEzsignofficeplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsignofficeplan

`func (o *ActivesessionResponseCompound) SetESystemconfigurationEzsignofficeplan(v FieldESystemconfigurationEzsignofficeplan)`

SetESystemconfigurationEzsignofficeplan sets ESystemconfigurationEzsignofficeplan field to given value.

### HasESystemconfigurationEzsignofficeplan

`func (o *ActivesessionResponseCompound) HasESystemconfigurationEzsignofficeplan() bool`

HasESystemconfigurationEzsignofficeplan returns a boolean if a field has been set.

### GetEUserEzsignaccess

`func (o *ActivesessionResponseCompound) GetEUserEzsignaccess() FieldEUserEzsignaccess`

GetEUserEzsignaccess returns the EUserEzsignaccess field if non-nil, zero value otherwise.

### GetEUserEzsignaccessOk

`func (o *ActivesessionResponseCompound) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool)`

GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignaccess

`func (o *ActivesessionResponseCompound) SetEUserEzsignaccess(v FieldEUserEzsignaccess)`

SetEUserEzsignaccess sets EUserEzsignaccess field to given value.


### GetEUserEzsignprepaid

`func (o *ActivesessionResponseCompound) GetEUserEzsignprepaid() FieldEUserEzsignprepaid`

GetEUserEzsignprepaid returns the EUserEzsignprepaid field if non-nil, zero value otherwise.

### GetEUserEzsignprepaidOk

`func (o *ActivesessionResponseCompound) GetEUserEzsignprepaidOk() (*FieldEUserEzsignprepaid, bool)`

GetEUserEzsignprepaidOk returns a tuple with the EUserEzsignprepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignprepaid

`func (o *ActivesessionResponseCompound) SetEUserEzsignprepaid(v FieldEUserEzsignprepaid)`

SetEUserEzsignprepaid sets EUserEzsignprepaid field to given value.

### HasEUserEzsignprepaid

`func (o *ActivesessionResponseCompound) HasEUserEzsignprepaid() bool`

HasEUserEzsignprepaid returns a boolean if a field has been set.

### GetBUserEzsigntrial

`func (o *ActivesessionResponseCompound) GetBUserEzsigntrial() bool`

GetBUserEzsigntrial returns the BUserEzsigntrial field if non-nil, zero value otherwise.

### GetBUserEzsigntrialOk

`func (o *ActivesessionResponseCompound) GetBUserEzsigntrialOk() (*bool, bool)`

GetBUserEzsigntrialOk returns a tuple with the BUserEzsigntrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserEzsigntrial

`func (o *ActivesessionResponseCompound) SetBUserEzsigntrial(v bool)`

SetBUserEzsigntrial sets BUserEzsigntrial field to given value.

### HasBUserEzsigntrial

`func (o *ActivesessionResponseCompound) HasBUserEzsigntrial() bool`

HasBUserEzsigntrial returns a boolean if a field has been set.

### GetDtUserEzsignprepaidexpiration

`func (o *ActivesessionResponseCompound) GetDtUserEzsignprepaidexpiration() string`

GetDtUserEzsignprepaidexpiration returns the DtUserEzsignprepaidexpiration field if non-nil, zero value otherwise.

### GetDtUserEzsignprepaidexpirationOk

`func (o *ActivesessionResponseCompound) GetDtUserEzsignprepaidexpirationOk() (*string, bool)`

GetDtUserEzsignprepaidexpirationOk returns a tuple with the DtUserEzsignprepaidexpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtUserEzsignprepaidexpiration

`func (o *ActivesessionResponseCompound) SetDtUserEzsignprepaidexpiration(v string)`

SetDtUserEzsignprepaidexpiration sets DtUserEzsignprepaidexpiration field to given value.

### HasDtUserEzsignprepaidexpiration

`func (o *ActivesessionResponseCompound) HasDtUserEzsignprepaidexpiration() bool`

HasDtUserEzsignprepaidexpiration returns a boolean if a field has been set.

### GetAPkiPermissionID

`func (o *ActivesessionResponseCompound) GetAPkiPermissionID() []int32`

GetAPkiPermissionID returns the APkiPermissionID field if non-nil, zero value otherwise.

### GetAPkiPermissionIDOk

`func (o *ActivesessionResponseCompound) GetAPkiPermissionIDOk() (*[]int32, bool)`

GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiPermissionID

`func (o *ActivesessionResponseCompound) SetAPkiPermissionID(v []int32)`

SetAPkiPermissionID sets APkiPermissionID field to given value.


### GetObjUserReal

`func (o *ActivesessionResponseCompound) GetObjUserReal() ActivesessionResponseCompoundUser`

GetObjUserReal returns the ObjUserReal field if non-nil, zero value otherwise.

### GetObjUserRealOk

`func (o *ActivesessionResponseCompound) GetObjUserRealOk() (*ActivesessionResponseCompoundUser, bool)`

GetObjUserRealOk returns a tuple with the ObjUserReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserReal

`func (o *ActivesessionResponseCompound) SetObjUserReal(v ActivesessionResponseCompoundUser)`

SetObjUserReal sets ObjUserReal field to given value.


### GetObjUserCloned

`func (o *ActivesessionResponseCompound) GetObjUserCloned() ActivesessionResponseCompoundUser`

GetObjUserCloned returns the ObjUserCloned field if non-nil, zero value otherwise.

### GetObjUserClonedOk

`func (o *ActivesessionResponseCompound) GetObjUserClonedOk() (*ActivesessionResponseCompoundUser, bool)`

GetObjUserClonedOk returns a tuple with the ObjUserCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserCloned

`func (o *ActivesessionResponseCompound) SetObjUserCloned(v ActivesessionResponseCompoundUser)`

SetObjUserCloned sets ObjUserCloned field to given value.

### HasObjUserCloned

`func (o *ActivesessionResponseCompound) HasObjUserCloned() bool`

HasObjUserCloned returns a boolean if a field has been set.

### GetObjApikey

`func (o *ActivesessionResponseCompound) GetObjApikey() ActivesessionResponseCompoundApikey`

GetObjApikey returns the ObjApikey field if non-nil, zero value otherwise.

### GetObjApikeyOk

`func (o *ActivesessionResponseCompound) GetObjApikeyOk() (*ActivesessionResponseCompoundApikey, bool)`

GetObjApikeyOk returns a tuple with the ObjApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikey

`func (o *ActivesessionResponseCompound) SetObjApikey(v ActivesessionResponseCompoundApikey)`

SetObjApikey sets ObjApikey field to given value.

### HasObjApikey

`func (o *ActivesessionResponseCompound) HasObjApikey() bool`

HasObjApikey returns a boolean if a field has been set.

### GetAEModuleInternalname

`func (o *ActivesessionResponseCompound) GetAEModuleInternalname() []string`

GetAEModuleInternalname returns the AEModuleInternalname field if non-nil, zero value otherwise.

### GetAEModuleInternalnameOk

`func (o *ActivesessionResponseCompound) GetAEModuleInternalnameOk() (*[]string, bool)`

GetAEModuleInternalnameOk returns a tuple with the AEModuleInternalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAEModuleInternalname

`func (o *ActivesessionResponseCompound) SetAEModuleInternalname(v []string)`

SetAEModuleInternalname sets AEModuleInternalname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


