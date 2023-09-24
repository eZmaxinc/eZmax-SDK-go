# UserResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserID** | **int32** | The unique ID of the User | 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiAssistantID** | Pointer to **int32** | The unique ID of the Assistant. | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiCompanyIDDefault** | **int32** | The unique ID of the Company | 
**SCompanyNameX** | **string** | The Name of the Company in the language of the requester | 
**FkiDepartmentIDDefault** | **int32** | The unique ID of the Department | 
**SDepartmentNameX** | **string** | The Name of the Department in the language of the requester | 
**FkiTimezoneID** | **int32** | The unique ID of the Timezone | 
**STimezoneName** | **string** | The description of the Timezone | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**SLanguageNameX** | **string** | The Name of the Language in the language of the requester | 
**ObjEmail** | [**EmailResponseCompound**](EmailResponseCompound.md) |  | 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**ObjPhoneHome** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**ObjPhoneSMS** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| | [optional] 
**FkiModuleIDForm** | Pointer to **int32** | The unique ID of the Module | [optional] 
**SModuleNameX** | Pointer to **string** | The Name of the Module in the language of the requester | [optional] 
**EUserOrigin** | [**FieldEUserOrigin**](FieldEUserOrigin.md) |  | 
**EUserType** | [**FieldEUserType**](FieldEUserType.md) |  | 
**EUserLogintype** | [**FieldEUserLogintype**](FieldEUserLogintype.md) |  | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SUserLoginname** | **string** | The login name of the User. | 
**EUserEzsignaccess** | [**FieldEUserEzsignaccess**](FieldEUserEzsignaccess.md) |  | 
**DtUserLastlogondate** | Pointer to **string** | The last logon date of the User | [optional] 
**DtUserPasswordchanged** | Pointer to **string** | The date at which the User&#39;s password was last changed | [optional] 
**DtUserEzsignprepaidexpiration** | Pointer to **string** | The eZsign prepaid expiration date | [optional] 
**BUserIsactive** | **bool** | Whether the User is active or not | 
**BUserValidatebyadministration** | Pointer to **bool** | Whether if the transactions in which the User is implicated must be validated by administrative personnel or not | [optional] 
**BUserValidatebydirector** | Pointer to **bool** | Whether if the transactions in which the User is implicated must be validated by a director or not | [optional] 
**BUserAttachmentautoverified** | Pointer to **bool** | Whether if Attachments uploaded by the User must be validated or not | [optional] 
**BUserChangepassword** | **bool** | Whether if the User is forced to change its password | 
**ObjAudit** | [**CommonAudit**](CommonAudit.md) |  | 

## Methods

### NewUserResponseCompound

`func NewUserResponseCompound(pkiUserID int32, fkiCompanyIDDefault int32, sCompanyNameX string, fkiDepartmentIDDefault int32, sDepartmentNameX string, fkiTimezoneID int32, sTimezoneName string, fkiLanguageID int32, sLanguageNameX string, objEmail EmailResponseCompound, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, eUserOrigin FieldEUserOrigin, eUserType FieldEUserType, eUserLogintype FieldEUserLogintype, sUserFirstname string, sUserLastname string, sUserLoginname string, eUserEzsignaccess FieldEUserEzsignaccess, bUserIsactive bool, bUserChangepassword bool, objAudit CommonAudit, ) *UserResponseCompound`

NewUserResponseCompound instantiates a new UserResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseCompoundWithDefaults

`func NewUserResponseCompoundWithDefaults() *UserResponseCompound`

NewUserResponseCompoundWithDefaults instantiates a new UserResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserID

`func (o *UserResponseCompound) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *UserResponseCompound) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *UserResponseCompound) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.


### GetFkiAgentID

`func (o *UserResponseCompound) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *UserResponseCompound) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *UserResponseCompound) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *UserResponseCompound) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *UserResponseCompound) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *UserResponseCompound) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *UserResponseCompound) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *UserResponseCompound) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiAssistantID

`func (o *UserResponseCompound) GetFkiAssistantID() int32`

GetFkiAssistantID returns the FkiAssistantID field if non-nil, zero value otherwise.

### GetFkiAssistantIDOk

`func (o *UserResponseCompound) GetFkiAssistantIDOk() (*int32, bool)`

GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAssistantID

`func (o *UserResponseCompound) SetFkiAssistantID(v int32)`

SetFkiAssistantID sets FkiAssistantID field to given value.

### HasFkiAssistantID

`func (o *UserResponseCompound) HasFkiAssistantID() bool`

HasFkiAssistantID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *UserResponseCompound) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *UserResponseCompound) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *UserResponseCompound) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *UserResponseCompound) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiCompanyIDDefault

`func (o *UserResponseCompound) GetFkiCompanyIDDefault() int32`

GetFkiCompanyIDDefault returns the FkiCompanyIDDefault field if non-nil, zero value otherwise.

### GetFkiCompanyIDDefaultOk

`func (o *UserResponseCompound) GetFkiCompanyIDDefaultOk() (*int32, bool)`

GetFkiCompanyIDDefaultOk returns a tuple with the FkiCompanyIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyIDDefault

`func (o *UserResponseCompound) SetFkiCompanyIDDefault(v int32)`

SetFkiCompanyIDDefault sets FkiCompanyIDDefault field to given value.


### GetSCompanyNameX

`func (o *UserResponseCompound) GetSCompanyNameX() string`

GetSCompanyNameX returns the SCompanyNameX field if non-nil, zero value otherwise.

### GetSCompanyNameXOk

`func (o *UserResponseCompound) GetSCompanyNameXOk() (*string, bool)`

GetSCompanyNameXOk returns a tuple with the SCompanyNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyNameX

`func (o *UserResponseCompound) SetSCompanyNameX(v string)`

SetSCompanyNameX sets SCompanyNameX field to given value.


### GetFkiDepartmentIDDefault

`func (o *UserResponseCompound) GetFkiDepartmentIDDefault() int32`

GetFkiDepartmentIDDefault returns the FkiDepartmentIDDefault field if non-nil, zero value otherwise.

### GetFkiDepartmentIDDefaultOk

`func (o *UserResponseCompound) GetFkiDepartmentIDDefaultOk() (*int32, bool)`

GetFkiDepartmentIDDefaultOk returns a tuple with the FkiDepartmentIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentIDDefault

`func (o *UserResponseCompound) SetFkiDepartmentIDDefault(v int32)`

SetFkiDepartmentIDDefault sets FkiDepartmentIDDefault field to given value.


### GetSDepartmentNameX

`func (o *UserResponseCompound) GetSDepartmentNameX() string`

GetSDepartmentNameX returns the SDepartmentNameX field if non-nil, zero value otherwise.

### GetSDepartmentNameXOk

`func (o *UserResponseCompound) GetSDepartmentNameXOk() (*string, bool)`

GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDepartmentNameX

`func (o *UserResponseCompound) SetSDepartmentNameX(v string)`

SetSDepartmentNameX sets SDepartmentNameX field to given value.


### GetFkiTimezoneID

`func (o *UserResponseCompound) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *UserResponseCompound) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *UserResponseCompound) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.


### GetSTimezoneName

`func (o *UserResponseCompound) GetSTimezoneName() string`

GetSTimezoneName returns the STimezoneName field if non-nil, zero value otherwise.

### GetSTimezoneNameOk

`func (o *UserResponseCompound) GetSTimezoneNameOk() (*string, bool)`

GetSTimezoneNameOk returns a tuple with the STimezoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTimezoneName

`func (o *UserResponseCompound) SetSTimezoneName(v string)`

SetSTimezoneName sets STimezoneName field to given value.


### GetFkiLanguageID

`func (o *UserResponseCompound) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *UserResponseCompound) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *UserResponseCompound) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetSLanguageNameX

`func (o *UserResponseCompound) GetSLanguageNameX() string`

GetSLanguageNameX returns the SLanguageNameX field if non-nil, zero value otherwise.

### GetSLanguageNameXOk

`func (o *UserResponseCompound) GetSLanguageNameXOk() (*string, bool)`

GetSLanguageNameXOk returns a tuple with the SLanguageNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSLanguageNameX

`func (o *UserResponseCompound) SetSLanguageNameX(v string)`

SetSLanguageNameX sets SLanguageNameX field to given value.


### GetObjEmail

`func (o *UserResponseCompound) GetObjEmail() EmailResponseCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *UserResponseCompound) GetObjEmailOk() (*EmailResponseCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *UserResponseCompound) SetObjEmail(v EmailResponseCompound)`

SetObjEmail sets ObjEmail field to given value.


### GetFkiBillingentityinternalID

`func (o *UserResponseCompound) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *UserResponseCompound) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *UserResponseCompound) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *UserResponseCompound) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *UserResponseCompound) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *UserResponseCompound) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.


### GetObjPhoneHome

`func (o *UserResponseCompound) GetObjPhoneHome() PhoneResponseCompound`

GetObjPhoneHome returns the ObjPhoneHome field if non-nil, zero value otherwise.

### GetObjPhoneHomeOk

`func (o *UserResponseCompound) GetObjPhoneHomeOk() (*PhoneResponseCompound, bool)`

GetObjPhoneHomeOk returns a tuple with the ObjPhoneHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneHome

`func (o *UserResponseCompound) SetObjPhoneHome(v PhoneResponseCompound)`

SetObjPhoneHome sets ObjPhoneHome field to given value.

### HasObjPhoneHome

`func (o *UserResponseCompound) HasObjPhoneHome() bool`

HasObjPhoneHome returns a boolean if a field has been set.

### GetObjPhoneSMS

`func (o *UserResponseCompound) GetObjPhoneSMS() PhoneResponseCompound`

GetObjPhoneSMS returns the ObjPhoneSMS field if non-nil, zero value otherwise.

### GetObjPhoneSMSOk

`func (o *UserResponseCompound) GetObjPhoneSMSOk() (*PhoneResponseCompound, bool)`

GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneSMS

`func (o *UserResponseCompound) SetObjPhoneSMS(v PhoneResponseCompound)`

SetObjPhoneSMS sets ObjPhoneSMS field to given value.

### HasObjPhoneSMS

`func (o *UserResponseCompound) HasObjPhoneSMS() bool`

HasObjPhoneSMS returns a boolean if a field has been set.

### GetFkiSecretquestionID

`func (o *UserResponseCompound) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *UserResponseCompound) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *UserResponseCompound) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *UserResponseCompound) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetFkiModuleIDForm

`func (o *UserResponseCompound) GetFkiModuleIDForm() int32`

GetFkiModuleIDForm returns the FkiModuleIDForm field if non-nil, zero value otherwise.

### GetFkiModuleIDFormOk

`func (o *UserResponseCompound) GetFkiModuleIDFormOk() (*int32, bool)`

GetFkiModuleIDFormOk returns a tuple with the FkiModuleIDForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleIDForm

`func (o *UserResponseCompound) SetFkiModuleIDForm(v int32)`

SetFkiModuleIDForm sets FkiModuleIDForm field to given value.

### HasFkiModuleIDForm

`func (o *UserResponseCompound) HasFkiModuleIDForm() bool`

HasFkiModuleIDForm returns a boolean if a field has been set.

### GetSModuleNameX

`func (o *UserResponseCompound) GetSModuleNameX() string`

GetSModuleNameX returns the SModuleNameX field if non-nil, zero value otherwise.

### GetSModuleNameXOk

`func (o *UserResponseCompound) GetSModuleNameXOk() (*string, bool)`

GetSModuleNameXOk returns a tuple with the SModuleNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSModuleNameX

`func (o *UserResponseCompound) SetSModuleNameX(v string)`

SetSModuleNameX sets SModuleNameX field to given value.

### HasSModuleNameX

`func (o *UserResponseCompound) HasSModuleNameX() bool`

HasSModuleNameX returns a boolean if a field has been set.

### GetEUserOrigin

`func (o *UserResponseCompound) GetEUserOrigin() FieldEUserOrigin`

GetEUserOrigin returns the EUserOrigin field if non-nil, zero value otherwise.

### GetEUserOriginOk

`func (o *UserResponseCompound) GetEUserOriginOk() (*FieldEUserOrigin, bool)`

GetEUserOriginOk returns a tuple with the EUserOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserOrigin

`func (o *UserResponseCompound) SetEUserOrigin(v FieldEUserOrigin)`

SetEUserOrigin sets EUserOrigin field to given value.


### GetEUserType

`func (o *UserResponseCompound) GetEUserType() FieldEUserType`

GetEUserType returns the EUserType field if non-nil, zero value otherwise.

### GetEUserTypeOk

`func (o *UserResponseCompound) GetEUserTypeOk() (*FieldEUserType, bool)`

GetEUserTypeOk returns a tuple with the EUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserType

`func (o *UserResponseCompound) SetEUserType(v FieldEUserType)`

SetEUserType sets EUserType field to given value.


### GetEUserLogintype

`func (o *UserResponseCompound) GetEUserLogintype() FieldEUserLogintype`

GetEUserLogintype returns the EUserLogintype field if non-nil, zero value otherwise.

### GetEUserLogintypeOk

`func (o *UserResponseCompound) GetEUserLogintypeOk() (*FieldEUserLogintype, bool)`

GetEUserLogintypeOk returns a tuple with the EUserLogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserLogintype

`func (o *UserResponseCompound) SetEUserLogintype(v FieldEUserLogintype)`

SetEUserLogintype sets EUserLogintype field to given value.


### GetSUserFirstname

`func (o *UserResponseCompound) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UserResponseCompound) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UserResponseCompound) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UserResponseCompound) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UserResponseCompound) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UserResponseCompound) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UserResponseCompound) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UserResponseCompound) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UserResponseCompound) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetEUserEzsignaccess

`func (o *UserResponseCompound) GetEUserEzsignaccess() FieldEUserEzsignaccess`

GetEUserEzsignaccess returns the EUserEzsignaccess field if non-nil, zero value otherwise.

### GetEUserEzsignaccessOk

`func (o *UserResponseCompound) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool)`

GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignaccess

`func (o *UserResponseCompound) SetEUserEzsignaccess(v FieldEUserEzsignaccess)`

SetEUserEzsignaccess sets EUserEzsignaccess field to given value.


### GetDtUserLastlogondate

`func (o *UserResponseCompound) GetDtUserLastlogondate() string`

GetDtUserLastlogondate returns the DtUserLastlogondate field if non-nil, zero value otherwise.

### GetDtUserLastlogondateOk

`func (o *UserResponseCompound) GetDtUserLastlogondateOk() (*string, bool)`

GetDtUserLastlogondateOk returns a tuple with the DtUserLastlogondate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtUserLastlogondate

`func (o *UserResponseCompound) SetDtUserLastlogondate(v string)`

SetDtUserLastlogondate sets DtUserLastlogondate field to given value.

### HasDtUserLastlogondate

`func (o *UserResponseCompound) HasDtUserLastlogondate() bool`

HasDtUserLastlogondate returns a boolean if a field has been set.

### GetDtUserPasswordchanged

`func (o *UserResponseCompound) GetDtUserPasswordchanged() string`

GetDtUserPasswordchanged returns the DtUserPasswordchanged field if non-nil, zero value otherwise.

### GetDtUserPasswordchangedOk

`func (o *UserResponseCompound) GetDtUserPasswordchangedOk() (*string, bool)`

GetDtUserPasswordchangedOk returns a tuple with the DtUserPasswordchanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtUserPasswordchanged

`func (o *UserResponseCompound) SetDtUserPasswordchanged(v string)`

SetDtUserPasswordchanged sets DtUserPasswordchanged field to given value.

### HasDtUserPasswordchanged

`func (o *UserResponseCompound) HasDtUserPasswordchanged() bool`

HasDtUserPasswordchanged returns a boolean if a field has been set.

### GetDtUserEzsignprepaidexpiration

`func (o *UserResponseCompound) GetDtUserEzsignprepaidexpiration() string`

GetDtUserEzsignprepaidexpiration returns the DtUserEzsignprepaidexpiration field if non-nil, zero value otherwise.

### GetDtUserEzsignprepaidexpirationOk

`func (o *UserResponseCompound) GetDtUserEzsignprepaidexpirationOk() (*string, bool)`

GetDtUserEzsignprepaidexpirationOk returns a tuple with the DtUserEzsignprepaidexpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtUserEzsignprepaidexpiration

`func (o *UserResponseCompound) SetDtUserEzsignprepaidexpiration(v string)`

SetDtUserEzsignprepaidexpiration sets DtUserEzsignprepaidexpiration field to given value.

### HasDtUserEzsignprepaidexpiration

`func (o *UserResponseCompound) HasDtUserEzsignprepaidexpiration() bool`

HasDtUserEzsignprepaidexpiration returns a boolean if a field has been set.

### GetBUserIsactive

`func (o *UserResponseCompound) GetBUserIsactive() bool`

GetBUserIsactive returns the BUserIsactive field if non-nil, zero value otherwise.

### GetBUserIsactiveOk

`func (o *UserResponseCompound) GetBUserIsactiveOk() (*bool, bool)`

GetBUserIsactiveOk returns a tuple with the BUserIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserIsactive

`func (o *UserResponseCompound) SetBUserIsactive(v bool)`

SetBUserIsactive sets BUserIsactive field to given value.


### GetBUserValidatebyadministration

`func (o *UserResponseCompound) GetBUserValidatebyadministration() bool`

GetBUserValidatebyadministration returns the BUserValidatebyadministration field if non-nil, zero value otherwise.

### GetBUserValidatebyadministrationOk

`func (o *UserResponseCompound) GetBUserValidatebyadministrationOk() (*bool, bool)`

GetBUserValidatebyadministrationOk returns a tuple with the BUserValidatebyadministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserValidatebyadministration

`func (o *UserResponseCompound) SetBUserValidatebyadministration(v bool)`

SetBUserValidatebyadministration sets BUserValidatebyadministration field to given value.

### HasBUserValidatebyadministration

`func (o *UserResponseCompound) HasBUserValidatebyadministration() bool`

HasBUserValidatebyadministration returns a boolean if a field has been set.

### GetBUserValidatebydirector

`func (o *UserResponseCompound) GetBUserValidatebydirector() bool`

GetBUserValidatebydirector returns the BUserValidatebydirector field if non-nil, zero value otherwise.

### GetBUserValidatebydirectorOk

`func (o *UserResponseCompound) GetBUserValidatebydirectorOk() (*bool, bool)`

GetBUserValidatebydirectorOk returns a tuple with the BUserValidatebydirector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserValidatebydirector

`func (o *UserResponseCompound) SetBUserValidatebydirector(v bool)`

SetBUserValidatebydirector sets BUserValidatebydirector field to given value.

### HasBUserValidatebydirector

`func (o *UserResponseCompound) HasBUserValidatebydirector() bool`

HasBUserValidatebydirector returns a boolean if a field has been set.

### GetBUserAttachmentautoverified

`func (o *UserResponseCompound) GetBUserAttachmentautoverified() bool`

GetBUserAttachmentautoverified returns the BUserAttachmentautoverified field if non-nil, zero value otherwise.

### GetBUserAttachmentautoverifiedOk

`func (o *UserResponseCompound) GetBUserAttachmentautoverifiedOk() (*bool, bool)`

GetBUserAttachmentautoverifiedOk returns a tuple with the BUserAttachmentautoverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserAttachmentautoverified

`func (o *UserResponseCompound) SetBUserAttachmentautoverified(v bool)`

SetBUserAttachmentautoverified sets BUserAttachmentautoverified field to given value.

### HasBUserAttachmentautoverified

`func (o *UserResponseCompound) HasBUserAttachmentautoverified() bool`

HasBUserAttachmentautoverified returns a boolean if a field has been set.

### GetBUserChangepassword

`func (o *UserResponseCompound) GetBUserChangepassword() bool`

GetBUserChangepassword returns the BUserChangepassword field if non-nil, zero value otherwise.

### GetBUserChangepasswordOk

`func (o *UserResponseCompound) GetBUserChangepasswordOk() (*bool, bool)`

GetBUserChangepasswordOk returns a tuple with the BUserChangepassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserChangepassword

`func (o *UserResponseCompound) SetBUserChangepassword(v bool)`

SetBUserChangepassword sets BUserChangepassword field to given value.


### GetObjAudit

`func (o *UserResponseCompound) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *UserResponseCompound) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *UserResponseCompound) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


