/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsignbulksendsignermappingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendsignermappingResponse{}

// EzsignbulksendsignermappingResponse A Ezsignbulksendsignermapping Object
type EzsignbulksendsignermappingResponse struct {
	// The unique ID of the Ezsignbulksendsignermapping
	PkiEzsignbulksendsignermappingID int32 `json:"pkiEzsignbulksendsignermappingID"`
	// The unique ID of the Ezsignbulksend
	FkiEzsignbulksendID int32 `json:"fkiEzsignbulksendID"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The description of the Ezsignbulksendsignermapping
	SEzsignbulksendsignermappingDescription string `json:"sEzsignbulksendsignermappingDescription"`
}

type _EzsignbulksendsignermappingResponse EzsignbulksendsignermappingResponse

// NewEzsignbulksendsignermappingResponse instantiates a new EzsignbulksendsignermappingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendsignermappingResponse(pkiEzsignbulksendsignermappingID int32, fkiEzsignbulksendID int32, sEzsignbulksendsignermappingDescription string) *EzsignbulksendsignermappingResponse {
	this := EzsignbulksendsignermappingResponse{}
	this.PkiEzsignbulksendsignermappingID = pkiEzsignbulksendsignermappingID
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	this.SEzsignbulksendsignermappingDescription = sEzsignbulksendsignermappingDescription
	return &this
}

// NewEzsignbulksendsignermappingResponseWithDefaults instantiates a new EzsignbulksendsignermappingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendsignermappingResponseWithDefaults() *EzsignbulksendsignermappingResponse {
	this := EzsignbulksendsignermappingResponse{}
	return &this
}

// GetPkiEzsignbulksendsignermappingID returns the PkiEzsignbulksendsignermappingID field value
func (o *EzsignbulksendsignermappingResponse) GetPkiEzsignbulksendsignermappingID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignbulksendsignermappingID
}

// GetPkiEzsignbulksendsignermappingIDOk returns a tuple with the PkiEzsignbulksendsignermappingID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponse) GetPkiEzsignbulksendsignermappingIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignbulksendsignermappingID, true
}

// SetPkiEzsignbulksendsignermappingID sets field value
func (o *EzsignbulksendsignermappingResponse) SetPkiEzsignbulksendsignermappingID(v int32) {
	o.PkiEzsignbulksendsignermappingID = v
}

// GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field value
func (o *EzsignbulksendsignermappingResponse) GetFkiEzsignbulksendID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignbulksendID
}

// GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponse) GetFkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignbulksendID, true
}

// SetFkiEzsignbulksendID sets field value
func (o *EzsignbulksendsignermappingResponse) SetFkiEzsignbulksendID(v int32) {
	o.FkiEzsignbulksendID = v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignbulksendsignermappingResponse) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignbulksendsignermappingResponse) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignbulksendsignermappingResponse) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetSEzsignbulksendsignermappingDescription returns the SEzsignbulksendsignermappingDescription field value
func (o *EzsignbulksendsignermappingResponse) GetSEzsignbulksendsignermappingDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignbulksendsignermappingDescription
}

// GetSEzsignbulksendsignermappingDescriptionOk returns a tuple with the SEzsignbulksendsignermappingDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendsignermappingResponse) GetSEzsignbulksendsignermappingDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignbulksendsignermappingDescription, true
}

// SetSEzsignbulksendsignermappingDescription sets field value
func (o *EzsignbulksendsignermappingResponse) SetSEzsignbulksendsignermappingDescription(v string) {
	o.SEzsignbulksendsignermappingDescription = v
}

func (o EzsignbulksendsignermappingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendsignermappingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignbulksendsignermappingID"] = o.PkiEzsignbulksendsignermappingID
	toSerialize["fkiEzsignbulksendID"] = o.FkiEzsignbulksendID
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	toSerialize["sEzsignbulksendsignermappingDescription"] = o.SEzsignbulksendsignermappingDescription
	return toSerialize, nil
}

func (o *EzsignbulksendsignermappingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignbulksendsignermappingID",
		"fkiEzsignbulksendID",
		"sEzsignbulksendsignermappingDescription",
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

	varEzsignbulksendsignermappingResponse := _EzsignbulksendsignermappingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendsignermappingResponse)

	if err != nil {
		return err
	}

	*o = EzsignbulksendsignermappingResponse(varEzsignbulksendsignermappingResponse)

	return err
}

type NullableEzsignbulksendsignermappingResponse struct {
	value *EzsignbulksendsignermappingResponse
	isSet bool
}

func (v NullableEzsignbulksendsignermappingResponse) Get() *EzsignbulksendsignermappingResponse {
	return v.value
}

func (v *NullableEzsignbulksendsignermappingResponse) Set(val *EzsignbulksendsignermappingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendsignermappingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendsignermappingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendsignermappingResponse(val *EzsignbulksendsignermappingResponse) *NullableEzsignbulksendsignermappingResponse {
	return &NullableEzsignbulksendsignermappingResponse{value: val, isSet: true}
}

func (v NullableEzsignbulksendsignermappingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendsignermappingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


