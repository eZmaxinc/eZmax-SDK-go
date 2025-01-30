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
	"bytes"
	"fmt"
)

// checks if the ApikeyListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyListElement{}

// ApikeyListElement A Branding List Element
type ApikeyListElement struct {
	// The unique ID of the Apikey
	PkiApikeyID int32 `json:"pkiApikeyID"`
	// The description of the Apikey in the language of the requester
	SApikeyDescriptionX string `json:"sApikeyDescriptionX"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// Whether the apikey is active or not
	BApikeyIsactive bool `json:"bApikeyIsactive"`
	// Whether the apikey is signed or not
	BApikeyIssigned bool `json:"bApikeyIssigned"`
}

type _ApikeyListElement ApikeyListElement

// NewApikeyListElement instantiates a new ApikeyListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyListElement(pkiApikeyID int32, sApikeyDescriptionX string, sUserFirstname string, sUserLastname string, bApikeyIsactive bool, bApikeyIssigned bool) *ApikeyListElement {
	this := ApikeyListElement{}
	this.PkiApikeyID = pkiApikeyID
	this.SApikeyDescriptionX = sApikeyDescriptionX
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.BApikeyIsactive = bApikeyIsactive
	this.BApikeyIssigned = bApikeyIssigned
	return &this
}

// NewApikeyListElementWithDefaults instantiates a new ApikeyListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyListElementWithDefaults() *ApikeyListElement {
	this := ApikeyListElement{}
	return &this
}

// GetPkiApikeyID returns the PkiApikeyID field value
func (o *ApikeyListElement) GetPkiApikeyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiApikeyID
}

// GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetPkiApikeyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiApikeyID, true
}

// SetPkiApikeyID sets field value
func (o *ApikeyListElement) SetPkiApikeyID(v int32) {
	o.PkiApikeyID = v
}

// GetSApikeyDescriptionX returns the SApikeyDescriptionX field value
func (o *ApikeyListElement) GetSApikeyDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SApikeyDescriptionX
}

// GetSApikeyDescriptionXOk returns a tuple with the SApikeyDescriptionX field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetSApikeyDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SApikeyDescriptionX, true
}

// SetSApikeyDescriptionX sets field value
func (o *ApikeyListElement) SetSApikeyDescriptionX(v string) {
	o.SApikeyDescriptionX = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *ApikeyListElement) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *ApikeyListElement) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *ApikeyListElement) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *ApikeyListElement) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetBApikeyIsactive returns the BApikeyIsactive field value
func (o *ApikeyListElement) GetBApikeyIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BApikeyIsactive
}

// GetBApikeyIsactiveOk returns a tuple with the BApikeyIsactive field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetBApikeyIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BApikeyIsactive, true
}

// SetBApikeyIsactive sets field value
func (o *ApikeyListElement) SetBApikeyIsactive(v bool) {
	o.BApikeyIsactive = v
}

// GetBApikeyIssigned returns the BApikeyIssigned field value
func (o *ApikeyListElement) GetBApikeyIssigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BApikeyIssigned
}

// GetBApikeyIssignedOk returns a tuple with the BApikeyIssigned field value
// and a boolean to check if the value has been set.
func (o *ApikeyListElement) GetBApikeyIssignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BApikeyIssigned, true
}

// SetBApikeyIssigned sets field value
func (o *ApikeyListElement) SetBApikeyIssigned(v bool) {
	o.BApikeyIssigned = v
}

func (o ApikeyListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiApikeyID"] = o.PkiApikeyID
	toSerialize["sApikeyDescriptionX"] = o.SApikeyDescriptionX
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["bApikeyIsactive"] = o.BApikeyIsactive
	toSerialize["bApikeyIssigned"] = o.BApikeyIssigned
	return toSerialize, nil
}

func (o *ApikeyListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiApikeyID",
		"sApikeyDescriptionX",
		"sUserFirstname",
		"sUserLastname",
		"bApikeyIsactive",
		"bApikeyIssigned",
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

	varApikeyListElement := _ApikeyListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyListElement)

	if err != nil {
		return err
	}

	*o = ApikeyListElement(varApikeyListElement)

	return err
}

type NullableApikeyListElement struct {
	value *ApikeyListElement
	isSet bool
}

func (v NullableApikeyListElement) Get() *ApikeyListElement {
	return v.value
}

func (v *NullableApikeyListElement) Set(val *ApikeyListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyListElement(val *ApikeyListElement) *NullableApikeyListElement {
	return &NullableApikeyListElement{value: val, isSet: true}
}

func (v NullableApikeyListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


