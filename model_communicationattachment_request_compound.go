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
)

// checks if the CommunicationattachmentRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationattachmentRequestCompound{}

// CommunicationattachmentRequestCompound A Communicationattachment Object and children
type CommunicationattachmentRequestCompound struct {
	// The unique ID of the Communicationattachment
	PkiCommunicationattachmentID *int32 `json:"pkiCommunicationattachmentID,omitempty"`
	// The unique ID of the Attachment.
	FkiAttachmentID *int32 `json:"fkiAttachmentID,omitempty"`
	// The unique ID of the Invoice.
	FkiInvoiceID *int32 `json:"fkiInvoiceID,omitempty"`
	// The unique ID of the Salarypreparation.
	FkiSalarypreparationID *int32 `json:"fkiSalarypreparationID,omitempty"`
}

// NewCommunicationattachmentRequestCompound instantiates a new CommunicationattachmentRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationattachmentRequestCompound() *CommunicationattachmentRequestCompound {
	this := CommunicationattachmentRequestCompound{}
	return &this
}

// NewCommunicationattachmentRequestCompoundWithDefaults instantiates a new CommunicationattachmentRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationattachmentRequestCompoundWithDefaults() *CommunicationattachmentRequestCompound {
	this := CommunicationattachmentRequestCompound{}
	return &this
}

// GetPkiCommunicationattachmentID returns the PkiCommunicationattachmentID field value if set, zero value otherwise.
func (o *CommunicationattachmentRequestCompound) GetPkiCommunicationattachmentID() int32 {
	if o == nil || IsNil(o.PkiCommunicationattachmentID) {
		var ret int32
		return ret
	}
	return *o.PkiCommunicationattachmentID
}

// GetPkiCommunicationattachmentIDOk returns a tuple with the PkiCommunicationattachmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationattachmentRequestCompound) GetPkiCommunicationattachmentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiCommunicationattachmentID) {
		return nil, false
	}
	return o.PkiCommunicationattachmentID, true
}

// HasPkiCommunicationattachmentID returns a boolean if a field has been set.
func (o *CommunicationattachmentRequestCompound) HasPkiCommunicationattachmentID() bool {
	if o != nil && !IsNil(o.PkiCommunicationattachmentID) {
		return true
	}

	return false
}

// SetPkiCommunicationattachmentID gets a reference to the given int32 and assigns it to the PkiCommunicationattachmentID field.
func (o *CommunicationattachmentRequestCompound) SetPkiCommunicationattachmentID(v int32) {
	o.PkiCommunicationattachmentID = &v
}

// GetFkiAttachmentID returns the FkiAttachmentID field value if set, zero value otherwise.
func (o *CommunicationattachmentRequestCompound) GetFkiAttachmentID() int32 {
	if o == nil || IsNil(o.FkiAttachmentID) {
		var ret int32
		return ret
	}
	return *o.FkiAttachmentID
}

// GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationattachmentRequestCompound) GetFkiAttachmentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiAttachmentID) {
		return nil, false
	}
	return o.FkiAttachmentID, true
}

// HasFkiAttachmentID returns a boolean if a field has been set.
func (o *CommunicationattachmentRequestCompound) HasFkiAttachmentID() bool {
	if o != nil && !IsNil(o.FkiAttachmentID) {
		return true
	}

	return false
}

// SetFkiAttachmentID gets a reference to the given int32 and assigns it to the FkiAttachmentID field.
func (o *CommunicationattachmentRequestCompound) SetFkiAttachmentID(v int32) {
	o.FkiAttachmentID = &v
}

// GetFkiInvoiceID returns the FkiInvoiceID field value if set, zero value otherwise.
func (o *CommunicationattachmentRequestCompound) GetFkiInvoiceID() int32 {
	if o == nil || IsNil(o.FkiInvoiceID) {
		var ret int32
		return ret
	}
	return *o.FkiInvoiceID
}

// GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationattachmentRequestCompound) GetFkiInvoiceIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiInvoiceID) {
		return nil, false
	}
	return o.FkiInvoiceID, true
}

// HasFkiInvoiceID returns a boolean if a field has been set.
func (o *CommunicationattachmentRequestCompound) HasFkiInvoiceID() bool {
	if o != nil && !IsNil(o.FkiInvoiceID) {
		return true
	}

	return false
}

// SetFkiInvoiceID gets a reference to the given int32 and assigns it to the FkiInvoiceID field.
func (o *CommunicationattachmentRequestCompound) SetFkiInvoiceID(v int32) {
	o.FkiInvoiceID = &v
}

// GetFkiSalarypreparationID returns the FkiSalarypreparationID field value if set, zero value otherwise.
func (o *CommunicationattachmentRequestCompound) GetFkiSalarypreparationID() int32 {
	if o == nil || IsNil(o.FkiSalarypreparationID) {
		var ret int32
		return ret
	}
	return *o.FkiSalarypreparationID
}

// GetFkiSalarypreparationIDOk returns a tuple with the FkiSalarypreparationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationattachmentRequestCompound) GetFkiSalarypreparationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiSalarypreparationID) {
		return nil, false
	}
	return o.FkiSalarypreparationID, true
}

// HasFkiSalarypreparationID returns a boolean if a field has been set.
func (o *CommunicationattachmentRequestCompound) HasFkiSalarypreparationID() bool {
	if o != nil && !IsNil(o.FkiSalarypreparationID) {
		return true
	}

	return false
}

// SetFkiSalarypreparationID gets a reference to the given int32 and assigns it to the FkiSalarypreparationID field.
func (o *CommunicationattachmentRequestCompound) SetFkiSalarypreparationID(v int32) {
	o.FkiSalarypreparationID = &v
}

func (o CommunicationattachmentRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationattachmentRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiCommunicationattachmentID) {
		toSerialize["pkiCommunicationattachmentID"] = o.PkiCommunicationattachmentID
	}
	if !IsNil(o.FkiAttachmentID) {
		toSerialize["fkiAttachmentID"] = o.FkiAttachmentID
	}
	if !IsNil(o.FkiInvoiceID) {
		toSerialize["fkiInvoiceID"] = o.FkiInvoiceID
	}
	if !IsNil(o.FkiSalarypreparationID) {
		toSerialize["fkiSalarypreparationID"] = o.FkiSalarypreparationID
	}
	return toSerialize, nil
}

type NullableCommunicationattachmentRequestCompound struct {
	value *CommunicationattachmentRequestCompound
	isSet bool
}

func (v NullableCommunicationattachmentRequestCompound) Get() *CommunicationattachmentRequestCompound {
	return v.value
}

func (v *NullableCommunicationattachmentRequestCompound) Set(val *CommunicationattachmentRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationattachmentRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationattachmentRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationattachmentRequestCompound(val *CommunicationattachmentRequestCompound) *NullableCommunicationattachmentRequestCompound {
	return &NullableCommunicationattachmentRequestCompound{value: val, isSet: true}
}

func (v NullableCommunicationattachmentRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationattachmentRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


