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

// checks if the CustomCommunicationrecipientsgroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCommunicationrecipientsgroupResponse{}

// CustomCommunicationrecipientsgroupResponse Generic CommunicationrecipientsGroup Response
type CustomCommunicationrecipientsgroupResponse struct {
	// The label for the Communicationrecipientsgroup
	SCommunicationrecipientsgroupLabel string `json:"sCommunicationrecipientsgroupLabel"`
	AObjCommunicationrecipientsrecipient []CustomCommunicationrecipientsrecipientResponse `json:"a_objCommunicationrecipientsrecipient"`
}

// NewCustomCommunicationrecipientsgroupResponse instantiates a new CustomCommunicationrecipientsgroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCommunicationrecipientsgroupResponse(sCommunicationrecipientsgroupLabel string, aObjCommunicationrecipientsrecipient []CustomCommunicationrecipientsrecipientResponse) *CustomCommunicationrecipientsgroupResponse {
	this := CustomCommunicationrecipientsgroupResponse{}
	this.SCommunicationrecipientsgroupLabel = sCommunicationrecipientsgroupLabel
	this.AObjCommunicationrecipientsrecipient = aObjCommunicationrecipientsrecipient
	return &this
}

// NewCustomCommunicationrecipientsgroupResponseWithDefaults instantiates a new CustomCommunicationrecipientsgroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCommunicationrecipientsgroupResponseWithDefaults() *CustomCommunicationrecipientsgroupResponse {
	this := CustomCommunicationrecipientsgroupResponse{}
	return &this
}

// GetSCommunicationrecipientsgroupLabel returns the SCommunicationrecipientsgroupLabel field value
func (o *CustomCommunicationrecipientsgroupResponse) GetSCommunicationrecipientsgroupLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCommunicationrecipientsgroupLabel
}

// GetSCommunicationrecipientsgroupLabelOk returns a tuple with the SCommunicationrecipientsgroupLabel field value
// and a boolean to check if the value has been set.
func (o *CustomCommunicationrecipientsgroupResponse) GetSCommunicationrecipientsgroupLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCommunicationrecipientsgroupLabel, true
}

// SetSCommunicationrecipientsgroupLabel sets field value
func (o *CustomCommunicationrecipientsgroupResponse) SetSCommunicationrecipientsgroupLabel(v string) {
	o.SCommunicationrecipientsgroupLabel = v
}

// GetAObjCommunicationrecipientsrecipient returns the AObjCommunicationrecipientsrecipient field value
func (o *CustomCommunicationrecipientsgroupResponse) GetAObjCommunicationrecipientsrecipient() []CustomCommunicationrecipientsrecipientResponse {
	if o == nil {
		var ret []CustomCommunicationrecipientsrecipientResponse
		return ret
	}

	return o.AObjCommunicationrecipientsrecipient
}

// GetAObjCommunicationrecipientsrecipientOk returns a tuple with the AObjCommunicationrecipientsrecipient field value
// and a boolean to check if the value has been set.
func (o *CustomCommunicationrecipientsgroupResponse) GetAObjCommunicationrecipientsrecipientOk() ([]CustomCommunicationrecipientsrecipientResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunicationrecipientsrecipient, true
}

// SetAObjCommunicationrecipientsrecipient sets field value
func (o *CustomCommunicationrecipientsgroupResponse) SetAObjCommunicationrecipientsrecipient(v []CustomCommunicationrecipientsrecipientResponse) {
	o.AObjCommunicationrecipientsrecipient = v
}

func (o CustomCommunicationrecipientsgroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCommunicationrecipientsgroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sCommunicationrecipientsgroupLabel"] = o.SCommunicationrecipientsgroupLabel
	toSerialize["a_objCommunicationrecipientsrecipient"] = o.AObjCommunicationrecipientsrecipient
	return toSerialize, nil
}

type NullableCustomCommunicationrecipientsgroupResponse struct {
	value *CustomCommunicationrecipientsgroupResponse
	isSet bool
}

func (v NullableCustomCommunicationrecipientsgroupResponse) Get() *CustomCommunicationrecipientsgroupResponse {
	return v.value
}

func (v *NullableCustomCommunicationrecipientsgroupResponse) Set(val *CustomCommunicationrecipientsgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCommunicationrecipientsgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCommunicationrecipientsgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCommunicationrecipientsgroupResponse(val *CustomCommunicationrecipientsgroupResponse) *NullableCustomCommunicationrecipientsgroupResponse {
	return &NullableCustomCommunicationrecipientsgroupResponse{value: val, isSet: true}
}

func (v NullableCustomCommunicationrecipientsgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCommunicationrecipientsgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


