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

// checks if the EzsigntemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateResponse{}

// EzsigntemplateResponse A Ezsigntemplate Object
type EzsigntemplateResponse struct {
	// The unique ID of the Ezsigntemplate
	PkiEzsigntemplateID int32 `json:"pkiEzsigntemplateID"`
	// The unique ID of the Ezsigntemplatedocument
	FkiEzsigntemplatedocumentID *int32 `json:"fkiEzsigntemplatedocumentID,omitempty"`
	// The unique ID of the Ezsignfoldertype.
	FkiEzsignfoldertypeID *int32 `json:"fkiEzsignfoldertypeID,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The Name of the Language in the language of the requester
	SLanguageNameX string `json:"sLanguageNameX"`
	// The description of the Ezsigntemplate
	SEzsigntemplateDescription string `json:"sEzsigntemplateDescription"`
	// The filename pattern of the Ezsigntemplate
	SEzsigntemplateFilenamepattern *string `json:"sEzsigntemplateFilenamepattern,omitempty"`
	// Whether the Ezsigntemplate can be accessed by admin users only (eUserType=Normal)
	BEzsigntemplateAdminonly bool `json:"bEzsigntemplateAdminonly"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX *string `json:"sEzsignfoldertypeNameX,omitempty"`
	ObjAudit CommonAudit `json:"objAudit"`
	// Whether the Ezsigntemplate if allowed to edit or not
	BEzsigntemplateEditallowed bool `json:"bEzsigntemplateEditallowed"`
	EEzsigntemplateType *FieldEEzsigntemplateType `json:"eEzsigntemplateType,omitempty"`
}

type _EzsigntemplateResponse EzsigntemplateResponse

// NewEzsigntemplateResponse instantiates a new EzsigntemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateResponse(pkiEzsigntemplateID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, objAudit CommonAudit, bEzsigntemplateEditallowed bool) *EzsigntemplateResponse {
	this := EzsigntemplateResponse{}
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	this.ObjAudit = objAudit
	this.BEzsigntemplateEditallowed = bEzsigntemplateEditallowed
	return &this
}

// NewEzsigntemplateResponseWithDefaults instantiates a new EzsigntemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateResponseWithDefaults() *EzsigntemplateResponse {
	this := EzsigntemplateResponse{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value
func (o *EzsigntemplateResponse) GetPkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateID, true
}

// SetPkiEzsigntemplateID sets field value
func (o *EzsigntemplateResponse) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplateResponse) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzsigntemplatedocumentID, true
}

// HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplateResponse) HasFkiEzsigntemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatedocumentID field.
func (o *EzsigntemplateResponse) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateResponse) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateResponse) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsigntemplateResponse) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateResponse) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSLanguageNameX returns the SLanguageNameX field value
func (o *EzsigntemplateResponse) GetSLanguageNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetSLanguageNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLanguageNameX, true
}

// SetSLanguageNameX sets field value
func (o *EzsigntemplateResponse) SetSLanguageNameX(v string) {
	o.SLanguageNameX = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateResponse) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateResponse) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetSEzsigntemplateFilenamepattern returns the SEzsigntemplateFilenamepattern field value if set, zero value otherwise.
func (o *EzsigntemplateResponse) GetSEzsigntemplateFilenamepattern() string {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateFilenamepattern
}

// GetSEzsigntemplateFilenamepatternOk returns a tuple with the SEzsigntemplateFilenamepattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetSEzsigntemplateFilenamepatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		return nil, false
	}
	return o.SEzsigntemplateFilenamepattern, true
}

// HasSEzsigntemplateFilenamepattern returns a boolean if a field has been set.
func (o *EzsigntemplateResponse) HasSEzsigntemplateFilenamepattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateFilenamepattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateFilenamepattern gets a reference to the given string and assigns it to the SEzsigntemplateFilenamepattern field.
func (o *EzsigntemplateResponse) SetSEzsigntemplateFilenamepattern(v string) {
	o.SEzsigntemplateFilenamepattern = &v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateResponse) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateResponse) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *EzsigntemplateResponse) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzsigntemplateResponse) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *EzsigntemplateResponse) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsigntemplateResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsigntemplateResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

// GetBEzsigntemplateEditallowed returns the BEzsigntemplateEditallowed field value
func (o *EzsigntemplateResponse) GetBEzsigntemplateEditallowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateEditallowed
}

// GetBEzsigntemplateEditallowedOk returns a tuple with the BEzsigntemplateEditallowed field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetBEzsigntemplateEditallowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateEditallowed, true
}

// SetBEzsigntemplateEditallowed sets field value
func (o *EzsigntemplateResponse) SetBEzsigntemplateEditallowed(v bool) {
	o.BEzsigntemplateEditallowed = v
}

// GetEEzsigntemplateType returns the EEzsigntemplateType field value if set, zero value otherwise.
func (o *EzsigntemplateResponse) GetEEzsigntemplateType() FieldEEzsigntemplateType {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		var ret FieldEEzsigntemplateType
		return ret
	}
	return *o.EEzsigntemplateType
}

// GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateResponse) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool) {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		return nil, false
	}
	return o.EEzsigntemplateType, true
}

// HasEEzsigntemplateType returns a boolean if a field has been set.
func (o *EzsigntemplateResponse) HasEEzsigntemplateType() bool {
	if o != nil && !IsNil(o.EEzsigntemplateType) {
		return true
	}

	return false
}

// SetEEzsigntemplateType gets a reference to the given FieldEEzsigntemplateType and assigns it to the EEzsigntemplateType field.
func (o *EzsigntemplateResponse) SetEEzsigntemplateType(v FieldEEzsigntemplateType) {
	o.EEzsigntemplateType = &v
}

func (o EzsigntemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateID"] = o.PkiEzsigntemplateID
	if !IsNil(o.FkiEzsigntemplatedocumentID) {
		toSerialize["fkiEzsigntemplatedocumentID"] = o.FkiEzsigntemplatedocumentID
	}
	if !IsNil(o.FkiEzsignfoldertypeID) {
		toSerialize["fkiEzsignfoldertypeID"] = o.FkiEzsignfoldertypeID
	}
	toSerialize["fkiLanguageID"] = o.FkiLanguageID
	toSerialize["sLanguageNameX"] = o.SLanguageNameX
	toSerialize["sEzsigntemplateDescription"] = o.SEzsigntemplateDescription
	if !IsNil(o.SEzsigntemplateFilenamepattern) {
		toSerialize["sEzsigntemplateFilenamepattern"] = o.SEzsigntemplateFilenamepattern
	}
	toSerialize["bEzsigntemplateAdminonly"] = o.BEzsigntemplateAdminonly
	if !IsNil(o.SEzsignfoldertypeNameX) {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	toSerialize["objAudit"] = o.ObjAudit
	toSerialize["bEzsigntemplateEditallowed"] = o.BEzsigntemplateEditallowed
	if !IsNil(o.EEzsigntemplateType) {
		toSerialize["eEzsigntemplateType"] = o.EEzsigntemplateType
	}
	return toSerialize, nil
}

func (o *EzsigntemplateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateID",
		"fkiLanguageID",
		"sLanguageNameX",
		"sEzsigntemplateDescription",
		"bEzsigntemplateAdminonly",
		"objAudit",
		"bEzsigntemplateEditallowed",
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

	varEzsigntemplateResponse := _EzsigntemplateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplateResponse(varEzsigntemplateResponse)

	return err
}

type NullableEzsigntemplateResponse struct {
	value *EzsigntemplateResponse
	isSet bool
}

func (v NullableEzsigntemplateResponse) Get() *EzsigntemplateResponse {
	return v.value
}

func (v *NullableEzsigntemplateResponse) Set(val *EzsigntemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateResponse(val *EzsigntemplateResponse) *NullableEzsigntemplateResponse {
	return &NullableEzsigntemplateResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


