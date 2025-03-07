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

// checks if the EzmaxinvoicingsummaryexternalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingsummaryexternalResponseCompound{}

// EzmaxinvoicingsummaryexternalResponseCompound A Ezmaxinvoicingsummaryexternal Object
type EzmaxinvoicingsummaryexternalResponseCompound struct {
	// The unique ID of the Ezmaxinvoicingsummaryexternal
	PkiEzmaxinvoicingsummaryexternalID *int32 `json:"pkiEzmaxinvoicingsummaryexternalID,omitempty"`
	// The unique ID of the Ezmaxinvoicing
	FkiEzmaxinvoicingID *int32 `json:"fkiEzmaxinvoicingID,omitempty"`
	// The unique ID of the Billingentityexternal
	FkiBillingentityexternalID int32 `json:"fkiBillingentityexternalID"`
	// The description of the Billingentityexternal
	SBillingentityexternalDescription string `json:"sBillingentityexternalDescription"`
	// The description of the Ezmaxinvoicingsummaryexternal
	SEzmaxinvoicingsummaryexternalDescription string `json:"sEzmaxinvoicingsummaryexternalDescription"`
	AObjEzmaxinvoicingsummaryexternaldetail []EzmaxinvoicingsummaryexternaldetailResponseCompound `json:"a_objEzmaxinvoicingsummaryexternaldetail"`
}

type _EzmaxinvoicingsummaryexternalResponseCompound EzmaxinvoicingsummaryexternalResponseCompound

// NewEzmaxinvoicingsummaryexternalResponseCompound instantiates a new EzmaxinvoicingsummaryexternalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingsummaryexternalResponseCompound(fkiBillingentityexternalID int32, sBillingentityexternalDescription string, sEzmaxinvoicingsummaryexternalDescription string, aObjEzmaxinvoicingsummaryexternaldetail []EzmaxinvoicingsummaryexternaldetailResponseCompound) *EzmaxinvoicingsummaryexternalResponseCompound {
	this := EzmaxinvoicingsummaryexternalResponseCompound{}
	this.FkiBillingentityexternalID = fkiBillingentityexternalID
	this.SBillingentityexternalDescription = sBillingentityexternalDescription
	this.SEzmaxinvoicingsummaryexternalDescription = sEzmaxinvoicingsummaryexternalDescription
	this.AObjEzmaxinvoicingsummaryexternaldetail = aObjEzmaxinvoicingsummaryexternaldetail
	return &this
}

// NewEzmaxinvoicingsummaryexternalResponseCompoundWithDefaults instantiates a new EzmaxinvoicingsummaryexternalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingsummaryexternalResponseCompoundWithDefaults() *EzmaxinvoicingsummaryexternalResponseCompound {
	this := EzmaxinvoicingsummaryexternalResponseCompound{}
	return &this
}

// GetPkiEzmaxinvoicingsummaryexternalID returns the PkiEzmaxinvoicingsummaryexternalID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetPkiEzmaxinvoicingsummaryexternalID() int32 {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		var ret int32
		return ret
	}
	return *o.PkiEzmaxinvoicingsummaryexternalID
}

// GetPkiEzmaxinvoicingsummaryexternalIDOk returns a tuple with the PkiEzmaxinvoicingsummaryexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetPkiEzmaxinvoicingsummaryexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		return nil, false
	}
	return o.PkiEzmaxinvoicingsummaryexternalID, true
}

// HasPkiEzmaxinvoicingsummaryexternalID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) HasPkiEzmaxinvoicingsummaryexternalID() bool {
	if o != nil && !IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		return true
	}

	return false
}

// SetPkiEzmaxinvoicingsummaryexternalID gets a reference to the given int32 and assigns it to the PkiEzmaxinvoicingsummaryexternalID field.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetPkiEzmaxinvoicingsummaryexternalID(v int32) {
	o.PkiEzmaxinvoicingsummaryexternalID = &v
}

// GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field value if set, zero value otherwise.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetFkiEzmaxinvoicingID() int32 {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		var ret int32
		return ret
	}
	return *o.FkiEzmaxinvoicingID
}

// GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetFkiEzmaxinvoicingIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzmaxinvoicingID) {
		return nil, false
	}
	return o.FkiEzmaxinvoicingID, true
}

// HasFkiEzmaxinvoicingID returns a boolean if a field has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) HasFkiEzmaxinvoicingID() bool {
	if o != nil && !IsNil(o.FkiEzmaxinvoicingID) {
		return true
	}

	return false
}

// SetFkiEzmaxinvoicingID gets a reference to the given int32 and assigns it to the FkiEzmaxinvoicingID field.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetFkiEzmaxinvoicingID(v int32) {
	o.FkiEzmaxinvoicingID = &v
}

// GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetFkiBillingentityexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityexternalID
}

// GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetFkiBillingentityexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityexternalID, true
}

// SetFkiBillingentityexternalID sets field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetFkiBillingentityexternalID(v int32) {
	o.FkiBillingentityexternalID = v
}

// GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetSBillingentityexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityexternalDescription
}

// GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetSBillingentityexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityexternalDescription, true
}

// SetSBillingentityexternalDescription sets field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetSBillingentityexternalDescription(v string) {
	o.SBillingentityexternalDescription = v
}

// GetSEzmaxinvoicingsummaryexternalDescription returns the SEzmaxinvoicingsummaryexternalDescription field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetSEzmaxinvoicingsummaryexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxinvoicingsummaryexternalDescription
}

// GetSEzmaxinvoicingsummaryexternalDescriptionOk returns a tuple with the SEzmaxinvoicingsummaryexternalDescription field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetSEzmaxinvoicingsummaryexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxinvoicingsummaryexternalDescription, true
}

// SetSEzmaxinvoicingsummaryexternalDescription sets field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetSEzmaxinvoicingsummaryexternalDescription(v string) {
	o.SEzmaxinvoicingsummaryexternalDescription = v
}

// GetAObjEzmaxinvoicingsummaryexternaldetail returns the AObjEzmaxinvoicingsummaryexternaldetail field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetAObjEzmaxinvoicingsummaryexternaldetail() []EzmaxinvoicingsummaryexternaldetailResponseCompound {
	if o == nil {
		var ret []EzmaxinvoicingsummaryexternaldetailResponseCompound
		return ret
	}

	return o.AObjEzmaxinvoicingsummaryexternaldetail
}

// GetAObjEzmaxinvoicingsummaryexternaldetailOk returns a tuple with the AObjEzmaxinvoicingsummaryexternaldetail field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingsummaryexternalResponseCompound) GetAObjEzmaxinvoicingsummaryexternaldetailOk() ([]EzmaxinvoicingsummaryexternaldetailResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzmaxinvoicingsummaryexternaldetail, true
}

// SetAObjEzmaxinvoicingsummaryexternaldetail sets field value
func (o *EzmaxinvoicingsummaryexternalResponseCompound) SetAObjEzmaxinvoicingsummaryexternaldetail(v []EzmaxinvoicingsummaryexternaldetailResponseCompound) {
	o.AObjEzmaxinvoicingsummaryexternaldetail = v
}

func (o EzmaxinvoicingsummaryexternalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingsummaryexternalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzmaxinvoicingsummaryexternalID) {
		toSerialize["pkiEzmaxinvoicingsummaryexternalID"] = o.PkiEzmaxinvoicingsummaryexternalID
	}
	if !IsNil(o.FkiEzmaxinvoicingID) {
		toSerialize["fkiEzmaxinvoicingID"] = o.FkiEzmaxinvoicingID
	}
	toSerialize["fkiBillingentityexternalID"] = o.FkiBillingentityexternalID
	toSerialize["sBillingentityexternalDescription"] = o.SBillingentityexternalDescription
	toSerialize["sEzmaxinvoicingsummaryexternalDescription"] = o.SEzmaxinvoicingsummaryexternalDescription
	toSerialize["a_objEzmaxinvoicingsummaryexternaldetail"] = o.AObjEzmaxinvoicingsummaryexternaldetail
	return toSerialize, nil
}

func (o *EzmaxinvoicingsummaryexternalResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiBillingentityexternalID",
		"sBillingentityexternalDescription",
		"sEzmaxinvoicingsummaryexternalDescription",
		"a_objEzmaxinvoicingsummaryexternaldetail",
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

	varEzmaxinvoicingsummaryexternalResponseCompound := _EzmaxinvoicingsummaryexternalResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingsummaryexternalResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingsummaryexternalResponseCompound(varEzmaxinvoicingsummaryexternalResponseCompound)

	return err
}

type NullableEzmaxinvoicingsummaryexternalResponseCompound struct {
	value *EzmaxinvoicingsummaryexternalResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingsummaryexternalResponseCompound) Get() *EzmaxinvoicingsummaryexternalResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingsummaryexternalResponseCompound) Set(val *EzmaxinvoicingsummaryexternalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingsummaryexternalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingsummaryexternalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingsummaryexternalResponseCompound(val *EzmaxinvoicingsummaryexternalResponseCompound) *NullableEzmaxinvoicingsummaryexternalResponseCompound {
	return &NullableEzmaxinvoicingsummaryexternalResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingsummaryexternalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingsummaryexternalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


