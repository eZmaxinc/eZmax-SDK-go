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

// checks if the EzsigntemplateGetObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateGetObjectV1ResponseMPayload{}

// EzsigntemplateGetObjectV1ResponseMPayload Payload for GET /1/object/ezsigntemplate/{pkiEzsigntemplateID}
type EzsigntemplateGetObjectV1ResponseMPayload struct {
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
	ObjEzsigntemplatedocument *EzsigntemplatedocumentResponse `json:"objEzsigntemplatedocument,omitempty"`
	AObjEzsigntemplatesigner []EzsigntemplatesignerResponseCompound `json:"a_objEzsigntemplatesigner"`
}

type _EzsigntemplateGetObjectV1ResponseMPayload EzsigntemplateGetObjectV1ResponseMPayload

// NewEzsigntemplateGetObjectV1ResponseMPayload instantiates a new EzsigntemplateGetObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateGetObjectV1ResponseMPayload(pkiEzsigntemplateID int32, fkiLanguageID int32, sLanguageNameX string, sEzsigntemplateDescription string, bEzsigntemplateAdminonly bool, objAudit CommonAudit, bEzsigntemplateEditallowed bool, aObjEzsigntemplatesigner []EzsigntemplatesignerResponseCompound) *EzsigntemplateGetObjectV1ResponseMPayload {
	this := EzsigntemplateGetObjectV1ResponseMPayload{}
	this.PkiEzsigntemplateID = pkiEzsigntemplateID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.SEzsigntemplateDescription = sEzsigntemplateDescription
	this.BEzsigntemplateAdminonly = bEzsigntemplateAdminonly
	this.ObjAudit = objAudit
	this.BEzsigntemplateEditallowed = bEzsigntemplateEditallowed
	this.AObjEzsigntemplatesigner = aObjEzsigntemplatesigner
	return &this
}

// NewEzsigntemplateGetObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplateGetObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateGetObjectV1ResponseMPayloadWithDefaults() *EzsigntemplateGetObjectV1ResponseMPayload {
	this := EzsigntemplateGetObjectV1ResponseMPayload{}
	return &this
}

// GetPkiEzsigntemplateID returns the PkiEzsigntemplateID field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetPkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateID
}

// GetPkiEzsigntemplateIDOk returns a tuple with the PkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetPkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateID, true
}

// SetPkiEzsigntemplateID sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetPkiEzsigntemplateID(v int32) {
	o.PkiEzsigntemplateID = v
}

// GetFkiEzsigntemplatedocumentID returns the FkiEzsigntemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsigntemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatedocumentID
}

// GetFkiEzsigntemplatedocumentIDOk returns a tuple with the FkiEzsigntemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsigntemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzsigntemplatedocumentID, true
}

// HasFkiEzsigntemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasFkiEzsigntemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatedocumentID field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiEzsigntemplatedocumentID(v int32) {
	o.FkiEzsigntemplatedocumentID = &v
}

// GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeID() int32 {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfoldertypeID
}

// GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.FkiEzsignfoldertypeID, true
}

// HasFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.FkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetFkiEzsignfoldertypeID gets a reference to the given int32 and assigns it to the FkiEzsignfoldertypeID field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiEzsignfoldertypeID(v int32) {
	o.FkiEzsignfoldertypeID = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSLanguageNameX returns the SLanguageNameX field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSLanguageNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSLanguageNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLanguageNameX, true
}

// SetSLanguageNameX sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSLanguageNameX(v string) {
	o.SLanguageNameX = v
}

// GetSEzsigntemplateDescription returns the SEzsigntemplateDescription field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateDescription
}

// GetSEzsigntemplateDescriptionOk returns a tuple with the SEzsigntemplateDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateDescription, true
}

// SetSEzsigntemplateDescription sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSEzsigntemplateDescription(v string) {
	o.SEzsigntemplateDescription = v
}

// GetSEzsigntemplateFilenamepattern returns the SEzsigntemplateFilenamepattern field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateFilenamepattern() string {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateFilenamepattern
}

// GetSEzsigntemplateFilenamepatternOk returns a tuple with the SEzsigntemplateFilenamepattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsigntemplateFilenamepatternOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateFilenamepattern) {
		return nil, false
	}
	return o.SEzsigntemplateFilenamepattern, true
}

// HasSEzsigntemplateFilenamepattern returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasSEzsigntemplateFilenamepattern() bool {
	if o != nil && !IsNil(o.SEzsigntemplateFilenamepattern) {
		return true
	}

	return false
}

// SetSEzsigntemplateFilenamepattern gets a reference to the given string and assigns it to the SEzsigntemplateFilenamepattern field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSEzsigntemplateFilenamepattern(v string) {
	o.SEzsigntemplateFilenamepattern = &v
}

// GetBEzsigntemplateAdminonly returns the BEzsigntemplateAdminonly field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateAdminonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateAdminonly
}

// GetBEzsigntemplateAdminonlyOk returns a tuple with the BEzsigntemplateAdminonly field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateAdminonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateAdminonly, true
}

// SetBEzsigntemplateAdminonly sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetBEzsigntemplateAdminonly(v bool) {
	o.BEzsigntemplateAdminonly = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsignfoldertypeNameX() string {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldertypeNameX) {
		return nil, false
	}
	return o.SEzsignfoldertypeNameX, true
}

// HasSEzsignfoldertypeNameX returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasSEzsignfoldertypeNameX() bool {
	if o != nil && !IsNil(o.SEzsignfoldertypeNameX) {
		return true
	}

	return false
}

// SetSEzsignfoldertypeNameX gets a reference to the given string and assigns it to the SEzsignfoldertypeNameX field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = &v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

// GetBEzsigntemplateEditallowed returns the BEzsigntemplateEditallowed field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateEditallowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplateEditallowed
}

// GetBEzsigntemplateEditallowedOk returns a tuple with the BEzsigntemplateEditallowed field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetBEzsigntemplateEditallowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplateEditallowed, true
}

// SetBEzsigntemplateEditallowed sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetBEzsigntemplateEditallowed(v bool) {
	o.BEzsigntemplateEditallowed = v
}

// GetEEzsigntemplateType returns the EEzsigntemplateType field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetEEzsigntemplateType() FieldEEzsigntemplateType {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		var ret FieldEEzsigntemplateType
		return ret
	}
	return *o.EEzsigntemplateType
}

// GetEEzsigntemplateTypeOk returns a tuple with the EEzsigntemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetEEzsigntemplateTypeOk() (*FieldEEzsigntemplateType, bool) {
	if o == nil || IsNil(o.EEzsigntemplateType) {
		return nil, false
	}
	return o.EEzsigntemplateType, true
}

// HasEEzsigntemplateType returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasEEzsigntemplateType() bool {
	if o != nil && !IsNil(o.EEzsigntemplateType) {
		return true
	}

	return false
}

// SetEEzsigntemplateType gets a reference to the given FieldEEzsigntemplateType and assigns it to the EEzsigntemplateType field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetEEzsigntemplateType(v FieldEEzsigntemplateType) {
	o.EEzsigntemplateType = &v
}

// GetObjEzsigntemplatedocument returns the ObjEzsigntemplatedocument field value if set, zero value otherwise.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjEzsigntemplatedocument() EzsigntemplatedocumentResponse {
	if o == nil || IsNil(o.ObjEzsigntemplatedocument) {
		var ret EzsigntemplatedocumentResponse
		return ret
	}
	return *o.ObjEzsigntemplatedocument
}

// GetObjEzsigntemplatedocumentOk returns a tuple with the ObjEzsigntemplatedocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetObjEzsigntemplatedocumentOk() (*EzsigntemplatedocumentResponse, bool) {
	if o == nil || IsNil(o.ObjEzsigntemplatedocument) {
		return nil, false
	}
	return o.ObjEzsigntemplatedocument, true
}

// HasObjEzsigntemplatedocument returns a boolean if a field has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) HasObjEzsigntemplatedocument() bool {
	if o != nil && !IsNil(o.ObjEzsigntemplatedocument) {
		return true
	}

	return false
}

// SetObjEzsigntemplatedocument gets a reference to the given EzsigntemplatedocumentResponse and assigns it to the ObjEzsigntemplatedocument field.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetObjEzsigntemplatedocument(v EzsigntemplatedocumentResponse) {
	o.ObjEzsigntemplatedocument = &v
}

// GetAObjEzsigntemplatesigner returns the AObjEzsigntemplatesigner field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetAObjEzsigntemplatesigner() []EzsigntemplatesignerResponseCompound {
	if o == nil {
		var ret []EzsigntemplatesignerResponseCompound
		return ret
	}

	return o.AObjEzsigntemplatesigner
}

// GetAObjEzsigntemplatesignerOk returns a tuple with the AObjEzsigntemplatesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetObjectV1ResponseMPayload) GetAObjEzsigntemplatesignerOk() ([]EzsigntemplatesignerResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatesigner, true
}

// SetAObjEzsigntemplatesigner sets field value
func (o *EzsigntemplateGetObjectV1ResponseMPayload) SetAObjEzsigntemplatesigner(v []EzsigntemplatesignerResponseCompound) {
	o.AObjEzsigntemplatesigner = v
}

func (o EzsigntemplateGetObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateGetObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ObjEzsigntemplatedocument) {
		toSerialize["objEzsigntemplatedocument"] = o.ObjEzsigntemplatedocument
	}
	toSerialize["a_objEzsigntemplatesigner"] = o.AObjEzsigntemplatesigner
	return toSerialize, nil
}

func (o *EzsigntemplateGetObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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
		"a_objEzsigntemplatesigner",
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

	varEzsigntemplateGetObjectV1ResponseMPayload := _EzsigntemplateGetObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateGetObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplateGetObjectV1ResponseMPayload(varEzsigntemplateGetObjectV1ResponseMPayload)

	return err
}

type NullableEzsigntemplateGetObjectV1ResponseMPayload struct {
	value *EzsigntemplateGetObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateGetObjectV1ResponseMPayload) Get() *EzsigntemplateGetObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateGetObjectV1ResponseMPayload) Set(val *EzsigntemplateGetObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateGetObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateGetObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateGetObjectV1ResponseMPayload(val *EzsigntemplateGetObjectV1ResponseMPayload) *NullableEzsigntemplateGetObjectV1ResponseMPayload {
	return &NullableEzsigntemplateGetObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateGetObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateGetObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


