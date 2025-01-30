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

// checks if the EzsigntemplatesignaturecustomdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignaturecustomdateRequest{}

// EzsigntemplatesignaturecustomdateRequest An Ezsigntemplatesignaturecustomdate Object
type EzsigntemplatesignaturecustomdateRequest struct {
	// The unique ID of the Ezsigntemplatesignaturecustomdate
	PkiEzsigntemplatesignaturecustomdateID *int32 `json:"pkiEzsigntemplatesignaturecustomdateID,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	// Deprecated
	IEzsigntemplatesignaturecustomdateX *int32 `json:"iEzsigntemplatesignaturecustomdateX,omitempty"`
	// The Y coordinate (Vertical) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	// Deprecated
	IEzsigntemplatesignaturecustomdateY *int32 `json:"iEzsigntemplatesignaturecustomdateY,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the left of the signature, you would use \"200\" for the X coordinate.
	IEzsigntemplatesignaturecustomdateOffsetx *int32 `json:"iEzsigntemplatesignaturecustomdateOffsetx,omitempty"`
	// The X coordinate (Horizontal) where to put the Ezsigntemplatesignaturecustomdate on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the Ezsigntemplatesignaturecustomdate 2 inches from the top of the signature, you would use \"200\" for the Y coordinate.
	IEzsigntemplatesignaturecustomdateOffsety *int32 `json:"iEzsigntemplatesignaturecustomdateOffsety,omitempty"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 | | {YY} | 22 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsigntemplatesignaturecustomdateFormat string `json:"sEzsigntemplatesignaturecustomdateFormat"`
}

type _EzsigntemplatesignaturecustomdateRequest EzsigntemplatesignaturecustomdateRequest

// NewEzsigntemplatesignaturecustomdateRequest instantiates a new EzsigntemplatesignaturecustomdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignaturecustomdateRequest(sEzsigntemplatesignaturecustomdateFormat string) *EzsigntemplatesignaturecustomdateRequest {
	this := EzsigntemplatesignaturecustomdateRequest{}
	this.SEzsigntemplatesignaturecustomdateFormat = sEzsigntemplatesignaturecustomdateFormat
	return &this
}

// NewEzsigntemplatesignaturecustomdateRequestWithDefaults instantiates a new EzsigntemplatesignaturecustomdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignaturecustomdateRequestWithDefaults() *EzsigntemplatesignaturecustomdateRequest {
	this := EzsigntemplatesignaturecustomdateRequest{}
	return &this
}

// GetPkiEzsigntemplatesignaturecustomdateID returns the PkiEzsigntemplatesignaturecustomdateID field value if set, zero value otherwise.
func (o *EzsigntemplatesignaturecustomdateRequest) GetPkiEzsigntemplatesignaturecustomdateID() int32 {
	if o == nil || IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsigntemplatesignaturecustomdateID
}

// GetPkiEzsigntemplatesignaturecustomdateIDOk returns a tuple with the PkiEzsigntemplatesignaturecustomdateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) GetPkiEzsigntemplatesignaturecustomdateIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		return nil, false
	}
	return o.PkiEzsigntemplatesignaturecustomdateID, true
}

// HasPkiEzsigntemplatesignaturecustomdateID returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) HasPkiEzsigntemplatesignaturecustomdateID() bool {
	if o != nil && !IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		return true
	}

	return false
}

// SetPkiEzsigntemplatesignaturecustomdateID gets a reference to the given int32 and assigns it to the PkiEzsigntemplatesignaturecustomdateID field.
func (o *EzsigntemplatesignaturecustomdateRequest) SetPkiEzsigntemplatesignaturecustomdateID(v int32) {
	o.PkiEzsigntemplatesignaturecustomdateID = &v
}

// GetIEzsigntemplatesignaturecustomdateX returns the IEzsigntemplatesignaturecustomdateX field value if set, zero value otherwise.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateX() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateX) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignaturecustomdateX
}

// GetIEzsigntemplatesignaturecustomdateXOk returns a tuple with the IEzsigntemplatesignaturecustomdateX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateXOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateX) {
		return nil, false
	}
	return o.IEzsigntemplatesignaturecustomdateX, true
}

// HasIEzsigntemplatesignaturecustomdateX returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) HasIEzsigntemplatesignaturecustomdateX() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignaturecustomdateX) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignaturecustomdateX gets a reference to the given int32 and assigns it to the IEzsigntemplatesignaturecustomdateX field.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) SetIEzsigntemplatesignaturecustomdateX(v int32) {
	o.IEzsigntemplatesignaturecustomdateX = &v
}

// GetIEzsigntemplatesignaturecustomdateY returns the IEzsigntemplatesignaturecustomdateY field value if set, zero value otherwise.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateY() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateY) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignaturecustomdateY
}

// GetIEzsigntemplatesignaturecustomdateYOk returns a tuple with the IEzsigntemplatesignaturecustomdateY field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateYOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateY) {
		return nil, false
	}
	return o.IEzsigntemplatesignaturecustomdateY, true
}

// HasIEzsigntemplatesignaturecustomdateY returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) HasIEzsigntemplatesignaturecustomdateY() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignaturecustomdateY) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignaturecustomdateY gets a reference to the given int32 and assigns it to the IEzsigntemplatesignaturecustomdateY field.
// Deprecated
func (o *EzsigntemplatesignaturecustomdateRequest) SetIEzsigntemplatesignaturecustomdateY(v int32) {
	o.IEzsigntemplatesignaturecustomdateY = &v
}

// GetIEzsigntemplatesignaturecustomdateOffsetx returns the IEzsigntemplatesignaturecustomdateOffsetx field value if set, zero value otherwise.
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateOffsetx() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateOffsetx) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignaturecustomdateOffsetx
}

// GetIEzsigntemplatesignaturecustomdateOffsetxOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsetx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateOffsetxOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateOffsetx) {
		return nil, false
	}
	return o.IEzsigntemplatesignaturecustomdateOffsetx, true
}

// HasIEzsigntemplatesignaturecustomdateOffsetx returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) HasIEzsigntemplatesignaturecustomdateOffsetx() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignaturecustomdateOffsetx) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignaturecustomdateOffsetx gets a reference to the given int32 and assigns it to the IEzsigntemplatesignaturecustomdateOffsetx field.
func (o *EzsigntemplatesignaturecustomdateRequest) SetIEzsigntemplatesignaturecustomdateOffsetx(v int32) {
	o.IEzsigntemplatesignaturecustomdateOffsetx = &v
}

// GetIEzsigntemplatesignaturecustomdateOffsety returns the IEzsigntemplatesignaturecustomdateOffsety field value if set, zero value otherwise.
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateOffsety() int32 {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateOffsety) {
		var ret int32
		return ret
	}
	return *o.IEzsigntemplatesignaturecustomdateOffsety
}

// GetIEzsigntemplatesignaturecustomdateOffsetyOk returns a tuple with the IEzsigntemplatesignaturecustomdateOffsety field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) GetIEzsigntemplatesignaturecustomdateOffsetyOk() (*int32, bool) {
	if o == nil || IsNil(o.IEzsigntemplatesignaturecustomdateOffsety) {
		return nil, false
	}
	return o.IEzsigntemplatesignaturecustomdateOffsety, true
}

// HasIEzsigntemplatesignaturecustomdateOffsety returns a boolean if a field has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) HasIEzsigntemplatesignaturecustomdateOffsety() bool {
	if o != nil && !IsNil(o.IEzsigntemplatesignaturecustomdateOffsety) {
		return true
	}

	return false
}

// SetIEzsigntemplatesignaturecustomdateOffsety gets a reference to the given int32 and assigns it to the IEzsigntemplatesignaturecustomdateOffsety field.
func (o *EzsigntemplatesignaturecustomdateRequest) SetIEzsigntemplatesignaturecustomdateOffsety(v int32) {
	o.IEzsigntemplatesignaturecustomdateOffsety = &v
}

// GetSEzsigntemplatesignaturecustomdateFormat returns the SEzsigntemplatesignaturecustomdateFormat field value
func (o *EzsigntemplatesignaturecustomdateRequest) GetSEzsigntemplatesignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignaturecustomdateFormat
}

// GetSEzsigntemplatesignaturecustomdateFormatOk returns a tuple with the SEzsigntemplatesignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignaturecustomdateRequest) GetSEzsigntemplatesignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignaturecustomdateFormat, true
}

// SetSEzsigntemplatesignaturecustomdateFormat sets field value
func (o *EzsigntemplatesignaturecustomdateRequest) SetSEzsigntemplatesignaturecustomdateFormat(v string) {
	o.SEzsigntemplatesignaturecustomdateFormat = v
}

func (o EzsigntemplatesignaturecustomdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignaturecustomdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsigntemplatesignaturecustomdateID) {
		toSerialize["pkiEzsigntemplatesignaturecustomdateID"] = o.PkiEzsigntemplatesignaturecustomdateID
	}
	if !IsNil(o.IEzsigntemplatesignaturecustomdateX) {
		toSerialize["iEzsigntemplatesignaturecustomdateX"] = o.IEzsigntemplatesignaturecustomdateX
	}
	if !IsNil(o.IEzsigntemplatesignaturecustomdateY) {
		toSerialize["iEzsigntemplatesignaturecustomdateY"] = o.IEzsigntemplatesignaturecustomdateY
	}
	if !IsNil(o.IEzsigntemplatesignaturecustomdateOffsetx) {
		toSerialize["iEzsigntemplatesignaturecustomdateOffsetx"] = o.IEzsigntemplatesignaturecustomdateOffsetx
	}
	if !IsNil(o.IEzsigntemplatesignaturecustomdateOffsety) {
		toSerialize["iEzsigntemplatesignaturecustomdateOffsety"] = o.IEzsigntemplatesignaturecustomdateOffsety
	}
	toSerialize["sEzsigntemplatesignaturecustomdateFormat"] = o.SEzsigntemplatesignaturecustomdateFormat
	return toSerialize, nil
}

func (o *EzsigntemplatesignaturecustomdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsigntemplatesignaturecustomdateRequest := _EzsigntemplatesignaturecustomdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignaturecustomdateRequest)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignaturecustomdateRequest(varEzsigntemplatesignaturecustomdateRequest)

	return err
}

type NullableEzsigntemplatesignaturecustomdateRequest struct {
	value *EzsigntemplatesignaturecustomdateRequest
	isSet bool
}

func (v NullableEzsigntemplatesignaturecustomdateRequest) Get() *EzsigntemplatesignaturecustomdateRequest {
	return v.value
}

func (v *NullableEzsigntemplatesignaturecustomdateRequest) Set(val *EzsigntemplatesignaturecustomdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignaturecustomdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignaturecustomdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignaturecustomdateRequest(val *EzsigntemplatesignaturecustomdateRequest) *NullableEzsigntemplatesignaturecustomdateRequest {
	return &NullableEzsigntemplatesignaturecustomdateRequest{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignaturecustomdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignaturecustomdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


