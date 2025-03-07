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

// checks if the EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload{}

// EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload Payload for POST /1/object/ezsigntemplatepackagemembership
type EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigntemplatepackagemembershipID []int32 `json:"a_pkiEzsigntemplatepackagemembershipID"`
	// Whether the Ezsignbulksend was automatically modified and needs a manual validation
	BEzsigntemplatepackageNeedvalidation bool `json:"bEzsigntemplatepackageNeedvalidation"`
	// Whether the Ezsigntemplatepackage was automatically modified and needs a manual validation
	BEzsignbulksendNeedvalidation bool `json:"bEzsignbulksendNeedvalidation"`
}

type _EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload

// NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload instantiates a new EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload(aPkiEzsigntemplatepackagemembershipID []int32, bEzsigntemplatepackageNeedvalidation bool, bEzsignbulksendNeedvalidation bool) *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload{}
	this.APkiEzsigntemplatepackagemembershipID = aPkiEzsigntemplatepackagemembershipID
	this.BEzsigntemplatepackageNeedvalidation = bEzsigntemplatepackageNeedvalidation
	this.BEzsignbulksendNeedvalidation = bEzsignbulksendNeedvalidation
	return &this
}

// NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayloadWithDefaults() *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload {
	this := EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplatepackagemembershipID returns the APkiEzsigntemplatepackagemembershipID field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagemembershipID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplatepackagemembershipID
}

// GetAPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the APkiEzsigntemplatepackagemembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetAPkiEzsigntemplatepackagemembershipIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplatepackagemembershipID, true
}

// SetAPkiEzsigntemplatepackagemembershipID sets field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetAPkiEzsigntemplatepackagemembershipID(v []int32) {
	o.APkiEzsigntemplatepackagemembershipID = v
}

// GetBEzsigntemplatepackageNeedvalidation returns the BEzsigntemplatepackageNeedvalidation field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsigntemplatepackageNeedvalidation
}

// GetBEzsigntemplatepackageNeedvalidationOk returns a tuple with the BEzsigntemplatepackageNeedvalidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsigntemplatepackageNeedvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsigntemplatepackageNeedvalidation, true
}

// SetBEzsigntemplatepackageNeedvalidation sets field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetBEzsigntemplatepackageNeedvalidation(v bool) {
	o.BEzsigntemplatepackageNeedvalidation = v
}

// GetBEzsignbulksendNeedvalidation returns the BEzsignbulksendNeedvalidation field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignbulksendNeedvalidation
}

// GetBEzsignbulksendNeedvalidationOk returns a tuple with the BEzsignbulksendNeedvalidation field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) GetBEzsignbulksendNeedvalidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignbulksendNeedvalidation, true
}

// SetBEzsignbulksendNeedvalidation sets field value
func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) SetBEzsignbulksendNeedvalidation(v bool) {
	o.BEzsignbulksendNeedvalidation = v
}

func (o EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplatepackagemembershipID"] = o.APkiEzsigntemplatepackagemembershipID
	toSerialize["bEzsigntemplatepackageNeedvalidation"] = o.BEzsigntemplatepackageNeedvalidation
	toSerialize["bEzsignbulksendNeedvalidation"] = o.BEzsignbulksendNeedvalidation
	return toSerialize, nil
}

func (o *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsigntemplatepackagemembershipID",
		"bEzsigntemplatepackageNeedvalidation",
		"bEzsignbulksendNeedvalidation",
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

	varEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload := _EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload(varEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload)

	return err
}

type NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload struct {
	value *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) Get() *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) Set(val *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload(val *EzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) *NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload {
	return &NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


