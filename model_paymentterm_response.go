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

// checks if the PaymenttermResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymenttermResponse{}

// PaymenttermResponse A Paymentterm Object
type PaymenttermResponse struct {
	// The unique ID of the Paymentterm
	PkiPaymenttermID int32 `json:"pkiPaymenttermID"`
	// The code of the Paymentterm
	SPaymenttermCode string `json:"sPaymenttermCode"`
	EPaymenttermType FieldEPaymenttermType `json:"ePaymenttermType"`
	// The day of the Paymentterm
	IPaymenttermDay int32 `json:"iPaymenttermDay"`
	ObjPaymenttermDescription MultilingualPaymenttermDescription `json:"objPaymenttermDescription"`
	// Whether the Paymentterm is active or not
	BPaymenttermIsactive bool `json:"bPaymenttermIsactive"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _PaymenttermResponse PaymenttermResponse

// NewPaymenttermResponse instantiates a new PaymenttermResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymenttermResponse(pkiPaymenttermID int32, sPaymenttermCode string, ePaymenttermType FieldEPaymenttermType, iPaymenttermDay int32, objPaymenttermDescription MultilingualPaymenttermDescription, bPaymenttermIsactive bool, objAudit CommonAudit) *PaymenttermResponse {
	this := PaymenttermResponse{}
	this.PkiPaymenttermID = pkiPaymenttermID
	this.SPaymenttermCode = sPaymenttermCode
	this.EPaymenttermType = ePaymenttermType
	this.IPaymenttermDay = iPaymenttermDay
	this.ObjPaymenttermDescription = objPaymenttermDescription
	this.BPaymenttermIsactive = bPaymenttermIsactive
	this.ObjAudit = objAudit
	return &this
}

// NewPaymenttermResponseWithDefaults instantiates a new PaymenttermResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymenttermResponseWithDefaults() *PaymenttermResponse {
	this := PaymenttermResponse{}
	return &this
}

// GetPkiPaymenttermID returns the PkiPaymenttermID field value
func (o *PaymenttermResponse) GetPkiPaymenttermID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiPaymenttermID
}

// GetPkiPaymenttermIDOk returns a tuple with the PkiPaymenttermID field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetPkiPaymenttermIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiPaymenttermID, true
}

// SetPkiPaymenttermID sets field value
func (o *PaymenttermResponse) SetPkiPaymenttermID(v int32) {
	o.PkiPaymenttermID = v
}

// GetSPaymenttermCode returns the SPaymenttermCode field value
func (o *PaymenttermResponse) GetSPaymenttermCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SPaymenttermCode
}

// GetSPaymenttermCodeOk returns a tuple with the SPaymenttermCode field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetSPaymenttermCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SPaymenttermCode, true
}

// SetSPaymenttermCode sets field value
func (o *PaymenttermResponse) SetSPaymenttermCode(v string) {
	o.SPaymenttermCode = v
}

// GetEPaymenttermType returns the EPaymenttermType field value
func (o *PaymenttermResponse) GetEPaymenttermType() FieldEPaymenttermType {
	if o == nil {
		var ret FieldEPaymenttermType
		return ret
	}

	return o.EPaymenttermType
}

// GetEPaymenttermTypeOk returns a tuple with the EPaymenttermType field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetEPaymenttermTypeOk() (*FieldEPaymenttermType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EPaymenttermType, true
}

// SetEPaymenttermType sets field value
func (o *PaymenttermResponse) SetEPaymenttermType(v FieldEPaymenttermType) {
	o.EPaymenttermType = v
}

// GetIPaymenttermDay returns the IPaymenttermDay field value
func (o *PaymenttermResponse) GetIPaymenttermDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IPaymenttermDay
}

// GetIPaymenttermDayOk returns a tuple with the IPaymenttermDay field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetIPaymenttermDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPaymenttermDay, true
}

// SetIPaymenttermDay sets field value
func (o *PaymenttermResponse) SetIPaymenttermDay(v int32) {
	o.IPaymenttermDay = v
}

// GetObjPaymenttermDescription returns the ObjPaymenttermDescription field value
func (o *PaymenttermResponse) GetObjPaymenttermDescription() MultilingualPaymenttermDescription {
	if o == nil {
		var ret MultilingualPaymenttermDescription
		return ret
	}

	return o.ObjPaymenttermDescription
}

// GetObjPaymenttermDescriptionOk returns a tuple with the ObjPaymenttermDescription field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetObjPaymenttermDescriptionOk() (*MultilingualPaymenttermDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjPaymenttermDescription, true
}

// SetObjPaymenttermDescription sets field value
func (o *PaymenttermResponse) SetObjPaymenttermDescription(v MultilingualPaymenttermDescription) {
	o.ObjPaymenttermDescription = v
}

// GetBPaymenttermIsactive returns the BPaymenttermIsactive field value
func (o *PaymenttermResponse) GetBPaymenttermIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BPaymenttermIsactive
}

// GetBPaymenttermIsactiveOk returns a tuple with the BPaymenttermIsactive field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetBPaymenttermIsactiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BPaymenttermIsactive, true
}

// SetBPaymenttermIsactive sets field value
func (o *PaymenttermResponse) SetBPaymenttermIsactive(v bool) {
	o.BPaymenttermIsactive = v
}

// GetObjAudit returns the ObjAudit field value
func (o *PaymenttermResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *PaymenttermResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *PaymenttermResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o PaymenttermResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymenttermResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiPaymenttermID"] = o.PkiPaymenttermID
	toSerialize["sPaymenttermCode"] = o.SPaymenttermCode
	toSerialize["ePaymenttermType"] = o.EPaymenttermType
	toSerialize["iPaymenttermDay"] = o.IPaymenttermDay
	toSerialize["objPaymenttermDescription"] = o.ObjPaymenttermDescription
	toSerialize["bPaymenttermIsactive"] = o.BPaymenttermIsactive
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *PaymenttermResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiPaymenttermID",
		"sPaymenttermCode",
		"ePaymenttermType",
		"iPaymenttermDay",
		"objPaymenttermDescription",
		"bPaymenttermIsactive",
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

	varPaymenttermResponse := _PaymenttermResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymenttermResponse)

	if err != nil {
		return err
	}

	*o = PaymenttermResponse(varPaymenttermResponse)

	return err
}

type NullablePaymenttermResponse struct {
	value *PaymenttermResponse
	isSet bool
}

func (v NullablePaymenttermResponse) Get() *PaymenttermResponse {
	return v.value
}

func (v *NullablePaymenttermResponse) Set(val *PaymenttermResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymenttermResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymenttermResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymenttermResponse(val *PaymenttermResponse) *NullablePaymenttermResponse {
	return &NullablePaymenttermResponse{value: val, isSet: true}
}

func (v NullablePaymenttermResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymenttermResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


