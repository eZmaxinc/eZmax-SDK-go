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

// checks if the CustomNotificationtestgetnotificationtestsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomNotificationtestgetnotificationtestsResponse{}

// CustomNotificationtestgetnotificationtestsResponse A Notificationtest Object in the context of getNotificationtests
type CustomNotificationtestgetnotificationtestsResponse struct {
	// The unique ID of the Notificationtest
	PkiNotificationtestID int32 `json:"pkiNotificationtestID"`
	ObjNotificationtestName MultilingualNotificationtestName `json:"objNotificationtestName"`
	// The unique ID of the Notificationsubsection
	FkiNotificationsubsectionID int32 `json:"fkiNotificationsubsectionID"`
	// The function name of the Notificationtest
	SNotificationtestFunction string `json:"sNotificationtestFunction"`
	// The name of the Notificationtest in the language of the requester
	SNotificationtestNameX string `json:"sNotificationtestNameX"`
	ENotificationpreferenceStatus FieldENotificationpreferenceStatus `json:"eNotificationpreferenceStatus"`
	// The number of elements returned by the Notificationtest
	INotificationtest int32 `json:"iNotificationtest"`
}

type _CustomNotificationtestgetnotificationtestsResponse CustomNotificationtestgetnotificationtestsResponse

// NewCustomNotificationtestgetnotificationtestsResponse instantiates a new CustomNotificationtestgetnotificationtestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomNotificationtestgetnotificationtestsResponse(pkiNotificationtestID int32, objNotificationtestName MultilingualNotificationtestName, fkiNotificationsubsectionID int32, sNotificationtestFunction string, sNotificationtestNameX string, eNotificationpreferenceStatus FieldENotificationpreferenceStatus, iNotificationtest int32) *CustomNotificationtestgetnotificationtestsResponse {
	this := CustomNotificationtestgetnotificationtestsResponse{}
	this.PkiNotificationtestID = pkiNotificationtestID
	this.ObjNotificationtestName = objNotificationtestName
	this.FkiNotificationsubsectionID = fkiNotificationsubsectionID
	this.SNotificationtestFunction = sNotificationtestFunction
	this.SNotificationtestNameX = sNotificationtestNameX
	this.ENotificationpreferenceStatus = eNotificationpreferenceStatus
	this.INotificationtest = iNotificationtest
	return &this
}

// NewCustomNotificationtestgetnotificationtestsResponseWithDefaults instantiates a new CustomNotificationtestgetnotificationtestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomNotificationtestgetnotificationtestsResponseWithDefaults() *CustomNotificationtestgetnotificationtestsResponse {
	this := CustomNotificationtestgetnotificationtestsResponse{}
	return &this
}

// GetPkiNotificationtestID returns the PkiNotificationtestID field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetPkiNotificationtestID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiNotificationtestID
}

// GetPkiNotificationtestIDOk returns a tuple with the PkiNotificationtestID field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetPkiNotificationtestIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiNotificationtestID, true
}

// SetPkiNotificationtestID sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetPkiNotificationtestID(v int32) {
	o.PkiNotificationtestID = v
}

// GetObjNotificationtestName returns the ObjNotificationtestName field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetObjNotificationtestName() MultilingualNotificationtestName {
	if o == nil {
		var ret MultilingualNotificationtestName
		return ret
	}

	return o.ObjNotificationtestName
}

// GetObjNotificationtestNameOk returns a tuple with the ObjNotificationtestName field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetObjNotificationtestNameOk() (*MultilingualNotificationtestName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjNotificationtestName, true
}

// SetObjNotificationtestName sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetObjNotificationtestName(v MultilingualNotificationtestName) {
	o.ObjNotificationtestName = v
}

// GetFkiNotificationsubsectionID returns the FkiNotificationsubsectionID field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetFkiNotificationsubsectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiNotificationsubsectionID
}

// GetFkiNotificationsubsectionIDOk returns a tuple with the FkiNotificationsubsectionID field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetFkiNotificationsubsectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiNotificationsubsectionID, true
}

// SetFkiNotificationsubsectionID sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetFkiNotificationsubsectionID(v int32) {
	o.FkiNotificationsubsectionID = v
}

// GetSNotificationtestFunction returns the SNotificationtestFunction field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SNotificationtestFunction
}

// GetSNotificationtestFunctionOk returns a tuple with the SNotificationtestFunction field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestFunctionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNotificationtestFunction, true
}

// SetSNotificationtestFunction sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetSNotificationtestFunction(v string) {
	o.SNotificationtestFunction = v
}

// GetSNotificationtestNameX returns the SNotificationtestNameX field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SNotificationtestNameX
}

// GetSNotificationtestNameXOk returns a tuple with the SNotificationtestNameX field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNotificationtestNameX, true
}

// SetSNotificationtestNameX sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetSNotificationtestNameX(v string) {
	o.SNotificationtestNameX = v
}

// GetENotificationpreferenceStatus returns the ENotificationpreferenceStatus field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetENotificationpreferenceStatus() FieldENotificationpreferenceStatus {
	if o == nil {
		var ret FieldENotificationpreferenceStatus
		return ret
	}

	return o.ENotificationpreferenceStatus
}

// GetENotificationpreferenceStatusOk returns a tuple with the ENotificationpreferenceStatus field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetENotificationpreferenceStatusOk() (*FieldENotificationpreferenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ENotificationpreferenceStatus, true
}

// SetENotificationpreferenceStatus sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetENotificationpreferenceStatus(v FieldENotificationpreferenceStatus) {
	o.ENotificationpreferenceStatus = v
}

// GetINotificationtest returns the INotificationtest field value
func (o *CustomNotificationtestgetnotificationtestsResponse) GetINotificationtest() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.INotificationtest
}

// GetINotificationtestOk returns a tuple with the INotificationtest field value
// and a boolean to check if the value has been set.
func (o *CustomNotificationtestgetnotificationtestsResponse) GetINotificationtestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.INotificationtest, true
}

// SetINotificationtest sets field value
func (o *CustomNotificationtestgetnotificationtestsResponse) SetINotificationtest(v int32) {
	o.INotificationtest = v
}

func (o CustomNotificationtestgetnotificationtestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomNotificationtestgetnotificationtestsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiNotificationtestID"] = o.PkiNotificationtestID
	toSerialize["objNotificationtestName"] = o.ObjNotificationtestName
	toSerialize["fkiNotificationsubsectionID"] = o.FkiNotificationsubsectionID
	toSerialize["sNotificationtestFunction"] = o.SNotificationtestFunction
	toSerialize["sNotificationtestNameX"] = o.SNotificationtestNameX
	toSerialize["eNotificationpreferenceStatus"] = o.ENotificationpreferenceStatus
	toSerialize["iNotificationtest"] = o.INotificationtest
	return toSerialize, nil
}

func (o *CustomNotificationtestgetnotificationtestsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiNotificationtestID",
		"objNotificationtestName",
		"fkiNotificationsubsectionID",
		"sNotificationtestFunction",
		"sNotificationtestNameX",
		"eNotificationpreferenceStatus",
		"iNotificationtest",
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

	varCustomNotificationtestgetnotificationtestsResponse := _CustomNotificationtestgetnotificationtestsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomNotificationtestgetnotificationtestsResponse)

	if err != nil {
		return err
	}

	*o = CustomNotificationtestgetnotificationtestsResponse(varCustomNotificationtestgetnotificationtestsResponse)

	return err
}

type NullableCustomNotificationtestgetnotificationtestsResponse struct {
	value *CustomNotificationtestgetnotificationtestsResponse
	isSet bool
}

func (v NullableCustomNotificationtestgetnotificationtestsResponse) Get() *CustomNotificationtestgetnotificationtestsResponse {
	return v.value
}

func (v *NullableCustomNotificationtestgetnotificationtestsResponse) Set(val *CustomNotificationtestgetnotificationtestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomNotificationtestgetnotificationtestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomNotificationtestgetnotificationtestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomNotificationtestgetnotificationtestsResponse(val *CustomNotificationtestgetnotificationtestsResponse) *NullableCustomNotificationtestgetnotificationtestsResponse {
	return &NullableCustomNotificationtestgetnotificationtestsResponse{value: val, isSet: true}
}

func (v NullableCustomNotificationtestgetnotificationtestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomNotificationtestgetnotificationtestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


