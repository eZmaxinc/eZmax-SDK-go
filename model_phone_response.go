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

// checks if the PhoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneResponse{}

// PhoneResponse A Phone Object
type PhoneResponse struct {
	// The unique ID of the Phone.
	PkiPhoneID int32 `json:"pkiPhoneID"`
	// The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free|
	FkiPhonetypeID int32 `json:"fkiPhonetypeID"`
	// Deprecated
	EPhoneType *FieldEPhoneType `json:"ePhoneType,omitempty"`
	// A phone number in E.164 Format
	SPhoneE164 *string `json:"sPhoneE164,omitempty" validate:"regexp=^\\\\+[1-9]\\\\d{1,14}$"`
	// The extension of the phone number.  The extension is the \"123\" section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers
	SPhoneExtension *string `json:"sPhoneExtension,omitempty"`
}

type _PhoneResponse PhoneResponse

// NewPhoneResponse instantiates a new PhoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneResponse(pkiPhoneID int32, fkiPhonetypeID int32) *PhoneResponse {
	this := PhoneResponse{}
	this.PkiPhoneID = pkiPhoneID
	this.FkiPhonetypeID = fkiPhonetypeID
	return &this
}

// NewPhoneResponseWithDefaults instantiates a new PhoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneResponseWithDefaults() *PhoneResponse {
	this := PhoneResponse{}
	return &this
}

// GetPkiPhoneID returns the PkiPhoneID field value
func (o *PhoneResponse) GetPkiPhoneID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPhoneID
}

// GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field value
// and a boolean to check if the value has been set.
func (o *PhoneResponse) GetPkiPhoneIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPhoneID, true
}

// SetPkiPhoneID sets field value
func (o *PhoneResponse) SetPkiPhoneID(v int32) {
	o.PkiPhoneID = v
}

// GetFkiPhonetypeID returns the FkiPhonetypeID field value
func (o *PhoneResponse) GetFkiPhonetypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiPhonetypeID
}

// GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field value
// and a boolean to check if the value has been set.
func (o *PhoneResponse) GetFkiPhonetypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiPhonetypeID, true
}

// SetFkiPhonetypeID sets field value
func (o *PhoneResponse) SetFkiPhonetypeID(v int32) {
	o.FkiPhonetypeID = v
}

// GetEPhoneType returns the EPhoneType field value if set, zero value otherwise.
// Deprecated
func (o *PhoneResponse) GetEPhoneType() FieldEPhoneType {
	if o == nil || IsNil(o.EPhoneType) {
		var ret FieldEPhoneType
		return ret
	}
	return *o.EPhoneType
}

// GetEPhoneTypeOk returns a tuple with the EPhoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneResponse) GetEPhoneTypeOk() (*FieldEPhoneType, bool) {
	if o == nil || IsNil(o.EPhoneType) {
		return nil, false
	}
	return o.EPhoneType, true
}

// HasEPhoneType returns a boolean if a field has been set.
func (o *PhoneResponse) HasEPhoneType() bool {
	if o != nil && !IsNil(o.EPhoneType) {
		return true
	}

	return false
}

// SetEPhoneType gets a reference to the given FieldEPhoneType and assigns it to the EPhoneType field.
// Deprecated
func (o *PhoneResponse) SetEPhoneType(v FieldEPhoneType) {
	o.EPhoneType = &v
}

// GetSPhoneE164 returns the SPhoneE164 field value if set, zero value otherwise.
func (o *PhoneResponse) GetSPhoneE164() string {
	if o == nil || IsNil(o.SPhoneE164) {
		var ret string
		return ret
	}
	return *o.SPhoneE164
}

// GetSPhoneE164Ok returns a tuple with the SPhoneE164 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneResponse) GetSPhoneE164Ok() (*string, bool) {
	if o == nil || IsNil(o.SPhoneE164) {
		return nil, false
	}
	return o.SPhoneE164, true
}

// HasSPhoneE164 returns a boolean if a field has been set.
func (o *PhoneResponse) HasSPhoneE164() bool {
	if o != nil && !IsNil(o.SPhoneE164) {
		return true
	}

	return false
}

// SetSPhoneE164 gets a reference to the given string and assigns it to the SPhoneE164 field.
func (o *PhoneResponse) SetSPhoneE164(v string) {
	o.SPhoneE164 = &v
}

// GetSPhoneExtension returns the SPhoneExtension field value if set, zero value otherwise.
func (o *PhoneResponse) GetSPhoneExtension() string {
	if o == nil || IsNil(o.SPhoneExtension) {
		var ret string
		return ret
	}
	return *o.SPhoneExtension
}

// GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneResponse) GetSPhoneExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneExtension) {
		return nil, false
	}
	return o.SPhoneExtension, true
}

// HasSPhoneExtension returns a boolean if a field has been set.
func (o *PhoneResponse) HasSPhoneExtension() bool {
	if o != nil && !IsNil(o.SPhoneExtension) {
		return true
	}

	return false
}

// SetSPhoneExtension gets a reference to the given string and assigns it to the SPhoneExtension field.
func (o *PhoneResponse) SetSPhoneExtension(v string) {
	o.SPhoneExtension = &v
}

func (o PhoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPhoneID"] = o.PkiPhoneID
	toSerialize["fkiPhonetypeID"] = o.FkiPhonetypeID
	if !IsNil(o.EPhoneType) {
		toSerialize["ePhoneType"] = o.EPhoneType
	}
	if !IsNil(o.SPhoneE164) {
		toSerialize["sPhoneE164"] = o.SPhoneE164
	}
	if !IsNil(o.SPhoneExtension) {
		toSerialize["sPhoneExtension"] = o.SPhoneExtension
	}
	return toSerialize, nil
}

func (o *PhoneResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPhoneID",
		"fkiPhonetypeID",
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

	varPhoneResponse := _PhoneResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneResponse)

	if err != nil {
		return err
	}

	*o = PhoneResponse(varPhoneResponse)

	return err
}

type NullablePhoneResponse struct {
	value *PhoneResponse
	isSet bool
}

func (v NullablePhoneResponse) Get() *PhoneResponse {
	return v.value
}

func (v *NullablePhoneResponse) Set(val *PhoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneResponse(val *PhoneResponse) *NullablePhoneResponse {
	return &NullablePhoneResponse{value: val, isSet: true}
}

func (v NullablePhoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


