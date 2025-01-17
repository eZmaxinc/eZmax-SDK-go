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

// checks if the EzmaxinvoicingsummaryglobalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryglobalResponseCompound{}

// EzmaxinvoicingsummaryglobalResponseCompound A Ezmaxinvoicingsummaryglobal Object
type EzmaxinvoicingsummaryglobalResponseCompound struct {
	EzmaxinvoicingsummaryglobalResponse
	AObjEzmaxinvoicingcommission []EzmaxinvoicingcommissionResponseCompound `json:"a_objEzmaxinvoicingcommission,omitempty"`
}

type _EzmaxinvoicingsummaryglobalResponseCompound EzmaxinvoicingsummaryglobalResponseCompound

// NewEzmaxinvoicingsummaryglobalResponseCompound instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryglobalResponseCompound(fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, dtEzmaxinvoicingsummaryglobalStart string, dtEzmaxinvoicingsummaryglobalEnd string, iEzmaxinvoicingsummaryglobalDays int32, dEzmaxinvoicingsummaryglobalCountreal string, dEzmaxinvoicingsummaryglobalCountbilled string, dEzmaxinvoicingsummaryglobalSubtotal string, dEzmaxinvoicingsummaryglobalRebateamount string, dEzmaxinvoicingsummaryglobalRebatepercent string, dEzmaxinvoicingsummaryglobalRebatetotal string, dEzmaxinvoicingsummaryglobalTotal string, bEzmaxinvoicingsummaryglobalAdjustment bool, tEzmaxproductHelpX string) *EzmaxinvoicingsummaryglobalResponseCompound {
	this := EzmaxinvoicingsummaryglobalResponseCompound{}
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.SEzmaxproductDescriptionX = sEzmaxproductDescriptionX
	this.DtEzmaxinvoicingsummaryglobalStart = dtEzmaxinvoicingsummaryglobalStart
	this.DtEzmaxinvoicingsummaryglobalEnd = dtEzmaxinvoicingsummaryglobalEnd
	this.IEzmaxinvoicingsummaryglobalDays = iEzmaxinvoicingsummaryglobalDays
	this.DEzmaxinvoicingsummaryglobalCountreal = dEzmaxinvoicingsummaryglobalCountreal
	this.DEzmaxinvoicingsummaryglobalCountbilled = dEzmaxinvoicingsummaryglobalCountbilled
	this.DEzmaxinvoicingsummaryglobalSubtotal = dEzmaxinvoicingsummaryglobalSubtotal
	this.DEzmaxinvoicingsummaryglobalRebateamount = dEzmaxinvoicingsummaryglobalRebateamount
	this.DEzmaxinvoicingsummaryglobalRebatepercent = dEzmaxinvoicingsummaryglobalRebatepercent
	this.DEzmaxinvoicingsummaryglobalRebatetotal = dEzmaxinvoicingsummaryglobalRebatetotal
	this.DEzmaxinvoicingsummaryglobalTotal = dEzmaxinvoicingsummaryglobalTotal
	this.BEzmaxinvoicingsummaryglobalAdjustment = bEzmaxinvoicingsummaryglobalAdjustment
	this.TEzmaxproductHelpX = tEzmaxproductHelpX
	return &this
}

// NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryglobalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryglobalResponseCompoundWithDefaults() *EzmaxinvoicingsummaryglobalResponseCompound {
	this := EzmaxinvoicingsummaryglobalResponseCompound{}
	return &this
}

// GetAObjEzmaxinvoicingcommission returns the AObjEzmaxinvoicingcommission field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommission() []EzmaxinvoicingcommissionResponseCompound {
	if o == nil || IsNil(o.AObjEzmaxinvoicingcommission) {
		var ret []EzmaxinvoicingcommissionResponseCompound
		return ret
	}
	return o.AObjEzmaxinvoicingcommission
}

// GetAObjEzmaxinvoicingcommissionOk returns a tuple with the AObjEzmaxinvoicingcommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) GetAObjEzmaxinvoicingcommissionOk() ([]EzmaxinvoicingcommissionResponseCompound, bool) {
	if o == nil || IsNil(o.AObjEzmaxinvoicingcommission) {
		return nil, false
	}
	return o.AObjEzmaxinvoicingcommission, true
}

// HasAObjEzmaxinvoicingcommission returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) HasAObjEzmaxinvoicingcommission() bool {
	if o != nil && !IsNil(o.AObjEzmaxinvoicingcommission) {
		return true
	}

	return false
}

// SetAObjEzmaxinvoicingcommission gets a reference to the given []EzmaxinvoicingcommissionResponseCompound and assigns it to the AObjEzmaxinvoicingcommission field.
func (o *EzmaxinvoicingsummaryglobalResponseCompound) SetAObjEzmaxinvoicingcommission(v []EzmaxinvoicingcommissionResponseCompound) {
	o.AObjEzmaxinvoicingcommission = v
}

func (o EzmaxinvoicingsummaryglobalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryglobalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AObjEzmaxinvoicingcommission) {
		toSerialize["a_objEzmaxinvoicingcommission"] = o.AObjEzmaxinvoicingcommission
	}
	return toSerialize, nil
}

func (o *EzmaxinvoicingsummaryglobalResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxproductID",
		"sEzmaxproductDescriptionX",
		"dtEzmaxinvoicingsummaryglobalStart",
		"dtEzmaxinvoicingsummaryglobalEnd",
		"iEzmaxinvoicingsummaryglobalDays",
		"dEzmaxinvoicingsummaryglobalCountreal",
		"dEzmaxinvoicingsummaryglobalCountbilled",
		"dEzmaxinvoicingsummaryglobalSubtotal",
		"dEzmaxinvoicingsummaryglobalRebateamount",
		"dEzmaxinvoicingsummaryglobalRebatepercent",
		"dEzmaxinvoicingsummaryglobalRebatetotal",
		"dEzmaxinvoicingsummaryglobalTotal",
		"bEzmaxinvoicingsummaryglobalAdjustment",
		"tEzmaxproductHelpX",
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

	varEzmaxinvoicingsummaryglobalResponseCompound := _EzmaxinvoicingsummaryglobalResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingsummaryglobalResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingsummaryglobalResponseCompound(varEzmaxinvoicingsummaryglobalResponseCompound)

	return err
}

type NullableEzmaxinvoicingsummaryglobalResponseCompound struct {
	value *EzmaxinvoicingsummaryglobalResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) Get() *EzmaxinvoicingsummaryglobalResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) Set(val *EzmaxinvoicingsummaryglobalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryglobalResponseCompound(val *EzmaxinvoicingsummaryglobalResponseCompound) *NullableEzmaxinvoicingsummaryglobalResponseCompound {
	return &NullableEzmaxinvoicingsummaryglobalResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryglobalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryglobalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


