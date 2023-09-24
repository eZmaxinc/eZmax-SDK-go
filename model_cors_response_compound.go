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

// checks if the CorsResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorsResponseCompound{}

// CorsResponseCompound A Cors Object
type CorsResponseCompound struct {
	// The unique ID of the Cors
	PkiCorsID int32 `json:"pkiCorsID"`
	// The unique ID of the Apikey
	FkiApikeyID int32 `json:"fkiApikeyID"`
	// The entryurl of the Cors
	SCorsEntryurl string `json:"sCorsEntryurl"`
}

// NewCorsResponseCompound instantiates a new CorsResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorsResponseCompound(pkiCorsID int32, fkiApikeyID int32, sCorsEntryurl string) *CorsResponseCompound {
	this := CorsResponseCompound{}
	this.PkiCorsID = pkiCorsID
	this.FkiApikeyID = fkiApikeyID
	this.SCorsEntryurl = sCorsEntryurl
	return &this
}

// NewCorsResponseCompoundWithDefaults instantiates a new CorsResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsResponseCompoundWithDefaults() *CorsResponseCompound {
	this := CorsResponseCompound{}
	return &this
}

// GetPkiCorsID returns the PkiCorsID field value
func (o *CorsResponseCompound) GetPkiCorsID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCorsID
}

// GetPkiCorsIDOk returns a tuple with the PkiCorsID field value
// and a boolean to check if the value has been set.
func (o *CorsResponseCompound) GetPkiCorsIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCorsID, true
}

// SetPkiCorsID sets field value
func (o *CorsResponseCompound) SetPkiCorsID(v int32) {
	o.PkiCorsID = v
}

// GetFkiApikeyID returns the FkiApikeyID field value
func (o *CorsResponseCompound) GetFkiApikeyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value
// and a boolean to check if the value has been set.
func (o *CorsResponseCompound) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiApikeyID, true
}

// SetFkiApikeyID sets field value
func (o *CorsResponseCompound) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = v
}

// GetSCorsEntryurl returns the SCorsEntryurl field value
func (o *CorsResponseCompound) GetSCorsEntryurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCorsEntryurl
}

// GetSCorsEntryurlOk returns a tuple with the SCorsEntryurl field value
// and a boolean to check if the value has been set.
func (o *CorsResponseCompound) GetSCorsEntryurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCorsEntryurl, true
}

// SetSCorsEntryurl sets field value
func (o *CorsResponseCompound) SetSCorsEntryurl(v string) {
	o.SCorsEntryurl = v
}

func (o CorsResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorsResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCorsID"] = o.PkiCorsID
	toSerialize["fkiApikeyID"] = o.FkiApikeyID
	toSerialize["sCorsEntryurl"] = o.SCorsEntryurl
	return toSerialize, nil
}

type NullableCorsResponseCompound struct {
	value *CorsResponseCompound
	isSet bool
}

func (v NullableCorsResponseCompound) Get() *CorsResponseCompound {
	return v.value
}

func (v *NullableCorsResponseCompound) Set(val *CorsResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCorsResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCorsResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorsResponseCompound(val *CorsResponseCompound) *NullableCorsResponseCompound {
	return &NullableCorsResponseCompound{value: val, isSet: true}
}

func (v NullableCorsResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorsResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

