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

// checks if the EzsignbulksenddocumentmappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksenddocumentmappingRequest{}

// EzsignbulksenddocumentmappingRequest A Ezsignbulksenddocumentmapping Object
type EzsignbulksenddocumentmappingRequest struct {
	// The unique ID of the Ezsignbulksenddocumentmapping.
	PkiEzsignbulksenddocumentmappingID *int32 `json:"pkiEzsignbulksenddocumentmappingID,omitempty"`
	// The unique ID of the Ezsignbulksend
	FkiEzsignbulksendID int32 `json:"fkiEzsignbulksendID"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID *int32 `json:"fkiEzsigntemplatepackageID,omitempty"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID *int32 `json:"fkiEzsigntemplateID,omitempty"`
}

type _EzsignbulksenddocumentmappingRequest EzsignbulksenddocumentmappingRequest

// NewEzsignbulksenddocumentmappingRequest instantiates a new EzsignbulksenddocumentmappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksenddocumentmappingRequest(fkiEzsignbulksendID int32) *EzsignbulksenddocumentmappingRequest {
	this := EzsignbulksenddocumentmappingRequest{}
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	return &this
}

// NewEzsignbulksenddocumentmappingRequestWithDefaults instantiates a new EzsignbulksenddocumentmappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksenddocumentmappingRequestWithDefaults() *EzsignbulksenddocumentmappingRequest {
	this := EzsignbulksenddocumentmappingRequest{}
	return &this
}

// GetPkiEzsignbulksenddocumentmappingID returns the PkiEzsignbulksenddocumentmappingID field value if set, zero value otherwise.
func (o *EzsignbulksenddocumentmappingRequest) GetPkiEzsignbulksenddocumentmappingID() int32 {
	if o == nil || IsNil(o.PkiEzsignbulksenddocumentmappingID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignbulksenddocumentmappingID
}

// GetPkiEzsignbulksenddocumentmappingIDOk returns a tuple with the PkiEzsignbulksenddocumentmappingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingRequest) GetPkiEzsignbulksenddocumentmappingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignbulksenddocumentmappingID) {
		return nil, false
	}
	return o.PkiEzsignbulksenddocumentmappingID, true
}

// HasPkiEzsignbulksenddocumentmappingID returns a boolean if a field has been set.
func (o *EzsignbulksenddocumentmappingRequest) HasPkiEzsignbulksenddocumentmappingID() bool {
	if o != nil && !IsNil(o.PkiEzsignbulksenddocumentmappingID) {
		return true
	}

	return false
}

// SetPkiEzsignbulksenddocumentmappingID gets a reference to the given int32 and assigns it to the PkiEzsignbulksenddocumentmappingID field.
func (o *EzsignbulksenddocumentmappingRequest) SetPkiEzsignbulksenddocumentmappingID(v int32) {
	o.PkiEzsignbulksenddocumentmappingID = &v
}

// GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field value
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsignbulksendID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignbulksendID
}

// GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignbulksendID, true
}

// SetFkiEzsignbulksendID sets field value
func (o *EzsignbulksenddocumentmappingRequest) SetFkiEzsignbulksendID(v int32) {
	o.FkiEzsignbulksendID = v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value if set, zero value otherwise.
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplatepackageID) {
		return nil, false
	}
	return o.FkiEzsigntemplatepackageID, true
}

// HasFkiEzsigntemplatepackageID returns a boolean if a field has been set.
func (o *EzsignbulksenddocumentmappingRequest) HasFkiEzsigntemplatepackageID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplatepackageID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplatepackageID gets a reference to the given int32 and assigns it to the FkiEzsigntemplatepackageID field.
func (o *EzsignbulksenddocumentmappingRequest) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = &v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value if set, zero value otherwise.
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsigntemplateID() int32 {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksenddocumentmappingRequest) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigntemplateID) {
		return nil, false
	}
	return o.FkiEzsigntemplateID, true
}

// HasFkiEzsigntemplateID returns a boolean if a field has been set.
func (o *EzsignbulksenddocumentmappingRequest) HasFkiEzsigntemplateID() bool {
	if o != nil && !IsNil(o.FkiEzsigntemplateID) {
		return true
	}

	return false
}

// SetFkiEzsigntemplateID gets a reference to the given int32 and assigns it to the FkiEzsigntemplateID field.
func (o *EzsignbulksenddocumentmappingRequest) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = &v
}

func (o EzsignbulksenddocumentmappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksenddocumentmappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignbulksenddocumentmappingID) {
		toSerialize["pkiEzsignbulksenddocumentmappingID"] = o.PkiEzsignbulksenddocumentmappingID
	}
	toSerialize["fkiEzsignbulksendID"] = o.FkiEzsignbulksendID
	if !IsNil(o.FkiEzsigntemplatepackageID) {
		toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	}
	if !IsNil(o.FkiEzsigntemplateID) {
		toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	}
	return toSerialize, nil
}

func (o *EzsignbulksenddocumentmappingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignbulksendID",
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

	varEzsignbulksenddocumentmappingRequest := _EzsignbulksenddocumentmappingRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksenddocumentmappingRequest)

	if err != nil {
		return err
	}

	*o = EzsignbulksenddocumentmappingRequest(varEzsignbulksenddocumentmappingRequest)

	return err
}

type NullableEzsignbulksenddocumentmappingRequest struct {
	value *EzsignbulksenddocumentmappingRequest
	isSet bool
}

func (v NullableEzsignbulksenddocumentmappingRequest) Get() *EzsignbulksenddocumentmappingRequest {
	return v.value
}

func (v *NullableEzsignbulksenddocumentmappingRequest) Set(val *EzsignbulksenddocumentmappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksenddocumentmappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksenddocumentmappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksenddocumentmappingRequest(val *EzsignbulksenddocumentmappingRequest) *NullableEzsignbulksenddocumentmappingRequest {
	return &NullableEzsignbulksenddocumentmappingRequest{value: val, isSet: true}
}

func (v NullableEzsignbulksenddocumentmappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksenddocumentmappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


