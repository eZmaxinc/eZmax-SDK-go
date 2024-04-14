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

// checks if the EzsignsignergroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignergroupResponse{}

// EzsignsignergroupResponse An Ezsignsignergroup Object
type EzsignsignergroupResponse struct {
	// The unique ID of the Ezsignsignergroup
	PkiEzsignsignergroupID int32 `json:"pkiEzsignsignergroupID"`
	ObjEzsignsignergroupDescription MultilingualEzsignsignergroupDescription `json:"objEzsignsignergroupDescription"`
	// The Description of the Ezsignsignergroup in the language of the requester
	SEzsignsignergroupDescriptionX *string `json:"sEzsignsignergroupDescriptionX,omitempty"`
}

type _EzsignsignergroupResponse EzsignsignergroupResponse

// NewEzsignsignergroupResponse instantiates a new EzsignsignergroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignergroupResponse(pkiEzsignsignergroupID int32, objEzsignsignergroupDescription MultilingualEzsignsignergroupDescription) *EzsignsignergroupResponse {
	this := EzsignsignergroupResponse{}
	this.PkiEzsignsignergroupID = pkiEzsignsignergroupID
	this.ObjEzsignsignergroupDescription = objEzsignsignergroupDescription
	return &this
}

// NewEzsignsignergroupResponseWithDefaults instantiates a new EzsignsignergroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignergroupResponseWithDefaults() *EzsignsignergroupResponse {
	this := EzsignsignergroupResponse{}
	return &this
}

// GetPkiEzsignsignergroupID returns the PkiEzsignsignergroupID field value
func (o *EzsignsignergroupResponse) GetPkiEzsignsignergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignergroupID
}

// GetPkiEzsignsignergroupIDOk returns a tuple with the PkiEzsignsignergroupID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponse) GetPkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignergroupID, true
}

// SetPkiEzsignsignergroupID sets field value
func (o *EzsignsignergroupResponse) SetPkiEzsignsignergroupID(v int32) {
	o.PkiEzsignsignergroupID = v
}

// GetObjEzsignsignergroupDescription returns the ObjEzsignsignergroupDescription field value
func (o *EzsignsignergroupResponse) GetObjEzsignsignergroupDescription() MultilingualEzsignsignergroupDescription {
	if o == nil {
		var ret MultilingualEzsignsignergroupDescription
		return ret
	}

	return o.ObjEzsignsignergroupDescription
}

// GetObjEzsignsignergroupDescriptionOk returns a tuple with the ObjEzsignsignergroupDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponse) GetObjEzsignsignergroupDescriptionOk() (*MultilingualEzsignsignergroupDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignergroupDescription, true
}

// SetObjEzsignsignergroupDescription sets field value
func (o *EzsignsignergroupResponse) SetObjEzsignsignergroupDescription(v MultilingualEzsignsignergroupDescription) {
	o.ObjEzsignsignergroupDescription = v
}

// GetSEzsignsignergroupDescriptionX returns the SEzsignsignergroupDescriptionX field value if set, zero value otherwise.
func (o *EzsignsignergroupResponse) GetSEzsignsignergroupDescriptionX() string {
	if o == nil || IsNil(o.SEzsignsignergroupDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzsignsignergroupDescriptionX
}

// GetSEzsignsignergroupDescriptionXOk returns a tuple with the SEzsignsignergroupDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignergroupResponse) GetSEzsignsignergroupDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignsignergroupDescriptionX) {
		return nil, false
	}
	return o.SEzsignsignergroupDescriptionX, true
}

// HasSEzsignsignergroupDescriptionX returns a boolean if a field has been set.
func (o *EzsignsignergroupResponse) HasSEzsignsignergroupDescriptionX() bool {
	if o != nil && !IsNil(o.SEzsignsignergroupDescriptionX) {
		return true
	}

	return false
}

// SetSEzsignsignergroupDescriptionX gets a reference to the given string and assigns it to the SEzsignsignergroupDescriptionX field.
func (o *EzsignsignergroupResponse) SetSEzsignsignergroupDescriptionX(v string) {
	o.SEzsignsignergroupDescriptionX = &v
}

func (o EzsignsignergroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignergroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignergroupID"] = o.PkiEzsignsignergroupID
	toSerialize["objEzsignsignergroupDescription"] = o.ObjEzsignsignergroupDescription
	if !IsNil(o.SEzsignsignergroupDescriptionX) {
		toSerialize["sEzsignsignergroupDescriptionX"] = o.SEzsignsignergroupDescriptionX
	}
	return toSerialize, nil
}

func (o *EzsignsignergroupResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignergroupID",
		"objEzsignsignergroupDescription",
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

	varEzsignsignergroupResponse := _EzsignsignergroupResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignergroupResponse)

	if err != nil {
		return err
	}

	*o = EzsignsignergroupResponse(varEzsignsignergroupResponse)

	return err
}

type NullableEzsignsignergroupResponse struct {
	value *EzsignsignergroupResponse
	isSet bool
}

func (v NullableEzsignsignergroupResponse) Get() *EzsignsignergroupResponse {
	return v.value
}

func (v *NullableEzsignsignergroupResponse) Set(val *EzsignsignergroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignergroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignergroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignergroupResponse(val *EzsignsignergroupResponse) *NullableEzsignsignergroupResponse {
	return &NullableEzsignsignergroupResponse{value: val, isSet: true}
}

func (v NullableEzsignsignergroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignergroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


