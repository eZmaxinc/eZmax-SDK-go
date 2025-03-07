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

// checks if the CreditcardmerchantResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardmerchantResponseCompound{}

// CreditcardmerchantResponseCompound A Creditcardmerchant Object
type CreditcardmerchantResponseCompound struct {
	// The unique ID of the Creditcardmerchant
	PkiCreditcardmerchantID int32 `json:"pkiCreditcardmerchantID"`
	// The unique ID of the Bankaccount
	FkiBankaccountID int32 `json:"fkiBankaccountID"`
	// The name of the bank
	SBankaccountBankname *string `json:"sBankaccountBankname,omitempty"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID *int32 `json:"fkiLanguageID,omitempty"`
	// The Name of the Language in the language of the requester
	SLanguageNameX *string `json:"sLanguageNameX,omitempty"`
	// Whether if visa are denied
	BCreditcardmerchantDenyvisa bool `json:"bCreditcardmerchantDenyvisa"`
	// Whether if mastercard are denied
	BCreditcardmerchantDenymastercard bool `json:"bCreditcardmerchantDenymastercard"`
	// Whether if amex are denied
	BCreditcardmerchantDenyamex bool `json:"bCreditcardmerchantDenyamex"`
	// Whether the creditcardmerchant is active or not
	BCreditcardmerchantIsactive bool `json:"bCreditcardmerchantIsactive"`
	// The description of the Creditcardmerchant
	SCreditcardmerchantDescription string `json:"sCreditcardmerchantDescription" validate:"regexp=^.{0,25}$"`
	// The storeid of the Creditcardmerchant
	SCreditcardmerchantStoreid string `json:"sCreditcardmerchantStoreid" validate:"regexp=^.{0,25}$"`
}

type _CreditcardmerchantResponseCompound CreditcardmerchantResponseCompound

// NewCreditcardmerchantResponseCompound instantiates a new CreditcardmerchantResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardmerchantResponseCompound(pkiCreditcardmerchantID int32, fkiBankaccountID int32, bCreditcardmerchantDenyvisa bool, bCreditcardmerchantDenymastercard bool, bCreditcardmerchantDenyamex bool, bCreditcardmerchantIsactive bool, sCreditcardmerchantDescription string, sCreditcardmerchantStoreid string) *CreditcardmerchantResponseCompound {
	this := CreditcardmerchantResponseCompound{}
	this.PkiCreditcardmerchantID = pkiCreditcardmerchantID
	this.FkiBankaccountID = fkiBankaccountID
	this.BCreditcardmerchantDenyvisa = bCreditcardmerchantDenyvisa
	this.BCreditcardmerchantDenymastercard = bCreditcardmerchantDenymastercard
	this.BCreditcardmerchantDenyamex = bCreditcardmerchantDenyamex
	this.BCreditcardmerchantIsactive = bCreditcardmerchantIsactive
	this.SCreditcardmerchantDescription = sCreditcardmerchantDescription
	this.SCreditcardmerchantStoreid = sCreditcardmerchantStoreid
	return &this
}

// NewCreditcardmerchantResponseCompoundWithDefaults instantiates a new CreditcardmerchantResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardmerchantResponseCompoundWithDefaults() *CreditcardmerchantResponseCompound {
	this := CreditcardmerchantResponseCompound{}
	return &this
}

// GetPkiCreditcardmerchantID returns the PkiCreditcardmerchantID field value
func (o *CreditcardmerchantResponseCompound) GetPkiCreditcardmerchantID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCreditcardmerchantID
}

// GetPkiCreditcardmerchantIDOk returns a tuple with the PkiCreditcardmerchantID field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetPkiCreditcardmerchantIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCreditcardmerchantID, true
}

// SetPkiCreditcardmerchantID sets field value
func (o *CreditcardmerchantResponseCompound) SetPkiCreditcardmerchantID(v int32) {
	o.PkiCreditcardmerchantID = v
}

// GetFkiBankaccountID returns the FkiBankaccountID field value
func (o *CreditcardmerchantResponseCompound) GetFkiBankaccountID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBankaccountID
}

// GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetFkiBankaccountIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBankaccountID, true
}

// SetFkiBankaccountID sets field value
func (o *CreditcardmerchantResponseCompound) SetFkiBankaccountID(v int32) {
	o.FkiBankaccountID = v
}

// GetSBankaccountBankname returns the SBankaccountBankname field value if set, zero value otherwise.
func (o *CreditcardmerchantResponseCompound) GetSBankaccountBankname() string {
	if o == nil || IsNil(o.SBankaccountBankname) {
		var ret string
		return ret
	}
	return *o.SBankaccountBankname
}

// GetSBankaccountBanknameOk returns a tuple with the SBankaccountBankname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetSBankaccountBanknameOk() (*string, bool) {
	if o == nil || IsNil(o.SBankaccountBankname) {
		return nil, false
	}
	return o.SBankaccountBankname, true
}

// HasSBankaccountBankname returns a boolean if a field has been set.
func (o *CreditcardmerchantResponseCompound) HasSBankaccountBankname() bool {
	if o != nil && !IsNil(o.SBankaccountBankname) {
		return true
	}

	return false
}

// SetSBankaccountBankname gets a reference to the given string and assigns it to the SBankaccountBankname field.
func (o *CreditcardmerchantResponseCompound) SetSBankaccountBankname(v string) {
	o.SBankaccountBankname = &v
}

// GetFkiLanguageID returns the FkiLanguageID field value if set, zero value otherwise.
func (o *CreditcardmerchantResponseCompound) GetFkiLanguageID() int32 {
	if o == nil || IsNil(o.FkiLanguageID) {
		var ret int32
		return ret
	}
	return *o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiLanguageID) {
		return nil, false
	}
	return o.FkiLanguageID, true
}

// HasFkiLanguageID returns a boolean if a field has been set.
func (o *CreditcardmerchantResponseCompound) HasFkiLanguageID() bool {
	if o != nil && !IsNil(o.FkiLanguageID) {
		return true
	}

	return false
}

// SetFkiLanguageID gets a reference to the given int32 and assigns it to the FkiLanguageID field.
func (o *CreditcardmerchantResponseCompound) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = &v
}

// GetSLanguageNameX returns the SLanguageNameX field value if set, zero value otherwise.
func (o *CreditcardmerchantResponseCompound) GetSLanguageNameX() string {
	if o == nil || IsNil(o.SLanguageNameX) {
		var ret string
		return ret
	}
	return *o.SLanguageNameX
}

// GetSLanguageNameXOk returns a tuple with the SLanguageNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetSLanguageNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SLanguageNameX) {
		return nil, false
	}
	return o.SLanguageNameX, true
}

// HasSLanguageNameX returns a boolean if a field has been set.
func (o *CreditcardmerchantResponseCompound) HasSLanguageNameX() bool {
	if o != nil && !IsNil(o.SLanguageNameX) {
		return true
	}

	return false
}

// SetSLanguageNameX gets a reference to the given string and assigns it to the SLanguageNameX field.
func (o *CreditcardmerchantResponseCompound) SetSLanguageNameX(v string) {
	o.SLanguageNameX = &v
}

// GetBCreditcardmerchantDenyvisa returns the BCreditcardmerchantDenyvisa field value
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyvisa() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardmerchantDenyvisa
}

// GetBCreditcardmerchantDenyvisaOk returns a tuple with the BCreditcardmerchantDenyvisa field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyvisaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardmerchantDenyvisa, true
}

// SetBCreditcardmerchantDenyvisa sets field value
func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenyvisa(v bool) {
	o.BCreditcardmerchantDenyvisa = v
}

// GetBCreditcardmerchantDenymastercard returns the BCreditcardmerchantDenymastercard field value
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenymastercard() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardmerchantDenymastercard
}

// GetBCreditcardmerchantDenymastercardOk returns a tuple with the BCreditcardmerchantDenymastercard field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenymastercardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardmerchantDenymastercard, true
}

// SetBCreditcardmerchantDenymastercard sets field value
func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenymastercard(v bool) {
	o.BCreditcardmerchantDenymastercard = v
}

// GetBCreditcardmerchantDenyamex returns the BCreditcardmerchantDenyamex field value
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyamex() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardmerchantDenyamex
}

// GetBCreditcardmerchantDenyamexOk returns a tuple with the BCreditcardmerchantDenyamex field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantDenyamexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardmerchantDenyamex, true
}

// SetBCreditcardmerchantDenyamex sets field value
func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantDenyamex(v bool) {
	o.BCreditcardmerchantDenyamex = v
}

// GetBCreditcardmerchantIsactive returns the BCreditcardmerchantIsactive field value
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardmerchantIsactive
}

// GetBCreditcardmerchantIsactiveOk returns a tuple with the BCreditcardmerchantIsactive field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetBCreditcardmerchantIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardmerchantIsactive, true
}

// SetBCreditcardmerchantIsactive sets field value
func (o *CreditcardmerchantResponseCompound) SetBCreditcardmerchantIsactive(v bool) {
	o.BCreditcardmerchantIsactive = v
}

// GetSCreditcardmerchantDescription returns the SCreditcardmerchantDescription field value
func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardmerchantDescription
}

// GetSCreditcardmerchantDescriptionOk returns a tuple with the SCreditcardmerchantDescription field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardmerchantDescription, true
}

// SetSCreditcardmerchantDescription sets field value
func (o *CreditcardmerchantResponseCompound) SetSCreditcardmerchantDescription(v string) {
	o.SCreditcardmerchantDescription = v
}

// GetSCreditcardmerchantStoreid returns the SCreditcardmerchantStoreid field value
func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantStoreid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardmerchantStoreid
}

// GetSCreditcardmerchantStoreidOk returns a tuple with the SCreditcardmerchantStoreid field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantResponseCompound) GetSCreditcardmerchantStoreidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardmerchantStoreid, true
}

// SetSCreditcardmerchantStoreid sets field value
func (o *CreditcardmerchantResponseCompound) SetSCreditcardmerchantStoreid(v string) {
	o.SCreditcardmerchantStoreid = v
}

func (o CreditcardmerchantResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardmerchantResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCreditcardmerchantID"] = o.PkiCreditcardmerchantID
	toSerialize["fkiBankaccountID"] = o.FkiBankaccountID
	if !IsNil(o.SBankaccountBankname) {
		toSerialize["sBankaccountBankname"] = o.SBankaccountBankname
	}
	if !IsNil(o.FkiLanguageID) {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if !IsNil(o.SLanguageNameX) {
		toSerialize["sLanguageNameX"] = o.SLanguageNameX
	}
	toSerialize["bCreditcardmerchantDenyvisa"] = o.BCreditcardmerchantDenyvisa
	toSerialize["bCreditcardmerchantDenymastercard"] = o.BCreditcardmerchantDenymastercard
	toSerialize["bCreditcardmerchantDenyamex"] = o.BCreditcardmerchantDenyamex
	toSerialize["bCreditcardmerchantIsactive"] = o.BCreditcardmerchantIsactive
	toSerialize["sCreditcardmerchantDescription"] = o.SCreditcardmerchantDescription
	toSerialize["sCreditcardmerchantStoreid"] = o.SCreditcardmerchantStoreid
	return toSerialize, nil
}

func (o *CreditcardmerchantResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiCreditcardmerchantID",
		"fkiBankaccountID",
		"bCreditcardmerchantDenyvisa",
		"bCreditcardmerchantDenymastercard",
		"bCreditcardmerchantDenyamex",
		"bCreditcardmerchantIsactive",
		"sCreditcardmerchantDescription",
		"sCreditcardmerchantStoreid",
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

	varCreditcardmerchantResponseCompound := _CreditcardmerchantResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardmerchantResponseCompound)

	if err != nil {
		return err
	}

	*o = CreditcardmerchantResponseCompound(varCreditcardmerchantResponseCompound)

	return err
}

type NullableCreditcardmerchantResponseCompound struct {
	value *CreditcardmerchantResponseCompound
	isSet bool
}

func (v NullableCreditcardmerchantResponseCompound) Get() *CreditcardmerchantResponseCompound {
	return v.value
}

func (v *NullableCreditcardmerchantResponseCompound) Set(val *CreditcardmerchantResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardmerchantResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardmerchantResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardmerchantResponseCompound(val *CreditcardmerchantResponseCompound) *NullableCreditcardmerchantResponseCompound {
	return &NullableCreditcardmerchantResponseCompound{value: val, isSet: true}
}

func (v NullableCreditcardmerchantResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardmerchantResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


