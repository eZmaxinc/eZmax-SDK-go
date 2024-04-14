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

// checks if the EzsignsigningreasonListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonListElement{}

// EzsignsigningreasonListElement A Ezsignsigningreason List Element
type EzsignsigningreasonListElement struct {
	// The unique ID of the Ezsignsigningreason
	PkiEzsignsigningreasonID int32 `json:"pkiEzsignsigningreasonID"`
	// The description of the Ezsignsigningreason in the language of the requester
	SEzsignsigningreasonDescriptionX string `json:"sEzsignsigningreasonDescriptionX"`
	// Whether the ezsignsigningreason is active or not
	BEzsignsigningreasonIsactive bool `json:"bEzsignsigningreasonIsactive"`
}

type _EzsignsigningreasonListElement EzsignsigningreasonListElement

// NewEzsignsigningreasonListElement instantiates a new EzsignsigningreasonListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonListElement(pkiEzsignsigningreasonID int32, sEzsignsigningreasonDescriptionX string, bEzsignsigningreasonIsactive bool) *EzsignsigningreasonListElement {
	this := EzsignsigningreasonListElement{}
	this.PkiEzsignsigningreasonID = pkiEzsignsigningreasonID
	this.SEzsignsigningreasonDescriptionX = sEzsignsigningreasonDescriptionX
	this.BEzsignsigningreasonIsactive = bEzsignsigningreasonIsactive
	return &this
}

// NewEzsignsigningreasonListElementWithDefaults instantiates a new EzsignsigningreasonListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonListElementWithDefaults() *EzsignsigningreasonListElement {
	this := EzsignsigningreasonListElement{}
	return &this
}

// GetPkiEzsignsigningreasonID returns the PkiEzsignsigningreasonID field value
func (o *EzsignsigningreasonListElement) GetPkiEzsignsigningreasonID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsigningreasonID
}

// GetPkiEzsignsigningreasonIDOk returns a tuple with the PkiEzsignsigningreasonID field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonListElement) GetPkiEzsignsigningreasonIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsigningreasonID, true
}

// SetPkiEzsignsigningreasonID sets field value
func (o *EzsignsigningreasonListElement) SetPkiEzsignsigningreasonID(v int32) {
	o.PkiEzsignsigningreasonID = v
}

// GetSEzsignsigningreasonDescriptionX returns the SEzsignsigningreasonDescriptionX field value
func (o *EzsignsigningreasonListElement) GetSEzsignsigningreasonDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsigningreasonDescriptionX
}

// GetSEzsignsigningreasonDescriptionXOk returns a tuple with the SEzsignsigningreasonDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonListElement) GetSEzsignsigningreasonDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsigningreasonDescriptionX, true
}

// SetSEzsignsigningreasonDescriptionX sets field value
func (o *EzsignsigningreasonListElement) SetSEzsignsigningreasonDescriptionX(v string) {
	o.SEzsignsigningreasonDescriptionX = v
}

// GetBEzsignsigningreasonIsactive returns the BEzsignsigningreasonIsactive field value
func (o *EzsignsigningreasonListElement) GetBEzsignsigningreasonIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignsigningreasonIsactive
}

// GetBEzsignsigningreasonIsactiveOk returns a tuple with the BEzsignsigningreasonIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonListElement) GetBEzsignsigningreasonIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignsigningreasonIsactive, true
}

// SetBEzsignsigningreasonIsactive sets field value
func (o *EzsignsigningreasonListElement) SetBEzsignsigningreasonIsactive(v bool) {
	o.BEzsignsigningreasonIsactive = v
}

func (o EzsignsigningreasonListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsigningreasonID"] = o.PkiEzsignsigningreasonID
	toSerialize["sEzsignsigningreasonDescriptionX"] = o.SEzsignsigningreasonDescriptionX
	toSerialize["bEzsignsigningreasonIsactive"] = o.BEzsignsigningreasonIsactive
	return toSerialize, nil
}

func (o *EzsignsigningreasonListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsigningreasonID",
		"sEzsignsigningreasonDescriptionX",
		"bEzsignsigningreasonIsactive",
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

	varEzsignsigningreasonListElement := _EzsignsigningreasonListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonListElement)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonListElement(varEzsignsigningreasonListElement)

	return err
}

type NullableEzsignsigningreasonListElement struct {
	value *EzsignsigningreasonListElement
	isSet bool
}

func (v NullableEzsignsigningreasonListElement) Get() *EzsignsigningreasonListElement {
	return v.value
}

func (v *NullableEzsignsigningreasonListElement) Set(val *EzsignsigningreasonListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonListElement(val *EzsignsigningreasonListElement) *NullableEzsignsigningreasonListElement {
	return &NullableEzsignsigningreasonListElement{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


