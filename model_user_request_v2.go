/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRequestV2{}

// UserRequestV2 A User Object
type UserRequestV2 struct {
	// The unique ID of the User
	PkiUserID *int32 `json:"pkiUserID,omitempty"`
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
	// The unique ID of the Department
	FkiDepartmentIDDefault int32 `json:"fkiDepartmentIDDefault"`
	// The unique ID of the Timezone
	FkiTimezoneID int32 `json:"fkiTimezoneID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// An Email Object and children to create a complete structure
	ObjEmail EmailRequest `json:"objEmail"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	// A Phone Object and children to create a complete structure
	ObjPhoneHome *PhoneRequestV2 `json:"objPhoneHome,omitempty"`
	// A Phone Object and children to create a complete structure
	ObjPhoneSMS *PhoneRequestV2 `json:"objPhoneSMS,omitempty"`
	// The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father's middle name| |15|Your mother's maiden name| |16|Name of your eldest child| |17|Your spouse's middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat's name| |22|Date of Birth (YYYY-MM-DD)| |22|Secret Code| |22|Your reference code|
	FkiSecretquestionID *int32 `json:"fkiSecretquestionID,omitempty"`
	// The answer to the Secretquestion
	SUserSecretresponse *string `json:"sUserSecretresponse,omitempty"`
	// The unique ID of the Module
	FkiModuleIDForm *int32 `json:"fkiModuleIDForm,omitempty"`
	EUserType FieldEUserType `json:"eUserType"`
	EUserLogintype FieldEUserLogintype `json:"eUserLogintype"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The login name of the User.
	SUserLoginname string "json:\"sUserLoginname\" validate:\"regexp=^(?:([\\\\w.%+\\\\-!#$%&'*+\\/=?^`{|}~]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,20})|([a-zA-Z0-9]){1,32})$\""
	// The job title of the user
	SUserJobtitle *string `json:"sUserJobtitle,omitempty" validate:"regexp=^.{0,50}$"`
	EUserEzsignaccess FieldEUserEzsignaccess `json:"eUserEzsignaccess"`
	// Whether the User is active or not
	BUserIsactive bool `json:"bUserIsactive"`
	// Whether if the transactions in which the User is implicated must be validated by administrative personnel or not
	BUserValidatebyadministration *bool `json:"bUserValidatebyadministration,omitempty"`
	// Whether if the transactions in which the User is implicated must be validated by a director or not
	BUserValidatebydirector *bool `json:"bUserValidatebydirector,omitempty"`
	// Whether if Attachments uploaded by the User must be validated or not
	BUserAttachmentautoverified *bool `json:"bUserAttachmentautoverified,omitempty"`
	// Whether if the User is forced to change its password
	BUserChangepassword *bool `json:"bUserChangepassword,omitempty"`
}

type _UserRequestV2 UserRequestV2

// NewUserRequestV2 instantiates a new UserRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRequestV2(fkiCompanyIDDefault int32, fkiDepartmentIDDefault int32, fkiTimezoneID int32, fkiLanguageID int32, objEmail EmailRequest, fkiBillingentityinternalID int32, eUserType FieldEUserType, eUserLogintype FieldEUserLogintype, sUserFirstname string, sUserLastname string, sUserLoginname string, eUserEzsignaccess FieldEUserEzsignaccess, bUserIsactive bool) *UserRequestV2 {
	this := UserRequestV2{}
	this.FkiCompanyIDDefault = fkiCompanyIDDefault
	this.FkiDepartmentIDDefault = fkiDepartmentIDDefault
	this.FkiTimezoneID = fkiTimezoneID
	this.FkiLanguageID = fkiLanguageID
	this.ObjEmail = objEmail
	this.FkiBillingentityinternalID = fkiBillingentityinternalID
	this.EUserType = eUserType
	this.EUserLogintype = eUserLogintype
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.EUserEzsignaccess = eUserEzsignaccess
	this.BUserIsactive = bUserIsactive
	return &this
}

// NewUserRequestV2WithDefaults instantiates a new UserRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRequestV2WithDefaults() *UserRequestV2 {
	this := UserRequestV2{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value if set, zero value otherwise.
func (o *UserRequestV2) GetPkiUserID() int32 {
	if o == nil || IsNil(o.PkiUserID) {
		var ret int32
		return ret
	}
	return *o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetPkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUserID) {
		return nil, false
	}
	return o.PkiUserID, true
}

// HasPkiUserID returns a boolean if a field has been set.
func (o *UserRequestV2) HasPkiUserID() bool {
	if o != nil && !IsNil(o.PkiUserID) {
		return true
	}

	return false
}

// SetPkiUserID gets a reference to the given int32 and assigns it to the PkiUserID field.
func (o *UserRequestV2) SetPkiUserID(v int32) {
	o.PkiUserID = &v
}

// GetFkiAgentID returns the FkiAgentID field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiAgentID() int32 {
	if o == nil || IsNil(o.FkiAgentID) {
		var ret int32
		return ret
	}
	return *o.FkiAgentID
}

// GetFkiAgentIDOk returns a tuple with the FkiAgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiAgentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAgentID) {
		return nil, false
	}
	return o.FkiAgentID, true
}

// HasFkiAgentID returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiAgentID() bool {
	if o != nil && !IsNil(o.FkiAgentID) {
		return true
	}

	return false
}

// SetFkiAgentID gets a reference to the given int32 and assigns it to the FkiAgentID field.
func (o *UserRequestV2) SetFkiAgentID(v int32) {
	o.FkiAgentID = &v
}

// GetFkiBrokerID returns the FkiBrokerID field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiBrokerID() int32 {
	if o == nil || IsNil(o.FkiBrokerID) {
		var ret int32
		return ret
	}
	return *o.FkiBrokerID
}

// GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiBrokerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrokerID) {
		return nil, false
	}
	return o.FkiBrokerID, true
}

// HasFkiBrokerID returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiBrokerID() bool {
	if o != nil && !IsNil(o.FkiBrokerID) {
		return true
	}

	return false
}

// SetFkiBrokerID gets a reference to the given int32 and assigns it to the FkiBrokerID field.
func (o *UserRequestV2) SetFkiBrokerID(v int32) {
	o.FkiBrokerID = &v
}

// GetFkiAssistantID returns the FkiAssistantID field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiAssistantID() int32 {
	if o == nil || IsNil(o.FkiAssistantID) {
		var ret int32
		return ret
	}
	return *o.FkiAssistantID
}

// GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiAssistantIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAssistantID) {
		return nil, false
	}
	return o.FkiAssistantID, true
}

// HasFkiAssistantID returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiAssistantID() bool {
	if o != nil && !IsNil(o.FkiAssistantID) {
		return true
	}

	return false
}

// SetFkiAssistantID gets a reference to the given int32 and assigns it to the FkiAssistantID field.
func (o *UserRequestV2) SetFkiAssistantID(v int32) {
	o.FkiAssistantID = &v
}

// GetFkiEmployeeID returns the FkiEmployeeID field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiEmployeeID() int32 {
	if o == nil || IsNil(o.FkiEmployeeID) {
		var ret int32
		return ret
	}
	return *o.FkiEmployeeID
}

// GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiEmployeeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEmployeeID) {
		return nil, false
	}
	return o.FkiEmployeeID, true
}

// HasFkiEmployeeID returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiEmployeeID() bool {
	if o != nil && !IsNil(o.FkiEmployeeID) {
		return true
	}

	return false
}

// SetFkiEmployeeID gets a reference to the given int32 and assigns it to the FkiEmployeeID field.
func (o *UserRequestV2) SetFkiEmployeeID(v int32) {
	o.FkiEmployeeID = &v
}

// GetFkiCompanyIDDefault returns the FkiCompanyIDDefault field value
func (o *UserRequestV2) GetFkiCompanyIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCompanyIDDefault
}

// GetFkiCompanyIDDefaultOk returns a tuple with the FkiCompanyIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiCompanyIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCompanyIDDefault, true
}

// SetFkiCompanyIDDefault sets field value
func (o *UserRequestV2) SetFkiCompanyIDDefault(v int32) {
	o.FkiCompanyIDDefault = v
}

// GetFkiDepartmentIDDefault returns the FkiDepartmentIDDefault field value
func (o *UserRequestV2) GetFkiDepartmentIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDepartmentIDDefault
}

// GetFkiDepartmentIDDefaultOk returns a tuple with the FkiDepartmentIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiDepartmentIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDepartmentIDDefault, true
}

// SetFkiDepartmentIDDefault sets field value
func (o *UserRequestV2) SetFkiDepartmentIDDefault(v int32) {
	o.FkiDepartmentIDDefault = v
}

// GetFkiTimezoneID returns the FkiTimezoneID field value
func (o *UserRequestV2) GetFkiTimezoneID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiTimezoneID
}

// GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiTimezoneIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiTimezoneID, true
}

// SetFkiTimezoneID sets field value
func (o *UserRequestV2) SetFkiTimezoneID(v int32) {
	o.FkiTimezoneID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserRequestV2) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserRequestV2) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetObjEmail returns the ObjEmail field value
func (o *UserRequestV2) GetObjEmail() EmailRequest {
	if o == nil {
		var ret EmailRequest
		return ret
	}

	return o.ObjEmail
}

// GetObjEmailOk returns a tuple with the ObjEmail field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetObjEmailOk() (*EmailRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEmail, true
}

// SetObjEmail sets field value
func (o *UserRequestV2) SetObjEmail(v EmailRequest) {
	o.ObjEmail = v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *UserRequestV2) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *UserRequestV2) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetObjPhoneHome returns the ObjPhoneHome field value if set, zero value otherwise.
func (o *UserRequestV2) GetObjPhoneHome() PhoneRequestV2 {
	if o == nil || IsNil(o.ObjPhoneHome) {
		var ret PhoneRequestV2
		return ret
	}
	return *o.ObjPhoneHome
}

// GetObjPhoneHomeOk returns a tuple with the ObjPhoneHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetObjPhoneHomeOk() (*PhoneRequestV2, bool) {
	if o == nil || IsNil(o.ObjPhoneHome) {
		return nil, false
	}
	return o.ObjPhoneHome, true
}

// HasObjPhoneHome returns a boolean if a field has been set.
func (o *UserRequestV2) HasObjPhoneHome() bool {
	if o != nil && !IsNil(o.ObjPhoneHome) {
		return true
	}

	return false
}

// SetObjPhoneHome gets a reference to the given PhoneRequestV2 and assigns it to the ObjPhoneHome field.
func (o *UserRequestV2) SetObjPhoneHome(v PhoneRequestV2) {
	o.ObjPhoneHome = &v
}

// GetObjPhoneSMS returns the ObjPhoneSMS field value if set, zero value otherwise.
func (o *UserRequestV2) GetObjPhoneSMS() PhoneRequestV2 {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		var ret PhoneRequestV2
		return ret
	}
	return *o.ObjPhoneSMS
}

// GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetObjPhoneSMSOk() (*PhoneRequestV2, bool) {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		return nil, false
	}
	return o.ObjPhoneSMS, true
}

// HasObjPhoneSMS returns a boolean if a field has been set.
func (o *UserRequestV2) HasObjPhoneSMS() bool {
	if o != nil && !IsNil(o.ObjPhoneSMS) {
		return true
	}

	return false
}

// SetObjPhoneSMS gets a reference to the given PhoneRequestV2 and assigns it to the ObjPhoneSMS field.
func (o *UserRequestV2) SetObjPhoneSMS(v PhoneRequestV2) {
	o.ObjPhoneSMS = &v
}

// GetFkiSecretquestionID returns the FkiSecretquestionID field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiSecretquestionID() int32 {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		var ret int32
		return ret
	}
	return *o.FkiSecretquestionID
}

// GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiSecretquestionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		return nil, false
	}
	return o.FkiSecretquestionID, true
}

// HasFkiSecretquestionID returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiSecretquestionID() bool {
	if o != nil && !IsNil(o.FkiSecretquestionID) {
		return true
	}

	return false
}

// SetFkiSecretquestionID gets a reference to the given int32 and assigns it to the FkiSecretquestionID field.
func (o *UserRequestV2) SetFkiSecretquestionID(v int32) {
	o.FkiSecretquestionID = &v
}

// GetSUserSecretresponse returns the SUserSecretresponse field value if set, zero value otherwise.
func (o *UserRequestV2) GetSUserSecretresponse() string {
	if o == nil || IsNil(o.SUserSecretresponse) {
		var ret string
		return ret
	}
	return *o.SUserSecretresponse
}

// GetSUserSecretresponseOk returns a tuple with the SUserSecretresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetSUserSecretresponseOk() (*string, bool) {
	if o == nil || IsNil(o.SUserSecretresponse) {
		return nil, false
	}
	return o.SUserSecretresponse, true
}

// HasSUserSecretresponse returns a boolean if a field has been set.
func (o *UserRequestV2) HasSUserSecretresponse() bool {
	if o != nil && !IsNil(o.SUserSecretresponse) {
		return true
	}

	return false
}

// SetSUserSecretresponse gets a reference to the given string and assigns it to the SUserSecretresponse field.
func (o *UserRequestV2) SetSUserSecretresponse(v string) {
	o.SUserSecretresponse = &v
}

// GetFkiModuleIDForm returns the FkiModuleIDForm field value if set, zero value otherwise.
func (o *UserRequestV2) GetFkiModuleIDForm() int32 {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		var ret int32
		return ret
	}
	return *o.FkiModuleIDForm
}

// GetFkiModuleIDFormOk returns a tuple with the FkiModuleIDForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetFkiModuleIDFormOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		return nil, false
	}
	return o.FkiModuleIDForm, true
}

// HasFkiModuleIDForm returns a boolean if a field has been set.
func (o *UserRequestV2) HasFkiModuleIDForm() bool {
	if o != nil && !IsNil(o.FkiModuleIDForm) {
		return true
	}

	return false
}

// SetFkiModuleIDForm gets a reference to the given int32 and assigns it to the FkiModuleIDForm field.
func (o *UserRequestV2) SetFkiModuleIDForm(v int32) {
	o.FkiModuleIDForm = &v
}

// GetEUserType returns the EUserType field value
func (o *UserRequestV2) GetEUserType() FieldEUserType {
	if o == nil {
		var ret FieldEUserType
		return ret
	}

	return o.EUserType
}

// GetEUserTypeOk returns a tuple with the EUserType field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetEUserTypeOk() (*FieldEUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserType, true
}

// SetEUserType sets field value
func (o *UserRequestV2) SetEUserType(v FieldEUserType) {
	o.EUserType = v
}

// GetEUserLogintype returns the EUserLogintype field value
func (o *UserRequestV2) GetEUserLogintype() FieldEUserLogintype {
	if o == nil {
		var ret FieldEUserLogintype
		return ret
	}

	return o.EUserLogintype
}

// GetEUserLogintypeOk returns a tuple with the EUserLogintype field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetEUserLogintypeOk() (*FieldEUserLogintype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserLogintype, true
}

// SetEUserLogintype sets field value
func (o *UserRequestV2) SetEUserLogintype(v FieldEUserLogintype) {
	o.EUserLogintype = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserRequestV2) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserRequestV2) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserRequestV2) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserRequestV2) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UserRequestV2) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UserRequestV2) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetSUserJobtitle returns the SUserJobtitle field value if set, zero value otherwise.
func (o *UserRequestV2) GetSUserJobtitle() string {
	if o == nil || IsNil(o.SUserJobtitle) {
		var ret string
		return ret
	}
	return *o.SUserJobtitle
}

// GetSUserJobtitleOk returns a tuple with the SUserJobtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetSUserJobtitleOk() (*string, bool) {
	if o == nil || IsNil(o.SUserJobtitle) {
		return nil, false
	}
	return o.SUserJobtitle, true
}

// HasSUserJobtitle returns a boolean if a field has been set.
func (o *UserRequestV2) HasSUserJobtitle() bool {
	if o != nil && !IsNil(o.SUserJobtitle) {
		return true
	}

	return false
}

// SetSUserJobtitle gets a reference to the given string and assigns it to the SUserJobtitle field.
func (o *UserRequestV2) SetSUserJobtitle(v string) {
	o.SUserJobtitle = &v
}

// GetEUserEzsignaccess returns the EUserEzsignaccess field value
func (o *UserRequestV2) GetEUserEzsignaccess() FieldEUserEzsignaccess {
	if o == nil {
		var ret FieldEUserEzsignaccess
		return ret
	}

	return o.EUserEzsignaccess
}

// GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserEzsignaccess, true
}

// SetEUserEzsignaccess sets field value
func (o *UserRequestV2) SetEUserEzsignaccess(v FieldEUserEzsignaccess) {
	o.EUserEzsignaccess = v
}

// GetBUserIsactive returns the BUserIsactive field value
func (o *UserRequestV2) GetBUserIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BUserIsactive
}

// GetBUserIsactiveOk returns a tuple with the BUserIsactive field value
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetBUserIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUserIsactive, true
}

// SetBUserIsactive sets field value
func (o *UserRequestV2) SetBUserIsactive(v bool) {
	o.BUserIsactive = v
}

// GetBUserValidatebyadministration returns the BUserValidatebyadministration field value if set, zero value otherwise.
func (o *UserRequestV2) GetBUserValidatebyadministration() bool {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebyadministration
}

// GetBUserValidatebyadministrationOk returns a tuple with the BUserValidatebyadministration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetBUserValidatebyadministrationOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		return nil, false
	}
	return o.BUserValidatebyadministration, true
}

// HasBUserValidatebyadministration returns a boolean if a field has been set.
func (o *UserRequestV2) HasBUserValidatebyadministration() bool {
	if o != nil && !IsNil(o.BUserValidatebyadministration) {
		return true
	}

	return false
}

// SetBUserValidatebyadministration gets a reference to the given bool and assigns it to the BUserValidatebyadministration field.
func (o *UserRequestV2) SetBUserValidatebyadministration(v bool) {
	o.BUserValidatebyadministration = &v
}

// GetBUserValidatebydirector returns the BUserValidatebydirector field value if set, zero value otherwise.
func (o *UserRequestV2) GetBUserValidatebydirector() bool {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebydirector
}

// GetBUserValidatebydirectorOk returns a tuple with the BUserValidatebydirector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetBUserValidatebydirectorOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		return nil, false
	}
	return o.BUserValidatebydirector, true
}

// HasBUserValidatebydirector returns a boolean if a field has been set.
func (o *UserRequestV2) HasBUserValidatebydirector() bool {
	if o != nil && !IsNil(o.BUserValidatebydirector) {
		return true
	}

	return false
}

// SetBUserValidatebydirector gets a reference to the given bool and assigns it to the BUserValidatebydirector field.
func (o *UserRequestV2) SetBUserValidatebydirector(v bool) {
	o.BUserValidatebydirector = &v
}

// GetBUserAttachmentautoverified returns the BUserAttachmentautoverified field value if set, zero value otherwise.
func (o *UserRequestV2) GetBUserAttachmentautoverified() bool {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		var ret bool
		return ret
	}
	return *o.BUserAttachmentautoverified
}

// GetBUserAttachmentautoverifiedOk returns a tuple with the BUserAttachmentautoverified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetBUserAttachmentautoverifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		return nil, false
	}
	return o.BUserAttachmentautoverified, true
}

// HasBUserAttachmentautoverified returns a boolean if a field has been set.
func (o *UserRequestV2) HasBUserAttachmentautoverified() bool {
	if o != nil && !IsNil(o.BUserAttachmentautoverified) {
		return true
	}

	return false
}

// SetBUserAttachmentautoverified gets a reference to the given bool and assigns it to the BUserAttachmentautoverified field.
func (o *UserRequestV2) SetBUserAttachmentautoverified(v bool) {
	o.BUserAttachmentautoverified = &v
}

// GetBUserChangepassword returns the BUserChangepassword field value if set, zero value otherwise.
func (o *UserRequestV2) GetBUserChangepassword() bool {
	if o == nil || IsNil(o.BUserChangepassword) {
		var ret bool
		return ret
	}
	return *o.BUserChangepassword
}

// GetBUserChangepasswordOk returns a tuple with the BUserChangepassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequestV2) GetBUserChangepasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserChangepassword) {
		return nil, false
	}
	return o.BUserChangepassword, true
}

// HasBUserChangepassword returns a boolean if a field has been set.
func (o *UserRequestV2) HasBUserChangepassword() bool {
	if o != nil && !IsNil(o.BUserChangepassword) {
		return true
	}

	return false
}

// SetBUserChangepassword gets a reference to the given bool and assigns it to the BUserChangepassword field.
func (o *UserRequestV2) SetBUserChangepassword(v bool) {
	o.BUserChangepassword = &v
}

func (o UserRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiUserID) {
		toSerialize["pkiUserID"] = o.PkiUserID
	}
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
	toSerialize["fkiDepartmentIDDefault"] = o.FkiDepartmentIDDefault
	toSerialize["fkiTimezoneID"] = o.FkiTimezoneID
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["objEmail"] = o.ObjEmail
	toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	if !IsNil(o.ObjPhoneHome) {
		toSerialize["objPhoneHome"] = o.ObjPhoneHome
	}
	if !IsNil(o.ObjPhoneSMS) {
		toSerialize["objPhoneSMS"] = o.ObjPhoneSMS
	}
	if !IsNil(o.FkiSecretquestionID) {
		toSerialize["fkiSecretquestionID"] = o.FkiSecretquestionID
	}
	if !IsNil(o.SUserSecretresponse) {
		toSerialize["sUserSecretresponse"] = o.SUserSecretresponse
	}
	if !IsNil(o.FkiModuleIDForm) {
		toSerialize["fkiModuleIDForm"] = o.FkiModuleIDForm
	}
	toSerialize["eUserType"] = o.EUserType
	toSerialize["eUserLogintype"] = o.EUserLogintype
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserLoginname"] = o.SUserLoginname
	if !IsNil(o.SUserJobtitle) {
		toSerialize["sUserJobtitle"] = o.SUserJobtitle
	}
	toSerialize["eUserEzsignaccess"] = o.EUserEzsignaccess
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
	if !IsNil(o.BUserChangepassword) {
		toSerialize["bUserChangepassword"] = o.BUserChangepassword
	}
	return toSerialize, nil
}

func (o *UserRequestV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiCompanyIDDefault",
		"fkiDepartmentIDDefault",
		"fkiTimezoneID",
		"fkiLanguageID",
		"objEmail",
		"fkiBillingentityinternalID",
		"eUserType",
		"eUserLogintype",
		"sUserFirstname",
		"sUserLastname",
		"sUserLoginname",
		"eUserEzsignaccess",
		"bUserIsactive",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserRequestV2 := _UserRequestV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRequestV2)

	if err != nil {
		return err
	}

	*o = UserRequestV2(varUserRequestV2)

	return err
}

type NullableUserRequestV2 struct {
	value *UserRequestV2
	isSet bool
}

func (v NullableUserRequestV2) Get() *UserRequestV2 {
	return v.value
}

func (v *NullableUserRequestV2) Set(val *UserRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRequestV2(val *UserRequestV2) *NullableUserRequestV2 {
	return &NullableUserRequestV2{value: val, isSet: true}
}

func (v NullableUserRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


