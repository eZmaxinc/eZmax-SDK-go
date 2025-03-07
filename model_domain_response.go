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

// checks if the DomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainResponse{}

// DomainResponse A Domain Object
type DomainResponse struct {
	// The unique ID of the Domain
	PkiDomainID int32 `json:"pkiDomainID"`
	// The name of the Domain
	SDomainName string `json:"sDomainName" validate:"regexp=^(?=.{4,75}$)([a-zA-Z0-9-]+\\\\.)+[a-zA-Z]{2,63}$"`
	// Whether the DKIM is valid or not
	BDomainValiddkim bool `json:"bDomainValiddkim"`
	// Whether the mail from is valid or not
	BDomainValidmailfrom bool `json:"bDomainValidmailfrom"`
	// Whether the customer has access to it or not
	BDomainValidcustomer bool `json:"bDomainValidcustomer"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _DomainResponse DomainResponse

// NewDomainResponse instantiates a new DomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainResponse(pkiDomainID int32, sDomainName string, bDomainValiddkim bool, bDomainValidmailfrom bool, bDomainValidcustomer bool, objAudit CommonAudit) *DomainResponse {
	this := DomainResponse{}
	this.PkiDomainID = pkiDomainID
	this.SDomainName = sDomainName
	this.BDomainValiddkim = bDomainValiddkim
	this.BDomainValidmailfrom = bDomainValidmailfrom
	this.BDomainValidcustomer = bDomainValidcustomer
	this.ObjAudit = objAudit
	return &this
}

// NewDomainResponseWithDefaults instantiates a new DomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainResponseWithDefaults() *DomainResponse {
	this := DomainResponse{}
	return &this
}

// GetPkiDomainID returns the PkiDomainID field value
func (o *DomainResponse) GetPkiDomainID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiDomainID
}

// GetPkiDomainIDOk returns a tuple with the PkiDomainID field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetPkiDomainIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiDomainID, true
}

// SetPkiDomainID sets field value
func (o *DomainResponse) SetPkiDomainID(v int32) {
	o.PkiDomainID = v
}

// GetSDomainName returns the SDomainName field value
func (o *DomainResponse) GetSDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDomainName
}

// GetSDomainNameOk returns a tuple with the SDomainName field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetSDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDomainName, true
}

// SetSDomainName sets field value
func (o *DomainResponse) SetSDomainName(v string) {
	o.SDomainName = v
}

// GetBDomainValiddkim returns the BDomainValiddkim field value
func (o *DomainResponse) GetBDomainValiddkim() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BDomainValiddkim
}

// GetBDomainValiddkimOk returns a tuple with the BDomainValiddkim field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetBDomainValiddkimOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BDomainValiddkim, true
}

// SetBDomainValiddkim sets field value
func (o *DomainResponse) SetBDomainValiddkim(v bool) {
	o.BDomainValiddkim = v
}

// GetBDomainValidmailfrom returns the BDomainValidmailfrom field value
func (o *DomainResponse) GetBDomainValidmailfrom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BDomainValidmailfrom
}

// GetBDomainValidmailfromOk returns a tuple with the BDomainValidmailfrom field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetBDomainValidmailfromOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BDomainValidmailfrom, true
}

// SetBDomainValidmailfrom sets field value
func (o *DomainResponse) SetBDomainValidmailfrom(v bool) {
	o.BDomainValidmailfrom = v
}

// GetBDomainValidcustomer returns the BDomainValidcustomer field value
func (o *DomainResponse) GetBDomainValidcustomer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BDomainValidcustomer
}

// GetBDomainValidcustomerOk returns a tuple with the BDomainValidcustomer field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetBDomainValidcustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BDomainValidcustomer, true
}

// SetBDomainValidcustomer sets field value
func (o *DomainResponse) SetBDomainValidcustomer(v bool) {
	o.BDomainValidcustomer = v
}

// GetObjAudit returns the ObjAudit field value
func (o *DomainResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *DomainResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *DomainResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o DomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiDomainID"] = o.PkiDomainID
	toSerialize["sDomainName"] = o.SDomainName
	toSerialize["bDomainValiddkim"] = o.BDomainValiddkim
	toSerialize["bDomainValidmailfrom"] = o.BDomainValidmailfrom
	toSerialize["bDomainValidcustomer"] = o.BDomainValidcustomer
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *DomainResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiDomainID",
		"sDomainName",
		"bDomainValiddkim",
		"bDomainValidmailfrom",
		"bDomainValidcustomer",
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

	varDomainResponse := _DomainResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDomainResponse)

	if err != nil {
		return err
	}

	*o = DomainResponse(varDomainResponse)

	return err
}

type NullableDomainResponse struct {
	value *DomainResponse
	isSet bool
}

func (v NullableDomainResponse) Get() *DomainResponse {
	return v.value
}

func (v *NullableDomainResponse) Set(val *DomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainResponse(val *DomainResponse) *NullableDomainResponse {
	return &NullableDomainResponse{value: val, isSet: true}
}

func (v NullableDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


