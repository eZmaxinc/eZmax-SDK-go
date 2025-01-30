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

// checks if the EzsigntemplatedocumentpagerecognitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentpagerecognitionRequest{}

// EzsigntemplatedocumentpagerecognitionRequest A Ezsigntemplatedocumentpagerecognition Object
type EzsigntemplatedocumentpagerecognitionRequest struct {
	// The unique ID of the Ezsigntemplatedocumentpagerecognition
	PkiEzsigntemplatedocumentpagerecognitionID *int32 `json:"pkiEzsigntemplatedocumentpagerecognitionID,omitempty"`
	// The unique ID of the Ezsigntemplatedocumentpage
	FkiEzsigntemplatedocumentpageID int32 `json:"fkiEzsigntemplatedocumentpageID"`
	EEzsigntemplatedocumentpagerecognitionOperator FieldEEzsigntemplatedocumentpagerecognitionOperator `json:"eEzsigntemplatedocumentpagerecognitionOperator"`
	EEzsigntemplatedocumentpagerecognitionSection FieldEEzsigntemplatedocumentpagerecognitionSection `json:"eEzsigntemplatedocumentpagerecognitionSection"`
	// The similarpercentage of the Ezsigntemplatedocumentpagerecognition
	IEzsigntemplatedocumentpagerecognitionSimilarpercentage *int32 `json:"iEzsigntemplatedocumentpagerecognitionSimilarpercentage,omitempty"`
	// The x of the Ezsigntemplatedocumentpagerecognition
	IEzsigntemplatedocumentpagerecognitionX *int32 `json:"iEzsigntemplatedocumentpagerecognitionX,omitempty"`
	// The y of the Ezsigntemplatedocumentpagerecognition
	IEzsigntemplatedocumentpagerecognitionY *int32 `json:"iEzsigntemplatedocumentpagerecognitionY,omitempty"`
	// The width of the Ezsigntemplatedocumentpagerecognition
	IEzsigntemplatedocumentpagerecognitionWidth *int32 `json:"iEzsigntemplatedocumentpagerecognitionWidth,omitempty"`
	// The height of the Ezsigntemplatedocumentpagerecognition
	IEzsigntemplatedocumentpagerecognitionHeight *int32 `json:"iEzsigntemplatedocumentpagerecognitionHeight,omitempty"`
	// The text of the Ezsigntemplatedocumentpagerecognition
	TEzsigntemplatedocumentpagerecognitionText string `json:"tEzsigntemplatedocumentpagerecognitionText" validate:"regexp=^[.\\\\D\\\\d]{0,65535}$"`
}

type _EzsigntemplatedocumentpagerecognitionRequest EzsigntemplatedocumentpagerecognitionRequest

// NewEzsigntemplatedocumentpagerecognitionRequest instantiates a new EzsigntemplatedocumentpagerecognitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentpagerecognitionRequest(fkiEzsigntemplatedocumentpageID int32, eEzsigntemplatedocumentpagerecognitionOperator FieldEEzsigntemplatedocumentpagerecognitionOperator, eEzsigntemplatedocumentpagerecognitionSection FieldEEzsigntemplatedocumentpagerecognitionSection, tEzsigntemplatedocumentpagerecognitionText string) *EzsigntemplatedocumentpagerecognitionRequest {
	this := EzsigntemplatedocumentpagerecognitionRequest{}
	this.FkiEzsigntemplatedocumentpageID = fkiEzsigntemplatedocumentpageID
	this.EEzsigntemplatedocumentpagerecognitionOperator = eEzsigntemplatedocumentpagerecognitionOperator
	this.EEzsigntemplatedocumentpagerecognitionSection = eEzsigntemplatedocumentpagerecognitionSection
	this.TEzsigntemplatedocumentpagerecognitionText = tEzsigntemplatedocumentpagerecognitionText
	return &this
}

// NewEzsigntemplatedocumentpagerecognitionRequestWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentpagerecognitionRequestWithDefaults() *EzsigntemplatedocumentpagerecognitionRequest {
	this := EzsigntemplatedocumentpagerecognitionRequest{}
	return &this
}

// GetPkiEzsigntemplatedocumentpagerecognitionID returns the PkiEzsigntemplatedocumentpagerecognitionID field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetPkiEzsigntemplatedocumentpagerecognitionID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatedocumentpagerecognitionID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatedocumentpagerecognitionID
}

// GetPkiEzsigntemplatedocumentpagerecognitionIDOk returns a tuple with the PkiEzsigntemplatedocumentpagerecognitionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetPkiEzsigntemplatedocumentpagerecognitionIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatedocumentpagerecognitionID) {
		return nil, false
	}
	return o.PkiEzsigntemplatedocumentpagerecognitionID, true
}

// HasPkiEzsigntemplatedocumentpagerecognitionID returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasPkiEzsigntemplatedocumentpagerecognitionID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatedocumentpagerecognitionID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatedocumentpagerecognitionID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatedocumentpagerecognitionID field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetPkiEzsigntemplatedocumentpagerecognitionID(v int32) {
	o.PkiEzsigntemplatedocumentpagerecognitionID = &v
}

// GetFkiEzsigntemplatedocumentpageID returns the FkiEzsigntemplatedocumentpageID field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetFkiEzsigntemplatedocumentpageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatedocumentpageID
}

// GetFkiEzsigntemplatedocumentpageIDOk returns a tuple with the FkiEzsigntemplatedocumentpageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetFkiEzsigntemplatedocumentpageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatedocumentpageID, true
}

// SetFkiEzsigntemplatedocumentpageID sets field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetFkiEzsigntemplatedocumentpageID(v int32) {
	o.FkiEzsigntemplatedocumentpageID = v
}

// GetEEzsigntemplatedocumentpagerecognitionOperator returns the EEzsigntemplatedocumentpagerecognitionOperator field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionOperator() FieldEEzsigntemplatedocumentpagerecognitionOperator {
	if o == nil {
		var ret FieldEEzsigntemplatedocumentpagerecognitionOperator
		return ret
	}

	return o.EEzsigntemplatedocumentpagerecognitionOperator
}

// GetEEzsigntemplatedocumentpagerecognitionOperatorOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionOperator field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionOperatorOk() (*FieldEEzsigntemplatedocumentpagerecognitionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatedocumentpagerecognitionOperator, true
}

// SetEEzsigntemplatedocumentpagerecognitionOperator sets field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetEEzsigntemplatedocumentpagerecognitionOperator(v FieldEEzsigntemplatedocumentpagerecognitionOperator) {
	o.EEzsigntemplatedocumentpagerecognitionOperator = v
}

// GetEEzsigntemplatedocumentpagerecognitionSection returns the EEzsigntemplatedocumentpagerecognitionSection field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionSection() FieldEEzsigntemplatedocumentpagerecognitionSection {
	if o == nil {
		var ret FieldEEzsigntemplatedocumentpagerecognitionSection
		return ret
	}

	return o.EEzsigntemplatedocumentpagerecognitionSection
}

// GetEEzsigntemplatedocumentpagerecognitionSectionOk returns a tuple with the EEzsigntemplatedocumentpagerecognitionSection field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetEEzsigntemplatedocumentpagerecognitionSectionOk() (*FieldEEzsigntemplatedocumentpagerecognitionSection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsigntemplatedocumentpagerecognitionSection, true
}

// SetEEzsigntemplatedocumentpagerecognitionSection sets field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetEEzsigntemplatedocumentpagerecognitionSection(v FieldEEzsigntemplatedocumentpagerecognitionSection) {
	o.EEzsigntemplatedocumentpagerecognitionSection = v
}

// GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentage() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage
}

// GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionSimilarpercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage, true
}

// HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionSimilarpercentage() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentpagerecognitionSimilarpercentage field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionSimilarpercentage(v int32) {
	o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage = &v
}

// GetIEzsigntemplatedocumentpagerecognitionX returns the IEzsigntemplatedocumentpagerecognitionX field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionX() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionX) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentpagerecognitionX
}

// GetIEzsigntemplatedocumentpagerecognitionXOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionXOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionX) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentpagerecognitionX, true
}

// HasIEzsigntemplatedocumentpagerecognitionX returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionX() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentpagerecognitionX) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentpagerecognitionX gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentpagerecognitionX field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionX(v int32) {
	o.IEzsigntemplatedocumentpagerecognitionX = &v
}

// GetIEzsigntemplatedocumentpagerecognitionY returns the IEzsigntemplatedocumentpagerecognitionY field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionY() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionY) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentpagerecognitionY
}

// GetIEzsigntemplatedocumentpagerecognitionYOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionYOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionY) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentpagerecognitionY, true
}

// HasIEzsigntemplatedocumentpagerecognitionY returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionY() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentpagerecognitionY) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentpagerecognitionY gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentpagerecognitionY field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionY(v int32) {
	o.IEzsigntemplatedocumentpagerecognitionY = &v
}

// GetIEzsigntemplatedocumentpagerecognitionWidth returns the IEzsigntemplatedocumentpagerecognitionWidth field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionWidth() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionWidth) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentpagerecognitionWidth
}

// GetIEzsigntemplatedocumentpagerecognitionWidthOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionWidth) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentpagerecognitionWidth, true
}

// HasIEzsigntemplatedocumentpagerecognitionWidth returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionWidth() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentpagerecognitionWidth) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentpagerecognitionWidth gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentpagerecognitionWidth field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionWidth(v int32) {
	o.IEzsigntemplatedocumentpagerecognitionWidth = &v
}

// GetIEzsigntemplatedocumentpagerecognitionHeight returns the IEzsigntemplatedocumentpagerecognitionHeight field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionHeight() int32 {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionHeight) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatedocumentpagerecognitionHeight
}

// GetIEzsigntemplatedocumentpagerecognitionHeightOk returns a tuple with the IEzsigntemplatedocumentpagerecognitionHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetIEzsigntemplatedocumentpagerecognitionHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatedocumentpagerecognitionHeight) {
		return nil, false
	}
	return o.IEzsigntemplatedocumentpagerecognitionHeight, true
}

// HasIEzsigntemplatedocumentpagerecognitionHeight returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) HasIEzsigntemplatedocumentpagerecognitionHeight() bool {
	if o != nil && !IsNil(o.IEzsigntemplatedocumentpagerecognitionHeight) {
		return true
	}

	return false
}

// SetIEzsigntemplatedocumentpagerecognitionHeight gets a reference to the given int32 and assigns it to the IEzsigntemplatedocumentpagerecognitionHeight field.
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetIEzsigntemplatedocumentpagerecognitionHeight(v int32) {
	o.IEzsigntemplatedocumentpagerecognitionHeight = &v
}

// GetTEzsigntemplatedocumentpagerecognitionText returns the TEzsigntemplatedocumentpagerecognitionText field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetTEzsigntemplatedocumentpagerecognitionText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsigntemplatedocumentpagerecognitionText
}

// GetTEzsigntemplatedocumentpagerecognitionTextOk returns a tuple with the TEzsigntemplatedocumentpagerecognitionText field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionRequest) GetTEzsigntemplatedocumentpagerecognitionTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzsigntemplatedocumentpagerecognitionText, true
}

// SetTEzsigntemplatedocumentpagerecognitionText sets field value
func (o *EzsigntemplatedocumentpagerecognitionRequest) SetTEzsigntemplatedocumentpagerecognitionText(v string) {
	o.TEzsigntemplatedocumentpagerecognitionText = v
}

func (o EzsigntemplatedocumentpagerecognitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentpagerecognitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatedocumentpagerecognitionID) {
		toSerialize["pkiEzsigntemplatedocumentpagerecognitionID"] = o.PkiEzsigntemplatedocumentpagerecognitionID
	}
	toSerialize["fkiEzsigntemplatedocumentpageID"] = o.FkiEzsigntemplatedocumentpageID
	toSerialize["eEzsigntemplatedocumentpagerecognitionOperator"] = o.EEzsigntemplatedocumentpagerecognitionOperator
	toSerialize["eEzsigntemplatedocumentpagerecognitionSection"] = o.EEzsigntemplatedocumentpagerecognitionSection
	if !IsNil(o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage) {
		toSerialize["iEzsigntemplatedocumentpagerecognitionSimilarpercentage"] = o.IEzsigntemplatedocumentpagerecognitionSimilarpercentage
	}
	if !IsNil(o.IEzsigntemplatedocumentpagerecognitionX) {
		toSerialize["iEzsigntemplatedocumentpagerecognitionX"] = o.IEzsigntemplatedocumentpagerecognitionX
	}
	if !IsNil(o.IEzsigntemplatedocumentpagerecognitionY) {
		toSerialize["iEzsigntemplatedocumentpagerecognitionY"] = o.IEzsigntemplatedocumentpagerecognitionY
	}
	if !IsNil(o.IEzsigntemplatedocumentpagerecognitionWidth) {
		toSerialize["iEzsigntemplatedocumentpagerecognitionWidth"] = o.IEzsigntemplatedocumentpagerecognitionWidth
	}
	if !IsNil(o.IEzsigntemplatedocumentpagerecognitionHeight) {
		toSerialize["iEzsigntemplatedocumentpagerecognitionHeight"] = o.IEzsigntemplatedocumentpagerecognitionHeight
	}
	toSerialize["tEzsigntemplatedocumentpagerecognitionText"] = o.TEzsigntemplatedocumentpagerecognitionText
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentpagerecognitionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigntemplatedocumentpageID",
		"eEzsigntemplatedocumentpagerecognitionOperator",
		"eEzsigntemplatedocumentpagerecognitionSection",
		"tEzsigntemplatedocumentpagerecognitionText",
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

	varEzsigntemplatedocumentpagerecognitionRequest := _EzsigntemplatedocumentpagerecognitionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentpagerecognitionRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentpagerecognitionRequest(varEzsigntemplatedocumentpagerecognitionRequest)

	return err
}

type NullableEzsigntemplatedocumentpagerecognitionRequest struct {
	value *EzsigntemplatedocumentpagerecognitionRequest
	isSet bool
}

func (v NullableEzsigntemplatedocumentpagerecognitionRequest) Get() *EzsigntemplatedocumentpagerecognitionRequest {
	return v.value
}

func (v *NullableEzsigntemplatedocumentpagerecognitionRequest) Set(val *EzsigntemplatedocumentpagerecognitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentpagerecognitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentpagerecognitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentpagerecognitionRequest(val *EzsigntemplatedocumentpagerecognitionRequest) *NullableEzsigntemplatedocumentpagerecognitionRequest {
	return &NullableEzsigntemplatedocumentpagerecognitionRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentpagerecognitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentpagerecognitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


