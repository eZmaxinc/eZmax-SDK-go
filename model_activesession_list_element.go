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

// checks if the ActivesessionListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivesessionListElement{}

// ActivesessionListElement A Activesession List Element
type ActivesessionListElement struct {
	// The unique ID of the Activesession
	PkiActivesessionID int32 `json:"pkiActivesessionID"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The unique ID of the Computer
	FkiComputerID int32 `json:"fkiComputerID"`
	// The unique ID of the Company
	FkiCompanyID int32 `json:"fkiCompanyID"`
	// The unique ID of the Department
	FkiDepartmentID int32 `json:"fkiDepartmentID"`
	// The Name of the Company in the language of the requester
	SCompanyNameX string `json:"sCompanyNameX"`
	// The Name of the Department in the language of the requester
	SDepartmentNameX string `json:"sDepartmentNameX"`
	// The loginname of the Activesession
	SActivesessionLoginname string `json:"sActivesessionLoginname" validate:"regexp=^.{0,32}$"`
	// The description of the Computer
	SComputerDescription string `json:"sComputerDescription" validate:"regexp=^.{0,50}$"`
	// The first hit of the Activesession
	DtActivesessionFirsthit string `json:"dtActivesessionFirsthit" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	// The last hit of the Activesession
	DtActivesessionLasthit string `json:"dtActivesessionLasthit" validate:"regexp=^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1]) ([01]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"`
	// Represent an IP address.
	SActivesessionIP string `json:"sActivesessionIP"`
}

type _ActivesessionListElement ActivesessionListElement

// NewActivesessionListElement instantiates a new ActivesessionListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionListElement(pkiActivesessionID int32, fkiUserID int32, fkiComputerID int32, fkiCompanyID int32, fkiDepartmentID int32, sCompanyNameX string, sDepartmentNameX string, sActivesessionLoginname string, sComputerDescription string, dtActivesessionFirsthit string, dtActivesessionLasthit string, sActivesessionIP string) *ActivesessionListElement {
	this := ActivesessionListElement{}
	this.PkiActivesessionID = pkiActivesessionID
	this.FkiUserID = fkiUserID
	this.FkiComputerID = fkiComputerID
	this.FkiCompanyID = fkiCompanyID
	this.FkiDepartmentID = fkiDepartmentID
	this.SCompanyNameX = sCompanyNameX
	this.SDepartmentNameX = sDepartmentNameX
	this.SActivesessionLoginname = sActivesessionLoginname
	this.SComputerDescription = sComputerDescription
	this.DtActivesessionFirsthit = dtActivesessionFirsthit
	this.DtActivesessionLasthit = dtActivesessionLasthit
	this.SActivesessionIP = sActivesessionIP
	return &this
}

// NewActivesessionListElementWithDefaults instantiates a new ActivesessionListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionListElementWithDefaults() *ActivesessionListElement {
	this := ActivesessionListElement{}
	return &this
}

// GetPkiActivesessionID returns the PkiActivesessionID field value
func (o *ActivesessionListElement) GetPkiActivesessionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiActivesessionID
}

// GetPkiActivesessionIDOk returns a tuple with the PkiActivesessionID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetPkiActivesessionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiActivesessionID, true
}

// SetPkiActivesessionID sets field value
func (o *ActivesessionListElement) SetPkiActivesessionID(v int32) {
	o.PkiActivesessionID = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *ActivesessionListElement) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *ActivesessionListElement) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetFkiComputerID returns the FkiComputerID field value
func (o *ActivesessionListElement) GetFkiComputerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiComputerID
}

// GetFkiComputerIDOk returns a tuple with the FkiComputerID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetFkiComputerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiComputerID, true
}

// SetFkiComputerID sets field value
func (o *ActivesessionListElement) SetFkiComputerID(v int32) {
	o.FkiComputerID = v
}

// GetFkiCompanyID returns the FkiCompanyID field value
func (o *ActivesessionListElement) GetFkiCompanyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCompanyID
}

// GetFkiCompanyIDOk returns a tuple with the FkiCompanyID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetFkiCompanyIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiCompanyID, true
}

// SetFkiCompanyID sets field value
func (o *ActivesessionListElement) SetFkiCompanyID(v int32) {
	o.FkiCompanyID = v
}

// GetFkiDepartmentID returns the FkiDepartmentID field value
func (o *ActivesessionListElement) GetFkiDepartmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDepartmentID
}

// GetFkiDepartmentIDOk returns a tuple with the FkiDepartmentID field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetFkiDepartmentIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDepartmentID, true
}

// SetFkiDepartmentID sets field value
func (o *ActivesessionListElement) SetFkiDepartmentID(v int32) {
	o.FkiDepartmentID = v
}

// GetSCompanyNameX returns the SCompanyNameX field value
func (o *ActivesessionListElement) GetSCompanyNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCompanyNameX
}

// GetSCompanyNameXOk returns a tuple with the SCompanyNameX field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetSCompanyNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCompanyNameX, true
}

// SetSCompanyNameX sets field value
func (o *ActivesessionListElement) SetSCompanyNameX(v string) {
	o.SCompanyNameX = v
}

// GetSDepartmentNameX returns the SDepartmentNameX field value
func (o *ActivesessionListElement) GetSDepartmentNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDepartmentNameX
}

// GetSDepartmentNameXOk returns a tuple with the SDepartmentNameX field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetSDepartmentNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDepartmentNameX, true
}

// SetSDepartmentNameX sets field value
func (o *ActivesessionListElement) SetSDepartmentNameX(v string) {
	o.SDepartmentNameX = v
}

// GetSActivesessionLoginname returns the SActivesessionLoginname field value
func (o *ActivesessionListElement) GetSActivesessionLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SActivesessionLoginname
}

// GetSActivesessionLoginnameOk returns a tuple with the SActivesessionLoginname field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetSActivesessionLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SActivesessionLoginname, true
}

// SetSActivesessionLoginname sets field value
func (o *ActivesessionListElement) SetSActivesessionLoginname(v string) {
	o.SActivesessionLoginname = v
}

// GetSComputerDescription returns the SComputerDescription field value
func (o *ActivesessionListElement) GetSComputerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SComputerDescription
}

// GetSComputerDescriptionOk returns a tuple with the SComputerDescription field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetSComputerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SComputerDescription, true
}

// SetSComputerDescription sets field value
func (o *ActivesessionListElement) SetSComputerDescription(v string) {
	o.SComputerDescription = v
}

// GetDtActivesessionFirsthit returns the DtActivesessionFirsthit field value
func (o *ActivesessionListElement) GetDtActivesessionFirsthit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtActivesessionFirsthit
}

// GetDtActivesessionFirsthitOk returns a tuple with the DtActivesessionFirsthit field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetDtActivesessionFirsthitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtActivesessionFirsthit, true
}

// SetDtActivesessionFirsthit sets field value
func (o *ActivesessionListElement) SetDtActivesessionFirsthit(v string) {
	o.DtActivesessionFirsthit = v
}

// GetDtActivesessionLasthit returns the DtActivesessionLasthit field value
func (o *ActivesessionListElement) GetDtActivesessionLasthit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtActivesessionLasthit
}

// GetDtActivesessionLasthitOk returns a tuple with the DtActivesessionLasthit field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetDtActivesessionLasthitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DtActivesessionLasthit, true
}

// SetDtActivesessionLasthit sets field value
func (o *ActivesessionListElement) SetDtActivesessionLasthit(v string) {
	o.DtActivesessionLasthit = v
}

// GetSActivesessionIP returns the SActivesessionIP field value
func (o *ActivesessionListElement) GetSActivesessionIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SActivesessionIP
}

// GetSActivesessionIPOk returns a tuple with the SActivesessionIP field value
// and a boolean to check if the value has been set.
func (o *ActivesessionListElement) GetSActivesessionIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SActivesessionIP, true
}

// SetSActivesessionIP sets field value
func (o *ActivesessionListElement) SetSActivesessionIP(v string) {
	o.SActivesessionIP = v
}

func (o ActivesessionListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivesessionListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiActivesessionID"] = o.PkiActivesessionID
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["fkiComputerID"] = o.FkiComputerID
	toSerialize["fkiCompanyID"] = o.FkiCompanyID
	toSerialize["fkiDepartmentID"] = o.FkiDepartmentID
	toSerialize["sCompanyNameX"] = o.SCompanyNameX
	toSerialize["sDepartmentNameX"] = o.SDepartmentNameX
	toSerialize["sActivesessionLoginname"] = o.SActivesessionLoginname
	toSerialize["sComputerDescription"] = o.SComputerDescription
	toSerialize["dtActivesessionFirsthit"] = o.DtActivesessionFirsthit
	toSerialize["dtActivesessionLasthit"] = o.DtActivesessionLasthit
	toSerialize["sActivesessionIP"] = o.SActivesessionIP
	return toSerialize, nil
}

func (o *ActivesessionListElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiActivesessionID",
		"fkiUserID",
		"fkiComputerID",
		"fkiCompanyID",
		"fkiDepartmentID",
		"sCompanyNameX",
		"sDepartmentNameX",
		"sActivesessionLoginname",
		"sComputerDescription",
		"dtActivesessionFirsthit",
		"dtActivesessionLasthit",
		"sActivesessionIP",
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

	varActivesessionListElement := _ActivesessionListElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivesessionListElement)

	if err != nil {
		return err
	}

	*o = ActivesessionListElement(varActivesessionListElement)

	return err
}

type NullableActivesessionListElement struct {
	value *ActivesessionListElement
	isSet bool
}

func (v NullableActivesessionListElement) Get() *ActivesessionListElement {
	return v.value
}

func (v *NullableActivesessionListElement) Set(val *ActivesessionListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionListElement(val *ActivesessionListElement) *NullableActivesessionListElement {
	return &NullableActivesessionListElement{value: val, isSet: true}
}

func (v NullableActivesessionListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


