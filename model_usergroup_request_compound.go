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

// checks if the UsergroupRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupRequestCompound{}

// UsergroupRequestCompound A Usergroup Object and children
type UsergroupRequestCompound struct {
	// The unique ID of the Usergroup
	PkiUsergroupID *int32 `json:"pkiUsergroupID,omitempty"`
	ObjEmail *EmailRequest `json:"objEmail,omitempty"`
	ObjUsergroupName MultilingualUsergroupName `json:"objUsergroupName"`
}

type _UsergroupRequestCompound UsergroupRequestCompound

// NewUsergroupRequestCompound instantiates a new UsergroupRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupRequestCompound(objUsergroupName MultilingualUsergroupName) *UsergroupRequestCompound {
	this := UsergroupRequestCompound{}
	this.ObjUsergroupName = objUsergroupName
	return &this
}

// NewUsergroupRequestCompoundWithDefaults instantiates a new UsergroupRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupRequestCompoundWithDefaults() *UsergroupRequestCompound {
	this := UsergroupRequestCompound{}
	return &this
}

// GetPkiUsergroupID returns the PkiUsergroupID field value if set, zero value otherwise.
func (o *UsergroupRequestCompound) GetPkiUsergroupID() int32 {
	if o == nil || IsNil(o.PkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.PkiUsergroupID
}

// GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupRequestCompound) GetPkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUsergroupID) {
		return nil, false
	}
	return o.PkiUsergroupID, true
}

// HasPkiUsergroupID returns a boolean if a field has been set.
func (o *UsergroupRequestCompound) HasPkiUsergroupID() bool {
	if o != nil && !IsNil(o.PkiUsergroupID) {
		return true
	}

	return false
}

// SetPkiUsergroupID gets a reference to the given int32 and assigns it to the PkiUsergroupID field.
func (o *UsergroupRequestCompound) SetPkiUsergroupID(v int32) {
	o.PkiUsergroupID = &v
}

// GetObjEmail returns the ObjEmail field value if set, zero value otherwise.
func (o *UsergroupRequestCompound) GetObjEmail() EmailRequest {
	if o == nil || IsNil(o.ObjEmail) {
		var ret EmailRequest
		return ret
	}
	return *o.ObjEmail
}

// GetObjEmailOk returns a tuple with the ObjEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupRequestCompound) GetObjEmailOk() (*EmailRequest, bool) {
	if o == nil || IsNil(o.ObjEmail) {
		return nil, false
	}
	return o.ObjEmail, true
}

// HasObjEmail returns a boolean if a field has been set.
func (o *UsergroupRequestCompound) HasObjEmail() bool {
	if o != nil && !IsNil(o.ObjEmail) {
		return true
	}

	return false
}

// SetObjEmail gets a reference to the given EmailRequest and assigns it to the ObjEmail field.
func (o *UsergroupRequestCompound) SetObjEmail(v EmailRequest) {
	o.ObjEmail = &v
}

// GetObjUsergroupName returns the ObjUsergroupName field value
func (o *UsergroupRequestCompound) GetObjUsergroupName() MultilingualUsergroupName {
	if o == nil {
		var ret MultilingualUsergroupName
		return ret
	}

	return o.ObjUsergroupName
}

// GetObjUsergroupNameOk returns a tuple with the ObjUsergroupName field value
// and a boolean to check if the value has been set.
func (o *UsergroupRequestCompound) GetObjUsergroupNameOk() (*MultilingualUsergroupName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUsergroupName, true
}

// SetObjUsergroupName sets field value
func (o *UsergroupRequestCompound) SetObjUsergroupName(v MultilingualUsergroupName) {
	o.ObjUsergroupName = v
}

func (o UsergroupRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiUsergroupID) {
		toSerialize["pkiUsergroupID"] = o.PkiUsergroupID
	}
	if !IsNil(o.ObjEmail) {
		toSerialize["objEmail"] = o.ObjEmail
	}
	toSerialize["objUsergroupName"] = o.ObjUsergroupName
	return toSerialize, nil
}

func (o *UsergroupRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objUsergroupName",
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

	varUsergroupRequestCompound := _UsergroupRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupRequestCompound)

	if err != nil {
		return err
	}

	*o = UsergroupRequestCompound(varUsergroupRequestCompound)

	return err
}

type NullableUsergroupRequestCompound struct {
	value *UsergroupRequestCompound
	isSet bool
}

func (v NullableUsergroupRequestCompound) Get() *UsergroupRequestCompound {
	return v.value
}

func (v *NullableUsergroupRequestCompound) Set(val *UsergroupRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupRequestCompound(val *UsergroupRequestCompound) *NullableUsergroupRequestCompound {
	return &NullableUsergroupRequestCompound{value: val, isSet: true}
}

func (v NullableUsergroupRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


