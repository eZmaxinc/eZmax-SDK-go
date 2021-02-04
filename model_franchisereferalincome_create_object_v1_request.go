/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.28
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// FranchisereferalincomeCreateObjectV1Request Request for the /1/object/franchisereferalincome/createObject API Request
type FranchisereferalincomeCreateObjectV1Request struct {
	ObjFranchisereferalincome *FranchisereferalincomeRequest `json:"objFranchisereferalincome,omitempty"`
	ObjFranchisereferalincomeCompound *FranchisereferalincomeRequestCompound `json:"objFranchisereferalincomeCompound,omitempty"`
}

// NewFranchisereferalincomeCreateObjectV1Request instantiates a new FranchisereferalincomeCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV1Request() *FranchisereferalincomeCreateObjectV1Request {
	this := FranchisereferalincomeCreateObjectV1Request{}
	return &this
}

// NewFranchisereferalincomeCreateObjectV1RequestWithDefaults instantiates a new FranchisereferalincomeCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV1RequestWithDefaults() *FranchisereferalincomeCreateObjectV1Request {
	this := FranchisereferalincomeCreateObjectV1Request{}
	return &this
}

// GetObjFranchisereferalincome returns the ObjFranchisereferalincome field value if set, zero value otherwise.
func (o *FranchisereferalincomeCreateObjectV1Request) GetObjFranchisereferalincome() FranchisereferalincomeRequest {
	if o == nil || o.ObjFranchisereferalincome == nil {
		var ret FranchisereferalincomeRequest
		return ret
	}
	return *o.ObjFranchisereferalincome
}

// GetObjFranchisereferalincomeOk returns a tuple with the ObjFranchisereferalincome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1Request) GetObjFranchisereferalincomeOk() (*FranchisereferalincomeRequest, bool) {
	if o == nil || o.ObjFranchisereferalincome == nil {
		return nil, false
	}
	return o.ObjFranchisereferalincome, true
}

// HasObjFranchisereferalincome returns a boolean if a field has been set.
func (o *FranchisereferalincomeCreateObjectV1Request) HasObjFranchisereferalincome() bool {
	if o != nil && o.ObjFranchisereferalincome != nil {
		return true
	}

	return false
}

// SetObjFranchisereferalincome gets a reference to the given FranchisereferalincomeRequest and assigns it to the ObjFranchisereferalincome field.
func (o *FranchisereferalincomeCreateObjectV1Request) SetObjFranchisereferalincome(v FranchisereferalincomeRequest) {
	o.ObjFranchisereferalincome = &v
}

// GetObjFranchisereferalincomeCompound returns the ObjFranchisereferalincomeCompound field value if set, zero value otherwise.
func (o *FranchisereferalincomeCreateObjectV1Request) GetObjFranchisereferalincomeCompound() FranchisereferalincomeRequestCompound {
	if o == nil || o.ObjFranchisereferalincomeCompound == nil {
		var ret FranchisereferalincomeRequestCompound
		return ret
	}
	return *o.ObjFranchisereferalincomeCompound
}

// GetObjFranchisereferalincomeCompoundOk returns a tuple with the ObjFranchisereferalincomeCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1Request) GetObjFranchisereferalincomeCompoundOk() (*FranchisereferalincomeRequestCompound, bool) {
	if o == nil || o.ObjFranchisereferalincomeCompound == nil {
		return nil, false
	}
	return o.ObjFranchisereferalincomeCompound, true
}

// HasObjFranchisereferalincomeCompound returns a boolean if a field has been set.
func (o *FranchisereferalincomeCreateObjectV1Request) HasObjFranchisereferalincomeCompound() bool {
	if o != nil && o.ObjFranchisereferalincomeCompound != nil {
		return true
	}

	return false
}

// SetObjFranchisereferalincomeCompound gets a reference to the given FranchisereferalincomeRequestCompound and assigns it to the ObjFranchisereferalincomeCompound field.
func (o *FranchisereferalincomeCreateObjectV1Request) SetObjFranchisereferalincomeCompound(v FranchisereferalincomeRequestCompound) {
	o.ObjFranchisereferalincomeCompound = &v
}

func (o FranchisereferalincomeCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjFranchisereferalincome != nil {
		toSerialize["objFranchisereferalincome"] = o.ObjFranchisereferalincome
	}
	if o.ObjFranchisereferalincomeCompound != nil {
		toSerialize["objFranchisereferalincomeCompound"] = o.ObjFranchisereferalincomeCompound
	}
	return json.Marshal(toSerialize)
}

type NullableFranchisereferalincomeCreateObjectV1Request struct {
	value *FranchisereferalincomeCreateObjectV1Request
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV1Request) Get() *FranchisereferalincomeCreateObjectV1Request {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV1Request) Set(val *FranchisereferalincomeCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV1Request(val *FranchisereferalincomeCreateObjectV1Request) *NullableFranchisereferalincomeCreateObjectV1Request {
	return &NullableFranchisereferalincomeCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

