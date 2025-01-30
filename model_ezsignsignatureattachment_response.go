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

// checks if the EzsignsignatureattachmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureattachmentResponse{}

// EzsignsignatureattachmentResponse An Ezsignsignatureattachment Object
type EzsignsignatureattachmentResponse struct {
	// The unique ID of the Ezsignsignatureattachment
	PkiEzsignsignatureattachmentID int32 `json:"pkiEzsignsignatureattachmentID"`
	// The unique ID of the Ezsignsignature
	FkiEzsignsignatureID int32 `json:"fkiEzsignsignatureID"`
	// The md5 of the Ezsignsignatureattachment
	BinEzsignsignatureattachmentMD5 string `json:"binEzsignsignatureattachmentMD5"`
	// The name of the Ezsignsignatureattachment
	SEzsignsignatureattachmentName string `json:"sEzsignsignatureattachmentName" validate:"regexp=^.{0,75}$"`
	// The Url to the requested document.  Url will expire after 3 hours.
	SDownloadUrl string `json:"sDownloadUrl"`
}

type _EzsignsignatureattachmentResponse EzsignsignatureattachmentResponse

// NewEzsignsignatureattachmentResponse instantiates a new EzsignsignatureattachmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureattachmentResponse(pkiEzsignsignatureattachmentID int32, fkiEzsignsignatureID int32, binEzsignsignatureattachmentMD5 string, sEzsignsignatureattachmentName string, sDownloadUrl string) *EzsignsignatureattachmentResponse {
	this := EzsignsignatureattachmentResponse{}
	this.PkiEzsignsignatureattachmentID = pkiEzsignsignatureattachmentID
	this.FkiEzsignsignatureID = fkiEzsignsignatureID
	this.BinEzsignsignatureattachmentMD5 = binEzsignsignatureattachmentMD5
	this.SEzsignsignatureattachmentName = sEzsignsignatureattachmentName
	this.SDownloadUrl = sDownloadUrl
	return &this
}

// NewEzsignsignatureattachmentResponseWithDefaults instantiates a new EzsignsignatureattachmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureattachmentResponseWithDefaults() *EzsignsignatureattachmentResponse {
	this := EzsignsignatureattachmentResponse{}
	return &this
}

// GetPkiEzsignsignatureattachmentID returns the PkiEzsignsignatureattachmentID field value
func (o *EzsignsignatureattachmentResponse) GetPkiEzsignsignatureattachmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignsignatureattachmentID
}

// GetPkiEzsignsignatureattachmentIDOk returns a tuple with the PkiEzsignsignatureattachmentID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureattachmentResponse) GetPkiEzsignsignatureattachmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignsignatureattachmentID, true
}

// SetPkiEzsignsignatureattachmentID sets field value
func (o *EzsignsignatureattachmentResponse) SetPkiEzsignsignatureattachmentID(v int32) {
	o.PkiEzsignsignatureattachmentID = v
}

// GetFkiEzsignsignatureID returns the FkiEzsignsignatureID field value
func (o *EzsignsignatureattachmentResponse) GetFkiEzsignsignatureID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignsignatureID
}

// GetFkiEzsignsignatureIDOk returns a tuple with the FkiEzsignsignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureattachmentResponse) GetFkiEzsignsignatureIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignsignatureID, true
}

// SetFkiEzsignsignatureID sets field value
func (o *EzsignsignatureattachmentResponse) SetFkiEzsignsignatureID(v int32) {
	o.FkiEzsignsignatureID = v
}

// GetBinEzsignsignatureattachmentMD5 returns the BinEzsignsignatureattachmentMD5 field value
func (o *EzsignsignatureattachmentResponse) GetBinEzsignsignatureattachmentMD5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinEzsignsignatureattachmentMD5
}

// GetBinEzsignsignatureattachmentMD5Ok returns a tuple with the BinEzsignsignatureattachmentMD5 field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureattachmentResponse) GetBinEzsignsignatureattachmentMD5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BinEzsignsignatureattachmentMD5, true
}

// SetBinEzsignsignatureattachmentMD5 sets field value
func (o *EzsignsignatureattachmentResponse) SetBinEzsignsignatureattachmentMD5(v string) {
	o.BinEzsignsignatureattachmentMD5 = v
}

// GetSEzsignsignatureattachmentName returns the SEzsignsignatureattachmentName field value
func (o *EzsignsignatureattachmentResponse) GetSEzsignsignatureattachmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignsignatureattachmentName
}

// GetSEzsignsignatureattachmentNameOk returns a tuple with the SEzsignsignatureattachmentName field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureattachmentResponse) GetSEzsignsignatureattachmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignsignatureattachmentName, true
}

// SetSEzsignsignatureattachmentName sets field value
func (o *EzsignsignatureattachmentResponse) SetSEzsignsignatureattachmentName(v string) {
	o.SEzsignsignatureattachmentName = v
}

// GetSDownloadUrl returns the SDownloadUrl field value
func (o *EzsignsignatureattachmentResponse) GetSDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDownloadUrl
}

// GetSDownloadUrlOk returns a tuple with the SDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureattachmentResponse) GetSDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDownloadUrl, true
}

// SetSDownloadUrl sets field value
func (o *EzsignsignatureattachmentResponse) SetSDownloadUrl(v string) {
	o.SDownloadUrl = v
}

func (o EzsignsignatureattachmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureattachmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignsignatureattachmentID"] = o.PkiEzsignsignatureattachmentID
	toSerialize["fkiEzsignsignatureID"] = o.FkiEzsignsignatureID
	toSerialize["binEzsignsignatureattachmentMD5"] = o.BinEzsignsignatureattachmentMD5
	toSerialize["sEzsignsignatureattachmentName"] = o.SEzsignsignatureattachmentName
	toSerialize["sDownloadUrl"] = o.SDownloadUrl
	return toSerialize, nil
}

func (o *EzsignsignatureattachmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignsignatureattachmentID",
		"fkiEzsignsignatureID",
		"binEzsignsignatureattachmentMD5",
		"sEzsignsignatureattachmentName",
		"sDownloadUrl",
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

	varEzsignsignatureattachmentResponse := _EzsignsignatureattachmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignatureattachmentResponse)

	if err != nil {
		return err
	}

	*o = EzsignsignatureattachmentResponse(varEzsignsignatureattachmentResponse)

	return err
}

type NullableEzsignsignatureattachmentResponse struct {
	value *EzsignsignatureattachmentResponse
	isSet bool
}

func (v NullableEzsignsignatureattachmentResponse) Get() *EzsignsignatureattachmentResponse {
	return v.value
}

func (v *NullableEzsignsignatureattachmentResponse) Set(val *EzsignsignatureattachmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureattachmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureattachmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureattachmentResponse(val *EzsignsignatureattachmentResponse) *NullableEzsignsignatureattachmentResponse {
	return &NullableEzsignsignatureattachmentResponse{value: val, isSet: true}
}

func (v NullableEzsignsignatureattachmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureattachmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


