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

// checks if the CorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorsResponse{}

// CorsResponse A Cors Object
type CorsResponse struct {
	// The unique ID of the Cors
	PkiCorsID int32 `json:"pkiCorsID"`
	// The unique ID of the Apikey
	FkiApikeyID int32 `json:"fkiApikeyID"`
	// The entryurl of the Cors
	SCorsEntryurl string `json:"sCorsEntryurl"`
}

// NewCorsResponse instantiates a new CorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorsResponse(pkiCorsID int32, fkiApikeyID int32, sCorsEntryurl string) *CorsResponse {
	this := CorsResponse{}
	this.PkiCorsID = pkiCorsID
	this.FkiApikeyID = fkiApikeyID
	this.SCorsEntryurl = sCorsEntryurl
	return &this
}

// NewCorsResponseWithDefaults instantiates a new CorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorsResponseWithDefaults() *CorsResponse {
	this := CorsResponse{}
	return &this
}

// GetPkiCorsID returns the PkiCorsID field value
func (o *CorsResponse) GetPkiCorsID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCorsID
}

// GetPkiCorsIDOk returns a tuple with the PkiCorsID field value
// and a boolean to check if the value has been set.
func (o *CorsResponse) GetPkiCorsIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCorsID, true
}

// SetPkiCorsID sets field value
func (o *CorsResponse) SetPkiCorsID(v int32) {
	o.PkiCorsID = v
}

// GetFkiApikeyID returns the FkiApikeyID field value
func (o *CorsResponse) GetFkiApikeyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiApikeyID
}

// GetFkiApikeyIDOk returns a tuple with the FkiApikeyID field value
// and a boolean to check if the value has been set.
func (o *CorsResponse) GetFkiApikeyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiApikeyID, true
}

// SetFkiApikeyID sets field value
func (o *CorsResponse) SetFkiApikeyID(v int32) {
	o.FkiApikeyID = v
}

// GetSCorsEntryurl returns the SCorsEntryurl field value
func (o *CorsResponse) GetSCorsEntryurl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCorsEntryurl
}

// GetSCorsEntryurlOk returns a tuple with the SCorsEntryurl field value
// and a boolean to check if the value has been set.
func (o *CorsResponse) GetSCorsEntryurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCorsEntryurl, true
}

// SetSCorsEntryurl sets field value
func (o *CorsResponse) SetSCorsEntryurl(v string) {
	o.SCorsEntryurl = v
}

func (o CorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCorsID"] = o.PkiCorsID
	toSerialize["fkiApikeyID"] = o.FkiApikeyID
	toSerialize["sCorsEntryurl"] = o.SCorsEntryurl
	return toSerialize, nil
}

type NullableCorsResponse struct {
	value *CorsResponse
	isSet bool
}

func (v NullableCorsResponse) Get() *CorsResponse {
	return v.value
}

func (v *NullableCorsResponse) Set(val *CorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorsResponse(val *CorsResponse) *NullableCorsResponse {
	return &NullableCorsResponse{value: val, isSet: true}
}

func (v NullableCorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


