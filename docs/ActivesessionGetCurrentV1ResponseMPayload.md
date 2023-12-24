# ActivesessionGetCurrentV1ResponseMPayload

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
**EUserEzsignaccess** | [**FieldEUserEzsignaccess**](FieldEUserEzsignaccess.md) |  | 
**EUserEzsignprepaid** | Pointer to [**FieldEUserEzsignprepaid**](FieldEUserEzsignprepaid.md) |  | [optional] 
**DtUserEzsignprepaidexpiration** | Pointer to **string** | The eZsign prepaid expiration date | [optional] 
**APkiPermissionID** | **[]int32** | An array of permissions granted to the user or api key | 
**ObjUserReal** | [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | 
**ObjUserCloned** | Pointer to [**ActivesessionResponseCompoundUser**](ActivesessionResponseCompoundUser.md) |  | [optional] 
**ObjApikey** | Pointer to [**ActivesessionResponseCompoundApikey**](ActivesessionResponseCompoundApikey.md) |  | [optional] 
**AEModuleInternalname** | **[]string** | An Array of Registered modules.  These are the modules that are Licensed to be used by the User or the API Key. | 

## Methods

### NewActivesessionGetCurrentV1ResponseMPayload

`func NewActivesessionGetCurrentV1ResponseMPayload(eActivesessionUsertype FieldEActivesessionUsertype, eActivesessionOrigin FieldEActivesessionOrigin, eActivesessionWeekdaystart FieldEActivesessionWeekdaystart, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, bActivesessionDebug bool, bActivesessionIssuperadmin bool, pksCustomerCode string, fkiSystemconfigurationtypeID int32, eUserEzsignaccess FieldEUserEzsignaccess, aPkiPermissionID []int32, objUserReal ActivesessionResponseCompoundUser, aEModuleInternalname []string, ) *ActivesessionGetCurrentV1ResponseMPayload`

NewActivesessionGetCurrentV1ResponseMPayload instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults

`func NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults() *ActivesessionGetCurrentV1ResponseMPayload`

NewActivesessionGetCurrentV1ResponseMPayloadWithDefaults instantiates a new ActivesessionGetCurrentV1ResponseMPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEActivesessionUsertype

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionUsertype() FieldEActivesessionUsertype`

GetEActivesessionUsertype returns the EActivesessionUsertype field if non-nil, zero value otherwise.

### GetEActivesessionUsertypeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionUsertypeOk() (*FieldEActivesessionUsertype, bool)`

GetEActivesessionUsertypeOk returns a tuple with the EActivesessionUsertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionUsertype

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionUsertype(v FieldEActivesessionUsertype)`

SetEActivesessionUsertype sets EActivesessionUsertype field to given value.


### GetEActivesessionOrigin

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionOrigin() FieldEActivesessionOrigin`

GetEActivesessionOrigin returns the EActivesessionOrigin field if non-nil, zero value otherwise.

### GetEActivesessionOriginOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionOriginOk() (*FieldEActivesessionOrigin, bool)`

GetEActivesessionOriginOk returns a tuple with the EActivesessionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionOrigin

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionOrigin(v FieldEActivesessionOrigin)`

SetEActivesessionOrigin sets EActivesessionOrigin field to given value.


### GetEActivesessionWeekdaystart

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionWeekdaystart() FieldEActivesessionWeekdaystart`

GetEActivesessionWeekdaystart returns the EActivesessionWeekdaystart field if non-nil, zero value otherwise.

### GetEActivesessionWeekdaystartOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionWeekdaystartOk() (*FieldEActivesessionWeekdaystart, bool)`

GetEActivesessionWeekdaystartOk returns a tuple with the EActivesessionWeekdaystart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionWeekdaystart

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionWeekdaystart(v FieldEActivesessionWeekdaystart)`

SetEActivesessionWeekdaystart sets EActivesessionWeekdaystart field to given value.


### GetFkiLanguageID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSCompanyNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetSDepartmentNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetBActivesessionDebug

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionDebug() bool`

GetBActivesessionDebug returns the BActivesessionDebug field if non-nil, zero value otherwise.

### GetBActivesessionDebugOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionDebugOk() (*bool, bool)`

GetBActivesessionDebugOk returns a tuple with the BActivesessionDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionDebug

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionDebug(v bool)`

SetBActivesessionDebug sets BActivesessionDebug field to given value.


### GetBActivesessionIssuperadmin

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionIssuperadmin() bool`

GetBActivesessionIssuperadmin returns the BActivesessionIssuperadmin field if non-nil, zero value otherwise.

### GetBActivesessionIssuperadminOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionIssuperadminOk() (*bool, bool)`

GetBActivesessionIssuperadminOk returns a tuple with the BActivesessionIssuperadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionIssuperadmin

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionIssuperadmin(v bool)`

SetBActivesessionIssuperadmin sets BActivesessionIssuperadmin field to given value.


### GetPksCustomerCode

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetFkiSignatureID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiSignatureID() int32`

GetFkiSignatureID returns the FkiSignatureID field if non-nil, zero value otherwise.

### GetFkiSignatureIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiSignatureIDOk() (*int32, bool)`

GetFkiSignatureIDOk returns a tuple with the FkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSignatureID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiSignatureID(v int32)`

SetFkiSignatureID sets FkiSignatureID field to given value.

### HasFkiSignatureID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasFkiSignatureID() bool`

HasFkiSignatureID returns a boolean if a field has been set.

### GetEUserEzsignaccess

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEUserEzsignaccess() FieldEUserEzsignaccess`

GetEUserEzsignaccess returns the EUserEzsignaccess field if non-nil, zero value otherwise.

### GetEUserEzsignaccessOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool)`

GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignaccess

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEUserEzsignaccess(v FieldEUserEzsignaccess)`

SetEUserEzsignaccess sets EUserEzsignaccess field to given value.


### GetEUserEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEUserEzsignprepaid() FieldEUserEzsignprepaid`

GetEUserEzsignprepaid returns the EUserEzsignprepaid field if non-nil, zero value otherwise.

### GetEUserEzsignprepaidOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEUserEzsignprepaidOk() (*FieldEUserEzsignprepaid, bool)`

GetEUserEzsignprepaidOk returns a tuple with the EUserEzsignprepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEUserEzsignprepaid(v FieldEUserEzsignprepaid)`

SetEUserEzsignprepaid sets EUserEzsignprepaid field to given value.

### HasEUserEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasEUserEzsignprepaid() bool`

HasEUserEzsignprepaid returns a boolean if a field has been set.

### GetDtUserEzsignprepaidexpiration

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetDtUserEzsignprepaidexpiration() string`

GetDtUserEzsignprepaidexpiration returns the DtUserEzsignprepaidexpiration field if non-nil, zero value otherwise.

### GetDtUserEzsignprepaidexpirationOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetDtUserEzsignprepaidexpirationOk() (*string, bool)`

GetDtUserEzsignprepaidexpirationOk returns a tuple with the DtUserEzsignprepaidexpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtUserEzsignprepaidexpiration

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetDtUserEzsignprepaidexpiration(v string)`

SetDtUserEzsignprepaidexpiration sets DtUserEzsignprepaidexpiration field to given value.

### HasDtUserEzsignprepaidexpiration

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasDtUserEzsignprepaidexpiration() bool`

HasDtUserEzsignprepaidexpiration returns a boolean if a field has been set.

### GetAPkiPermissionID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPkiPermissionID() []int32`

GetAPkiPermissionID returns the APkiPermissionID field if non-nil, zero value otherwise.

### GetAPkiPermissionIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAPkiPermissionIDOk() (*[]int32, bool)`

GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiPermissionID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetAPkiPermissionID(v []int32)`

SetAPkiPermissionID sets APkiPermissionID field to given value.


### GetObjUserReal

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjUserReal() ActivesessionResponseCompoundUser`

GetObjUserReal returns the ObjUserReal field if non-nil, zero value otherwise.

### GetObjUserRealOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjUserRealOk() (*ActivesessionResponseCompoundUser, bool)`

GetObjUserRealOk returns a tuple with the ObjUserReal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserReal

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetObjUserReal(v ActivesessionResponseCompoundUser)`

SetObjUserReal sets ObjUserReal field to given value.


### GetObjUserCloned

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjUserCloned() ActivesessionResponseCompoundUser`

GetObjUserCloned returns the ObjUserCloned field if non-nil, zero value otherwise.

### GetObjUserClonedOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjUserClonedOk() (*ActivesessionResponseCompoundUser, bool)`

GetObjUserClonedOk returns a tuple with the ObjUserCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjUserCloned

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetObjUserCloned(v ActivesessionResponseCompoundUser)`

SetObjUserCloned sets ObjUserCloned field to given value.

### HasObjUserCloned

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasObjUserCloned() bool`

HasObjUserCloned returns a boolean if a field has been set.

### GetObjApikey

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjApikey() ActivesessionResponseCompoundApikey`

GetObjApikey returns the ObjApikey field if non-nil, zero value otherwise.

### GetObjApikeyOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetObjApikeyOk() (*ActivesessionResponseCompoundApikey, bool)`

GetObjApikeyOk returns a tuple with the ObjApikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjApikey

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetObjApikey(v ActivesessionResponseCompoundApikey)`

SetObjApikey sets ObjApikey field to given value.

### HasObjApikey

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasObjApikey() bool`

HasObjApikey returns a boolean if a field has been set.

### GetAEModuleInternalname

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAEModuleInternalname() []string`

GetAEModuleInternalname returns the AEModuleInternalname field if non-nil, zero value otherwise.

### GetAEModuleInternalnameOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetAEModuleInternalnameOk() (*[]string, bool)`

GetAEModuleInternalnameOk returns a tuple with the AEModuleInternalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAEModuleInternalname

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetAEModuleInternalname(v []string)`

SetAEModuleInternalname sets AEModuleInternalname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


