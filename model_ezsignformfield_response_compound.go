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

// checks if the EzsignformfieldResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldResponseCompound{}

// EzsignformfieldResponseCompound An Ezsignformfield Object and children to create a complete structure
type EzsignformfieldResponseCompound struct {
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
	EEzsignformfieldDependencyrequirement *FieldEEzsignformfieldDependencyrequirement `json:"eEzsignformfieldDependencyrequirement,omitempty"`
	AObjEzsignelementdependency []EzsignelementdependencyResponseCompound `json:"a_objEzsignelementdependency,omitempty"`
}

type _EzsignformfieldResponseCompound EzsignformfieldResponseCompound

// NewEzsignformfieldResponseCompound instantiates a new EzsignformfieldResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldResponseCompound(pkiEzsignformfieldID int32, iEzsignpagePagenumber int32, sEzsignformfieldLabel string, iEzsignformfieldX int32, iEzsignformfieldY int32, iEzsignformfieldWidth int32, iEzsignformfieldHeight int32) *EzsignformfieldResponseCompound {
	this := EzsignformfieldResponseCompound{}
	this.PkiEzsignformfieldID = pkiEzsignformfieldID
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.IEzsignformfieldX = iEzsignformfieldX
	this.IEzsignformfieldY = iEzsignformfieldY
	this.IEzsignformfieldWidth = iEzsignformfieldWidth
	this.IEzsignformfieldHeight = iEzsignformfieldHeight
	return &this
}

// NewEzsignformfieldResponseCompoundWithDefaults instantiates a new EzsignformfieldResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldResponseCompoundWithDefaults() *EzsignformfieldResponseCompound {
	this := EzsignformfieldResponseCompound{}
	return &this
}

// GetPkiEzsignformfieldID returns the PkiEzsignformfieldID field value
func (o *EzsignformfieldResponseCompound) GetPkiEzsignformfieldID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignformfieldID
}

// GetPkiEzsignformfieldIDOk returns a tuple with the PkiEzsignformfieldID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetPkiEzsignformfieldIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignformfieldID, true
}

// SetPkiEzsignformfieldID sets field value
func (o *EzsignformfieldResponseCompound) SetPkiEzsignformfieldID(v int32) {
	o.PkiEzsignformfieldID = v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignformfieldResponseCompound) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignformfieldResponseCompound) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetSEzsignformfieldValue returns the SEzsignformfieldValue field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldValue() string {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldValue
}

// GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldValue) {
		return nil, false
	}
	return o.SEzsignformfieldValue, true
}

// HasSEzsignformfieldValue returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasSEzsignformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldValue) {
		return true
	}

	return false
}

// SetSEzsignformfieldValue gets a reference to the given string and assigns it to the SEzsignformfieldValue field.
func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldValue(v string) {
	o.SEzsignformfieldValue = &v
}

// GetIEzsignformfieldX returns the IEzsignformfieldX field value
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldX
}

// GetIEzsignformfieldXOk returns a tuple with the IEzsignformfieldX field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldX, true
}

// SetIEzsignformfieldX sets field value
func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldX(v int32) {
	o.IEzsignformfieldX = v
}

// GetIEzsignformfieldY returns the IEzsignformfieldY field value
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldY
}

// GetIEzsignformfieldYOk returns a tuple with the IEzsignformfieldY field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldY, true
}

// SetIEzsignformfieldY sets field value
func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldY(v int32) {
	o.IEzsignformfieldY = v
}

// GetIEzsignformfieldWidth returns the IEzsignformfieldWidth field value
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldWidth
}

// GetIEzsignformfieldWidthOk returns a tuple with the IEzsignformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldWidth, true
}

// SetIEzsignformfieldWidth sets field value
func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldWidth(v int32) {
	o.IEzsignformfieldWidth = v
}

// GetIEzsignformfieldHeight returns the IEzsignformfieldHeight field value
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignformfieldHeight
}

// GetIEzsignformfieldHeightOk returns a tuple with the IEzsignformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetIEzsignformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignformfieldHeight, true
}

// SetIEzsignformfieldHeight sets field value
func (o *EzsignformfieldResponseCompound) SetIEzsignformfieldHeight(v int32) {
	o.IEzsignformfieldHeight = v
}

// GetBEzsignformfieldAutocomplete returns the BEzsignformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldAutocomplete
}

// GetBEzsignformfieldAutocompleteOk returns a tuple with the BEzsignformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsignformfieldAutocomplete, true
}

// HasBEzsignformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasBEzsignformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsignformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsignformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsignformfieldAutocomplete field.
func (o *EzsignformfieldResponseCompound) SetBEzsignformfieldAutocomplete(v bool) {
	o.BEzsignformfieldAutocomplete = &v
}

// GetBEzsignformfieldSelected returns the BEzsignformfieldSelected field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsignformfieldSelected
}

// GetBEzsignformfieldSelectedOk returns a tuple with the BEzsignformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetBEzsignformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignformfieldSelected) {
		return nil, false
	}
	return o.BEzsignformfieldSelected, true
}

// HasBEzsignformfieldSelected returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasBEzsignformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsignformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsignformfieldSelected gets a reference to the given bool and assigns it to the BEzsignformfieldSelected field.
func (o *EzsignformfieldResponseCompound) SetBEzsignformfieldSelected(v bool) {
	o.BEzsignformfieldSelected = &v
}

// GetSEzsignformfieldEnteredvalue returns the SEzsignformfieldEnteredvalue field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldEnteredvalue() string {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		var ret string
		return ret
	}
	return *o.SEzsignformfieldEnteredvalue
}

// GetSEzsignformfieldEnteredvalueOk returns a tuple with the SEzsignformfieldEnteredvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetSEzsignformfieldEnteredvalueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignformfieldEnteredvalue) {
		return nil, false
	}
	return o.SEzsignformfieldEnteredvalue, true
}

// HasSEzsignformfieldEnteredvalue returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasSEzsignformfieldEnteredvalue() bool {
	if o != nil && !IsNil(o.SEzsignformfieldEnteredvalue) {
		return true
	}

	return false
}

// SetSEzsignformfieldEnteredvalue gets a reference to the given string and assigns it to the SEzsignformfieldEnteredvalue field.
func (o *EzsignformfieldResponseCompound) SetSEzsignformfieldEnteredvalue(v string) {
	o.SEzsignformfieldEnteredvalue = &v
}

// GetEEzsignformfieldDependencyrequirement returns the EEzsignformfieldDependencyrequirement field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetEEzsignformfieldDependencyrequirement() FieldEEzsignformfieldDependencyrequirement {
	if o == nil || IsNil(o.EEzsignformfieldDependencyrequirement) {
		var ret FieldEEzsignformfieldDependencyrequirement
		return ret
	}
	return *o.EEzsignformfieldDependencyrequirement
}

// GetEEzsignformfieldDependencyrequirementOk returns a tuple with the EEzsignformfieldDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetEEzsignformfieldDependencyrequirementOk() (*FieldEEzsignformfieldDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsignformfieldDependencyrequirement) {
		return nil, false
	}
	return o.EEzsignformfieldDependencyrequirement, true
}

// HasEEzsignformfieldDependencyrequirement returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasEEzsignformfieldDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsignformfieldDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsignformfieldDependencyrequirement gets a reference to the given FieldEEzsignformfieldDependencyrequirement and assigns it to the EEzsignformfieldDependencyrequirement field.
func (o *EzsignformfieldResponseCompound) SetEEzsignformfieldDependencyrequirement(v FieldEEzsignformfieldDependencyrequirement) {
	o.EEzsignformfieldDependencyrequirement = &v
}

// GetAObjEzsignelementdependency returns the AObjEzsignelementdependency field value if set, zero value otherwise.
func (o *EzsignformfieldResponseCompound) GetAObjEzsignelementdependency() []EzsignelementdependencyResponseCompound {
	if o == nil || IsNil(o.AObjEzsignelementdependency) {
		var ret []EzsignelementdependencyResponseCompound
		return ret
	}
	return o.AObjEzsignelementdependency
}

// GetAObjEzsignelementdependencyOk returns a tuple with the AObjEzsignelementdependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponseCompound) GetAObjEzsignelementdependencyOk() ([]EzsignelementdependencyResponseCompound, bool) {
	if o == nil || IsNil(o.AObjEzsignelementdependency) {
		return nil, false
	}
	return o.AObjEzsignelementdependency, true
}

// HasAObjEzsignelementdependency returns a boolean if a field has been set.
func (o *EzsignformfieldResponseCompound) HasAObjEzsignelementdependency() bool {
	if o != nil && !IsNil(o.AObjEzsignelementdependency) {
		return true
	}

	return false
}

// SetAObjEzsignelementdependency gets a reference to the given []EzsignelementdependencyResponseCompound and assigns it to the AObjEzsignelementdependency field.
func (o *EzsignformfieldResponseCompound) SetAObjEzsignelementdependency(v []EzsignelementdependencyResponseCompound) {
	o.AObjEzsignelementdependency = v
}

func (o EzsignformfieldResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldResponseCompound) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EEzsignformfieldDependencyrequirement) {
		toSerialize["eEzsignformfieldDependencyrequirement"] = o.EEzsignformfieldDependencyrequirement
	}
	if !IsNil(o.AObjEzsignelementdependency) {
		toSerialize["a_objEzsignelementdependency"] = o.AObjEzsignelementdependency
	}
	return toSerialize, nil
}

func (o *EzsignformfieldResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignformfieldID",
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

	varEzsignformfieldResponseCompound := _EzsignformfieldResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignformfieldResponseCompound(varEzsignformfieldResponseCompound)

	return err
}

type NullableEzsignformfieldResponseCompound struct {
	value *EzsignformfieldResponseCompound
	isSet bool
}

func (v NullableEzsignformfieldResponseCompound) Get() *EzsignformfieldResponseCompound {
	return v.value
}

func (v *NullableEzsignformfieldResponseCompound) Set(val *EzsignformfieldResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldResponseCompound(val *EzsignformfieldResponseCompound) *NullableEzsignformfieldResponseCompound {
	return &NullableEzsignformfieldResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignformfieldResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


