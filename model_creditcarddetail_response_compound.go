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

// checks if the CreditcarddetailResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcarddetailResponseCompound{}

// CreditcarddetailResponseCompound A Creditcarddetail Object
type CreditcarddetailResponseCompound struct {
	// The unique ID of the Creditcarddetail
	PkiCreditcarddetailID int32 `json:"pkiCreditcarddetailID"`
	// The unique ID of the Creditcardtype
	FkiCreditcardtypeID int32 `json:"fkiCreditcardtypeID"`
	// The numbermasked of the Creditcarddetail
	SCreditcarddetailNumbermasked string `json:"sCreditcarddetailNumbermasked"`
	// The expirationmonth of the Creditcarddetail
	ICreditcarddetailExpirationmonth int32 `json:"iCreditcarddetailExpirationmonth"`
	// The expirationyear of the Creditcarddetail
	ICreditcarddetailExpirationyear int32 `json:"iCreditcarddetailExpirationyear"`
	// The civic of the Creditcarddetail
	SCreditcarddetailCivic string `json:"sCreditcarddetailCivic"`
	// The street of the Creditcarddetail
	SCreditcarddetailStreet string `json:"sCreditcarddetailStreet"`
	// The zip of the Creditcarddetail
	SCreditcarddetailZip string `json:"sCreditcarddetailZip"`
}

type _CreditcarddetailResponseCompound CreditcarddetailResponseCompound

// NewCreditcarddetailResponseCompound instantiates a new CreditcarddetailResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcarddetailResponseCompound(pkiCreditcarddetailID int32, fkiCreditcardtypeID int32, sCreditcarddetailNumbermasked string, iCreditcarddetailExpirationmonth int32, iCreditcarddetailExpirationyear int32, sCreditcarddetailCivic string, sCreditcarddetailStreet string, sCreditcarddetailZip string) *CreditcarddetailResponseCompound {
	this := CreditcarddetailResponseCompound{}
	this.PkiCreditcarddetailID = pkiCreditcarddetailID
	this.FkiCreditcardtypeID = fkiCreditcardtypeID
	this.SCreditcarddetailNumbermasked = sCreditcarddetailNumbermasked
	this.ICreditcarddetailExpirationmonth = iCreditcarddetailExpirationmonth
	this.ICreditcarddetailExpirationyear = iCreditcarddetailExpirationyear
	this.SCreditcarddetailCivic = sCreditcarddetailCivic
	this.SCreditcarddetailStreet = sCreditcarddetailStreet
	this.SCreditcarddetailZip = sCreditcarddetailZip
	return &this
}

// NewCreditcarddetailResponseCompoundWithDefaults instantiates a new CreditcarddetailResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcarddetailResponseCompoundWithDefaults() *CreditcarddetailResponseCompound {
	this := CreditcarddetailResponseCompound{}
	return &this
}

// GetPkiCreditcarddetailID returns the PkiCreditcarddetailID field value
func (o *CreditcarddetailResponseCompound) GetPkiCreditcarddetailID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCreditcarddetailID
}

// GetPkiCreditcarddetailIDOk returns a tuple with the PkiCreditcarddetailID field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetPkiCreditcarddetailIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCreditcarddetailID, true
}

// SetPkiCreditcarddetailID sets field value
func (o *CreditcarddetailResponseCompound) SetPkiCreditcarddetailID(v int32) {
	o.PkiCreditcarddetailID = v
}

// GetFkiCreditcardtypeID returns the FkiCreditcardtypeID field value
func (o *CreditcarddetailResponseCompound) GetFkiCreditcardtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCreditcardtypeID
}

// GetFkiCreditcardtypeIDOk returns a tuple with the FkiCreditcardtypeID field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetFkiCreditcardtypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCreditcardtypeID, true
}

// SetFkiCreditcardtypeID sets field value
func (o *CreditcarddetailResponseCompound) SetFkiCreditcardtypeID(v int32) {
	o.FkiCreditcardtypeID = v
}

// GetSCreditcarddetailNumbermasked returns the SCreditcarddetailNumbermasked field value
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailNumbermasked() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcarddetailNumbermasked
}

// GetSCreditcarddetailNumbermaskedOk returns a tuple with the SCreditcarddetailNumbermasked field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailNumbermaskedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcarddetailNumbermasked, true
}

// SetSCreditcarddetailNumbermasked sets field value
func (o *CreditcarddetailResponseCompound) SetSCreditcarddetailNumbermasked(v string) {
	o.SCreditcarddetailNumbermasked = v
}

// GetICreditcarddetailExpirationmonth returns the ICreditcarddetailExpirationmonth field value
func (o *CreditcarddetailResponseCompound) GetICreditcarddetailExpirationmonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICreditcarddetailExpirationmonth
}

// GetICreditcarddetailExpirationmonthOk returns a tuple with the ICreditcarddetailExpirationmonth field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetICreditcarddetailExpirationmonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICreditcarddetailExpirationmonth, true
}

// SetICreditcarddetailExpirationmonth sets field value
func (o *CreditcarddetailResponseCompound) SetICreditcarddetailExpirationmonth(v int32) {
	o.ICreditcarddetailExpirationmonth = v
}

// GetICreditcarddetailExpirationyear returns the ICreditcarddetailExpirationyear field value
func (o *CreditcarddetailResponseCompound) GetICreditcarddetailExpirationyear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICreditcarddetailExpirationyear
}

// GetICreditcarddetailExpirationyearOk returns a tuple with the ICreditcarddetailExpirationyear field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetICreditcarddetailExpirationyearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICreditcarddetailExpirationyear, true
}

// SetICreditcarddetailExpirationyear sets field value
func (o *CreditcarddetailResponseCompound) SetICreditcarddetailExpirationyear(v int32) {
	o.ICreditcarddetailExpirationyear = v
}

// GetSCreditcarddetailCivic returns the SCreditcarddetailCivic field value
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailCivic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcarddetailCivic
}

// GetSCreditcarddetailCivicOk returns a tuple with the SCreditcarddetailCivic field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailCivicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcarddetailCivic, true
}

// SetSCreditcarddetailCivic sets field value
func (o *CreditcarddetailResponseCompound) SetSCreditcarddetailCivic(v string) {
	o.SCreditcarddetailCivic = v
}

// GetSCreditcarddetailStreet returns the SCreditcarddetailStreet field value
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcarddetailStreet
}

// GetSCreditcarddetailStreetOk returns a tuple with the SCreditcarddetailStreet field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcarddetailStreet, true
}

// SetSCreditcarddetailStreet sets field value
func (o *CreditcarddetailResponseCompound) SetSCreditcarddetailStreet(v string) {
	o.SCreditcarddetailStreet = v
}

// GetSCreditcarddetailZip returns the SCreditcarddetailZip field value
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcarddetailZip
}

// GetSCreditcarddetailZipOk returns a tuple with the SCreditcarddetailZip field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailResponseCompound) GetSCreditcarddetailZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcarddetailZip, true
}

// SetSCreditcarddetailZip sets field value
func (o *CreditcarddetailResponseCompound) SetSCreditcarddetailZip(v string) {
	o.SCreditcarddetailZip = v
}

func (o CreditcarddetailResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcarddetailResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiCreditcarddetailID"] = o.PkiCreditcarddetailID
	toSerialize["fkiCreditcardtypeID"] = o.FkiCreditcardtypeID
	toSerialize["sCreditcarddetailNumbermasked"] = o.SCreditcarddetailNumbermasked
	toSerialize["iCreditcarddetailExpirationmonth"] = o.ICreditcarddetailExpirationmonth
	toSerialize["iCreditcarddetailExpirationyear"] = o.ICreditcarddetailExpirationyear
	toSerialize["sCreditcarddetailCivic"] = o.SCreditcarddetailCivic
	toSerialize["sCreditcarddetailStreet"] = o.SCreditcarddetailStreet
	toSerialize["sCreditcarddetailZip"] = o.SCreditcarddetailZip
	return toSerialize, nil
}

func (o *CreditcarddetailResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiCreditcarddetailID",
		"fkiCreditcardtypeID",
		"sCreditcarddetailNumbermasked",
		"iCreditcarddetailExpirationmonth",
		"iCreditcarddetailExpirationyear",
		"sCreditcarddetailCivic",
		"sCreditcarddetailStreet",
		"sCreditcarddetailZip",
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

	varCreditcarddetailResponseCompound := _CreditcarddetailResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcarddetailResponseCompound)

	if err != nil {
		return err
	}

	*o = CreditcarddetailResponseCompound(varCreditcarddetailResponseCompound)

	return err
}

type NullableCreditcarddetailResponseCompound struct {
	value *CreditcarddetailResponseCompound
	isSet bool
}

func (v NullableCreditcarddetailResponseCompound) Get() *CreditcarddetailResponseCompound {
	return v.value
}

func (v *NullableCreditcarddetailResponseCompound) Set(val *CreditcarddetailResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcarddetailResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcarddetailResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcarddetailResponseCompound(val *CreditcarddetailResponseCompound) *NullableCreditcarddetailResponseCompound {
	return &NullableCreditcarddetailResponseCompound{value: val, isSet: true}
}

func (v NullableCreditcarddetailResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcarddetailResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


