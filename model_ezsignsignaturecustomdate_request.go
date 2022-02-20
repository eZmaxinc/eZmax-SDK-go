/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignsignaturecustomdateRequest An Ezsignsignaturecustomdate Object
type EzsignsignaturecustomdateRequest struct {
	// The unique ID of the Ezsignsignaturecustomdate
	PkiEzsignsignaturecustomdateID *int32 `json:"pkiEzsignsignaturecustomdateID,omitempty"`
	// The X coordinate (Horizontal) where to put the custom date block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the custom date block 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignsignaturecustomdateX int32 `json:"iEzsignsignaturecustomdateX"`
	// The Y coordinate (Vertical) where to put the custom date block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the custom date block 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignsignaturecustomdateY int32 `json:"iEzsignsignaturecustomdateY"`
	// The custom date format to use  You can use the codes below and they will be replaced at signature time. Text values like month and day names will be rendered in the proper language. Other text will be left as-is.  The codes examples below are based on the following datetime: Thursday, January 6, 2022 at 08:07:09 EST  For example, the format \"Signature date: {MM}/{DD}/{YYYY} {hh}:{mm}\" would become \"Signature date: 01/06/2022 08:07\"  **Year**  | Code | Example | | - | - | | {YYYY} | 2022 |  **Month**  | Code | Example | | - | - | | {MonthCapitalize} | Janvier | | {Month} | janvier | | {MM} | 01 | | {M} | 1 |  **Day**  | Code | Example | | - | - | | {DayCapitalize} | Jeudi | | {Day} | jeudi | | {DD} | 06 | | {D} | 6 |  **Hour**  | Code | Example | | - | - | | {hh} | 08 |  **Minute**  | Code | Example | | - | - | | {mm} | 07 |  **Second**  | Code | Example | | - | - | | {ss} | 09 |        **Timezone**  | Code | Example | | - | - | | {Z} | EST |       **Time**  | Code | Example | | - | - | | {Time} | 08:07:09 |   | {TimeZ} | 08:07:09 EST |     **Date**  | Code | Example | | - | - | | {Date} | 2022-01-06 |   | {DateText} | 1er Janvier 2022 |  **Full**  | Code | Example | | - | - | | {DateTime} | 2022-01-06 08:07:09 |   | {DateTimeZ} | 2022-01-06 08:07:09 EST | 
	SEzsignsignaturecustomdateFormat string `json:"sEzsignsignaturecustomdateFormat"`
}

// NewEzsignsignaturecustomdateRequest instantiates a new EzsignsignaturecustomdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignaturecustomdateRequest(iEzsignsignaturecustomdateX int32, iEzsignsignaturecustomdateY int32, sEzsignsignaturecustomdateFormat string) *EzsignsignaturecustomdateRequest {
	this := EzsignsignaturecustomdateRequest{}
	this.IEzsignsignaturecustomdateX = iEzsignsignaturecustomdateX
	this.IEzsignsignaturecustomdateY = iEzsignsignaturecustomdateY
	this.SEzsignsignaturecustomdateFormat = sEzsignsignaturecustomdateFormat
	return &this
}

// NewEzsignsignaturecustomdateRequestWithDefaults instantiates a new EzsignsignaturecustomdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignaturecustomdateRequestWithDefaults() *EzsignsignaturecustomdateRequest {
	this := EzsignsignaturecustomdateRequest{}
	return &this
}

// GetPkiEzsignsignaturecustomdateID returns the PkiEzsignsignaturecustomdateID field value if set, zero value otherwise.
func (o *EzsignsignaturecustomdateRequest) GetPkiEzsignsignaturecustomdateID() int32 {
	if o == nil || o.PkiEzsignsignaturecustomdateID == nil {
		var ret int32
		return ret
	}
	return *o.PkiEzsignsignaturecustomdateID
}

// GetPkiEzsignsignaturecustomdateIDOk returns a tuple with the PkiEzsignsignaturecustomdateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateRequest) GetPkiEzsignsignaturecustomdateIDOk() (*int32, bool) {
	if o == nil || o.PkiEzsignsignaturecustomdateID == nil {
		return nil, false
	}
	return o.PkiEzsignsignaturecustomdateID, true
}

// HasPkiEzsignsignaturecustomdateID returns a boolean if a field has been set.
func (o *EzsignsignaturecustomdateRequest) HasPkiEzsignsignaturecustomdateID() bool {
	if o != nil && o.PkiEzsignsignaturecustomdateID != nil {
		return true
	}

	return false
}

// SetPkiEzsignsignaturecustomdateID gets a reference to the given int32 and assigns it to the PkiEzsignsignaturecustomdateID field.
func (o *EzsignsignaturecustomdateRequest) SetPkiEzsignsignaturecustomdateID(v int32) {
	o.PkiEzsignsignaturecustomdateID = &v
}

// GetIEzsignsignaturecustomdateX returns the IEzsignsignaturecustomdateX field value
func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateX
}

// GetIEzsignsignaturecustomdateXOk returns a tuple with the IEzsignsignaturecustomdateX field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateX, true
}

// SetIEzsignsignaturecustomdateX sets field value
func (o *EzsignsignaturecustomdateRequest) SetIEzsignsignaturecustomdateX(v int32) {
	o.IEzsignsignaturecustomdateX = v
}

// GetIEzsignsignaturecustomdateY returns the IEzsignsignaturecustomdateY field value
func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignaturecustomdateY
}

// GetIEzsignsignaturecustomdateYOk returns a tuple with the IEzsignsignaturecustomdateY field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateRequest) GetIEzsignsignaturecustomdateYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignsignaturecustomdateY, true
}

// SetIEzsignsignaturecustomdateY sets field value
func (o *EzsignsignaturecustomdateRequest) SetIEzsignsignaturecustomdateY(v int32) {
	o.IEzsignsignaturecustomdateY = v
}

// GetSEzsignsignaturecustomdateFormat returns the SEzsignsignaturecustomdateFormat field value
func (o *EzsignsignaturecustomdateRequest) GetSEzsignsignaturecustomdateFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsignaturecustomdateFormat
}

// GetSEzsignsignaturecustomdateFormatOk returns a tuple with the SEzsignsignaturecustomdateFormat field value
// and a boolean to check if the value has been set.
func (o *EzsignsignaturecustomdateRequest) GetSEzsignsignaturecustomdateFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsignaturecustomdateFormat, true
}

// SetSEzsignsignaturecustomdateFormat sets field value
func (o *EzsignsignaturecustomdateRequest) SetSEzsignsignaturecustomdateFormat(v string) {
	o.SEzsignsignaturecustomdateFormat = v
}

func (o EzsignsignaturecustomdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PkiEzsignsignaturecustomdateID != nil {
		toSerialize["pkiEzsignsignaturecustomdateID"] = o.PkiEzsignsignaturecustomdateID
	}
	if true {
		toSerialize["iEzsignsignaturecustomdateX"] = o.IEzsignsignaturecustomdateX
	}
	if true {
		toSerialize["iEzsignsignaturecustomdateY"] = o.IEzsignsignaturecustomdateY
	}
	if true {
		toSerialize["sEzsignsignaturecustomdateFormat"] = o.SEzsignsignaturecustomdateFormat
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignaturecustomdateRequest struct {
	value *EzsignsignaturecustomdateRequest
	isSet bool
}

func (v NullableEzsignsignaturecustomdateRequest) Get() *EzsignsignaturecustomdateRequest {
	return v.value
}

func (v *NullableEzsignsignaturecustomdateRequest) Set(val *EzsignsignaturecustomdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignaturecustomdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignaturecustomdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignaturecustomdateRequest(val *EzsignsignaturecustomdateRequest) *NullableEzsignsignaturecustomdateRequest {
	return &NullableEzsignsignaturecustomdateRequest{value: val, isSet: true}
}

func (v NullableEzsignsignaturecustomdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignaturecustomdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


