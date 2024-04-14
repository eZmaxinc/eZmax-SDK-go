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

// checks if the UserlogintypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserlogintypeResponse{}

// UserlogintypeResponse An Userlogintype Object
type UserlogintypeResponse struct {
	// The unique ID of the Userlogintype  Valid values:  |Value|Description|Detail| |-|-|-| |1|**Email Only**|The Ezsignsigner will receive a secure link by email| |2|**Email and phone or SMS**|The Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**| |3|**Email and secret question**|The Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer| |4|**In person only**|The Ezsignsigner will only be able to sign \"In-Person\" and there won't be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type| |5|**In person with phone or SMS**|The Ezsignsigner will only be able to sign \"In-Person\" and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**|
	PkiUserlogintypeID int32 `json:"pkiUserlogintypeID"`
	ObjUserlogintypeDescription MultilingualUserlogintypeDescription `json:"objUserlogintypeDescription"`
	// The description of the Userlogintype in the language of the requester
	SUserlogintypeDescriptionX string `json:"sUserlogintypeDescriptionX"`
}

type _UserlogintypeResponse UserlogintypeResponse

// NewUserlogintypeResponse instantiates a new UserlogintypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserlogintypeResponse(pkiUserlogintypeID int32, objUserlogintypeDescription MultilingualUserlogintypeDescription, sUserlogintypeDescriptionX string) *UserlogintypeResponse {
	this := UserlogintypeResponse{}
	this.PkiUserlogintypeID = pkiUserlogintypeID
	this.ObjUserlogintypeDescription = objUserlogintypeDescription
	this.SUserlogintypeDescriptionX = sUserlogintypeDescriptionX
	return &this
}

// NewUserlogintypeResponseWithDefaults instantiates a new UserlogintypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserlogintypeResponseWithDefaults() *UserlogintypeResponse {
	this := UserlogintypeResponse{}
	return &this
}

// GetPkiUserlogintypeID returns the PkiUserlogintypeID field value
func (o *UserlogintypeResponse) GetPkiUserlogintypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserlogintypeID
}

// GetPkiUserlogintypeIDOk returns a tuple with the PkiUserlogintypeID field value
// and a boolean to check if the value has been set.
func (o *UserlogintypeResponse) GetPkiUserlogintypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUserlogintypeID, true
}

// SetPkiUserlogintypeID sets field value
func (o *UserlogintypeResponse) SetPkiUserlogintypeID(v int32) {
	o.PkiUserlogintypeID = v
}

// GetObjUserlogintypeDescription returns the ObjUserlogintypeDescription field value
func (o *UserlogintypeResponse) GetObjUserlogintypeDescription() MultilingualUserlogintypeDescription {
	if o == nil {
		var ret MultilingualUserlogintypeDescription
		return ret
	}

	return o.ObjUserlogintypeDescription
}

// GetObjUserlogintypeDescriptionOk returns a tuple with the ObjUserlogintypeDescription field value
// and a boolean to check if the value has been set.
func (o *UserlogintypeResponse) GetObjUserlogintypeDescriptionOk() (*MultilingualUserlogintypeDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjUserlogintypeDescription, true
}

// SetObjUserlogintypeDescription sets field value
func (o *UserlogintypeResponse) SetObjUserlogintypeDescription(v MultilingualUserlogintypeDescription) {
	o.ObjUserlogintypeDescription = v
}

// GetSUserlogintypeDescriptionX returns the SUserlogintypeDescriptionX field value
func (o *UserlogintypeResponse) GetSUserlogintypeDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserlogintypeDescriptionX
}

// GetSUserlogintypeDescriptionXOk returns a tuple with the SUserlogintypeDescriptionX field value
// and a boolean to check if the value has been set.
func (o *UserlogintypeResponse) GetSUserlogintypeDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserlogintypeDescriptionX, true
}

// SetSUserlogintypeDescriptionX sets field value
func (o *UserlogintypeResponse) SetSUserlogintypeDescriptionX(v string) {
	o.SUserlogintypeDescriptionX = v
}

func (o UserlogintypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserlogintypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUserlogintypeID"] = o.PkiUserlogintypeID
	toSerialize["objUserlogintypeDescription"] = o.ObjUserlogintypeDescription
	toSerialize["sUserlogintypeDescriptionX"] = o.SUserlogintypeDescriptionX
	return toSerialize, nil
}

func (o *UserlogintypeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUserlogintypeID",
		"objUserlogintypeDescription",
		"sUserlogintypeDescriptionX",
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

	varUserlogintypeResponse := _UserlogintypeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserlogintypeResponse)

	if err != nil {
		return err
	}

	*o = UserlogintypeResponse(varUserlogintypeResponse)

	return err
}

type NullableUserlogintypeResponse struct {
	value *UserlogintypeResponse
	isSet bool
}

func (v NullableUserlogintypeResponse) Get() *UserlogintypeResponse {
	return v.value
}

func (v *NullableUserlogintypeResponse) Set(val *UserlogintypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserlogintypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserlogintypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserlogintypeResponse(val *UserlogintypeResponse) *NullableUserlogintypeResponse {
	return &NullableUserlogintypeResponse{value: val, isSet: true}
}

func (v NullableUserlogintypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserlogintypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


