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
)

// checks if the DiscussionmessageRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmessageRequestPatch{}

// DiscussionmessageRequestPatch A Discussionmessage Object
type DiscussionmessageRequestPatch struct {
	// The unique ID of the Discussionmembership
	FkiDiscussionmembershipIDActionrequired *int32 `json:"fkiDiscussionmembershipIDActionrequired,omitempty"`
	// The content of the Discussionmessage
	TDiscussionmessageContent *string `json:"tDiscussionmessageContent,omitempty" validate:"regexp=^.{0,65535}$"`
}

// NewDiscussionmessageRequestPatch instantiates a new DiscussionmessageRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmessageRequestPatch() *DiscussionmessageRequestPatch {
	this := DiscussionmessageRequestPatch{}
	return &this
}

// NewDiscussionmessageRequestPatchWithDefaults instantiates a new DiscussionmessageRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmessageRequestPatchWithDefaults() *DiscussionmessageRequestPatch {
	this := DiscussionmessageRequestPatch{}
	return &this
}

// GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field value if set, zero value otherwise.
func (o *DiscussionmessageRequestPatch) GetFkiDiscussionmembershipIDActionrequired() int32 {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		var ret int32
		return ret
	}
	return *o.FkiDiscussionmembershipIDActionrequired
}

// GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequestPatch) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return nil, false
	}
	return o.FkiDiscussionmembershipIDActionrequired, true
}

// HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.
func (o *DiscussionmessageRequestPatch) HasFkiDiscussionmembershipIDActionrequired() bool {
	if o != nil && !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return true
	}

	return false
}

// SetFkiDiscussionmembershipIDActionrequired gets a reference to the given int32 and assigns it to the FkiDiscussionmembershipIDActionrequired field.
func (o *DiscussionmessageRequestPatch) SetFkiDiscussionmembershipIDActionrequired(v int32) {
	o.FkiDiscussionmembershipIDActionrequired = &v
}

// GetTDiscussionmessageContent returns the TDiscussionmessageContent field value if set, zero value otherwise.
func (o *DiscussionmessageRequestPatch) GetTDiscussionmessageContent() string {
	if o == nil || IsNil(o.TDiscussionmessageContent) {
		var ret string
		return ret
	}
	return *o.TDiscussionmessageContent
}

// GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequestPatch) GetTDiscussionmessageContentOk() (*string, bool) {
	if o == nil || IsNil(o.TDiscussionmessageContent) {
		return nil, false
	}
	return o.TDiscussionmessageContent, true
}

// HasTDiscussionmessageContent returns a boolean if a field has been set.
func (o *DiscussionmessageRequestPatch) HasTDiscussionmessageContent() bool {
	if o != nil && !IsNil(o.TDiscussionmessageContent) {
		return true
	}

	return false
}

// SetTDiscussionmessageContent gets a reference to the given string and assigns it to the TDiscussionmessageContent field.
func (o *DiscussionmessageRequestPatch) SetTDiscussionmessageContent(v string) {
	o.TDiscussionmessageContent = &v
}

func (o DiscussionmessageRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmessageRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		toSerialize["fkiDiscussionmembershipIDActionrequired"] = o.FkiDiscussionmembershipIDActionrequired
	}
	if !IsNil(o.TDiscussionmessageContent) {
		toSerialize["tDiscussionmessageContent"] = o.TDiscussionmessageContent
	}
	return toSerialize, nil
}

type NullableDiscussionmessageRequestPatch struct {
	value *DiscussionmessageRequestPatch
	isSet bool
}

func (v NullableDiscussionmessageRequestPatch) Get() *DiscussionmessageRequestPatch {
	return v.value
}

func (v *NullableDiscussionmessageRequestPatch) Set(val *DiscussionmessageRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmessageRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmessageRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmessageRequestPatch(val *DiscussionmessageRequestPatch) *NullableDiscussionmessageRequestPatch {
	return &NullableDiscussionmessageRequestPatch{value: val, isSet: true}
}

func (v NullableDiscussionmessageRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmessageRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


