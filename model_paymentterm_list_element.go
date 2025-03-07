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

// checks if the PaymenttermListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermListElement{}

// PaymenttermListElement A Paymentterm List Element
type PaymenttermListElement struct {
	// The unique ID of the Paymentterm
	PkiPaymenttermID int32 `json:"pkiPaymenttermID"`
	// The code of the Paymentterm
	SPaymenttermCode string `json:"sPaymenttermCode" validate:"regexp=^[A-Z0-9]{1,4}$"`
	EPaymenttermType FieldEPaymenttermType `json:"ePaymenttermType"`
	// The day of the Paymentterm
	IPaymenttermDay int32 `json:"iPaymenttermDay"`
	// The description of the Paymentterm in the language of the requester
	SPaymenttermDescriptionX string `json:"sPaymenttermDescriptionX" validate:"regexp=^.{1,40}$"`
	// Whether the Paymentterm is active or not
	BPaymenttermIsactive bool `json:"bPaymenttermIsactive"`
}

type _PaymenttermListElement PaymenttermListElement

// NewPaymenttermListElement instantiates a new PaymenttermListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermListElement(pkiPaymenttermID int32, sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, sPaymenttermDescriptionX string, bPaymenttermIsactive bool) *PaymenttermListElement {
	this := PaymenttermListElement{}
	this.PkiPaymenttermID = pkiPaymenttermID
	this.SPaymenttermCode = sPaymenttermCode
	this.EPaymenttermType = ePaymenttermType
	this.IPaymenttermDay = iPaymenttermDay
	this.SPaymenttermDescriptionX = sPaymenttermDescriptionX
	this.BPaymenttermIsactive = bPaymenttermIsactive
	return &this
}

// NewPaymenttermListElementWithDefaults instantiates a new PaymenttermListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermListElementWithDefaults() *PaymenttermListElement {
	this := PaymenttermListElement{}
	return &this
}

// GetPkiPaymenttermID returns the PkiPaymenttermID field value
func (o *PaymenttermListElement) GetPkiPaymenttermID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPaymenttermID
}

// GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetPkiPaymenttermIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPaymenttermID, true
}

// SetPkiPaymenttermID sets field value
func (o *PaymenttermListElement) SetPkiPaymenttermID(v int32) {
	o.PkiPaymenttermID = v
}

// GetSPaymenttermCode returns the SPaymenttermCode field value
func (o *PaymenttermListElement) GetSPaymenttermCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPaymenttermCode
}

// GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetSPaymenttermCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SPaymenttermCode, true
}

// SetSPaymenttermCode sets field value
func (o *PaymenttermListElement) SetSPaymenttermCode(v string) {
	o.SPaymenttermCode = v
}

// GetEPaymenttermType returns the EPaymenttermType field value
func (o *PaymenttermListElement) GetEPaymenttermType() FieldEPaymenttermType {
	if o == nil {
		var ret FieldEPaymenttermType
		return ret
	}

	return o.EPaymenttermType
}

// GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EPaymenttermType, true
}

// SetEPaymenttermType sets field value
func (o *PaymenttermListElement) SetEPaymenttermType(v FieldEPaymenttermType) {
	o.EPaymenttermType = v
}

// GetIPaymenttermDay returns the IPaymenttermDay field value
func (o *PaymenttermListElement) GetIPaymenttermDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IPaymenttermDay
}

// GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetIPaymenttermDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPaymenttermDay, true
}

// SetIPaymenttermDay sets field value
func (o *PaymenttermListElement) SetIPaymenttermDay(v int32) {
	o.IPaymenttermDay = v
}

// GetSPaymenttermDescriptionX returns the SPaymenttermDescriptionX field value
func (o *PaymenttermListElement) GetSPaymenttermDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPaymenttermDescriptionX
}

// GetSPaymenttermDescriptionXOk returns a tuple with the SPaymenttermDescriptionX field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetSPaymenttermDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SPaymenttermDescriptionX, true
}

// SetSPaymenttermDescriptionX sets field value
func (o *PaymenttermListElement) SetSPaymenttermDescriptionX(v string) {
	o.SPaymenttermDescriptionX = v
}

// GetBPaymenttermIsactive returns the BPaymenttermIsactive field value
func (o *PaymenttermListElement) GetBPaymenttermIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BPaymenttermIsactive
}

// GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field value
// and a boolean to check if the value has been set.
func (o *PaymenttermListElement) GetBPaymenttermIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BPaymenttermIsactive, true
}

// SetBPaymenttermIsactive sets field value
func (o *PaymenttermListElement) SetBPaymenttermIsactive(v bool) {
	o.BPaymenttermIsactive = v
}

func (o PaymenttermListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPaymenttermID"] = o.PkiPaymenttermID
	toSerialize["sPaymenttermCode"] = o.SPaymenttermCode
	toSerialize["ePaymenttermType"] = o.EPaymenttermType
	toSerialize["iPaymenttermDay"] = o.IPaymenttermDay
	toSerialize["sPaymenttermDescriptionX"] = o.SPaymenttermDescriptionX
	toSerialize["bPaymenttermIsactive"] = o.BPaymenttermIsactive
	return toSerialize, nil
}

func (o *PaymenttermListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPaymenttermID",
		"sPaymenttermCode",
		"ePaymenttermType",
		"iPaymenttermDay",
		"sPaymenttermDescriptionX",
		"bPaymenttermIsactive",
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

	varPaymenttermListElement := _PaymenttermListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermListElement)

	if err != nil {
		return err
	}

	*o = PaymenttermListElement(varPaymenttermListElement)

	return err
}

type NullablePaymenttermListElement struct {
	value *PaymenttermListElement
	isSet bool
}

func (v NullablePaymenttermListElement) Get() *PaymenttermListElement {
	return v.value
}

func (v *NullablePaymenttermListElement) Set(val *PaymenttermListElement) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermListElement) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermListElement(val *PaymenttermListElement) *NullablePaymenttermListElement {
	return &NullablePaymenttermListElement{value: val, isSet: true}
}

func (v NullablePaymenttermListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


