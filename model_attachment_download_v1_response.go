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

// checks if the AttachmentDownloadV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentDownloadV1Response{}

// AttachmentDownloadV1Response Response for POST /1/object/ezsignfolder/{pkiEzsignfolderID}/send
type AttachmentDownloadV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewAttachmentDownloadV1Response instantiates a new AttachmentDownloadV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentDownloadV1Response(objDebugPayload CommonResponseObjDebugPayload) *AttachmentDownloadV1Response {
	this := AttachmentDownloadV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewAttachmentDownloadV1ResponseWithDefaults instantiates a new AttachmentDownloadV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentDownloadV1ResponseWithDefaults() *AttachmentDownloadV1Response {
	this := AttachmentDownloadV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *AttachmentDownloadV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *AttachmentDownloadV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *AttachmentDownloadV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *AttachmentDownloadV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentDownloadV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *AttachmentDownloadV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *AttachmentDownloadV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o AttachmentDownloadV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentDownloadV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

type NullableAttachmentDownloadV1Response struct {
	value *AttachmentDownloadV1Response
	isSet bool
}

func (v NullableAttachmentDownloadV1Response) Get() *AttachmentDownloadV1Response {
	return v.value
}

func (v *NullableAttachmentDownloadV1Response) Set(val *AttachmentDownloadV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentDownloadV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentDownloadV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentDownloadV1Response(val *AttachmentDownloadV1Response) *NullableAttachmentDownloadV1Response {
	return &NullableAttachmentDownloadV1Response{value: val, isSet: true}
}

func (v NullableAttachmentDownloadV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentDownloadV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


