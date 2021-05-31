/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.44
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// PhoneRequest A Phone Object
type PhoneRequest struct {
	// The unique ID of the Phonetype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Mobile| |4|Fax| |5|Pager| |6|Toll Free|
	FkiPhonetypeID int32 `json:"fkiPhonetypeID"`
	EPhoneType FieldEPhoneType `json:"ePhoneType"`
	// The region of the phone number. (For a North America Number only)  The region is the \"514\" section in this sample phone number: (514) 990-1516 x123
	SPhoneRegion *string `json:"sPhoneRegion,omitempty"`
	// The exchange of the phone number. (For a North America Number only)  The exchange is the \"990\" section in this sample phone number: (514) 990-1516 x123
	SPhoneExchange *string `json:"sPhoneExchange,omitempty"`
	// The number of the phone number. (For a North America Number only)  The number is the \"1516\" section in this sample phone number: (514) 990-1516 x123
	SPhoneNumber *string `json:"sPhoneNumber,omitempty"`
	// The international phone number. (For phone numbers outside of North)  Do not specify the \"011\" part of the phone number used to dial an international phone number from North America.  For example for this sample phone number \"+442071838750\", you would send \"442071838750\" without the \"+\" sign.
	SPhoneInternational *string `json:"sPhoneInternational,omitempty"`
	// The extension of the phone number.  The extension is the \"123\" section in this sample phone number: (514) 990-1516 x123.  It can also be used with international phone numbers
	SPhoneExtension *string `json:"sPhoneExtension,omitempty"`
}

// NewPhoneRequest instantiates a new PhoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneRequest(fkiPhonetypeID int32, ePhoneType FieldEPhoneType) *PhoneRequest {
	this := PhoneRequest{}
	this.FkiPhonetypeID = fkiPhonetypeID
	this.EPhoneType = ePhoneType
	return &this
}

// NewPhoneRequestWithDefaults instantiates a new PhoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneRequestWithDefaults() *PhoneRequest {
	this := PhoneRequest{}
	return &this
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
	if o == nil  {
		return nil, false
	}
	return &o.FkiPhonetypeID, true
}

// SetFkiPhonetypeID sets field value
func (o *PhoneRequest) SetFkiPhonetypeID(v int32) {
	o.FkiPhonetypeID = v
}

// GetEPhoneType returns the EPhoneType field value
func (o *PhoneRequest) GetEPhoneType() FieldEPhoneType {
	if o == nil {
		var ret FieldEPhoneType
		return ret
	}

	return o.EPhoneType
}

// GetEPhoneTypeOk returns a tuple with the EPhoneType field value
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetEPhoneTypeOk() (*FieldEPhoneType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EPhoneType, true
}

// SetEPhoneType sets field value
func (o *PhoneRequest) SetEPhoneType(v FieldEPhoneType) {
	o.EPhoneType = v
}

// GetSPhoneRegion returns the SPhoneRegion field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneRegion() string {
	if o == nil || o.SPhoneRegion == nil {
		var ret string
		return ret
	}
	return *o.SPhoneRegion
}

// GetSPhoneRegionOk returns a tuple with the SPhoneRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneRegionOk() (*string, bool) {
	if o == nil || o.SPhoneRegion == nil {
		return nil, false
	}
	return o.SPhoneRegion, true
}

// HasSPhoneRegion returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneRegion() bool {
	if o != nil && o.SPhoneRegion != nil {
		return true
	}

	return false
}

// SetSPhoneRegion gets a reference to the given string and assigns it to the SPhoneRegion field.
func (o *PhoneRequest) SetSPhoneRegion(v string) {
	o.SPhoneRegion = &v
}

// GetSPhoneExchange returns the SPhoneExchange field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneExchange() string {
	if o == nil || o.SPhoneExchange == nil {
		var ret string
		return ret
	}
	return *o.SPhoneExchange
}

// GetSPhoneExchangeOk returns a tuple with the SPhoneExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneExchangeOk() (*string, bool) {
	if o == nil || o.SPhoneExchange == nil {
		return nil, false
	}
	return o.SPhoneExchange, true
}

// HasSPhoneExchange returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneExchange() bool {
	if o != nil && o.SPhoneExchange != nil {
		return true
	}

	return false
}

// SetSPhoneExchange gets a reference to the given string and assigns it to the SPhoneExchange field.
func (o *PhoneRequest) SetSPhoneExchange(v string) {
	o.SPhoneExchange = &v
}

// GetSPhoneNumber returns the SPhoneNumber field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneNumber() string {
	if o == nil || o.SPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.SPhoneNumber
}

// GetSPhoneNumberOk returns a tuple with the SPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneNumberOk() (*string, bool) {
	if o == nil || o.SPhoneNumber == nil {
		return nil, false
	}
	return o.SPhoneNumber, true
}

// HasSPhoneNumber returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneNumber() bool {
	if o != nil && o.SPhoneNumber != nil {
		return true
	}

	return false
}

// SetSPhoneNumber gets a reference to the given string and assigns it to the SPhoneNumber field.
func (o *PhoneRequest) SetSPhoneNumber(v string) {
	o.SPhoneNumber = &v
}

// GetSPhoneInternational returns the SPhoneInternational field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneInternational() string {
	if o == nil || o.SPhoneInternational == nil {
		var ret string
		return ret
	}
	return *o.SPhoneInternational
}

// GetSPhoneInternationalOk returns a tuple with the SPhoneInternational field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneInternationalOk() (*string, bool) {
	if o == nil || o.SPhoneInternational == nil {
		return nil, false
	}
	return o.SPhoneInternational, true
}

// HasSPhoneInternational returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneInternational() bool {
	if o != nil && o.SPhoneInternational != nil {
		return true
	}

	return false
}

// SetSPhoneInternational gets a reference to the given string and assigns it to the SPhoneInternational field.
func (o *PhoneRequest) SetSPhoneInternational(v string) {
	o.SPhoneInternational = &v
}

// GetSPhoneExtension returns the SPhoneExtension field value if set, zero value otherwise.
func (o *PhoneRequest) GetSPhoneExtension() string {
	if o == nil || o.SPhoneExtension == nil {
		var ret string
		return ret
	}
	return *o.SPhoneExtension
}

// GetSPhoneExtensionOk returns a tuple with the SPhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneRequest) GetSPhoneExtensionOk() (*string, bool) {
	if o == nil || o.SPhoneExtension == nil {
		return nil, false
	}
	return o.SPhoneExtension, true
}

// HasSPhoneExtension returns a boolean if a field has been set.
func (o *PhoneRequest) HasSPhoneExtension() bool {
	if o != nil && o.SPhoneExtension != nil {
		return true
	}

	return false
}

// SetSPhoneExtension gets a reference to the given string and assigns it to the SPhoneExtension field.
func (o *PhoneRequest) SetSPhoneExtension(v string) {
	o.SPhoneExtension = &v
}

func (o PhoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiPhonetypeID"] = o.FkiPhonetypeID
	}
	if true {
		toSerialize["ePhoneType"] = o.EPhoneType
	}
	if o.SPhoneRegion != nil {
		toSerialize["sPhoneRegion"] = o.SPhoneRegion
	}
	if o.SPhoneExchange != nil {
		toSerialize["sPhoneExchange"] = o.SPhoneExchange
	}
	if o.SPhoneNumber != nil {
		toSerialize["sPhoneNumber"] = o.SPhoneNumber
	}
	if o.SPhoneInternational != nil {
		toSerialize["sPhoneInternational"] = o.SPhoneInternational
	}
	if o.SPhoneExtension != nil {
		toSerialize["sPhoneExtension"] = o.SPhoneExtension
	}
	return json.Marshal(toSerialize)
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


