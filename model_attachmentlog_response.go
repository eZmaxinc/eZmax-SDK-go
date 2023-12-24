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

// checks if the AttachmentlogResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachmentlogResponse{}

// AttachmentlogResponse An Attachmentlog Object
type AttachmentlogResponse struct {
	// The unique ID of the Attachment.
	FkiAttachmentID int32 `json:"fkiAttachmentID"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The created date
	DtAttachmentlogDatetime string `json:"dtAttachmentlogDatetime"`
	EAttachmentlogType FieldEAttachmentlogType `json:"eAttachmentlogType"`
	// The additionnal detail
	SAttachmentlogDetail *string `json:"sAttachmentlogDetail,omitempty"`
}

type _AttachmentlogResponse AttachmentlogResponse

// NewAttachmentlogResponse instantiates a new AttachmentlogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentlogResponse(fkiAttachmentID int32, fkiUserID int32, dtAttachmentlogDatetime string, eAttachmentlogType FieldEAttachmentlogType) *AttachmentlogResponse {
	this := AttachmentlogResponse{}
	this.FkiAttachmentID = fkiAttachmentID
	this.FkiUserID = fkiUserID
	this.DtAttachmentlogDatetime = dtAttachmentlogDatetime
	this.EAttachmentlogType = eAttachmentlogType
	return &this
}

// NewAttachmentlogResponseWithDefaults instantiates a new AttachmentlogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentlogResponseWithDefaults() *AttachmentlogResponse {
	this := AttachmentlogResponse{}
	return &this
}

// GetFkiAttachmentID returns the FkiAttachmentID field value
func (o *AttachmentlogResponse) GetFkiAttachmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiAttachmentID
}

// GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field value
// and a boolean to check if the value has been set.
func (o *AttachmentlogResponse) GetFkiAttachmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiAttachmentID, true
}

// SetFkiAttachmentID sets field value
func (o *AttachmentlogResponse) SetFkiAttachmentID(v int32) {
	o.FkiAttachmentID = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *AttachmentlogResponse) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *AttachmentlogResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *AttachmentlogResponse) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetDtAttachmentlogDatetime returns the DtAttachmentlogDatetime field value
func (o *AttachmentlogResponse) GetDtAttachmentlogDatetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtAttachmentlogDatetime
}

// GetDtAttachmentlogDatetimeOk returns a tuple with the DtAttachmentlogDatetime field value
// and a boolean to check if the value has been set.
func (o *AttachmentlogResponse) GetDtAttachmentlogDatetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtAttachmentlogDatetime, true
}

// SetDtAttachmentlogDatetime sets field value
func (o *AttachmentlogResponse) SetDtAttachmentlogDatetime(v string) {
	o.DtAttachmentlogDatetime = v
}

// GetEAttachmentlogType returns the EAttachmentlogType field value
func (o *AttachmentlogResponse) GetEAttachmentlogType() FieldEAttachmentlogType {
	if o == nil {
		var ret FieldEAttachmentlogType
		return ret
	}

	return o.EAttachmentlogType
}

// GetEAttachmentlogTypeOk returns a tuple with the EAttachmentlogType field value
// and a boolean to check if the value has been set.
func (o *AttachmentlogResponse) GetEAttachmentlogTypeOk() (*FieldEAttachmentlogType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EAttachmentlogType, true
}

// SetEAttachmentlogType sets field value
func (o *AttachmentlogResponse) SetEAttachmentlogType(v FieldEAttachmentlogType) {
	o.EAttachmentlogType = v
}

// GetSAttachmentlogDetail returns the SAttachmentlogDetail field value if set, zero value otherwise.
func (o *AttachmentlogResponse) GetSAttachmentlogDetail() string {
	if o == nil || IsNil(o.SAttachmentlogDetail) {
		var ret string
		return ret
	}
	return *o.SAttachmentlogDetail
}

// GetSAttachmentlogDetailOk returns a tuple with the SAttachmentlogDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentlogResponse) GetSAttachmentlogDetailOk() (*string, bool) {
	if o == nil || IsNil(o.SAttachmentlogDetail) {
		return nil, false
	}
	return o.SAttachmentlogDetail, true
}

// HasSAttachmentlogDetail returns a boolean if a field has been set.
func (o *AttachmentlogResponse) HasSAttachmentlogDetail() bool {
	if o != nil && !IsNil(o.SAttachmentlogDetail) {
		return true
	}

	return false
}

// SetSAttachmentlogDetail gets a reference to the given string and assigns it to the SAttachmentlogDetail field.
func (o *AttachmentlogResponse) SetSAttachmentlogDetail(v string) {
	o.SAttachmentlogDetail = &v
}

func (o AttachmentlogResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentlogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiAttachmentID"] = o.FkiAttachmentID
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["dtAttachmentlogDatetime"] = o.DtAttachmentlogDatetime
	toSerialize["eAttachmentlogType"] = o.EAttachmentlogType
	if !IsNil(o.SAttachmentlogDetail) {
		toSerialize["sAttachmentlogDetail"] = o.SAttachmentlogDetail
	}
	return toSerialize, nil
}

func (o *AttachmentlogResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiAttachmentID",
		"fkiUserID",
		"dtAttachmentlogDatetime",
		"eAttachmentlogType",
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

	varAttachmentlogResponse := _AttachmentlogResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachmentlogResponse)

	if err != nil {
		return err
	}

	*o = AttachmentlogResponse(varAttachmentlogResponse)

	return err
}

type NullableAttachmentlogResponse struct {
	value *AttachmentlogResponse
	isSet bool
}

func (v NullableAttachmentlogResponse) Get() *AttachmentlogResponse {
	return v.value
}

func (v *NullableAttachmentlogResponse) Set(val *AttachmentlogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentlogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentlogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentlogResponse(val *AttachmentlogResponse) *NullableAttachmentlogResponse {
	return &NullableAttachmentlogResponse{value: val, isSet: true}
}

func (v NullableAttachmentlogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentlogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


