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

// checks if the EzsignformfieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldResponse{}

// EzsignformfieldResponse An Ezsignformfield Object
type EzsignformfieldResponse struct {
	// The unique ID of the Ezsignformfield
	PkiEzsignformfieldID int32 `json:"pkiEzsignformfieldID"`
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
}

// NewEzsignformfieldResponse instantiates a new EzsignformfieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldResponse(pkiEzsignformfieldID int32, iEzsignpagePagenumber int32, sEzsignformfieldLabel string, iEzsignformfieldX int32, iEzsignformfieldY int32, iEzsignformfieldWidth int32, iEzsignformfieldHeight int32) *EzsignformfieldResponse {
	this := EzsignformfieldResponse{}
	this.PkiEzsignformfieldID = pkiEzsignformfieldID
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.IEzsignformfieldX = iEzsignformfieldX
	this.IEzsignformfieldY = iEzsignformfieldY
	this.IEzsignformfieldWidth = iEzsignformfieldWidth
	this.IEzsignformfieldHeight = iEzsignformfieldHeight
	return &this
}

// NewEzsignformfieldResponseWithDefaults instantiates a new EzsignformfieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldResponseWithDefaults() *EzsignformfieldResponse {
	this := EzsignformfieldResponse{}
	return &this
}

// GetPkiEzsignformfieldID returns the PkiEzsignformfieldID field value
func (o *EzsignformfieldResponse) GetPkiEzsignformfieldID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignformfieldID
}

// GetPkiEzsignformfieldIDOk returns a tuple with the PkiEzsignformfieldID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetPkiEzsignformfieldIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignformfieldID, true
}

// SetPkiEzsignformfieldID sets field value
func (o *EzsignformfieldResponse) SetPkiEzsignformfieldID(v int32) {
	o.PkiEzsignformfieldID = v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignformfieldResponse) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignformfieldResponse) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *EzsignformfieldResponse) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *EzsignformfieldResponse) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetSEzsignformfieldValue returns the SEzsignformfieldValue field value if set, zero value otherwise.
func (o *EzsignformfieldResponse) GetSEzsignformfieldValue() string {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldValue
}

// GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetSEzsignformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		return nil, false
	}
	return o.SEzsignformfieldValue, true
}

// HasSEzsignformfieldValue returns a boolean if a field has been set.
func (o *EzsignformfieldResponse) HasSEzsignformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldValue) {
		return true
	}

	return false
}

// SetSEzsignformfieldValue gets a reference to the given string and assigns it to the SEzsignformfieldValue field.
func (o *EzsignformfieldResponse) SetSEzsignformfieldValue(v string) {
	o.SEzsignformfieldValue = &v
}

// GetIEzsignformfieldX returns the IEzsignformfieldX field value
func (o *EzsignformfieldResponse) GetIEzsignformfieldX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldX
}

// GetIEzsignformfieldXOk returns a tuple with the IEzsignformfieldX field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetIEzsignformfieldXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldX, true
}

// SetIEzsignformfieldX sets field value
func (o *EzsignformfieldResponse) SetIEzsignformfieldX(v int32) {
	o.IEzsignformfieldX = v
}

// GetIEzsignformfieldY returns the IEzsignformfieldY field value
func (o *EzsignformfieldResponse) GetIEzsignformfieldY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldY
}

// GetIEzsignformfieldYOk returns a tuple with the IEzsignformfieldY field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetIEzsignformfieldYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldY, true
}

// SetIEzsignformfieldY sets field value
func (o *EzsignformfieldResponse) SetIEzsignformfieldY(v int32) {
	o.IEzsignformfieldY = v
}

// GetIEzsignformfieldWidth returns the IEzsignformfieldWidth field value
func (o *EzsignformfieldResponse) GetIEzsignformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldWidth
}

// GetIEzsignformfieldWidthOk returns a tuple with the IEzsignformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetIEzsignformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldWidth, true
}

// SetIEzsignformfieldWidth sets field value
func (o *EzsignformfieldResponse) SetIEzsignformfieldWidth(v int32) {
	o.IEzsignformfieldWidth = v
}

// GetIEzsignformfieldHeight returns the IEzsignformfieldHeight field value
func (o *EzsignformfieldResponse) GetIEzsignformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldHeight
}

// GetIEzsignformfieldHeightOk returns a tuple with the IEzsignformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetIEzsignformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldHeight, true
}

// SetIEzsignformfieldHeight sets field value
func (o *EzsignformfieldResponse) SetIEzsignformfieldHeight(v int32) {
	o.IEzsignformfieldHeight = v
}

// GetBEzsignformfieldAutocomplete returns the BEzsignformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsignformfieldResponse) GetBEzsignformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldAutocomplete
}

// GetBEzsignformfieldAutocompleteOk returns a tuple with the BEzsignformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetBEzsignformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsignformfieldAutocomplete, true
}

// HasBEzsignformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsignformfieldResponse) HasBEzsignformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsignformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsignformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsignformfieldAutocomplete field.
func (o *EzsignformfieldResponse) SetBEzsignformfieldAutocomplete(v bool) {
	o.BEzsignformfieldAutocomplete = &v
}

// GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field value if set, zero value otherwise.
func (o *EzsignformfieldResponse) GetBEzsignformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldSelected
}

// GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetBEzsignformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		return nil, false
	}
	return o.BEzsignformfieldSelected, true
}

// HasBEzsignformfieldSelected returns a boolean if a field has been set.
func (o *EzsignformfieldResponse) HasBEzsignformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsignformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsignformfieldSelected gets a reference to the given bool and assigns it to the BEzsignformfieldSelected field.
func (o *EzsignformfieldResponse) SetBEzsignformfieldSelected(v bool) {
	o.BEzsignformfieldSelected = &v
}

// GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field value if set, zero value otherwise.
func (o *EzsignformfieldResponse) GetSEzsignformfieldEnteredvalue() string {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldEnteredvalue
}

// GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetSEzsignformfieldEnteredvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		return nil, false
	}
	return o.SEzsignformfieldEnteredvalue, true
}

// HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.
func (o *EzsignformfieldResponse) HasSEzsignformfieldEnteredvalue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldEnteredvalue) {
		return true
	}

	return false
}

// SetSEzsignformfieldEnteredvalue gets a reference to the given string and assigns it to the SEzsignformfieldEnteredvalue field.
func (o *EzsignformfieldResponse) SetSEzsignformfieldEnteredvalue(v string) {
	o.SEzsignformfieldEnteredvalue = &v
}

func (o EzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignformfieldID"] = o.PkiEzsignformfieldID
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
	return toSerialize, nil
}

type NullableEzsignformfieldResponse struct {
	value *EzsignformfieldResponse
	isSet bool
}

func (v NullableEzsignformfieldResponse) Get() *EzsignformfieldResponse {
	return v.value
}

func (v *NullableEzsignformfieldResponse) Set(val *EzsignformfieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldResponse(val *EzsignformfieldResponse) *NullableEzsignformfieldResponse {
	return &NullableEzsignformfieldResponse{value: val, isSet: true}
}

func (v NullableEzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


