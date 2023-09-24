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

// checks if the EzsigntemplatepackagemembershipResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagemembershipResponseCompound{}

// EzsigntemplatepackagemembershipResponseCompound A Ezsigntemplatepackagemembership Object
type EzsigntemplatepackagemembershipResponseCompound struct {
	// The unique ID of the Ezsigntemplatepackagemembership
	PkiEzsigntemplatepackagemembershipID int32 `json:"pkiEzsigntemplatepackagemembershipID"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The order in which the Ezsigntemplate will be imported when using an Ezsigntemplatepackage.
	IEzsigntemplatepackagemembershipOrder int32 `json:"iEzsigntemplatepackagemembershipOrder"`
	ObjEzsigntemplate EzsigntemplateResponseCompound `json:"objEzsigntemplate"`
	AObjEzsigntemplatepackagesignermembership []EzsigntemplatepackagesignermembershipResponseCompound `json:"a_objEzsigntemplatepackagesignermembership"`
}

// NewEzsigntemplatepackagemembershipResponseCompound instantiates a new EzsigntemplatepackagemembershipResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagemembershipResponseCompound(pkiEzsigntemplatepackagemembershipID int32, fkiEzsigntemplatepackageID int32, fkiEzsigntemplateID int32, iEzsigntemplatepackagemembershipOrder int32, objEzsigntemplate EzsigntemplateResponseCompound, aObjEzsigntemplatepackagesignermembership []EzsigntemplatepackagesignermembershipResponseCompound) *EzsigntemplatepackagemembershipResponseCompound {
	this := EzsigntemplatepackagemembershipResponseCompound{}
	this.PkiEzsigntemplatepackagemembershipID = pkiEzsigntemplatepackagemembershipID
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.IEzsigntemplatepackagemembershipOrder = iEzsigntemplatepackagemembershipOrder
	this.ObjEzsigntemplate = objEzsigntemplate
	this.AObjEzsigntemplatepackagesignermembership = aObjEzsigntemplatepackagesignermembership
	return &this
}

// NewEzsigntemplatepackagemembershipResponseCompoundWithDefaults instantiates a new EzsigntemplatepackagemembershipResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagemembershipResponseCompoundWithDefaults() *EzsigntemplatepackagemembershipResponseCompound {
	this := EzsigntemplatepackagemembershipResponseCompound{}
	return &this
}

// GetPkiEzsigntemplatepackagemembershipID returns the PkiEzsigntemplatepackagemembershipID field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetPkiEzsigntemplatepackagemembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackagemembershipID
}

// GetPkiEzsigntemplatepackagemembershipIDOk returns a tuple with the PkiEzsigntemplatepackagemembershipID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetPkiEzsigntemplatepackagemembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackagemembershipID, true
}

// SetPkiEzsigntemplatepackagemembershipID sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetPkiEzsigntemplatepackagemembershipID(v int32) {
	o.PkiEzsigntemplatepackagemembershipID = v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetIEzsigntemplatepackagemembershipOrder returns the IEzsigntemplatepackagemembershipOrder field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetIEzsigntemplatepackagemembershipOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsigntemplatepackagemembershipOrder
}

// GetIEzsigntemplatepackagemembershipOrderOk returns a tuple with the IEzsigntemplatepackagemembershipOrder field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetIEzsigntemplatepackagemembershipOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEzsigntemplatepackagemembershipOrder, true
}

// SetIEzsigntemplatepackagemembershipOrder sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetIEzsigntemplatepackagemembershipOrder(v int32) {
	o.IEzsigntemplatepackagemembershipOrder = v
}

// GetObjEzsigntemplate returns the ObjEzsigntemplate field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetObjEzsigntemplate() EzsigntemplateResponseCompound {
	if o == nil {
		var ret EzsigntemplateResponseCompound
		return ret
	}

	return o.ObjEzsigntemplate
}

// GetObjEzsigntemplateOk returns a tuple with the ObjEzsigntemplate field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetObjEzsigntemplateOk() (*EzsigntemplateResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsigntemplate, true
}

// SetObjEzsigntemplate sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetObjEzsigntemplate(v EzsigntemplateResponseCompound) {
	o.ObjEzsigntemplate = v
}

// GetAObjEzsigntemplatepackagesignermembership returns the AObjEzsigntemplatepackagesignermembership field value
func (o *EzsigntemplatepackagemembershipResponseCompound) GetAObjEzsigntemplatepackagesignermembership() []EzsigntemplatepackagesignermembershipResponseCompound {
	if o == nil {
		var ret []EzsigntemplatepackagesignermembershipResponseCompound
		return ret
	}

	return o.AObjEzsigntemplatepackagesignermembership
}

// GetAObjEzsigntemplatepackagesignermembershipOk returns a tuple with the AObjEzsigntemplatepackagesignermembership field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagemembershipResponseCompound) GetAObjEzsigntemplatepackagesignermembershipOk() ([]EzsigntemplatepackagesignermembershipResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatepackagesignermembership, true
}

// SetAObjEzsigntemplatepackagesignermembership sets field value
func (o *EzsigntemplatepackagemembershipResponseCompound) SetAObjEzsigntemplatepackagesignermembership(v []EzsigntemplatepackagesignermembershipResponseCompound) {
	o.AObjEzsigntemplatepackagesignermembership = v
}

func (o EzsigntemplatepackagemembershipResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagemembershipResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatepackagemembershipID"] = o.PkiEzsigntemplatepackagemembershipID
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["iEzsigntemplatepackagemembershipOrder"] = o.IEzsigntemplatepackagemembershipOrder
	toSerialize["objEzsigntemplate"] = o.ObjEzsigntemplate
	toSerialize["a_objEzsigntemplatepackagesignermembership"] = o.AObjEzsigntemplatepackagesignermembership
	return toSerialize, nil
}

type NullableEzsigntemplatepackagemembershipResponseCompound struct {
	value *EzsigntemplatepackagemembershipResponseCompound
	isSet bool
}

func (v NullableEzsigntemplatepackagemembershipResponseCompound) Get() *EzsigntemplatepackagemembershipResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplatepackagemembershipResponseCompound) Set(val *EzsigntemplatepackagemembershipResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagemembershipResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagemembershipResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagemembershipResponseCompound(val *EzsigntemplatepackagemembershipResponseCompound) *NullableEzsigntemplatepackagemembershipResponseCompound {
	return &NullableEzsigntemplatepackagemembershipResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagemembershipResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagemembershipResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


