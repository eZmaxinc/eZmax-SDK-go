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

// checks if the BuyercontractGetCommunicationListV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyercontractGetCommunicationListV1Response{}

// BuyercontractGetCommunicationListV1Response Response for GET /1/object/buyercontract/{pkiBuyercontractID}/getCommunicationList
type BuyercontractGetCommunicationListV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayloadGetList `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload BuyercontractGetCommunicationListV1ResponseMPayload `json:"mPayload"`
}

type _BuyercontractGetCommunicationListV1Response BuyercontractGetCommunicationListV1Response

// NewBuyercontractGetCommunicationListV1Response instantiates a new BuyercontractGetCommunicationListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyercontractGetCommunicationListV1Response(objDebugPayload CommonResponseObjDebugPayloadGetList, mPayload BuyercontractGetCommunicationListV1ResponseMPayload) *BuyercontractGetCommunicationListV1Response {
	this := BuyercontractGetCommunicationListV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewBuyercontractGetCommunicationListV1ResponseWithDefaults instantiates a new BuyercontractGetCommunicationListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyercontractGetCommunicationListV1ResponseWithDefaults() *BuyercontractGetCommunicationListV1Response {
	this := BuyercontractGetCommunicationListV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *BuyercontractGetCommunicationListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *BuyercontractGetCommunicationListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *BuyercontractGetCommunicationListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *BuyercontractGetCommunicationListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyercontractGetCommunicationListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *BuyercontractGetCommunicationListV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *BuyercontractGetCommunicationListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *BuyercontractGetCommunicationListV1Response) GetMPayload() BuyercontractGetCommunicationListV1ResponseMPayload {
	if o == nil {
		var ret BuyercontractGetCommunicationListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *BuyercontractGetCommunicationListV1Response) GetMPayloadOk() (*BuyercontractGetCommunicationListV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *BuyercontractGetCommunicationListV1Response) SetMPayload(v BuyercontractGetCommunicationListV1ResponseMPayload) {
	o.MPayload = v
}

func (o BuyercontractGetCommunicationListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyercontractGetCommunicationListV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *BuyercontractGetCommunicationListV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objDebugPayload",
		"mPayload",
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

	varBuyercontractGetCommunicationListV1Response := _BuyercontractGetCommunicationListV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuyercontractGetCommunicationListV1Response)

	if err != nil {
		return err
	}

	*o = BuyercontractGetCommunicationListV1Response(varBuyercontractGetCommunicationListV1Response)

	return err
}

type NullableBuyercontractGetCommunicationListV1Response struct {
	value *BuyercontractGetCommunicationListV1Response
	isSet bool
}

func (v NullableBuyercontractGetCommunicationListV1Response) Get() *BuyercontractGetCommunicationListV1Response {
	return v.value
}

func (v *NullableBuyercontractGetCommunicationListV1Response) Set(val *BuyercontractGetCommunicationListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyercontractGetCommunicationListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyercontractGetCommunicationListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyercontractGetCommunicationListV1Response(val *BuyercontractGetCommunicationListV1Response) *NullableBuyercontractGetCommunicationListV1Response {
	return &NullableBuyercontractGetCommunicationListV1Response{value: val, isSet: true}
}

func (v NullableBuyercontractGetCommunicationListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyercontractGetCommunicationListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


