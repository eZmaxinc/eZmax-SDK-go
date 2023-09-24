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

// checks if the CommunicationexternalrecipientResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationexternalrecipientResponseCompound{}

// CommunicationexternalrecipientResponseCompound A Communicationexternalrecipient Object
type CommunicationexternalrecipientResponseCompound struct {
	// The unique ID of the Communicationexternalrecipient
	PkiCommunicationexternalrecipientID int32 `json:"pkiCommunicationexternalrecipientID"`
	ECommunicationexternalrecipientType FieldECommunicationexternalrecipientType `json:"eCommunicationexternalrecipientType"`
	ObjDescriptionstatic DescriptionstaticResponseCompound `json:"objDescriptionstatic"`
	ObjEmailstatic *EmailstaticResponseCompound `json:"objEmailstatic,omitempty"`
	ObjPhonestatic *PhonestaticResponseCompound `json:"objPhonestatic,omitempty"`
}

// NewCommunicationexternalrecipientResponseCompound instantiates a new CommunicationexternalrecipientResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationexternalrecipientResponseCompound(pkiCommunicationexternalrecipientID int32, eCommunicationexternalrecipientType FieldECommunicationexternalrecipientType, objDescriptionstatic DescriptionstaticResponseCompound) *CommunicationexternalrecipientResponseCompound {
	this := CommunicationexternalrecipientResponseCompound{}
	this.PkiCommunicationexternalrecipientID = pkiCommunicationexternalrecipientID
	this.ECommunicationexternalrecipientType = eCommunicationexternalrecipientType
	this.ObjDescriptionstatic = objDescriptionstatic
	return &this
}

// NewCommunicationexternalrecipientResponseCompoundWithDefaults instantiates a new CommunicationexternalrecipientResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationexternalrecipientResponseCompoundWithDefaults() *CommunicationexternalrecipientResponseCompound {
	this := CommunicationexternalrecipientResponseCompound{}
	return &this
}

// GetPkiCommunicationexternalrecipientID returns the PkiCommunicationexternalrecipientID field value
func (o *CommunicationexternalrecipientResponseCompound) GetPkiCommunicationexternalrecipientID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCommunicationexternalrecipientID
}

// GetPkiCommunicationexternalrecipientIDOk returns a tuple with the PkiCommunicationexternalrecipientID field value
// and a boolean to check if the value has been set.
func (o *CommunicationexternalrecipientResponseCompound) GetPkiCommunicationexternalrecipientIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCommunicationexternalrecipientID, true
}

// SetPkiCommunicationexternalrecipientID sets field value
func (o *CommunicationexternalrecipientResponseCompound) SetPkiCommunicationexternalrecipientID(v int32) {
	o.PkiCommunicationexternalrecipientID = v
}

// GetECommunicationexternalrecipientType returns the ECommunicationexternalrecipientType field value
func (o *CommunicationexternalrecipientResponseCompound) GetECommunicationexternalrecipientType() FieldECommunicationexternalrecipientType {
	if o == nil {
		var ret FieldECommunicationexternalrecipientType
		return ret
	}

	return o.ECommunicationexternalrecipientType
}

// GetECommunicationexternalrecipientTypeOk returns a tuple with the ECommunicationexternalrecipientType field value
// and a boolean to check if the value has been set.
func (o *CommunicationexternalrecipientResponseCompound) GetECommunicationexternalrecipientTypeOk() (*FieldECommunicationexternalrecipientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ECommunicationexternalrecipientType, true
}

// SetECommunicationexternalrecipientType sets field value
func (o *CommunicationexternalrecipientResponseCompound) SetECommunicationexternalrecipientType(v FieldECommunicationexternalrecipientType) {
	o.ECommunicationexternalrecipientType = v
}

// GetObjDescriptionstatic returns the ObjDescriptionstatic field value
func (o *CommunicationexternalrecipientResponseCompound) GetObjDescriptionstatic() DescriptionstaticResponseCompound {
	if o == nil {
		var ret DescriptionstaticResponseCompound
		return ret
	}

	return o.ObjDescriptionstatic
}

// GetObjDescriptionstaticOk returns a tuple with the ObjDescriptionstatic field value
// and a boolean to check if the value has been set.
func (o *CommunicationexternalrecipientResponseCompound) GetObjDescriptionstaticOk() (*DescriptionstaticResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDescriptionstatic, true
}

// SetObjDescriptionstatic sets field value
func (o *CommunicationexternalrecipientResponseCompound) SetObjDescriptionstatic(v DescriptionstaticResponseCompound) {
	o.ObjDescriptionstatic = v
}

// GetObjEmailstatic returns the ObjEmailstatic field value if set, zero value otherwise.
func (o *CommunicationexternalrecipientResponseCompound) GetObjEmailstatic() EmailstaticResponseCompound {
	if o == nil || IsNil(o.ObjEmailstatic) {
		var ret EmailstaticResponseCompound
		return ret
	}
	return *o.ObjEmailstatic
}

// GetObjEmailstaticOk returns a tuple with the ObjEmailstatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationexternalrecipientResponseCompound) GetObjEmailstaticOk() (*EmailstaticResponseCompound, bool) {
	if o == nil || IsNil(o.ObjEmailstatic) {
		return nil, false
	}
	return o.ObjEmailstatic, true
}

// HasObjEmailstatic returns a boolean if a field has been set.
func (o *CommunicationexternalrecipientResponseCompound) HasObjEmailstatic() bool {
	if o != nil && !IsNil(o.ObjEmailstatic) {
		return true
	}

	return false
}

// SetObjEmailstatic gets a reference to the given EmailstaticResponseCompound and assigns it to the ObjEmailstatic field.
func (o *CommunicationexternalrecipientResponseCompound) SetObjEmailstatic(v EmailstaticResponseCompound) {
	o.ObjEmailstatic = &v
}

// GetObjPhonestatic returns the ObjPhonestatic field value if set, zero value otherwise.
func (o *CommunicationexternalrecipientResponseCompound) GetObjPhonestatic() PhonestaticResponseCompound {
	if o == nil || IsNil(o.ObjPhonestatic) {
		var ret PhonestaticResponseCompound
		return ret
	}
	return *o.ObjPhonestatic
}

// GetObjPhonestaticOk returns a tuple with the ObjPhonestatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationexternalrecipientResponseCompound) GetObjPhonestaticOk() (*PhonestaticResponseCompound, bool) {
	if o == nil || IsNil(o.ObjPhonestatic) {
		return nil, false
	}
	return o.ObjPhonestatic, true
}

// HasObjPhonestatic returns a boolean if a field has been set.
func (o *CommunicationexternalrecipientResponseCompound) HasObjPhonestatic() bool {
	if o != nil && !IsNil(o.ObjPhonestatic) {
		return true
	}

	return false
}

// SetObjPhonestatic gets a reference to the given PhonestaticResponseCompound and assigns it to the ObjPhonestatic field.
func (o *CommunicationexternalrecipientResponseCompound) SetObjPhonestatic(v PhonestaticResponseCompound) {
	o.ObjPhonestatic = &v
}

func (o CommunicationexternalrecipientResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationexternalrecipientResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCommunicationexternalrecipientID"] = o.PkiCommunicationexternalrecipientID
	toSerialize["eCommunicationexternalrecipientType"] = o.ECommunicationexternalrecipientType
	toSerialize["objDescriptionstatic"] = o.ObjDescriptionstatic
	if !IsNil(o.ObjEmailstatic) {
		toSerialize["objEmailstatic"] = o.ObjEmailstatic
	}
	if !IsNil(o.ObjPhonestatic) {
		toSerialize["objPhonestatic"] = o.ObjPhonestatic
	}
	return toSerialize, nil
}

type NullableCommunicationexternalrecipientResponseCompound struct {
	value *CommunicationexternalrecipientResponseCompound
	isSet bool
}

func (v NullableCommunicationexternalrecipientResponseCompound) Get() *CommunicationexternalrecipientResponseCompound {
	return v.value
}

func (v *NullableCommunicationexternalrecipientResponseCompound) Set(val *CommunicationexternalrecipientResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationexternalrecipientResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationexternalrecipientResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationexternalrecipientResponseCompound(val *CommunicationexternalrecipientResponseCompound) *NullableCommunicationexternalrecipientResponseCompound {
	return &NullableCommunicationexternalrecipientResponseCompound{value: val, isSet: true}
}

func (v NullableCommunicationexternalrecipientResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationexternalrecipientResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


