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

// checks if the EmailRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailRequestCompound{}

// EmailRequestCompound An Email Object and children to create a complete structure
type EmailRequestCompound struct {
	// The unique ID of the Email
	PkiEmailID *int32 `json:"pkiEmailID,omitempty"`
	// The unique ID of the Emailtype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home|
	FkiEmailtypeID int32 `json:"fkiEmailtypeID"`
	// The email address.
	SEmailAddress string "json:\"sEmailAddress\" validate:\"regexp=^[\\\\w.%+\\\\-!#$%&'*+\\/=?^`{|}~]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,20}$\""
}

type _EmailRequestCompound EmailRequestCompound

// NewEmailRequestCompound instantiates a new EmailRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailRequestCompound(fkiEmailtypeID int32, sEmailAddress string) *EmailRequestCompound {
	this := EmailRequestCompound{}
	this.FkiEmailtypeID = fkiEmailtypeID
	this.SEmailAddress = sEmailAddress
	return &this
}

// NewEmailRequestCompoundWithDefaults instantiates a new EmailRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailRequestCompoundWithDefaults() *EmailRequestCompound {
	this := EmailRequestCompound{}
	return &this
}

// GetPkiEmailID returns the PkiEmailID field value if set, zero value otherwise.
func (o *EmailRequestCompound) GetPkiEmailID() int32 {
	if o == nil || IsNil(o.PkiEmailID) {
		var ret int32
		return ret
	}
	return *o.PkiEmailID
}

// GetPkiEmailIDOk returns a tuple with the PkiEmailID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailRequestCompound) GetPkiEmailIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEmailID) {
		return nil, false
	}
	return o.PkiEmailID, true
}

// HasPkiEmailID returns a boolean if a field has been set.
func (o *EmailRequestCompound) HasPkiEmailID() bool {
	if o != nil && !IsNil(o.PkiEmailID) {
		return true
	}

	return false
}

// SetPkiEmailID gets a reference to the given int32 and assigns it to the PkiEmailID field.
func (o *EmailRequestCompound) SetPkiEmailID(v int32) {
	o.PkiEmailID = &v
}

// GetFkiEmailtypeID returns the FkiEmailtypeID field value
func (o *EmailRequestCompound) GetFkiEmailtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEmailtypeID
}

// GetFkiEmailtypeIDOk returns a tuple with the FkiEmailtypeID field value
// and a boolean to check if the value has been set.
func (o *EmailRequestCompound) GetFkiEmailtypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEmailtypeID, true
}

// SetFkiEmailtypeID sets field value
func (o *EmailRequestCompound) SetFkiEmailtypeID(v int32) {
	o.FkiEmailtypeID = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *EmailRequestCompound) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *EmailRequestCompound) GetSEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *EmailRequestCompound) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

func (o EmailRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEmailID) {
		toSerialize["pkiEmailID"] = o.PkiEmailID
	}
	toSerialize["fkiEmailtypeID"] = o.FkiEmailtypeID
	toSerialize["sEmailAddress"] = o.SEmailAddress
	return toSerialize, nil
}

func (o *EmailRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEmailtypeID",
		"sEmailAddress",
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

	varEmailRequestCompound := _EmailRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailRequestCompound)

	if err != nil {
		return err
	}

	*o = EmailRequestCompound(varEmailRequestCompound)

	return err
}

type NullableEmailRequestCompound struct {
	value *EmailRequestCompound
	isSet bool
}

func (v NullableEmailRequestCompound) Get() *EmailRequestCompound {
	return v.value
}

func (v *NullableEmailRequestCompound) Set(val *EmailRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailRequestCompound(val *EmailRequestCompound) *NullableEmailRequestCompound {
	return &NullableEmailRequestCompound{value: val, isSet: true}
}

func (v NullableEmailRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


