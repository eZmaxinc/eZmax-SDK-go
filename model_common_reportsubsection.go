/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CommonReportsubsection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonReportsubsection{}

// CommonReportsubsection A Subsection in a Reportsection. It contains 3 Reportsubsectionparts (Header, Body and Footer) 
type CommonReportsubsection struct {
	ObjReportsubsectionpartHeader CommonReportsubsectionpart `json:"objReportsubsectionpartHeader"`
	ObjReportsubsectionpartBody CommonReportsubsectionpart `json:"objReportsubsectionpartBody"`
	ObjReportsubsectionpartFooter CommonReportsubsectionpart `json:"objReportsubsectionpartFooter"`
}

type _CommonReportsubsection CommonReportsubsection

// NewCommonReportsubsection instantiates a new CommonReportsubsection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonReportsubsection(objReportsubsectionpartHeader CommonReportsubsectionpart, objReportsubsectionpartBody CommonReportsubsectionpart, objReportsubsectionpartFooter CommonReportsubsectionpart) *CommonReportsubsection {
	this := CommonReportsubsection{}
	this.ObjReportsubsectionpartHeader = objReportsubsectionpartHeader
	this.ObjReportsubsectionpartBody = objReportsubsectionpartBody
	this.ObjReportsubsectionpartFooter = objReportsubsectionpartFooter
	return &this
}

// NewCommonReportsubsectionWithDefaults instantiates a new CommonReportsubsection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonReportsubsectionWithDefaults() *CommonReportsubsection {
	this := CommonReportsubsection{}
	return &this
}

// GetObjReportsubsectionpartHeader returns the ObjReportsubsectionpartHeader field value
func (o *CommonReportsubsection) GetObjReportsubsectionpartHeader() CommonReportsubsectionpart {
	if o == nil {
		var ret CommonReportsubsectionpart
		return ret
	}

	return o.ObjReportsubsectionpartHeader
}

// GetObjReportsubsectionpartHeaderOk returns a tuple with the ObjReportsubsectionpartHeader field value
// and a boolean to check if the value has been set.
func (o *CommonReportsubsection) GetObjReportsubsectionpartHeaderOk() (*CommonReportsubsectionpart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjReportsubsectionpartHeader, true
}

// SetObjReportsubsectionpartHeader sets field value
func (o *CommonReportsubsection) SetObjReportsubsectionpartHeader(v CommonReportsubsectionpart) {
	o.ObjReportsubsectionpartHeader = v
}

// GetObjReportsubsectionpartBody returns the ObjReportsubsectionpartBody field value
func (o *CommonReportsubsection) GetObjReportsubsectionpartBody() CommonReportsubsectionpart {
	if o == nil {
		var ret CommonReportsubsectionpart
		return ret
	}

	return o.ObjReportsubsectionpartBody
}

// GetObjReportsubsectionpartBodyOk returns a tuple with the ObjReportsubsectionpartBody field value
// and a boolean to check if the value has been set.
func (o *CommonReportsubsection) GetObjReportsubsectionpartBodyOk() (*CommonReportsubsectionpart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjReportsubsectionpartBody, true
}

// SetObjReportsubsectionpartBody sets field value
func (o *CommonReportsubsection) SetObjReportsubsectionpartBody(v CommonReportsubsectionpart) {
	o.ObjReportsubsectionpartBody = v
}

// GetObjReportsubsectionpartFooter returns the ObjReportsubsectionpartFooter field value
func (o *CommonReportsubsection) GetObjReportsubsectionpartFooter() CommonReportsubsectionpart {
	if o == nil {
		var ret CommonReportsubsectionpart
		return ret
	}

	return o.ObjReportsubsectionpartFooter
}

// GetObjReportsubsectionpartFooterOk returns a tuple with the ObjReportsubsectionpartFooter field value
// and a boolean to check if the value has been set.
func (o *CommonReportsubsection) GetObjReportsubsectionpartFooterOk() (*CommonReportsubsectionpart, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjReportsubsectionpartFooter, true
}

// SetObjReportsubsectionpartFooter sets field value
func (o *CommonReportsubsection) SetObjReportsubsectionpartFooter(v CommonReportsubsectionpart) {
	o.ObjReportsubsectionpartFooter = v
}

func (o CommonReportsubsection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonReportsubsection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objReportsubsectionpartHeader"] = o.ObjReportsubsectionpartHeader
	toSerialize["objReportsubsectionpartBody"] = o.ObjReportsubsectionpartBody
	toSerialize["objReportsubsectionpartFooter"] = o.ObjReportsubsectionpartFooter
	return toSerialize, nil
}

func (o *CommonReportsubsection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objReportsubsectionpartHeader",
		"objReportsubsectionpartBody",
		"objReportsubsectionpartFooter",
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

	varCommonReportsubsection := _CommonReportsubsection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonReportsubsection)

	if err != nil {
		return err
	}

	*o = CommonReportsubsection(varCommonReportsubsection)

	return err
}

type NullableCommonReportsubsection struct {
	value *CommonReportsubsection
	isSet bool
}

func (v NullableCommonReportsubsection) Get() *CommonReportsubsection {
	return v.value
}

func (v *NullableCommonReportsubsection) Set(val *CommonReportsubsection) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonReportsubsection) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonReportsubsection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonReportsubsection(val *CommonReportsubsection) *NullableCommonReportsubsection {
	return &NullableCommonReportsubsection{value: val, isSet: true}
}

func (v NullableCommonReportsubsection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonReportsubsection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


