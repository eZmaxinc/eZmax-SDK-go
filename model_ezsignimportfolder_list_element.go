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
	"bytes"
	"fmt"
)

// checks if the EzsignimportfolderListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignimportfolderListElement{}

// EzsignimportfolderListElement A Ezsignimportfolder List Element
type EzsignimportfolderListElement struct {
	// The unique ID of the Ezsignimportfolder
	PkiEzsignimportfolderID int32 `json:"pkiEzsignimportfolderID"`
	// The name of the Ezsignimportfolder
	SEzsignimportfolderName string `json:"sEzsignimportfolderName" validate:"regexp=^.{0,150}$"`
	// The date and time at which the object was created
	DtCreatedDate *string `json:"dtCreatedDate,omitempty"`
	// The date and time at which the object was last modified
	DtModifiedDate *string `json:"dtModifiedDate,omitempty"`
	// The count of Ezsignimportdocument.
	ITotalEzsignimportdocument *int32 `json:"iTotalEzsignimportdocument,omitempty"`
	// The count of Ezsignimportdocument not imported in an Ezsignfolder.
	ITotalEzsignimportdocumentNotImported *int32 `json:"iTotalEzsignimportdocumentNotImported,omitempty"`
	EEzsignimportfolderStatus *ComputedEEzsignimportfolderStatus `json:"eEzsignimportfolderStatus,omitempty"`
}

type _EzsignimportfolderListElement EzsignimportfolderListElement

// NewEzsignimportfolderListElement instantiates a new EzsignimportfolderListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignimportfolderListElement(pkiEzsignimportfolderID int32, sEzsignimportfolderName string) *EzsignimportfolderListElement {
	this := EzsignimportfolderListElement{}
	this.PkiEzsignimportfolderID = pkiEzsignimportfolderID
	this.SEzsignimportfolderName = sEzsignimportfolderName
	return &this
}

// NewEzsignimportfolderListElementWithDefaults instantiates a new EzsignimportfolderListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignimportfolderListElementWithDefaults() *EzsignimportfolderListElement {
	this := EzsignimportfolderListElement{}
	return &this
}

// GetPkiEzsignimportfolderID returns the PkiEzsignimportfolderID field value
func (o *EzsignimportfolderListElement) GetPkiEzsignimportfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignimportfolderID
}

// GetPkiEzsignimportfolderIDOk returns a tuple with the PkiEzsignimportfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetPkiEzsignimportfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignimportfolderID, true
}

// SetPkiEzsignimportfolderID sets field value
func (o *EzsignimportfolderListElement) SetPkiEzsignimportfolderID(v int32) {
	o.PkiEzsignimportfolderID = v
}

// GetSEzsignimportfolderName returns the SEzsignimportfolderName field value
func (o *EzsignimportfolderListElement) GetSEzsignimportfolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignimportfolderName
}

// GetSEzsignimportfolderNameOk returns a tuple with the SEzsignimportfolderName field value
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetSEzsignimportfolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignimportfolderName, true
}

// SetSEzsignimportfolderName sets field value
func (o *EzsignimportfolderListElement) SetSEzsignimportfolderName(v string) {
	o.SEzsignimportfolderName = v
}

// GetDtCreatedDate returns the DtCreatedDate field value if set, zero value otherwise.
func (o *EzsignimportfolderListElement) GetDtCreatedDate() string {
	if o == nil || IsNil(o.DtCreatedDate) {
		var ret string
		return ret
	}
	return *o.DtCreatedDate
}

// GetDtCreatedDateOk returns a tuple with the DtCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetDtCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.DtCreatedDate) {
		return nil, false
	}
	return o.DtCreatedDate, true
}

// HasDtCreatedDate returns a boolean if a field has been set.
func (o *EzsignimportfolderListElement) HasDtCreatedDate() bool {
	if o != nil && !IsNil(o.DtCreatedDate) {
		return true
	}

	return false
}

// SetDtCreatedDate gets a reference to the given string and assigns it to the DtCreatedDate field.
func (o *EzsignimportfolderListElement) SetDtCreatedDate(v string) {
	o.DtCreatedDate = &v
}

// GetDtModifiedDate returns the DtModifiedDate field value if set, zero value otherwise.
func (o *EzsignimportfolderListElement) GetDtModifiedDate() string {
	if o == nil || IsNil(o.DtModifiedDate) {
		var ret string
		return ret
	}
	return *o.DtModifiedDate
}

// GetDtModifiedDateOk returns a tuple with the DtModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetDtModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.DtModifiedDate) {
		return nil, false
	}
	return o.DtModifiedDate, true
}

// HasDtModifiedDate returns a boolean if a field has been set.
func (o *EzsignimportfolderListElement) HasDtModifiedDate() bool {
	if o != nil && !IsNil(o.DtModifiedDate) {
		return true
	}

	return false
}

// SetDtModifiedDate gets a reference to the given string and assigns it to the DtModifiedDate field.
func (o *EzsignimportfolderListElement) SetDtModifiedDate(v string) {
	o.DtModifiedDate = &v
}

// GetITotalEzsignimportdocument returns the ITotalEzsignimportdocument field value if set, zero value otherwise.
func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocument() int32 {
	if o == nil || IsNil(o.ITotalEzsignimportdocument) {
		var ret int32
		return ret
	}
	return *o.ITotalEzsignimportdocument
}

// GetITotalEzsignimportdocumentOk returns a tuple with the ITotalEzsignimportdocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentOk() (*int32, bool) {
	if o == nil || IsNil(o.ITotalEzsignimportdocument) {
		return nil, false
	}
	return o.ITotalEzsignimportdocument, true
}

// HasITotalEzsignimportdocument returns a boolean if a field has been set.
func (o *EzsignimportfolderListElement) HasITotalEzsignimportdocument() bool {
	if o != nil && !IsNil(o.ITotalEzsignimportdocument) {
		return true
	}

	return false
}

// SetITotalEzsignimportdocument gets a reference to the given int32 and assigns it to the ITotalEzsignimportdocument field.
func (o *EzsignimportfolderListElement) SetITotalEzsignimportdocument(v int32) {
	o.ITotalEzsignimportdocument = &v
}

// GetITotalEzsignimportdocumentNotImported returns the ITotalEzsignimportdocumentNotImported field value if set, zero value otherwise.
func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentNotImported() int32 {
	if o == nil || IsNil(o.ITotalEzsignimportdocumentNotImported) {
		var ret int32
		return ret
	}
	return *o.ITotalEzsignimportdocumentNotImported
}

// GetITotalEzsignimportdocumentNotImportedOk returns a tuple with the ITotalEzsignimportdocumentNotImported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetITotalEzsignimportdocumentNotImportedOk() (*int32, bool) {
	if o == nil || IsNil(o.ITotalEzsignimportdocumentNotImported) {
		return nil, false
	}
	return o.ITotalEzsignimportdocumentNotImported, true
}

// HasITotalEzsignimportdocumentNotImported returns a boolean if a field has been set.
func (o *EzsignimportfolderListElement) HasITotalEzsignimportdocumentNotImported() bool {
	if o != nil && !IsNil(o.ITotalEzsignimportdocumentNotImported) {
		return true
	}

	return false
}

// SetITotalEzsignimportdocumentNotImported gets a reference to the given int32 and assigns it to the ITotalEzsignimportdocumentNotImported field.
func (o *EzsignimportfolderListElement) SetITotalEzsignimportdocumentNotImported(v int32) {
	o.ITotalEzsignimportdocumentNotImported = &v
}

// GetEEzsignimportfolderStatus returns the EEzsignimportfolderStatus field value if set, zero value otherwise.
func (o *EzsignimportfolderListElement) GetEEzsignimportfolderStatus() ComputedEEzsignimportfolderStatus {
	if o == nil || IsNil(o.EEzsignimportfolderStatus) {
		var ret ComputedEEzsignimportfolderStatus
		return ret
	}
	return *o.EEzsignimportfolderStatus
}

// GetEEzsignimportfolderStatusOk returns a tuple with the EEzsignimportfolderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignimportfolderListElement) GetEEzsignimportfolderStatusOk() (*ComputedEEzsignimportfolderStatus, bool) {
	if o == nil || IsNil(o.EEzsignimportfolderStatus) {
		return nil, false
	}
	return o.EEzsignimportfolderStatus, true
}

// HasEEzsignimportfolderStatus returns a boolean if a field has been set.
func (o *EzsignimportfolderListElement) HasEEzsignimportfolderStatus() bool {
	if o != nil && !IsNil(o.EEzsignimportfolderStatus) {
		return true
	}

	return false
}

// SetEEzsignimportfolderStatus gets a reference to the given ComputedEEzsignimportfolderStatus and assigns it to the EEzsignimportfolderStatus field.
func (o *EzsignimportfolderListElement) SetEEzsignimportfolderStatus(v ComputedEEzsignimportfolderStatus) {
	o.EEzsignimportfolderStatus = &v
}

func (o EzsignimportfolderListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignimportfolderListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignimportfolderID"] = o.PkiEzsignimportfolderID
	toSerialize["sEzsignimportfolderName"] = o.SEzsignimportfolderName
	if !IsNil(o.DtCreatedDate) {
		toSerialize["dtCreatedDate"] = o.DtCreatedDate
	}
	if !IsNil(o.DtModifiedDate) {
		toSerialize["dtModifiedDate"] = o.DtModifiedDate
	}
	if !IsNil(o.ITotalEzsignimportdocument) {
		toSerialize["iTotalEzsignimportdocument"] = o.ITotalEzsignimportdocument
	}
	if !IsNil(o.ITotalEzsignimportdocumentNotImported) {
		toSerialize["iTotalEzsignimportdocumentNotImported"] = o.ITotalEzsignimportdocumentNotImported
	}
	if !IsNil(o.EEzsignimportfolderStatus) {
		toSerialize["eEzsignimportfolderStatus"] = o.EEzsignimportfolderStatus
	}
	return toSerialize, nil
}

func (o *EzsignimportfolderListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignimportfolderID",
		"sEzsignimportfolderName",
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

	varEzsignimportfolderListElement := _EzsignimportfolderListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignimportfolderListElement)

	if err != nil {
		return err
	}

	*o = EzsignimportfolderListElement(varEzsignimportfolderListElement)

	return err
}

type NullableEzsignimportfolderListElement struct {
	value *EzsignimportfolderListElement
	isSet bool
}

func (v NullableEzsignimportfolderListElement) Get() *EzsignimportfolderListElement {
	return v.value
}

func (v *NullableEzsignimportfolderListElement) Set(val *EzsignimportfolderListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignimportfolderListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignimportfolderListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignimportfolderListElement(val *EzsignimportfolderListElement) *NullableEzsignimportfolderListElement {
	return &NullableEzsignimportfolderListElement{value: val, isSet: true}
}

func (v NullableEzsignimportfolderListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignimportfolderListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


