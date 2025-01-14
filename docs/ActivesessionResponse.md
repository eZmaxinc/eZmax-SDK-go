# ActivesessionResponse

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

## Methods

### NewActivesessionResponse

`func NewActivesessionResponse(eActivesessionUsertype FieldEActivesessionUsertype, eActivesessionOrigin FieldEActivesessionOrigin, eActivesessionWeekdaystart FieldEActivesessionWeekdaystart, fkiLanguageID int32, sCompanyNameX string, sDepartmentNameX string, bActivesessionDebug bool, bActivesessionIssuperadmin bool, eActivesessionEzsignaccess FieldEActivesessionEzsignaccess, pksCustomerCode string, fkiSystemconfigurationtypeID int32, ) *ActivesessionResponse`

NewActivesessionResponse instantiates a new ActivesessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivesessionResponseWithDefaults

`func NewActivesessionResponseWithDefaults() *ActivesessionResponse`

NewActivesessionResponseWithDefaults instantiates a new ActivesessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEActivesessionUsertype

`func (o *ActivesessionResponse) GetEActivesessionUsertype() FieldEActivesessionUsertype`

GetEActivesessionUsertype returns the EActivesessionUsertype field if non-nil, zero value otherwise.

### GetEActivesessionUsertypeOk

`func (o *ActivesessionResponse) GetEActivesessionUsertypeOk() (*FieldEActivesessionUsertype, bool)`

GetEActivesessionUsertypeOk returns a tuple with the EActivesessionUsertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionUsertype

`func (o *ActivesessionResponse) SetEActivesessionUsertype(v FieldEActivesessionUsertype)`

SetEActivesessionUsertype sets EActivesessionUsertype field to given value.


### GetEActivesessionOrigin

`func (o *ActivesessionResponse) GetEActivesessionOrigin() FieldEActivesessionOrigin`

GetEActivesessionOrigin returns the EActivesessionOrigin field if non-nil, zero value otherwise.

### GetEActivesessionOriginOk

`func (o *ActivesessionResponse) GetEActivesessionOriginOk() (*FieldEActivesessionOrigin, bool)`

GetEActivesessionOriginOk returns a tuple with the EActivesessionOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionOrigin

`func (o *ActivesessionResponse) SetEActivesessionOrigin(v FieldEActivesessionOrigin)`

SetEActivesessionOrigin sets EActivesessionOrigin field to given value.


### GetEActivesessionWeekdaystart

`func (o *ActivesessionResponse) GetEActivesessionWeekdaystart() FieldEActivesessionWeekdaystart`

GetEActivesessionWeekdaystart returns the EActivesessionWeekdaystart field if non-nil, zero value otherwise.

### GetEActivesessionWeekdaystartOk

`func (o *ActivesessionResponse) GetEActivesessionWeekdaystartOk() (*FieldEActivesessionWeekdaystart, bool)`

GetEActivesessionWeekdaystartOk returns a tuple with the EActivesessionWeekdaystart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionWeekdaystart

`func (o *ActivesessionResponse) SetEActivesessionWeekdaystart(v FieldEActivesessionWeekdaystart)`

SetEActivesessionWeekdaystart sets EActivesessionWeekdaystart field to given value.


### GetFkiLanguageID

`func (o *ActivesessionResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *ActivesessionResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *ActivesessionResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSCompanyNameX

`func (o *ActivesessionResponse) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *ActivesessionResponse) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *ActivesessionResponse) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetSDepartmentNameX

`func (o *ActivesessionResponse) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *ActivesessionResponse) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *ActivesessionResponse) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetBActivesessionDebug

`func (o *ActivesessionResponse) GetBActivesessionDebug() bool`

GetBActivesessionDebug returns the BActivesessionDebug field if non-nil, zero value otherwise.

### GetBActivesessionDebugOk

`func (o *ActivesessionResponse) GetBActivesessionDebugOk() (*bool, bool)`

GetBActivesessionDebugOk returns a tuple with the BActivesessionDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionDebug

`func (o *ActivesessionResponse) SetBActivesessionDebug(v bool)`

SetBActivesessionDebug sets BActivesessionDebug field to given value.


### GetBActivesessionIssuperadmin

`func (o *ActivesessionResponse) GetBActivesessionIssuperadmin() bool`

GetBActivesessionIssuperadmin returns the BActivesessionIssuperadmin field if non-nil, zero value otherwise.

### GetBActivesessionIssuperadminOk

`func (o *ActivesessionResponse) GetBActivesessionIssuperadminOk() (*bool, bool)`

GetBActivesessionIssuperadminOk returns a tuple with the BActivesessionIssuperadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionIssuperadmin

`func (o *ActivesessionResponse) SetBActivesessionIssuperadmin(v bool)`

SetBActivesessionIssuperadmin sets BActivesessionIssuperadmin field to given value.


### GetBActivesessionAttachment

`func (o *ActivesessionResponse) GetBActivesessionAttachment() bool`

GetBActivesessionAttachment returns the BActivesessionAttachment field if non-nil, zero value otherwise.

### GetBActivesessionAttachmentOk

`func (o *ActivesessionResponse) GetBActivesessionAttachmentOk() (*bool, bool)`

GetBActivesessionAttachmentOk returns a tuple with the BActivesessionAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionAttachment

`func (o *ActivesessionResponse) SetBActivesessionAttachment(v bool)`

SetBActivesessionAttachment sets BActivesessionAttachment field to given value.

### HasBActivesessionAttachment

`func (o *ActivesessionResponse) HasBActivesessionAttachment() bool`

HasBActivesessionAttachment returns a boolean if a field has been set.

### GetBActivesessionCanafe

`func (o *ActivesessionResponse) GetBActivesessionCanafe() bool`

GetBActivesessionCanafe returns the BActivesessionCanafe field if non-nil, zero value otherwise.

### GetBActivesessionCanafeOk

`func (o *ActivesessionResponse) GetBActivesessionCanafeOk() (*bool, bool)`

GetBActivesessionCanafeOk returns a tuple with the BActivesessionCanafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionCanafe

`func (o *ActivesessionResponse) SetBActivesessionCanafe(v bool)`

SetBActivesessionCanafe sets BActivesessionCanafe field to given value.

### HasBActivesessionCanafe

`func (o *ActivesessionResponse) HasBActivesessionCanafe() bool`

HasBActivesessionCanafe returns a boolean if a field has been set.

### GetBActivesessionFinancial

`func (o *ActivesessionResponse) GetBActivesessionFinancial() bool`

GetBActivesessionFinancial returns the BActivesessionFinancial field if non-nil, zero value otherwise.

### GetBActivesessionFinancialOk

`func (o *ActivesessionResponse) GetBActivesessionFinancialOk() (*bool, bool)`

GetBActivesessionFinancialOk returns a tuple with the BActivesessionFinancial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionFinancial

`func (o *ActivesessionResponse) SetBActivesessionFinancial(v bool)`

SetBActivesessionFinancial sets BActivesessionFinancial field to given value.

### HasBActivesessionFinancial

`func (o *ActivesessionResponse) HasBActivesessionFinancial() bool`

HasBActivesessionFinancial returns a boolean if a field has been set.

### GetBActivesessionRealestatecompleted

`func (o *ActivesessionResponse) GetBActivesessionRealestatecompleted() bool`

GetBActivesessionRealestatecompleted returns the BActivesessionRealestatecompleted field if non-nil, zero value otherwise.

### GetBActivesessionRealestatecompletedOk

`func (o *ActivesessionResponse) GetBActivesessionRealestatecompletedOk() (*bool, bool)`

GetBActivesessionRealestatecompletedOk returns a tuple with the BActivesessionRealestatecompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBActivesessionRealestatecompleted

`func (o *ActivesessionResponse) SetBActivesessionRealestatecompleted(v bool)`

SetBActivesessionRealestatecompleted sets BActivesessionRealestatecompleted field to given value.

### HasBActivesessionRealestatecompleted

`func (o *ActivesessionResponse) HasBActivesessionRealestatecompleted() bool`

HasBActivesessionRealestatecompleted returns a boolean if a field has been set.

### GetEActivesessionEzsign

`func (o *ActivesessionResponse) GetEActivesessionEzsign() FieldEActivesessionEzsign`

GetEActivesessionEzsign returns the EActivesessionEzsign field if non-nil, zero value otherwise.

### GetEActivesessionEzsignOk

`func (o *ActivesessionResponse) GetEActivesessionEzsignOk() (*FieldEActivesessionEzsign, bool)`

GetEActivesessionEzsignOk returns a tuple with the EActivesessionEzsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsign

`func (o *ActivesessionResponse) SetEActivesessionEzsign(v FieldEActivesessionEzsign)`

SetEActivesessionEzsign sets EActivesessionEzsign field to given value.

### HasEActivesessionEzsign

`func (o *ActivesessionResponse) HasEActivesessionEzsign() bool`

HasEActivesessionEzsign returns a boolean if a field has been set.

### GetEActivesessionEzsignaccess

`func (o *ActivesessionResponse) GetEActivesessionEzsignaccess() FieldEActivesessionEzsignaccess`

GetEActivesessionEzsignaccess returns the EActivesessionEzsignaccess field if non-nil, zero value otherwise.

### GetEActivesessionEzsignaccessOk

`func (o *ActivesessionResponse) GetEActivesessionEzsignaccessOk() (*FieldEActivesessionEzsignaccess, bool)`

GetEActivesessionEzsignaccessOk returns a tuple with the EActivesessionEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsignaccess

`func (o *ActivesessionResponse) SetEActivesessionEzsignaccess(v FieldEActivesessionEzsignaccess)`

SetEActivesessionEzsignaccess sets EActivesessionEzsignaccess field to given value.


### GetEActivesessionEzsignprepaid

`func (o *ActivesessionResponse) GetEActivesessionEzsignprepaid() FieldEActivesessionEzsignprepaid`

GetEActivesessionEzsignprepaid returns the EActivesessionEzsignprepaid field if non-nil, zero value otherwise.

### GetEActivesessionEzsignprepaidOk

`func (o *ActivesessionResponse) GetEActivesessionEzsignprepaidOk() (*FieldEActivesessionEzsignprepaid, bool)`

GetEActivesessionEzsignprepaidOk returns a tuple with the EActivesessionEzsignprepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionEzsignprepaid

`func (o *ActivesessionResponse) SetEActivesessionEzsignprepaid(v FieldEActivesessionEzsignprepaid)`

SetEActivesessionEzsignprepaid sets EActivesessionEzsignprepaid field to given value.

### HasEActivesessionEzsignprepaid

`func (o *ActivesessionResponse) HasEActivesessionEzsignprepaid() bool`

HasEActivesessionEzsignprepaid returns a boolean if a field has been set.

### GetEActivesessionRealestateinprogress

`func (o *ActivesessionResponse) GetEActivesessionRealestateinprogress() FieldEActivesessionRealestateinprogress`

GetEActivesessionRealestateinprogress returns the EActivesessionRealestateinprogress field if non-nil, zero value otherwise.

### GetEActivesessionRealestateinprogressOk

`func (o *ActivesessionResponse) GetEActivesessionRealestateinprogressOk() (*FieldEActivesessionRealestateinprogress, bool)`

GetEActivesessionRealestateinprogressOk returns a tuple with the EActivesessionRealestateinprogress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEActivesessionRealestateinprogress

`func (o *ActivesessionResponse) SetEActivesessionRealestateinprogress(v FieldEActivesessionRealestateinprogress)`

SetEActivesessionRealestateinprogress sets EActivesessionRealestateinprogress field to given value.

### HasEActivesessionRealestateinprogress

`func (o *ActivesessionResponse) HasEActivesessionRealestateinprogress() bool`

HasEActivesessionRealestateinprogress returns a boolean if a field has been set.

### GetPksCustomerCode

`func (o *ActivesessionResponse) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *ActivesessionResponse) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *ActivesessionResponse) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *ActivesessionResponse) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *ActivesessionResponse) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *ActivesessionResponse) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetFkiSignatureID

`func (o *ActivesessionResponse) GetFkiSignatureID() int32`

GetFkiSignatureID returns the FkiSignatureID field if non-nil, zero value otherwise.

### GetFkiSignatureIDOk

`func (o *ActivesessionResponse) GetFkiSignatureIDOk() (*int32, bool)`

GetFkiSignatureIDOk returns a tuple with the FkiSignatureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSignatureID

`func (o *ActivesessionResponse) SetFkiSignatureID(v int32)`

SetFkiSignatureID sets FkiSignatureID field to given value.

### HasFkiSignatureID

`func (o *ActivesessionResponse) HasFkiSignatureID() bool`

HasFkiSignatureID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


