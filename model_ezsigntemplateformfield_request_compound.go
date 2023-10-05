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

// checks if the EzsigntemplateformfieldRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateformfieldRequestCompound{}

// EzsigntemplateformfieldRequestCompound An Ezsigntemplateformfield Object and children to create a complete structure
type EzsigntemplateformfieldRequestCompound struct {
	// The unique ID of the Ezsigntemplateformfield
	PkiEzsigntemplateformfieldID *int32 `json:"pkiEzsigntemplateformfieldID,omitempty"`
	// The page number in the Ezsigntemplatedocument
	IEzsigntemplatedocumentpagePagenumber int32 `json:"iEzsigntemplatedocumentpagePagenumber"`
	// The Label for the Ezsigntemplateformfield
	SEzsigntemplateformfieldLabel string `json:"sEzsigntemplateformfieldLabel"`
	// The value for the Ezsigntemplateformfield
	SEzsigntemplateformfieldValue *string `json:"sEzsigntemplateformfieldValue,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplateformfieldX int32 `json:"iEzsigntemplateformfieldX"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplateformfield on the Ezsigntemplatepage.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplateformfield 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplateformfieldY int32 `json:"iEzsigntemplateformfieldY"`
	// The Width of the Ezsigntemplateformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsigntemplateformfieldgroupType.  | eEzsigntemplateformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22-65535     | | Radio                     | 22           | | Text                      | 22-65535     | | Textarea                  | 22-65535     |
	IEzsigntemplateformfieldWidth int32 `json:"iEzsigntemplateformfieldWidth"`
	// The Height of the Ezsigntemplateformfield in pixels calculated at 100 DPI  The allowed values are varying based on the eEzsigntemplateformfieldgroupType.  | eEzsigntemplateformfieldgroupType | Valid values | | ------------------------- | ------------ | | Checkbox                  | 22           | | Dropdown                  | 22           | | Radio                     | 22           | | Text                      | 22           | | Textarea                  | 22-65535     | 
	IEzsigntemplateformfieldHeight int32 `json:"iEzsigntemplateformfieldHeight"`
	// Whether the Ezsigntemplateformfield allows the use of the autocomplete of the browser.  This can only be set if eEzsigntemplateformfieldgroupType is **Text**
	BEzsigntemplateformfieldAutocomplete *bool `json:"bEzsigntemplateformfieldAutocomplete,omitempty"`
	// Whether the Ezsigntemplateformfield is selected or not by default.  This can only be set if eEzsigntemplateformfieldgroupType is **Checkbox** or **Radio**
	BEzsigntemplateformfieldSelected *bool `json:"bEzsigntemplateformfieldSelected,omitempty"`
	EEzsigntemplateformfieldDependencyrequirement *FieldEEzsigntemplateformfieldDependencyrequirement `json:"eEzsigntemplateformfieldDependencyrequirement,omitempty"`
	AObjEzsigntemplateelementdependency []EzsigntemplateelementdependencyRequestCompound `json:"a_objEzsigntemplateelementdependency,omitempty"`
}

// NewEzsigntemplateformfieldRequestCompound instantiates a new EzsigntemplateformfieldRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateformfieldRequestCompound(iEzsigntemplatedocumentpagePagenumber int32, sEzsigntemplateformfieldLabel string, iEzsigntemplateformfieldX int32, iEzsigntemplateformfieldY int32, iEzsigntemplateformfieldWidth int32, iEzsigntemplateformfieldHeight int32) *EzsigntemplateformfieldRequestCompound {
	this := EzsigntemplateformfieldRequestCompound{}
	this.IEzsigntemplatedocumentpagePagenumber = iEzsigntemplatedocumentpagePagenumber
	this.SEzsigntemplateformfieldLabel = sEzsigntemplateformfieldLabel
	this.IEzsigntemplateformfieldX = iEzsigntemplateformfieldX
	this.IEzsigntemplateformfieldY = iEzsigntemplateformfieldY
	this.IEzsigntemplateformfieldWidth = iEzsigntemplateformfieldWidth
	this.IEzsigntemplateformfieldHeight = iEzsigntemplateformfieldHeight
	return &this
}

// NewEzsigntemplateformfieldRequestCompoundWithDefaults instantiates a new EzsigntemplateformfieldRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateformfieldRequestCompoundWithDefaults() *EzsigntemplateformfieldRequestCompound {
	this := EzsigntemplateformfieldRequestCompound{}
	return &this
}

// GetPkiEzsigntemplateformfieldID returns the PkiEzsigntemplateformfieldID field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplateformfieldID
}

// GetPkiEzsigntemplateformfieldIDOk returns a tuple with the PkiEzsigntemplateformfieldID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetPkiEzsigntemplateformfieldIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplateformfieldID) {
		return nil, false
	}
	return o.PkiEzsigntemplateformfieldID, true
}

// HasPkiEzsigntemplateformfieldID returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasPkiEzsigntemplateformfieldID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplateformfieldID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplateformfieldID gets a reference to the given int32 and assigns it to the PkiEzsigntemplateformfieldID field.
func (o *EzsigntemplateformfieldRequestCompound) SetPkiEzsigntemplateformfieldID(v int32) {
	o.PkiEzsigntemplateformfieldID = &v
}

// GetIEzsigntemplatedocumentpagePagenumber returns the IEzsigntemplatedocumentpagePagenumber field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatedocumentpagePagenumber
}

// GetIEzsigntemplatedocumentpagePagenumberOk returns a tuple with the IEzsigntemplatedocumentpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplatedocumentpagePagenumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatedocumentpagePagenumber, true
}

// SetIEzsigntemplatedocumentpagePagenumber sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplatedocumentpagePagenumber(v int32) {
	o.IEzsigntemplatedocumentpagePagenumber = v
}

// GetSEzsigntemplateformfieldLabel returns the SEzsigntemplateformfieldLabel field value
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateformfieldLabel
}

// GetSEzsigntemplateformfieldLabelOk returns a tuple with the SEzsigntemplateformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateformfieldLabel, true
}

// SetSEzsigntemplateformfieldLabel sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldLabel(v string) {
	o.SEzsigntemplateformfieldLabel = v
}

// GetSEzsigntemplateformfieldValue returns the SEzsigntemplateformfieldValue field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValue() string {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		var ret string
		return ret
	}
	return *o.SEzsigntemplateformfieldValue
}

// GetSEzsigntemplateformfieldValueOk returns a tuple with the SEzsigntemplateformfieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetSEzsigntemplateformfieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigntemplateformfieldValue) {
		return nil, false
	}
	return o.SEzsigntemplateformfieldValue, true
}

// HasSEzsigntemplateformfieldValue returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasSEzsigntemplateformfieldValue() bool {
	if o != nil && !IsNil(o.SEzsigntemplateformfieldValue) {
		return true
	}

	return false
}

// SetSEzsigntemplateformfieldValue gets a reference to the given string and assigns it to the SEzsigntemplateformfieldValue field.
func (o *EzsigntemplateformfieldRequestCompound) SetSEzsigntemplateformfieldValue(v string) {
	o.SEzsigntemplateformfieldValue = &v
}

// GetIEzsigntemplateformfieldX returns the IEzsigntemplateformfieldX field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldX
}

// GetIEzsigntemplateformfieldXOk returns a tuple with the IEzsigntemplateformfieldX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldX, true
}

// SetIEzsigntemplateformfieldX sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldX(v int32) {
	o.IEzsigntemplateformfieldX = v
}

// GetIEzsigntemplateformfieldY returns the IEzsigntemplateformfieldY field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldY
}

// GetIEzsigntemplateformfieldYOk returns a tuple with the IEzsigntemplateformfieldY field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldY, true
}

// SetIEzsigntemplateformfieldY sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldY(v int32) {
	o.IEzsigntemplateformfieldY = v
}

// GetIEzsigntemplateformfieldWidth returns the IEzsigntemplateformfieldWidth field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldWidth
}

// GetIEzsigntemplateformfieldWidthOk returns a tuple with the IEzsigntemplateformfieldWidth field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldWidth, true
}

// SetIEzsigntemplateformfieldWidth sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldWidth(v int32) {
	o.IEzsigntemplateformfieldWidth = v
}

// GetIEzsigntemplateformfieldHeight returns the IEzsigntemplateformfieldHeight field value
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplateformfieldHeight
}

// GetIEzsigntemplateformfieldHeightOk returns a tuple with the IEzsigntemplateformfieldHeight field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetIEzsigntemplateformfieldHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplateformfieldHeight, true
}

// SetIEzsigntemplateformfieldHeight sets field value
func (o *EzsigntemplateformfieldRequestCompound) SetIEzsigntemplateformfieldHeight(v int32) {
	o.IEzsigntemplateformfieldHeight = v
}

// GetBEzsigntemplateformfieldAutocomplete returns the BEzsigntemplateformfieldAutocomplete field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocomplete() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldAutocomplete
}

// GetBEzsigntemplateformfieldAutocompleteOk returns a tuple with the BEzsigntemplateformfieldAutocomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldAutocompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldAutocomplete, true
}

// HasBEzsigntemplateformfieldAutocomplete returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldAutocomplete() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldAutocomplete gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldAutocomplete field.
func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldAutocomplete(v bool) {
	o.BEzsigntemplateformfieldAutocomplete = &v
}

// GetBEzsigntemplateformfieldSelected returns the BEzsigntemplateformfieldSelected field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelected() bool {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplateformfieldSelected
}

// GetBEzsigntemplateformfieldSelectedOk returns a tuple with the BEzsigntemplateformfieldSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetBEzsigntemplateformfieldSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplateformfieldSelected) {
		return nil, false
	}
	return o.BEzsigntemplateformfieldSelected, true
}

// HasBEzsigntemplateformfieldSelected returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasBEzsigntemplateformfieldSelected() bool {
	if o != nil && !IsNil(o.BEzsigntemplateformfieldSelected) {
		return true
	}

	return false
}

// SetBEzsigntemplateformfieldSelected gets a reference to the given bool and assigns it to the BEzsigntemplateformfieldSelected field.
func (o *EzsigntemplateformfieldRequestCompound) SetBEzsigntemplateformfieldSelected(v bool) {
	o.BEzsigntemplateformfieldSelected = &v
}

// GetEEzsigntemplateformfieldDependencyrequirement returns the EEzsigntemplateformfieldDependencyrequirement field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirement() FieldEEzsigntemplateformfieldDependencyrequirement {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		var ret FieldEEzsigntemplateformfieldDependencyrequirement
		return ret
	}
	return *o.EEzsigntemplateformfieldDependencyrequirement
}

// GetEEzsigntemplateformfieldDependencyrequirementOk returns a tuple with the EEzsigntemplateformfieldDependencyrequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetEEzsigntemplateformfieldDependencyrequirementOk() (*FieldEEzsigntemplateformfieldDependencyrequirement, bool) {
	if o == nil || IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return nil, false
	}
	return o.EEzsigntemplateformfieldDependencyrequirement, true
}

// HasEEzsigntemplateformfieldDependencyrequirement returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasEEzsigntemplateformfieldDependencyrequirement() bool {
	if o != nil && !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		return true
	}

	return false
}

// SetEEzsigntemplateformfieldDependencyrequirement gets a reference to the given FieldEEzsigntemplateformfieldDependencyrequirement and assigns it to the EEzsigntemplateformfieldDependencyrequirement field.
func (o *EzsigntemplateformfieldRequestCompound) SetEEzsigntemplateformfieldDependencyrequirement(v FieldEEzsigntemplateformfieldDependencyrequirement) {
	o.EEzsigntemplateformfieldDependencyrequirement = &v
}

// GetAObjEzsigntemplateelementdependency returns the AObjEzsigntemplateelementdependency field value if set, zero value otherwise.
func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependency() []EzsigntemplateelementdependencyRequestCompound {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		var ret []EzsigntemplateelementdependencyRequestCompound
		return ret
	}
	return o.AObjEzsigntemplateelementdependency
}

// GetAObjEzsigntemplateelementdependencyOk returns a tuple with the AObjEzsigntemplateelementdependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateformfieldRequestCompound) GetAObjEzsigntemplateelementdependencyOk() ([]EzsigntemplateelementdependencyRequestCompound, bool) {
	if o == nil || IsNil(o.AObjEzsigntemplateelementdependency) {
		return nil, false
	}
	return o.AObjEzsigntemplateelementdependency, true
}

// HasAObjEzsigntemplateelementdependency returns a boolean if a field has been set.
func (o *EzsigntemplateformfieldRequestCompound) HasAObjEzsigntemplateelementdependency() bool {
	if o != nil && !IsNil(o.AObjEzsigntemplateelementdependency) {
		return true
	}

	return false
}

// SetAObjEzsigntemplateelementdependency gets a reference to the given []EzsigntemplateelementdependencyRequestCompound and assigns it to the AObjEzsigntemplateelementdependency field.
func (o *EzsigntemplateformfieldRequestCompound) SetAObjEzsigntemplateelementdependency(v []EzsigntemplateelementdependencyRequestCompound) {
	o.AObjEzsigntemplateelementdependency = v
}

func (o EzsigntemplateformfieldRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateformfieldRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplateformfieldID) {
		toSerialize["pkiEzsigntemplateformfieldID"] = o.PkiEzsigntemplateformfieldID
	}
	toSerialize["iEzsigntemplatedocumentpagePagenumber"] = o.IEzsigntemplatedocumentpagePagenumber
	toSerialize["sEzsigntemplateformfieldLabel"] = o.SEzsigntemplateformfieldLabel
	if !IsNil(o.SEzsigntemplateformfieldValue) {
		toSerialize["sEzsigntemplateformfieldValue"] = o.SEzsigntemplateformfieldValue
	}
	toSerialize["iEzsigntemplateformfieldX"] = o.IEzsigntemplateformfieldX
	toSerialize["iEzsigntemplateformfieldY"] = o.IEzsigntemplateformfieldY
	toSerialize["iEzsigntemplateformfieldWidth"] = o.IEzsigntemplateformfieldWidth
	toSerialize["iEzsigntemplateformfieldHeight"] = o.IEzsigntemplateformfieldHeight
	if !IsNil(o.BEzsigntemplateformfieldAutocomplete) {
		toSerialize["bEzsigntemplateformfieldAutocomplete"] = o.BEzsigntemplateformfieldAutocomplete
	}
	if !IsNil(o.BEzsigntemplateformfieldSelected) {
		toSerialize["bEzsigntemplateformfieldSelected"] = o.BEzsigntemplateformfieldSelected
	}
	if !IsNil(o.EEzsigntemplateformfieldDependencyrequirement) {
		toSerialize["eEzsigntemplateformfieldDependencyrequirement"] = o.EEzsigntemplateformfieldDependencyrequirement
	}
	if !IsNil(o.AObjEzsigntemplateelementdependency) {
		toSerialize["a_objEzsigntemplateelementdependency"] = o.AObjEzsigntemplateelementdependency
	}
	return toSerialize, nil
}

type NullableEzsigntemplateformfieldRequestCompound struct {
	value *EzsigntemplateformfieldRequestCompound
	isSet bool
}

func (v NullableEzsigntemplateformfieldRequestCompound) Get() *EzsigntemplateformfieldRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplateformfieldRequestCompound) Set(val *EzsigntemplateformfieldRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateformfieldRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateformfieldRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateformfieldRequestCompound(val *EzsigntemplateformfieldRequestCompound) *NullableEzsigntemplateformfieldRequestCompound {
	return &NullableEzsigntemplateformfieldRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateformfieldRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateformfieldRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


