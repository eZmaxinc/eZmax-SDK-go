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

// checks if the EzsignbulksendtransmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendtransmissionResponse{}

// EzsignbulksendtransmissionResponse An Ezsignbulksendtransmission Object
type EzsignbulksendtransmissionResponse struct {
	// The unique ID of the Ezsignbulksendtransmission
	PkiEzsignbulksendtransmissionID int32 `json:"pkiEzsignbulksendtransmissionID"`
	// The unique ID of the Ezsignbulksend
	FkiEzsignbulksendID int32 `json:"fkiEzsignbulksendID"`
	// The description of the Ezsignbulksendtransmission
	SEzsignbulksendtransmissionDescription string `json:"sEzsignbulksendtransmissionDescription"`
	// The number of errors during the Ezsignbulksendtransmission
	IEzsignbulksendtransmissionErrors int32 `json:"iEzsignbulksendtransmissionErrors"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _EzsignbulksendtransmissionResponse EzsignbulksendtransmissionResponse

// NewEzsignbulksendtransmissionResponse instantiates a new EzsignbulksendtransmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendtransmissionResponse(pkiEzsignbulksendtransmissionID int32, fkiEzsignbulksendID int32, sEzsignbulksendtransmissionDescription string, iEzsignbulksendtransmissionErrors int32, objAudit CommonAudit) *EzsignbulksendtransmissionResponse {
	this := EzsignbulksendtransmissionResponse{}
	this.PkiEzsignbulksendtransmissionID = pkiEzsignbulksendtransmissionID
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	this.SEzsignbulksendtransmissionDescription = sEzsignbulksendtransmissionDescription
	this.IEzsignbulksendtransmissionErrors = iEzsignbulksendtransmissionErrors
	this.ObjAudit = objAudit
	return &this
}

// NewEzsignbulksendtransmissionResponseWithDefaults instantiates a new EzsignbulksendtransmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendtransmissionResponseWithDefaults() *EzsignbulksendtransmissionResponse {
	this := EzsignbulksendtransmissionResponse{}
	return &this
}

// GetPkiEzsignbulksendtransmissionID returns the PkiEzsignbulksendtransmissionID field value
func (o *EzsignbulksendtransmissionResponse) GetPkiEzsignbulksendtransmissionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignbulksendtransmissionID
}

// GetPkiEzsignbulksendtransmissionIDOk returns a tuple with the PkiEzsignbulksendtransmissionID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionResponse) GetPkiEzsignbulksendtransmissionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignbulksendtransmissionID, true
}

// SetPkiEzsignbulksendtransmissionID sets field value
func (o *EzsignbulksendtransmissionResponse) SetPkiEzsignbulksendtransmissionID(v int32) {
	o.PkiEzsignbulksendtransmissionID = v
}

// GetFkiEzsignbulksendID returns the FkiEzsignbulksendID field value
func (o *EzsignbulksendtransmissionResponse) GetFkiEzsignbulksendID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignbulksendID
}

// GetFkiEzsignbulksendIDOk returns a tuple with the FkiEzsignbulksendID field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionResponse) GetFkiEzsignbulksendIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignbulksendID, true
}

// SetFkiEzsignbulksendID sets field value
func (o *EzsignbulksendtransmissionResponse) SetFkiEzsignbulksendID(v int32) {
	o.FkiEzsignbulksendID = v
}

// GetSEzsignbulksendtransmissionDescription returns the SEzsignbulksendtransmissionDescription field value
func (o *EzsignbulksendtransmissionResponse) GetSEzsignbulksendtransmissionDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignbulksendtransmissionDescription
}

// GetSEzsignbulksendtransmissionDescriptionOk returns a tuple with the SEzsignbulksendtransmissionDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionResponse) GetSEzsignbulksendtransmissionDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignbulksendtransmissionDescription, true
}

// SetSEzsignbulksendtransmissionDescription sets field value
func (o *EzsignbulksendtransmissionResponse) SetSEzsignbulksendtransmissionDescription(v string) {
	o.SEzsignbulksendtransmissionDescription = v
}

// GetIEzsignbulksendtransmissionErrors returns the IEzsignbulksendtransmissionErrors field value
func (o *EzsignbulksendtransmissionResponse) GetIEzsignbulksendtransmissionErrors() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignbulksendtransmissionErrors
}

// GetIEzsignbulksendtransmissionErrorsOk returns a tuple with the IEzsignbulksendtransmissionErrors field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionResponse) GetIEzsignbulksendtransmissionErrorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignbulksendtransmissionErrors, true
}

// SetIEzsignbulksendtransmissionErrors sets field value
func (o *EzsignbulksendtransmissionResponse) SetIEzsignbulksendtransmissionErrors(v int32) {
	o.IEzsignbulksendtransmissionErrors = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsignbulksendtransmissionResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsignbulksendtransmissionResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzsignbulksendtransmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendtransmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignbulksendtransmissionID"] = o.PkiEzsignbulksendtransmissionID
	toSerialize["fkiEzsignbulksendID"] = o.FkiEzsignbulksendID
	toSerialize["sEzsignbulksendtransmissionDescription"] = o.SEzsignbulksendtransmissionDescription
	toSerialize["iEzsignbulksendtransmissionErrors"] = o.IEzsignbulksendtransmissionErrors
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *EzsignbulksendtransmissionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignbulksendtransmissionID",
		"fkiEzsignbulksendID",
		"sEzsignbulksendtransmissionDescription",
		"iEzsignbulksendtransmissionErrors",
		"objAudit",
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

	varEzsignbulksendtransmissionResponse := _EzsignbulksendtransmissionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendtransmissionResponse)

	if err != nil {
		return err
	}

	*o = EzsignbulksendtransmissionResponse(varEzsignbulksendtransmissionResponse)

	return err
}

type NullableEzsignbulksendtransmissionResponse struct {
	value *EzsignbulksendtransmissionResponse
	isSet bool
}

func (v NullableEzsignbulksendtransmissionResponse) Get() *EzsignbulksendtransmissionResponse {
	return v.value
}

func (v *NullableEzsignbulksendtransmissionResponse) Set(val *EzsignbulksendtransmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendtransmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendtransmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendtransmissionResponse(val *EzsignbulksendtransmissionResponse) *NullableEzsignbulksendtransmissionResponse {
	return &NullableEzsignbulksendtransmissionResponse{value: val, isSet: true}
}

func (v NullableEzsignbulksendtransmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendtransmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


