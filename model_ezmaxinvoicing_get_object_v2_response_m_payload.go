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

// checks if the EzmaxinvoicingGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingGetObjectV2ResponseMPayload{}

// EzmaxinvoicingGetObjectV2ResponseMPayload Payload for GET /2/object/ezmaxinvoicing/{pkiEzmaxinvoicingID}
type EzmaxinvoicingGetObjectV2ResponseMPayload struct {
	ObjEzmaxinvoicing EzmaxinvoicingResponseCompound `json:"objEzmaxinvoicing"`
}

type _EzmaxinvoicingGetObjectV2ResponseMPayload EzmaxinvoicingGetObjectV2ResponseMPayload

// NewEzmaxinvoicingGetObjectV2ResponseMPayload instantiates a new EzmaxinvoicingGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingGetObjectV2ResponseMPayload(objEzmaxinvoicing EzmaxinvoicingResponseCompound) *EzmaxinvoicingGetObjectV2ResponseMPayload {
	this := EzmaxinvoicingGetObjectV2ResponseMPayload{}
	this.ObjEzmaxinvoicing = objEzmaxinvoicing
	return &this
}

// NewEzmaxinvoicingGetObjectV2ResponseMPayloadWithDefaults instantiates a new EzmaxinvoicingGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingGetObjectV2ResponseMPayloadWithDefaults() *EzmaxinvoicingGetObjectV2ResponseMPayload {
	this := EzmaxinvoicingGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjEzmaxinvoicing returns the ObjEzmaxinvoicing field value
func (o *EzmaxinvoicingGetObjectV2ResponseMPayload) GetObjEzmaxinvoicing() EzmaxinvoicingResponseCompound {
	if o == nil {
		var ret EzmaxinvoicingResponseCompound
		return ret
	}

	return o.ObjEzmaxinvoicing
}

// GetObjEzmaxinvoicingOk returns a tuple with the ObjEzmaxinvoicing field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingGetObjectV2ResponseMPayload) GetObjEzmaxinvoicingOk() (*EzmaxinvoicingResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzmaxinvoicing, true
}

// SetObjEzmaxinvoicing sets field value
func (o *EzmaxinvoicingGetObjectV2ResponseMPayload) SetObjEzmaxinvoicing(v EzmaxinvoicingResponseCompound) {
	o.ObjEzmaxinvoicing = v
}

func (o EzmaxinvoicingGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzmaxinvoicing"] = o.ObjEzmaxinvoicing
	return toSerialize, nil
}

func (o *EzmaxinvoicingGetObjectV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objEzmaxinvoicing",
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

	varEzmaxinvoicingGetObjectV2ResponseMPayload := _EzmaxinvoicingGetObjectV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingGetObjectV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingGetObjectV2ResponseMPayload(varEzmaxinvoicingGetObjectV2ResponseMPayload)

	return err
}

type NullableEzmaxinvoicingGetObjectV2ResponseMPayload struct {
	value *EzmaxinvoicingGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzmaxinvoicingGetObjectV2ResponseMPayload) Get() *EzmaxinvoicingGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzmaxinvoicingGetObjectV2ResponseMPayload) Set(val *EzmaxinvoicingGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingGetObjectV2ResponseMPayload(val *EzmaxinvoicingGetObjectV2ResponseMPayload) *NullableEzmaxinvoicingGetObjectV2ResponseMPayload {
	return &NullableEzmaxinvoicingGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


