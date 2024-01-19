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

// checks if the DiscussionmessageDeleteObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmessageDeleteObjectV1Response{}

// DiscussionmessageDeleteObjectV1Response Response for DELETE /1/object/discussionmessage/{pkiDiscussionmessageID}
type DiscussionmessageDeleteObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

type _DiscussionmessageDeleteObjectV1Response DiscussionmessageDeleteObjectV1Response

// NewDiscussionmessageDeleteObjectV1Response instantiates a new DiscussionmessageDeleteObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmessageDeleteObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *DiscussionmessageDeleteObjectV1Response {
	this := DiscussionmessageDeleteObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewDiscussionmessageDeleteObjectV1ResponseWithDefaults instantiates a new DiscussionmessageDeleteObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmessageDeleteObjectV1ResponseWithDefaults() *DiscussionmessageDeleteObjectV1Response {
	this := DiscussionmessageDeleteObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *DiscussionmessageDeleteObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageDeleteObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *DiscussionmessageDeleteObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *DiscussionmessageDeleteObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageDeleteObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *DiscussionmessageDeleteObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *DiscussionmessageDeleteObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o DiscussionmessageDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmessageDeleteObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

func (o *DiscussionmessageDeleteObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
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

	varDiscussionmessageDeleteObjectV1Response := _DiscussionmessageDeleteObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmessageDeleteObjectV1Response)

	if err != nil {
		return err
	}

	*o = DiscussionmessageDeleteObjectV1Response(varDiscussionmessageDeleteObjectV1Response)

	return err
}

type NullableDiscussionmessageDeleteObjectV1Response struct {
	value *DiscussionmessageDeleteObjectV1Response
	isSet bool
}

func (v NullableDiscussionmessageDeleteObjectV1Response) Get() *DiscussionmessageDeleteObjectV1Response {
	return v.value
}

func (v *NullableDiscussionmessageDeleteObjectV1Response) Set(val *DiscussionmessageDeleteObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmessageDeleteObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmessageDeleteObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmessageDeleteObjectV1Response(val *DiscussionmessageDeleteObjectV1Response) *NullableDiscussionmessageDeleteObjectV1Response {
	return &NullableDiscussionmessageDeleteObjectV1Response{value: val, isSet: true}
}

func (v NullableDiscussionmessageDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmessageDeleteObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


