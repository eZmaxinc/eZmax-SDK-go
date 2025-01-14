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

// checks if the EzsignuserRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignuserRequestCompound{}

// EzsignuserRequestCompound A Ezsignuser Object and children
type EzsignuserRequestCompound struct {
	// The unique ID of the Ezsignuser
	PkiEzsignuserID *int32 `json:"pkiEzsignuserID,omitempty"`
	// The unique ID of the Contact
	FkiContactID int32 `json:"fkiContactID"`
	ObjContact ContactRequestCompoundV2 `json:"objContact"`
}

type _EzsignuserRequestCompound EzsignuserRequestCompound

// NewEzsignuserRequestCompound instantiates a new EzsignuserRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignuserRequestCompound(fkiContactID int32, objContact ContactRequestCompoundV2) *EzsignuserRequestCompound {
	this := EzsignuserRequestCompound{}
	this.FkiContactID = fkiContactID
	this.ObjContact = objContact
	return &this
}

// NewEzsignuserRequestCompoundWithDefaults instantiates a new EzsignuserRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignuserRequestCompoundWithDefaults() *EzsignuserRequestCompound {
	this := EzsignuserRequestCompound{}
	return &this
}

// GetPkiEzsignuserID returns the PkiEzsignuserID field value if set, zero value otherwise.
func (o *EzsignuserRequestCompound) GetPkiEzsignuserID() int32 {
	if o == nil || IsNil(o.PkiEzsignuserID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignuserID
}

// GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignuserRequestCompound) GetPkiEzsignuserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignuserID) {
		return nil, false
	}
	return o.PkiEzsignuserID, true
}

// HasPkiEzsignuserID returns a boolean if a field has been set.
func (o *EzsignuserRequestCompound) HasPkiEzsignuserID() bool {
	if o != nil && !IsNil(o.PkiEzsignuserID) {
		return true
	}

	return false
}

// SetPkiEzsignuserID gets a reference to the given int32 and assigns it to the PkiEzsignuserID field.
func (o *EzsignuserRequestCompound) SetPkiEzsignuserID(v int32) {
	o.PkiEzsignuserID = &v
}

// GetFkiContactID returns the FkiContactID field value
func (o *EzsignuserRequestCompound) GetFkiContactID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiContactID
}

// GetFkiContactIDOk returns a tuple with the FkiContactID field value
// and a boolean to check if the value has been set.
func (o *EzsignuserRequestCompound) GetFkiContactIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiContactID, true
}

// SetFkiContactID sets field value
func (o *EzsignuserRequestCompound) SetFkiContactID(v int32) {
	o.FkiContactID = v
}

// GetObjContact returns the ObjContact field value
func (o *EzsignuserRequestCompound) GetObjContact() ContactRequestCompoundV2 {
	if o == nil {
		var ret ContactRequestCompoundV2
		return ret
	}

	return o.ObjContact
}

// GetObjContactOk returns a tuple with the ObjContact field value
// and a boolean to check if the value has been set.
func (o *EzsignuserRequestCompound) GetObjContactOk() (*ContactRequestCompoundV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjContact, true
}

// SetObjContact sets field value
func (o *EzsignuserRequestCompound) SetObjContact(v ContactRequestCompoundV2) {
	o.ObjContact = v
}

func (o EzsignuserRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignuserRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignuserID) {
		toSerialize["pkiEzsignuserID"] = o.PkiEzsignuserID
	}
	toSerialize["fkiContactID"] = o.FkiContactID
	toSerialize["objContact"] = o.ObjContact
	return toSerialize, nil
}

func (o *EzsignuserRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiContactID",
		"objContact",
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

	varEzsignuserRequestCompound := _EzsignuserRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignuserRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsignuserRequestCompound(varEzsignuserRequestCompound)

	return err
}

type NullableEzsignuserRequestCompound struct {
	value *EzsignuserRequestCompound
	isSet bool
}

func (v NullableEzsignuserRequestCompound) Get() *EzsignuserRequestCompound {
	return v.value
}

func (v *NullableEzsignuserRequestCompound) Set(val *EzsignuserRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignuserRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignuserRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignuserRequestCompound(val *EzsignuserRequestCompound) *NullableEzsignuserRequestCompound {
	return &NullableEzsignuserRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignuserRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignuserRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


