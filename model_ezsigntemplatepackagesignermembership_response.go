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

// checks if the EzsigntemplatepackagesignermembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignermembershipResponse{}

// EzsigntemplatepackagesignermembershipResponse A Ezsigntemplatepackagesignermembership Object
type EzsigntemplatepackagesignermembershipResponse struct {
	// The unique ID of the Ezsigntemplatepackagesignermembership
	PkiEzsigntemplatepackagesignermembershipID int32 `json:"pkiEzsigntemplatepackagesignermembershipID"`
	// The unique ID of the Ezsigntemplatepackagemembership
	FkiEzsigntemplatepackagemembershipID int32 `json:"fkiEzsigntemplatepackagemembershipID"`
	// The unique ID of the Ezsigntemplatepackagesigner
	FkiEzsigntemplatepackagesignerID int32 `json:"fkiEzsigntemplatepackagesignerID"`
	// The unique ID of the Ezsigntemplatesigner
	FkiEzsigntemplatesignerID int32 `json:"fkiEzsigntemplatesignerID"`
	// The Copy number in case of multiple copies.
	IEzsigntemplatepackagesignermembershipCopy *int32 `json:"iEzsigntemplatepackagesignermembershipCopy,omitempty"`
}

type _EzsigntemplatepackagesignermembershipResponse EzsigntemplatepackagesignermembershipResponse

// NewEzsigntemplatepackagesignermembershipResponse instantiates a new EzsigntemplatepackagesignermembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignermembershipResponse(pkiEzsigntemplatepackagesignermembershipID int32, fkiEzsigntemplatepackagemembershipID int32, fkiEzsigntemplatepackagesignerID int32, fkiEzsigntemplatesignerID int32) *EzsigntemplatepackagesignermembershipResponse {
	this := EzsigntemplatepackagesignermembershipResponse{}
	this.PkiEzsigntemplatepackagesignermembershipID = pkiEzsigntemplatepackagesignermembershipID
	this.FkiEzsigntemplatepackagemembershipID = fkiEzsigntemplatepackagemembershipID
	this.FkiEzsigntemplatepackagesignerID = fkiEzsigntemplatepackagesignerID
	this.FkiEzsigntemplatesignerID = fkiEzsigntemplatesignerID
	return &this
}

// NewEzsigntemplatepackagesignermembershipResponseWithDefaults instantiates a new EzsigntemplatepackagesignermembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignermembershipResponseWithDefaults() *EzsigntemplatepackagesignermembershipResponse {
	this := EzsigntemplatepackagesignermembershipResponse{}
	return &this
}

// GetPkiEzsigntemplatepackagesignermembershipID returns the PkiEzsigntemplatepackagesignermembershipID field value
func (o *EzsigntemplatepackagesignermembershipResponse) GetPkiEzsigntemplatepackagesignermembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackagesignermembershipID
}

// GetPkiEzsigntemplatepackagesignermembershipIDOk returns a tuple with the PkiEzsigntemplatepackagesignermembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) GetPkiEzsigntemplatepackagesignermembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackagesignermembershipID, true
}

// SetPkiEzsigntemplatepackagesignermembershipID sets field value
func (o *EzsigntemplatepackagesignermembershipResponse) SetPkiEzsigntemplatepackagesignermembershipID(v int32) {
	o.PkiEzsigntemplatepackagesignermembershipID = v
}

// GetFkiEzsigntemplatepackagemembershipID returns the FkiEzsigntemplatepackagemembershipID field value
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatepackagemembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackagemembershipID
}

// GetFkiEzsigntemplatepackagemembershipIDOk returns a tuple with the FkiEzsigntemplatepackagemembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatepackagemembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackagemembershipID, true
}

// SetFkiEzsigntemplatepackagemembershipID sets field value
func (o *EzsigntemplatepackagesignermembershipResponse) SetFkiEzsigntemplatepackagemembershipID(v int32) {
	o.FkiEzsigntemplatepackagemembershipID = v
}

// GetFkiEzsigntemplatepackagesignerID returns the FkiEzsigntemplatepackagesignerID field value
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatepackagesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackagesignerID
}

// GetFkiEzsigntemplatepackagesignerIDOk returns a tuple with the FkiEzsigntemplatepackagesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatepackagesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackagesignerID, true
}

// SetFkiEzsigntemplatepackagesignerID sets field value
func (o *EzsigntemplatepackagesignermembershipResponse) SetFkiEzsigntemplatepackagesignerID(v int32) {
	o.FkiEzsigntemplatepackagesignerID = v
}

// GetFkiEzsigntemplatesignerID returns the FkiEzsigntemplatesignerID field value
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatesignerID
}

// GetFkiEzsigntemplatesignerIDOk returns a tuple with the FkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) GetFkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatesignerID, true
}

// SetFkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatepackagesignermembershipResponse) SetFkiEzsigntemplatesignerID(v int32) {
	o.FkiEzsigntemplatesignerID = v
}

// GetIEzsigntemplatepackagesignermembershipCopy returns the IEzsigntemplatepackagesignermembershipCopy field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignermembershipResponse) GetIEzsigntemplatepackagesignermembershipCopy() int32 {
	if o == nil || IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatepackagesignermembershipCopy
}

// GetIEzsigntemplatepackagesignermembershipCopyOk returns a tuple with the IEzsigntemplatepackagesignermembershipCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) GetIEzsigntemplatepackagesignermembershipCopyOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		return nil, false
	}
	return o.IEzsigntemplatepackagesignermembershipCopy, true
}

// HasIEzsigntemplatepackagesignermembershipCopy returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignermembershipResponse) HasIEzsigntemplatepackagesignermembershipCopy() bool {
	if o != nil && !IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		return true
	}

	return false
}

// SetIEzsigntemplatepackagesignermembershipCopy gets a reference to the given int32 and assigns it to the IEzsigntemplatepackagesignermembershipCopy field.
func (o *EzsigntemplatepackagesignermembershipResponse) SetIEzsigntemplatepackagesignermembershipCopy(v int32) {
	o.IEzsigntemplatepackagesignermembershipCopy = &v
}

func (o EzsigntemplatepackagesignermembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignermembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatepackagesignermembershipID"] = o.PkiEzsigntemplatepackagesignermembershipID
	toSerialize["fkiEzsigntemplatepackagemembershipID"] = o.FkiEzsigntemplatepackagemembershipID
	toSerialize["fkiEzsigntemplatepackagesignerID"] = o.FkiEzsigntemplatepackagesignerID
	toSerialize["fkiEzsigntemplatesignerID"] = o.FkiEzsigntemplatesignerID
	if !IsNil(o.IEzsigntemplatepackagesignermembershipCopy) {
		toSerialize["iEzsigntemplatepackagesignermembershipCopy"] = o.IEzsigntemplatepackagesignermembershipCopy
	}
	return toSerialize, nil
}

func (o *EzsigntemplatepackagesignermembershipResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatepackagesignermembershipID",
		"fkiEzsigntemplatepackagemembershipID",
		"fkiEzsigntemplatepackagesignerID",
		"fkiEzsigntemplatesignerID",
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

	varEzsigntemplatepackagesignermembershipResponse := _EzsigntemplatepackagesignermembershipResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagesignermembershipResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagesignermembershipResponse(varEzsigntemplatepackagesignermembershipResponse)

	return err
}

type NullableEzsigntemplatepackagesignermembershipResponse struct {
	value *EzsigntemplatepackagesignermembershipResponse
	isSet bool
}

func (v NullableEzsigntemplatepackagesignermembershipResponse) Get() *EzsigntemplatepackagesignermembershipResponse {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignermembershipResponse) Set(val *EzsigntemplatepackagesignermembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignermembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignermembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignermembershipResponse(val *EzsigntemplatepackagesignermembershipResponse) *NullableEzsigntemplatepackagesignermembershipResponse {
	return &NullableEzsigntemplatepackagesignermembershipResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignermembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignermembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


