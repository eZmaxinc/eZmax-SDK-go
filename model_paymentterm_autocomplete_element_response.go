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

// checks if the PaymenttermAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermAutocompleteElementResponse{}

// PaymenttermAutocompleteElementResponse A Paymentterm AutocompleteElement Response
type PaymenttermAutocompleteElementResponse struct {
	// The unique ID of the Paymentterm
	PkiPaymenttermID int32 `json:"pkiPaymenttermID"`
	// The description of the Paymentterm in the language of the requester
	SPaymenttermDescriptionX string `json:"sPaymenttermDescriptionX" validate:"regexp=^.{1,40}$"`
	// Whether the Paymentterm is active or not
	BPaymenttermIsactive bool `json:"bPaymenttermIsactive"`
}

type _PaymenttermAutocompleteElementResponse PaymenttermAutocompleteElementResponse

// NewPaymenttermAutocompleteElementResponse instantiates a new PaymenttermAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermAutocompleteElementResponse(pkiPaymenttermID int32, sPaymenttermDescriptionX string, bPaymenttermIsactive bool) *PaymenttermAutocompleteElementResponse {
	this := PaymenttermAutocompleteElementResponse{}
	this.PkiPaymenttermID = pkiPaymenttermID
	this.SPaymenttermDescriptionX = sPaymenttermDescriptionX
	this.BPaymenttermIsactive = bPaymenttermIsactive
	return &this
}

// NewPaymenttermAutocompleteElementResponseWithDefaults instantiates a new PaymenttermAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermAutocompleteElementResponseWithDefaults() *PaymenttermAutocompleteElementResponse {
	this := PaymenttermAutocompleteElementResponse{}
	return &this
}

// GetPkiPaymenttermID returns the PkiPaymenttermID field value
func (o *PaymenttermAutocompleteElementResponse) GetPkiPaymenttermID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPaymenttermID
}

// GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field value
// and a boolean to check if the value has been set.
func (o *PaymenttermAutocompleteElementResponse) GetPkiPaymenttermIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPaymenttermID, true
}

// SetPkiPaymenttermID sets field value
func (o *PaymenttermAutocompleteElementResponse) SetPkiPaymenttermID(v int32) {
	o.PkiPaymenttermID = v
}

// GetSPaymenttermDescriptionX returns the SPaymenttermDescriptionX field value
func (o *PaymenttermAutocompleteElementResponse) GetSPaymenttermDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPaymenttermDescriptionX
}

// GetSPaymenttermDescriptionXOk returns a tuple with the SPaymenttermDescriptionX field value
// and a boolean to check if the value has been set.
func (o *PaymenttermAutocompleteElementResponse) GetSPaymenttermDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SPaymenttermDescriptionX, true
}

// SetSPaymenttermDescriptionX sets field value
func (o *PaymenttermAutocompleteElementResponse) SetSPaymenttermDescriptionX(v string) {
	o.SPaymenttermDescriptionX = v
}

// GetBPaymenttermIsactive returns the BPaymenttermIsactive field value
func (o *PaymenttermAutocompleteElementResponse) GetBPaymenttermIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BPaymenttermIsactive
}

// GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field value
// and a boolean to check if the value has been set.
func (o *PaymenttermAutocompleteElementResponse) GetBPaymenttermIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BPaymenttermIsactive, true
}

// SetBPaymenttermIsactive sets field value
func (o *PaymenttermAutocompleteElementResponse) SetBPaymenttermIsactive(v bool) {
	o.BPaymenttermIsactive = v
}

func (o PaymenttermAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPaymenttermID"] = o.PkiPaymenttermID
	toSerialize["sPaymenttermDescriptionX"] = o.SPaymenttermDescriptionX
	toSerialize["bPaymenttermIsactive"] = o.BPaymenttermIsactive
	return toSerialize, nil
}

func (o *PaymenttermAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPaymenttermID",
		"sPaymenttermDescriptionX",
		"bPaymenttermIsactive",
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

	varPaymenttermAutocompleteElementResponse := _PaymenttermAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = PaymenttermAutocompleteElementResponse(varPaymenttermAutocompleteElementResponse)

	return err
}

type NullablePaymenttermAutocompleteElementResponse struct {
	value *PaymenttermAutocompleteElementResponse
	isSet bool
}

func (v NullablePaymenttermAutocompleteElementResponse) Get() *PaymenttermAutocompleteElementResponse {
	return v.value
}

func (v *NullablePaymenttermAutocompleteElementResponse) Set(val *PaymenttermAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermAutocompleteElementResponse(val *PaymenttermAutocompleteElementResponse) *NullablePaymenttermAutocompleteElementResponse {
	return &NullablePaymenttermAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullablePaymenttermAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


