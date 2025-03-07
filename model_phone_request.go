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

// checks if the PhoneRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneRequest{}

// PhoneRequest A Phone Object
type PhoneRequest struct {
	// The unique ID of the Phone.
	PkiPhoneID *int32 `json:"pkiPhoneID,omitempty"`
	// The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free|
	FkiPhonetypeID int32 `json:"fkiPhonetypeID"`
	// Deprecated
	EPhoneType *FieldEPhoneType `json:"ePhoneType,omitempty"`
	// The region of the phone number. (For a North America Number only)  The region is the \"514\" section in this sample phone number: (514) 990-1516 x123
	// Deprecated
	SPhoneRegion *string `json:"sPhoneRegion,omitempty"`
	// The exchange of the phone number. (For a North America Number only)  The exchange is the \"990\" section in this sample phone number: (514) 990-1516 x123
	// Deprecated
	SPhoneExchange *string `json:"sPhoneExchange,omitempty"`
	// The number of the phone number. (For a North America Number only)  The number is the \"1516\" section in this sample phone number: (514) 990-1516 x123
	// Deprecated
	SPhoneNumber *string `json:"sPhoneNumber,omitempty"`
	// The international phone number.
	// Deprecated
	SPhoneInternational *string `json:"sPhoneInternational,omitempty"`
	// The extension of the phone number.  The extension is the \"123\" section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers
	SPhoneExtension *string `json:"sPhoneExtension,omitempty"`
	// A phone number in E.164 Format
	SPhoneE164 *string `json:"sPhoneE164,omitempty" validate:"regexp=^\\\\+[1-9]\\\\d{1,14}$"`
}

type _PhoneRequest PhoneRequest

// NewPhoneRequest instantiates a new PhoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneRequest(fkiPhonetypeID int32) *PhoneRequest {
	this := PhoneRequest{}
	this.FkiPhonetypeID = fkiPhonetypeID
	return &this
}

// NewPhoneRequestWithDefaults instantiates a new PhoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneRequestWithDefaults() *PhoneRequest {
	this := PhoneRequest{}
	return &this
}

// GetPkiPhoneID returns the PkiPhoneID field value if set, zero value otherwise.
func (o *PhoneRequest) GetPkiPhoneID() int32 {
	if o == nil || IsNil(o.PkiPhoneID) {
		var ret int32
		return ret
	}
	return *o.PkiPhoneID
}

// GetPkiPhoneIDOk returns a tuple with the PkiPhoneID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetPkiPhoneIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiPhoneID) {
		return nil, false
	}
	return o.PkiPhoneID, true
}

// HasPkiPhoneID returns a boolean if a field has been set.
func (o *PhoneRequest) HasPkiPhoneID() bool {
	if o != nil && !IsNil(o.PkiPhoneID) {
		return true
	}

	return false
}

// SetPkiPhoneID gets a reference to the given int32 and assigns it to the PkiPhoneID field.
func (o *PhoneRequest) SetPkiPhoneID(v int32) {
	o.PkiPhoneID = &v
}

// GetFkiPhonetypeID returns the FkiPhonetypeID field value
func (o *PhoneRequest) GetFkiPhonetypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiPhonetypeID
}

// GetFkiPhonetypeIDOk returns a tuple with the FkiPhonetypeID field value
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetFkiPhonetypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiPhonetypeID, true
}

// SetFkiPhonetypeID sets field value
func (o *PhoneRequest) SetFkiPhonetypeID(v int32) {
	o.FkiPhonetypeID = v
}

// GetEPhoneType returns the EPhoneType field value if set, zero value otherwise.
// Deprecated
func (o *PhoneRequest) GetEPhoneType() FieldEPhoneType {
	if o == nil || IsNil(o.EPhoneType) {
		var ret FieldEPhoneType
		return ret
	}
	return *o.EPhoneType
}

// GetEPhoneTypeOk returns a tuple with the EPhoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneRequest) GetEPhoneTypeOk() (*FieldEPhoneType, bool) {
	if o == nil || IsNil(o.EPhoneType) {
		return nil, false
	}
	return o.EPhoneType, true
}

// HasEPhoneType returns a boolean if a field has been set.
func (o *PhoneRequest) HasEPhoneType() bool {
	if o != nil && !IsNil(o.EPhoneType) {
		return true
	}

	return false
}

// SetEPhoneType gets a reference to the given FieldEPhoneType and assigns it to the EPhoneType field.
// Deprecated
func (o *PhoneRequest) SetEPhoneType(v FieldEPhoneType) {
	o.EPhoneType = &v
}

// GetSPhoneRegion returns the SPhoneRegion field value if set, zero value otherwise.
// Deprecated
func (o *PhoneRequest) GetSPhoneRegion() string {
	if o == nil || IsNil(o.SPhoneRegion) {
		var ret string
		return ret
	}
	return *o.SPhoneRegion
}

// GetSPhoneRegionOk returns a tuple with the SPhoneRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneRequest) GetSPhoneRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneRegion) {
		return nil, false
	}
	return o.SPhoneRegion, true
}

// HasSPhoneRegion returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneRegion() bool {
	if o != nil && !IsNil(o.SPhoneRegion) {
		return true
	}

	return false
}

// SetSPhoneRegion gets a reference to the given string and assigns it to the SPhoneRegion field.
// Deprecated
func (o *PhoneRequest) SetSPhoneRegion(v string) {
	o.SPhoneRegion = &v
}

// GetSPhoneExchange returns the SPhoneExchange field value if set, zero value otherwise.
// Deprecated
func (o *PhoneRequest) GetSPhoneExchange() string {
	if o == nil || IsNil(o.SPhoneExchange) {
		var ret string
		return ret
	}
	return *o.SPhoneExchange
}

// GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneRequest) GetSPhoneExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneExchange) {
		return nil, false
	}
	return o.SPhoneExchange, true
}

// HasSPhoneExchange returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneExchange() bool {
	if o != nil && !IsNil(o.SPhoneExchange) {
		return true
	}

	return false
}

// SetSPhoneExchange gets a reference to the given string and assigns it to the SPhoneExchange field.
// Deprecated
func (o *PhoneRequest) SetSPhoneExchange(v string) {
	o.SPhoneExchange = &v
}

// GetSPhoneNumber returns the SPhoneNumber field value if set, zero value otherwise.
// Deprecated
func (o *PhoneRequest) GetSPhoneNumber() string {
	if o == nil || IsNil(o.SPhoneNumber) {
		var ret string
		return ret
	}
	return *o.SPhoneNumber
}

// GetSPhoneNumberOk returns a tuple with the SPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneRequest) GetSPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneNumber) {
		return nil, false
	}
	return o.SPhoneNumber, true
}

// HasSPhoneNumber returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneNumber() bool {
	if o != nil && !IsNil(o.SPhoneNumber) {
		return true
	}

	return false
}

// SetSPhoneNumber gets a reference to the given string and assigns it to the SPhoneNumber field.
// Deprecated
func (o *PhoneRequest) SetSPhoneNumber(v string) {
	o.SPhoneNumber = &v
}

// GetSPhoneInternational returns the SPhoneInternational field value if set, zero value otherwise.
// Deprecated
func (o *PhoneRequest) GetSPhoneInternational() string {
	if o == nil || IsNil(o.SPhoneInternational) {
		var ret string
		return ret
	}
	return *o.SPhoneInternational
}

// GetSPhoneInternationalOk returns a tuple with the SPhoneInternational field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhoneRequest) GetSPhoneInternationalOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneInternational) {
		return nil, false
	}
	return o.SPhoneInternational, true
}

// HasSPhoneInternational returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneInternational() bool {
	if o != nil && !IsNil(o.SPhoneInternational) {
		return true
	}

	return false
}

// SetSPhoneInternational gets a reference to the given string and assigns it to the SPhoneInternational field.
// Deprecated
func (o *PhoneRequest) SetSPhoneInternational(v string) {
	o.SPhoneInternational = &v
}

// GetSPhoneExtension returns the SPhoneExtension field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneExtension() string {
	if o == nil || IsNil(o.SPhoneExtension) {
		var ret string
		return ret
	}
	return *o.SPhoneExtension
}

// GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.SPhoneExtension) {
		return nil, false
	}
	return o.SPhoneExtension, true
}

// HasSPhoneExtension returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneExtension() bool {
	if o != nil && !IsNil(o.SPhoneExtension) {
		return true
	}

	return false
}

// SetSPhoneExtension gets a reference to the given string and assigns it to the SPhoneExtension field.
func (o *PhoneRequest) SetSPhoneExtension(v string) {
	o.SPhoneExtension = &v
}

// GetSPhoneE164 returns the SPhoneE164 field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneE164() string {
	if o == nil || IsNil(o.SPhoneE164) {
		var ret string
		return ret
	}
	return *o.SPhoneE164
}

// GetSPhoneE164Ok returns a tuple with the SPhoneE164 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneE164Ok() (*string, bool) {
	if o == nil || IsNil(o.SPhoneE164) {
		return nil, false
	}
	return o.SPhoneE164, true
}

// HasSPhoneE164 returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneE164() bool {
	if o != nil && !IsNil(o.SPhoneE164) {
		return true
	}

	return false
}

// SetSPhoneE164 gets a reference to the given string and assigns it to the SPhoneE164 field.
func (o *PhoneRequest) SetSPhoneE164(v string) {
	o.SPhoneE164 = &v
}

func (o PhoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiPhoneID) {
		toSerialize["pkiPhoneID"] = o.PkiPhoneID
	}
	toSerialize["fkiPhonetypeID"] = o.FkiPhonetypeID
	if !IsNil(o.EPhoneType) {
		toSerialize["ePhoneType"] = o.EPhoneType
	}
	if !IsNil(o.SPhoneRegion) {
		toSerialize["sPhoneRegion"] = o.SPhoneRegion
	}
	if !IsNil(o.SPhoneExchange) {
		toSerialize["sPhoneExchange"] = o.SPhoneExchange
	}
	if !IsNil(o.SPhoneNumber) {
		toSerialize["sPhoneNumber"] = o.SPhoneNumber
	}
	if !IsNil(o.SPhoneInternational) {
		toSerialize["sPhoneInternational"] = o.SPhoneInternational
	}
	if !IsNil(o.SPhoneExtension) {
		toSerialize["sPhoneExtension"] = o.SPhoneExtension
	}
	if !IsNil(o.SPhoneE164) {
		toSerialize["sPhoneE164"] = o.SPhoneE164
	}
	return toSerialize, nil
}

func (o *PhoneRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPhoneRequest := _PhoneRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneRequest)

	if err != nil {
		return err
	}

	*o = PhoneRequest(varPhoneRequest)

	return err
}

type NullablePhoneRequest struct {
	value *PhoneRequest
	isSet bool
}

func (v NullablePhoneRequest) Get() *PhoneRequest {
	return v.value
}

func (v *NullablePhoneRequest) Set(val *PhoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneRequest(val *PhoneRequest) *NullablePhoneRequest {
	return &NullablePhoneRequest{value: val, isSet: true}
}

func (v NullablePhoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


