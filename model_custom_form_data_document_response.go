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

// checks if the CustomFormDataDocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFormDataDocumentResponse{}

// CustomFormDataDocumentResponse A form Data Document Object 
type CustomFormDataDocumentResponse struct {
	// The unique ID of the Ezsigndocument
	PkiEzsigndocumentID int32 `json:"pkiEzsigndocumentID"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
	// The date and time at which the object was last modified
	DtModifiedDate string `json:"dtModifiedDate"`
	AObjFormDataSigner []CustomFormDataSignerResponse `json:"a_objFormDataSigner"`
}

type _CustomFormDataDocumentResponse CustomFormDataDocumentResponse

// NewCustomFormDataDocumentResponse instantiates a new CustomFormDataDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormDataDocumentResponse(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, sEzsigndocumentName string, dtModifiedDate string, aObjFormDataSigner []CustomFormDataSignerResponse) *CustomFormDataDocumentResponse {
	this := CustomFormDataDocumentResponse{}
	this.PkiEzsigndocumentID = pkiEzsigndocumentID
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.SEzsigndocumentName = sEzsigndocumentName
	this.DtModifiedDate = dtModifiedDate
	this.AObjFormDataSigner = aObjFormDataSigner
	return &this
}

// NewCustomFormDataDocumentResponseWithDefaults instantiates a new CustomFormDataDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormDataDocumentResponseWithDefaults() *CustomFormDataDocumentResponse {
	this := CustomFormDataDocumentResponse{}
	return &this
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value
func (o *CustomFormDataDocumentResponse) GetPkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataDocumentResponse) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigndocumentID, true
}

// SetPkiEzsigndocumentID sets field value
func (o *CustomFormDataDocumentResponse) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *CustomFormDataDocumentResponse) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataDocumentResponse) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *CustomFormDataDocumentResponse) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *CustomFormDataDocumentResponse) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataDocumentResponse) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *CustomFormDataDocumentResponse) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

// GetDtModifiedDate returns the DtModifiedDate field value
func (o *CustomFormDataDocumentResponse) GetDtModifiedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtModifiedDate
}

// GetDtModifiedDateOk returns a tuple with the DtModifiedDate field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataDocumentResponse) GetDtModifiedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtModifiedDate, true
}

// SetDtModifiedDate sets field value
func (o *CustomFormDataDocumentResponse) SetDtModifiedDate(v string) {
	o.DtModifiedDate = v
}

// GetAObjFormDataSigner returns the AObjFormDataSigner field value
func (o *CustomFormDataDocumentResponse) GetAObjFormDataSigner() []CustomFormDataSignerResponse {
	if o == nil {
		var ret []CustomFormDataSignerResponse
		return ret
	}

	return o.AObjFormDataSigner
}

// GetAObjFormDataSignerOk returns a tuple with the AObjFormDataSigner field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataDocumentResponse) GetAObjFormDataSignerOk() ([]CustomFormDataSignerResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFormDataSigner, true
}

// SetAObjFormDataSigner sets field value
func (o *CustomFormDataDocumentResponse) SetAObjFormDataSigner(v []CustomFormDataSignerResponse) {
	o.AObjFormDataSigner = v
}

func (o CustomFormDataDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFormDataDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigndocumentID"] = o.PkiEzsigndocumentID
	toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	toSerialize["dtModifiedDate"] = o.DtModifiedDate
	toSerialize["a_objFormDataSigner"] = o.AObjFormDataSigner
	return toSerialize, nil
}

func (o *CustomFormDataDocumentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigndocumentID",
		"fkiEzsignfolderID",
		"sEzsigndocumentName",
		"dtModifiedDate",
		"a_objFormDataSigner",
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

	varCustomFormDataDocumentResponse := _CustomFormDataDocumentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomFormDataDocumentResponse)

	if err != nil {
		return err
	}

	*o = CustomFormDataDocumentResponse(varCustomFormDataDocumentResponse)

	return err
}

type NullableCustomFormDataDocumentResponse struct {
	value *CustomFormDataDocumentResponse
	isSet bool
}

func (v NullableCustomFormDataDocumentResponse) Get() *CustomFormDataDocumentResponse {
	return v.value
}

func (v *NullableCustomFormDataDocumentResponse) Set(val *CustomFormDataDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormDataDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormDataDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormDataDocumentResponse(val *CustomFormDataDocumentResponse) *NullableCustomFormDataDocumentResponse {
	return &NullableCustomFormDataDocumentResponse{value: val, isSet: true}
}

func (v NullableCustomFormDataDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormDataDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


