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

// checks if the EzsignfolderBatchDownloadV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderBatchDownloadV1Request{}

// EzsignfolderBatchDownloadV1Request Request for POST /1/object/ezsignfolder/{pkiEzsignfolderID}/batchDownload
type EzsignfolderBatchDownloadV1Request struct {
	APkiEzsigndocumentID []int32 `json:"a_pkiEzsigndocumentID"`
	// The type of document to retrieve.  1. **Signed** Is the final document once all signatures were applied. 2. **Proofdocument** Is the evidence report. 3. **Proof** Is the complete evidence archive including all of the above and more.
	AEDocumentType []string `json:"a_eDocumentType"`
}

type _EzsignfolderBatchDownloadV1Request EzsignfolderBatchDownloadV1Request

// NewEzsignfolderBatchDownloadV1Request instantiates a new EzsignfolderBatchDownloadV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderBatchDownloadV1Request(aPkiEzsigndocumentID []int32, aEDocumentType []string) *EzsignfolderBatchDownloadV1Request {
	this := EzsignfolderBatchDownloadV1Request{}
	this.APkiEzsigndocumentID = aPkiEzsigndocumentID
	this.AEDocumentType = aEDocumentType
	return &this
}

// NewEzsignfolderBatchDownloadV1RequestWithDefaults instantiates a new EzsignfolderBatchDownloadV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderBatchDownloadV1RequestWithDefaults() *EzsignfolderBatchDownloadV1Request {
	this := EzsignfolderBatchDownloadV1Request{}
	return &this
}

// GetAPkiEzsigndocumentID returns the APkiEzsigndocumentID field value
func (o *EzsignfolderBatchDownloadV1Request) GetAPkiEzsigndocumentID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigndocumentID
}

// GetAPkiEzsigndocumentIDOk returns a tuple with the APkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderBatchDownloadV1Request) GetAPkiEzsigndocumentIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigndocumentID, true
}

// SetAPkiEzsigndocumentID sets field value
func (o *EzsignfolderBatchDownloadV1Request) SetAPkiEzsigndocumentID(v []int32) {
	o.APkiEzsigndocumentID = v
}

// GetAEDocumentType returns the AEDocumentType field value
func (o *EzsignfolderBatchDownloadV1Request) GetAEDocumentType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AEDocumentType
}

// GetAEDocumentTypeOk returns a tuple with the AEDocumentType field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderBatchDownloadV1Request) GetAEDocumentTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AEDocumentType, true
}

// SetAEDocumentType sets field value
func (o *EzsignfolderBatchDownloadV1Request) SetAEDocumentType(v []string) {
	o.AEDocumentType = v
}

func (o EzsignfolderBatchDownloadV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderBatchDownloadV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigndocumentID"] = o.APkiEzsigndocumentID
	toSerialize["a_eDocumentType"] = o.AEDocumentType
	return toSerialize, nil
}

func (o *EzsignfolderBatchDownloadV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsigndocumentID",
		"a_eDocumentType",
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

	varEzsignfolderBatchDownloadV1Request := _EzsignfolderBatchDownloadV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderBatchDownloadV1Request)

	if err != nil {
		return err
	}

	*o = EzsignfolderBatchDownloadV1Request(varEzsignfolderBatchDownloadV1Request)

	return err
}

type NullableEzsignfolderBatchDownloadV1Request struct {
	value *EzsignfolderBatchDownloadV1Request
	isSet bool
}

func (v NullableEzsignfolderBatchDownloadV1Request) Get() *EzsignfolderBatchDownloadV1Request {
	return v.value
}

func (v *NullableEzsignfolderBatchDownloadV1Request) Set(val *EzsignfolderBatchDownloadV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderBatchDownloadV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderBatchDownloadV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderBatchDownloadV1Request(val *EzsignfolderBatchDownloadV1Request) *NullableEzsignfolderBatchDownloadV1Request {
	return &NullableEzsignfolderBatchDownloadV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderBatchDownloadV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderBatchDownloadV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


