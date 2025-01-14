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

// checks if the CustomEzsignfolderezsigntemplatepublicResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignfolderezsigntemplatepublicResponse{}

// CustomEzsignfolderezsigntemplatepublicResponse An Ezsignfolder Object in the context of an Ezsigntemplatepublic
type CustomEzsignfolderezsigntemplatepublicResponse struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription" validate:"regexp=^.{0,75}$"`
	EEzsignfolderStep FieldEEzsignfolderStep `json:"eEzsignfolderStep"`
	// The number of total signatures that were requested in the Ezsignfolder
	IEzsignfolderSignaturetotal int32 `json:"iEzsignfolderSignaturetotal"`
	// The number of total form fields that were requested in the Ezsignfolder
	IEzsignfolderFormfieldtotal int32 `json:"iEzsignfolderFormfieldtotal"`
	// The number of signatures that were signed in the Ezsignfolder.
	IEzsignfolderSignaturesigned int32 `json:"iEzsignfolderSignaturesigned"`
	AObjEzsignfolderezsigntemplatepublicSigner []CustomEzsignfolderezsigntemplatepublicSignerResponse `json:"a_objEzsignfolderezsigntemplatepublicSigner,omitempty"`
}

type _CustomEzsignfolderezsigntemplatepublicResponse CustomEzsignfolderezsigntemplatepublicResponse

// NewCustomEzsignfolderezsigntemplatepublicResponse instantiates a new CustomEzsignfolderezsigntemplatepublicResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignfolderezsigntemplatepublicResponse(pkiEzsignfolderID int32, sEzsignfolderDescription string, eEzsignfolderStep FieldEEzsignfolderStep, iEzsignfolderSignaturetotal int32, iEzsignfolderFormfieldtotal int32, iEzsignfolderSignaturesigned int32) *CustomEzsignfolderezsigntemplatepublicResponse {
	this := CustomEzsignfolderezsigntemplatepublicResponse{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.EEzsignfolderStep = eEzsignfolderStep
	this.IEzsignfolderSignaturetotal = iEzsignfolderSignaturetotal
	this.IEzsignfolderFormfieldtotal = iEzsignfolderFormfieldtotal
	this.IEzsignfolderSignaturesigned = iEzsignfolderSignaturesigned
	return &this
}

// NewCustomEzsignfolderezsigntemplatepublicResponseWithDefaults instantiates a new CustomEzsignfolderezsigntemplatepublicResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignfolderezsigntemplatepublicResponseWithDefaults() *CustomEzsignfolderezsigntemplatepublicResponse {
	this := CustomEzsignfolderezsigntemplatepublicResponse{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetEEzsignfolderStep returns the EEzsignfolderStep field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetEEzsignfolderStep() FieldEEzsignfolderStep {
	if o == nil {
		var ret FieldEEzsignfolderStep
		return ret
	}

	return o.EEzsignfolderStep
}

// GetEEzsignfolderStepOk returns a tuple with the EEzsignfolderStep field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetEEzsignfolderStepOk() (*FieldEEzsignfolderStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzsignfolderStep, true
}

// SetEEzsignfolderStep sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetEEzsignfolderStep(v FieldEEzsignfolderStep) {
	o.EEzsignfolderStep = v
}

// GetIEzsignfolderSignaturetotal returns the IEzsignfolderSignaturetotal field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturetotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignfolderSignaturetotal
}

// GetIEzsignfolderSignaturetotalOk returns a tuple with the IEzsignfolderSignaturetotal field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturetotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignfolderSignaturetotal, true
}

// SetIEzsignfolderSignaturetotal sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderSignaturetotal(v int32) {
	o.IEzsignfolderSignaturetotal = v
}

// GetIEzsignfolderFormfieldtotal returns the IEzsignfolderFormfieldtotal field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderFormfieldtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignfolderFormfieldtotal
}

// GetIEzsignfolderFormfieldtotalOk returns a tuple with the IEzsignfolderFormfieldtotal field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderFormfieldtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignfolderFormfieldtotal, true
}

// SetIEzsignfolderFormfieldtotal sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderFormfieldtotal(v int32) {
	o.IEzsignfolderFormfieldtotal = v
}

// GetIEzsignfolderSignaturesigned returns the IEzsignfolderSignaturesigned field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturesigned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignfolderSignaturesigned
}

// GetIEzsignfolderSignaturesignedOk returns a tuple with the IEzsignfolderSignaturesigned field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetIEzsignfolderSignaturesignedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsignfolderSignaturesigned, true
}

// SetIEzsignfolderSignaturesigned sets field value
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetIEzsignfolderSignaturesigned(v int32) {
	o.IEzsignfolderSignaturesigned = v
}

// GetAObjEzsignfolderezsigntemplatepublicSigner returns the AObjEzsignfolderezsigntemplatepublicSigner field value if set, zero value otherwise.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetAObjEzsignfolderezsigntemplatepublicSigner() []CustomEzsignfolderezsigntemplatepublicSignerResponse {
	if o == nil || IsNil(o.AObjEzsignfolderezsigntemplatepublicSigner) {
		var ret []CustomEzsignfolderezsigntemplatepublicSignerResponse
		return ret
	}
	return o.AObjEzsignfolderezsigntemplatepublicSigner
}

// GetAObjEzsignfolderezsigntemplatepublicSignerOk returns a tuple with the AObjEzsignfolderezsigntemplatepublicSigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) GetAObjEzsignfolderezsigntemplatepublicSignerOk() ([]CustomEzsignfolderezsigntemplatepublicSignerResponse, bool) {
	if o == nil || IsNil(o.AObjEzsignfolderezsigntemplatepublicSigner) {
		return nil, false
	}
	return o.AObjEzsignfolderezsigntemplatepublicSigner, true
}

// HasAObjEzsignfolderezsigntemplatepublicSigner returns a boolean if a field has been set.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) HasAObjEzsignfolderezsigntemplatepublicSigner() bool {
	if o != nil && !IsNil(o.AObjEzsignfolderezsigntemplatepublicSigner) {
		return true
	}

	return false
}

// SetAObjEzsignfolderezsigntemplatepublicSigner gets a reference to the given []CustomEzsignfolderezsigntemplatepublicSignerResponse and assigns it to the AObjEzsignfolderezsigntemplatepublicSigner field.
func (o *CustomEzsignfolderezsigntemplatepublicResponse) SetAObjEzsignfolderezsigntemplatepublicSigner(v []CustomEzsignfolderezsigntemplatepublicSignerResponse) {
	o.AObjEzsignfolderezsigntemplatepublicSigner = v
}

func (o CustomEzsignfolderezsigntemplatepublicResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignfolderezsigntemplatepublicResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	toSerialize["eEzsignfolderStep"] = o.EEzsignfolderStep
	toSerialize["iEzsignfolderSignaturetotal"] = o.IEzsignfolderSignaturetotal
	toSerialize["iEzsignfolderFormfieldtotal"] = o.IEzsignfolderFormfieldtotal
	toSerialize["iEzsignfolderSignaturesigned"] = o.IEzsignfolderSignaturesigned
	if !IsNil(o.AObjEzsignfolderezsigntemplatepublicSigner) {
		toSerialize["a_objEzsignfolderezsigntemplatepublicSigner"] = o.AObjEzsignfolderezsigntemplatepublicSigner
	}
	return toSerialize, nil
}

func (o *CustomEzsignfolderezsigntemplatepublicResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignfolderID",
		"sEzsignfolderDescription",
		"eEzsignfolderStep",
		"iEzsignfolderSignaturetotal",
		"iEzsignfolderFormfieldtotal",
		"iEzsignfolderSignaturesigned",
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

	varCustomEzsignfolderezsigntemplatepublicResponse := _CustomEzsignfolderezsigntemplatepublicResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEzsignfolderezsigntemplatepublicResponse)

	if err != nil {
		return err
	}

	*o = CustomEzsignfolderezsigntemplatepublicResponse(varCustomEzsignfolderezsigntemplatepublicResponse)

	return err
}

type NullableCustomEzsignfolderezsigntemplatepublicResponse struct {
	value *CustomEzsignfolderezsigntemplatepublicResponse
	isSet bool
}

func (v NullableCustomEzsignfolderezsigntemplatepublicResponse) Get() *CustomEzsignfolderezsigntemplatepublicResponse {
	return v.value
}

func (v *NullableCustomEzsignfolderezsigntemplatepublicResponse) Set(val *CustomEzsignfolderezsigntemplatepublicResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignfolderezsigntemplatepublicResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignfolderezsigntemplatepublicResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignfolderezsigntemplatepublicResponse(val *CustomEzsignfolderezsigntemplatepublicResponse) *NullableCustomEzsignfolderezsigntemplatepublicResponse {
	return &NullableCustomEzsignfolderezsigntemplatepublicResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignfolderezsigntemplatepublicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignfolderezsigntemplatepublicResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


