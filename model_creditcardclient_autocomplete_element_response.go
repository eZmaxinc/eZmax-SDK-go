/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreditcardclientAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientAutocompleteElementResponse{}

// CreditcardclientAutocompleteElementResponse A Creditcardclient AutocompleteElement Response
type CreditcardclientAutocompleteElementResponse struct {
	// The unique ID of the Creditcardclient
	PkiCreditcardclientID int32 `json:"pkiCreditcardclientID"`
	// The description of the Creditcardclient
	SCreditcardclientDescription string `json:"sCreditcardclientDescription"`
	// Whether the creditcardclient is active or not
	BCreditcardclientIsactive bool `json:"bCreditcardclientIsactive"`
}

type _CreditcardclientAutocompleteElementResponse CreditcardclientAutocompleteElementResponse

// NewCreditcardclientAutocompleteElementResponse instantiates a new CreditcardclientAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientAutocompleteElementResponse(pkiCreditcardclientID int32, sCreditcardclientDescription string, bCreditcardclientIsactive bool) *CreditcardclientAutocompleteElementResponse {
	this := CreditcardclientAutocompleteElementResponse{}
	this.PkiCreditcardclientID = pkiCreditcardclientID
	this.SCreditcardclientDescription = sCreditcardclientDescription
	this.BCreditcardclientIsactive = bCreditcardclientIsactive
	return &this
}

// NewCreditcardclientAutocompleteElementResponseWithDefaults instantiates a new CreditcardclientAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientAutocompleteElementResponseWithDefaults() *CreditcardclientAutocompleteElementResponse {
	this := CreditcardclientAutocompleteElementResponse{}
	return &this
}

// GetPkiCreditcardclientID returns the PkiCreditcardclientID field value
func (o *CreditcardclientAutocompleteElementResponse) GetPkiCreditcardclientID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCreditcardclientID
}

// GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientAutocompleteElementResponse) GetPkiCreditcardclientIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCreditcardclientID, true
}

// SetPkiCreditcardclientID sets field value
func (o *CreditcardclientAutocompleteElementResponse) SetPkiCreditcardclientID(v int32) {
	o.PkiCreditcardclientID = v
}

// GetSCreditcardclientDescription returns the SCreditcardclientDescription field value
func (o *CreditcardclientAutocompleteElementResponse) GetSCreditcardclientDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardclientDescription
}

// GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientAutocompleteElementResponse) GetSCreditcardclientDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardclientDescription, true
}

// SetSCreditcardclientDescription sets field value
func (o *CreditcardclientAutocompleteElementResponse) SetSCreditcardclientDescription(v string) {
	o.SCreditcardclientDescription = v
}

// GetBCreditcardclientIsactive returns the BCreditcardclientIsactive field value
func (o *CreditcardclientAutocompleteElementResponse) GetBCreditcardclientIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientIsactive
}

// GetBCreditcardclientIsactiveOk returns a tuple with the BCreditcardclientIsactive field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientAutocompleteElementResponse) GetBCreditcardclientIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientIsactive, true
}

// SetBCreditcardclientIsactive sets field value
func (o *CreditcardclientAutocompleteElementResponse) SetBCreditcardclientIsactive(v bool) {
	o.BCreditcardclientIsactive = v
}

func (o CreditcardclientAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCreditcardclientID"] = o.PkiCreditcardclientID
	toSerialize["sCreditcardclientDescription"] = o.SCreditcardclientDescription
	toSerialize["bCreditcardclientIsactive"] = o.BCreditcardclientIsactive
	return toSerialize, nil
}

func (o *CreditcardclientAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiCreditcardclientID",
		"sCreditcardclientDescription",
		"bCreditcardclientIsactive",
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

	varCreditcardclientAutocompleteElementResponse := _CreditcardclientAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = CreditcardclientAutocompleteElementResponse(varCreditcardclientAutocompleteElementResponse)

	return err
}

type NullableCreditcardclientAutocompleteElementResponse struct {
	value *CreditcardclientAutocompleteElementResponse
	isSet bool
}

func (v NullableCreditcardclientAutocompleteElementResponse) Get() *CreditcardclientAutocompleteElementResponse {
	return v.value
}

func (v *NullableCreditcardclientAutocompleteElementResponse) Set(val *CreditcardclientAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientAutocompleteElementResponse(val *CreditcardclientAutocompleteElementResponse) *NullableCreditcardclientAutocompleteElementResponse {
	return &NullableCreditcardclientAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableCreditcardclientAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

