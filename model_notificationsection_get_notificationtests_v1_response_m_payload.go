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

// checks if the NotificationsectionGetNotificationtestsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsectionGetNotificationtestsV1ResponseMPayload{}

// NotificationsectionGetNotificationtestsV1ResponseMPayload Payload for GET /1/object/notificationsection/{pkiNotificationsectionID}/getNotificationtests
type NotificationsectionGetNotificationtestsV1ResponseMPayload struct {
	AObjNotificationsubsection []CustomNotificationsubsectiongetnotificationtestsResponse `json:"a_objNotificationsubsection"`
}

// NewNotificationsectionGetNotificationtestsV1ResponseMPayload instantiates a new NotificationsectionGetNotificationtestsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsectionGetNotificationtestsV1ResponseMPayload(aObjNotificationsubsection []CustomNotificationsubsectiongetnotificationtestsResponse) *NotificationsectionGetNotificationtestsV1ResponseMPayload {
	this := NotificationsectionGetNotificationtestsV1ResponseMPayload{}
	this.AObjNotificationsubsection = aObjNotificationsubsection
	return &this
}

// NewNotificationsectionGetNotificationtestsV1ResponseMPayloadWithDefaults instantiates a new NotificationsectionGetNotificationtestsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsectionGetNotificationtestsV1ResponseMPayloadWithDefaults() *NotificationsectionGetNotificationtestsV1ResponseMPayload {
	this := NotificationsectionGetNotificationtestsV1ResponseMPayload{}
	return &this
}

// GetAObjNotificationsubsection returns the AObjNotificationsubsection field value
func (o *NotificationsectionGetNotificationtestsV1ResponseMPayload) GetAObjNotificationsubsection() []CustomNotificationsubsectiongetnotificationtestsResponse {
	if o == nil {
		var ret []CustomNotificationsubsectiongetnotificationtestsResponse
		return ret
	}

	return o.AObjNotificationsubsection
}

// GetAObjNotificationsubsectionOk returns a tuple with the AObjNotificationsubsection field value
// and a boolean to check if the value has been set.
func (o *NotificationsectionGetNotificationtestsV1ResponseMPayload) GetAObjNotificationsubsectionOk() ([]CustomNotificationsubsectiongetnotificationtestsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjNotificationsubsection, true
}

// SetAObjNotificationsubsection sets field value
func (o *NotificationsectionGetNotificationtestsV1ResponseMPayload) SetAObjNotificationsubsection(v []CustomNotificationsubsectiongetnotificationtestsResponse) {
	o.AObjNotificationsubsection = v
}

func (o NotificationsectionGetNotificationtestsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsectionGetNotificationtestsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objNotificationsubsection"] = o.AObjNotificationsubsection
	return toSerialize, nil
}

type NullableNotificationsectionGetNotificationtestsV1ResponseMPayload struct {
	value *NotificationsectionGetNotificationtestsV1ResponseMPayload
	isSet bool
}

func (v NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) Get() *NotificationsectionGetNotificationtestsV1ResponseMPayload {
	return v.value
}

func (v *NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) Set(val *NotificationsectionGetNotificationtestsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsectionGetNotificationtestsV1ResponseMPayload(val *NotificationsectionGetNotificationtestsV1ResponseMPayload) *NullableNotificationsectionGetNotificationtestsV1ResponseMPayload {
	return &NullableNotificationsectionGetNotificationtestsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsectionGetNotificationtestsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


