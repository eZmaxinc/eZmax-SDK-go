# ActivesessionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EActivesessionUsertype** | [**FieldEActivesessionUsertype**](FieldEActivesessionUsertype.md) |  | 
**EActivesessionOrigin** | [**FieldEActivesessionOrigin**](FieldEActivesessionOrigin.md) |  | 
**EActivesessionWeekdaystart** | [**FieldEActivesessionWeekdaystart**](FieldEActivesessionWeekdaystart.md) |  | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SCompanyNameX** | **string** | The Name of the Company in the language of the requester | 
**SDepartmentNameX** | **string** | The Name of the Department in the language of the requester | 
**BActivesessionDebug** | **bool** | Whether the active session is in debug or not | 
**BActivesessionIssuperadmin** | **bool** | Whether the active session is superadmin or not | 
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**FkiSignatureID** | Pointer to **int32** | The unique ID of the Signature | [optional] 
**BSystemconfigurationEzsignpaidbyoffice** | Pointer to **bool** | Whether if Ezsign is paid by the company or not | [optional] 
**ESystemconfigurationEzsignofficeplan** | Pointer to [**FieldESystemconfigurationEzsignofficeplan**](FieldESystemconfigurationEzsignofficeplan.md) |  | [optional] 
**EUserEzsignaccess** | [**FieldEUserEzsignaccess**](FieldEUserEzsignaccess.md) |  | 
**EUserEzsignprepaid** | Pointer to [**FieldEUserEzsignprepaid**](FieldEUserEzsignprepaid.md) |  | [optional] 
**DtUserEzsignprepaidexpiration** | Pointer to **string** | The eZsign prepaid expiration date | [optional] 
**APkiPermissionID** | **[]int32** | An array of permissions granted to the user or api key | 
**ObjUserReal** | [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | 
**ObjUserCloned** | Pointer to [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | [optional] 
**ObjApikey** | Pointer to [**ActivesessionResponseCompoundApikey**](ActivesessionResponseCompoundApikey.md) |  | [optional] 
**AEModuleInternalname** | **[]string** | An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key. | 

## Methods

### NewActivesessionResponseCompound

`func NewActivesessionResponseCompound(eActivesessionUsertype FieldEActivesessionUsertype, eActivesessionOrigin FieldEActivesessionOrigin, eActivesessionWeekdaystart FieldEActivesessionWeekdaystart, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, bActivesessionDebug bool, bActivesessionIssuperadmin bool, pksCustomerCode string, fkiSystemconfigurationtypeID int32, eUserEzsignaccess FieldEUserEzsignaccess, aPkiPermissionID []int32, objUserReal ActivesessionResponseCompoundUser, aEModuleInternalname []string, ) *ActivesessionResponseCompound`

NewActivesessionResponseCompound instantiates a new ActivesessionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionResponseCompoundWithDefaults

`func NewActivesessionResponseCompoundWithDefaults() *ActivesessionResponseCompound`

NewActivesessionResponseCompoundWithDefaults instantiates a new ActivesessionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEActivesessionUsertype

`func (o *ActivesessionResponseCompound) GetEActivesessionUsertype() FieldEActivesessionUsertype`

GetEActivesessionUsertype returns the EActivesessionUsertype field if non-nil, zero value otherwise.

### GetEActivesessionUsertypeOk

`func (o *ActivesessionResponseCompound) GetEActivesessionUsertypeOk() (*FieldEActivesessionUsertype, bool)`

GetEActivesessionUsertypeOk returns a tuple with the EActivesessionUsertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionUsertype

`func (o *ActivesessionResponseCompound) SetEActivesessionUsertype(v FieldEActivesessionUsertype)`

SetEActivesessionUsertype sets EActivesessionUsertype field to given value.


### GetEActivesessionOrigin

`func (o *ActivesessionResponseCompound) GetEActivesessionOrigin() FieldEActivesessionOrigin`

GetEActivesessionOrigin returns the EActivesessionOrigin field if non-nil, zero value otherwise.

### GetEActivesessionOriginOk

`func (o *ActivesessionResponseCompound) GetEActivesessionOriginOk() (*FieldEActivesessionOrigin, bool)`

GetEActivesessionOriginOk returns a tuple with the EActivesessionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionOrigin

`func (o *ActivesessionResponseCompound) SetEActivesessionOrigin(v FieldEActivesessionOrigin)`

SetEActivesessionOrigin sets EActivesessionOrigin field to given value.


### GetEActivesessionWeekdaystart

`func (o *ActivesessionResponseCompound) GetEActivesessionWeekdaystart() FieldEActivesessionWeekdaystart`

GetEActivesessionWeekdaystart returns the EActivesessionWeekdaystart field if non-nil, zero value otherwise.

### GetEActivesessionWeekdaystartOk

`func (o *ActivesessionResponseCompound) GetEActivesessionWeekdaystartOk() (*FieldEActivesessionWeekdaystart, bool)`

GetEActivesessionWeekdaystartOk returns a tuple with the EActivesessionWeekdaystart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionWeekdaystart

`func (o *ActivesessionResponseCompound) SetEActivesessionWeekdaystart(v FieldEActivesessionWeekdaystart)`

SetEActivesessionWeekdaystart sets EActivesessionWeekdaystart field to given value.


### GetFkiLanguageID

`func (o *ActivesessionResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ActivesessionResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ActivesessionResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSCompanyNameX

`func (o *ActivesessionResponseCompound) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *ActivesessionResponseCompound) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *ActivesessionResponseCompound) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetSDepartmentNameX

`func (o *ActivesessionResponseCompound) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *ActivesessionResponseCompound) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *ActivesessionResponseCompound) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetBActivesessionDebug

`func (o *ActivesessionResponseCompound) GetBActivesessionDebug() bool`

GetBActivesessionDebug returns the BActivesessionDebug field if non-nil, zero value otherwise.

### GetBActivesessionDebugOk

`func (o *ActivesessionResponseCompound) GetBActivesessionDebugOk() (*bool, bool)`

GetBActivesessionDebugOk returns a tuple with the BActivesessionDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionDebug

`func (o *ActivesessionResponseCompound) SetBActivesessionDebug(v bool)`

SetBActivesessionDebug sets BActivesessionDebug field to given value.


### GetBActivesessionIssuperadmin

`func (o *ActivesessionResponseCompound) GetBActivesessionIssuperadmin() bool`

GetBActivesessionIssuperadmin returns the BActivesessionIssuperadmin field if non-nil, zero value otherwise.

### GetBActivesessionIssuperadminOk

`func (o *ActivesessionResponseCompound) GetBActivesessionIssuperadminOk() (*bool, bool)`

GetBActivesessionIssuperadminOk returns a tuple with the BActivesessionIssuperadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionIssuperadmin

`func (o *ActivesessionResponseCompound) SetBActivesessionIssuperadmin(v bool)`

SetBActivesessionIssuperadmin sets BActivesessionIssuperadmin field to given value.


### GetPksCustomerCode

`func (o *ActivesessionResponseCompound) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *ActivesessionResponseCompound) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *ActivesessionResponseCompound) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *ActivesessionResponseCompound) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *ActivesessionResponseCompound) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *ActivesessionResponseCompound) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetFkiSignatureID

`func (o *ActivesessionResponseCompound) GetFkiSignatureID() int32`

GetFkiSignatureID returns the FkiSignatureID field if non-nil, zero value otherwise.

### GetFkiSignatureIDOk

`func (o *ActivesessionResponseCompound) GetFkiSignatureIDOk() (*int32, bool)`

GetFkiSignatureIDOk returns a tuple with the FkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSignatureID

`func (o *ActivesessionResponseCompound) SetFkiSignatureID(v int32)`

SetFkiSignatureID sets FkiSignatureID field to given value.

### HasFkiSignatureID

`func (o *ActivesessionResponseCompound) HasFkiSignatureID() bool`

HasFkiSignatureID returns a boolean if a field has been set.

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


