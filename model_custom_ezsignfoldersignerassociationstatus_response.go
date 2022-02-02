/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CustomEzsignfoldersignerassociationstatusResponse A Ezsignfoldersignerassociationstatus Object and children to create a complete structure
type CustomEzsignfoldersignerassociationstatusResponse struct {
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
	// The last name of the Ezsignsigner
	SEzsignfoldersignerassociationstatusLastname string `json:"sEzsignfoldersignerassociationstatusLastname"`
	// The first name of the Ezsignsigner
	SEzsignfoldersignerassociationstatusFirstname string `json:"sEzsignfoldersignerassociationstatusFirstname"`
	AObjEzsignsignaturestatus []CustomEzsignsignaturestatusResponse `json:"a_objEzsignsignaturestatus"`
}

// NewCustomEzsignfoldersignerassociationstatusResponse instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignfoldersignerassociationstatusResponse(fkiEzsignfoldersignerassociationID int32, sEzsignfoldersignerassociationstatusLastname string, sEzsignfoldersignerassociationstatusFirstname string, aObjEzsignsignaturestatus []CustomEzsignsignaturestatusResponse) *CustomEzsignfoldersignerassociationstatusResponse {
	this := CustomEzsignfoldersignerassociationstatusResponse{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	this.SEzsignfoldersignerassociationstatusLastname = sEzsignfoldersignerassociationstatusLastname
	this.SEzsignfoldersignerassociationstatusFirstname = sEzsignfoldersignerassociationstatusFirstname
	this.AObjEzsignsignaturestatus = aObjEzsignsignaturestatus
	return &this
}

// NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults() *CustomEzsignfoldersignerassociationstatusResponse {
	this := CustomEzsignfoldersignerassociationstatusResponse{}
	return &this
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

// GetSEzsignfoldersignerassociationstatusLastname returns the SEzsignfoldersignerassociationstatusLastname field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldersignerassociationstatusLastname
}

// GetSEzsignfoldersignerassociationstatusLastnameOk returns a tuple with the SEzsignfoldersignerassociationstatusLastname field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfoldersignerassociationstatusLastname, true
}

// SetSEzsignfoldersignerassociationstatusLastname sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusLastname(v string) {
	o.SEzsignfoldersignerassociationstatusLastname = v
}

// GetSEzsignfoldersignerassociationstatusFirstname returns the SEzsignfoldersignerassociationstatusFirstname field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldersignerassociationstatusFirstname
}

// GetSEzsignfoldersignerassociationstatusFirstnameOk returns a tuple with the SEzsignfoldersignerassociationstatusFirstname field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfoldersignerassociationstatusFirstname, true
}

// SetSEzsignfoldersignerassociationstatusFirstname sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusFirstname(v string) {
	o.SEzsignfoldersignerassociationstatusFirstname = v
}

// GetAObjEzsignsignaturestatus returns the AObjEzsignsignaturestatus field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatus() []CustomEzsignsignaturestatusResponse {
	if o == nil {
		var ret []CustomEzsignsignaturestatusResponse
		return ret
	}

	return o.AObjEzsignsignaturestatus
}

// GetAObjEzsignsignaturestatusOk returns a tuple with the AObjEzsignsignaturestatus field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatusOk() ([]CustomEzsignsignaturestatusResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AObjEzsignsignaturestatus, true
}

// SetAObjEzsignsignaturestatus sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetAObjEzsignsignaturestatus(v []CustomEzsignsignaturestatusResponse) {
	o.AObjEzsignsignaturestatus = v
}

func (o CustomEzsignfoldersignerassociationstatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	}
	if true {
		toSerialize["sEzsignfoldersignerassociationstatusLastname"] = o.SEzsignfoldersignerassociationstatusLastname
	}
	if true {
		toSerialize["sEzsignfoldersignerassociationstatusFirstname"] = o.SEzsignfoldersignerassociationstatusFirstname
	}
	if true {
		toSerialize["a_objEzsignsignaturestatus"] = o.AObjEzsignsignaturestatus
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEzsignfoldersignerassociationstatusResponse struct {
	value *CustomEzsignfoldersignerassociationstatusResponse
	isSet bool
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) Get() *CustomEzsignfoldersignerassociationstatusResponse {
	return v.value
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) Set(val *CustomEzsignfoldersignerassociationstatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignfoldersignerassociationstatusResponse(val *CustomEzsignfoldersignerassociationstatusResponse) *NullableCustomEzsignfoldersignerassociationstatusResponse {
	return &NullableCustomEzsignfoldersignerassociationstatusResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

