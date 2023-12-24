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

// checks if the EzsignformfieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldRequest{}

// EzsignformfieldRequest A Ezsignformfield Object
type EzsignformfieldRequest struct {
	// The unique ID of the Ezsignformfield
	PkiEzsignformfieldID *int32 `json:"pkiEzsignformfieldID,omitempty"`
	// The page number in the Ezsigndocument
	IEzsignpagePagenumber int32 `json:"iEzsignpagePagenumber"`
	// The Label for the Ezsignformfield
	SEzsignformfieldLabel string `json:"sEzsignformfieldLabel"`
	// The value for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is Checkbox or Radio
	SEzsignformfieldValue *string `json:"sEzsignformfieldValue,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsignformfield on the Ezsignpage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignformfield 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignformfieldX int32 `json:"iEzsignformfieldX"`
	// The Y coordinate (Vertical) where to put the Ezsignformfield on the Ezsignpage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignformfield 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignformfieldY int32 `json:"iEzsignformfieldY"`
	// The Width of the Ezsignformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsignformfieldgroupType.  | eEzsignformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22-65535     | | Radio                     | 22           | | Text                      | 22-65535     | | Textarea                  | 22-65535     |
	IEzsignformfieldWidth int32 `json:"iEzsignformfieldWidth"`
	// The Height of the Ezsignformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsignformfieldgroupType.  | eEzsignformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22           | | Radio                     | 22           | | Text                      | 22           | | Textarea                  | 22-65535     | 
	IEzsignformfieldHeight int32 `json:"iEzsignformfieldHeight"`
	// Whether the Ezsignformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsignformfieldgroupType is **Text**
	BEzsignformfieldAutocomplete *bool `json:"bEzsignformfieldAutocomplete,omitempty"`
	// Whether the Ezsignformfield is selected or not by default.  This can only be set if eEzsignformfieldgroupType is **Checkbox** or **Radio**
	BEzsignformfieldSelected *bool `json:"bEzsignformfieldSelected,omitempty"`
	// This is the value enterred for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is **Dropdown**, **Text** or **Textarea**
	SEzsignformfieldEnteredvalue *string `json:"sEzsignformfieldEnteredvalue,omitempty"`
	EEzsignformfieldDependencyrequirement *FieldEEzsignformfieldDependencyrequirement `json:"eEzsignformfieldDependencyrequirement,omitempty"`
}

type _EzsignformfieldRequest EzsignformfieldRequest

// NewEzsignformfieldRequest instantiates a new EzsignformfieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldRequest(iEzsignpagePagenumber int32, sEzsignformfieldLabel string, iEzsignformfieldX int32, iEzsignformfieldY int32, iEzsignformfieldWidth int32, iEzsignformfieldHeight int32) *EzsignformfieldRequest {
	this := EzsignformfieldRequest{}
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.IEzsignformfieldX = iEzsignformfieldX
	this.IEzsignformfieldY = iEzsignformfieldY
	this.IEzsignformfieldWidth = iEzsignformfieldWidth
	this.IEzsignformfieldHeight = iEzsignformfieldHeight
	return &this
}

// NewEzsignformfieldRequestWithDefaults instantiates a new EzsignformfieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldRequestWithDefaults() *EzsignformfieldRequest {
	this := EzsignformfieldRequest{}
	return &this
}

// GetPkiEzsignformfieldID returns the PkiEzsignformfieldID field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetPkiEzsignformfieldID() int32 {
	if o == nil || IsNil(o.PkiEzsignformfieldID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignformfieldID
}

// GetPkiEzsignformfieldIDOk returns a tuple with the PkiEzsignformfieldID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetPkiEzsignformfieldIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignformfieldID) {
		return nil, false
	}
	return o.PkiEzsignformfieldID, true
}

// HasPkiEzsignformfieldID returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasPkiEzsignformfieldID() bool {
	if o != nil && !IsNil(o.PkiEzsignformfieldID) {
		return true
	}

	return false
}

// SetPkiEzsignformfieldID gets a reference to the given int32 and assigns it to the PkiEzsignformfieldID field.
func (o *EzsignformfieldRequest) SetPkiEzsignformfieldID(v int32) {
	o.PkiEzsignformfieldID = &v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignformfieldRequest) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignformfieldRequest) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *EzsignformfieldRequest) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *EzsignformfieldRequest) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetSEzsignformfieldValue returns the SEzsignformfieldValue field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetSEzsignformfieldValue() string {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldValue
}

// GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetSEzsignformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		return nil, false
	}
	return o.SEzsignformfieldValue, true
}

// HasSEzsignformfieldValue returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasSEzsignformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldValue) {
		return true
	}

	return false
}

// SetSEzsignformfieldValue gets a reference to the given string and assigns it to the SEzsignformfieldValue field.
func (o *EzsignformfieldRequest) SetSEzsignformfieldValue(v string) {
	o.SEzsignformfieldValue = &v
}

// GetIEzsignformfieldX returns the IEzsignformfieldX field value
func (o *EzsignformfieldRequest) GetIEzsignformfieldX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldX
}

// GetIEzsignformfieldXOk returns a tuple with the IEzsignformfieldX field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetIEzsignformfieldXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldX, true
}

// SetIEzsignformfieldX sets field value
func (o *EzsignformfieldRequest) SetIEzsignformfieldX(v int32) {
	o.IEzsignformfieldX = v
}

// GetIEzsignformfieldY returns the IEzsignformfieldY field value
func (o *EzsignformfieldRequest) GetIEzsignformfieldY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldY
}

// GetIEzsignformfieldYOk returns a tuple with the IEzsignformfieldY field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetIEzsignformfieldYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldY, true
}

// SetIEzsignformfieldY sets field value
func (o *EzsignformfieldRequest) SetIEzsignformfieldY(v int32) {
	o.IEzsignformfieldY = v
}

// GetIEzsignformfieldWidth returns the IEzsignformfieldWidth field value
func (o *EzsignformfieldRequest) GetIEzsignformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldWidth
}

// GetIEzsignformfieldWidthOk returns a tuple with the IEzsignformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetIEzsignformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldWidth, true
}

// SetIEzsignformfieldWidth sets field value
func (o *EzsignformfieldRequest) SetIEzsignformfieldWidth(v int32) {
	o.IEzsignformfieldWidth = v
}

// GetIEzsignformfieldHeight returns the IEzsignformfieldHeight field value
func (o *EzsignformfieldRequest) GetIEzsignformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldHeight
}

// GetIEzsignformfieldHeightOk returns a tuple with the IEzsignformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetIEzsignformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldHeight, true
}

// SetIEzsignformfieldHeight sets field value
func (o *EzsignformfieldRequest) SetIEzsignformfieldHeight(v int32) {
	o.IEzsignformfieldHeight = v
}

// GetBEzsignformfieldAutocomplete returns the BEzsignformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetBEzsignformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldAutocomplete
}

// GetBEzsignformfieldAutocompleteOk returns a tuple with the BEzsignformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetBEzsignformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsignformfieldAutocomplete, true
}

// HasBEzsignformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasBEzsignformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsignformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsignformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsignformfieldAutocomplete field.
func (o *EzsignformfieldRequest) SetBEzsignformfieldAutocomplete(v bool) {
	o.BEzsignformfieldAutocomplete = &v
}

// GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetBEzsignformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldSelected
}

// GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetBEzsignformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		return nil, false
	}
	return o.BEzsignformfieldSelected, true
}

// HasBEzsignformfieldSelected returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasBEzsignformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsignformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsignformfieldSelected gets a reference to the given bool and assigns it to the BEzsignformfieldSelected field.
func (o *EzsignformfieldRequest) SetBEzsignformfieldSelected(v bool) {
	o.BEzsignformfieldSelected = &v
}

// GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetSEzsignformfieldEnteredvalue() string {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldEnteredvalue
}

// GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetSEzsignformfieldEnteredvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		return nil, false
	}
	return o.SEzsignformfieldEnteredvalue, true
}

// HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasSEzsignformfieldEnteredvalue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldEnteredvalue) {
		return true
	}

	return false
}

// SetSEzsignformfieldEnteredvalue gets a reference to the given string and assigns it to the SEzsignformfieldEnteredvalue field.
func (o *EzsignformfieldRequest) SetSEzsignformfieldEnteredvalue(v string) {
	o.SEzsignformfieldEnteredvalue = &v
}

// GetEEzsignformfieldDependencyrequirement returns the EEzsignformfieldDependencyrequirement field value if set, zero value otherwise.
func (o *EzsignformfieldRequest) GetEEzsignformfieldDependencyrequirement() FieldEEzsignformfieldDependencyrequirement {
	if o == nil || IsNil(o.EEzsignformfieldDependencyrequirement) {
		var ret FieldEEzsignformfieldDependencyrequirement
		return ret
	}
	return *o.EEzsignformfieldDependencyrequirement
}

// GetEEzsignformfieldDependencyrequirementOk returns a tuple with the EEzsignformfieldDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldRequest) GetEEzsignformfieldDependencyrequirementOk() (*FieldEEzsignformfieldDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsignformfieldDependencyrequirement) {
		return nil, false
	}
	return o.EEzsignformfieldDependencyrequirement, true
}

// HasEEzsignformfieldDependencyrequirement returns a boolean if a field has been set.
func (o *EzsignformfieldRequest) HasEEzsignformfieldDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsignformfieldDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsignformfieldDependencyrequirement gets a reference to the given FieldEEzsignformfieldDependencyrequirement and assigns it to the EEzsignformfieldDependencyrequirement field.
func (o *EzsignformfieldRequest) SetEEzsignformfieldDependencyrequirement(v FieldEEzsignformfieldDependencyrequirement) {
	o.EEzsignformfieldDependencyrequirement = &v
}

func (o EzsignformfieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignformfieldID) {
		toSerialize["pkiEzsignformfieldID"] = o.PkiEzsignformfieldID
	}
	toSerialize["iEzsignpagePagenumber"] = o.IEzsignpagePagenumber
	toSerialize["sEzsignformfieldLabel"] = o.SEzsignformfieldLabel
	if !IsNil(o.SEzsignformfieldValue) {
		toSerialize["sEzsignformfieldValue"] = o.SEzsignformfieldValue
	}
	toSerialize["iEzsignformfieldX"] = o.IEzsignformfieldX
	toSerialize["iEzsignformfieldY"] = o.IEzsignformfieldY
	toSerialize["iEzsignformfieldWidth"] = o.IEzsignformfieldWidth
	toSerialize["iEzsignformfieldHeight"] = o.IEzsignformfieldHeight
	if !IsNil(o.BEzsignformfieldAutocomplete) {
		toSerialize["bEzsignformfieldAutocomplete"] = o.BEzsignformfieldAutocomplete
	}
	if !IsNil(o.BEzsignformfieldSelected) {
		toSerialize["bEzsignformfieldSelected"] = o.BEzsignformfieldSelected
	}
	if !IsNil(o.SEzsignformfieldEnteredvalue) {
		toSerialize["sEzsignformfieldEnteredvalue"] = o.SEzsignformfieldEnteredvalue
	}
	if !IsNil(o.EEzsignformfieldDependencyrequirement) {
		toSerialize["eEzsignformfieldDependencyrequirement"] = o.EEzsignformfieldDependencyrequirement
	}
	return toSerialize, nil
}

func (o *EzsignformfieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iEzsignpagePagenumber",
		"sEzsignformfieldLabel",
		"iEzsignformfieldX",
		"iEzsignformfieldY",
		"iEzsignformfieldWidth",
		"iEzsignformfieldHeight",
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

	varEzsignformfieldRequest := _EzsignformfieldRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldRequest)

	if err != nil {
		return err
	}

	*o = EzsignformfieldRequest(varEzsignformfieldRequest)

	return err
}

type NullableEzsignformfieldRequest struct {
	value *EzsignformfieldRequest
	isSet bool
}

func (v NullableEzsignformfieldRequest) Get() *EzsignformfieldRequest {
	return v.value
}

func (v *NullableEzsignformfieldRequest) Set(val *EzsignformfieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldRequest(val *EzsignformfieldRequest) *NullableEzsignformfieldRequest {
	return &NullableEzsignformfieldRequest{value: val, isSet: true}
}

func (v NullableEzsignformfieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


