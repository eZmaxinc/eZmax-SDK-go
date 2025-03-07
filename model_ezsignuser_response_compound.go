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

// checks if the EzsignuserResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignuserResponseCompound{}

// EzsignuserResponseCompound A Ezsignuser Object
type EzsignuserResponseCompound struct {
	// The unique ID of the Ezsignuser
	PkiEzsignuserID int32 `json:"pkiEzsignuserID"`
	// The unique ID of the Contact
	FkiContactID int32 `json:"fkiContactID"`
	ObjContact ContactResponseCompound `json:"objContact"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _EzsignuserResponseCompound EzsignuserResponseCompound

// NewEzsignuserResponseCompound instantiates a new EzsignuserResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignuserResponseCompound(pkiEzsignuserID int32, fkiContactID int32, objContact ContactResponseCompound, objAudit CommonAudit) *EzsignuserResponseCompound {
	this := EzsignuserResponseCompound{}
	this.PkiEzsignuserID = pkiEzsignuserID
	this.FkiContactID = fkiContactID
	this.ObjContact = objContact
	this.ObjAudit = objAudit
	return &this
}

// NewEzsignuserResponseCompoundWithDefaults instantiates a new EzsignuserResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignuserResponseCompoundWithDefaults() *EzsignuserResponseCompound {
	this := EzsignuserResponseCompound{}
	return &this
}

// GetPkiEzsignuserID returns the PkiEzsignuserID field value
func (o *EzsignuserResponseCompound) GetPkiEzsignuserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignuserID
}

// GetPkiEzsignuserIDOk returns a tuple with the PkiEzsignuserID field value
// and a boolean to check if the value has been set.
func (o *EzsignuserResponseCompound) GetPkiEzsignuserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignuserID, true
}

// SetPkiEzsignuserID sets field value
func (o *EzsignuserResponseCompound) SetPkiEzsignuserID(v int32) {
	o.PkiEzsignuserID = v
}

// GetFkiContactID returns the FkiContactID field value
func (o *EzsignuserResponseCompound) GetFkiContactID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiContactID
}

// GetFkiContactIDOk returns a tuple with the FkiContactID field value
// and a boolean to check if the value has been set.
func (o *EzsignuserResponseCompound) GetFkiContactIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiContactID, true
}

// SetFkiContactID sets field value
func (o *EzsignuserResponseCompound) SetFkiContactID(v int32) {
	o.FkiContactID = v
}

// GetObjContact returns the ObjContact field value
func (o *EzsignuserResponseCompound) GetObjContact() ContactResponseCompound {
	if o == nil {
		var ret ContactResponseCompound
		return ret
	}

	return o.ObjContact
}

// GetObjContactOk returns a tuple with the ObjContact field value
// and a boolean to check if the value has been set.
func (o *EzsignuserResponseCompound) GetObjContactOk() (*ContactResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjContact, true
}

// SetObjContact sets field value
func (o *EzsignuserResponseCompound) SetObjContact(v ContactResponseCompound) {
	o.ObjContact = v
}

// GetObjAudit returns the ObjAudit field value
func (o *EzsignuserResponseCompound) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *EzsignuserResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *EzsignuserResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o EzsignuserResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignuserResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignuserID"] = o.PkiEzsignuserID
	toSerialize["fkiContactID"] = o.FkiContactID
	toSerialize["objContact"] = o.ObjContact
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *EzsignuserResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignuserID",
		"fkiContactID",
		"objContact",
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

	varEzsignuserResponseCompound := _EzsignuserResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignuserResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignuserResponseCompound(varEzsignuserResponseCompound)

	return err
}

type NullableEzsignuserResponseCompound struct {
	value *EzsignuserResponseCompound
	isSet bool
}

func (v NullableEzsignuserResponseCompound) Get() *EzsignuserResponseCompound {
	return v.value
}

func (v *NullableEzsignuserResponseCompound) Set(val *EzsignuserResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignuserResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignuserResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignuserResponseCompound(val *EzsignuserResponseCompound) *NullableEzsignuserResponseCompound {
	return &NullableEzsignuserResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignuserResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignuserResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


