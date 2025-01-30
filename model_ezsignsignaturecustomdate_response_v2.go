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

// checks if the EzsignsignaturecustomdateResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignaturecustomdateResponseV2{}

// EzsignsignaturecustomdateResponseV2 An Ezsignsignaturecustomdate Object
type EzsignsignaturecustomdateResponseV2 struct {
	// The unique ID of the Ezsignsignaturecustomdate
	PkiEzsignsignaturecustomdateID int32 `json:"pkiEzsignsignaturecustomdateID"`
	// The X coordinate (Horizontal) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 2 inches from the left of the signature, you would use \"200\" for the X coordinate.
	IEzsignsignaturecustomdateOffsetx int32 `json:"iEzsignsignaturecustomdateOffsetx"`
	// The Y coordinate (Vertical) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 3 inches from the top of the signature, you would use \"300\" for the Y coordinate.
	IEzsignsignaturecustomdateOffsety int32 `json:"iEzsignsignaturecustomdateOffsety"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsignsignaturecustomdateFormat string `json:"sEzsignsignaturecustomdateFormat"`
}

type _EzsignsignaturecustomdateResponseV2 EzsignsignaturecustomdateResponseV2

// NewEzsignsignaturecustomdateResponseV2 instantiates a new EzsignsignaturecustomdateResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignaturecustomdateResponseV2(pkiEzsignsignaturecustomdateID int32, iEzsignsignaturecustomdateOffsetx int32, iEzsignsignaturecustomdateOffsety int32, sEzsignsignaturecustomdateFormat string) *EzsignsignaturecustomdateResponseV2 {
	this := EzsignsignaturecustomdateResponseV2{}
	this.PkiEzsignsignaturecustomdateID = pkiEzsignsignaturecustomdateID
	this.IEzsignsignaturecustomdateOffsetx = iEzsignsignaturecustomdateOffsetx
	this.IEzsignsignaturecustomdateOffsety = iEzsignsignaturecustomdateOffsety
	this.SEzsignsignaturecustomdateFormat = sEzsignsignaturecustomdateFormat
	return &this
}

// NewEzsignsignaturecustomdateResponseV2WithDefaults instantiates a new EzsignsignaturecustomdateResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignaturecustomdateResponseV2WithDefaults() *EzsignsignaturecustomdateResponseV2 {
	this := EzsignsignaturecustomdateResponseV2{}
	return &this
}

// GetPkiEzsignsignaturecustomdateID returns the PkiEzsignsignaturecustomdateID field value
func (o *EzsignsignaturecustomdateResponseV2) GetPkiEzsignsignaturecustomdateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignaturecustomdateID
}

// GetPkiEzsignsignaturecustomdateIDOk returns a tuple with the PkiEzsignsignaturecustomdateID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponseV2) GetPkiEzsignsignaturecustomdateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignaturecustomdateID, true
}

// SetPkiEzsignsignaturecustomdateID sets field value
func (o *EzsignsignaturecustomdateResponseV2) SetPkiEzsignsignaturecustomdateID(v int32) {
	o.PkiEzsignsignaturecustomdateID = v
}

// GetIEzsignsignaturecustomdateOffsetx returns the IEzsignsignaturecustomdateOffsetx field value
func (o *EzsignsignaturecustomdateResponseV2) GetIEzsignsignaturecustomdateOffsetx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateOffsetx
}

// GetIEzsignsignaturecustomdateOffsetxOk returns a tuple with the IEzsignsignaturecustomdateOffsetx field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponseV2) GetIEzsignsignaturecustomdateOffsetxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateOffsetx, true
}

// SetIEzsignsignaturecustomdateOffsetx sets field value
func (o *EzsignsignaturecustomdateResponseV2) SetIEzsignsignaturecustomdateOffsetx(v int32) {
	o.IEzsignsignaturecustomdateOffsetx = v
}

// GetIEzsignsignaturecustomdateOffsety returns the IEzsignsignaturecustomdateOffsety field value
func (o *EzsignsignaturecustomdateResponseV2) GetIEzsignsignaturecustomdateOffsety() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateOffsety
}

// GetIEzsignsignaturecustomdateOffsetyOk returns a tuple with the IEzsignsignaturecustomdateOffsety field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponseV2) GetIEzsignsignaturecustomdateOffsetyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateOffsety, true
}

// SetIEzsignsignaturecustomdateOffsety sets field value
func (o *EzsignsignaturecustomdateResponseV2) SetIEzsignsignaturecustomdateOffsety(v int32) {
	o.IEzsignsignaturecustomdateOffsety = v
}

// GetSEzsignsignaturecustomdateFormat returns the SEzsignsignaturecustomdateFormat field value
func (o *EzsignsignaturecustomdateResponseV2) GetSEzsignsignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsignaturecustomdateFormat
}

// GetSEzsignsignaturecustomdateFormatOk returns a tuple with the SEzsignsignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponseV2) GetSEzsignsignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsignaturecustomdateFormat, true
}

// SetSEzsignsignaturecustomdateFormat sets field value
func (o *EzsignsignaturecustomdateResponseV2) SetSEzsignsignaturecustomdateFormat(v string) {
	o.SEzsignsignaturecustomdateFormat = v
}

func (o EzsignsignaturecustomdateResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignaturecustomdateResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignaturecustomdateID"] = o.PkiEzsignsignaturecustomdateID
	toSerialize["iEzsignsignaturecustomdateOffsetx"] = o.IEzsignsignaturecustomdateOffsetx
	toSerialize["iEzsignsignaturecustomdateOffsety"] = o.IEzsignsignaturecustomdateOffsety
	toSerialize["sEzsignsignaturecustomdateFormat"] = o.SEzsignsignaturecustomdateFormat
	return toSerialize, nil
}

func (o *EzsignsignaturecustomdateResponseV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignaturecustomdateID",
		"iEzsignsignaturecustomdateOffsetx",
		"iEzsignsignaturecustomdateOffsety",
		"sEzsignsignaturecustomdateFormat",
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

	varEzsignsignaturecustomdateResponseV2 := _EzsignsignaturecustomdateResponseV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignaturecustomdateResponseV2)

	if err != nil {
		return err
	}

	*o = EzsignsignaturecustomdateResponseV2(varEzsignsignaturecustomdateResponseV2)

	return err
}

type NullableEzsignsignaturecustomdateResponseV2 struct {
	value *EzsignsignaturecustomdateResponseV2
	isSet bool
}

func (v NullableEzsignsignaturecustomdateResponseV2) Get() *EzsignsignaturecustomdateResponseV2 {
	return v.value
}

func (v *NullableEzsignsignaturecustomdateResponseV2) Set(val *EzsignsignaturecustomdateResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignaturecustomdateResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignaturecustomdateResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignaturecustomdateResponseV2(val *EzsignsignaturecustomdateResponseV2) *NullableEzsignsignaturecustomdateResponseV2 {
	return &NullableEzsignsignaturecustomdateResponseV2{value: val, isSet: true}
}

func (v NullableEzsignsignaturecustomdateResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignaturecustomdateResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


