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

// checks if the GlobalEzmaxcustomerGetConfigurationV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalEzmaxcustomerGetConfigurationV1Response{}

// GlobalEzmaxcustomerGetConfigurationV1Response Response for GET /1/ezmaxcustomer/{pksEzmaxcustomerCode}/getConfiguration
type GlobalEzmaxcustomerGetConfigurationV1Response struct {
	// The region code
	SInfrastructureregionCode string `json:"sInfrastructureregionCode"`
	// The region code
	SInfrastructureregionCodeWeb string `json:"sInfrastructureregionCodeWeb"`
	// The environment type Description
	SInfrastructureenvironmenttypeDescription string `json:"sInfrastructureenvironmenttypeDescription"`
	// The ID of the client in Cognito
	SCognitoClientIDExternal *string `json:"sCognitoClientIDExternal,omitempty"`
	// The ID of the client in Cognito
	SCognitoClientIDEzmaxpublic string `json:"sCognitoClientIDEzmaxpublic"`
}

type _GlobalEzmaxcustomerGetConfigurationV1Response GlobalEzmaxcustomerGetConfigurationV1Response

// NewGlobalEzmaxcustomerGetConfigurationV1Response instantiates a new GlobalEzmaxcustomerGetConfigurationV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalEzmaxcustomerGetConfigurationV1Response(sInfrastructureregionCode string, sInfrastructureregionCodeWeb string, sInfrastructureenvironmenttypeDescription string, sCognitoClientIDEzmaxpublic string) *GlobalEzmaxcustomerGetConfigurationV1Response {
	this := GlobalEzmaxcustomerGetConfigurationV1Response{}
	this.SInfrastructureregionCode = sInfrastructureregionCode
	this.SInfrastructureregionCodeWeb = sInfrastructureregionCodeWeb
	this.SInfrastructureenvironmenttypeDescription = sInfrastructureenvironmenttypeDescription
	this.SCognitoClientIDEzmaxpublic = sCognitoClientIDEzmaxpublic
	return &this
}

// NewGlobalEzmaxcustomerGetConfigurationV1ResponseWithDefaults instantiates a new GlobalEzmaxcustomerGetConfigurationV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalEzmaxcustomerGetConfigurationV1ResponseWithDefaults() *GlobalEzmaxcustomerGetConfigurationV1Response {
	this := GlobalEzmaxcustomerGetConfigurationV1Response{}
	return &this
}

// GetSInfrastructureregionCode returns the SInfrastructureregionCode field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureregionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SInfrastructureregionCode
}

// GetSInfrastructureregionCodeOk returns a tuple with the SInfrastructureregionCode field value
// and a boolean to check if the value has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureregionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SInfrastructureregionCode, true
}

// SetSInfrastructureregionCode sets field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) SetSInfrastructureregionCode(v string) {
	o.SInfrastructureregionCode = v
}

// GetSInfrastructureregionCodeWeb returns the SInfrastructureregionCodeWeb field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureregionCodeWeb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SInfrastructureregionCodeWeb
}

// GetSInfrastructureregionCodeWebOk returns a tuple with the SInfrastructureregionCodeWeb field value
// and a boolean to check if the value has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureregionCodeWebOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SInfrastructureregionCodeWeb, true
}

// SetSInfrastructureregionCodeWeb sets field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) SetSInfrastructureregionCodeWeb(v string) {
	o.SInfrastructureregionCodeWeb = v
}

// GetSInfrastructureenvironmenttypeDescription returns the SInfrastructureenvironmenttypeDescription field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureenvironmenttypeDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SInfrastructureenvironmenttypeDescription
}

// GetSInfrastructureenvironmenttypeDescriptionOk returns a tuple with the SInfrastructureenvironmenttypeDescription field value
// and a boolean to check if the value has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSInfrastructureenvironmenttypeDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SInfrastructureenvironmenttypeDescription, true
}

// SetSInfrastructureenvironmenttypeDescription sets field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) SetSInfrastructureenvironmenttypeDescription(v string) {
	o.SInfrastructureenvironmenttypeDescription = v
}

// GetSCognitoClientIDExternal returns the SCognitoClientIDExternal field value if set, zero value otherwise.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSCognitoClientIDExternal() string {
	if o == nil || IsNil(o.SCognitoClientIDExternal) {
		var ret string
		return ret
	}
	return *o.SCognitoClientIDExternal
}

// GetSCognitoClientIDExternalOk returns a tuple with the SCognitoClientIDExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSCognitoClientIDExternalOk() (*string, bool) {
	if o == nil || IsNil(o.SCognitoClientIDExternal) {
		return nil, false
	}
	return o.SCognitoClientIDExternal, true
}

// HasSCognitoClientIDExternal returns a boolean if a field has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) HasSCognitoClientIDExternal() bool {
	if o != nil && !IsNil(o.SCognitoClientIDExternal) {
		return true
	}

	return false
}

// SetSCognitoClientIDExternal gets a reference to the given string and assigns it to the SCognitoClientIDExternal field.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) SetSCognitoClientIDExternal(v string) {
	o.SCognitoClientIDExternal = &v
}

// GetSCognitoClientIDEzmaxpublic returns the SCognitoClientIDEzmaxpublic field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSCognitoClientIDEzmaxpublic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCognitoClientIDEzmaxpublic
}

// GetSCognitoClientIDEzmaxpublicOk returns a tuple with the SCognitoClientIDEzmaxpublic field value
// and a boolean to check if the value has been set.
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) GetSCognitoClientIDEzmaxpublicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCognitoClientIDEzmaxpublic, true
}

// SetSCognitoClientIDEzmaxpublic sets field value
func (o *GlobalEzmaxcustomerGetConfigurationV1Response) SetSCognitoClientIDEzmaxpublic(v string) {
	o.SCognitoClientIDEzmaxpublic = v
}

func (o GlobalEzmaxcustomerGetConfigurationV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalEzmaxcustomerGetConfigurationV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sInfrastructureregionCode"] = o.SInfrastructureregionCode
	toSerialize["sInfrastructureregionCodeWeb"] = o.SInfrastructureregionCodeWeb
	toSerialize["sInfrastructureenvironmenttypeDescription"] = o.SInfrastructureenvironmenttypeDescription
	if !IsNil(o.SCognitoClientIDExternal) {
		toSerialize["sCognitoClientIDExternal"] = o.SCognitoClientIDExternal
	}
	toSerialize["sCognitoClientIDEzmaxpublic"] = o.SCognitoClientIDEzmaxpublic
	return toSerialize, nil
}

func (o *GlobalEzmaxcustomerGetConfigurationV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sInfrastructureregionCode",
		"sInfrastructureregionCodeWeb",
		"sInfrastructureenvironmenttypeDescription",
		"sCognitoClientIDEzmaxpublic",
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

	varGlobalEzmaxcustomerGetConfigurationV1Response := _GlobalEzmaxcustomerGetConfigurationV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlobalEzmaxcustomerGetConfigurationV1Response)

	if err != nil {
		return err
	}

	*o = GlobalEzmaxcustomerGetConfigurationV1Response(varGlobalEzmaxcustomerGetConfigurationV1Response)

	return err
}

type NullableGlobalEzmaxcustomerGetConfigurationV1Response struct {
	value *GlobalEzmaxcustomerGetConfigurationV1Response
	isSet bool
}

func (v NullableGlobalEzmaxcustomerGetConfigurationV1Response) Get() *GlobalEzmaxcustomerGetConfigurationV1Response {
	return v.value
}

func (v *NullableGlobalEzmaxcustomerGetConfigurationV1Response) Set(val *GlobalEzmaxcustomerGetConfigurationV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalEzmaxcustomerGetConfigurationV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalEzmaxcustomerGetConfigurationV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalEzmaxcustomerGetConfigurationV1Response(val *GlobalEzmaxcustomerGetConfigurationV1Response) *NullableGlobalEzmaxcustomerGetConfigurationV1Response {
	return &NullableGlobalEzmaxcustomerGetConfigurationV1Response{value: val, isSet: true}
}

func (v NullableGlobalEzmaxcustomerGetConfigurationV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalEzmaxcustomerGetConfigurationV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


