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

// checks if the CustomEzsignfoldertransmissionSignerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignfoldertransmissionSignerResponse{}

// CustomEzsignfoldertransmissionSignerResponse A form Signer Object in the context of an Ezsignfoldertransmissions
type CustomEzsignfoldertransmissionSignerResponse struct {
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The First name of the contact
	SContactFirstname string `json:"sContactFirstname"`
	// The Last name of the contact
	SContactLastname string `json:"sContactLastname"`
}

type _CustomEzsignfoldertransmissionSignerResponse CustomEzsignfoldertransmissionSignerResponse

// NewCustomEzsignfoldertransmissionSignerResponse instantiates a new CustomEzsignfoldertransmissionSignerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignfoldertransmissionSignerResponse(sContactFirstname string, sContactLastname string) *CustomEzsignfoldertransmissionSignerResponse {
	this := CustomEzsignfoldertransmissionSignerResponse{}
	this.SContactFirstname = sContactFirstname
	this.SContactLastname = sContactLastname
	return &this
}

// NewCustomEzsignfoldertransmissionSignerResponseWithDefaults instantiates a new CustomEzsignfoldertransmissionSignerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignfoldertransmissionSignerResponseWithDefaults() *CustomEzsignfoldertransmissionSignerResponse {
	this := CustomEzsignfoldertransmissionSignerResponse{}
	return &this
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *CustomEzsignfoldertransmissionSignerResponse) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertransmissionSignerResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *CustomEzsignfoldertransmissionSignerResponse) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *CustomEzsignfoldertransmissionSignerResponse) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetSContactFirstname returns the SContactFirstname field value
func (o *CustomEzsignfoldertransmissionSignerResponse) GetSContactFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactFirstname
}

// GetSContactFirstnameOk returns a tuple with the SContactFirstname field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertransmissionSignerResponse) GetSContactFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactFirstname, true
}

// SetSContactFirstname sets field value
func (o *CustomEzsignfoldertransmissionSignerResponse) SetSContactFirstname(v string) {
	o.SContactFirstname = v
}

// GetSContactLastname returns the SContactLastname field value
func (o *CustomEzsignfoldertransmissionSignerResponse) GetSContactLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SContactLastname
}

// GetSContactLastnameOk returns a tuple with the SContactLastname field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertransmissionSignerResponse) GetSContactLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SContactLastname, true
}

// SetSContactLastname sets field value
func (o *CustomEzsignfoldertransmissionSignerResponse) SetSContactLastname(v string) {
	o.SContactLastname = v
}

func (o CustomEzsignfoldertransmissionSignerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignfoldertransmissionSignerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	toSerialize["sContactFirstname"] = o.SContactFirstname
	toSerialize["sContactLastname"] = o.SContactLastname
	return toSerialize, nil
}

func (o *CustomEzsignfoldertransmissionSignerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sContactFirstname",
		"sContactLastname",
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

	varCustomEzsignfoldertransmissionSignerResponse := _CustomEzsignfoldertransmissionSignerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEzsignfoldertransmissionSignerResponse)

	if err != nil {
		return err
	}

	*o = CustomEzsignfoldertransmissionSignerResponse(varCustomEzsignfoldertransmissionSignerResponse)

	return err
}

type NullableCustomEzsignfoldertransmissionSignerResponse struct {
	value *CustomEzsignfoldertransmissionSignerResponse
	isSet bool
}

func (v NullableCustomEzsignfoldertransmissionSignerResponse) Get() *CustomEzsignfoldertransmissionSignerResponse {
	return v.value
}

func (v *NullableCustomEzsignfoldertransmissionSignerResponse) Set(val *CustomEzsignfoldertransmissionSignerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignfoldertransmissionSignerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignfoldertransmissionSignerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignfoldertransmissionSignerResponse(val *CustomEzsignfoldertransmissionSignerResponse) *NullableCustomEzsignfoldertransmissionSignerResponse {
	return &NullableCustomEzsignfoldertransmissionSignerResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignfoldertransmissionSignerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignfoldertransmissionSignerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


