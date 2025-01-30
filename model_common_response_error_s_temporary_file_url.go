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

// checks if the CommonResponseErrorSTemporaryFileUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseErrorSTemporaryFileUrl{}

// CommonResponseErrorSTemporaryFileUrl Generic Error Message
type CommonResponseErrorSTemporaryFileUrl struct {
	// The message giving details about the error
	SErrorMessage string `json:"sErrorMessage" validate:"regexp=^.{0,500}$"`
	EErrorCode FieldEErrorCode `json:"eErrorCode"`
	// More error message detail
	ASErrorMessagedetail []string `json:"a_sErrorMessagedetail,omitempty"`
	// The Temporary File Url of the document that was uploaded. That url can be reused instead of uploading the file again.
	STemporaryFileUrl *string `json:"sTemporaryFileUrl,omitempty" validate:"regexp=^(https|http):\\/\\/[^\\\\s\\/$.?#].[^\\\\s]*$"`
}

type _CommonResponseErrorSTemporaryFileUrl CommonResponseErrorSTemporaryFileUrl

// NewCommonResponseErrorSTemporaryFileUrl instantiates a new CommonResponseErrorSTemporaryFileUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseErrorSTemporaryFileUrl(sErrorMessage string, eErrorCode FieldEErrorCode) *CommonResponseErrorSTemporaryFileUrl {
	this := CommonResponseErrorSTemporaryFileUrl{}
	this.SErrorMessage = sErrorMessage
	this.EErrorCode = eErrorCode
	return &this
}

// NewCommonResponseErrorSTemporaryFileUrlWithDefaults instantiates a new CommonResponseErrorSTemporaryFileUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseErrorSTemporaryFileUrlWithDefaults() *CommonResponseErrorSTemporaryFileUrl {
	this := CommonResponseErrorSTemporaryFileUrl{}
	return &this
}

// GetSErrorMessage returns the SErrorMessage field value
func (o *CommonResponseErrorSTemporaryFileUrl) GetSErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SErrorMessage
}

// GetSErrorMessageOk returns a tuple with the SErrorMessage field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) GetSErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SErrorMessage, true
}

// SetSErrorMessage sets field value
func (o *CommonResponseErrorSTemporaryFileUrl) SetSErrorMessage(v string) {
	o.SErrorMessage = v
}

// GetEErrorCode returns the EErrorCode field value
func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCode() FieldEErrorCode {
	if o == nil {
		var ret FieldEErrorCode
		return ret
	}

	return o.EErrorCode
}

// GetEErrorCodeOk returns a tuple with the EErrorCode field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) GetEErrorCodeOk() (*FieldEErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EErrorCode, true
}

// SetEErrorCode sets field value
func (o *CommonResponseErrorSTemporaryFileUrl) SetEErrorCode(v FieldEErrorCode) {
	o.EErrorCode = v
}

// GetASErrorMessagedetail returns the ASErrorMessagedetail field value if set, zero value otherwise.
func (o *CommonResponseErrorSTemporaryFileUrl) GetASErrorMessagedetail() []string {
	if o == nil || IsNil(o.ASErrorMessagedetail) {
		var ret []string
		return ret
	}
	return o.ASErrorMessagedetail
}

// GetASErrorMessagedetailOk returns a tuple with the ASErrorMessagedetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) GetASErrorMessagedetailOk() ([]string, bool) {
	if o == nil || IsNil(o.ASErrorMessagedetail) {
		return nil, false
	}
	return o.ASErrorMessagedetail, true
}

// HasASErrorMessagedetail returns a boolean if a field has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) HasASErrorMessagedetail() bool {
	if o != nil && !IsNil(o.ASErrorMessagedetail) {
		return true
	}

	return false
}

// SetASErrorMessagedetail gets a reference to the given []string and assigns it to the ASErrorMessagedetail field.
func (o *CommonResponseErrorSTemporaryFileUrl) SetASErrorMessagedetail(v []string) {
	o.ASErrorMessagedetail = v
}

// GetSTemporaryFileUrl returns the STemporaryFileUrl field value if set, zero value otherwise.
func (o *CommonResponseErrorSTemporaryFileUrl) GetSTemporaryFileUrl() string {
	if o == nil || IsNil(o.STemporaryFileUrl) {
		var ret string
		return ret
	}
	return *o.STemporaryFileUrl
}

// GetSTemporaryFileUrlOk returns a tuple with the STemporaryFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) GetSTemporaryFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.STemporaryFileUrl) {
		return nil, false
	}
	return o.STemporaryFileUrl, true
}

// HasSTemporaryFileUrl returns a boolean if a field has been set.
func (o *CommonResponseErrorSTemporaryFileUrl) HasSTemporaryFileUrl() bool {
	if o != nil && !IsNil(o.STemporaryFileUrl) {
		return true
	}

	return false
}

// SetSTemporaryFileUrl gets a reference to the given string and assigns it to the STemporaryFileUrl field.
func (o *CommonResponseErrorSTemporaryFileUrl) SetSTemporaryFileUrl(v string) {
	o.STemporaryFileUrl = &v
}

func (o CommonResponseErrorSTemporaryFileUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseErrorSTemporaryFileUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sErrorMessage"] = o.SErrorMessage
	toSerialize["eErrorCode"] = o.EErrorCode
	if !IsNil(o.ASErrorMessagedetail) {
		toSerialize["a_sErrorMessagedetail"] = o.ASErrorMessagedetail
	}
	if !IsNil(o.STemporaryFileUrl) {
		toSerialize["sTemporaryFileUrl"] = o.STemporaryFileUrl
	}
	return toSerialize, nil
}

func (o *CommonResponseErrorSTemporaryFileUrl) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sErrorMessage",
		"eErrorCode",
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

	varCommonResponseErrorSTemporaryFileUrl := _CommonResponseErrorSTemporaryFileUrl{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseErrorSTemporaryFileUrl)

	if err != nil {
		return err
	}

	*o = CommonResponseErrorSTemporaryFileUrl(varCommonResponseErrorSTemporaryFileUrl)

	return err
}

type NullableCommonResponseErrorSTemporaryFileUrl struct {
	value *CommonResponseErrorSTemporaryFileUrl
	isSet bool
}

func (v NullableCommonResponseErrorSTemporaryFileUrl) Get() *CommonResponseErrorSTemporaryFileUrl {
	return v.value
}

func (v *NullableCommonResponseErrorSTemporaryFileUrl) Set(val *CommonResponseErrorSTemporaryFileUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseErrorSTemporaryFileUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseErrorSTemporaryFileUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseErrorSTemporaryFileUrl(val *CommonResponseErrorSTemporaryFileUrl) *NullableCommonResponseErrorSTemporaryFileUrl {
	return &NullableCommonResponseErrorSTemporaryFileUrl{value: val, isSet: true}
}

func (v NullableCommonResponseErrorSTemporaryFileUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseErrorSTemporaryFileUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


