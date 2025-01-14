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
**BActivesessionAttachment** | Pointer to **bool** | Can access attachment when we clone a user | [optional] 
**BActivesessionCanafe** | Pointer to **bool** | Can access canafe when we clone a user | [optional] 
**BActivesessionFinancial** | Pointer to **bool** | Can access financial element when we clone a user | [optional] 
**BActivesessionRealestatecompleted** | Pointer to **bool** | Can access closed realestate folders when we clone a user | [optional] 
**EActivesessionEzsign** | Pointer to [**FieldEActivesessionEzsign**](FieldEActivesessionEzsign.md) |  | [optional] 
**EActivesessionEzsignaccess** | [**FieldEActivesessionEzsignaccess**](FieldEActivesessionEzsignaccess.md) |  | 
**EActivesessionEzsignprepaid** | Pointer to [**FieldEActivesessionEzsignprepaid**](FieldEActivesessionEzsignprepaid.md) |  | [optional] 
**EActivesessionRealestateinprogress** | Pointer to [**FieldEActivesessionRealestateinprogress**](FieldEActivesessionRealestateinprogress.md) |  | [optional] 
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**FkiSignatureID** | Pointer to **int32** | The unique ID of the Signature | [optional] 
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

### NewActivesessionGetCurrentV1ResponseMPayload

`func NewActivesessionGetCurrentV1ResponseMPayload(eActivesessionUsertype FieldEActivesessionUsertype, eActivesessionOrigin FieldEActivesessionOrigin, eActivesessionWeekdaystart FieldEActivesessionWeekdaystart, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, bActivesessionDebug bool, bActivesessionIssuperadmin bool, eActivesessionEzsignaccess FieldEActivesessionEzsignaccess, pksCustomerCode string, fkiSystemconfigurationtypeID int32, eUserEzsignaccess FieldEUserEzsignaccess, aPkiPermissionID []int32, objUserReal ActivesessionResponseCompoundUser, aEModuleInternalname []string, ) *ActivesessionGetCurrentV1ResponseMPayload`

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


### GetBActivesessionAttachment

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionAttachment() bool`

GetBActivesessionAttachment returns the BActivesessionAttachment field if non-nil, zero value otherwise.

### GetBActivesessionAttachmentOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionAttachmentOk() (*bool, bool)`

GetBActivesessionAttachmentOk returns a tuple with the BActivesessionAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionAttachment

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionAttachment(v bool)`

SetBActivesessionAttachment sets BActivesessionAttachment field to given value.

### HasBActivesessionAttachment

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBActivesessionAttachment() bool`

HasBActivesessionAttachment returns a boolean if a field has been set.

### GetBActivesessionCanafe

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionCanafe() bool`

GetBActivesessionCanafe returns the BActivesessionCanafe field if non-nil, zero value otherwise.

### GetBActivesessionCanafeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionCanafeOk() (*bool, bool)`

GetBActivesessionCanafeOk returns a tuple with the BActivesessionCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionCanafe

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionCanafe(v bool)`

SetBActivesessionCanafe sets BActivesessionCanafe field to given value.

### HasBActivesessionCanafe

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBActivesessionCanafe() bool`

HasBActivesessionCanafe returns a boolean if a field has been set.

### GetBActivesessionFinancial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionFinancial() bool`

GetBActivesessionFinancial returns the BActivesessionFinancial field if non-nil, zero value otherwise.

### GetBActivesessionFinancialOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionFinancialOk() (*bool, bool)`

GetBActivesessionFinancialOk returns a tuple with the BActivesessionFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionFinancial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionFinancial(v bool)`

SetBActivesessionFinancial sets BActivesessionFinancial field to given value.

### HasBActivesessionFinancial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBActivesessionFinancial() bool`

HasBActivesessionFinancial returns a boolean if a field has been set.

### GetBActivesessionRealestatecompleted

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionRealestatecompleted() bool`

GetBActivesessionRealestatecompleted returns the BActivesessionRealestatecompleted field if non-nil, zero value otherwise.

### GetBActivesessionRealestatecompletedOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBActivesessionRealestatecompletedOk() (*bool, bool)`

GetBActivesessionRealestatecompletedOk returns a tuple with the BActivesessionRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionRealestatecompleted

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBActivesessionRealestatecompleted(v bool)`

SetBActivesessionRealestatecompleted sets BActivesessionRealestatecompleted field to given value.

### HasBActivesessionRealestatecompleted

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBActivesessionRealestatecompleted() bool`

HasBActivesessionRealestatecompleted returns a boolean if a field has been set.

### GetEActivesessionEzsign

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsign() FieldEActivesessionEzsign`

GetEActivesessionEzsign returns the EActivesessionEzsign field if non-nil, zero value otherwise.

### GetEActivesessionEzsignOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsignOk() (*FieldEActivesessionEzsign, bool)`

GetEActivesessionEzsignOk returns a tuple with the EActivesessionEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsign

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionEzsign(v FieldEActivesessionEzsign)`

SetEActivesessionEzsign sets EActivesessionEzsign field to given value.

### HasEActivesessionEzsign

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasEActivesessionEzsign() bool`

HasEActivesessionEzsign returns a boolean if a field has been set.

### GetEActivesessionEzsignaccess

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsignaccess() FieldEActivesessionEzsignaccess`

GetEActivesessionEzsignaccess returns the EActivesessionEzsignaccess field if non-nil, zero value otherwise.

### GetEActivesessionEzsignaccessOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsignaccessOk() (*FieldEActivesessionEzsignaccess, bool)`

GetEActivesessionEzsignaccessOk returns a tuple with the EActivesessionEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsignaccess

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionEzsignaccess(v FieldEActivesessionEzsignaccess)`

SetEActivesessionEzsignaccess sets EActivesessionEzsignaccess field to given value.


### GetEActivesessionEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsignprepaid() FieldEActivesessionEzsignprepaid`

GetEActivesessionEzsignprepaid returns the EActivesessionEzsignprepaid field if non-nil, zero value otherwise.

### GetEActivesessionEzsignprepaidOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionEzsignprepaidOk() (*FieldEActivesessionEzsignprepaid, bool)`

GetEActivesessionEzsignprepaidOk returns a tuple with the EActivesessionEzsignprepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionEzsignprepaid(v FieldEActivesessionEzsignprepaid)`

SetEActivesessionEzsignprepaid sets EActivesessionEzsignprepaid field to given value.

### HasEActivesessionEzsignprepaid

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasEActivesessionEzsignprepaid() bool`

HasEActivesessionEzsignprepaid returns a boolean if a field has been set.

### GetEActivesessionRealestateinprogress

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionRealestateinprogress() FieldEActivesessionRealestateinprogress`

GetEActivesessionRealestateinprogress returns the EActivesessionRealestateinprogress field if non-nil, zero value otherwise.

### GetEActivesessionRealestateinprogressOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetEActivesessionRealestateinprogressOk() (*FieldEActivesessionRealestateinprogress, bool)`

GetEActivesessionRealestateinprogressOk returns a tuple with the EActivesessionRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionRealestateinprogress

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetEActivesessionRealestateinprogress(v FieldEActivesessionRealestateinprogress)`

SetEActivesessionRealestateinprogress sets EActivesessionRealestateinprogress field to given value.

### HasEActivesessionRealestateinprogress

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasEActivesessionRealestateinprogress() bool`

HasEActivesessionRealestateinprogress returns a boolean if a field has been set.

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

### GetFkiEzsignuserID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiEzsignuserID() int32`

GetFkiEzsignuserID returns the FkiEzsignuserID field if non-nil, zero value otherwise.

### GetFkiEzsignuserIDOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetFkiEzsignuserIDOk() (*int32, bool)`

GetFkiEzsignuserIDOk returns a tuple with the FkiEzsignuserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignuserID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetFkiEzsignuserID(v int32)`

SetFkiEzsignuserID sets FkiEzsignuserID field to given value.

### HasFkiEzsignuserID

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasFkiEzsignuserID() bool`

HasFkiEzsignuserID returns a boolean if a field has been set.

### GetBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBSystemconfigurationEzsignpaidbyoffice() bool`

GetBSystemconfigurationEzsignpaidbyoffice returns the BSystemconfigurationEzsignpaidbyoffice field if non-nil, zero value otherwise.

### GetBSystemconfigurationEzsignpaidbyofficeOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBSystemconfigurationEzsignpaidbyofficeOk() (*bool, bool)`

GetBSystemconfigurationEzsignpaidbyofficeOk returns a tuple with the BSystemconfigurationEzsignpaidbyoffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBSystemconfigurationEzsignpaidbyoffice(v bool)`

SetBSystemconfigurationEzsignpaidbyoffice sets BSystemconfigurationEzsignpaidbyoffice field to given value.

### HasBSystemconfigurationEzsignpaidbyoffice

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBSystemconfigurationEzsignpaidbyoffice() bool`

HasBSystemconfigurationEzsignpaidbyoffice returns a boolean if a field has been set.

### GetESystemconfigurationEzsignofficeplan

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetESystemconfigurationEzsignofficeplan() FieldESystemconfigurationEzsignofficeplan`

GetESystemconfigurationEzsignofficeplan returns the ESystemconfigurationEzsignofficeplan field if non-nil, zero value otherwise.

### GetESystemconfigurationEzsignofficeplanOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetESystemconfigurationEzsignofficeplanOk() (*FieldESystemconfigurationEzsignofficeplan, bool)`

GetESystemconfigurationEzsignofficeplanOk returns a tuple with the ESystemconfigurationEzsignofficeplan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESystemconfigurationEzsignofficeplan

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetESystemconfigurationEzsignofficeplan(v FieldESystemconfigurationEzsignofficeplan)`

SetESystemconfigurationEzsignofficeplan sets ESystemconfigurationEzsignofficeplan field to given value.

### HasESystemconfigurationEzsignofficeplan

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasESystemconfigurationEzsignofficeplan() bool`

HasESystemconfigurationEzsignofficeplan returns a boolean if a field has been set.

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

### GetBUserEzsigntrial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBUserEzsigntrial() bool`

GetBUserEzsigntrial returns the BUserEzsigntrial field if non-nil, zero value otherwise.

### GetBUserEzsigntrialOk

`func (o *ActivesessionGetCurrentV1ResponseMPayload) GetBUserEzsigntrialOk() (*bool, bool)`

GetBUserEzsigntrialOk returns a tuple with the BUserEzsigntrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserEzsigntrial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) SetBUserEzsigntrial(v bool)`

SetBUserEzsigntrial sets BUserEzsigntrial field to given value.

### HasBUserEzsigntrial

`func (o *ActivesessionGetCurrentV1ResponseMPayload) HasBUserEzsigntrial() bool`

HasBUserEzsigntrial returns a boolean if a field has been set.

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


