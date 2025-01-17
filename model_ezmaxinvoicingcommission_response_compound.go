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

// checks if the EzmaxinvoicingcommissionResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingcommissionResponseCompound{}

// EzmaxinvoicingcommissionResponseCompound A Ezmaxinvoicingcommission Object
type EzmaxinvoicingcommissionResponseCompound struct {
	EzmaxinvoicingcommissionResponse
	ObjContactName *CustomContactNameResponse `json:"objContactName,omitempty"`
}

type _EzmaxinvoicingcommissionResponseCompound EzmaxinvoicingcommissionResponseCompound

// NewEzmaxinvoicingcommissionResponseCompound instantiates a new EzmaxinvoicingcommissionResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingcommissionResponseCompound(dtEzmaxinvoicingcommissionStart string, dtEzmaxinvoicingcommissionEnd string, iEzmaxinvoicingcommissionDays int32, dEzmaxinvoicingcommissionAmount string) *EzmaxinvoicingcommissionResponseCompound {
	this := EzmaxinvoicingcommissionResponseCompound{}
	this.DtEzmaxinvoicingcommissionStart = dtEzmaxinvoicingcommissionStart
	this.DtEzmaxinvoicingcommissionEnd = dtEzmaxinvoicingcommissionEnd
	this.IEzmaxinvoicingcommissionDays = iEzmaxinvoicingcommissionDays
	this.DEzmaxinvoicingcommissionAmount = dEzmaxinvoicingcommissionAmount
	return &this
}

// NewEzmaxinvoicingcommissionResponseCompoundWithDefaults instantiates a new EzmaxinvoicingcommissionResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingcommissionResponseCompoundWithDefaults() *EzmaxinvoicingcommissionResponseCompound {
	this := EzmaxinvoicingcommissionResponseCompound{}
	return &this
}

// GetObjContactName returns the ObjContactName field value if set, zero value otherwise.
func (o *EzmaxinvoicingcommissionResponseCompound) GetObjContactName() CustomContactNameResponse {
	if o == nil || IsNil(o.ObjContactName) {
		var ret CustomContactNameResponse
		return ret
	}
	return *o.ObjContactName
}

// GetObjContactNameOk returns a tuple with the ObjContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcommissionResponseCompound) GetObjContactNameOk() (*CustomContactNameResponse, bool) {
	if o == nil || IsNil(o.ObjContactName) {
		return nil, false
	}
	return o.ObjContactName, true
}

// HasObjContactName returns a boolean if a field has been set.
func (o *EzmaxinvoicingcommissionResponseCompound) HasObjContactName() bool {
	if o != nil && !IsNil(o.ObjContactName) {
		return true
	}

	return false
}

// SetObjContactName gets a reference to the given CustomContactNameResponse and assigns it to the ObjContactName field.
func (o *EzmaxinvoicingcommissionResponseCompound) SetObjContactName(v CustomContactNameResponse) {
	o.ObjContactName = &v
}

func (o EzmaxinvoicingcommissionResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingcommissionResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjContactName) {
		toSerialize["objContactName"] = o.ObjContactName
	}
	return toSerialize, nil
}

func (o *EzmaxinvoicingcommissionResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dtEzmaxinvoicingcommissionStart",
		"dtEzmaxinvoicingcommissionEnd",
		"iEzmaxinvoicingcommissionDays",
		"dEzmaxinvoicingcommissionAmount",
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

	varEzmaxinvoicingcommissionResponseCompound := _EzmaxinvoicingcommissionResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingcommissionResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingcommissionResponseCompound(varEzmaxinvoicingcommissionResponseCompound)

	return err
}

type NullableEzmaxinvoicingcommissionResponseCompound struct {
	value *EzmaxinvoicingcommissionResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingcommissionResponseCompound) Get() *EzmaxinvoicingcommissionResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingcommissionResponseCompound) Set(val *EzmaxinvoicingcommissionResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingcommissionResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingcommissionResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingcommissionResponseCompound(val *EzmaxinvoicingcommissionResponseCompound) *NullableEzmaxinvoicingcommissionResponseCompound {
	return &NullableEzmaxinvoicingcommissionResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingcommissionResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingcommissionResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


