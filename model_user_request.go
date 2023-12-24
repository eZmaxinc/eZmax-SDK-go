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
	"bytes"
	"fmt"
)

// checks if the UserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRequest{}

// UserRequest A User Object
type UserRequest struct {
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
	ObjEmail EmailRequestCompound `json:"objEmail"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	ObjPhoneHome *PhoneRequestCompound `json:"objPhoneHome,omitempty"`
	ObjPhoneSMS *PhoneRequestCompound `json:"objPhoneSMS,omitempty"`
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
	SUserLoginname string `json:"sUserLoginname"`
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

type _UserRequest UserRequest

// NewUserRequest instantiates a new UserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRequest(fkiCompanyIDDefault int32, fkiDepartmentIDDefault int32, fkiTimezoneID int32, fkiLanguageID int32, objEmail EmailRequestCompound, fkiBillingentityinternalID int32, eUserType FieldEUserType, eUserLogintype FieldEUserLogintype, sUserFirstname string, sUserLastname string, sUserLoginname string, eUserEzsignaccess FieldEUserEzsignaccess, bUserIsactive bool) *UserRequest {
	this := UserRequest{}
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

// NewUserRequestWithDefaults instantiates a new UserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRequestWithDefaults() *UserRequest {
	this := UserRequest{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value if set, zero value otherwise.
func (o *UserRequest) GetPkiUserID() int32 {
	if o == nil || IsNil(o.PkiUserID) {
		var ret int32
		return ret
	}
	return *o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetPkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUserID) {
		return nil, false
	}
	return o.PkiUserID, true
}

// HasPkiUserID returns a boolean if a field has been set.
func (o *UserRequest) HasPkiUserID() bool {
	if o != nil && !IsNil(o.PkiUserID) {
		return true
	}

	return false
}

// SetPkiUserID gets a reference to the given int32 and assigns it to the PkiUserID field.
func (o *UserRequest) SetPkiUserID(v int32) {
	o.PkiUserID = &v
}

// GetFkiAgentID returns the FkiAgentID field value if set, zero value otherwise.
func (o *UserRequest) GetFkiAgentID() int32 {
	if o == nil || IsNil(o.FkiAgentID) {
		var ret int32
		return ret
	}
	return *o.FkiAgentID
}

// GetFkiAgentIDOk returns a tuple with the FkiAgentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiAgentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAgentID) {
		return nil, false
	}
	return o.FkiAgentID, true
}

// HasFkiAgentID returns a boolean if a field has been set.
func (o *UserRequest) HasFkiAgentID() bool {
	if o != nil && !IsNil(o.FkiAgentID) {
		return true
	}

	return false
}

// SetFkiAgentID gets a reference to the given int32 and assigns it to the FkiAgentID field.
func (o *UserRequest) SetFkiAgentID(v int32) {
	o.FkiAgentID = &v
}

// GetFkiBrokerID returns the FkiBrokerID field value if set, zero value otherwise.
func (o *UserRequest) GetFkiBrokerID() int32 {
	if o == nil || IsNil(o.FkiBrokerID) {
		var ret int32
		return ret
	}
	return *o.FkiBrokerID
}

// GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiBrokerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiBrokerID) {
		return nil, false
	}
	return o.FkiBrokerID, true
}

// HasFkiBrokerID returns a boolean if a field has been set.
func (o *UserRequest) HasFkiBrokerID() bool {
	if o != nil && !IsNil(o.FkiBrokerID) {
		return true
	}

	return false
}

// SetFkiBrokerID gets a reference to the given int32 and assigns it to the FkiBrokerID field.
func (o *UserRequest) SetFkiBrokerID(v int32) {
	o.FkiBrokerID = &v
}

// GetFkiAssistantID returns the FkiAssistantID field value if set, zero value otherwise.
func (o *UserRequest) GetFkiAssistantID() int32 {
	if o == nil || IsNil(o.FkiAssistantID) {
		var ret int32
		return ret
	}
	return *o.FkiAssistantID
}

// GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiAssistantIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAssistantID) {
		return nil, false
	}
	return o.FkiAssistantID, true
}

// HasFkiAssistantID returns a boolean if a field has been set.
func (o *UserRequest) HasFkiAssistantID() bool {
	if o != nil && !IsNil(o.FkiAssistantID) {
		return true
	}

	return false
}

// SetFkiAssistantID gets a reference to the given int32 and assigns it to the FkiAssistantID field.
func (o *UserRequest) SetFkiAssistantID(v int32) {
	o.FkiAssistantID = &v
}

// GetFkiEmployeeID returns the FkiEmployeeID field value if set, zero value otherwise.
func (o *UserRequest) GetFkiEmployeeID() int32 {
	if o == nil || IsNil(o.FkiEmployeeID) {
		var ret int32
		return ret
	}
	return *o.FkiEmployeeID
}

// GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiEmployeeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEmployeeID) {
		return nil, false
	}
	return o.FkiEmployeeID, true
}

// HasFkiEmployeeID returns a boolean if a field has been set.
func (o *UserRequest) HasFkiEmployeeID() bool {
	if o != nil && !IsNil(o.FkiEmployeeID) {
		return true
	}

	return false
}

// SetFkiEmployeeID gets a reference to the given int32 and assigns it to the FkiEmployeeID field.
func (o *UserRequest) SetFkiEmployeeID(v int32) {
	o.FkiEmployeeID = &v
}

// GetFkiCompanyIDDefault returns the FkiCompanyIDDefault field value
func (o *UserRequest) GetFkiCompanyIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCompanyIDDefault
}

// GetFkiCompanyIDDefaultOk returns a tuple with the FkiCompanyIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiCompanyIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCompanyIDDefault, true
}

// SetFkiCompanyIDDefault sets field value
func (o *UserRequest) SetFkiCompanyIDDefault(v int32) {
	o.FkiCompanyIDDefault = v
}

// GetFkiDepartmentIDDefault returns the FkiDepartmentIDDefault field value
func (o *UserRequest) GetFkiDepartmentIDDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDepartmentIDDefault
}

// GetFkiDepartmentIDDefaultOk returns a tuple with the FkiDepartmentIDDefault field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiDepartmentIDDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDepartmentIDDefault, true
}

// SetFkiDepartmentIDDefault sets field value
func (o *UserRequest) SetFkiDepartmentIDDefault(v int32) {
	o.FkiDepartmentIDDefault = v
}

// GetFkiTimezoneID returns the FkiTimezoneID field value
func (o *UserRequest) GetFkiTimezoneID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiTimezoneID
}

// GetFkiTimezoneIDOk returns a tuple with the FkiTimezoneID field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiTimezoneIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiTimezoneID, true
}

// SetFkiTimezoneID sets field value
func (o *UserRequest) SetFkiTimezoneID(v int32) {
	o.FkiTimezoneID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetObjEmail returns the ObjEmail field value
func (o *UserRequest) GetObjEmail() EmailRequestCompound {
	if o == nil {
		var ret EmailRequestCompound
		return ret
	}

	return o.ObjEmail
}

// GetObjEmailOk returns a tuple with the ObjEmail field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetObjEmailOk() (*EmailRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEmail, true
}

// SetObjEmail sets field value
func (o *UserRequest) SetObjEmail(v EmailRequestCompound) {
	o.ObjEmail = v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *UserRequest) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *UserRequest) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetObjPhoneHome returns the ObjPhoneHome field value if set, zero value otherwise.
func (o *UserRequest) GetObjPhoneHome() PhoneRequestCompound {
	if o == nil || IsNil(o.ObjPhoneHome) {
		var ret PhoneRequestCompound
		return ret
	}
	return *o.ObjPhoneHome
}

// GetObjPhoneHomeOk returns a tuple with the ObjPhoneHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetObjPhoneHomeOk() (*PhoneRequestCompound, bool) {
	if o == nil || IsNil(o.ObjPhoneHome) {
		return nil, false
	}
	return o.ObjPhoneHome, true
}

// HasObjPhoneHome returns a boolean if a field has been set.
func (o *UserRequest) HasObjPhoneHome() bool {
	if o != nil && !IsNil(o.ObjPhoneHome) {
		return true
	}

	return false
}

// SetObjPhoneHome gets a reference to the given PhoneRequestCompound and assigns it to the ObjPhoneHome field.
func (o *UserRequest) SetObjPhoneHome(v PhoneRequestCompound) {
	o.ObjPhoneHome = &v
}

// GetObjPhoneSMS returns the ObjPhoneSMS field value if set, zero value otherwise.
func (o *UserRequest) GetObjPhoneSMS() PhoneRequestCompound {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		var ret PhoneRequestCompound
		return ret
	}
	return *o.ObjPhoneSMS
}

// GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetObjPhoneSMSOk() (*PhoneRequestCompound, bool) {
	if o == nil || IsNil(o.ObjPhoneSMS) {
		return nil, false
	}
	return o.ObjPhoneSMS, true
}

// HasObjPhoneSMS returns a boolean if a field has been set.
func (o *UserRequest) HasObjPhoneSMS() bool {
	if o != nil && !IsNil(o.ObjPhoneSMS) {
		return true
	}

	return false
}

// SetObjPhoneSMS gets a reference to the given PhoneRequestCompound and assigns it to the ObjPhoneSMS field.
func (o *UserRequest) SetObjPhoneSMS(v PhoneRequestCompound) {
	o.ObjPhoneSMS = &v
}

// GetFkiSecretquestionID returns the FkiSecretquestionID field value if set, zero value otherwise.
func (o *UserRequest) GetFkiSecretquestionID() int32 {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		var ret int32
		return ret
	}
	return *o.FkiSecretquestionID
}

// GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiSecretquestionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiSecretquestionID) {
		return nil, false
	}
	return o.FkiSecretquestionID, true
}

// HasFkiSecretquestionID returns a boolean if a field has been set.
func (o *UserRequest) HasFkiSecretquestionID() bool {
	if o != nil && !IsNil(o.FkiSecretquestionID) {
		return true
	}

	return false
}

// SetFkiSecretquestionID gets a reference to the given int32 and assigns it to the FkiSecretquestionID field.
func (o *UserRequest) SetFkiSecretquestionID(v int32) {
	o.FkiSecretquestionID = &v
}

// GetSUserSecretresponse returns the SUserSecretresponse field value if set, zero value otherwise.
func (o *UserRequest) GetSUserSecretresponse() string {
	if o == nil || IsNil(o.SUserSecretresponse) {
		var ret string
		return ret
	}
	return *o.SUserSecretresponse
}

// GetSUserSecretresponseOk returns a tuple with the SUserSecretresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetSUserSecretresponseOk() (*string, bool) {
	if o == nil || IsNil(o.SUserSecretresponse) {
		return nil, false
	}
	return o.SUserSecretresponse, true
}

// HasSUserSecretresponse returns a boolean if a field has been set.
func (o *UserRequest) HasSUserSecretresponse() bool {
	if o != nil && !IsNil(o.SUserSecretresponse) {
		return true
	}

	return false
}

// SetSUserSecretresponse gets a reference to the given string and assigns it to the SUserSecretresponse field.
func (o *UserRequest) SetSUserSecretresponse(v string) {
	o.SUserSecretresponse = &v
}

// GetFkiModuleIDForm returns the FkiModuleIDForm field value if set, zero value otherwise.
func (o *UserRequest) GetFkiModuleIDForm() int32 {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		var ret int32
		return ret
	}
	return *o.FkiModuleIDForm
}

// GetFkiModuleIDFormOk returns a tuple with the FkiModuleIDForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetFkiModuleIDFormOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiModuleIDForm) {
		return nil, false
	}
	return o.FkiModuleIDForm, true
}

// HasFkiModuleIDForm returns a boolean if a field has been set.
func (o *UserRequest) HasFkiModuleIDForm() bool {
	if o != nil && !IsNil(o.FkiModuleIDForm) {
		return true
	}

	return false
}

// SetFkiModuleIDForm gets a reference to the given int32 and assigns it to the FkiModuleIDForm field.
func (o *UserRequest) SetFkiModuleIDForm(v int32) {
	o.FkiModuleIDForm = &v
}

// GetEUserType returns the EUserType field value
func (o *UserRequest) GetEUserType() FieldEUserType {
	if o == nil {
		var ret FieldEUserType
		return ret
	}

	return o.EUserType
}

// GetEUserTypeOk returns a tuple with the EUserType field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetEUserTypeOk() (*FieldEUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserType, true
}

// SetEUserType sets field value
func (o *UserRequest) SetEUserType(v FieldEUserType) {
	o.EUserType = v
}

// GetEUserLogintype returns the EUserLogintype field value
func (o *UserRequest) GetEUserLogintype() FieldEUserLogintype {
	if o == nil {
		var ret FieldEUserLogintype
		return ret
	}

	return o.EUserLogintype
}

// GetEUserLogintypeOk returns a tuple with the EUserLogintype field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetEUserLogintypeOk() (*FieldEUserLogintype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserLogintype, true
}

// SetEUserLogintype sets field value
func (o *UserRequest) SetEUserLogintype(v FieldEUserLogintype) {
	o.EUserLogintype = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserRequest) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserRequest) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserRequest) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserRequest) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UserRequest) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UserRequest) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetEUserEzsignaccess returns the EUserEzsignaccess field value
func (o *UserRequest) GetEUserEzsignaccess() FieldEUserEzsignaccess {
	if o == nil {
		var ret FieldEUserEzsignaccess
		return ret
	}

	return o.EUserEzsignaccess
}

// GetEUserEzsignaccessOk returns a tuple with the EUserEzsignaccess field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetEUserEzsignaccessOk() (*FieldEUserEzsignaccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserEzsignaccess, true
}

// SetEUserEzsignaccess sets field value
func (o *UserRequest) SetEUserEzsignaccess(v FieldEUserEzsignaccess) {
	o.EUserEzsignaccess = v
}

// GetBUserIsactive returns the BUserIsactive field value
func (o *UserRequest) GetBUserIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BUserIsactive
}

// GetBUserIsactiveOk returns a tuple with the BUserIsactive field value
// and a boolean to check if the value has been set.
func (o *UserRequest) GetBUserIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUserIsactive, true
}

// SetBUserIsactive sets field value
func (o *UserRequest) SetBUserIsactive(v bool) {
	o.BUserIsactive = v
}

// GetBUserValidatebyadministration returns the BUserValidatebyadministration field value if set, zero value otherwise.
func (o *UserRequest) GetBUserValidatebyadministration() bool {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebyadministration
}

// GetBUserValidatebyadministrationOk returns a tuple with the BUserValidatebyadministration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetBUserValidatebyadministrationOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebyadministration) {
		return nil, false
	}
	return o.BUserValidatebyadministration, true
}

// HasBUserValidatebyadministration returns a boolean if a field has been set.
func (o *UserRequest) HasBUserValidatebyadministration() bool {
	if o != nil && !IsNil(o.BUserValidatebyadministration) {
		return true
	}

	return false
}

// SetBUserValidatebyadministration gets a reference to the given bool and assigns it to the BUserValidatebyadministration field.
func (o *UserRequest) SetBUserValidatebyadministration(v bool) {
	o.BUserValidatebyadministration = &v
}

// GetBUserValidatebydirector returns the BUserValidatebydirector field value if set, zero value otherwise.
func (o *UserRequest) GetBUserValidatebydirector() bool {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		var ret bool
		return ret
	}
	return *o.BUserValidatebydirector
}

// GetBUserValidatebydirectorOk returns a tuple with the BUserValidatebydirector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetBUserValidatebydirectorOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserValidatebydirector) {
		return nil, false
	}
	return o.BUserValidatebydirector, true
}

// HasBUserValidatebydirector returns a boolean if a field has been set.
func (o *UserRequest) HasBUserValidatebydirector() bool {
	if o != nil && !IsNil(o.BUserValidatebydirector) {
		return true
	}

	return false
}

// SetBUserValidatebydirector gets a reference to the given bool and assigns it to the BUserValidatebydirector field.
func (o *UserRequest) SetBUserValidatebydirector(v bool) {
	o.BUserValidatebydirector = &v
}

// GetBUserAttachmentautoverified returns the BUserAttachmentautoverified field value if set, zero value otherwise.
func (o *UserRequest) GetBUserAttachmentautoverified() bool {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		var ret bool
		return ret
	}
	return *o.BUserAttachmentautoverified
}

// GetBUserAttachmentautoverifiedOk returns a tuple with the BUserAttachmentautoverified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetBUserAttachmentautoverifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserAttachmentautoverified) {
		return nil, false
	}
	return o.BUserAttachmentautoverified, true
}

// HasBUserAttachmentautoverified returns a boolean if a field has been set.
func (o *UserRequest) HasBUserAttachmentautoverified() bool {
	if o != nil && !IsNil(o.BUserAttachmentautoverified) {
		return true
	}

	return false
}

// SetBUserAttachmentautoverified gets a reference to the given bool and assigns it to the BUserAttachmentautoverified field.
func (o *UserRequest) SetBUserAttachmentautoverified(v bool) {
	o.BUserAttachmentautoverified = &v
}

// GetBUserChangepassword returns the BUserChangepassword field value if set, zero value otherwise.
func (o *UserRequest) GetBUserChangepassword() bool {
	if o == nil || IsNil(o.BUserChangepassword) {
		var ret bool
		return ret
	}
	return *o.BUserChangepassword
}

// GetBUserChangepasswordOk returns a tuple with the BUserChangepassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetBUserChangepasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.BUserChangepassword) {
		return nil, false
	}
	return o.BUserChangepassword, true
}

// HasBUserChangepassword returns a boolean if a field has been set.
func (o *UserRequest) HasBUserChangepassword() bool {
	if o != nil && !IsNil(o.BUserChangepassword) {
		return true
	}

	return false
}

// SetBUserChangepassword gets a reference to the given bool and assigns it to the BUserChangepassword field.
func (o *UserRequest) SetBUserChangepassword(v bool) {
	o.BUserChangepassword = &v
}

func (o UserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRequest) ToMap() (map[string]interface{}, error) {
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

func (o *UserRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUserRequest := _UserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRequest)

	if err != nil {
		return err
	}

	*o = UserRequest(varUserRequest)

	return err
}

type NullableUserRequest struct {
	value *UserRequest
	isSet bool
}

func (v NullableUserRequest) Get() *UserRequest {
	return v.value
}

func (v *NullableUserRequest) Set(val *UserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRequest(val *UserRequest) *NullableUserRequest {
	return &NullableUserRequest{value: val, isSet: true}
}

func (v NullableUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


