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

// checks if the EzsignfolderSendV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderSendV2Request{}

// EzsignfolderSendV2Request Request for POST /2/object/ezsignfolder/{pkiEzsignfolderID}/send
type EzsignfolderSendV2Request struct {
	// A custom text message that will be added to the email sent.
	TEzsignfolderMessage string `json:"tEzsignfolderMessage"`
	AFkiEzsignfoldersignerassociationID []int32 `json:"a_fkiEzsignfoldersignerassociationID"`
	AObjEzsignfoldersignerassociationmessage []CustomEzsignfoldersignerassociationmessageRequest `json:"a_objEzsignfoldersignerassociationmessage"`
}

type _EzsignfolderSendV2Request EzsignfolderSendV2Request

// NewEzsignfolderSendV2Request instantiates a new EzsignfolderSendV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderSendV2Request(tEzsignfolderMessage string, aFkiEzsignfoldersignerassociationID []int32, aObjEzsignfoldersignerassociationmessage []CustomEzsignfoldersignerassociationmessageRequest) *EzsignfolderSendV2Request {
	this := EzsignfolderSendV2Request{}
	this.TEzsignfolderMessage = tEzsignfolderMessage
	this.AFkiEzsignfoldersignerassociationID = aFkiEzsignfoldersignerassociationID
	this.AObjEzsignfoldersignerassociationmessage = aObjEzsignfoldersignerassociationmessage
	return &this
}

// NewEzsignfolderSendV2RequestWithDefaults instantiates a new EzsignfolderSendV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderSendV2RequestWithDefaults() *EzsignfolderSendV2Request {
	this := EzsignfolderSendV2Request{}
	return &this
}

// GetTEzsignfolderMessage returns the TEzsignfolderMessage field value
func (o *EzsignfolderSendV2Request) GetTEzsignfolderMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TEzsignfolderMessage
}

// GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV2Request) GetTEzsignfolderMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TEzsignfolderMessage, true
}

// SetTEzsignfolderMessage sets field value
func (o *EzsignfolderSendV2Request) SetTEzsignfolderMessage(v string) {
	o.TEzsignfolderMessage = v
}

// GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field value
func (o *EzsignfolderSendV2Request) GetAFkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AFkiEzsignfoldersignerassociationID
}

// GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV2Request) GetAFkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AFkiEzsignfoldersignerassociationID, true
}

// SetAFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfolderSendV2Request) SetAFkiEzsignfoldersignerassociationID(v []int32) {
	o.AFkiEzsignfoldersignerassociationID = v
}

// GetAObjEzsignfoldersignerassociationmessage returns the AObjEzsignfoldersignerassociationmessage field value
func (o *EzsignfolderSendV2Request) GetAObjEzsignfoldersignerassociationmessage() []CustomEzsignfoldersignerassociationmessageRequest {
	if o == nil {
		var ret []CustomEzsignfoldersignerassociationmessageRequest
		return ret
	}

	return o.AObjEzsignfoldersignerassociationmessage
}

// GetAObjEzsignfoldersignerassociationmessageOk returns a tuple with the AObjEzsignfoldersignerassociationmessage field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV2Request) GetAObjEzsignfoldersignerassociationmessageOk() ([]CustomEzsignfoldersignerassociationmessageRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfoldersignerassociationmessage, true
}

// SetAObjEzsignfoldersignerassociationmessage sets field value
func (o *EzsignfolderSendV2Request) SetAObjEzsignfoldersignerassociationmessage(v []CustomEzsignfoldersignerassociationmessageRequest) {
	o.AObjEzsignfoldersignerassociationmessage = v
}

func (o EzsignfolderSendV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderSendV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tEzsignfolderMessage"] = o.TEzsignfolderMessage
	toSerialize["a_fkiEzsignfoldersignerassociationID"] = o.AFkiEzsignfoldersignerassociationID
	toSerialize["a_objEzsignfoldersignerassociationmessage"] = o.AObjEzsignfoldersignerassociationmessage
	return toSerialize, nil
}

func (o *EzsignfolderSendV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tEzsignfolderMessage",
		"a_fkiEzsignfoldersignerassociationID",
		"a_objEzsignfoldersignerassociationmessage",
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

	varEzsignfolderSendV2Request := _EzsignfolderSendV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderSendV2Request)

	if err != nil {
		return err
	}

	*o = EzsignfolderSendV2Request(varEzsignfolderSendV2Request)

	return err
}

type NullableEzsignfolderSendV2Request struct {
	value *EzsignfolderSendV2Request
	isSet bool
}

func (v NullableEzsignfolderSendV2Request) Get() *EzsignfolderSendV2Request {
	return v.value
}

func (v *NullableEzsignfolderSendV2Request) Set(val *EzsignfolderSendV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderSendV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderSendV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderSendV2Request(val *EzsignfolderSendV2Request) *NullableEzsignfolderSendV2Request {
	return &NullableEzsignfolderSendV2Request{value: val, isSet: true}
}

func (v NullableEzsignfolderSendV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderSendV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


