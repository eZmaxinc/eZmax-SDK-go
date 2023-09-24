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

// checks if the CustomEzsignfoldertypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomEzsignfoldertypeResponse{}

// CustomEzsignfoldertypeResponse A Custom Ezsignfoldertype Object
type CustomEzsignfoldertypeResponse struct {
	// The unique ID of the Ezsignfoldertype.
	PkiEzsignfoldertypeID int32 `json:"pkiEzsignfoldertypeID"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
	// Whether we include the proof with the signed Ezsigndocument for Ezsignsigners
	BEzsignfoldertypeIncludeproofsigner bool `json:"bEzsignfoldertypeIncludeproofsigner"`
	// Whether we include the proof with the signed Ezsigndocument for users
	BEzsignfoldertypeIncludeproofuser bool `json:"bEzsignfoldertypeIncludeproofuser"`
	// Wheter if delegation of signature is allowed to another user or not
	BEzsignfoldertypeDelegate *bool `json:"bEzsignfoldertypeDelegate,omitempty"`
	// Wheter if Reassignment of signature is allowed to another signatory or not
	BEzsignfoldertypeReassign *bool `json:"bEzsignfoldertypeReassign,omitempty"`
}

// NewCustomEzsignfoldertypeResponse instantiates a new CustomEzsignfoldertypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEzsignfoldertypeResponse(pkiEzsignfoldertypeID int32, sEzsignfoldertypeNameX string, bEzsignfoldertypeIncludeproofsigner bool, bEzsignfoldertypeIncludeproofuser bool) *CustomEzsignfoldertypeResponse {
	this := CustomEzsignfoldertypeResponse{}
	this.PkiEzsignfoldertypeID = pkiEzsignfoldertypeID
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.BEzsignfoldertypeIncludeproofsigner = bEzsignfoldertypeIncludeproofsigner
	this.BEzsignfoldertypeIncludeproofuser = bEzsignfoldertypeIncludeproofuser
	return &this
}

// NewCustomEzsignfoldertypeResponseWithDefaults instantiates a new CustomEzsignfoldertypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEzsignfoldertypeResponseWithDefaults() *CustomEzsignfoldertypeResponse {
	this := CustomEzsignfoldertypeResponse{}
	return &this
}

// GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field value
func (o *CustomEzsignfoldertypeResponse) GetPkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfoldertypeID
}

// GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetPkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignfoldertypeID, true
}

// SetPkiEzsignfoldertypeID sets field value
func (o *CustomEzsignfoldertypeResponse) SetPkiEzsignfoldertypeID(v int32) {
	o.PkiEzsignfoldertypeID = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *CustomEzsignfoldertypeResponse) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *CustomEzsignfoldertypeResponse) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

// GetBEzsignfoldertypeIncludeproofsigner returns the BEzsignfoldertypeIncludeproofsigner field value
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofsigner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignfoldertypeIncludeproofsigner
}

// GetBEzsignfoldertypeIncludeproofsignerOk returns a tuple with the BEzsignfoldertypeIncludeproofsigner field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofsignerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignfoldertypeIncludeproofsigner, true
}

// SetBEzsignfoldertypeIncludeproofsigner sets field value
func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeIncludeproofsigner(v bool) {
	o.BEzsignfoldertypeIncludeproofsigner = v
}

// GetBEzsignfoldertypeIncludeproofuser returns the BEzsignfoldertypeIncludeproofuser field value
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofuser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignfoldertypeIncludeproofuser
}

// GetBEzsignfoldertypeIncludeproofuserOk returns a tuple with the BEzsignfoldertypeIncludeproofuser field value
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeIncludeproofuserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BEzsignfoldertypeIncludeproofuser, true
}

// SetBEzsignfoldertypeIncludeproofuser sets field value
func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeIncludeproofuser(v bool) {
	o.BEzsignfoldertypeIncludeproofuser = v
}

// GetBEzsignfoldertypeDelegate returns the BEzsignfoldertypeDelegate field value if set, zero value otherwise.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDelegate() bool {
	if o == nil || IsNil(o.BEzsignfoldertypeDelegate) {
		var ret bool
		return ret
	}
	return *o.BEzsignfoldertypeDelegate
}

// GetBEzsignfoldertypeDelegateOk returns a tuple with the BEzsignfoldertypeDelegate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeDelegateOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignfoldertypeDelegate) {
		return nil, false
	}
	return o.BEzsignfoldertypeDelegate, true
}

// HasBEzsignfoldertypeDelegate returns a boolean if a field has been set.
func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeDelegate() bool {
	if o != nil && !IsNil(o.BEzsignfoldertypeDelegate) {
		return true
	}

	return false
}

// SetBEzsignfoldertypeDelegate gets a reference to the given bool and assigns it to the BEzsignfoldertypeDelegate field.
func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeDelegate(v bool) {
	o.BEzsignfoldertypeDelegate = &v
}

// GetBEzsignfoldertypeReassign returns the BEzsignfoldertypeReassign field value if set, zero value otherwise.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassign() bool {
	if o == nil || IsNil(o.BEzsignfoldertypeReassign) {
		var ret bool
		return ret
	}
	return *o.BEzsignfoldertypeReassign
}

// GetBEzsignfoldertypeReassignOk returns a tuple with the BEzsignfoldertypeReassign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEzsignfoldertypeResponse) GetBEzsignfoldertypeReassignOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignfoldertypeReassign) {
		return nil, false
	}
	return o.BEzsignfoldertypeReassign, true
}

// HasBEzsignfoldertypeReassign returns a boolean if a field has been set.
func (o *CustomEzsignfoldertypeResponse) HasBEzsignfoldertypeReassign() bool {
	if o != nil && !IsNil(o.BEzsignfoldertypeReassign) {
		return true
	}

	return false
}

// SetBEzsignfoldertypeReassign gets a reference to the given bool and assigns it to the BEzsignfoldertypeReassign field.
func (o *CustomEzsignfoldertypeResponse) SetBEzsignfoldertypeReassign(v bool) {
	o.BEzsignfoldertypeReassign = &v
}

func (o CustomEzsignfoldertypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomEzsignfoldertypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignfoldertypeID"] = o.PkiEzsignfoldertypeID
	toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	toSerialize["bEzsignfoldertypeIncludeproofsigner"] = o.BEzsignfoldertypeIncludeproofsigner
	toSerialize["bEzsignfoldertypeIncludeproofuser"] = o.BEzsignfoldertypeIncludeproofuser
	if !IsNil(o.BEzsignfoldertypeDelegate) {
		toSerialize["bEzsignfoldertypeDelegate"] = o.BEzsignfoldertypeDelegate
	}
	if !IsNil(o.BEzsignfoldertypeReassign) {
		toSerialize["bEzsignfoldertypeReassign"] = o.BEzsignfoldertypeReassign
	}
	return toSerialize, nil
}

type NullableCustomEzsignfoldertypeResponse struct {
	value *CustomEzsignfoldertypeResponse
	isSet bool
}

func (v NullableCustomEzsignfoldertypeResponse) Get() *CustomEzsignfoldertypeResponse {
	return v.value
}

func (v *NullableCustomEzsignfoldertypeResponse) Set(val *CustomEzsignfoldertypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEzsignfoldertypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEzsignfoldertypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEzsignfoldertypeResponse(val *CustomEzsignfoldertypeResponse) *NullableCustomEzsignfoldertypeResponse {
	return &NullableCustomEzsignfoldertypeResponse{value: val, isSet: true}
}

func (v NullableCustomEzsignfoldertypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEzsignfoldertypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


