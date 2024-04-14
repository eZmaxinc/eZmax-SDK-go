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

// checks if the WebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponse{}

// WebhookResponse A webhook object
type WebhookResponse struct {
	// The unique ID of the Webhook
	PkiWebhookID int32 `json:"pkiWebhookID"`
	// The description of the Webhook
	SWebhookDescription string `json:"sWebhookDescription"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	EWebhookModule FieldEWebhookModule `json:"eWebhookModule"`
	EWebhookEzsignevent *FieldEWebhookEzsignevent `json:"eWebhookEzsignevent,omitempty"`
	EWebhookManagementevent *FieldEWebhookManagementevent `json:"eWebhookManagementevent,omitempty"`
	// The URL of the Webhook callback
	SWebhookUrl string `json:"sWebhookUrl"`
	// The email that will receive the Webhook in case all attempts fail
	SWebhookEmailfailed string `json:"sWebhookEmailfailed"`
	// The Apikey for the Webhook.  This will be hidden if we are not creating or regenerating the Apikey.
	SWebhookApikey *string `json:"sWebhookApikey,omitempty"`
	// The Secret for the Webhook.  This will be hidden if we are not creating or regenerating the Apikey.
	SWebhookSecret *string `json:"sWebhookSecret,omitempty"`
	// Whether the Webhook is active or not
	BWebhookIsactive bool `json:"bWebhookIsactive"`
	// Whether the requests will be signed or not
	BWebhookIssigned bool `json:"bWebhookIssigned"`
	// Wheter the server's SSL certificate should be validated or not. Not recommended to skip for production use
	BWebhookSkipsslvalidation bool `json:"bWebhookSkipsslvalidation"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _WebhookResponse WebhookResponse

// NewWebhookResponse instantiates a new WebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponse(pkiWebhookID int32, sWebhookDescription string, eWebhookModule FieldEWebhookModule, sWebhookUrl string, sWebhookEmailfailed string, bWebhookIsactive bool, bWebhookIssigned bool, bWebhookSkipsslvalidation bool, objAudit CommonAudit) *WebhookResponse {
	this := WebhookResponse{}
	this.PkiWebhookID = pkiWebhookID
	this.SWebhookDescription = sWebhookDescription
	this.EWebhookModule = eWebhookModule
	this.SWebhookUrl = sWebhookUrl
	this.SWebhookEmailfailed = sWebhookEmailfailed
	this.BWebhookIsactive = bWebhookIsactive
	this.BWebhookIssigned = bWebhookIssigned
	this.BWebhookSkipsslvalidation = bWebhookSkipsslvalidation
	this.ObjAudit = objAudit
	return &this
}

// NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseWithDefaults() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// GetPkiWebhookID returns the PkiWebhookID field value
func (o *WebhookResponse) GetPkiWebhookID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiWebhookID
}

// GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetPkiWebhookIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiWebhookID, true
}

// SetPkiWebhookID sets field value
func (o *WebhookResponse) SetPkiWebhookID(v int32) {
	o.PkiWebhookID = v
}

// GetSWebhookDescription returns the SWebhookDescription field value
func (o *WebhookResponse) GetSWebhookDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookDescription
}

// GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookDescription, true
}

// SetSWebhookDescription sets field value
func (o *WebhookResponse) SetSWebhookDescription(v string) {
	o.SWebhookDescription = v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *WebhookResponse) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *WebhookResponse) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *WebhookResponse) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *WebhookResponse) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *WebhookResponse) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *WebhookResponse) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetEWebhookModule returns the EWebhookModule field value
func (o *WebhookResponse) GetEWebhookModule() FieldEWebhookModule {
	if o == nil {
		var ret FieldEWebhookModule
		return ret
	}

	return o.EWebhookModule
}

// GetEWebhookModuleOk returns a tuple with the EWebhookModule field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookModuleOk() (*FieldEWebhookModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebhookModule, true
}

// SetEWebhookModule sets field value
func (o *WebhookResponse) SetEWebhookModule(v FieldEWebhookModule) {
	o.EWebhookModule = v
}

// GetEWebhookEzsignevent returns the EWebhookEzsignevent field value if set, zero value otherwise.
func (o *WebhookResponse) GetEWebhookEzsignevent() FieldEWebhookEzsignevent {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		var ret FieldEWebhookEzsignevent
		return ret
	}
	return *o.EWebhookEzsignevent
}

// GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool) {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		return nil, false
	}
	return o.EWebhookEzsignevent, true
}

// HasEWebhookEzsignevent returns a boolean if a field has been set.
func (o *WebhookResponse) HasEWebhookEzsignevent() bool {
	if o != nil && !IsNil(o.EWebhookEzsignevent) {
		return true
	}

	return false
}

// SetEWebhookEzsignevent gets a reference to the given FieldEWebhookEzsignevent and assigns it to the EWebhookEzsignevent field.
func (o *WebhookResponse) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent) {
	o.EWebhookEzsignevent = &v
}

// GetEWebhookManagementevent returns the EWebhookManagementevent field value if set, zero value otherwise.
func (o *WebhookResponse) GetEWebhookManagementevent() FieldEWebhookManagementevent {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		var ret FieldEWebhookManagementevent
		return ret
	}
	return *o.EWebhookManagementevent
}

// GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool) {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		return nil, false
	}
	return o.EWebhookManagementevent, true
}

// HasEWebhookManagementevent returns a boolean if a field has been set.
func (o *WebhookResponse) HasEWebhookManagementevent() bool {
	if o != nil && !IsNil(o.EWebhookManagementevent) {
		return true
	}

	return false
}

// SetEWebhookManagementevent gets a reference to the given FieldEWebhookManagementevent and assigns it to the EWebhookManagementevent field.
func (o *WebhookResponse) SetEWebhookManagementevent(v FieldEWebhookManagementevent) {
	o.EWebhookManagementevent = &v
}

// GetSWebhookUrl returns the SWebhookUrl field value
func (o *WebhookResponse) GetSWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookUrl
}

// GetSWebhookUrlOk returns a tuple with the SWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookUrl, true
}

// SetSWebhookUrl sets field value
func (o *WebhookResponse) SetSWebhookUrl(v string) {
	o.SWebhookUrl = v
}

// GetSWebhookEmailfailed returns the SWebhookEmailfailed field value
func (o *WebhookResponse) GetSWebhookEmailfailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEmailfailed
}

// GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookEmailfailedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebhookEmailfailed, true
}

// SetSWebhookEmailfailed sets field value
func (o *WebhookResponse) SetSWebhookEmailfailed(v string) {
	o.SWebhookEmailfailed = v
}

// GetSWebhookApikey returns the SWebhookApikey field value if set, zero value otherwise.
func (o *WebhookResponse) GetSWebhookApikey() string {
	if o == nil || IsNil(o.SWebhookApikey) {
		var ret string
		return ret
	}
	return *o.SWebhookApikey
}

// GetSWebhookApikeyOk returns a tuple with the SWebhookApikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookApikeyOk() (*string, bool) {
	if o == nil || IsNil(o.SWebhookApikey) {
		return nil, false
	}
	return o.SWebhookApikey, true
}

// HasSWebhookApikey returns a boolean if a field has been set.
func (o *WebhookResponse) HasSWebhookApikey() bool {
	if o != nil && !IsNil(o.SWebhookApikey) {
		return true
	}

	return false
}

// SetSWebhookApikey gets a reference to the given string and assigns it to the SWebhookApikey field.
func (o *WebhookResponse) SetSWebhookApikey(v string) {
	o.SWebhookApikey = &v
}

// GetSWebhookSecret returns the SWebhookSecret field value if set, zero value otherwise.
func (o *WebhookResponse) GetSWebhookSecret() string {
	if o == nil || IsNil(o.SWebhookSecret) {
		var ret string
		return ret
	}
	return *o.SWebhookSecret
}

// GetSWebhookSecretOk returns a tuple with the SWebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SWebhookSecret) {
		return nil, false
	}
	return o.SWebhookSecret, true
}

// HasSWebhookSecret returns a boolean if a field has been set.
func (o *WebhookResponse) HasSWebhookSecret() bool {
	if o != nil && !IsNil(o.SWebhookSecret) {
		return true
	}

	return false
}

// SetSWebhookSecret gets a reference to the given string and assigns it to the SWebhookSecret field.
func (o *WebhookResponse) SetSWebhookSecret(v string) {
	o.SWebhookSecret = &v
}

// GetBWebhookIsactive returns the BWebhookIsactive field value
func (o *WebhookResponse) GetBWebhookIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookIsactive
}

// GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetBWebhookIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookIsactive, true
}

// SetBWebhookIsactive sets field value
func (o *WebhookResponse) SetBWebhookIsactive(v bool) {
	o.BWebhookIsactive = v
}

// GetBWebhookIssigned returns the BWebhookIssigned field value
func (o *WebhookResponse) GetBWebhookIssigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookIssigned
}

// GetBWebhookIssignedOk returns a tuple with the BWebhookIssigned field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetBWebhookIssignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookIssigned, true
}

// SetBWebhookIssigned sets field value
func (o *WebhookResponse) SetBWebhookIssigned(v bool) {
	o.BWebhookIssigned = v
}

// GetBWebhookSkipsslvalidation returns the BWebhookSkipsslvalidation field value
func (o *WebhookResponse) GetBWebhookSkipsslvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BWebhookSkipsslvalidation
}

// GetBWebhookSkipsslvalidationOk returns a tuple with the BWebhookSkipsslvalidation field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetBWebhookSkipsslvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BWebhookSkipsslvalidation, true
}

// SetBWebhookSkipsslvalidation sets field value
func (o *WebhookResponse) SetBWebhookSkipsslvalidation(v bool) {
	o.BWebhookSkipsslvalidation = v
}

// GetObjAudit returns the ObjAudit field value
func (o *WebhookResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *WebhookResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o WebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiWebhookID"] = o.PkiWebhookID
	toSerialize["sWebhookDescription"] = o.SWebhookDescription
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	toSerialize["eWebhookModule"] = o.EWebhookModule
	if !IsNil(o.EWebhookEzsignevent) {
		toSerialize["eWebhookEzsignevent"] = o.EWebhookEzsignevent
	}
	if !IsNil(o.EWebhookManagementevent) {
		toSerialize["eWebhookManagementevent"] = o.EWebhookManagementevent
	}
	toSerialize["sWebhookUrl"] = o.SWebhookUrl
	toSerialize["sWebhookEmailfailed"] = o.SWebhookEmailfailed
	if !IsNil(o.SWebhookApikey) {
		toSerialize["sWebhookApikey"] = o.SWebhookApikey
	}
	if !IsNil(o.SWebhookSecret) {
		toSerialize["sWebhookSecret"] = o.SWebhookSecret
	}
	toSerialize["bWebhookIsactive"] = o.BWebhookIsactive
	toSerialize["bWebhookIssigned"] = o.BWebhookIssigned
	toSerialize["bWebhookSkipsslvalidation"] = o.BWebhookSkipsslvalidation
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *WebhookResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiWebhookID",
		"sWebhookDescription",
		"eWebhookModule",
		"sWebhookUrl",
		"sWebhookEmailfailed",
		"bWebhookIsactive",
		"bWebhookIssigned",
		"bWebhookSkipsslvalidation",
		"objAudit",
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

	varWebhookResponse := _WebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResponse)

	if err != nil {
		return err
	}

	*o = WebhookResponse(varWebhookResponse)

	return err
}

type NullableWebhookResponse struct {
	value *WebhookResponse
	isSet bool
}

func (v NullableWebhookResponse) Get() *WebhookResponse {
	return v.value
}

func (v *NullableWebhookResponse) Set(val *WebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponse(val *WebhookResponse) *NullableWebhookResponse {
	return &NullableWebhookResponse{value: val, isSet: true}
}

func (v NullableWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


