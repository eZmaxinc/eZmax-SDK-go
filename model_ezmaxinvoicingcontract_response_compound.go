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

// checks if the EzmaxinvoicingcontractResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingcontractResponseCompound{}

// EzmaxinvoicingcontractResponseCompound A Ezmaxinvoicingcontract Object
type EzmaxinvoicingcontractResponseCompound struct {
	// The unique ID of the Ezmaxinvoicingcontract
	PkiEzmaxinvoicingcontractID int32 `json:"pkiEzmaxinvoicingcontractID"`
	EEzmaxinvoicingcontractPaymenttype FieldEEzmaxinvoicingcontractPaymenttype `json:"eEzmaxinvoicingcontractPaymenttype"`
	// The length in years of the Ezmaxinvoicingcontract
	IEzmaxinvoicingcontractLength int32 `json:"iEzmaxinvoicingcontractLength"`
	// The start date of the Ezmaxinvoicingcontract
	DtEzmaxinvoicingcontractStart string `json:"dtEzmaxinvoicingcontractStart"`
	// The end date of the Ezmaxinvoicingcontract
	DtEzmaxinvoicingcontractEnd string `json:"dtEzmaxinvoicingcontractEnd"`
	// The price of the license
	DEzmaxinvoicingcontractLicense string `json:"dEzmaxinvoicingcontractLicense"`
	// The price for 121QA
	DEzmaxinvoicingcontract121qa string `json:"dEzmaxinvoicingcontract121qa"`
	// Whether eZsign is for all agents
	BEzmaxinvoicingcontractEzsignallagents bool `json:"bEzmaxinvoicingcontractEzsignallagents"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _EzmaxinvoicingcontractResponseCompound EzmaxinvoicingcontractResponseCompound

// NewEzmaxinvoicingcontractResponseCompound instantiates a new EzmaxinvoicingcontractResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingcontractResponseCompound(pkiEzmaxinvoicingcontractID int32, eEzmaxinvoicingcontractPaymenttype FieldEEzmaxinvoicingcontractPaymenttype, iEzmaxinvoicingcontractLength int32, dtEzmaxinvoicingcontractStart string, dtEzmaxinvoicingcontractEnd string, dEzmaxinvoicingcontractLicense string, dEzmaxinvoicingcontract121qa string, bEzmaxinvoicingcontractEzsignallagents bool, objAudit CommonAudit) *EzmaxinvoicingcontractResponseCompound {
	this := EzmaxinvoicingcontractResponseCompound{}
	this.PkiEzmaxinvoicingcontractID = pkiEzmaxinvoicingcontractID
	this.EEzmaxinvoicingcontractPaymenttype = eEzmaxinvoicingcontractPaymenttype
	this.IEzmaxinvoicingcontractLength = iEzmaxinvoicingcontractLength
	this.DtEzmaxinvoicingcontractStart = dtEzmaxinvoicingcontractStart
	this.DtEzmaxinvoicingcontractEnd = dtEzmaxinvoicingcontractEnd
	this.DEzmaxinvoicingcontractLicense = dEzmaxinvoicingcontractLicense
	this.DEzmaxinvoicingcontract121qa = dEzmaxinvoicingcontract121qa
	this.BEzmaxinvoicingcontractEzsignallagents = bEzmaxinvoicingcontractEzsignallagents
	this.ObjAudit = objAudit
	return &this
}

// NewEzmaxinvoicingcontractResponseCompoundWithDefaults instantiates a new EzmaxinvoicingcontractResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingcontractResponseCompoundWithDefaults() *EzmaxinvoicingcontractResponseCompound {
	this := EzmaxinvoicingcontractResponseCompound{}
	return &this
}

// GetPkiEzmaxinvoicingcontractID returns the PkiEzmaxinvoicingcontractID field value
func (o *EzmaxinvoicingcontractResponseCompound) GetPkiEzmaxinvoicingcontractID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzmaxinvoicingcontractID
}

// GetPkiEzmaxinvoicingcontractIDOk returns a tuple with the PkiEzmaxinvoicingcontractID field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetPkiEzmaxinvoicingcontractIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzmaxinvoicingcontractID, true
}

// SetPkiEzmaxinvoicingcontractID sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetPkiEzmaxinvoicingcontractID(v int32) {
	o.PkiEzmaxinvoicingcontractID = v
}

// GetEEzmaxinvoicingcontractPaymenttype returns the EEzmaxinvoicingcontractPaymenttype field value
func (o *EzmaxinvoicingcontractResponseCompound) GetEEzmaxinvoicingcontractPaymenttype() FieldEEzmaxinvoicingcontractPaymenttype {
	if o == nil {
		var ret FieldEEzmaxinvoicingcontractPaymenttype
		return ret
	}

	return o.EEzmaxinvoicingcontractPaymenttype
}

// GetEEzmaxinvoicingcontractPaymenttypeOk returns a tuple with the EEzmaxinvoicingcontractPaymenttype field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetEEzmaxinvoicingcontractPaymenttypeOk() (*FieldEEzmaxinvoicingcontractPaymenttype, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EEzmaxinvoicingcontractPaymenttype, true
}

// SetEEzmaxinvoicingcontractPaymenttype sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetEEzmaxinvoicingcontractPaymenttype(v FieldEEzmaxinvoicingcontractPaymenttype) {
	o.EEzmaxinvoicingcontractPaymenttype = v
}

// GetIEzmaxinvoicingcontractLength returns the IEzmaxinvoicingcontractLength field value
func (o *EzmaxinvoicingcontractResponseCompound) GetIEzmaxinvoicingcontractLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzmaxinvoicingcontractLength
}

// GetIEzmaxinvoicingcontractLengthOk returns a tuple with the IEzmaxinvoicingcontractLength field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetIEzmaxinvoicingcontractLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzmaxinvoicingcontractLength, true
}

// SetIEzmaxinvoicingcontractLength sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetIEzmaxinvoicingcontractLength(v int32) {
	o.IEzmaxinvoicingcontractLength = v
}

// GetDtEzmaxinvoicingcontractStart returns the DtEzmaxinvoicingcontractStart field value
func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzmaxinvoicingcontractStart
}

// GetDtEzmaxinvoicingcontractStartOk returns a tuple with the DtEzmaxinvoicingcontractStart field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzmaxinvoicingcontractStart, true
}

// SetDtEzmaxinvoicingcontractStart sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetDtEzmaxinvoicingcontractStart(v string) {
	o.DtEzmaxinvoicingcontractStart = v
}

// GetDtEzmaxinvoicingcontractEnd returns the DtEzmaxinvoicingcontractEnd field value
func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzmaxinvoicingcontractEnd
}

// GetDtEzmaxinvoicingcontractEndOk returns a tuple with the DtEzmaxinvoicingcontractEnd field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetDtEzmaxinvoicingcontractEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtEzmaxinvoicingcontractEnd, true
}

// SetDtEzmaxinvoicingcontractEnd sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetDtEzmaxinvoicingcontractEnd(v string) {
	o.DtEzmaxinvoicingcontractEnd = v
}

// GetDEzmaxinvoicingcontractLicense returns the DEzmaxinvoicingcontractLicense field value
func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontractLicense() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingcontractLicense
}

// GetDEzmaxinvoicingcontractLicenseOk returns a tuple with the DEzmaxinvoicingcontractLicense field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontractLicenseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingcontractLicense, true
}

// SetDEzmaxinvoicingcontractLicense sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetDEzmaxinvoicingcontractLicense(v string) {
	o.DEzmaxinvoicingcontractLicense = v
}

// GetDEzmaxinvoicingcontract121qa returns the DEzmaxinvoicingcontract121qa field value
func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontract121qa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DEzmaxinvoicingcontract121qa
}

// GetDEzmaxinvoicingcontract121qaOk returns a tuple with the DEzmaxinvoicingcontract121qa field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetDEzmaxinvoicingcontract121qaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DEzmaxinvoicingcontract121qa, true
}

// SetDEzmaxinvoicingcontract121qa sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetDEzmaxinvoicingcontract121qa(v string) {
	o.DEzmaxinvoicingcontract121qa = v
}

// GetBEzmaxinvoicingcontractEzsignallagents returns the BEzmaxinvoicingcontractEzsignallagents field value
func (o *EzmaxinvoicingcontractResponseCompound) GetBEzmaxinvoicingcontractEzsignallagents() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzmaxinvoicingcontractEzsignallagents
}

// GetBEzmaxinvoicingcontractEzsignallagentsOk returns a tuple with the BEzmaxinvoicingcontractEzsignallagents field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetBEzmaxinvoicingcontractEzsignallagentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzmaxinvoicingcontractEzsignallagents, true
}

// SetBEzmaxinvoicingcontractEzsignallagents sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetBEzmaxinvoicingcontractEzsignallagents(v bool) {
	o.BEzmaxinvoicingcontractEzsignallagents = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzmaxinvoicingcontractResponseCompound) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzmaxinvoicingcontractResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzmaxinvoicingcontractResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzmaxinvoicingcontractResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingcontractResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzmaxinvoicingcontractID"] = o.PkiEzmaxinvoicingcontractID
	toSerialize["eEzmaxinvoicingcontractPaymenttype"] = o.EEzmaxinvoicingcontractPaymenttype
	toSerialize["iEzmaxinvoicingcontractLength"] = o.IEzmaxinvoicingcontractLength
	toSerialize["dtEzmaxinvoicingcontractStart"] = o.DtEzmaxinvoicingcontractStart
	toSerialize["dtEzmaxinvoicingcontractEnd"] = o.DtEzmaxinvoicingcontractEnd
	toSerialize["dEzmaxinvoicingcontractLicense"] = o.DEzmaxinvoicingcontractLicense
	toSerialize["dEzmaxinvoicingcontract121qa"] = o.DEzmaxinvoicingcontract121qa
	toSerialize["bEzmaxinvoicingcontractEzsignallagents"] = o.BEzmaxinvoicingcontractEzsignallagents
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *EzmaxinvoicingcontractResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzmaxinvoicingcontractID",
		"eEzmaxinvoicingcontractPaymenttype",
		"iEzmaxinvoicingcontractLength",
		"dtEzmaxinvoicingcontractStart",
		"dtEzmaxinvoicingcontractEnd",
		"dEzmaxinvoicingcontractLicense",
		"dEzmaxinvoicingcontract121qa",
		"bEzmaxinvoicingcontractEzsignallagents",
		"objAudit",
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

	varEzmaxinvoicingcontractResponseCompound := _EzmaxinvoicingcontractResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingcontractResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingcontractResponseCompound(varEzmaxinvoicingcontractResponseCompound)

	return err
}

type NullableEzmaxinvoicingcontractResponseCompound struct {
	value *EzmaxinvoicingcontractResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingcontractResponseCompound) Get() *EzmaxinvoicingcontractResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) Set(val *EzmaxinvoicingcontractResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingcontractResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingcontractResponseCompound(val *EzmaxinvoicingcontractResponseCompound) *NullableEzmaxinvoicingcontractResponseCompound {
	return &NullableEzmaxinvoicingcontractResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingcontractResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


