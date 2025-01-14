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

// checks if the EzdoctemplatefieldtypecategoryAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzdoctemplatefieldtypecategoryAutocompleteElementResponse{}

// EzdoctemplatefieldtypecategoryAutocompleteElementResponse A Ezdoctemplatefieldtypecategory AutocompleteElement Response
type EzdoctemplatefieldtypecategoryAutocompleteElementResponse struct {
	// The unique ID of the Ezdoctemplatefieldtypecategory
	PkiEzdoctemplatefieldtypecategoryID int32 `json:"pkiEzdoctemplatefieldtypecategoryID"`
	// The unique ID of the Ezdoctemplatetype
	FkiEzdoctemplatetypeID int32 `json:"fkiEzdoctemplatetypeID"`
	// The description of the Ezdoctemplatefieldtypecategory in the language of the requester
	SEzdoctemplatefieldtypecategoryDescriptionX string `json:"sEzdoctemplatefieldtypecategoryDescriptionX" validate:"regexp=^.{0,55}$"`
	// Whether the Ezdoctemplatefieldtypecategory is active or not
	BEzdoctemplatefieldtypecategoryIsactive bool `json:"bEzdoctemplatefieldtypecategoryIsactive"`
}

type _EzdoctemplatefieldtypecategoryAutocompleteElementResponse EzdoctemplatefieldtypecategoryAutocompleteElementResponse

// NewEzdoctemplatefieldtypecategoryAutocompleteElementResponse instantiates a new EzdoctemplatefieldtypecategoryAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzdoctemplatefieldtypecategoryAutocompleteElementResponse(pkiEzdoctemplatefieldtypecategoryID int32, fkiEzdoctemplatetypeID int32, sEzdoctemplatefieldtypecategoryDescriptionX string, bEzdoctemplatefieldtypecategoryIsactive bool) *EzdoctemplatefieldtypecategoryAutocompleteElementResponse {
	this := EzdoctemplatefieldtypecategoryAutocompleteElementResponse{}
	this.PkiEzdoctemplatefieldtypecategoryID = pkiEzdoctemplatefieldtypecategoryID
	this.FkiEzdoctemplatetypeID = fkiEzdoctemplatetypeID
	this.SEzdoctemplatefieldtypecategoryDescriptionX = sEzdoctemplatefieldtypecategoryDescriptionX
	this.BEzdoctemplatefieldtypecategoryIsactive = bEzdoctemplatefieldtypecategoryIsactive
	return &this
}

// NewEzdoctemplatefieldtypecategoryAutocompleteElementResponseWithDefaults instantiates a new EzdoctemplatefieldtypecategoryAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzdoctemplatefieldtypecategoryAutocompleteElementResponseWithDefaults() *EzdoctemplatefieldtypecategoryAutocompleteElementResponse {
	this := EzdoctemplatefieldtypecategoryAutocompleteElementResponse{}
	return &this
}

// GetPkiEzdoctemplatefieldtypecategoryID returns the PkiEzdoctemplatefieldtypecategoryID field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetPkiEzdoctemplatefieldtypecategoryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzdoctemplatefieldtypecategoryID
}

// GetPkiEzdoctemplatefieldtypecategoryIDOk returns a tuple with the PkiEzdoctemplatefieldtypecategoryID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetPkiEzdoctemplatefieldtypecategoryIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzdoctemplatefieldtypecategoryID, true
}

// SetPkiEzdoctemplatefieldtypecategoryID sets field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) SetPkiEzdoctemplatefieldtypecategoryID(v int32) {
	o.PkiEzdoctemplatefieldtypecategoryID = v
}

// GetFkiEzdoctemplatetypeID returns the FkiEzdoctemplatetypeID field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetFkiEzdoctemplatetypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzdoctemplatetypeID
}

// GetFkiEzdoctemplatetypeIDOk returns a tuple with the FkiEzdoctemplatetypeID field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetFkiEzdoctemplatetypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzdoctemplatetypeID, true
}

// SetFkiEzdoctemplatetypeID sets field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) SetFkiEzdoctemplatetypeID(v int32) {
	o.FkiEzdoctemplatetypeID = v
}

// GetSEzdoctemplatefieldtypecategoryDescriptionX returns the SEzdoctemplatefieldtypecategoryDescriptionX field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetSEzdoctemplatefieldtypecategoryDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzdoctemplatefieldtypecategoryDescriptionX
}

// GetSEzdoctemplatefieldtypecategoryDescriptionXOk returns a tuple with the SEzdoctemplatefieldtypecategoryDescriptionX field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetSEzdoctemplatefieldtypecategoryDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzdoctemplatefieldtypecategoryDescriptionX, true
}

// SetSEzdoctemplatefieldtypecategoryDescriptionX sets field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) SetSEzdoctemplatefieldtypecategoryDescriptionX(v string) {
	o.SEzdoctemplatefieldtypecategoryDescriptionX = v
}

// GetBEzdoctemplatefieldtypecategoryIsactive returns the BEzdoctemplatefieldtypecategoryIsactive field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetBEzdoctemplatefieldtypecategoryIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzdoctemplatefieldtypecategoryIsactive
}

// GetBEzdoctemplatefieldtypecategoryIsactiveOk returns a tuple with the BEzdoctemplatefieldtypecategoryIsactive field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) GetBEzdoctemplatefieldtypecategoryIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzdoctemplatefieldtypecategoryIsactive, true
}

// SetBEzdoctemplatefieldtypecategoryIsactive sets field value
func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) SetBEzdoctemplatefieldtypecategoryIsactive(v bool) {
	o.BEzdoctemplatefieldtypecategoryIsactive = v
}

func (o EzdoctemplatefieldtypecategoryAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzdoctemplatefieldtypecategoryAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzdoctemplatefieldtypecategoryID"] = o.PkiEzdoctemplatefieldtypecategoryID
	toSerialize["fkiEzdoctemplatetypeID"] = o.FkiEzdoctemplatetypeID
	toSerialize["sEzdoctemplatefieldtypecategoryDescriptionX"] = o.SEzdoctemplatefieldtypecategoryDescriptionX
	toSerialize["bEzdoctemplatefieldtypecategoryIsactive"] = o.BEzdoctemplatefieldtypecategoryIsactive
	return toSerialize, nil
}

func (o *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzdoctemplatefieldtypecategoryID",
		"fkiEzdoctemplatetypeID",
		"sEzdoctemplatefieldtypecategoryDescriptionX",
		"bEzdoctemplatefieldtypecategoryIsactive",
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

	varEzdoctemplatefieldtypecategoryAutocompleteElementResponse := _EzdoctemplatefieldtypecategoryAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzdoctemplatefieldtypecategoryAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = EzdoctemplatefieldtypecategoryAutocompleteElementResponse(varEzdoctemplatefieldtypecategoryAutocompleteElementResponse)

	return err
}

type NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse struct {
	value *EzdoctemplatefieldtypecategoryAutocompleteElementResponse
	isSet bool
}

func (v NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) Get() *EzdoctemplatefieldtypecategoryAutocompleteElementResponse {
	return v.value
}

func (v *NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) Set(val *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse(val *EzdoctemplatefieldtypecategoryAutocompleteElementResponse) *NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse {
	return &NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzdoctemplatefieldtypecategoryAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


