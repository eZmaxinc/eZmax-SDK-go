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

// checks if the UsergroupexternalAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupexternalAutocompleteElementResponse{}

// UsergroupexternalAutocompleteElementResponse A Usergroupexternal AutocompleteElement Response
type UsergroupexternalAutocompleteElementResponse struct {
	// The unique ID of the Usergroupexternal
	PkiUsergroupexternalID int32 `json:"pkiUsergroupexternalID"`
	// The name of the Usergroupexternal
	SUsergroupexternalName string `json:"sUsergroupexternalName"`
	// Whether the Usergroupexternal is active or not
	BUsergroupexternalIsactive bool `json:"bUsergroupexternalIsactive"`
}

type _UsergroupexternalAutocompleteElementResponse UsergroupexternalAutocompleteElementResponse

// NewUsergroupexternalAutocompleteElementResponse instantiates a new UsergroupexternalAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupexternalAutocompleteElementResponse(pkiUsergroupexternalID int32, sUsergroupexternalName string, bUsergroupexternalIsactive bool) *UsergroupexternalAutocompleteElementResponse {
	this := UsergroupexternalAutocompleteElementResponse{}
	this.PkiUsergroupexternalID = pkiUsergroupexternalID
	this.SUsergroupexternalName = sUsergroupexternalName
	this.BUsergroupexternalIsactive = bUsergroupexternalIsactive
	return &this
}

// NewUsergroupexternalAutocompleteElementResponseWithDefaults instantiates a new UsergroupexternalAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupexternalAutocompleteElementResponseWithDefaults() *UsergroupexternalAutocompleteElementResponse {
	this := UsergroupexternalAutocompleteElementResponse{}
	return &this
}

// GetPkiUsergroupexternalID returns the PkiUsergroupexternalID field value
func (o *UsergroupexternalAutocompleteElementResponse) GetPkiUsergroupexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupexternalID
}

// GetPkiUsergroupexternalIDOk returns a tuple with the PkiUsergroupexternalID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalAutocompleteElementResponse) GetPkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupexternalID, true
}

// SetPkiUsergroupexternalID sets field value
func (o *UsergroupexternalAutocompleteElementResponse) SetPkiUsergroupexternalID(v int32) {
	o.PkiUsergroupexternalID = v
}

// GetSUsergroupexternalName returns the SUsergroupexternalName field value
func (o *UsergroupexternalAutocompleteElementResponse) GetSUsergroupexternalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalName
}

// GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalAutocompleteElementResponse) GetSUsergroupexternalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalName, true
}

// SetSUsergroupexternalName sets field value
func (o *UsergroupexternalAutocompleteElementResponse) SetSUsergroupexternalName(v string) {
	o.SUsergroupexternalName = v
}

// GetBUsergroupexternalIsactive returns the BUsergroupexternalIsactive field value
func (o *UsergroupexternalAutocompleteElementResponse) GetBUsergroupexternalIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BUsergroupexternalIsactive
}

// GetBUsergroupexternalIsactiveOk returns a tuple with the BUsergroupexternalIsactive field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalAutocompleteElementResponse) GetBUsergroupexternalIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BUsergroupexternalIsactive, true
}

// SetBUsergroupexternalIsactive sets field value
func (o *UsergroupexternalAutocompleteElementResponse) SetBUsergroupexternalIsactive(v bool) {
	o.BUsergroupexternalIsactive = v
}

func (o UsergroupexternalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupexternalAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupexternalID"] = o.PkiUsergroupexternalID
	toSerialize["sUsergroupexternalName"] = o.SUsergroupexternalName
	toSerialize["bUsergroupexternalIsactive"] = o.BUsergroupexternalIsactive
	return toSerialize, nil
}

func (o *UsergroupexternalAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupexternalID",
		"sUsergroupexternalName",
		"bUsergroupexternalIsactive",
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

	varUsergroupexternalAutocompleteElementResponse := _UsergroupexternalAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupexternalAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = UsergroupexternalAutocompleteElementResponse(varUsergroupexternalAutocompleteElementResponse)

	return err
}

type NullableUsergroupexternalAutocompleteElementResponse struct {
	value *UsergroupexternalAutocompleteElementResponse
	isSet bool
}

func (v NullableUsergroupexternalAutocompleteElementResponse) Get() *UsergroupexternalAutocompleteElementResponse {
	return v.value
}

func (v *NullableUsergroupexternalAutocompleteElementResponse) Set(val *UsergroupexternalAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupexternalAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupexternalAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupexternalAutocompleteElementResponse(val *UsergroupexternalAutocompleteElementResponse) *NullableUsergroupexternalAutocompleteElementResponse {
	return &NullableUsergroupexternalAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableUsergroupexternalAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupexternalAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

