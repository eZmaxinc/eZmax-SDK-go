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

// checks if the EzsigndocumentRequestPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentRequestPatch{}

// EzsigndocumentRequestPatch An Ezsigndocument Object
type EzsigndocumentRequestPatch struct {
	// The maximum date and time at which the Ezsigndocument can be signed.
	DtEzsigndocumentDuedate *string `json:"dtEzsigndocumentDuedate,omitempty"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName *string `json:"sEzsigndocumentName,omitempty"`
}

// NewEzsigndocumentRequestPatch instantiates a new EzsigndocumentRequestPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentRequestPatch() *EzsigndocumentRequestPatch {
	this := EzsigndocumentRequestPatch{}
	return &this
}

// NewEzsigndocumentRequestPatchWithDefaults instantiates a new EzsigndocumentRequestPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentRequestPatchWithDefaults() *EzsigndocumentRequestPatch {
	this := EzsigndocumentRequestPatch{}
	return &this
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value if set, zero value otherwise.
func (o *EzsigndocumentRequestPatch) GetDtEzsigndocumentDuedate() string {
	if o == nil || IsNil(o.DtEzsigndocumentDuedate) {
		var ret string
		return ret
	}
	return *o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestPatch) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil || IsNil(o.DtEzsigndocumentDuedate) {
		return nil, false
	}
	return o.DtEzsigndocumentDuedate, true
}

// HasDtEzsigndocumentDuedate returns a boolean if a field has been set.
func (o *EzsigndocumentRequestPatch) HasDtEzsigndocumentDuedate() bool {
	if o != nil && !IsNil(o.DtEzsigndocumentDuedate) {
		return true
	}

	return false
}

// SetDtEzsigndocumentDuedate gets a reference to the given string and assigns it to the DtEzsigndocumentDuedate field.
func (o *EzsigndocumentRequestPatch) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = &v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value if set, zero value otherwise.
func (o *EzsigndocumentRequestPatch) GetSEzsigndocumentName() string {
	if o == nil || IsNil(o.SEzsigndocumentName) {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequestPatch) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsigndocumentName) {
		return nil, false
	}
	return o.SEzsigndocumentName, true
}

// HasSEzsigndocumentName returns a boolean if a field has been set.
func (o *EzsigndocumentRequestPatch) HasSEzsigndocumentName() bool {
	if o != nil && !IsNil(o.SEzsigndocumentName) {
		return true
	}

	return false
}

// SetSEzsigndocumentName gets a reference to the given string and assigns it to the SEzsigndocumentName field.
func (o *EzsigndocumentRequestPatch) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = &v
}

func (o EzsigndocumentRequestPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentRequestPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DtEzsigndocumentDuedate) {
		toSerialize["dtEzsigndocumentDuedate"] = o.DtEzsigndocumentDuedate
	}
	if !IsNil(o.SEzsigndocumentName) {
		toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	}
	return toSerialize, nil
}

type NullableEzsigndocumentRequestPatch struct {
	value *EzsigndocumentRequestPatch
	isSet bool
}

func (v NullableEzsigndocumentRequestPatch) Get() *EzsigndocumentRequestPatch {
	return v.value
}

func (v *NullableEzsigndocumentRequestPatch) Set(val *EzsigndocumentRequestPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentRequestPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentRequestPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentRequestPatch(val *EzsigndocumentRequestPatch) *NullableEzsigndocumentRequestPatch {
	return &NullableEzsigndocumentRequestPatch{value: val, isSet: true}
}

func (v NullableEzsigndocumentRequestPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentRequestPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


