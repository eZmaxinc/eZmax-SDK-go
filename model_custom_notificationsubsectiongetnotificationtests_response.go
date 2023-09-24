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
)

// checks if the CustomNotificationsubsectiongetnotificationtestsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomNotificationsubsectiongetnotificationtestsResponse{}

// CustomNotificationsubsectiongetnotificationtestsResponse A Notificationsubsection Object in the context of getNotificationtests
type CustomNotificationsubsectiongetnotificationtestsResponse struct {
	// The unique ID of the Notificationsubsection
	PkiNotificationsubsectionID int32 `json:"pkiNotificationsubsectionID"`
	// The unique ID of the Notificationsection
	FkiNotificationsectionID int32 `json:"fkiNotificationsectionID"`
	ObjNotificationsubsectionName *MultilingualNotificationsubsectionName `json:"objNotificationsubsectionName,omitempty"`
	// The name of the Notificationsection in the language of the requester
	SNotificationsectionNameX *string `json:"sNotificationsectionNameX,omitempty"`
	// The name of the Notificationsubsection in the language of the requester
	SNotificationsubsectionNameX string `json:"sNotificationsubsectionNameX"`
	AObjNotificationtest []CustomNotificationtestgetnotificationtestsResponse `json:"a_objNotificationtest"`
}

// NewCustomNotificationsubsectiongetnotificationtestsResponse instantiates a new CustomNotificationsubsectiongetnotificationtestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomNotificationsubsectiongetnotificationtestsResponse(pkiNotificationsubsectionID int32, fkiNotificationsectionID int32, sNotificationsubsectionNameX string, aObjNotificationtest []CustomNotificationtestgetnotificationtestsResponse) *CustomNotificationsubsectiongetnotificationtestsResponse {
	this := CustomNotificationsubsectiongetnotificationtestsResponse{}
	this.PkiNotificationsubsectionID = pkiNotificationsubsectionID
	this.FkiNotificationsectionID = fkiNotificationsectionID
	this.SNotificationsubsectionNameX = sNotificationsubsectionNameX
	this.AObjNotificationtest = aObjNotificationtest
	return &this
}

// NewCustomNotificationsubsectiongetnotificationtestsResponseWithDefaults instantiates a new CustomNotificationsubsectiongetnotificationtestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomNotificationsubsectiongetnotificationtestsResponseWithDefaults() *CustomNotificationsubsectiongetnotificationtestsResponse {
	this := CustomNotificationsubsectiongetnotificationtestsResponse{}
	return &this
}

// GetPkiNotificationsubsectionID returns the PkiNotificationsubsectionID field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetPkiNotificationsubsectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiNotificationsubsectionID
}

// GetPkiNotificationsubsectionIDOk returns a tuple with the PkiNotificationsubsectionID field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetPkiNotificationsubsectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiNotificationsubsectionID, true
}

// SetPkiNotificationsubsectionID sets field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetPkiNotificationsubsectionID(v int32) {
	o.PkiNotificationsubsectionID = v
}

// GetFkiNotificationsectionID returns the FkiNotificationsectionID field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetFkiNotificationsectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiNotificationsectionID
}

// GetFkiNotificationsectionIDOk returns a tuple with the FkiNotificationsectionID field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetFkiNotificationsectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiNotificationsectionID, true
}

// SetFkiNotificationsectionID sets field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetFkiNotificationsectionID(v int32) {
	o.FkiNotificationsectionID = v
}

// GetObjNotificationsubsectionName returns the ObjNotificationsubsectionName field value if set, zero value otherwise.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetObjNotificationsubsectionName() MultilingualNotificationsubsectionName {
	if o == nil || IsNil(o.ObjNotificationsubsectionName) {
		var ret MultilingualNotificationsubsectionName
		return ret
	}
	return *o.ObjNotificationsubsectionName
}

// GetObjNotificationsubsectionNameOk returns a tuple with the ObjNotificationsubsectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetObjNotificationsubsectionNameOk() (*MultilingualNotificationsubsectionName, bool) {
	if o == nil || IsNil(o.ObjNotificationsubsectionName) {
		return nil, false
	}
	return o.ObjNotificationsubsectionName, true
}

// HasObjNotificationsubsectionName returns a boolean if a field has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) HasObjNotificationsubsectionName() bool {
	if o != nil && !IsNil(o.ObjNotificationsubsectionName) {
		return true
	}

	return false
}

// SetObjNotificationsubsectionName gets a reference to the given MultilingualNotificationsubsectionName and assigns it to the ObjNotificationsubsectionName field.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetObjNotificationsubsectionName(v MultilingualNotificationsubsectionName) {
	o.ObjNotificationsubsectionName = &v
}

// GetSNotificationsectionNameX returns the SNotificationsectionNameX field value if set, zero value otherwise.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsectionNameX() string {
	if o == nil || IsNil(o.SNotificationsectionNameX) {
		var ret string
		return ret
	}
	return *o.SNotificationsectionNameX
}

// GetSNotificationsectionNameXOk returns a tuple with the SNotificationsectionNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsectionNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SNotificationsectionNameX) {
		return nil, false
	}
	return o.SNotificationsectionNameX, true
}

// HasSNotificationsectionNameX returns a boolean if a field has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) HasSNotificationsectionNameX() bool {
	if o != nil && !IsNil(o.SNotificationsectionNameX) {
		return true
	}

	return false
}

// SetSNotificationsectionNameX gets a reference to the given string and assigns it to the SNotificationsectionNameX field.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetSNotificationsectionNameX(v string) {
	o.SNotificationsectionNameX = &v
}

// GetSNotificationsubsectionNameX returns the SNotificationsubsectionNameX field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsubsectionNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SNotificationsubsectionNameX
}

// GetSNotificationsubsectionNameXOk returns a tuple with the SNotificationsubsectionNameX field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsubsectionNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNotificationsubsectionNameX, true
}

// SetSNotificationsubsectionNameX sets field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetSNotificationsubsectionNameX(v string) {
	o.SNotificationsubsectionNameX = v
}

// GetAObjNotificationtest returns the AObjNotificationtest field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetAObjNotificationtest() []CustomNotificationtestgetnotificationtestsResponse {
	if o == nil {
		var ret []CustomNotificationtestgetnotificationtestsResponse
		return ret
	}

	return o.AObjNotificationtest
}

// GetAObjNotificationtestOk returns a tuple with the AObjNotificationtest field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetAObjNotificationtestOk() ([]CustomNotificationtestgetnotificationtestsResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjNotificationtest, true
}

// SetAObjNotificationtest sets field value
func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetAObjNotificationtest(v []CustomNotificationtestgetnotificationtestsResponse) {
	o.AObjNotificationtest = v
}

func (o CustomNotificationsubsectiongetnotificationtestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomNotificationsubsectiongetnotificationtestsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiNotificationsubsectionID"] = o.PkiNotificationsubsectionID
	toSerialize["fkiNotificationsectionID"] = o.FkiNotificationsectionID
	if !IsNil(o.ObjNotificationsubsectionName) {
		toSerialize["objNotificationsubsectionName"] = o.ObjNotificationsubsectionName
	}
	if !IsNil(o.SNotificationsectionNameX) {
		toSerialize["sNotificationsectionNameX"] = o.SNotificationsectionNameX
	}
	toSerialize["sNotificationsubsectionNameX"] = o.SNotificationsubsectionNameX
	toSerialize["a_objNotificationtest"] = o.AObjNotificationtest
	return toSerialize, nil
}

type NullableCustomNotificationsubsectiongetnotificationtestsResponse struct {
	value *CustomNotificationsubsectiongetnotificationtestsResponse
	isSet bool
}

func (v NullableCustomNotificationsubsectiongetnotificationtestsResponse) Get() *CustomNotificationsubsectiongetnotificationtestsResponse {
	return v.value
}

func (v *NullableCustomNotificationsubsectiongetnotificationtestsResponse) Set(val *CustomNotificationsubsectiongetnotificationtestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomNotificationsubsectiongetnotificationtestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomNotificationsubsectiongetnotificationtestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomNotificationsubsectiongetnotificationtestsResponse(val *CustomNotificationsubsectiongetnotificationtestsResponse) *NullableCustomNotificationsubsectiongetnotificationtestsResponse {
	return &NullableCustomNotificationsubsectiongetnotificationtestsResponse{value: val, isSet: true}
}

func (v NullableCustomNotificationsubsectiongetnotificationtestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomNotificationsubsectiongetnotificationtestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

