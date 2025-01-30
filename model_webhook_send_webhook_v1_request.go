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

// checks if the WebhookSendWebhookV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSendWebhookV1Request{}

// WebhookSendWebhookV1Request Request for POST /1/object/webhook/sendWebhook
type WebhookSendWebhookV1Request struct {
	EWebhookModule FieldEWebhookModule `json:"eWebhookModule"`
	EWebhookEzsignevent *CustomEWebhookEzsignevent `json:"eWebhookEzsignevent,omitempty"`
	EWebhookManagementevent *FieldEWebhookManagementevent `json:"eWebhookManagementevent,omitempty"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID *int32 `json:"fkiEzsignfolderID,omitempty"`
	// The unique ID of the Ezsigndocument
	FkiEzsigndocumentID *int32 `json:"fkiEzsigndocumentID,omitempty"`
	// The unique ID of the Ezsignsigner
	FkiEzsignsignerID *int32 `json:"fkiEzsignsignerID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Userstaged
	FkiUserstagedID *int32 `json:"fkiUserstagedID,omitempty"`
}

type _WebhookSendWebhookV1Request WebhookSendWebhookV1Request

// NewWebhookSendWebhookV1Request instantiates a new WebhookSendWebhookV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSendWebhookV1Request(eWebhookModule FieldEWebhookModule) *WebhookSendWebhookV1Request {
	this := WebhookSendWebhookV1Request{}
	this.EWebhookModule = eWebhookModule
	return &this
}

// NewWebhookSendWebhookV1RequestWithDefaults instantiates a new WebhookSendWebhookV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSendWebhookV1RequestWithDefaults() *WebhookSendWebhookV1Request {
	this := WebhookSendWebhookV1Request{}
	return &this
}

// GetEWebhookModule returns the EWebhookModule field value
func (o *WebhookSendWebhookV1Request) GetEWebhookModule() FieldEWebhookModule {
	if o == nil {
		var ret FieldEWebhookModule
		return ret
	}

	return o.EWebhookModule
}

// GetEWebhookModuleOk returns a tuple with the EWebhookModule field value
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetEWebhookModuleOk() (*FieldEWebhookModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebhookModule, true
}

// SetEWebhookModule sets field value
func (o *WebhookSendWebhookV1Request) SetEWebhookModule(v FieldEWebhookModule) {
	o.EWebhookModule = v
}

// GetEWebhookEzsignevent returns the EWebhookEzsignevent field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetEWebhookEzsignevent() CustomEWebhookEzsignevent {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		var ret CustomEWebhookEzsignevent
		return ret
	}
	return *o.EWebhookEzsignevent
}

// GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetEWebhookEzsigneventOk() (*CustomEWebhookEzsignevent, bool) {
	if o == nil || IsNil(o.EWebhookEzsignevent) {
		return nil, false
	}
	return o.EWebhookEzsignevent, true
}

// HasEWebhookEzsignevent returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasEWebhookEzsignevent() bool {
	if o != nil && !IsNil(o.EWebhookEzsignevent) {
		return true
	}

	return false
}

// SetEWebhookEzsignevent gets a reference to the given CustomEWebhookEzsignevent and assigns it to the EWebhookEzsignevent field.
func (o *WebhookSendWebhookV1Request) SetEWebhookEzsignevent(v CustomEWebhookEzsignevent) {
	o.EWebhookEzsignevent = &v
}

// GetEWebhookManagementevent returns the EWebhookManagementevent field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetEWebhookManagementevent() FieldEWebhookManagementevent {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		var ret FieldEWebhookManagementevent
		return ret
	}
	return *o.EWebhookManagementevent
}

// GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool) {
	if o == nil || IsNil(o.EWebhookManagementevent) {
		return nil, false
	}
	return o.EWebhookManagementevent, true
}

// HasEWebhookManagementevent returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasEWebhookManagementevent() bool {
	if o != nil && !IsNil(o.EWebhookManagementevent) {
		return true
	}

	return false
}

// SetEWebhookManagementevent gets a reference to the given FieldEWebhookManagementevent and assigns it to the EWebhookManagementevent field.
func (o *WebhookSendWebhookV1Request) SetEWebhookManagementevent(v FieldEWebhookManagementevent) {
	o.EWebhookManagementevent = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetFkiEzsignfolderID() int32 {
	if o == nil || IsNil(o.FkiEzsignfolderID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignfolderID) {
		return nil, false
	}
	return o.FkiEzsignfolderID, true
}

// HasFkiEzsignfolderID returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasFkiEzsignfolderID() bool {
	if o != nil && !IsNil(o.FkiEzsignfolderID) {
		return true
	}

	return false
}

// SetFkiEzsignfolderID gets a reference to the given int32 and assigns it to the FkiEzsignfolderID field.
func (o *WebhookSendWebhookV1Request) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = &v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetFkiEzsigndocumentID() int32 {
	if o == nil || IsNil(o.FkiEzsigndocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsigndocumentID) {
		return nil, false
	}
	return o.FkiEzsigndocumentID, true
}

// HasFkiEzsigndocumentID returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasFkiEzsigndocumentID() bool {
	if o != nil && !IsNil(o.FkiEzsigndocumentID) {
		return true
	}

	return false
}

// SetFkiEzsigndocumentID gets a reference to the given int32 and assigns it to the FkiEzsigndocumentID field.
func (o *WebhookSendWebhookV1Request) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = &v
}

// GetFkiEzsignsignerID returns the FkiEzsignsignerID field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetFkiEzsignsignerID() int32 {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignsignerID
}

// GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetFkiEzsignsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignsignerID) {
		return nil, false
	}
	return o.FkiEzsignsignerID, true
}

// HasFkiEzsignsignerID returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasFkiEzsignsignerID() bool {
	if o != nil && !IsNil(o.FkiEzsignsignerID) {
		return true
	}

	return false
}

// SetFkiEzsignsignerID gets a reference to the given int32 and assigns it to the FkiEzsignsignerID field.
func (o *WebhookSendWebhookV1Request) SetFkiEzsignsignerID(v int32) {
	o.FkiEzsignsignerID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *WebhookSendWebhookV1Request) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUserstagedID returns the FkiUserstagedID field value if set, zero value otherwise.
func (o *WebhookSendWebhookV1Request) GetFkiUserstagedID() int32 {
	if o == nil || IsNil(o.FkiUserstagedID) {
		var ret int32
		return ret
	}
	return *o.FkiUserstagedID
}

// GetFkiUserstagedIDOk returns a tuple with the FkiUserstagedID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSendWebhookV1Request) GetFkiUserstagedIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserstagedID) {
		return nil, false
	}
	return o.FkiUserstagedID, true
}

// HasFkiUserstagedID returns a boolean if a field has been set.
func (o *WebhookSendWebhookV1Request) HasFkiUserstagedID() bool {
	if o != nil && !IsNil(o.FkiUserstagedID) {
		return true
	}

	return false
}

// SetFkiUserstagedID gets a reference to the given int32 and assigns it to the FkiUserstagedID field.
func (o *WebhookSendWebhookV1Request) SetFkiUserstagedID(v int32) {
	o.FkiUserstagedID = &v
}

func (o WebhookSendWebhookV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSendWebhookV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eWebhookModule"] = o.EWebhookModule
	if !IsNil(o.EWebhookEzsignevent) {
		toSerialize["eWebhookEzsignevent"] = o.EWebhookEzsignevent
	}
	if !IsNil(o.EWebhookManagementevent) {
		toSerialize["eWebhookManagementevent"] = o.EWebhookManagementevent
	}
	if !IsNil(o.FkiEzsignfolderID) {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	if !IsNil(o.FkiEzsigndocumentID) {
		toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	}
	if !IsNil(o.FkiEzsignsignerID) {
		toSerialize["fkiEzsignsignerID"] = o.FkiEzsignsignerID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUserstagedID) {
		toSerialize["fkiUserstagedID"] = o.FkiUserstagedID
	}
	return toSerialize, nil
}

func (o *WebhookSendWebhookV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eWebhookModule",
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

	varWebhookSendWebhookV1Request := _WebhookSendWebhookV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookSendWebhookV1Request)

	if err != nil {
		return err
	}

	*o = WebhookSendWebhookV1Request(varWebhookSendWebhookV1Request)

	return err
}

type NullableWebhookSendWebhookV1Request struct {
	value *WebhookSendWebhookV1Request
	isSet bool
}

func (v NullableWebhookSendWebhookV1Request) Get() *WebhookSendWebhookV1Request {
	return v.value
}

func (v *NullableWebhookSendWebhookV1Request) Set(val *WebhookSendWebhookV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSendWebhookV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSendWebhookV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSendWebhookV1Request(val *WebhookSendWebhookV1Request) *NullableWebhookSendWebhookV1Request {
	return &NullableWebhookSendWebhookV1Request{value: val, isSet: true}
}

func (v NullableWebhookSendWebhookV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSendWebhookV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


