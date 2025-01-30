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

// checks if the EzsigntemplatepackagesignermembershipRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignermembershipRequestCompound{}

// EzsigntemplatepackagesignermembershipRequestCompound A Ezsigntemplatepackagesignermembership Object and children
type EzsigntemplatepackagesignermembershipRequestCompound struct {
	// The unique ID of the Ezsigntemplatepackagesignermembership
	PkiEzsigntemplatepackagesignermembershipID *int32 `json:"pkiEzsigntemplatepackagesignermembershipID,omitempty"`
	// The unique ID of the Ezsigntemplatepackagemembership
	FkiEzsigntemplatepackagemembershipID int32 `json:"fkiEzsigntemplatepackagemembershipID"`
	// The unique ID of the Ezsigntemplatepackagesigner
	FkiEzsigntemplatepackagesignerID int32 `json:"fkiEzsigntemplatepackagesignerID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
	// The Copy number in case of multiple copies.
	IEzsigntemplatepackagesignermembershipCopy *int32 `json:"iEzsigntemplatepackagesignermembershipCopy,omitempty"`
}

type _EzsigntemplatepackagesignermembershipRequestCompound EzsigntemplatepackagesignermembershipRequestCompound

// NewEzsigntemplatepackagesignermembershipRequestCompound instantiates a new EzsigntemplatepackagesignermembershipRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignermembershipRequestCompound(fkiEzsigntemplatepackagemembershipID int32, fkiEzsigntemplatepackagesignerID int32, fkiEzsigntemplatesignerID int32) *EzsigntemplatepackagesignermembershipRequestCompound {
	this := EzsigntemplatepackagesignermembershipRequestCompound{}
	this.FkiEzsigntemplatepackagemembershipID = fkiEzsigntemplatepackagemembershipID
	this.FkiEzsigntemplatepackagesignerID = fkiEzsigntemplatepackagesignerID
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	return &this
}

// NewEzsigntemplatepackagesignermembershipRequestCompoundWithDefaults instantiates a new EzsigntemplatepackagesignermembershipRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignermembershipRequestCompoundWithDefaults() *EzsigntemplatepackagesignermembershipRequestCompound {
	this := EzsigntemplatepackagesignermembershipRequestCompound{}
	return &this
}

// GetPkiEzsigntemplatepackagesignermembershipID returns the PkiEzsigntemplatepackagesignermembershipID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetPkiEzsigntemplatepackagesignermembershipID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignermembershipID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatepackagesignermembershipID
}

// GetPkiEzsigntemplatepackagesignermembershipIDOk returns a tuple with the PkiEzsigntemplatepackagesignermembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetPkiEzsigntemplatepackagesignermembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatepackagesignermembershipID) {
		return nil, false
	}
	return o.PkiEzsigntemplatepackagesignermembershipID, true
}

// HasPkiEzsigntemplatepackagesignermembershipID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) HasPkiEzsigntemplatepackagesignermembershipID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatepackagesignermembershipID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatepackagesignermembershipID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatepackagesignermembershipID field.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) SetPkiEzsigntemplatepackagesignermembershipID(v int32) {
	o.PkiEzsigntemplatepackagesignermembershipID = &v
}

// GetFkiEzsigntemplatepackagemembershipID returns the FkiEzsigntemplatepackagemembershipID field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatepackagemembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackagemembershipID
}

// GetFkiEzsigntemplatepackagemembershipIDOk returns a tuple with the FkiEzsigntemplatepackagemembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatepackagemembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackagemembershipID, true
}

// SetFkiEzsigntemplatepackagemembershipID sets field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) SetFkiEzsigntemplatepackagemembershipID(v int32) {
	o.FkiEzsigntemplatepackagemembershipID = v
}

// GetFkiEzsigntemplatepackagesignerID returns the FkiEzsigntemplatepackagesignerID field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatepackagesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackagesignerID
}

// GetFkiEzsigntemplatepackagesignerIDOk returns a tuple with the FkiEzsigntemplatepackagesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatepackagesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackagesignerID, true
}

// SetFkiEzsigntemplatepackagesignerID sets field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) SetFkiEzsigntemplatepackagesignerID(v int32) {
	o.FkiEzsigntemplatepackagesignerID = v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatepackagesignermembershipRequestCompound) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

// GetIEzsigntemplatepackagesignermembershipCopy returns the IEzsigntemplatepackagesignermembershipCopy field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetIEzsigntemplatepackagesignermembershipCopy() int32 {
	if o == nil || IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatepackagesignermembershipCopy
}

// GetIEzsigntemplatepackagesignermembershipCopyOk returns a tuple with the IEzsigntemplatepackagesignermembershipCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) GetIEzsigntemplatepackagesignermembershipCopyOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		return nil, false
	}
	return o.IEzsigntemplatepackagesignermembershipCopy, true
}

// HasIEzsigntemplatepackagesignermembershipCopy returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) HasIEzsigntemplatepackagesignermembershipCopy() bool {
	if o != nil && !IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		return true
	}

	return false
}

// SetIEzsigntemplatepackagesignermembershipCopy gets a reference to the given int32 and assigns it to the IEzsigntemplatepackagesignermembershipCopy field.
func (o *EzsigntemplatepackagesignermembershipRequestCompound) SetIEzsigntemplatepackagesignermembershipCopy(v int32) {
	o.IEzsigntemplatepackagesignermembershipCopy = &v
}

func (o EzsigntemplatepackagesignermembershipRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignermembershipRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatepackagesignermembershipID) {
		toSerialize["pkiEzsigntemplatepackagesignermembershipID"] = o.PkiEzsigntemplatepackagesignermembershipID
	}
	toSerialize["fkiEzsigntemplatepackagemembershipID"] = o.FkiEzsigntemplatepackagemembershipID
	toSerialize["fkiEzsigntemplatepackagesignerID"] = o.FkiEzsigntemplatepackagesignerID
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	if !IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		toSerialize["iEzsigntemplatepackagesignermembershipCopy"] = o.IEzsigntemplatepackagesignermembershipCopy
	}
	return toSerialize, nil
}

func (o *EzsigntemplatepackagesignermembershipRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatepackagemembershipID",
		"fkiEzsigntemplatepackagesignerID",
		"fkiEzsigntemplatesignerID",
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

	varEzsigntemplatepackagesignermembershipRequestCompound := _EzsigntemplatepackagesignermembershipRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagesignermembershipRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagesignermembershipRequestCompound(varEzsigntemplatepackagesignermembershipRequestCompound)

	return err
}

type NullableEzsigntemplatepackagesignermembershipRequestCompound struct {
	value *EzsigntemplatepackagesignermembershipRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatepackagesignermembershipRequestCompound) Get() *EzsigntemplatepackagesignermembershipRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignermembershipRequestCompound) Set(val *EzsigntemplatepackagesignermembershipRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignermembershipRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignermembershipRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignermembershipRequestCompound(val *EzsigntemplatepackagesignermembershipRequestCompound) *NullableEzsigntemplatepackagesignermembershipRequestCompound {
	return &NullableEzsigntemplatepackagesignermembershipRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignermembershipRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignermembershipRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


