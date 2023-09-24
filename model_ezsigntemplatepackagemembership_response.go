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
)

// checks if the EzsigntemplatepackagemembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipResponse{}

// EzsigntemplatepackagemembershipResponse A Ezsigntemplatepackagemembership Object
type EzsigntemplatepackagemembershipResponse struct {
	// The unique ID of the Ezsigntemplatepackagemembership
	PkiEzsigntemplatepackagemembershipID int32 `json:"pkiEzsigntemplatepackagemembershipID"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The order in which the Ezsigntemplate will be imported when using an Ezsigntemplatepackage.
	IEzsigntemplatepackagemembershipOrder int32 `json:"iEzsigntemplatepackagemembershipOrder"`
}

// NewEzsigntemplatepackagemembershipResponse instantiates a new EzsigntemplatepackagemembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipResponse(pkiEzsigntemplatepackagemembershipID int32, fkiEzsigntemplatepackageID int32, fkiEzsigntemplateID int32, iEzsigntemplatepackagemembershipOrder int32) *EzsigntemplatepackagemembershipResponse {
	this := EzsigntemplatepackagemembershipResponse{}
	this.PkiEzsigntemplatepackagemembershipID = pkiEzsigntemplatepackagemembershipID
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.IEzsigntemplatepackagemembershipOrder = iEzsigntemplatepackagemembershipOrder
	return &this
}

// NewEzsigntemplatepackagemembershipResponseWithDefaults instantiates a new EzsigntemplatepackagemembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipResponseWithDefaults() *EzsigntemplatepackagemembershipResponse {
	this := EzsigntemplatepackagemembershipResponse{}
	return &this
}

// GetPkiEzsigntemplatepackagemembershipID returns the PkiEzsigntemplatepackagemembershipID field value
func (o *EzsigntemplatepackagemembershipResponse) GetPkiEzsigntemplatepackagemembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackagemembershipID
}

// GetPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the PkiEzsigntemplatepackagemembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponse) GetPkiEzsigntemplatepackagemembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackagemembershipID, true
}

// SetPkiEzsigntemplatepackagemembershipID sets field value
func (o *EzsigntemplatepackagemembershipResponse) SetPkiEzsigntemplatepackagemembershipID(v int32) {
	o.PkiEzsigntemplatepackagemembershipID = v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagemembershipResponse) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponse) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagemembershipResponse) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatepackagemembershipResponse) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponse) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatepackagemembershipResponse) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetIEzsigntemplatepackagemembershipOrder returns the IEzsigntemplatepackagemembershipOrder field value
func (o *EzsigntemplatepackagemembershipResponse) GetIEzsigntemplatepackagemembershipOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatepackagemembershipOrder
}

// GetIEzsigntemplatepackagemembershipOrderOk returns a tuple with the IEzsigntemplatepackagemembershipOrder field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponse) GetIEzsigntemplatepackagemembershipOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatepackagemembershipOrder, true
}

// SetIEzsigntemplatepackagemembershipOrder sets field value
func (o *EzsigntemplatepackagemembershipResponse) SetIEzsigntemplatepackagemembershipOrder(v int32) {
	o.IEzsigntemplatepackagemembershipOrder = v
}

func (o EzsigntemplatepackagemembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatepackagemembershipID"] = o.PkiEzsigntemplatepackagemembershipID
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["iEzsigntemplatepackagemembershipOrder"] = o.IEzsigntemplatepackagemembershipOrder
	return toSerialize, nil
}

type NullableEzsigntemplatepackagemembershipResponse struct {
	value *EzsigntemplatepackagemembershipResponse
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipResponse) Get() *EzsigntemplatepackagemembershipResponse {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipResponse) Set(val *EzsigntemplatepackagemembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipResponse(val *EzsigntemplatepackagemembershipResponse) *NullableEzsigntemplatepackagemembershipResponse {
	return &NullableEzsigntemplatepackagemembershipResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

