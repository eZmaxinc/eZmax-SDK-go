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

// checks if the EzsigntemplatesignaturecustomdateResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignaturecustomdateResponseV2{}

// EzsigntemplatesignaturecustomdateResponseV2 An Ezsigntemplatesignaturecustomdate Object
type EzsigntemplatesignaturecustomdateResponseV2 struct {
	// The unique ID of the Ezsigntemplatesignaturecustomdate
	PkiEzsigntemplatesignaturecustomdateID int32 `json:"pkiEzsigntemplatesignaturecustomdateID"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left of the signature, you would use \"200\" for the X coordinate.
	IEzsigntemplatesignaturecustomdateOffsetx int32 `json:"iEzsigntemplatesignaturecustomdateOffsetx"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the top of the signature, you would use \"200\" for the Y coordinate.
	IEzsigntemplatesignaturecustomdateOffsety int32 `json:"iEzsigntemplatesignaturecustomdateOffsety"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsigntemplatesignaturecustomdateFormat string `json:"sEzsigntemplatesignaturecustomdateFormat"`
}

type _EzsigntemplatesignaturecustomdateResponseV2 EzsigntemplatesignaturecustomdateResponseV2

// NewEzsigntemplatesignaturecustomdateResponseV2 instantiates a new EzsigntemplatesignaturecustomdateResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignaturecustomdateResponseV2(pkiEzsigntemplatesignaturecustomdateID int32, iEzsigntemplatesignaturecustomdateOffsetx int32, iEzsigntemplatesignaturecustomdateOffsety int32, sEzsigntemplatesignaturecustomdateFormat string) *EzsigntemplatesignaturecustomdateResponseV2 {
	this := EzsigntemplatesignaturecustomdateResponseV2{}
	this.PkiEzsigntemplatesignaturecustomdateID = pkiEzsigntemplatesignaturecustomdateID
	this.IEzsigntemplatesignaturecustomdateOffsetx = iEzsigntemplatesignaturecustomdateOffsetx
	this.IEzsigntemplatesignaturecustomdateOffsety = iEzsigntemplatesignaturecustomdateOffsety
	this.SEzsigntemplatesignaturecustomdateFormat = sEzsigntemplatesignaturecustomdateFormat
	return &this
}

// NewEzsigntemplatesignaturecustomdateResponseV2WithDefaults instantiates a new EzsigntemplatesignaturecustomdateResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignaturecustomdateResponseV2WithDefaults() *EzsigntemplatesignaturecustomdateResponseV2 {
	this := EzsigntemplatesignaturecustomdateResponseV2{}
	return &this
}

// GetPkiEzsigntemplatesignaturecustomdateID returns the PkiEzsigntemplatesignaturecustomdateID field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetPkiEzsigntemplatesignaturecustomdateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatesignaturecustomdateID
}

// GetPkiEzsigntemplatesignaturecustomdateIDOk returns a tuple with the PkiEzsigntemplatesignaturecustomdateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetPkiEzsigntemplatesignaturecustomdateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatesignaturecustomdateID, true
}

// SetPkiEzsigntemplatesignaturecustomdateID sets field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) SetPkiEzsigntemplatesignaturecustomdateID(v int32) {
	o.PkiEzsigntemplatesignaturecustomdateID = v
}

// GetIEzsigntemplatesignaturecustomdateOffsetx returns the IEzsigntemplatesignaturecustomdateOffsetx field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetIEzsigntemplatesignaturecustomdateOffsetx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignaturecustomdateOffsetx
}

// GetIEzsigntemplatesignaturecustomdateOffsetxOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsetx field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetIEzsigntemplatesignaturecustomdateOffsetxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignaturecustomdateOffsetx, true
}

// SetIEzsigntemplatesignaturecustomdateOffsetx sets field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) SetIEzsigntemplatesignaturecustomdateOffsetx(v int32) {
	o.IEzsigntemplatesignaturecustomdateOffsetx = v
}

// GetIEzsigntemplatesignaturecustomdateOffsety returns the IEzsigntemplatesignaturecustomdateOffsety field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetIEzsigntemplatesignaturecustomdateOffsety() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatesignaturecustomdateOffsety
}

// GetIEzsigntemplatesignaturecustomdateOffsetyOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsety field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetIEzsigntemplatesignaturecustomdateOffsetyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatesignaturecustomdateOffsety, true
}

// SetIEzsigntemplatesignaturecustomdateOffsety sets field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) SetIEzsigntemplatesignaturecustomdateOffsety(v int32) {
	o.IEzsigntemplatesignaturecustomdateOffsety = v
}

// GetSEzsigntemplatesignaturecustomdateFormat returns the SEzsigntemplatesignaturecustomdateFormat field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetSEzsigntemplatesignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignaturecustomdateFormat
}

// GetSEzsigntemplatesignaturecustomdateFormatOk returns a tuple with the SEzsigntemplatesignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateResponseV2) GetSEzsigntemplatesignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignaturecustomdateFormat, true
}

// SetSEzsigntemplatesignaturecustomdateFormat sets field value
func (o *EzsigntemplatesignaturecustomdateResponseV2) SetSEzsigntemplatesignaturecustomdateFormat(v string) {
	o.SEzsigntemplatesignaturecustomdateFormat = v
}

func (o EzsigntemplatesignaturecustomdateResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignaturecustomdateResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatesignaturecustomdateID"] = o.PkiEzsigntemplatesignaturecustomdateID
	toSerialize["iEzsigntemplatesignaturecustomdateOffsetx"] = o.IEzsigntemplatesignaturecustomdateOffsetx
	toSerialize["iEzsigntemplatesignaturecustomdateOffsety"] = o.IEzsigntemplatesignaturecustomdateOffsety
	toSerialize["sEzsigntemplatesignaturecustomdateFormat"] = o.SEzsigntemplatesignaturecustomdateFormat
	return toSerialize, nil
}

func (o *EzsigntemplatesignaturecustomdateResponseV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatesignaturecustomdateID",
		"iEzsigntemplatesignaturecustomdateOffsetx",
		"iEzsigntemplatesignaturecustomdateOffsety",
		"sEzsigntemplatesignaturecustomdateFormat",
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

	varEzsigntemplatesignaturecustomdateResponseV2 := _EzsigntemplatesignaturecustomdateResponseV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignaturecustomdateResponseV2)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignaturecustomdateResponseV2(varEzsigntemplatesignaturecustomdateResponseV2)

	return err
}

type NullableEzsigntemplatesignaturecustomdateResponseV2 struct {
	value *EzsigntemplatesignaturecustomdateResponseV2
	isSet bool
}

func (v NullableEzsigntemplatesignaturecustomdateResponseV2) Get() *EzsigntemplatesignaturecustomdateResponseV2 {
	return v.value
}

func (v *NullableEzsigntemplatesignaturecustomdateResponseV2) Set(val *EzsigntemplatesignaturecustomdateResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignaturecustomdateResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignaturecustomdateResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignaturecustomdateResponseV2(val *EzsigntemplatesignaturecustomdateResponseV2) *NullableEzsigntemplatesignaturecustomdateResponseV2 {
	return &NullableEzsigntemplatesignaturecustomdateResponseV2{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignaturecustomdateResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignaturecustomdateResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


