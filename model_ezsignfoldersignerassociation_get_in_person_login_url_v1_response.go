/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfoldersignerassociationGetInPersonLoginUrlV1Response Response for the /1/object/ezsignfoldersignerassociation/getInPersonLoginUrl API Request
type EzsignfoldersignerassociationGetInPersonLoginUrlV1Response struct {
	MPayload EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1Response instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1Response(mPayload EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseWithDefaults instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseWithDefaults() *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetMPayload() EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	if o == nil {
		var ret EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetMPayloadOk() (*EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) SetMPayload(v EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response struct {
	value *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response
	isSet bool
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) Get() *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) Set(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1Response) *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response {
	return &NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


