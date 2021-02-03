/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfolderEditObjectV1Request Request for the /1/object/ezsignfolder/editObject API Request
type EzsignfolderEditObjectV1Request struct {
	ObjEzsignfolder *EzsignfolderRequest `json:"objEzsignfolder,omitempty"`
}

// NewEzsignfolderEditObjectV1Request instantiates a new EzsignfolderEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderEditObjectV1Request() *EzsignfolderEditObjectV1Request {
	this := EzsignfolderEditObjectV1Request{}
	return &this
}

// NewEzsignfolderEditObjectV1RequestWithDefaults instantiates a new EzsignfolderEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderEditObjectV1RequestWithDefaults() *EzsignfolderEditObjectV1Request {
	this := EzsignfolderEditObjectV1Request{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value if set, zero value otherwise.
func (o *EzsignfolderEditObjectV1Request) GetObjEzsignfolder() EzsignfolderRequest {
	if o == nil || o.ObjEzsignfolder == nil {
		var ret EzsignfolderRequest
		return ret
	}
	return *o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderEditObjectV1Request) GetObjEzsignfolderOk() (*EzsignfolderRequest, bool) {
	if o == nil || o.ObjEzsignfolder == nil {
		return nil, false
	}
	return o.ObjEzsignfolder, true
}

// HasObjEzsignfolder returns a boolean if a field has been set.
func (o *EzsignfolderEditObjectV1Request) HasObjEzsignfolder() bool {
	if o != nil && o.ObjEzsignfolder != nil {
		return true
	}

	return false
}

// SetObjEzsignfolder gets a reference to the given EzsignfolderRequest and assigns it to the ObjEzsignfolder field.
func (o *EzsignfolderEditObjectV1Request) SetObjEzsignfolder(v EzsignfolderRequest) {
	o.ObjEzsignfolder = &v
}

func (o EzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignfolder != nil {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderEditObjectV1Request struct {
	value *EzsignfolderEditObjectV1Request
	isSet bool
}

func (v NullableEzsignfolderEditObjectV1Request) Get() *EzsignfolderEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignfolderEditObjectV1Request) Set(val *EzsignfolderEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderEditObjectV1Request(val *EzsignfolderEditObjectV1Request) *NullableEzsignfolderEditObjectV1Request {
	return &NullableEzsignfolderEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


