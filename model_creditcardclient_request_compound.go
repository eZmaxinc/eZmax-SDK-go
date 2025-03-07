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

// checks if the CreditcardclientRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientRequestCompound{}

// CreditcardclientRequestCompound A Creditcardclient Object and children
type CreditcardclientRequestCompound struct {
	// The unique ID of the Creditcardclient
	PkiCreditcardclientID *int32 `json:"pkiCreditcardclientID,omitempty"`
	// The creditcard token identifier
	FksCreditcardtokenID *string `json:"fksCreditcardtokenID,omitempty" validate:"regexp=^\\\\{?[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\\\\}?$"`
	// Whether if it's the creditcardclient is the default one
	BCreditcardclientrelationIsdefault bool `json:"bCreditcardclientrelationIsdefault"`
	// The description of the Creditcardclient
	SCreditcardclientDescription string `json:"sCreditcardclientDescription" validate:"regexp=^.{0,50}$"`
	// Whether if it's an allowedagencypayment
	BCreditcardclientAllowedcompanypayment bool `json:"bCreditcardclientAllowedcompanypayment"`
	// Whether if it's an allowedroyallepageprotection
	BCreditcardclientAllowedezsign bool `json:"bCreditcardclientAllowedezsign"`
	// Whether if it's an allowedtranquillit
	BCreditcardclientAllowedtranquillit bool `json:"bCreditcardclientAllowedtranquillit"`
	ObjCreditcarddetail CreditcarddetailRequest `json:"objCreditcarddetail"`
	// The creditcard card CVV
	SCreditcardclientCVV string `json:"sCreditcardclientCVV" validate:"regexp=^[0-9]{3,4}$"`
}

type _CreditcardclientRequestCompound CreditcardclientRequestCompound

// NewCreditcardclientRequestCompound instantiates a new CreditcardclientRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientRequestCompound(bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientAllowedcompanypayment bool, bCreditcardclientAllowedezsign bool, bCreditcardclientAllowedtranquillit bool, objCreditcarddetail CreditcarddetailRequest, sCreditcardclientCVV string) *CreditcardclientRequestCompound {
	this := CreditcardclientRequestCompound{}
	this.BCreditcardclientrelationIsdefault = bCreditcardclientrelationIsdefault
	this.SCreditcardclientDescription = sCreditcardclientDescription
	this.BCreditcardclientAllowedcompanypayment = bCreditcardclientAllowedcompanypayment
	this.BCreditcardclientAllowedezsign = bCreditcardclientAllowedezsign
	this.BCreditcardclientAllowedtranquillit = bCreditcardclientAllowedtranquillit
	this.ObjCreditcarddetail = objCreditcarddetail
	this.SCreditcardclientCVV = sCreditcardclientCVV
	return &this
}

// NewCreditcardclientRequestCompoundWithDefaults instantiates a new CreditcardclientRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientRequestCompoundWithDefaults() *CreditcardclientRequestCompound {
	this := CreditcardclientRequestCompound{}
	return &this
}

// GetPkiCreditcardclientID returns the PkiCreditcardclientID field value if set, zero value otherwise.
func (o *CreditcardclientRequestCompound) GetPkiCreditcardclientID() int32 {
	if o == nil || IsNil(o.PkiCreditcardclientID) {
		var ret int32
		return ret
	}
	return *o.PkiCreditcardclientID
}

// GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetPkiCreditcardclientIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiCreditcardclientID) {
		return nil, false
	}
	return o.PkiCreditcardclientID, true
}

// HasPkiCreditcardclientID returns a boolean if a field has been set.
func (o *CreditcardclientRequestCompound) HasPkiCreditcardclientID() bool {
	if o != nil && !IsNil(o.PkiCreditcardclientID) {
		return true
	}

	return false
}

// SetPkiCreditcardclientID gets a reference to the given int32 and assigns it to the PkiCreditcardclientID field.
func (o *CreditcardclientRequestCompound) SetPkiCreditcardclientID(v int32) {
	o.PkiCreditcardclientID = &v
}

// GetFksCreditcardtokenID returns the FksCreditcardtokenID field value if set, zero value otherwise.
func (o *CreditcardclientRequestCompound) GetFksCreditcardtokenID() string {
	if o == nil || IsNil(o.FksCreditcardtokenID) {
		var ret string
		return ret
	}
	return *o.FksCreditcardtokenID
}

// GetFksCreditcardtokenIDOk returns a tuple with the FksCreditcardtokenID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetFksCreditcardtokenIDOk() (*string, bool) {
	if o == nil || IsNil(o.FksCreditcardtokenID) {
		return nil, false
	}
	return o.FksCreditcardtokenID, true
}

// HasFksCreditcardtokenID returns a boolean if a field has been set.
func (o *CreditcardclientRequestCompound) HasFksCreditcardtokenID() bool {
	if o != nil && !IsNil(o.FksCreditcardtokenID) {
		return true
	}

	return false
}

// SetFksCreditcardtokenID gets a reference to the given string and assigns it to the FksCreditcardtokenID field.
func (o *CreditcardclientRequestCompound) SetFksCreditcardtokenID(v string) {
	o.FksCreditcardtokenID = &v
}

// GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field value
func (o *CreditcardclientRequestCompound) GetBCreditcardclientrelationIsdefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientrelationIsdefault
}

// GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientrelationIsdefault, true
}

// SetBCreditcardclientrelationIsdefault sets field value
func (o *CreditcardclientRequestCompound) SetBCreditcardclientrelationIsdefault(v bool) {
	o.BCreditcardclientrelationIsdefault = v
}

// GetSCreditcardclientDescription returns the SCreditcardclientDescription field value
func (o *CreditcardclientRequestCompound) GetSCreditcardclientDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardclientDescription
}

// GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetSCreditcardclientDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardclientDescription, true
}

// SetSCreditcardclientDescription sets field value
func (o *CreditcardclientRequestCompound) SetSCreditcardclientDescription(v string) {
	o.SCreditcardclientDescription = v
}

// GetBCreditcardclientAllowedcompanypayment returns the BCreditcardclientAllowedcompanypayment field value
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedcompanypayment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedcompanypayment
}

// GetBCreditcardclientAllowedcompanypaymentOk returns a tuple with the BCreditcardclientAllowedcompanypayment field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedcompanypaymentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedcompanypayment, true
}

// SetBCreditcardclientAllowedcompanypayment sets field value
func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedcompanypayment(v bool) {
	o.BCreditcardclientAllowedcompanypayment = v
}

// GetBCreditcardclientAllowedezsign returns the BCreditcardclientAllowedezsign field value
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedezsign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedezsign
}

// GetBCreditcardclientAllowedezsignOk returns a tuple with the BCreditcardclientAllowedezsign field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedezsignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedezsign, true
}

// SetBCreditcardclientAllowedezsign sets field value
func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedezsign(v bool) {
	o.BCreditcardclientAllowedezsign = v
}

// GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field value
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedtranquillit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedtranquillit
}

// GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedtranquillit, true
}

// SetBCreditcardclientAllowedtranquillit sets field value
func (o *CreditcardclientRequestCompound) SetBCreditcardclientAllowedtranquillit(v bool) {
	o.BCreditcardclientAllowedtranquillit = v
}

// GetObjCreditcarddetail returns the ObjCreditcarddetail field value
func (o *CreditcardclientRequestCompound) GetObjCreditcarddetail() CreditcarddetailRequest {
	if o == nil {
		var ret CreditcarddetailRequest
		return ret
	}

	return o.ObjCreditcarddetail
}

// GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetObjCreditcarddetailOk() (*CreditcarddetailRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjCreditcarddetail, true
}

// SetObjCreditcarddetail sets field value
func (o *CreditcardclientRequestCompound) SetObjCreditcarddetail(v CreditcarddetailRequest) {
	o.ObjCreditcarddetail = v
}

// GetSCreditcardclientCVV returns the SCreditcardclientCVV field value
func (o *CreditcardclientRequestCompound) GetSCreditcardclientCVV() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardclientCVV
}

// GetSCreditcardclientCVVOk returns a tuple with the SCreditcardclientCVV field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequestCompound) GetSCreditcardclientCVVOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardclientCVV, true
}

// SetSCreditcardclientCVV sets field value
func (o *CreditcardclientRequestCompound) SetSCreditcardclientCVV(v string) {
	o.SCreditcardclientCVV = v
}

func (o CreditcardclientRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiCreditcardclientID) {
		toSerialize["pkiCreditcardclientID"] = o.PkiCreditcardclientID
	}
	if !IsNil(o.FksCreditcardtokenID) {
		toSerialize["fksCreditcardtokenID"] = o.FksCreditcardtokenID
	}
	toSerialize["bCreditcardclientrelationIsdefault"] = o.BCreditcardclientrelationIsdefault
	toSerialize["sCreditcardclientDescription"] = o.SCreditcardclientDescription
	toSerialize["bCreditcardclientAllowedcompanypayment"] = o.BCreditcardclientAllowedcompanypayment
	toSerialize["bCreditcardclientAllowedezsign"] = o.BCreditcardclientAllowedezsign
	toSerialize["bCreditcardclientAllowedtranquillit"] = o.BCreditcardclientAllowedtranquillit
	toSerialize["objCreditcarddetail"] = o.ObjCreditcarddetail
	toSerialize["sCreditcardclientCVV"] = o.SCreditcardclientCVV
	return toSerialize, nil
}

func (o *CreditcardclientRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bCreditcardclientrelationIsdefault",
		"sCreditcardclientDescription",
		"bCreditcardclientAllowedcompanypayment",
		"bCreditcardclientAllowedezsign",
		"bCreditcardclientAllowedtranquillit",
		"objCreditcarddetail",
		"sCreditcardclientCVV",
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

	varCreditcardclientRequestCompound := _CreditcardclientRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientRequestCompound)

	if err != nil {
		return err
	}

	*o = CreditcardclientRequestCompound(varCreditcardclientRequestCompound)

	return err
}

type NullableCreditcardclientRequestCompound struct {
	value *CreditcardclientRequestCompound
	isSet bool
}

func (v NullableCreditcardclientRequestCompound) Get() *CreditcardclientRequestCompound {
	return v.value
}

func (v *NullableCreditcardclientRequestCompound) Set(val *CreditcardclientRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientRequestCompound(val *CreditcardclientRequestCompound) *NullableCreditcardclientRequestCompound {
	return &NullableCreditcardclientRequestCompound{value: val, isSet: true}
}

func (v NullableCreditcardclientRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


