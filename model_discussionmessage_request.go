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

// checks if the DiscussionmessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmessageRequest{}

// DiscussionmessageRequest A Discussionmessage Object
type DiscussionmessageRequest struct {
	// The unique ID of the Discussionmessage
	PkiDiscussionmessageID *int32 `json:"pkiDiscussionmessageID,omitempty"`
	// The unique ID of the Discussion
	FkiDiscussionID int32 `json:"fkiDiscussionID"`
	// The unique ID of the Discussionmembership
	FkiDiscussionmembershipIDActionrequired *int32 `json:"fkiDiscussionmembershipIDActionrequired,omitempty"`
	// The content of the Discussionmessage
	TDiscussionmessageContent string `json:"tDiscussionmessageContent"`
}

type _DiscussionmessageRequest DiscussionmessageRequest

// NewDiscussionmessageRequest instantiates a new DiscussionmessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmessageRequest(fkiDiscussionID int32, tDiscussionmessageContent string) *DiscussionmessageRequest {
	this := DiscussionmessageRequest{}
	this.FkiDiscussionID = fkiDiscussionID
	this.TDiscussionmessageContent = tDiscussionmessageContent
	return &this
}

// NewDiscussionmessageRequestWithDefaults instantiates a new DiscussionmessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmessageRequestWithDefaults() *DiscussionmessageRequest {
	this := DiscussionmessageRequest{}
	return &this
}

// GetPkiDiscussionmessageID returns the PkiDiscussionmessageID field value if set, zero value otherwise.
func (o *DiscussionmessageRequest) GetPkiDiscussionmessageID() int32 {
	if o == nil || IsNil(o.PkiDiscussionmessageID) {
		var ret int32
		return ret
	}
	return *o.PkiDiscussionmessageID
}

// GetPkiDiscussionmessageIDOk returns a tuple with the PkiDiscussionmessageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequest) GetPkiDiscussionmessageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiDiscussionmessageID) {
		return nil, false
	}
	return o.PkiDiscussionmessageID, true
}

// HasPkiDiscussionmessageID returns a boolean if a field has been set.
func (o *DiscussionmessageRequest) HasPkiDiscussionmessageID() bool {
	if o != nil && !IsNil(o.PkiDiscussionmessageID) {
		return true
	}

	return false
}

// SetPkiDiscussionmessageID gets a reference to the given int32 and assigns it to the PkiDiscussionmessageID field.
func (o *DiscussionmessageRequest) SetPkiDiscussionmessageID(v int32) {
	o.PkiDiscussionmessageID = &v
}

// GetFkiDiscussionID returns the FkiDiscussionID field value
func (o *DiscussionmessageRequest) GetFkiDiscussionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDiscussionID
}

// GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequest) GetFkiDiscussionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDiscussionID, true
}

// SetFkiDiscussionID sets field value
func (o *DiscussionmessageRequest) SetFkiDiscussionID(v int32) {
	o.FkiDiscussionID = v
}

// GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field value if set, zero value otherwise.
func (o *DiscussionmessageRequest) GetFkiDiscussionmembershipIDActionrequired() int32 {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		var ret int32
		return ret
	}
	return *o.FkiDiscussionmembershipIDActionrequired
}

// GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequest) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return nil, false
	}
	return o.FkiDiscussionmembershipIDActionrequired, true
}

// HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.
func (o *DiscussionmessageRequest) HasFkiDiscussionmembershipIDActionrequired() bool {
	if o != nil && !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return true
	}

	return false
}

// SetFkiDiscussionmembershipIDActionrequired gets a reference to the given int32 and assigns it to the FkiDiscussionmembershipIDActionrequired field.
func (o *DiscussionmessageRequest) SetFkiDiscussionmembershipIDActionrequired(v int32) {
	o.FkiDiscussionmembershipIDActionrequired = &v
}

// GetTDiscussionmessageContent returns the TDiscussionmessageContent field value
func (o *DiscussionmessageRequest) GetTDiscussionmessageContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TDiscussionmessageContent
}

// GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageRequest) GetTDiscussionmessageContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TDiscussionmessageContent, true
}

// SetTDiscussionmessageContent sets field value
func (o *DiscussionmessageRequest) SetTDiscussionmessageContent(v string) {
	o.TDiscussionmessageContent = v
}

func (o DiscussionmessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiDiscussionmessageID) {
		toSerialize["pkiDiscussionmessageID"] = o.PkiDiscussionmessageID
	}
	toSerialize["fkiDiscussionID"] = o.FkiDiscussionID
	if !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		toSerialize["fkiDiscussionmembershipIDActionrequired"] = o.FkiDiscussionmembershipIDActionrequired
	}
	toSerialize["tDiscussionmessageContent"] = o.TDiscussionmessageContent
	return toSerialize, nil
}

func (o *DiscussionmessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiDiscussionID",
		"tDiscussionmessageContent",
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

	varDiscussionmessageRequest := _DiscussionmessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmessageRequest)

	if err != nil {
		return err
	}

	*o = DiscussionmessageRequest(varDiscussionmessageRequest)

	return err
}

type NullableDiscussionmessageRequest struct {
	value *DiscussionmessageRequest
	isSet bool
}

func (v NullableDiscussionmessageRequest) Get() *DiscussionmessageRequest {
	return v.value
}

func (v *NullableDiscussionmessageRequest) Set(val *DiscussionmessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmessageRequest(val *DiscussionmessageRequest) *NullableDiscussionmessageRequest {
	return &NullableDiscussionmessageRequest{value: val, isSet: true}
}

func (v NullableDiscussionmessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

