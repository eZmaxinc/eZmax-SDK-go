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

// checks if the CustomEzsignfoldersignerassociationstatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignfoldersignerassociationstatusResponse{}

// CustomEzsignfoldersignerassociationstatusResponse A Ezsignfoldersignerassociationstatus Object and children to create a complete structure
type CustomEzsignfoldersignerassociationstatusResponse struct {
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
	// The last name of the Ezsignsigner
	// Deprecated
	SEzsignfoldersignerassociationstatusLastname *string `json:"sEzsignfoldersignerassociationstatusLastname,omitempty"`
	// The first name of the Ezsignsigner
	// Deprecated
	SEzsignfoldersignerassociationstatusFirstname *string `json:"sEzsignfoldersignerassociationstatusFirstname,omitempty"`
	// The description of the Ezsignsigner
	SEzsignfoldersignerassociationstatusDescriptionX *string `json:"sEzsignfoldersignerassociationstatusDescriptionX,omitempty"`
	AObjEzsignsignaturestatus []CustomEzsignsignaturestatusResponse `json:"a_objEzsignsignaturestatus"`
}

type _CustomEzsignfoldersignerassociationstatusResponse CustomEzsignfoldersignerassociationstatusResponse

// NewCustomEzsignfoldersignerassociationstatusResponse instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignfoldersignerassociationstatusResponse(fkiEzsignfoldersignerassociationID int32, aObjEzsignsignaturestatus []CustomEzsignsignaturestatusResponse) *CustomEzsignfoldersignerassociationstatusResponse {
	this := CustomEzsignfoldersignerassociationstatusResponse{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	this.AObjEzsignsignaturestatus = aObjEzsignsignaturestatus
	return &this
}

// NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults instantiates a new CustomEzsignfoldersignerassociationstatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignfoldersignerassociationstatusResponseWithDefaults() *CustomEzsignfoldersignerassociationstatusResponse {
	this := CustomEzsignfoldersignerassociationstatusResponse{}
	return &this
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

// GetSEzsignfoldersignerassociationstatusLastname returns the SEzsignfoldersignerassociationstatusLastname field value if set, zero value otherwise.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastname() string {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusLastname) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldersignerassociationstatusLastname
}

// GetSEzsignfoldersignerassociationstatusLastnameOk returns a tuple with the SEzsignfoldersignerassociationstatusLastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusLastname) {
		return nil, false
	}
	return o.SEzsignfoldersignerassociationstatusLastname, true
}

// HasSEzsignfoldersignerassociationstatusLastname returns a boolean if a field has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusLastname() bool {
	if o != nil && !IsNil(o.SEzsignfoldersignerassociationstatusLastname) {
		return true
	}

	return false
}

// SetSEzsignfoldersignerassociationstatusLastname gets a reference to the given string and assigns it to the SEzsignfoldersignerassociationstatusLastname field.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusLastname(v string) {
	o.SEzsignfoldersignerassociationstatusLastname = &v
}

// GetSEzsignfoldersignerassociationstatusFirstname returns the SEzsignfoldersignerassociationstatusFirstname field value if set, zero value otherwise.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstname() string {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusFirstname) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldersignerassociationstatusFirstname
}

// GetSEzsignfoldersignerassociationstatusFirstnameOk returns a tuple with the SEzsignfoldersignerassociationstatusFirstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusFirstname) {
		return nil, false
	}
	return o.SEzsignfoldersignerassociationstatusFirstname, true
}

// HasSEzsignfoldersignerassociationstatusFirstname returns a boolean if a field has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusFirstname() bool {
	if o != nil && !IsNil(o.SEzsignfoldersignerassociationstatusFirstname) {
		return true
	}

	return false
}

// SetSEzsignfoldersignerassociationstatusFirstname gets a reference to the given string and assigns it to the SEzsignfoldersignerassociationstatusFirstname field.
// Deprecated
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusFirstname(v string) {
	o.SEzsignfoldersignerassociationstatusFirstname = &v
}

// GetSEzsignfoldersignerassociationstatusDescriptionX returns the SEzsignfoldersignerassociationstatusDescriptionX field value if set, zero value otherwise.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusDescriptionX() string {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusDescriptionX) {
		var ret string
		return ret
	}
	return *o.SEzsignfoldersignerassociationstatusDescriptionX
}

// GetSEzsignfoldersignerassociationstatusDescriptionXOk returns a tuple with the SEzsignfoldersignerassociationstatusDescriptionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetSEzsignfoldersignerassociationstatusDescriptionXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzsignfoldersignerassociationstatusDescriptionX) {
		return nil, false
	}
	return o.SEzsignfoldersignerassociationstatusDescriptionX, true
}

// HasSEzsignfoldersignerassociationstatusDescriptionX returns a boolean if a field has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) HasSEzsignfoldersignerassociationstatusDescriptionX() bool {
	if o != nil && !IsNil(o.SEzsignfoldersignerassociationstatusDescriptionX) {
		return true
	}

	return false
}

// SetSEzsignfoldersignerassociationstatusDescriptionX gets a reference to the given string and assigns it to the SEzsignfoldersignerassociationstatusDescriptionX field.
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetSEzsignfoldersignerassociationstatusDescriptionX(v string) {
	o.SEzsignfoldersignerassociationstatusDescriptionX = &v
}

// GetAObjEzsignsignaturestatus returns the AObjEzsignsignaturestatus field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatus() []CustomEzsignsignaturestatusResponse {
	if o == nil {
		var ret []CustomEzsignsignaturestatusResponse
		return ret
	}

	return o.AObjEzsignsignaturestatus
}

// GetAObjEzsignsignaturestatusOk returns a tuple with the AObjEzsignsignaturestatus field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldersignerassociationstatusResponse) GetAObjEzsignsignaturestatusOk() ([]CustomEzsignsignaturestatusResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignaturestatus, true
}

// SetAObjEzsignsignaturestatus sets field value
func (o *CustomEzsignfoldersignerassociationstatusResponse) SetAObjEzsignsignaturestatus(v []CustomEzsignsignaturestatusResponse) {
	o.AObjEzsignsignaturestatus = v
}

func (o CustomEzsignfoldersignerassociationstatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignfoldersignerassociationstatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	if !IsNil(o.SEzsignfoldersignerassociationstatusLastname) {
		toSerialize["sEzsignfoldersignerassociationstatusLastname"] = o.SEzsignfoldersignerassociationstatusLastname
	}
	if !IsNil(o.SEzsignfoldersignerassociationstatusFirstname) {
		toSerialize["sEzsignfoldersignerassociationstatusFirstname"] = o.SEzsignfoldersignerassociationstatusFirstname
	}
	if !IsNil(o.SEzsignfoldersignerassociationstatusDescriptionX) {
		toSerialize["sEzsignfoldersignerassociationstatusDescriptionX"] = o.SEzsignfoldersignerassociationstatusDescriptionX
	}
	toSerialize["a_objEzsignsignaturestatus"] = o.AObjEzsignsignaturestatus
	return toSerialize, nil
}

func (o *CustomEzsignfoldersignerassociationstatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfoldersignerassociationID",
		"a_objEzsignsignaturestatus",
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

	varCustomEzsignfoldersignerassociationstatusResponse := _CustomEzsignfoldersignerassociationstatusResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomEzsignfoldersignerassociationstatusResponse)

	if err != nil {
		return err
	}

	*o = CustomEzsignfoldersignerassociationstatusResponse(varCustomEzsignfoldersignerassociationstatusResponse)

	return err
}

type NullableCustomEzsignfoldersignerassociationstatusResponse struct {
	value *CustomEzsignfoldersignerassociationstatusResponse
	isSet bool
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) Get() *CustomEzsignfoldersignerassociationstatusResponse {
	return v.value
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) Set(val *CustomEzsignfoldersignerassociationstatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignfoldersignerassociationstatusResponse(val *CustomEzsignfoldersignerassociationstatusResponse) *NullableCustomEzsignfoldersignerassociationstatusResponse {
	return &NullableCustomEzsignfoldersignerassociationstatusResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignfoldersignerassociationstatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignfoldersignerassociationstatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


