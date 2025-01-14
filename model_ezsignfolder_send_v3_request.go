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

// checks if the EzsignfolderSendV3Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderSendV3Request{}

// EzsignfolderSendV3Request Request for POST /3/object/ezsignfolder/{pkiEzsignfolderID}/send
type EzsignfolderSendV3Request struct {
	// A custom text message that will be added to the email sent.
	TEzsignfolderMessage *string `json:"tEzsignfolderMessage,omitempty"`
	EEzsignfolderMessageorder *FieldEEzsignfolderMessageorder `json:"eEzsignfolderMessageorder,omitempty"`
	// The date and time at which the Ezsignfolder will be sent in the future.
	DtEzsignfolderDelayedsenddate *string `json:"dtEzsignfolderDelayedsenddate,omitempty"`
	AFkiEzsignfoldersignerassociationID []int32 `json:"a_fkiEzsignfoldersignerassociationID"`
}

type _EzsignfolderSendV3Request EzsignfolderSendV3Request

// NewEzsignfolderSendV3Request instantiates a new EzsignfolderSendV3Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderSendV3Request(aFkiEzsignfoldersignerassociationID []int32) *EzsignfolderSendV3Request {
	this := EzsignfolderSendV3Request{}
	var eEzsignfolderMessageorder FieldEEzsignfolderMessageorder = GLOBAL_FIRST
	this.EEzsignfolderMessageorder = &eEzsignfolderMessageorder
	this.AFkiEzsignfoldersignerassociationID = aFkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignfolderSendV3RequestWithDefaults instantiates a new EzsignfolderSendV3Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderSendV3RequestWithDefaults() *EzsignfolderSendV3Request {
	this := EzsignfolderSendV3Request{}
	var eEzsignfolderMessageorder FieldEEzsignfolderMessageorder = GLOBAL_FIRST
	this.EEzsignfolderMessageorder = &eEzsignfolderMessageorder
	return &this
}

// GetTEzsignfolderMessage returns the TEzsignfolderMessage field value if set, zero value otherwise.
func (o *EzsignfolderSendV3Request) GetTEzsignfolderMessage() string {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		var ret string
		return ret
	}
	return *o.TEzsignfolderMessage
}

// GetTEzsignfolderMessageOk returns a tuple with the TEzsignfolderMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Request) GetTEzsignfolderMessageOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfolderMessage) {
		return nil, false
	}
	return o.TEzsignfolderMessage, true
}

// HasTEzsignfolderMessage returns a boolean if a field has been set.
func (o *EzsignfolderSendV3Request) HasTEzsignfolderMessage() bool {
	if o != nil && !IsNil(o.TEzsignfolderMessage) {
		return true
	}

	return false
}

// SetTEzsignfolderMessage gets a reference to the given string and assigns it to the TEzsignfolderMessage field.
func (o *EzsignfolderSendV3Request) SetTEzsignfolderMessage(v string) {
	o.TEzsignfolderMessage = &v
}

// GetEEzsignfolderMessageorder returns the EEzsignfolderMessageorder field value if set, zero value otherwise.
func (o *EzsignfolderSendV3Request) GetEEzsignfolderMessageorder() FieldEEzsignfolderMessageorder {
	if o == nil || IsNil(o.EEzsignfolderMessageorder) {
		var ret FieldEEzsignfolderMessageorder
		return ret
	}
	return *o.EEzsignfolderMessageorder
}

// GetEEzsignfolderMessageorderOk returns a tuple with the EEzsignfolderMessageorder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Request) GetEEzsignfolderMessageorderOk() (*FieldEEzsignfolderMessageorder, bool) {
	if o == nil || IsNil(o.EEzsignfolderMessageorder) {
		return nil, false
	}
	return o.EEzsignfolderMessageorder, true
}

// HasEEzsignfolderMessageorder returns a boolean if a field has been set.
func (o *EzsignfolderSendV3Request) HasEEzsignfolderMessageorder() bool {
	if o != nil && !IsNil(o.EEzsignfolderMessageorder) {
		return true
	}

	return false
}

// SetEEzsignfolderMessageorder gets a reference to the given FieldEEzsignfolderMessageorder and assigns it to the EEzsignfolderMessageorder field.
func (o *EzsignfolderSendV3Request) SetEEzsignfolderMessageorder(v FieldEEzsignfolderMessageorder) {
	o.EEzsignfolderMessageorder = &v
}

// GetDtEzsignfolderDelayedsenddate returns the DtEzsignfolderDelayedsenddate field value if set, zero value otherwise.
func (o *EzsignfolderSendV3Request) GetDtEzsignfolderDelayedsenddate() string {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		var ret string
		return ret
	}
	return *o.DtEzsignfolderDelayedsenddate
}

// GetDtEzsignfolderDelayedsenddateOk returns a tuple with the DtEzsignfolderDelayedsenddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Request) GetDtEzsignfolderDelayedsenddateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsignfolderDelayedsenddate) {
		return nil, false
	}
	return o.DtEzsignfolderDelayedsenddate, true
}

// HasDtEzsignfolderDelayedsenddate returns a boolean if a field has been set.
func (o *EzsignfolderSendV3Request) HasDtEzsignfolderDelayedsenddate() bool {
	if o != nil && !IsNil(o.DtEzsignfolderDelayedsenddate) {
		return true
	}

	return false
}

// SetDtEzsignfolderDelayedsenddate gets a reference to the given string and assigns it to the DtEzsignfolderDelayedsenddate field.
func (o *EzsignfolderSendV3Request) SetDtEzsignfolderDelayedsenddate(v string) {
	o.DtEzsignfolderDelayedsenddate = &v
}

// GetAFkiEzsignfoldersignerassociationID returns the AFkiEzsignfoldersignerassociationID field value
func (o *EzsignfolderSendV3Request) GetAFkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AFkiEzsignfoldersignerassociationID
}

// GetAFkiEzsignfoldersignerassociationIDOk returns a tuple with the AFkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Request) GetAFkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AFkiEzsignfoldersignerassociationID, true
}

// SetAFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfolderSendV3Request) SetAFkiEzsignfoldersignerassociationID(v []int32) {
	o.AFkiEzsignfoldersignerassociationID = v
}

func (o EzsignfolderSendV3Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderSendV3Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TEzsignfolderMessage) {
		toSerialize["tEzsignfolderMessage"] = o.TEzsignfolderMessage
	}
	if !IsNil(o.EEzsignfolderMessageorder) {
		toSerialize["eEzsignfolderMessageorder"] = o.EEzsignfolderMessageorder
	}
	if !IsNil(o.DtEzsignfolderDelayedsenddate) {
		toSerialize["dtEzsignfolderDelayedsenddate"] = o.DtEzsignfolderDelayedsenddate
	}
	toSerialize["a_fkiEzsignfoldersignerassociationID"] = o.AFkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignfolderSendV3Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_fkiEzsignfoldersignerassociationID",
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

	varEzsignfolderSendV3Request := _EzsignfolderSendV3Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderSendV3Request)

	if err != nil {
		return err
	}

	*o = EzsignfolderSendV3Request(varEzsignfolderSendV3Request)

	return err
}

type NullableEzsignfolderSendV3Request struct {
	value *EzsignfolderSendV3Request
	isSet bool
}

func (v NullableEzsignfolderSendV3Request) Get() *EzsignfolderSendV3Request {
	return v.value
}

func (v *NullableEzsignfolderSendV3Request) Set(val *EzsignfolderSendV3Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderSendV3Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderSendV3Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderSendV3Request(val *EzsignfolderSendV3Request) *NullableEzsignfolderSendV3Request {
	return &NullableEzsignfolderSendV3Request{value: val, isSet: true}
}

func (v NullableEzsignfolderSendV3Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderSendV3Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


