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

// checks if the EzsignannotationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignannotationRequest{}

// EzsignannotationRequest A Ezsignannotation Object
type EzsignannotationRequest struct {
	// The unique ID of the Ezsignannotation
	PkiEzsignannotationID *int32 `json:"pkiEzsignannotationID,omitempty"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID int32 `json:"fkiEzsigndocumentID"`
	EEzsignannotationHorizontalalignment *EnumHorizontalalignment `json:"eEzsignannotationHorizontalalignment,omitempty"`
	EEzsignannotationVerticalalignment *EnumVerticalalignment `json:"eEzsignannotationVerticalalignment,omitempty"`
	EEzsignannotationType FieldEEzsignannotationType `json:"eEzsignannotationType"`
	// The X coordinate (Horizontal) where to put the Ezsignannotation on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignannotation 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignannotationX int32 `json:"iEzsignannotationX"`
	// The Y coordinate (Vertical) where to put the Ezsignannotation on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignannotation 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignannotationY int32 `json:"iEzsignannotationY"`
	// The Width of the Ezsignannotation.  Width is calculated at 100dpi (dot per inch). So for example, if you want to have the width of the Ezsignannotation to be 3 inches, you would use \"300\" for the Width.
	IEzsignannotationWidth *int32 `json:"iEzsignannotationWidth,omitempty"`
	// The Height of the Ezsignannotation.  Height is calculated at 100dpi (dot per inch). So for example, if you want to have the height of the Ezsignannotation to be 2 inches, you would use \"200\" for the Height.  This can only be set if eEzsignannotationType is **StrikethroughBlock** or **Text**
	IEzsignannotationHeight *int32 `json:"iEzsignannotationHeight,omitempty"`
	// The Text of the Ezsignannotation
	SEzsignannotationText *string `json:"sEzsignannotationText,omitempty"`
	// The page number in the Ezsigndocument
	IEzsignpagePagenumber int32 `json:"iEzsignpagePagenumber"`
}

type _EzsignannotationRequest EzsignannotationRequest

// NewEzsignannotationRequest instantiates a new EzsignannotationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignannotationRequest(fkiEzsigndocumentID int32, eEzsignannotationType FieldEEzsignannotationType, iEzsignannotationX int32, iEzsignannotationY int32, iEzsignpagePagenumber int32) *EzsignannotationRequest {
	this := EzsignannotationRequest{}
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	this.EEzsignannotationType = eEzsignannotationType
	this.IEzsignannotationX = iEzsignannotationX
	this.IEzsignannotationY = iEzsignannotationY
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	return &this
}

// NewEzsignannotationRequestWithDefaults instantiates a new EzsignannotationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignannotationRequestWithDefaults() *EzsignannotationRequest {
	this := EzsignannotationRequest{}
	return &this
}

// GetPkiEzsignannotationID returns the PkiEzsignannotationID field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetPkiEzsignannotationID() int32 {
	if o == nil || IsNil(o.PkiEzsignannotationID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignannotationID
}

// GetPkiEzsignannotationIDOk returns a tuple with the PkiEzsignannotationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetPkiEzsignannotationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignannotationID) {
		return nil, false
	}
	return o.PkiEzsignannotationID, true
}

// HasPkiEzsignannotationID returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasPkiEzsignannotationID() bool {
	if o != nil && !IsNil(o.PkiEzsignannotationID) {
		return true
	}

	return false
}

// SetPkiEzsignannotationID gets a reference to the given int32 and assigns it to the PkiEzsignannotationID field.
func (o *EzsignannotationRequest) SetPkiEzsignannotationID(v int32) {
	o.PkiEzsignannotationID = &v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value
func (o *EzsignannotationRequest) GetFkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigndocumentID, true
}

// SetFkiEzsigndocumentID sets field value
func (o *EzsignannotationRequest) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = v
}

// GetEEzsignannotationHorizontalalignment returns the EEzsignannotationHorizontalalignment field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetEEzsignannotationHorizontalalignment() EnumHorizontalalignment {
	if o == nil || IsNil(o.EEzsignannotationHorizontalalignment) {
		var ret EnumHorizontalalignment
		return ret
	}
	return *o.EEzsignannotationHorizontalalignment
}

// GetEEzsignannotationHorizontalalignmentOk returns a tuple with the EEzsignannotationHorizontalalignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetEEzsignannotationHorizontalalignmentOk() (*EnumHorizontalalignment, bool) {
	if o == nil || IsNil(o.EEzsignannotationHorizontalalignment) {
		return nil, false
	}
	return o.EEzsignannotationHorizontalalignment, true
}

// HasEEzsignannotationHorizontalalignment returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasEEzsignannotationHorizontalalignment() bool {
	if o != nil && !IsNil(o.EEzsignannotationHorizontalalignment) {
		return true
	}

	return false
}

// SetEEzsignannotationHorizontalalignment gets a reference to the given EnumHorizontalalignment and assigns it to the EEzsignannotationHorizontalalignment field.
func (o *EzsignannotationRequest) SetEEzsignannotationHorizontalalignment(v EnumHorizontalalignment) {
	o.EEzsignannotationHorizontalalignment = &v
}

// GetEEzsignannotationVerticalalignment returns the EEzsignannotationVerticalalignment field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetEEzsignannotationVerticalalignment() EnumVerticalalignment {
	if o == nil || IsNil(o.EEzsignannotationVerticalalignment) {
		var ret EnumVerticalalignment
		return ret
	}
	return *o.EEzsignannotationVerticalalignment
}

// GetEEzsignannotationVerticalalignmentOk returns a tuple with the EEzsignannotationVerticalalignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetEEzsignannotationVerticalalignmentOk() (*EnumVerticalalignment, bool) {
	if o == nil || IsNil(o.EEzsignannotationVerticalalignment) {
		return nil, false
	}
	return o.EEzsignannotationVerticalalignment, true
}

// HasEEzsignannotationVerticalalignment returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasEEzsignannotationVerticalalignment() bool {
	if o != nil && !IsNil(o.EEzsignannotationVerticalalignment) {
		return true
	}

	return false
}

// SetEEzsignannotationVerticalalignment gets a reference to the given EnumVerticalalignment and assigns it to the EEzsignannotationVerticalalignment field.
func (o *EzsignannotationRequest) SetEEzsignannotationVerticalalignment(v EnumVerticalalignment) {
	o.EEzsignannotationVerticalalignment = &v
}

// GetEEzsignannotationType returns the EEzsignannotationType field value
func (o *EzsignannotationRequest) GetEEzsignannotationType() FieldEEzsignannotationType {
	if o == nil {
		var ret FieldEEzsignannotationType
		return ret
	}

	return o.EEzsignannotationType
}

// GetEEzsignannotationTypeOk returns a tuple with the EEzsignannotationType field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetEEzsignannotationTypeOk() (*FieldEEzsignannotationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignannotationType, true
}

// SetEEzsignannotationType sets field value
func (o *EzsignannotationRequest) SetEEzsignannotationType(v FieldEEzsignannotationType) {
	o.EEzsignannotationType = v
}

// GetIEzsignannotationX returns the IEzsignannotationX field value
func (o *EzsignannotationRequest) GetIEzsignannotationX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignannotationX
}

// GetIEzsignannotationXOk returns a tuple with the IEzsignannotationX field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetIEzsignannotationXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignannotationX, true
}

// SetIEzsignannotationX sets field value
func (o *EzsignannotationRequest) SetIEzsignannotationX(v int32) {
	o.IEzsignannotationX = v
}

// GetIEzsignannotationY returns the IEzsignannotationY field value
func (o *EzsignannotationRequest) GetIEzsignannotationY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignannotationY
}

// GetIEzsignannotationYOk returns a tuple with the IEzsignannotationY field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetIEzsignannotationYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignannotationY, true
}

// SetIEzsignannotationY sets field value
func (o *EzsignannotationRequest) SetIEzsignannotationY(v int32) {
	o.IEzsignannotationY = v
}

// GetIEzsignannotationWidth returns the IEzsignannotationWidth field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetIEzsignannotationWidth() int32 {
	if o == nil || IsNil(o.IEzsignannotationWidth) {
		var ret int32
		return ret
	}
	return *o.IEzsignannotationWidth
}

// GetIEzsignannotationWidthOk returns a tuple with the IEzsignannotationWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetIEzsignannotationWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsignannotationWidth) {
		return nil, false
	}
	return o.IEzsignannotationWidth, true
}

// HasIEzsignannotationWidth returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasIEzsignannotationWidth() bool {
	if o != nil && !IsNil(o.IEzsignannotationWidth) {
		return true
	}

	return false
}

// SetIEzsignannotationWidth gets a reference to the given int32 and assigns it to the IEzsignannotationWidth field.
func (o *EzsignannotationRequest) SetIEzsignannotationWidth(v int32) {
	o.IEzsignannotationWidth = &v
}

// GetIEzsignannotationHeight returns the IEzsignannotationHeight field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetIEzsignannotationHeight() int32 {
	if o == nil || IsNil(o.IEzsignannotationHeight) {
		var ret int32
		return ret
	}
	return *o.IEzsignannotationHeight
}

// GetIEzsignannotationHeightOk returns a tuple with the IEzsignannotationHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetIEzsignannotationHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsignannotationHeight) {
		return nil, false
	}
	return o.IEzsignannotationHeight, true
}

// HasIEzsignannotationHeight returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasIEzsignannotationHeight() bool {
	if o != nil && !IsNil(o.IEzsignannotationHeight) {
		return true
	}

	return false
}

// SetIEzsignannotationHeight gets a reference to the given int32 and assigns it to the IEzsignannotationHeight field.
func (o *EzsignannotationRequest) SetIEzsignannotationHeight(v int32) {
	o.IEzsignannotationHeight = &v
}

// GetSEzsignannotationText returns the SEzsignannotationText field value if set, zero value otherwise.
func (o *EzsignannotationRequest) GetSEzsignannotationText() string {
	if o == nil || IsNil(o.SEzsignannotationText) {
		var ret string
		return ret
	}
	return *o.SEzsignannotationText
}

// GetSEzsignannotationTextOk returns a tuple with the SEzsignannotationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetSEzsignannotationTextOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignannotationText) {
		return nil, false
	}
	return o.SEzsignannotationText, true
}

// HasSEzsignannotationText returns a boolean if a field has been set.
func (o *EzsignannotationRequest) HasSEzsignannotationText() bool {
	if o != nil && !IsNil(o.SEzsignannotationText) {
		return true
	}

	return false
}

// SetSEzsignannotationText gets a reference to the given string and assigns it to the SEzsignannotationText field.
func (o *EzsignannotationRequest) SetSEzsignannotationText(v string) {
	o.SEzsignannotationText = &v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignannotationRequest) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignannotationRequest) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignannotationRequest) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

func (o EzsignannotationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignannotationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignannotationID) {
		toSerialize["pkiEzsignannotationID"] = o.PkiEzsignannotationID
	}
	toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	if !IsNil(o.EEzsignannotationHorizontalalignment) {
		toSerialize["eEzsignannotationHorizontalalignment"] = o.EEzsignannotationHorizontalalignment
	}
	if !IsNil(o.EEzsignannotationVerticalalignment) {
		toSerialize["eEzsignannotationVerticalalignment"] = o.EEzsignannotationVerticalalignment
	}
	toSerialize["eEzsignannotationType"] = o.EEzsignannotationType
	toSerialize["iEzsignannotationX"] = o.IEzsignannotationX
	toSerialize["iEzsignannotationY"] = o.IEzsignannotationY
	if !IsNil(o.IEzsignannotationWidth) {
		toSerialize["iEzsignannotationWidth"] = o.IEzsignannotationWidth
	}
	if !IsNil(o.IEzsignannotationHeight) {
		toSerialize["iEzsignannotationHeight"] = o.IEzsignannotationHeight
	}
	if !IsNil(o.SEzsignannotationText) {
		toSerialize["sEzsignannotationText"] = o.SEzsignannotationText
	}
	toSerialize["iEzsignpagePagenumber"] = o.IEzsignpagePagenumber
	return toSerialize, nil
}

func (o *EzsignannotationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsigndocumentID",
		"eEzsignannotationType",
		"iEzsignannotationX",
		"iEzsignannotationY",
		"iEzsignpagePagenumber",
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

	varEzsignannotationRequest := _EzsignannotationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignannotationRequest)

	if err != nil {
		return err
	}

	*o = EzsignannotationRequest(varEzsignannotationRequest)

	return err
}

type NullableEzsignannotationRequest struct {
	value *EzsignannotationRequest
	isSet bool
}

func (v NullableEzsignannotationRequest) Get() *EzsignannotationRequest {
	return v.value
}

func (v *NullableEzsignannotationRequest) Set(val *EzsignannotationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignannotationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignannotationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignannotationRequest(val *EzsignannotationRequest) *NullableEzsignannotationRequest {
	return &NullableEzsignannotationRequest{value: val, isSet: true}
}

func (v NullableEzsignannotationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignannotationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


