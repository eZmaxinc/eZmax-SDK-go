# UserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiAssistantID** | Pointer to **int32** | The unique ID of the Assistant. | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiCompanyIDDefault** | **int32** | The unique ID of the Company | 
**FkiDepartmentIDDefault** | **int32** | The unique ID of the Department | 
**FkiTimezoneID** | **int32** | The unique ID of the Timezone | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**ObjEmail** | [**EmailRequestCompound**](EmailRequestCompound.md) |  | 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**ObjPhoneHome** | Pointer to [**PhoneRequestCompound**](PhoneRequestCompound.md) |  | [optional] 
**ObjPhoneSMS** | Pointer to [**PhoneRequestCompound**](PhoneRequestCompound.md) |  | [optional] 
**FkiSecretquestionID** | Pointer to **int32** | The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father&#39;s middle name| |15|Your mother&#39;s maiden name| |16|Name of your eldest child| |17|Your spouse&#39;s middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat&#39;s name| |22|Date of Birth (YYYY-MM-DD)| |22|Secret Code| |22|Your reference code| | [optional] 
**SUserSecretresponse** | Pointer to **string** | The answer to the Secretquestion | [optional] 
**FkiModuleIDForm** | Pointer to **int32** | The unique ID of the Module | [optional] 
**EUserType** | [**FieldEUserType**](FieldEUserType.md) |  | 
**EUserLogintype** | [**FieldEUserLogintype**](FieldEUserLogintype.md) |  | 
**SUserFirstname** | **string** | The first name of the user | 
**SUserLastname** | **string** | The last name of the user | 
**SUserLoginname** | **string** | The login name of the User. | 
**EUserEzsignaccess** | [**FieldEUserEzsignaccess**](FieldEUserEzsignaccess.md) |  | 
**BUserIsactive** | **bool** | Whether the User is active or not | 
**BUserValidatebyadministration** | Pointer to **bool** | Whether if the transactions in which the User is implicated must be validated by administrative personnel or not | [optional] 
**BUserValidatebydirector** | Pointer to **bool** | Whether if the transactions in which the User is implicated must be validated by a director or not | [optional] 
**BUserAttachmentautoverified** | Pointer to **bool** | Whether if Attachments uploaded by the User must be validated or not | [optional] 
**BUserChangepassword** | Pointer to **bool** | Whether if the User is forced to change its password | [optional] 

## Methods

### NewUserRequest

`func NewUserRequest(fkiCompanyIDDefault int32, fkiDepartmentIDDefault int32, fkiTimezoneID int32, fkiLanguageID int32, objEmail EmailRequestCompound, fkiBillingentityinternalID int32, eUserType FieldEUserType, eUserLogintype FieldEUserLogintype, sUserFirstname string, sUserLastname string, sUserLoginname string, eUserEzsignaccess FieldEUserEzsignaccess, bUserIsactive bool, ) *UserRequest`

NewUserRequest instantiates a new UserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestWithDefaults

`func NewUserRequestWithDefaults() *UserRequest`

NewUserRequestWithDefaults instantiates a new UserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiUserID

`func (o *UserRequest) GetPkiUserID() int32`

GetPkiUserID returns the PkiUserID field if non-nil, zero value otherwise.

### GetPkiUserIDOk

`func (o *UserRequest) GetPkiUserIDOk() (*int32, bool)`

GetPkiUserIDOk returns a tuple with the PkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiUserID

`func (o *UserRequest) SetPkiUserID(v int32)`

SetPkiUserID sets PkiUserID field to given value.

### HasPkiUserID

`func (o *UserRequest) HasPkiUserID() bool`

HasPkiUserID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *UserRequest) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *UserRequest) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *UserRequest) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *UserRequest) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *UserRequest) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *UserRequest) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *UserRequest) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *UserRequest) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiAssistantID

`func (o *UserRequest) GetFkiAssistantID() int32`

GetFkiAssistantID returns the FkiAssistantID field if non-nil, zero value otherwise.

### GetFkiAssistantIDOk

`func (o *UserRequest) GetFkiAssistantIDOk() (*int32, bool)`

GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAssistantID

`func (o *UserRequest) SetFkiAssistantID(v int32)`

SetFkiAssistantID sets FkiAssistantID field to given value.

### HasFkiAssistantID

`func (o *UserRequest) HasFkiAssistantID() bool`

HasFkiAssistantID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *UserRequest) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *UserRequest) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *UserRequest) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *UserRequest) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiCompanyIDDefault

`func (o *UserRequest) GetFkiCompanyIDDefault() int32`

GetFkiCompanyIDDefault returns the FkiCompanyIDDefault field if non-nil, zero value otherwise.

### GetFkiCompanyIDDefaultOk

`func (o *UserRequest) GetFkiCompanyIDDefaultOk() (*int32, bool)`

GetFkiCompanyIDDefaultOk returns a tuple with the FkiCompanyIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCompanyIDDefault

`func (o *UserRequest) SetFkiCompanyIDDefault(v int32)`

SetFkiCompanyIDDefault sets FkiCompanyIDDefault field to given value.


### GetFkiDepartmentIDDefault

`func (o *UserRequest) GetFkiDepartmentIDDefault() int32`

GetFkiDepartmentIDDefault returns the FkiDepartmentIDDefault field if non-nil, zero value otherwise.

### GetFkiDepartmentIDDefaultOk

`func (o *UserRequest) GetFkiDepartmentIDDefaultOk() (*int32, bool)`

GetFkiDepartmentIDDefaultOk returns a tuple with the FkiDepartmentIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepartmentIDDefault

`func (o *UserRequest) SetFkiDepartmentIDDefault(v int32)`

SetFkiDepartmentIDDefault sets FkiDepartmentIDDefault field to given value.


### GetFkiTimezoneID

`func (o *UserRequest) GetFkiTimezoneID() int32`

GetFkiTimezoneID returns the FkiTimezoneID field if non-nil, zero value otherwise.

### GetFkiTimezoneIDOk

`func (o *UserRequest) GetFkiTimezoneIDOk() (*int32, bool)`

GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTimezoneID

`func (o *UserRequest) SetFkiTimezoneID(v int32)`

SetFkiTimezoneID sets FkiTimezoneID field to given value.


### GetFkiLanguageID

`func (o *UserRequest) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *UserRequest) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *UserRequest) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetObjEmail

`func (o *UserRequest) GetObjEmail() EmailRequestCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *UserRequest) GetObjEmailOk() (*EmailRequestCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *UserRequest) SetObjEmail(v EmailRequestCompound)`

SetObjEmail sets ObjEmail field to given value.


### GetFkiBillingentityinternalID

`func (o *UserRequest) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *UserRequest) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *UserRequest) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetObjPhoneHome

`func (o *UserRequest) GetObjPhoneHome() PhoneRequestCompound`

GetObjPhoneHome returns the ObjPhoneHome field if non-nil, zero value otherwise.

### GetObjPhoneHomeOk

`func (o *UserRequest) GetObjPhoneHomeOk() (*PhoneRequestCompound, bool)`

GetObjPhoneHomeOk returns a tuple with the ObjPhoneHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneHome

`func (o *UserRequest) SetObjPhoneHome(v PhoneRequestCompound)`

SetObjPhoneHome sets ObjPhoneHome field to given value.

### HasObjPhoneHome

`func (o *UserRequest) HasObjPhoneHome() bool`

HasObjPhoneHome returns a boolean if a field has been set.

### GetObjPhoneSMS

`func (o *UserRequest) GetObjPhoneSMS() PhoneRequestCompound`

GetObjPhoneSMS returns the ObjPhoneSMS field if non-nil, zero value otherwise.

### GetObjPhoneSMSOk

`func (o *UserRequest) GetObjPhoneSMSOk() (*PhoneRequestCompound, bool)`

GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneSMS

`func (o *UserRequest) SetObjPhoneSMS(v PhoneRequestCompound)`

SetObjPhoneSMS sets ObjPhoneSMS field to given value.

### HasObjPhoneSMS

`func (o *UserRequest) HasObjPhoneSMS() bool`

HasObjPhoneSMS returns a boolean if a field has been set.

### GetFkiSecretquestionID

`func (o *UserRequest) GetFkiSecretquestionID() int32`

GetFkiSecretquestionID returns the FkiSecretquestionID field if non-nil, zero value otherwise.

### GetFkiSecretquestionIDOk

`func (o *UserRequest) GetFkiSecretquestionIDOk() (*int32, bool)`

GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSecretquestionID

`func (o *UserRequest) SetFkiSecretquestionID(v int32)`

SetFkiSecretquestionID sets FkiSecretquestionID field to given value.

### HasFkiSecretquestionID

`func (o *UserRequest) HasFkiSecretquestionID() bool`

HasFkiSecretquestionID returns a boolean if a field has been set.

### GetSUserSecretresponse

`func (o *UserRequest) GetSUserSecretresponse() string`

GetSUserSecretresponse returns the SUserSecretresponse field if non-nil, zero value otherwise.

### GetSUserSecretresponseOk

`func (o *UserRequest) GetSUserSecretresponseOk() (*string, bool)`

GetSUserSecretresponseOk returns a tuple with the SUserSecretresponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserSecretresponse

`func (o *UserRequest) SetSUserSecretresponse(v string)`

SetSUserSecretresponse sets SUserSecretresponse field to given value.

### HasSUserSecretresponse

`func (o *UserRequest) HasSUserSecretresponse() bool`

HasSUserSecretresponse returns a boolean if a field has been set.

### GetFkiModuleIDForm

`func (o *UserRequest) GetFkiModuleIDForm() int32`

GetFkiModuleIDForm returns the FkiModuleIDForm field if non-nil, zero value otherwise.

### GetFkiModuleIDFormOk

`func (o *UserRequest) GetFkiModuleIDFormOk() (*int32, bool)`

GetFkiModuleIDFormOk returns a tuple with the FkiModuleIDForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiModuleIDForm

`func (o *UserRequest) SetFkiModuleIDForm(v int32)`

SetFkiModuleIDForm sets FkiModuleIDForm field to given value.

### HasFkiModuleIDForm

`func (o *UserRequest) HasFkiModuleIDForm() bool`

HasFkiModuleIDForm returns a boolean if a field has been set.

### GetEUserType

`func (o *UserRequest) GetEUserType() FieldEUserType`

GetEUserType returns the EUserType field if non-nil, zero value otherwise.

### GetEUserTypeOk

`func (o *UserRequest) GetEUserTypeOk() (*FieldEUserType, bool)`

GetEUserTypeOk returns a tuple with the EUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserType

`func (o *UserRequest) SetEUserType(v FieldEUserType)`

SetEUserType sets EUserType field to given value.


### GetEUserLogintype

`func (o *UserRequest) GetEUserLogintype() FieldEUserLogintype`

GetEUserLogintype returns the EUserLogintype field if non-nil, zero value otherwise.

### GetEUserLogintypeOk

`func (o *UserRequest) GetEUserLogintypeOk() (*FieldEUserLogintype, bool)`

GetEUserLogintypeOk returns a tuple with the EUserLogintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserLogintype

`func (o *UserRequest) SetEUserLogintype(v FieldEUserLogintype)`

SetEUserLogintype sets EUserLogintype field to given value.


### GetSUserFirstname

`func (o *UserRequest) GetSUserFirstname() string`

GetSUserFirstname returns the SUserFirstname field if non-nil, zero value otherwise.

### GetSUserFirstnameOk

`func (o *UserRequest) GetSUserFirstnameOk() (*string, bool)`

GetSUserFirstnameOk returns a tuple with the SUserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstname

`func (o *UserRequest) SetSUserFirstname(v string)`

SetSUserFirstname sets SUserFirstname field to given value.


### GetSUserLastname

`func (o *UserRequest) GetSUserLastname() string`

GetSUserLastname returns the SUserLastname field if non-nil, zero value otherwise.

### GetSUserLastnameOk

`func (o *UserRequest) GetSUserLastnameOk() (*string, bool)`

GetSUserLastnameOk returns a tuple with the SUserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastname

`func (o *UserRequest) SetSUserLastname(v string)`

SetSUserLastname sets SUserLastname field to given value.


### GetSUserLoginname

`func (o *UserRequest) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *UserRequest) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *UserRequest) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.


### GetEUserEzsignaccess

`func (o *UserRequest) GetEUserEzsignaccess() FieldEUserEzsignaccess`

GetEUserEzsignaccess returns the EUserEzsignaccess field if non-nil, zero value otherwise.

### GetEUserEzsignaccessOk

`func (o *UserRequest) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool)`

GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserEzsignaccess

`func (o *UserRequest) SetEUserEzsignaccess(v FieldEUserEzsignaccess)`

SetEUserEzsignaccess sets EUserEzsignaccess field to given value.


### GetBUserIsactive

`func (o *UserRequest) GetBUserIsactive() bool`

GetBUserIsactive returns the BUserIsactive field if non-nil, zero value otherwise.

### GetBUserIsactiveOk

`func (o *UserRequest) GetBUserIsactiveOk() (*bool, bool)`

GetBUserIsactiveOk returns a tuple with the BUserIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserIsactive

`func (o *UserRequest) SetBUserIsactive(v bool)`

SetBUserIsactive sets BUserIsactive field to given value.


### GetBUserValidatebyadministration

`func (o *UserRequest) GetBUserValidatebyadministration() bool`

GetBUserValidatebyadministration returns the BUserValidatebyadministration field if non-nil, zero value otherwise.

### GetBUserValidatebyadministrationOk

`func (o *UserRequest) GetBUserValidatebyadministrationOk() (*bool, bool)`

GetBUserValidatebyadministrationOk returns a tuple with the BUserValidatebyadministration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserValidatebyadministration

`func (o *UserRequest) SetBUserValidatebyadministration(v bool)`

SetBUserValidatebyadministration sets BUserValidatebyadministration field to given value.

### HasBUserValidatebyadministration

`func (o *UserRequest) HasBUserValidatebyadministration() bool`

HasBUserValidatebyadministration returns a boolean if a field has been set.

### GetBUserValidatebydirector

`func (o *UserRequest) GetBUserValidatebydirector() bool`

GetBUserValidatebydirector returns the BUserValidatebydirector field if non-nil, zero value otherwise.

### GetBUserValidatebydirectorOk

`func (o *UserRequest) GetBUserValidatebydirectorOk() (*bool, bool)`

GetBUserValidatebydirectorOk returns a tuple with the BUserValidatebydirector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserValidatebydirector

`func (o *UserRequest) SetBUserValidatebydirector(v bool)`

SetBUserValidatebydirector sets BUserValidatebydirector field to given value.

### HasBUserValidatebydirector

`func (o *UserRequest) HasBUserValidatebydirector() bool`

HasBUserValidatebydirector returns a boolean if a field has been set.

### GetBUserAttachmentautoverified

`func (o *UserRequest) GetBUserAttachmentautoverified() bool`

GetBUserAttachmentautoverified returns the BUserAttachmentautoverified field if non-nil, zero value otherwise.

### GetBUserAttachmentautoverifiedOk

`func (o *UserRequest) GetBUserAttachmentautoverifiedOk() (*bool, bool)`

GetBUserAttachmentautoverifiedOk returns a tuple with the BUserAttachmentautoverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserAttachmentautoverified

`func (o *UserRequest) SetBUserAttachmentautoverified(v bool)`

SetBUserAttachmentautoverified sets BUserAttachmentautoverified field to given value.

### HasBUserAttachmentautoverified

`func (o *UserRequest) HasBUserAttachmentautoverified() bool`

HasBUserAttachmentautoverified returns a boolean if a field has been set.

### GetBUserChangepassword

`func (o *UserRequest) GetBUserChangepassword() bool`

GetBUserChangepassword returns the BUserChangepassword field if non-nil, zero value otherwise.

### GetBUserChangepasswordOk

`func (o *UserRequest) GetBUserChangepasswordOk() (*bool, bool)`

GetBUserChangepasswordOk returns a tuple with the BUserChangepassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBUserChangepassword

`func (o *UserRequest) SetBUserChangepassword(v bool)`

SetBUserChangepassword sets BUserChangepassword field to given value.

### HasBUserChangepassword

`func (o *UserRequest) HasBUserChangepassword() bool`

HasBUserChangepassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


