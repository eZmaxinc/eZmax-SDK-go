/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the UserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponse{}

// UserResponse A User Object
type UserResponse struct {
	// The unique ID of the User
	PkiUserID int32 `json:"pkiUserID"`
	// The unique ID of the Agent.
	FkiAgentID *int32 `json:"fkiAgentID,omitempty"`
	// The unique ID of the Broker.
	FkiBrokerID *int32 `json:"fkiBrokerID,omitempty"`
	// The unique ID of the Assistant.
	FkiAssistantID *int32 `json:"fkiAssistantID,omitempty"`
	// The unique ID of the Employee.
	FkiEmployeeID *int32 `json:"fkiEmployeeID,omitempty"`
	// The unique ID of the Company
	FkiCompanyIDDefault int32 `json:"fkiCompanyIDDefault"`
	// The Name of the Company in the language of the requester
	SCompanyNameX string `json:"sCompanyNameX"`
	// The unique ID of the Department
	FkiDepartmentIDDefault int32 `json:"fkiDepartmentIDDefault"`
	// The Name of the Department in the language of the requester
	SDepartmentNameX string `json:"sDepartmentNameX"`
	// The unique ID of the Timezone
	FkiTimezoneID int32 `json:"fkiTimezoneID"`
	// The description of the Timezone
	STimezoneName string `json:"sTimezoneName"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The Name of the Language in the language of the requester
	SLanguageNameX string `json:"sLanguageNameX"`
	ObjEmail EmailResponseCompound `json:"objEmail"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX string `json:"sBillingentityinternalDescriptionX"`
	ObjPhoneHome *PhoneResponseCompound `json:"objPhoneHome,omitempty"`
	ObjPhoneSMS *PhoneResponseCompound `json:"objPhoneSMS,omitempty"`
	// The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father's middle name| |15|Your mother's maiden name| |16|Name of your eldest child| |17|Your spouse's middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat's name| |22|Date of Birth (YYYY-MM-DD)|
	FkiSecretquestionID *int32 `json:"fkiSecretquestionID,omitempty"`
	// The unique ID of the Module
	FkiModuleIDForm *int32 `json:"fkiModuleIDForm,omitempty"`
	// The Name of the Module in the language of the requester
	SModuleNameX *string `json:"sModuleNameX,omitempty"`
	EUserOrigin FieldEUserOrigin `json:"eUserOrigin"`
	EUserType FieldEUserType `json:"eUserType"`
	EUserLogintype FieldEUserLogintype `json:"eUserLogintype"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	EUserEzsignaccess FieldEUserEzsignaccess `json:"eUserEzsignaccess"`
	// The last logon date of the User
	DtUserLastlogondate *string `json:"dtUserLastlogondate,omitempty"`
	// The date at which the User's password was last changed
	DtUserPasswordchanged *string `json:"dtUserPasswordchanged,omitempty"`
	// The eZsign prepaid expiration date
	DtUserEzsignprepaidexpiration *string `json:"dtUserEzsignprepaidexpiration,omitempty"`
	// Whether the User is active or not
	BUserIsactive bool `json:"bUserIsactive"`
	// Whether if the transactions in which the User is implicated must be validated by administrative personnel or not
	BUserValidatebyadministration *bool `json:"bUserValidatebyadministration,omitempty"`
	// Whether if the transactions in which the User is implicated must be validated by a director or not
	BUserValidatebydirector *bool `json:"bUserValidatebydirector,omitempty"`
	// Whether if Attachments uploaded by the User must be validated or not
	BUserAttachmentautoverified *bool `json:"bUserAttachmentautoverified,omitempty"`
	// Whether if the User is forced to change its password
	BUserChangepassword bool `json:"bUserChangepassword"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse(pkiUserID int32, fkiCompanyIDDefault int32, sCompanyNameX string, fkiDepartmentIDDefault int32, sDepartmentNameX string, fkiTimezoneID int32, sTimezoneName string, fkiLanguageID int32, sLanguageNameX string, objEmail EmailResponseCompound, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, eUserOrigin FieldEUserOrigin, eUserType FieldEUserType, eUserLogintype FieldEUserLogintype, sUserFirstname string, sUserLastname string, sUserLoginname string, eUserEzsignaccess FieldEUserEzsignaccess, bUserIsactive bool, bUserChangepassword bool, objAudit CommonAudit) *UserResponse {
	this := UserResponse{}
	this.PkiUserID = pkiUserID
	this.FkiCompanyIDDefault = fkiCompanyIDDefault
	this.SCompanyNameX = sCompanyNameX
	this.FkiDepartmentIDDefault = fkiDepartmentIDDefault
	this.SDepartmentNameX = sDepartmentNameX
	this.FkiTimezoneID = fkiTimezoneID
	this.STimezoneName = sTimezoneName
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.ObjEmail = objEmail
	this.FkiBillingentityinternalID = fkiBillingentityinternalID
	this.SBillingentityinternalDescriptionX = sBillingentityinternalDescriptionX
	this.EUserOrigin = eUserOrigin
	this.EUserType = eUserType
	this.EUserLogintype = eUserLogintype
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.EUserEzsignaccess = eUserEzsignaccess
	this.BUserIsactive = bUserIsactive
	this.BUserChangepassword = bUserChangepassword
	this.ObjAudit = objAudit
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value
func (o *UserResponse) GetPkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUserID, true
}

// SetPkiUserID sets field value
func (o *UserResponse) SetPkiUserID(v int32) {
	o.PkiUserID = v
}

// GetFkiAgentID returns the FkiAgentID field value if set, zero value otherwise.
func (o *UserResponse) GetFkiAgentID() int32 {
	if o == nil || IsNil(o.FkiAgentID) {
		var ret int32
		return ret
	}
	return *o.FkiAgentID
}

// GetFkiAgentIDOk returns a tuple with the FkiAgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiAgentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAgentID) {
		return nil, false
	}
	return o.FkiAgentID, true
}

// HasFkiAgentID returns a boolean if a field has been set.
func (o *UserResponse) HasFkiAgentID() bool {
	if o != nil && !IsNil(o.FkiAgentID) {
		return true
	}

	return false
}

// SetFkiAgentID gets a reference to the given int32 and assigns it to the FkiAgentID field.
func (o *UserResponse) SetFkiAgentID(v int32) {
	o.FkiAgentID = &v
}

// GetFkiBrokerID returns the FkiBrokerID field value if set, zero value otherwise.
func (o *UserResponse) GetFkiBrokerID() int32 {
	if o == nil || IsNil(o.FkiBrokerID) {
		var ret int32
		return ret
	}
	return *o.FkiBrokerID
}

// GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiBrokerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrokerID) {
		return nil, false
	}
	return o.FkiBrokerID, true
}

// HasFkiBrokerID returns a boolean if a field has been set.
func (o *UserResponse) HasFkiBrokerID() bool {
	if o != nil && !IsNil(o.FkiBrokerID) {
		return true
	}

	return false
}

// SetFkiBrokerID gets a reference to the given int32 and assigns it to the FkiBrokerID field.
func (o *UserResponse) SetFkiBrokerID(v int32) {
	o.FkiBrokerID = &v
}

// GetFkiAssistantID returns the FkiAssistantID field value if set, zero value otherwise.
func (o *UserResponse) GetFkiAssistantID() int32 {
	if o == nil || IsNil(o.FkiAssistantID) {
		var ret int32
		return ret
	}
	return *o.FkiAssistantID
}

// GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiAssistantIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAssistantID) {
		return nil, false
	}
	return o.FkiAssistantID, true
}

// HasFkiAssistantID returns a boolean if a field has been set.
func (o *UserResponse) HasFkiAssistantID() bool {
	if o != nil && !IsNil(o.FkiAssistantID) {
		return true
	}

	return false
}

// SetFkiAssistantID gets a reference to the given int32 and assigns it to the FkiAssistantID field.
func (o *UserResponse) SetFkiAssistantID(v int32) {
	o.FkiAssistantID = &v
}

// GetFkiEmployeeID returns the FkiEmployeeID field value if set, zero value otherwise.
func (o *UserResponse) GetFkiEmployeeID() int32 {
	if o == nil || IsNil(o.FkiEmployeeID) {
		var ret int32
		return ret
	}
	return *o.FkiEmployeeID
}

// GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiEmployeeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEmployeeID) {
		return nil, false
	}
	return o.FkiEmployeeID, true
}

// HasFkiEmployeeID returns a boolean if a field has been set.
func (o *UserResponse) HasFkiEmployeeID() bool {
	if o != nil && !IsNil(o.FkiEmployeeID) {
		return true
	}

	return false
}

// SetFkiEmployeeID gets a reference to the given int32 and assigns it to the FkiEmployeeID field.
func (o *UserResponse) SetFkiEmployeeID(v int32) {
	o.FkiEmployeeID = &v
}

// GetFkiCompanyIDDefault returns the FkiCompanyIDDefault field value
func (o *UserResponse) GetFkiCompanyIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCompanyIDDefault
}

// GetFkiCompanyIDDefaultOk returns a tuple with the FkiCompanyIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiCompanyIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCompanyIDDefault, true
}

// SetFkiCompanyIDDefault sets field value
func (o *UserResponse) SetFkiCompanyIDDefault(v int32) {
	o.FkiCompanyIDDefault = v
}

// GetSCompanyNameX returns the SCompanyNameX field value
func (o *UserResponse) GetSCompanyNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCompanyNameX
}

// GetSCompanyNameXOk returns a tuple with the SCompanyNameX field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSCompanyNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCompanyNameX, true
}

// SetSCompanyNameX sets field value
func (o *UserResponse) SetSCompanyNameX(v string) {
	o.SCompanyNameX = v
}

// GetFkiDepartmentIDDefault returns the FkiDepartmentIDDefault field value
func (o *UserResponse) GetFkiDepartmentIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDepartmentIDDefault
}

// GetFkiDepartmentIDDefaultOk returns a tuple with the FkiDepartmentIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiDepartmentIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDepartmentIDDefault, true
}

// SetFkiDepartmentIDDefault sets field value
func (o *UserResponse) SetFkiDepartmentIDDefault(v int32) {
	o.FkiDepartmentIDDefault = v
}

// GetSDepartmentNameX returns the SDepartmentNameX field value
func (o *UserResponse) GetSDepartmentNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDepartmentNameX
}

// GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSDepartmentNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDepartmentNameX, true
}

// SetSDepartmentNameX sets field value
func (o *UserResponse) SetSDepartmentNameX(v string) {
	o.SDepartmentNameX = v
}

// GetFkiTimezoneID returns the FkiTimezoneID field value
func (o *UserResponse) GetFkiTimezoneID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiTimezoneID
}

// GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiTimezoneIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiTimezoneID, true
}

// SetFkiTimezoneID sets field value
func (o *UserResponse) SetFkiTimezoneID(v int32) {
	o.FkiTimezoneID = v
}

// GetSTimezoneName returns the STimezoneName field value
func (o *UserResponse) GetSTimezoneName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.STimezoneName
}

// GetSTimezoneNameOk returns a tuple with the STimezoneName field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSTimezoneNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.STimezoneName, true
}

// SetSTimezoneName sets field value
func (o *UserResponse) SetSTimezoneName(v string) {
	o.STimezoneName = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserResponse) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSLanguageNameX returns the SLanguageNameX field value
func (o *UserResponse) GetSLanguageNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSLanguageNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLanguageNameX, true
}

// SetSLanguageNameX sets field value
func (o *UserResponse) SetSLanguageNameX(v string) {
	o.SLanguageNameX = v
}

// GetObjEmail returns the ObjEmail field value
func (o *UserResponse) GetObjEmail() EmailResponseCompound {
	if o == nil {
		var ret EmailResponseCompound
		return ret
	}

	return o.ObjEmail
}

// GetObjEmailOk returns a tuple with the ObjEmail field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetObjEmailOk() (*EmailResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEmail, true
}

// SetObjEmail sets field value
func (o *UserResponse) SetObjEmail(v EmailResponseCompound) {
	o.ObjEmail = v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *UserResponse) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *UserResponse) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value
func (o *UserResponse) GetSBillingentityinternalDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityinternalDescriptionX, true
}

// SetSBillingentityinternalDescriptionX sets field value
func (o *UserResponse) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = v
}

// GetObjPhoneHome returns the ObjPhoneHome field value if set, zero value otherwise.
func (o *UserResponse) GetObjPhoneHome() PhoneResponseCompound {
	if o == nil || IsNil(o.ObjPhoneHome) {
		var ret PhoneResponseCompound
		return ret
	}
	return *o.ObjPhoneHome
}

// GetObjPhoneHomeOk returns a tuple with the ObjPhoneHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetObjPhoneHomeOk() (*PhoneResponseCompound, bool) {
	if o == nil || IsNil(o.ObjPhoneHome) {
		return nil, false
	}
	return o.ObjPhoneHome, true
}

// HasObjPhoneHome returns a boolean if a field has been set.
func (o *UserResponse) HasObjPhoneHome() bool {
	if o != nil && !IsNil(o.ObjPhoneHome) {
		return true
	}

	return false
}

// SetObjPhoneHome gets a reference to the given PhoneResponseCompound and assigns it to the ObjPhoneHome field.
func (o *UserResponse) SetObjPhoneHome(v PhoneResponseCompound) {
	o.ObjPhoneHome = &v
}

// GetObjPhoneSMS returns the ObjPhoneSMS field value if set, zero value otherwise.
func (o *UserResponse) GetObjPhoneSMS() PhoneResponseCompound {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		var ret PhoneResponseCompound
		return ret
	}
	return *o.ObjPhoneSMS
}

// GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetObjPhoneSMSOk() (*PhoneResponseCompound, bool) {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		return nil, false
	}
	return o.ObjPhoneSMS, true
}

// HasObjPhoneSMS returns a boolean if a field has been set.
func (o *UserResponse) HasObjPhoneSMS() bool {
	if o != nil && !IsNil(o.ObjPhoneSMS) {
		return true
	}

	return false
}

// SetObjPhoneSMS gets a reference to the given PhoneResponseCompound and assigns it to the ObjPhoneSMS field.
func (o *UserResponse) SetObjPhoneSMS(v PhoneResponseCompound) {
	o.ObjPhoneSMS = &v
}

// GetFkiSecretquestionID returns the FkiSecretquestionID field value if set, zero value otherwise.
func (o *UserResponse) GetFkiSecretquestionID() int32 {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		var ret int32
		return ret
	}
	return *o.FkiSecretquestionID
}

// GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiSecretquestionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		return nil, false
	}
	return o.FkiSecretquestionID, true
}

// HasFkiSecretquestionID returns a boolean if a field has been set.
func (o *UserResponse) HasFkiSecretquestionID() bool {
	if o != nil && !IsNil(o.FkiSecretquestionID) {
		return true
	}

	return false
}

// SetFkiSecretquestionID gets a reference to the given int32 and assigns it to the FkiSecretquestionID field.
func (o *UserResponse) SetFkiSecretquestionID(v int32) {
	o.FkiSecretquestionID = &v
}

// GetFkiModuleIDForm returns the FkiModuleIDForm field value if set, zero value otherwise.
func (o *UserResponse) GetFkiModuleIDForm() int32 {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		var ret int32
		return ret
	}
	return *o.FkiModuleIDForm
}

// GetFkiModuleIDFormOk returns a tuple with the FkiModuleIDForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiModuleIDFormOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		return nil, false
	}
	return o.FkiModuleIDForm, true
}

// HasFkiModuleIDForm returns a boolean if a field has been set.
func (o *UserResponse) HasFkiModuleIDForm() bool {
	if o != nil && !IsNil(o.FkiModuleIDForm) {
		return true
	}

	return false
}

// SetFkiModuleIDForm gets a reference to the given int32 and assigns it to the FkiModuleIDForm field.
func (o *UserResponse) SetFkiModuleIDForm(v int32) {
	o.FkiModuleIDForm = &v
}

// GetSModuleNameX returns the SModuleNameX field value if set, zero value otherwise.
func (o *UserResponse) GetSModuleNameX() string {
	if o == nil || IsNil(o.SModuleNameX) {
		var ret string
		return ret
	}
	return *o.SModuleNameX
}

// GetSModuleNameXOk returns a tuple with the SModuleNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSModuleNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SModuleNameX) {
		return nil, false
	}
	return o.SModuleNameX, true
}

// HasSModuleNameX returns a boolean if a field has been set.
func (o *UserResponse) HasSModuleNameX() bool {
	if o != nil && !IsNil(o.SModuleNameX) {
		return true
	}

	return false
}

// SetSModuleNameX gets a reference to the given string and assigns it to the SModuleNameX field.
func (o *UserResponse) SetSModuleNameX(v string) {
	o.SModuleNameX = &v
}

// GetEUserOrigin returns the EUserOrigin field value
func (o *UserResponse) GetEUserOrigin() FieldEUserOrigin {
	if o == nil {
		var ret FieldEUserOrigin
		return ret
	}

	return o.EUserOrigin
}

// GetEUserOriginOk returns a tuple with the EUserOrigin field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEUserOriginOk() (*FieldEUserOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserOrigin, true
}

// SetEUserOrigin sets field value
func (o *UserResponse) SetEUserOrigin(v FieldEUserOrigin) {
	o.EUserOrigin = v
}

// GetEUserType returns the EUserType field value
func (o *UserResponse) GetEUserType() FieldEUserType {
	if o == nil {
		var ret FieldEUserType
		return ret
	}

	return o.EUserType
}

// GetEUserTypeOk returns a tuple with the EUserType field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEUserTypeOk() (*FieldEUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserType, true
}

// SetEUserType sets field value
func (o *UserResponse) SetEUserType(v FieldEUserType) {
	o.EUserType = v
}

// GetEUserLogintype returns the EUserLogintype field value
func (o *UserResponse) GetEUserLogintype() FieldEUserLogintype {
	if o == nil {
		var ret FieldEUserLogintype
		return ret
	}

	return o.EUserLogintype
}

// GetEUserLogintypeOk returns a tuple with the EUserLogintype field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEUserLogintypeOk() (*FieldEUserLogintype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserLogintype, true
}

// SetEUserLogintype sets field value
func (o *UserResponse) SetEUserLogintype(v FieldEUserLogintype) {
	o.EUserLogintype = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserResponse) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserResponse) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserResponse) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UserResponse) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UserResponse) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetEUserEzsignaccess returns the EUserEzsignaccess field value
func (o *UserResponse) GetEUserEzsignaccess() FieldEUserEzsignaccess {
	if o == nil {
		var ret FieldEUserEzsignaccess
		return ret
	}

	return o.EUserEzsignaccess
}

// GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserEzsignaccess, true
}

// SetEUserEzsignaccess sets field value
func (o *UserResponse) SetEUserEzsignaccess(v FieldEUserEzsignaccess) {
	o.EUserEzsignaccess = v
}

// GetDtUserLastlogondate returns the DtUserLastlogondate field value if set, zero value otherwise.
func (o *UserResponse) GetDtUserLastlogondate() string {
	if o == nil || IsNil(o.DtUserLastlogondate) {
		var ret string
		return ret
	}
	return *o.DtUserLastlogondate
}

// GetDtUserLastlogondateOk returns a tuple with the DtUserLastlogondate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDtUserLastlogondateOk() (*string, bool) {
	if o == nil || IsNil(o.DtUserLastlogondate) {
		return nil, false
	}
	return o.DtUserLastlogondate, true
}

// HasDtUserLastlogondate returns a boolean if a field has been set.
func (o *UserResponse) HasDtUserLastlogondate() bool {
	if o != nil && !IsNil(o.DtUserLastlogondate) {
		return true
	}

	return false
}

// SetDtUserLastlogondate gets a reference to the given string and assigns it to the DtUserLastlogondate field.
func (o *UserResponse) SetDtUserLastlogondate(v string) {
	o.DtUserLastlogondate = &v
}

// GetDtUserPasswordchanged returns the DtUserPasswordchanged field value if set, zero value otherwise.
func (o *UserResponse) GetDtUserPasswordchanged() string {
	if o == nil || IsNil(o.DtUserPasswordchanged) {
		var ret string
		return ret
	}
	return *o.DtUserPasswordchanged
}

// GetDtUserPasswordchangedOk returns a tuple with the DtUserPasswordchanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDtUserPasswordchangedOk() (*string, bool) {
	if o == nil || IsNil(o.DtUserPasswordchanged) {
		return nil, false
	}
	return o.DtUserPasswordchanged, true
}

// HasDtUserPasswordchanged returns a boolean if a field has been set.
func (o *UserResponse) HasDtUserPasswordchanged() bool {
	if o != nil && !IsNil(o.DtUserPasswordchanged) {
		return true
	}

	return false
}

// SetDtUserPasswordchanged gets a reference to the given string and assigns it to the DtUserPasswordchanged field.
func (o *UserResponse) SetDtUserPasswordchanged(v string) {
	o.DtUserPasswordchanged = &v
}

// GetDtUserEzsignprepaidexpiration returns the DtUserEzsignprepaidexpiration field value if set, zero value otherwise.
func (o *UserResponse) GetDtUserEzsignprepaidexpiration() string {
	if o == nil || IsNil(o.DtUserEzsignprepaidexpiration) {
		var ret string
		return ret
	}
	return *o.DtUserEzsignprepaidexpiration
}

// GetDtUserEzsignprepaidexpirationOk returns a tuple with the DtUserEzsignprepaidexpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetDtUserEzsignprepaidexpirationOk() (*string, bool) {
	if o == nil || IsNil(o.DtUserEzsignprepaidexpiration) {
		return nil, false
	}
	return o.DtUserEzsignprepaidexpiration, true
}

// HasDtUserEzsignprepaidexpiration returns a boolean if a field has been set.
func (o *UserResponse) HasDtUserEzsignprepaidexpiration() bool {
	if o != nil && !IsNil(o.DtUserEzsignprepaidexpiration) {
		return true
	}

	return false
}

// SetDtUserEzsignprepaidexpiration gets a reference to the given string and assigns it to the DtUserEzsignprepaidexpiration field.
func (o *UserResponse) SetDtUserEzsignprepaidexpiration(v string) {
	o.DtUserEzsignprepaidexpiration = &v
}

// GetBUserIsactive returns the BUserIsactive field value
func (o *UserResponse) GetBUserIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BUserIsactive
}

// GetBUserIsactiveOk returns a tuple with the BUserIsactive field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBUserIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUserIsactive, true
}

// SetBUserIsactive sets field value
func (o *UserResponse) SetBUserIsactive(v bool) {
	o.BUserIsactive = v
}

// GetBUserValidatebyadministration returns the BUserValidatebyadministration field value if set, zero value otherwise.
func (o *UserResponse) GetBUserValidatebyadministration() bool {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebyadministration
}

// GetBUserValidatebyadministrationOk returns a tuple with the BUserValidatebyadministration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBUserValidatebyadministrationOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		return nil, false
	}
	return o.BUserValidatebyadministration, true
}

// HasBUserValidatebyadministration returns a boolean if a field has been set.
func (o *UserResponse) HasBUserValidatebyadministration() bool {
	if o != nil && !IsNil(o.BUserValidatebyadministration) {
		return true
	}

	return false
}

// SetBUserValidatebyadministration gets a reference to the given bool and assigns it to the BUserValidatebyadministration field.
func (o *UserResponse) SetBUserValidatebyadministration(v bool) {
	o.BUserValidatebyadministration = &v
}

// GetBUserValidatebydirector returns the BUserValidatebydirector field value if set, zero value otherwise.
func (o *UserResponse) GetBUserValidatebydirector() bool {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebydirector
}

// GetBUserValidatebydirectorOk returns a tuple with the BUserValidatebydirector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBUserValidatebydirectorOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		return nil, false
	}
	return o.BUserValidatebydirector, true
}

// HasBUserValidatebydirector returns a boolean if a field has been set.
func (o *UserResponse) HasBUserValidatebydirector() bool {
	if o != nil && !IsNil(o.BUserValidatebydirector) {
		return true
	}

	return false
}

// SetBUserValidatebydirector gets a reference to the given bool and assigns it to the BUserValidatebydirector field.
func (o *UserResponse) SetBUserValidatebydirector(v bool) {
	o.BUserValidatebydirector = &v
}

// GetBUserAttachmentautoverified returns the BUserAttachmentautoverified field value if set, zero value otherwise.
func (o *UserResponse) GetBUserAttachmentautoverified() bool {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		var ret bool
		return ret
	}
	return *o.BUserAttachmentautoverified
}

// GetBUserAttachmentautoverifiedOk returns a tuple with the BUserAttachmentautoverified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBUserAttachmentautoverifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		return nil, false
	}
	return o.BUserAttachmentautoverified, true
}

// HasBUserAttachmentautoverified returns a boolean if a field has been set.
func (o *UserResponse) HasBUserAttachmentautoverified() bool {
	if o != nil && !IsNil(o.BUserAttachmentautoverified) {
		return true
	}

	return false
}

// SetBUserAttachmentautoverified gets a reference to the given bool and assigns it to the BUserAttachmentautoverified field.
func (o *UserResponse) SetBUserAttachmentautoverified(v bool) {
	o.BUserAttachmentautoverified = &v
}

// GetBUserChangepassword returns the BUserChangepassword field value
func (o *UserResponse) GetBUserChangepassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BUserChangepassword
}

// GetBUserChangepasswordOk returns a tuple with the BUserChangepassword field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetBUserChangepasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUserChangepassword, true
}

// SetBUserChangepassword sets field value
func (o *UserResponse) SetBUserChangepassword(v bool) {
	o.BUserChangepassword = v
}

// GetObjAudit returns the ObjAudit field value
func (o *UserResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *UserResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUserID"] = o.PkiUserID
	if !IsNil(o.FkiAgentID) {
		toSerialize["fkiAgentID"] = o.FkiAgentID
	}
	if !IsNil(o.FkiBrokerID) {
		toSerialize["fkiBrokerID"] = o.FkiBrokerID
	}
	if !IsNil(o.FkiAssistantID) {
		toSerialize["fkiAssistantID"] = o.FkiAssistantID
	}
	if !IsNil(o.FkiEmployeeID) {
		toSerialize["fkiEmployeeID"] = o.FkiEmployeeID
	}
	toSerialize["fkiCompanyIDDefault"] = o.FkiCompanyIDDefault
	toSerialize["sCompanyNameX"] = o.SCompanyNameX
	toSerialize["fkiDepartmentIDDefault"] = o.FkiDepartmentIDDefault
	toSerialize["sDepartmentNameX"] = o.SDepartmentNameX
	toSerialize["fkiTimezoneID"] = o.FkiTimezoneID
	toSerialize["sTimezoneName"] = o.STimezoneName
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sLanguageNameX"] = o.SLanguageNameX
	toSerialize["objEmail"] = o.ObjEmail
	toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	if !IsNil(o.ObjPhoneHome) {
		toSerialize["objPhoneHome"] = o.ObjPhoneHome
	}
	if !IsNil(o.ObjPhoneSMS) {
		toSerialize["objPhoneSMS"] = o.ObjPhoneSMS
	}
	if !IsNil(o.FkiSecretquestionID) {
		toSerialize["fkiSecretquestionID"] = o.FkiSecretquestionID
	}
	if !IsNil(o.FkiModuleIDForm) {
		toSerialize["fkiModuleIDForm"] = o.FkiModuleIDForm
	}
	if !IsNil(o.SModuleNameX) {
		toSerialize["sModuleNameX"] = o.SModuleNameX
	}
	toSerialize["eUserOrigin"] = o.EUserOrigin
	toSerialize["eUserType"] = o.EUserType
	toSerialize["eUserLogintype"] = o.EUserLogintype
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserLoginname"] = o.SUserLoginname
	toSerialize["eUserEzsignaccess"] = o.EUserEzsignaccess
	if !IsNil(o.DtUserLastlogondate) {
		toSerialize["dtUserLastlogondate"] = o.DtUserLastlogondate
	}
	if !IsNil(o.DtUserPasswordchanged) {
		toSerialize["dtUserPasswordchanged"] = o.DtUserPasswordchanged
	}
	if !IsNil(o.DtUserEzsignprepaidexpiration) {
		toSerialize["dtUserEzsignprepaidexpiration"] = o.DtUserEzsignprepaidexpiration
	}
	toSerialize["bUserIsactive"] = o.BUserIsactive
	if !IsNil(o.BUserValidatebyadministration) {
		toSerialize["bUserValidatebyadministration"] = o.BUserValidatebyadministration
	}
	if !IsNil(o.BUserValidatebydirector) {
		toSerialize["bUserValidatebydirector"] = o.BUserValidatebydirector
	}
	if !IsNil(o.BUserAttachmentautoverified) {
		toSerialize["bUserAttachmentautoverified"] = o.BUserAttachmentautoverified
	}
	toSerialize["bUserChangepassword"] = o.BUserChangepassword
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


