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

// checks if the NotificationtestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationtestResponse{}

// NotificationtestResponse A Notificationtest Object
type NotificationtestResponse struct {
	// The unique ID of the Notificationtest
	PkiNotificationtestID int32 `json:"pkiNotificationtestID"`
	ObjNotificationtestName MultilingualNotificationtestName `json:"objNotificationtestName"`
	// The unique ID of the Notificationsubsection
	FkiNotificationsubsectionID int32 `json:"fkiNotificationsubsectionID"`
	// The function name of the Notificationtest
	SNotificationtestFunction string `json:"sNotificationtestFunction"`
	// The name of the Notificationtest in the language of the requester
	SNotificationtestNameX string `json:"sNotificationtestNameX"`
}

type _NotificationtestResponse NotificationtestResponse

// NewNotificationtestResponse instantiates a new NotificationtestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationtestResponse(pkiNotificationtestID int32, objNotificationtestName MultilingualNotificationtestName, fkiNotificationsubsectionID int32, sNotificationtestFunction string, sNotificationtestNameX string) *NotificationtestResponse {
	this := NotificationtestResponse{}
	this.PkiNotificationtestID = pkiNotificationtestID
	this.ObjNotificationtestName = objNotificationtestName
	this.FkiNotificationsubsectionID = fkiNotificationsubsectionID
	this.SNotificationtestFunction = sNotificationtestFunction
	this.SNotificationtestNameX = sNotificationtestNameX
	return &this
}

// NewNotificationtestResponseWithDefaults instantiates a new NotificationtestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationtestResponseWithDefaults() *NotificationtestResponse {
	this := NotificationtestResponse{}
	return &this
}

// GetPkiNotificationtestID returns the PkiNotificationtestID field value
func (o *NotificationtestResponse) GetPkiNotificationtestID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiNotificationtestID
}

// GetPkiNotificationtestIDOk returns a tuple with the PkiNotificationtestID field value
// and a boolean to check if the value has been set.
func (o *NotificationtestResponse) GetPkiNotificationtestIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiNotificationtestID, true
}

// SetPkiNotificationtestID sets field value
func (o *NotificationtestResponse) SetPkiNotificationtestID(v int32) {
	o.PkiNotificationtestID = v
}

// GetObjNotificationtestName returns the ObjNotificationtestName field value
func (o *NotificationtestResponse) GetObjNotificationtestName() MultilingualNotificationtestName {
	if o == nil {
		var ret MultilingualNotificationtestName
		return ret
	}

	return o.ObjNotificationtestName
}

// GetObjNotificationtestNameOk returns a tuple with the ObjNotificationtestName field value
// and a boolean to check if the value has been set.
func (o *NotificationtestResponse) GetObjNotificationtestNameOk() (*MultilingualNotificationtestName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjNotificationtestName, true
}

// SetObjNotificationtestName sets field value
func (o *NotificationtestResponse) SetObjNotificationtestName(v MultilingualNotificationtestName) {
	o.ObjNotificationtestName = v
}

// GetFkiNotificationsubsectionID returns the FkiNotificationsubsectionID field value
func (o *NotificationtestResponse) GetFkiNotificationsubsectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiNotificationsubsectionID
}

// GetFkiNotificationsubsectionIDOk returns a tuple with the FkiNotificationsubsectionID field value
// and a boolean to check if the value has been set.
func (o *NotificationtestResponse) GetFkiNotificationsubsectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiNotificationsubsectionID, true
}

// SetFkiNotificationsubsectionID sets field value
func (o *NotificationtestResponse) SetFkiNotificationsubsectionID(v int32) {
	o.FkiNotificationsubsectionID = v
}

// GetSNotificationtestFunction returns the SNotificationtestFunction field value
func (o *NotificationtestResponse) GetSNotificationtestFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SNotificationtestFunction
}

// GetSNotificationtestFunctionOk returns a tuple with the SNotificationtestFunction field value
// and a boolean to check if the value has been set.
func (o *NotificationtestResponse) GetSNotificationtestFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNotificationtestFunction, true
}

// SetSNotificationtestFunction sets field value
func (o *NotificationtestResponse) SetSNotificationtestFunction(v string) {
	o.SNotificationtestFunction = v
}

// GetSNotificationtestNameX returns the SNotificationtestNameX field value
func (o *NotificationtestResponse) GetSNotificationtestNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SNotificationtestNameX
}

// GetSNotificationtestNameXOk returns a tuple with the SNotificationtestNameX field value
// and a boolean to check if the value has been set.
func (o *NotificationtestResponse) GetSNotificationtestNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNotificationtestNameX, true
}

// SetSNotificationtestNameX sets field value
func (o *NotificationtestResponse) SetSNotificationtestNameX(v string) {
	o.SNotificationtestNameX = v
}

func (o NotificationtestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationtestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiNotificationtestID"] = o.PkiNotificationtestID
	toSerialize["objNotificationtestName"] = o.ObjNotificationtestName
	toSerialize["fkiNotificationsubsectionID"] = o.FkiNotificationsubsectionID
	toSerialize["sNotificationtestFunction"] = o.SNotificationtestFunction
	toSerialize["sNotificationtestNameX"] = o.SNotificationtestNameX
	return toSerialize, nil
}

func (o *NotificationtestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiNotificationtestID",
		"objNotificationtestName",
		"fkiNotificationsubsectionID",
		"sNotificationtestFunction",
		"sNotificationtestNameX",
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

	varNotificationtestResponse := _NotificationtestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationtestResponse)

	if err != nil {
		return err
	}

	*o = NotificationtestResponse(varNotificationtestResponse)

	return err
}

type NullableNotificationtestResponse struct {
	value *NotificationtestResponse
	isSet bool
}

func (v NullableNotificationtestResponse) Get() *NotificationtestResponse {
	return v.value
}

func (v *NullableNotificationtestResponse) Set(val *NotificationtestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationtestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationtestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationtestResponse(val *NotificationtestResponse) *NullableNotificationtestResponse {
	return &NullableNotificationtestResponse{value: val, isSet: true}
}

func (v NullableNotificationtestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationtestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


