/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the EzsigntemplateCopyV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateCopyV1Request{}

// EzsigntemplateCopyV1Request Request for POST /1/object/ezsigntemplate/{pkiEzsigntemplateID}/copy
type EzsigntemplateCopyV1Request struct {
	AFkiEzsignfoldertypeID []int32 `json:"a_fkiEzsignfoldertypeID,omitempty"`
	// Whether we shall copy the Ezsigntemplate as a company Ezsigntemplate
	BCopyCompany *bool `json:"bCopyCompany,omitempty"`
	// Whether we shall copy the Ezsigntemplate as a user Ezsigntemplate
	BCopyUser *bool `json:"bCopyUser,omitempty"`
}

// NewEzsigntemplateCopyV1Request instantiates a new EzsigntemplateCopyV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateCopyV1Request() *EzsigntemplateCopyV1Request {
	this := EzsigntemplateCopyV1Request{}
	return &this
}

// NewEzsigntemplateCopyV1RequestWithDefaults instantiates a new EzsigntemplateCopyV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateCopyV1RequestWithDefaults() *EzsigntemplateCopyV1Request {
	this := EzsigntemplateCopyV1Request{}
	return &this
}

// GetAFkiEzsignfoldertypeID returns the AFkiEzsignfoldertypeID field value if set, zero value otherwise.
func (o *EzsigntemplateCopyV1Request) GetAFkiEzsignfoldertypeID() []int32 {
	if o == nil || IsNil(o.AFkiEzsignfoldertypeID) {
		var ret []int32
		return ret
	}
	return o.AFkiEzsignfoldertypeID
}

// GetAFkiEzsignfoldertypeIDOk returns a tuple with the AFkiEzsignfoldertypeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCopyV1Request) GetAFkiEzsignfoldertypeIDOk() ([]int32, bool) {
	if o == nil || IsNil(o.AFkiEzsignfoldertypeID) {
		return nil, false
	}
	return o.AFkiEzsignfoldertypeID, true
}

// HasAFkiEzsignfoldertypeID returns a boolean if a field has been set.
func (o *EzsigntemplateCopyV1Request) HasAFkiEzsignfoldertypeID() bool {
	if o != nil && !IsNil(o.AFkiEzsignfoldertypeID) {
		return true
	}

	return false
}

// SetAFkiEzsignfoldertypeID gets a reference to the given []int32 and assigns it to the AFkiEzsignfoldertypeID field.
func (o *EzsigntemplateCopyV1Request) SetAFkiEzsignfoldertypeID(v []int32) {
	o.AFkiEzsignfoldertypeID = v
}

// GetBCopyCompany returns the BCopyCompany field value if set, zero value otherwise.
func (o *EzsigntemplateCopyV1Request) GetBCopyCompany() bool {
	if o == nil || IsNil(o.BCopyCompany) {
		var ret bool
		return ret
	}
	return *o.BCopyCompany
}

// GetBCopyCompanyOk returns a tuple with the BCopyCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCopyV1Request) GetBCopyCompanyOk() (*bool, bool) {
	if o == nil || IsNil(o.BCopyCompany) {
		return nil, false
	}
	return o.BCopyCompany, true
}

// HasBCopyCompany returns a boolean if a field has been set.
func (o *EzsigntemplateCopyV1Request) HasBCopyCompany() bool {
	if o != nil && !IsNil(o.BCopyCompany) {
		return true
	}

	return false
}

// SetBCopyCompany gets a reference to the given bool and assigns it to the BCopyCompany field.
func (o *EzsigntemplateCopyV1Request) SetBCopyCompany(v bool) {
	o.BCopyCompany = &v
}

// GetBCopyUser returns the BCopyUser field value if set, zero value otherwise.
func (o *EzsigntemplateCopyV1Request) GetBCopyUser() bool {
	if o == nil || IsNil(o.BCopyUser) {
		var ret bool
		return ret
	}
	return *o.BCopyUser
}

// GetBCopyUserOk returns a tuple with the BCopyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplateCopyV1Request) GetBCopyUserOk() (*bool, bool) {
	if o == nil || IsNil(o.BCopyUser) {
		return nil, false
	}
	return o.BCopyUser, true
}

// HasBCopyUser returns a boolean if a field has been set.
func (o *EzsigntemplateCopyV1Request) HasBCopyUser() bool {
	if o != nil && !IsNil(o.BCopyUser) {
		return true
	}

	return false
}

// SetBCopyUser gets a reference to the given bool and assigns it to the BCopyUser field.
func (o *EzsigntemplateCopyV1Request) SetBCopyUser(v bool) {
	o.BCopyUser = &v
}

func (o EzsigntemplateCopyV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateCopyV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AFkiEzsignfoldertypeID) {
		toSerialize["a_fkiEzsignfoldertypeID"] = o.AFkiEzsignfoldertypeID
	}
	if !IsNil(o.BCopyCompany) {
		toSerialize["bCopyCompany"] = o.BCopyCompany
	}
	if !IsNil(o.BCopyUser) {
		toSerialize["bCopyUser"] = o.BCopyUser
	}
	return toSerialize, nil
}

type NullableEzsigntemplateCopyV1Request struct {
	value *EzsigntemplateCopyV1Request
	isSet bool
}

func (v NullableEzsigntemplateCopyV1Request) Get() *EzsigntemplateCopyV1Request {
	return v.value
}

func (v *NullableEzsigntemplateCopyV1Request) Set(val *EzsigntemplateCopyV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateCopyV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateCopyV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateCopyV1Request(val *EzsigntemplateCopyV1Request) *NullableEzsigntemplateCopyV1Request {
	return &NullableEzsigntemplateCopyV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplateCopyV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateCopyV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


