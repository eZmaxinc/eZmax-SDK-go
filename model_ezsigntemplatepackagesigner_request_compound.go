/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsigntemplatepackagesignerRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignerRequestCompound{}

// EzsigntemplatepackagesignerRequestCompound A Ezsigntemplatepackagesigner Object and children
type EzsigntemplatepackagesignerRequestCompound struct {
	// The unique ID of the Ezsigntemplatepackagesigner
	PkiEzsigntemplatepackagesignerID *int32 `json:"pkiEzsigntemplatepackagesignerID,omitempty"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The unique ID of the Ezdoctemplatedocument
	FkiEzdoctemplatedocumentID *int32 `json:"fkiEzdoctemplatedocumentID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
	// If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain't required to sign the document.
	BEzsigntemplatepackagesignerReceivecopy *bool `json:"bEzsigntemplatepackagesignerReceivecopy,omitempty"`
	EEzsigntemplatepackagesignerMapping *FieldEEzsigntemplatepackagesignerMapping `json:"eEzsigntemplatepackagesignerMapping,omitempty"`
	// The description of the Ezsigntemplatepackagesigner
	SEzsigntemplatepackagesignerDescription string `json:"sEzsigntemplatepackagesignerDescription"`
}

type _EzsigntemplatepackagesignerRequestCompound EzsigntemplatepackagesignerRequestCompound

// NewEzsigntemplatepackagesignerRequestCompound instantiates a new EzsigntemplatepackagesignerRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignerRequestCompound(fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string) *EzsigntemplatepackagesignerRequestCompound {
	this := EzsigntemplatepackagesignerRequestCompound{}
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	var eEzsigntemplatepackagesignerMapping FieldEEzsigntemplatepackagesignerMapping = MANUAL
	this.EEzsigntemplatepackagesignerMapping = &eEzsigntemplatepackagesignerMapping
	this.SEzsigntemplatepackagesignerDescription = sEzsigntemplatepackagesignerDescription
	return &this
}

// NewEzsigntemplatepackagesignerRequestCompoundWithDefaults instantiates a new EzsigntemplatepackagesignerRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignerRequestCompoundWithDefaults() *EzsigntemplatepackagesignerRequestCompound {
	this := EzsigntemplatepackagesignerRequestCompound{}
	var eEzsigntemplatepackagesignerMapping FieldEEzsigntemplatepackagesignerMapping = MANUAL
	this.EEzsigntemplatepackagesignerMapping = &eEzsigntemplatepackagesignerMapping
	return &this
}

// GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetPkiEzsigntemplatepackagesignerID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackagesignerID
}

// GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignerID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackagesignerID, true
}

// HasPkiEzsigntemplatepackagesignerID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasPkiEzsigntemplatepackagesignerID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackagesignerID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackagesignerID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackagesignerID field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetPkiEzsigntemplatepackagesignerID(v int32) {
	o.PkiEzsigntemplatepackagesignerID = &v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagesignerRequestCompound) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiEzdoctemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzdoctemplatedocumentID
}

// GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzdoctemplatedocumentID, true
}

// HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasFkiEzdoctemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzdoctemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzdoctemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzdoctemplatedocumentID field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetFkiEzdoctemplatedocumentID(v int32) {
	o.FkiEzdoctemplatedocumentID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetBEzsigntemplatepackagesignerReceivecopy returns the BEzsigntemplatepackagesignerReceivecopy field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetBEzsigntemplatepackagesignerReceivecopy() bool {
	if o == nil || IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatepackagesignerReceivecopy
}

// GetBEzsigntemplatepackagesignerReceivecopyOk returns a tuple with the BEzsigntemplatepackagesignerReceivecopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetBEzsigntemplatepackagesignerReceivecopyOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		return nil, false
	}
	return o.BEzsigntemplatepackagesignerReceivecopy, true
}

// HasBEzsigntemplatepackagesignerReceivecopy returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasBEzsigntemplatepackagesignerReceivecopy() bool {
	if o != nil && !IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		return true
	}

	return false
}

// SetBEzsigntemplatepackagesignerReceivecopy gets a reference to the given bool and assigns it to the BEzsigntemplatepackagesignerReceivecopy field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetBEzsigntemplatepackagesignerReceivecopy(v bool) {
	o.BEzsigntemplatepackagesignerReceivecopy = &v
}

// GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerRequestCompound) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping {
	if o == nil || IsNil(o.EEzsigntemplatepackagesignerMapping) {
		var ret FieldEEzsigntemplatepackagesignerMapping
		return ret
	}
	return *o.EEzsigntemplatepackagesignerMapping
}

// GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool) {
	if o == nil || IsNil(o.EEzsigntemplatepackagesignerMapping) {
		return nil, false
	}
	return o.EEzsigntemplatepackagesignerMapping, true
}

// HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) HasEEzsigntemplatepackagesignerMapping() bool {
	if o != nil && !IsNil(o.EEzsigntemplatepackagesignerMapping) {
		return true
	}

	return false
}

// SetEEzsigntemplatepackagesignerMapping gets a reference to the given FieldEEzsigntemplatepackagesignerMapping and assigns it to the EEzsigntemplatepackagesignerMapping field.
func (o *EzsigntemplatepackagesignerRequestCompound) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping) {
	o.EEzsigntemplatepackagesignerMapping = &v
}

// GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field value
func (o *EzsigntemplatepackagesignerRequestCompound) GetSEzsigntemplatepackagesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackagesignerDescription
}

// GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerRequestCompound) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackagesignerDescription, true
}

// SetSEzsigntemplatepackagesignerDescription sets field value
func (o *EzsigntemplatepackagesignerRequestCompound) SetSEzsigntemplatepackagesignerDescription(v string) {
	o.SEzsigntemplatepackagesignerDescription = v
}

func (o EzsigntemplatepackagesignerRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignerRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackagesignerID) {
		toSerialize["pkiEzsigntemplatepackagesignerID"] = o.PkiEzsigntemplatepackagesignerID
	}
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	if !IsNil(o.FkiEzdoctemplatedocumentID) {
		toSerialize["fkiEzdoctemplatedocumentID"] = o.FkiEzdoctemplatedocumentID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	if !IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		toSerialize["bEzsigntemplatepackagesignerReceivecopy"] = o.BEzsigntemplatepackagesignerReceivecopy
	}
	if !IsNil(o.EEzsigntemplatepackagesignerMapping) {
		toSerialize["eEzsigntemplatepackagesignerMapping"] = o.EEzsigntemplatepackagesignerMapping
	}
	toSerialize["sEzsigntemplatepackagesignerDescription"] = o.SEzsigntemplatepackagesignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplatepackagesignerRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatepackageID",
		"sEzsigntemplatepackagesignerDescription",
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

	varEzsigntemplatepackagesignerRequestCompound := _EzsigntemplatepackagesignerRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagesignerRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagesignerRequestCompound(varEzsigntemplatepackagesignerRequestCompound)

	return err
}

type NullableEzsigntemplatepackagesignerRequestCompound struct {
	value *EzsigntemplatepackagesignerRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatepackagesignerRequestCompound) Get() *EzsigntemplatepackagesignerRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignerRequestCompound) Set(val *EzsigntemplatepackagesignerRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignerRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignerRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignerRequestCompound(val *EzsigntemplatepackagesignerRequestCompound) *NullableEzsigntemplatepackagesignerRequestCompound {
	return &NullableEzsigntemplatepackagesignerRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignerRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignerRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


