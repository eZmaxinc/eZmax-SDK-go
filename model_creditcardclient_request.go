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

// checks if the CreditcardclientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardclientRequest{}

// CreditcardclientRequest A Creditcardclient Object
type CreditcardclientRequest struct {
	// The unique ID of the Creditcardclient
	PkiCreditcardclientID *int32 `json:"pkiCreditcardclientID,omitempty"`
	// The creditcard token identifier
	FksCreditcardtokenID *string `json:"fksCreditcardtokenID,omitempty"`
	// Whether if it's an relationisdefault
	BCreditcardclientrelationIsdefault bool `json:"bCreditcardclientrelationIsdefault"`
	// The description of the Creditcardclient
	SCreditcardclientDescription string `json:"sCreditcardclientDescription"`
	// Whether the creditcardclient is active or not
	BCreditcardclientIsactive bool `json:"bCreditcardclientIsactive"`
	// Whether if it's an allowedagencypayment
	BCreditcardclientAllowedagencypayment bool `json:"bCreditcardclientAllowedagencypayment"`
	// Whether if it's an allowedroyallepageprotection
	BCreditcardclientAllowedroyallepageprotection bool `json:"bCreditcardclientAllowedroyallepageprotection"`
	// Whether if it's an allowedtranquillit
	BCreditcardclientAllowedtranquillit bool `json:"bCreditcardclientAllowedtranquillit"`
	ObjCreditcarddetail CreditcarddetailRequest `json:"objCreditcarddetail"`
	// The creditcard card CVV
	SCreditcardclientCVV string `json:"sCreditcardclientCVV"`
}

type _CreditcardclientRequest CreditcardclientRequest

// NewCreditcardclientRequest instantiates a new CreditcardclientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardclientRequest(bCreditcardclientrelationIsdefault bool, sCreditcardclientDescription string, bCreditcardclientIsactive bool, bCreditcardclientAllowedagencypayment bool, bCreditcardclientAllowedroyallepageprotection bool, bCreditcardclientAllowedtranquillit bool, objCreditcarddetail CreditcarddetailRequest, sCreditcardclientCVV string) *CreditcardclientRequest {
	this := CreditcardclientRequest{}
	this.BCreditcardclientrelationIsdefault = bCreditcardclientrelationIsdefault
	this.SCreditcardclientDescription = sCreditcardclientDescription
	this.BCreditcardclientIsactive = bCreditcardclientIsactive
	this.BCreditcardclientAllowedagencypayment = bCreditcardclientAllowedagencypayment
	this.BCreditcardclientAllowedroyallepageprotection = bCreditcardclientAllowedroyallepageprotection
	this.BCreditcardclientAllowedtranquillit = bCreditcardclientAllowedtranquillit
	this.ObjCreditcarddetail = objCreditcarddetail
	this.SCreditcardclientCVV = sCreditcardclientCVV
	return &this
}

// NewCreditcardclientRequestWithDefaults instantiates a new CreditcardclientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardclientRequestWithDefaults() *CreditcardclientRequest {
	this := CreditcardclientRequest{}
	return &this
}

// GetPkiCreditcardclientID returns the PkiCreditcardclientID field value if set, zero value otherwise.
func (o *CreditcardclientRequest) GetPkiCreditcardclientID() int32 {
	if o == nil || IsNil(o.PkiCreditcardclientID) {
		var ret int32
		return ret
	}
	return *o.PkiCreditcardclientID
}

// GetPkiCreditcardclientIDOk returns a tuple with the PkiCreditcardclientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetPkiCreditcardclientIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiCreditcardclientID) {
		return nil, false
	}
	return o.PkiCreditcardclientID, true
}

// HasPkiCreditcardclientID returns a boolean if a field has been set.
func (o *CreditcardclientRequest) HasPkiCreditcardclientID() bool {
	if o != nil && !IsNil(o.PkiCreditcardclientID) {
		return true
	}

	return false
}

// SetPkiCreditcardclientID gets a reference to the given int32 and assigns it to the PkiCreditcardclientID field.
func (o *CreditcardclientRequest) SetPkiCreditcardclientID(v int32) {
	o.PkiCreditcardclientID = &v
}

// GetFksCreditcardtokenID returns the FksCreditcardtokenID field value if set, zero value otherwise.
func (o *CreditcardclientRequest) GetFksCreditcardtokenID() string {
	if o == nil || IsNil(o.FksCreditcardtokenID) {
		var ret string
		return ret
	}
	return *o.FksCreditcardtokenID
}

// GetFksCreditcardtokenIDOk returns a tuple with the FksCreditcardtokenID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetFksCreditcardtokenIDOk() (*string, bool) {
	if o == nil || IsNil(o.FksCreditcardtokenID) {
		return nil, false
	}
	return o.FksCreditcardtokenID, true
}

// HasFksCreditcardtokenID returns a boolean if a field has been set.
func (o *CreditcardclientRequest) HasFksCreditcardtokenID() bool {
	if o != nil && !IsNil(o.FksCreditcardtokenID) {
		return true
	}

	return false
}

// SetFksCreditcardtokenID gets a reference to the given string and assigns it to the FksCreditcardtokenID field.
func (o *CreditcardclientRequest) SetFksCreditcardtokenID(v string) {
	o.FksCreditcardtokenID = &v
}

// GetBCreditcardclientrelationIsdefault returns the BCreditcardclientrelationIsdefault field value
func (o *CreditcardclientRequest) GetBCreditcardclientrelationIsdefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientrelationIsdefault
}

// GetBCreditcardclientrelationIsdefaultOk returns a tuple with the BCreditcardclientrelationIsdefault field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetBCreditcardclientrelationIsdefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientrelationIsdefault, true
}

// SetBCreditcardclientrelationIsdefault sets field value
func (o *CreditcardclientRequest) SetBCreditcardclientrelationIsdefault(v bool) {
	o.BCreditcardclientrelationIsdefault = v
}

// GetSCreditcardclientDescription returns the SCreditcardclientDescription field value
func (o *CreditcardclientRequest) GetSCreditcardclientDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardclientDescription
}

// GetSCreditcardclientDescriptionOk returns a tuple with the SCreditcardclientDescription field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetSCreditcardclientDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardclientDescription, true
}

// SetSCreditcardclientDescription sets field value
func (o *CreditcardclientRequest) SetSCreditcardclientDescription(v string) {
	o.SCreditcardclientDescription = v
}

// GetBCreditcardclientIsactive returns the BCreditcardclientIsactive field value
func (o *CreditcardclientRequest) GetBCreditcardclientIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientIsactive
}

// GetBCreditcardclientIsactiveOk returns a tuple with the BCreditcardclientIsactive field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetBCreditcardclientIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientIsactive, true
}

// SetBCreditcardclientIsactive sets field value
func (o *CreditcardclientRequest) SetBCreditcardclientIsactive(v bool) {
	o.BCreditcardclientIsactive = v
}

// GetBCreditcardclientAllowedagencypayment returns the BCreditcardclientAllowedagencypayment field value
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedagencypayment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedagencypayment
}

// GetBCreditcardclientAllowedagencypaymentOk returns a tuple with the BCreditcardclientAllowedagencypayment field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedagencypaymentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedagencypayment, true
}

// SetBCreditcardclientAllowedagencypayment sets field value
func (o *CreditcardclientRequest) SetBCreditcardclientAllowedagencypayment(v bool) {
	o.BCreditcardclientAllowedagencypayment = v
}

// GetBCreditcardclientAllowedroyallepageprotection returns the BCreditcardclientAllowedroyallepageprotection field value
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedroyallepageprotection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedroyallepageprotection
}

// GetBCreditcardclientAllowedroyallepageprotectionOk returns a tuple with the BCreditcardclientAllowedroyallepageprotection field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedroyallepageprotectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedroyallepageprotection, true
}

// SetBCreditcardclientAllowedroyallepageprotection sets field value
func (o *CreditcardclientRequest) SetBCreditcardclientAllowedroyallepageprotection(v bool) {
	o.BCreditcardclientAllowedroyallepageprotection = v
}

// GetBCreditcardclientAllowedtranquillit returns the BCreditcardclientAllowedtranquillit field value
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedtranquillit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BCreditcardclientAllowedtranquillit
}

// GetBCreditcardclientAllowedtranquillitOk returns a tuple with the BCreditcardclientAllowedtranquillit field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetBCreditcardclientAllowedtranquillitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BCreditcardclientAllowedtranquillit, true
}

// SetBCreditcardclientAllowedtranquillit sets field value
func (o *CreditcardclientRequest) SetBCreditcardclientAllowedtranquillit(v bool) {
	o.BCreditcardclientAllowedtranquillit = v
}

// GetObjCreditcarddetail returns the ObjCreditcarddetail field value
func (o *CreditcardclientRequest) GetObjCreditcarddetail() CreditcarddetailRequest {
	if o == nil {
		var ret CreditcarddetailRequest
		return ret
	}

	return o.ObjCreditcarddetail
}

// GetObjCreditcarddetailOk returns a tuple with the ObjCreditcarddetail field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetObjCreditcarddetailOk() (*CreditcarddetailRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjCreditcarddetail, true
}

// SetObjCreditcarddetail sets field value
func (o *CreditcardclientRequest) SetObjCreditcarddetail(v CreditcarddetailRequest) {
	o.ObjCreditcarddetail = v
}

// GetSCreditcardclientCVV returns the SCreditcardclientCVV field value
func (o *CreditcardclientRequest) GetSCreditcardclientCVV() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardclientCVV
}

// GetSCreditcardclientCVVOk returns a tuple with the SCreditcardclientCVV field value
// and a boolean to check if the value has been set.
func (o *CreditcardclientRequest) GetSCreditcardclientCVVOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardclientCVV, true
}

// SetSCreditcardclientCVV sets field value
func (o *CreditcardclientRequest) SetSCreditcardclientCVV(v string) {
	o.SCreditcardclientCVV = v
}

func (o CreditcardclientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardclientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiCreditcardclientID) {
		toSerialize["pkiCreditcardclientID"] = o.PkiCreditcardclientID
	}
	if !IsNil(o.FksCreditcardtokenID) {
		toSerialize["fksCreditcardtokenID"] = o.FksCreditcardtokenID
	}
	toSerialize["bCreditcardclientrelationIsdefault"] = o.BCreditcardclientrelationIsdefault
	toSerialize["sCreditcardclientDescription"] = o.SCreditcardclientDescription
	toSerialize["bCreditcardclientIsactive"] = o.BCreditcardclientIsactive
	toSerialize["bCreditcardclientAllowedagencypayment"] = o.BCreditcardclientAllowedagencypayment
	toSerialize["bCreditcardclientAllowedroyallepageprotection"] = o.BCreditcardclientAllowedroyallepageprotection
	toSerialize["bCreditcardclientAllowedtranquillit"] = o.BCreditcardclientAllowedtranquillit
	toSerialize["objCreditcarddetail"] = o.ObjCreditcarddetail
	toSerialize["sCreditcardclientCVV"] = o.SCreditcardclientCVV
	return toSerialize, nil
}

func (o *CreditcardclientRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bCreditcardclientrelationIsdefault",
		"sCreditcardclientDescription",
		"bCreditcardclientIsactive",
		"bCreditcardclientAllowedagencypayment",
		"bCreditcardclientAllowedroyallepageprotection",
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

	varCreditcardclientRequest := _CreditcardclientRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardclientRequest)

	if err != nil {
		return err
	}

	*o = CreditcardclientRequest(varCreditcardclientRequest)

	return err
}

type NullableCreditcardclientRequest struct {
	value *CreditcardclientRequest
	isSet bool
}

func (v NullableCreditcardclientRequest) Get() *CreditcardclientRequest {
	return v.value
}

func (v *NullableCreditcardclientRequest) Set(val *CreditcardclientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardclientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardclientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardclientRequest(val *CreditcardclientRequest) *NullableCreditcardclientRequest {
	return &NullableCreditcardclientRequest{value: val, isSet: true}
}

func (v NullableCreditcardclientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardclientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

