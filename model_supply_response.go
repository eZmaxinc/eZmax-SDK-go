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

// checks if the SupplyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyResponse{}

// SupplyResponse A Supply Object
type SupplyResponse struct {
	// The unique ID of the Supply
	PkiSupplyID int32 `json:"pkiSupplyID"`
	// The unique ID of the Glaccount
	FkiGlaccountID *int32 `json:"fkiGlaccountID,omitempty"`
	// The unique ID of the Glaccountcontainer
	FkiGlaccountcontainerID *int32 `json:"fkiGlaccountcontainerID,omitempty"`
	// The unique ID of the Variableexpense
	FkiVariableexpenseID int32 `json:"fkiVariableexpenseID"`
	// The code of the Supply
	SSupplyCode string `json:"sSupplyCode" validate:"regexp=^.{0,5}$"`
	ObjSupplyDescription MultilingualSupplyDescription `json:"objSupplyDescription"`
	// The unit price of the Supply
	DSupplyUnitprice string `json:"dSupplyUnitprice" validate:"regexp=^-{0,1}[\\\\d]{1,9}?\\\\.[\\\\d]{2}$"`
	// Whether the supply is active or not
	BSupplyIsactive bool `json:"bSupplyIsactive"`
	// Whether if the price is variable
	BSupplyVariableprice bool `json:"bSupplyVariableprice"`
	// The Description for the Glaccount in the language of the requester
	SGlaccountDescriptionX *string `json:"sGlaccountDescriptionX,omitempty"`
	// The Description for the Glaccountcontainer in the language of the requester
	SGlaccountcontainerLongdescriptionX *string `json:"sGlaccountcontainerLongdescriptionX,omitempty"`
	// The description of the Variableexpense in the language of the requester
	SVariableexpenseDescriptionX *string `json:"sVariableexpenseDescriptionX,omitempty" validate:"regexp=^.{0,40}$"`
}

type _SupplyResponse SupplyResponse

// NewSupplyResponse instantiates a new SupplyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyResponse(pkiSupplyID int32, fkiVariableexpenseID int32, sSupplyCode string, objSupplyDescription MultilingualSupplyDescription, dSupplyUnitprice string, bSupplyIsactive bool, bSupplyVariableprice bool) *SupplyResponse {
	this := SupplyResponse{}
	this.PkiSupplyID = pkiSupplyID
	this.FkiVariableexpenseID = fkiVariableexpenseID
	this.SSupplyCode = sSupplyCode
	this.ObjSupplyDescription = objSupplyDescription
	this.DSupplyUnitprice = dSupplyUnitprice
	this.BSupplyIsactive = bSupplyIsactive
	this.BSupplyVariableprice = bSupplyVariableprice
	return &this
}

// NewSupplyResponseWithDefaults instantiates a new SupplyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyResponseWithDefaults() *SupplyResponse {
	this := SupplyResponse{}
	return &this
}

// GetPkiSupplyID returns the PkiSupplyID field value
func (o *SupplyResponse) GetPkiSupplyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSupplyID
}

// GetPkiSupplyIDOk returns a tuple with the PkiSupplyID field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetPkiSupplyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSupplyID, true
}

// SetPkiSupplyID sets field value
func (o *SupplyResponse) SetPkiSupplyID(v int32) {
	o.PkiSupplyID = v
}

// GetFkiGlaccountID returns the FkiGlaccountID field value if set, zero value otherwise.
func (o *SupplyResponse) GetFkiGlaccountID() int32 {
	if o == nil || IsNil(o.FkiGlaccountID) {
		var ret int32
		return ret
	}
	return *o.FkiGlaccountID
}

// GetFkiGlaccountIDOk returns a tuple with the FkiGlaccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetFkiGlaccountIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiGlaccountID) {
		return nil, false
	}
	return o.FkiGlaccountID, true
}

// HasFkiGlaccountID returns a boolean if a field has been set.
func (o *SupplyResponse) HasFkiGlaccountID() bool {
	if o != nil && !IsNil(o.FkiGlaccountID) {
		return true
	}

	return false
}

// SetFkiGlaccountID gets a reference to the given int32 and assigns it to the FkiGlaccountID field.
func (o *SupplyResponse) SetFkiGlaccountID(v int32) {
	o.FkiGlaccountID = &v
}

// GetFkiGlaccountcontainerID returns the FkiGlaccountcontainerID field value if set, zero value otherwise.
func (o *SupplyResponse) GetFkiGlaccountcontainerID() int32 {
	if o == nil || IsNil(o.FkiGlaccountcontainerID) {
		var ret int32
		return ret
	}
	return *o.FkiGlaccountcontainerID
}

// GetFkiGlaccountcontainerIDOk returns a tuple with the FkiGlaccountcontainerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetFkiGlaccountcontainerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiGlaccountcontainerID) {
		return nil, false
	}
	return o.FkiGlaccountcontainerID, true
}

// HasFkiGlaccountcontainerID returns a boolean if a field has been set.
func (o *SupplyResponse) HasFkiGlaccountcontainerID() bool {
	if o != nil && !IsNil(o.FkiGlaccountcontainerID) {
		return true
	}

	return false
}

// SetFkiGlaccountcontainerID gets a reference to the given int32 and assigns it to the FkiGlaccountcontainerID field.
func (o *SupplyResponse) SetFkiGlaccountcontainerID(v int32) {
	o.FkiGlaccountcontainerID = &v
}

// GetFkiVariableexpenseID returns the FkiVariableexpenseID field value
func (o *SupplyResponse) GetFkiVariableexpenseID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiVariableexpenseID
}

// GetFkiVariableexpenseIDOk returns a tuple with the FkiVariableexpenseID field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetFkiVariableexpenseIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiVariableexpenseID, true
}

// SetFkiVariableexpenseID sets field value
func (o *SupplyResponse) SetFkiVariableexpenseID(v int32) {
	o.FkiVariableexpenseID = v
}

// GetSSupplyCode returns the SSupplyCode field value
func (o *SupplyResponse) GetSSupplyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSupplyCode
}

// GetSSupplyCodeOk returns a tuple with the SSupplyCode field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetSSupplyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSupplyCode, true
}

// SetSSupplyCode sets field value
func (o *SupplyResponse) SetSSupplyCode(v string) {
	o.SSupplyCode = v
}

// GetObjSupplyDescription returns the ObjSupplyDescription field value
func (o *SupplyResponse) GetObjSupplyDescription() MultilingualSupplyDescription {
	if o == nil {
		var ret MultilingualSupplyDescription
		return ret
	}

	return o.ObjSupplyDescription
}

// GetObjSupplyDescriptionOk returns a tuple with the ObjSupplyDescription field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetObjSupplyDescriptionOk() (*MultilingualSupplyDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSupplyDescription, true
}

// SetObjSupplyDescription sets field value
func (o *SupplyResponse) SetObjSupplyDescription(v MultilingualSupplyDescription) {
	o.ObjSupplyDescription = v
}

// GetDSupplyUnitprice returns the DSupplyUnitprice field value
func (o *SupplyResponse) GetDSupplyUnitprice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DSupplyUnitprice
}

// GetDSupplyUnitpriceOk returns a tuple with the DSupplyUnitprice field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetDSupplyUnitpriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DSupplyUnitprice, true
}

// SetDSupplyUnitprice sets field value
func (o *SupplyResponse) SetDSupplyUnitprice(v string) {
	o.DSupplyUnitprice = v
}

// GetBSupplyIsactive returns the BSupplyIsactive field value
func (o *SupplyResponse) GetBSupplyIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSupplyIsactive
}

// GetBSupplyIsactiveOk returns a tuple with the BSupplyIsactive field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetBSupplyIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSupplyIsactive, true
}

// SetBSupplyIsactive sets field value
func (o *SupplyResponse) SetBSupplyIsactive(v bool) {
	o.BSupplyIsactive = v
}

// GetBSupplyVariableprice returns the BSupplyVariableprice field value
func (o *SupplyResponse) GetBSupplyVariableprice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSupplyVariableprice
}

// GetBSupplyVariablepriceOk returns a tuple with the BSupplyVariableprice field value
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetBSupplyVariablepriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSupplyVariableprice, true
}

// SetBSupplyVariableprice sets field value
func (o *SupplyResponse) SetBSupplyVariableprice(v bool) {
	o.BSupplyVariableprice = v
}

// GetSGlaccountDescriptionX returns the SGlaccountDescriptionX field value if set, zero value otherwise.
func (o *SupplyResponse) GetSGlaccountDescriptionX() string {
	if o == nil || IsNil(o.SGlaccountDescriptionX) {
		var ret string
		return ret
	}
	return *o.SGlaccountDescriptionX
}

// GetSGlaccountDescriptionXOk returns a tuple with the SGlaccountDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetSGlaccountDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SGlaccountDescriptionX) {
		return nil, false
	}
	return o.SGlaccountDescriptionX, true
}

// HasSGlaccountDescriptionX returns a boolean if a field has been set.
func (o *SupplyResponse) HasSGlaccountDescriptionX() bool {
	if o != nil && !IsNil(o.SGlaccountDescriptionX) {
		return true
	}

	return false
}

// SetSGlaccountDescriptionX gets a reference to the given string and assigns it to the SGlaccountDescriptionX field.
func (o *SupplyResponse) SetSGlaccountDescriptionX(v string) {
	o.SGlaccountDescriptionX = &v
}

// GetSGlaccountcontainerLongdescriptionX returns the SGlaccountcontainerLongdescriptionX field value if set, zero value otherwise.
func (o *SupplyResponse) GetSGlaccountcontainerLongdescriptionX() string {
	if o == nil || IsNil(o.SGlaccountcontainerLongdescriptionX) {
		var ret string
		return ret
	}
	return *o.SGlaccountcontainerLongdescriptionX
}

// GetSGlaccountcontainerLongdescriptionXOk returns a tuple with the SGlaccountcontainerLongdescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetSGlaccountcontainerLongdescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SGlaccountcontainerLongdescriptionX) {
		return nil, false
	}
	return o.SGlaccountcontainerLongdescriptionX, true
}

// HasSGlaccountcontainerLongdescriptionX returns a boolean if a field has been set.
func (o *SupplyResponse) HasSGlaccountcontainerLongdescriptionX() bool {
	if o != nil && !IsNil(o.SGlaccountcontainerLongdescriptionX) {
		return true
	}

	return false
}

// SetSGlaccountcontainerLongdescriptionX gets a reference to the given string and assigns it to the SGlaccountcontainerLongdescriptionX field.
func (o *SupplyResponse) SetSGlaccountcontainerLongdescriptionX(v string) {
	o.SGlaccountcontainerLongdescriptionX = &v
}

// GetSVariableexpenseDescriptionX returns the SVariableexpenseDescriptionX field value if set, zero value otherwise.
func (o *SupplyResponse) GetSVariableexpenseDescriptionX() string {
	if o == nil || IsNil(o.SVariableexpenseDescriptionX) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseDescriptionX
}

// GetSVariableexpenseDescriptionXOk returns a tuple with the SVariableexpenseDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplyResponse) GetSVariableexpenseDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseDescriptionX) {
		return nil, false
	}
	return o.SVariableexpenseDescriptionX, true
}

// HasSVariableexpenseDescriptionX returns a boolean if a field has been set.
func (o *SupplyResponse) HasSVariableexpenseDescriptionX() bool {
	if o != nil && !IsNil(o.SVariableexpenseDescriptionX) {
		return true
	}

	return false
}

// SetSVariableexpenseDescriptionX gets a reference to the given string and assigns it to the SVariableexpenseDescriptionX field.
func (o *SupplyResponse) SetSVariableexpenseDescriptionX(v string) {
	o.SVariableexpenseDescriptionX = &v
}

func (o SupplyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSupplyID"] = o.PkiSupplyID
	if !IsNil(o.FkiGlaccountID) {
		toSerialize["fkiGlaccountID"] = o.FkiGlaccountID
	}
	if !IsNil(o.FkiGlaccountcontainerID) {
		toSerialize["fkiGlaccountcontainerID"] = o.FkiGlaccountcontainerID
	}
	toSerialize["fkiVariableexpenseID"] = o.FkiVariableexpenseID
	toSerialize["sSupplyCode"] = o.SSupplyCode
	toSerialize["objSupplyDescription"] = o.ObjSupplyDescription
	toSerialize["dSupplyUnitprice"] = o.DSupplyUnitprice
	toSerialize["bSupplyIsactive"] = o.BSupplyIsactive
	toSerialize["bSupplyVariableprice"] = o.BSupplyVariableprice
	if !IsNil(o.SGlaccountDescriptionX) {
		toSerialize["sGlaccountDescriptionX"] = o.SGlaccountDescriptionX
	}
	if !IsNil(o.SGlaccountcontainerLongdescriptionX) {
		toSerialize["sGlaccountcontainerLongdescriptionX"] = o.SGlaccountcontainerLongdescriptionX
	}
	if !IsNil(o.SVariableexpenseDescriptionX) {
		toSerialize["sVariableexpenseDescriptionX"] = o.SVariableexpenseDescriptionX
	}
	return toSerialize, nil
}

func (o *SupplyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSupplyID",
		"fkiVariableexpenseID",
		"sSupplyCode",
		"objSupplyDescription",
		"dSupplyUnitprice",
		"bSupplyIsactive",
		"bSupplyVariableprice",
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

	varSupplyResponse := _SupplyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyResponse)

	if err != nil {
		return err
	}

	*o = SupplyResponse(varSupplyResponse)

	return err
}

type NullableSupplyResponse struct {
	value *SupplyResponse
	isSet bool
}

func (v NullableSupplyResponse) Get() *SupplyResponse {
	return v.value
}

func (v *NullableSupplyResponse) Set(val *SupplyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyResponse(val *SupplyResponse) *NullableSupplyResponse {
	return &NullableSupplyResponse{value: val, isSet: true}
}

func (v NullableSupplyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


