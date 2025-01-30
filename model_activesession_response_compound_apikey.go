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

// checks if the ActivesessionResponseCompoundApikey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivesessionResponseCompoundApikey{}

// ActivesessionResponseCompoundApikey An Activesession->Apikey object and children to create a complete structure
type ActivesessionResponseCompoundApikey struct {
	// The unique ID of the Apikey
	PkiApikeyID int32 `json:"pkiApikeyID"`
	// The description of the Apikey in the language of the requester
	SApikeyDescriptionX string `json:"sApikeyDescriptionX"`
}

type _ActivesessionResponseCompoundApikey ActivesessionResponseCompoundApikey

// NewActivesessionResponseCompoundApikey instantiates a new ActivesessionResponseCompoundApikey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionResponseCompoundApikey(pkiApikeyID int32, sApikeyDescriptionX string) *ActivesessionResponseCompoundApikey {
	this := ActivesessionResponseCompoundApikey{}
	this.PkiApikeyID = pkiApikeyID
	this.SApikeyDescriptionX = sApikeyDescriptionX
	return &this
}

// NewActivesessionResponseCompoundApikeyWithDefaults instantiates a new ActivesessionResponseCompoundApikey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionResponseCompoundApikeyWithDefaults() *ActivesessionResponseCompoundApikey {
	this := ActivesessionResponseCompoundApikey{}
	return &this
}

// GetPkiApikeyID returns the PkiApikeyID field value
func (o *ActivesessionResponseCompoundApikey) GetPkiApikeyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiApikeyID
}

// GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundApikey) GetPkiApikeyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiApikeyID, true
}

// SetPkiApikeyID sets field value
func (o *ActivesessionResponseCompoundApikey) SetPkiApikeyID(v int32) {
	o.PkiApikeyID = v
}

// GetSApikeyDescriptionX returns the SApikeyDescriptionX field value
func (o *ActivesessionResponseCompoundApikey) GetSApikeyDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SApikeyDescriptionX
}

// GetSApikeyDescriptionXOk returns a tuple with the SApikeyDescriptionX field value
// and a boolean to check if the value has been set.
func (o *ActivesessionResponseCompoundApikey) GetSApikeyDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SApikeyDescriptionX, true
}

// SetSApikeyDescriptionX sets field value
func (o *ActivesessionResponseCompoundApikey) SetSApikeyDescriptionX(v string) {
	o.SApikeyDescriptionX = v
}

func (o ActivesessionResponseCompoundApikey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivesessionResponseCompoundApikey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiApikeyID"] = o.PkiApikeyID
	toSerialize["sApikeyDescriptionX"] = o.SApikeyDescriptionX
	return toSerialize, nil
}

func (o *ActivesessionResponseCompoundApikey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiApikeyID",
		"sApikeyDescriptionX",
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

	varActivesessionResponseCompoundApikey := _ActivesessionResponseCompoundApikey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivesessionResponseCompoundApikey)

	if err != nil {
		return err
	}

	*o = ActivesessionResponseCompoundApikey(varActivesessionResponseCompoundApikey)

	return err
}

type NullableActivesessionResponseCompoundApikey struct {
	value *ActivesessionResponseCompoundApikey
	isSet bool
}

func (v NullableActivesessionResponseCompoundApikey) Get() *ActivesessionResponseCompoundApikey {
	return v.value
}

func (v *NullableActivesessionResponseCompoundApikey) Set(val *ActivesessionResponseCompoundApikey) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionResponseCompoundApikey) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionResponseCompoundApikey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionResponseCompoundApikey(val *ActivesessionResponseCompoundApikey) *NullableActivesessionResponseCompoundApikey {
	return &NullableActivesessionResponseCompoundApikey{value: val, isSet: true}
}

func (v NullableActivesessionResponseCompoundApikey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionResponseCompoundApikey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


