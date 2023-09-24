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

// checks if the EzsigntemplatesignaturecustomdateRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignaturecustomdateRequestCompound{}

// EzsigntemplatesignaturecustomdateRequestCompound An Ezsigntemplatesignaturecustomdate Object and children to create a complete structure
type EzsigntemplatesignaturecustomdateRequestCompound struct {
	// The unique ID of the Ezsigntemplatesignaturecustomdate
	PkiEzsigntemplatesignaturecustomdateID *int32 `json:"pkiEzsigntemplatesignaturecustomdateID,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsigntemplatesignaturecustomdateX int32 `json:"iEzsigntemplatesignaturecustomdateX"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsigntemplatesignaturecustomdateY int32 `json:"iEzsigntemplatesignaturecustomdateY"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsigntemplatesignaturecustomdateFormat string `json:"sEzsigntemplatesignaturecustomdateFormat"`
}

// NewEzsigntemplatesignaturecustomdateRequestCompound instantiates a new EzsigntemplatesignaturecustomdateRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignaturecustomdateRequestCompound(iEzsigntemplatesignaturecustomdateX int32, iEzsigntemplatesignaturecustomdateY int32, sEzsigntemplatesignaturecustomdateFormat string) *EzsigntemplatesignaturecustomdateRequestCompound {
	this := EzsigntemplatesignaturecustomdateRequestCompound{}
	this.IEzsigntemplatesignaturecustomdateX = iEzsigntemplatesignaturecustomdateX
	this.IEzsigntemplatesignaturecustomdateY = iEzsigntemplatesignaturecustomdateY
	this.SEzsigntemplatesignaturecustomdateFormat = sEzsigntemplatesignaturecustomdateFormat
	return &this
}

// NewEzsigntemplatesignaturecustomdateRequestCompoundWithDefaults instantiates a new EzsigntemplatesignaturecustomdateRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignaturecustomdateRequestCompoundWithDefaults() *EzsigntemplatesignaturecustomdateRequestCompound {
	this := EzsigntemplatesignaturecustomdateRequestCompound{}
	return &this
}

// GetPkiEzsigntemplatesignaturecustomdateID returns the PkiEzsigntemplatesignaturecustomdateID field value if set, zero value otherwise.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetPkiEzsigntemplatesignaturecustomdateID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatesignaturecustomdateID
}

// GetPkiEzsigntemplatesignaturecustomdateIDOk returns a tuple with the PkiEzsigntemplatesignaturecustomdateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetPkiEzsigntemplatesignaturecustomdateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		return nil, false
	}
	return o.PkiEzsigntemplatesignaturecustomdateID, true
}

// HasPkiEzsigntemplatesignaturecustomdateID returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) HasPkiEzsigntemplatesignaturecustomdateID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatesignaturecustomdateID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatesignaturecustomdateID field.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetPkiEzsigntemplatesignaturecustomdateID(v int32) {
	o.PkiEzsigntemplatesignaturecustomdateID = &v
}

// GetIEzsigntemplatesignaturecustomdateX returns the IEzsigntemplatesignaturecustomdateX field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignaturecustomdateX
}

// GetIEzsigntemplatesignaturecustomdateXOk returns a tuple with the IEzsigntemplatesignaturecustomdateX field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignaturecustomdateX, true
}

// SetIEzsigntemplatesignaturecustomdateX sets field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetIEzsigntemplatesignaturecustomdateX(v int32) {
	o.IEzsigntemplatesignaturecustomdateX = v
}

// GetIEzsigntemplatesignaturecustomdateY returns the IEzsigntemplatesignaturecustomdateY field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignaturecustomdateY
}

// GetIEzsigntemplatesignaturecustomdateYOk returns a tuple with the IEzsigntemplatesignaturecustomdateY field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetIEzsigntemplatesignaturecustomdateYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignaturecustomdateY, true
}

// SetIEzsigntemplatesignaturecustomdateY sets field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetIEzsigntemplatesignaturecustomdateY(v int32) {
	o.IEzsigntemplatesignaturecustomdateY = v
}

// GetSEzsigntemplatesignaturecustomdateFormat returns the SEzsigntemplatesignaturecustomdateFormat field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetSEzsigntemplatesignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignaturecustomdateFormat
}

// GetSEzsigntemplatesignaturecustomdateFormatOk returns a tuple with the SEzsigntemplatesignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequestCompound) GetSEzsigntemplatesignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignaturecustomdateFormat, true
}

// SetSEzsigntemplatesignaturecustomdateFormat sets field value
func (o *EzsigntemplatesignaturecustomdateRequestCompound) SetSEzsigntemplatesignaturecustomdateFormat(v string) {
	o.SEzsigntemplatesignaturecustomdateFormat = v
}

func (o EzsigntemplatesignaturecustomdateRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignaturecustomdateRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		toSerialize["pkiEzsigntemplatesignaturecustomdateID"] = o.PkiEzsigntemplatesignaturecustomdateID
	}
	toSerialize["iEzsigntemplatesignaturecustomdateX"] = o.IEzsigntemplatesignaturecustomdateX
	toSerialize["iEzsigntemplatesignaturecustomdateY"] = o.IEzsigntemplatesignaturecustomdateY
	toSerialize["sEzsigntemplatesignaturecustomdateFormat"] = o.SEzsigntemplatesignaturecustomdateFormat
	return toSerialize, nil
}

type NullableEzsigntemplatesignaturecustomdateRequestCompound struct {
	value *EzsigntemplatesignaturecustomdateRequestCompound
	isSet bool
}

func (v NullableEzsigntemplatesignaturecustomdateRequestCompound) Get() *EzsigntemplatesignaturecustomdateRequestCompound {
	return v.value
}

func (v *NullableEzsigntemplatesignaturecustomdateRequestCompound) Set(val *EzsigntemplatesignaturecustomdateRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignaturecustomdateRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignaturecustomdateRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignaturecustomdateRequestCompound(val *EzsigntemplatesignaturecustomdateRequestCompound) *NullableEzsigntemplatesignaturecustomdateRequestCompound {
	return &NullableEzsigntemplatesignaturecustomdateRequestCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignaturecustomdateRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignaturecustomdateRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

