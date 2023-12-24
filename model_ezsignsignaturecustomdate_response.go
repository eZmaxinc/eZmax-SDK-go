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

// checks if the EzsignsignaturecustomdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignaturecustomdateResponse{}

// EzsignsignaturecustomdateResponse An Ezsignsignaturecustomdate Object
type EzsignsignaturecustomdateResponse struct {
	// The unique ID of the Ezsignsignaturecustomdate
	PkiEzsignsignaturecustomdateID int32 `json:"pkiEzsignsignaturecustomdateID"`
	// The X coordinate (Horizontal) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignsignaturecustomdateX int32 `json:"iEzsignsignaturecustomdateX"`
	// The Y coordinate (Vertical) where to put the Ezsignsignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsignsignaturecustomdate block 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignsignaturecustomdateY int32 `json:"iEzsignsignaturecustomdateY"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsignsignaturecustomdateFormat string `json:"sEzsignsignaturecustomdateFormat"`
}

type _EzsignsignaturecustomdateResponse EzsignsignaturecustomdateResponse

// NewEzsignsignaturecustomdateResponse instantiates a new EzsignsignaturecustomdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignaturecustomdateResponse(pkiEzsignsignaturecustomdateID int32, iEzsignsignaturecustomdateX int32, iEzsignsignaturecustomdateY int32, sEzsignsignaturecustomdateFormat string) *EzsignsignaturecustomdateResponse {
	this := EzsignsignaturecustomdateResponse{}
	this.PkiEzsignsignaturecustomdateID = pkiEzsignsignaturecustomdateID
	this.IEzsignsignaturecustomdateX = iEzsignsignaturecustomdateX
	this.IEzsignsignaturecustomdateY = iEzsignsignaturecustomdateY
	this.SEzsignsignaturecustomdateFormat = sEzsignsignaturecustomdateFormat
	return &this
}

// NewEzsignsignaturecustomdateResponseWithDefaults instantiates a new EzsignsignaturecustomdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignaturecustomdateResponseWithDefaults() *EzsignsignaturecustomdateResponse {
	this := EzsignsignaturecustomdateResponse{}
	return &this
}

// GetPkiEzsignsignaturecustomdateID returns the PkiEzsignsignaturecustomdateID field value
func (o *EzsignsignaturecustomdateResponse) GetPkiEzsignsignaturecustomdateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignaturecustomdateID
}

// GetPkiEzsignsignaturecustomdateIDOk returns a tuple with the PkiEzsignsignaturecustomdateID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponse) GetPkiEzsignsignaturecustomdateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignaturecustomdateID, true
}

// SetPkiEzsignsignaturecustomdateID sets field value
func (o *EzsignsignaturecustomdateResponse) SetPkiEzsignsignaturecustomdateID(v int32) {
	o.PkiEzsignsignaturecustomdateID = v
}

// GetIEzsignsignaturecustomdateX returns the IEzsignsignaturecustomdateX field value
func (o *EzsignsignaturecustomdateResponse) GetIEzsignsignaturecustomdateX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateX
}

// GetIEzsignsignaturecustomdateXOk returns a tuple with the IEzsignsignaturecustomdateX field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponse) GetIEzsignsignaturecustomdateXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateX, true
}

// SetIEzsignsignaturecustomdateX sets field value
func (o *EzsignsignaturecustomdateResponse) SetIEzsignsignaturecustomdateX(v int32) {
	o.IEzsignsignaturecustomdateX = v
}

// GetIEzsignsignaturecustomdateY returns the IEzsignsignaturecustomdateY field value
func (o *EzsignsignaturecustomdateResponse) GetIEzsignsignaturecustomdateY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateY
}

// GetIEzsignsignaturecustomdateYOk returns a tuple with the IEzsignsignaturecustomdateY field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponse) GetIEzsignsignaturecustomdateYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateY, true
}

// SetIEzsignsignaturecustomdateY sets field value
func (o *EzsignsignaturecustomdateResponse) SetIEzsignsignaturecustomdateY(v int32) {
	o.IEzsignsignaturecustomdateY = v
}

// GetSEzsignsignaturecustomdateFormat returns the SEzsignsignaturecustomdateFormat field value
func (o *EzsignsignaturecustomdateResponse) GetSEzsignsignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsignaturecustomdateFormat
}

// GetSEzsignsignaturecustomdateFormatOk returns a tuple with the SEzsignsignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateResponse) GetSEzsignsignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsignaturecustomdateFormat, true
}

// SetSEzsignsignaturecustomdateFormat sets field value
func (o *EzsignsignaturecustomdateResponse) SetSEzsignsignaturecustomdateFormat(v string) {
	o.SEzsignsignaturecustomdateFormat = v
}

func (o EzsignsignaturecustomdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignaturecustomdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignaturecustomdateID"] = o.PkiEzsignsignaturecustomdateID
	toSerialize["iEzsignsignaturecustomdateX"] = o.IEzsignsignaturecustomdateX
	toSerialize["iEzsignsignaturecustomdateY"] = o.IEzsignsignaturecustomdateY
	toSerialize["sEzsignsignaturecustomdateFormat"] = o.SEzsignsignaturecustomdateFormat
	return toSerialize, nil
}

func (o *EzsignsignaturecustomdateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignaturecustomdateID",
		"iEzsignsignaturecustomdateX",
		"iEzsignsignaturecustomdateY",
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

	varEzsignsignaturecustomdateResponse := _EzsignsignaturecustomdateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignaturecustomdateResponse)

	if err != nil {
		return err
	}

	*o = EzsignsignaturecustomdateResponse(varEzsignsignaturecustomdateResponse)

	return err
}

type NullableEzsignsignaturecustomdateResponse struct {
	value *EzsignsignaturecustomdateResponse
	isSet bool
}

func (v NullableEzsignsignaturecustomdateResponse) Get() *EzsignsignaturecustomdateResponse {
	return v.value
}

func (v *NullableEzsignsignaturecustomdateResponse) Set(val *EzsignsignaturecustomdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignaturecustomdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignaturecustomdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignaturecustomdateResponse(val *EzsignsignaturecustomdateResponse) *NullableEzsignsignaturecustomdateResponse {
	return &NullableEzsignsignaturecustomdateResponse{value: val, isSet: true}
}

func (v NullableEzsignsignaturecustomdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignaturecustomdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


