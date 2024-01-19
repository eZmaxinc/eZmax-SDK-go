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
	"bytes"
	"fmt"
)

// checks if the DiscussionmessageResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscussionmessageResponseCompound{}

// DiscussionmessageResponseCompound A Discussionmessage Object and children
type DiscussionmessageResponseCompound struct {
	// The unique ID of the Discussionmessage
	PkiDiscussionmessageID int32 `json:"pkiDiscussionmessageID"`
	// The unique ID of the Discussion
	FkiDiscussionID int32 `json:"fkiDiscussionID"`
	// The unique ID of the Discussionmembership
	FkiDiscussionmembershipID *int32 `json:"fkiDiscussionmembershipID,omitempty"`
	// The unique ID of the Discussionmembership
	FkiDiscussionmembershipIDActionrequired *int32 `json:"fkiDiscussionmembershipIDActionrequired,omitempty"`
	EDiscussionmessageStatus FieldEDiscussionmessageStatus `json:"eDiscussionmessageStatus"`
	// The content of the Discussionmessage
	TDiscussionmessageContent string `json:"tDiscussionmessageContent"`
	// The name the creator of the Discussionmessage.
	SDiscussionmessageCreatorname string `json:"sDiscussionmessageCreatorname"`
	// The name the Actionrequired of the Discussionmessage.
	SDiscussionmessageActionrequiredname *string `json:"sDiscussionmessageActionrequiredname,omitempty"`
	ObjAudit CommonAudit `json:"objAudit"`
}

type _DiscussionmessageResponseCompound DiscussionmessageResponseCompound

// NewDiscussionmessageResponseCompound instantiates a new DiscussionmessageResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscussionmessageResponseCompound(pkiDiscussionmessageID int32, fkiDiscussionID int32, eDiscussionmessageStatus FieldEDiscussionmessageStatus, tDiscussionmessageContent string, sDiscussionmessageCreatorname string, objAudit CommonAudit) *DiscussionmessageResponseCompound {
	this := DiscussionmessageResponseCompound{}
	this.PkiDiscussionmessageID = pkiDiscussionmessageID
	this.FkiDiscussionID = fkiDiscussionID
	this.EDiscussionmessageStatus = eDiscussionmessageStatus
	this.TDiscussionmessageContent = tDiscussionmessageContent
	this.SDiscussionmessageCreatorname = sDiscussionmessageCreatorname
	this.ObjAudit = objAudit
	return &this
}

// NewDiscussionmessageResponseCompoundWithDefaults instantiates a new DiscussionmessageResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscussionmessageResponseCompoundWithDefaults() *DiscussionmessageResponseCompound {
	this := DiscussionmessageResponseCompound{}
	return &this
}

// GetPkiDiscussionmessageID returns the PkiDiscussionmessageID field value
func (o *DiscussionmessageResponseCompound) GetPkiDiscussionmessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiDiscussionmessageID
}

// GetPkiDiscussionmessageIDOk returns a tuple with the PkiDiscussionmessageID field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetPkiDiscussionmessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiDiscussionmessageID, true
}

// SetPkiDiscussionmessageID sets field value
func (o *DiscussionmessageResponseCompound) SetPkiDiscussionmessageID(v int32) {
	o.PkiDiscussionmessageID = v
}

// GetFkiDiscussionID returns the FkiDiscussionID field value
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiDiscussionID
}

// GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiDiscussionID, true
}

// SetFkiDiscussionID sets field value
func (o *DiscussionmessageResponseCompound) SetFkiDiscussionID(v int32) {
	o.FkiDiscussionID = v
}

// GetFkiDiscussionmembershipID returns the FkiDiscussionmembershipID field value if set, zero value otherwise.
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipID() int32 {
	if o == nil || IsNil(o.FkiDiscussionmembershipID) {
		var ret int32
		return ret
	}
	return *o.FkiDiscussionmembershipID
}

// GetFkiDiscussionmembershipIDOk returns a tuple with the FkiDiscussionmembershipID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiDiscussionmembershipID) {
		return nil, false
	}
	return o.FkiDiscussionmembershipID, true
}

// HasFkiDiscussionmembershipID returns a boolean if a field has been set.
func (o *DiscussionmessageResponseCompound) HasFkiDiscussionmembershipID() bool {
	if o != nil && !IsNil(o.FkiDiscussionmembershipID) {
		return true
	}

	return false
}

// SetFkiDiscussionmembershipID gets a reference to the given int32 and assigns it to the FkiDiscussionmembershipID field.
func (o *DiscussionmessageResponseCompound) SetFkiDiscussionmembershipID(v int32) {
	o.FkiDiscussionmembershipID = &v
}

// GetFkiDiscussionmembershipIDActionrequired returns the FkiDiscussionmembershipIDActionrequired field value if set, zero value otherwise.
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDActionrequired() int32 {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		var ret int32
		return ret
	}
	return *o.FkiDiscussionmembershipIDActionrequired
}

// GetFkiDiscussionmembershipIDActionrequiredOk returns a tuple with the FkiDiscussionmembershipIDActionrequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetFkiDiscussionmembershipIDActionrequiredOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return nil, false
	}
	return o.FkiDiscussionmembershipIDActionrequired, true
}

// HasFkiDiscussionmembershipIDActionrequired returns a boolean if a field has been set.
func (o *DiscussionmessageResponseCompound) HasFkiDiscussionmembershipIDActionrequired() bool {
	if o != nil && !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		return true
	}

	return false
}

// SetFkiDiscussionmembershipIDActionrequired gets a reference to the given int32 and assigns it to the FkiDiscussionmembershipIDActionrequired field.
func (o *DiscussionmessageResponseCompound) SetFkiDiscussionmembershipIDActionrequired(v int32) {
	o.FkiDiscussionmembershipIDActionrequired = &v
}

// GetEDiscussionmessageStatus returns the EDiscussionmessageStatus field value
func (o *DiscussionmessageResponseCompound) GetEDiscussionmessageStatus() FieldEDiscussionmessageStatus {
	if o == nil {
		var ret FieldEDiscussionmessageStatus
		return ret
	}

	return o.EDiscussionmessageStatus
}

// GetEDiscussionmessageStatusOk returns a tuple with the EDiscussionmessageStatus field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetEDiscussionmessageStatusOk() (*FieldEDiscussionmessageStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EDiscussionmessageStatus, true
}

// SetEDiscussionmessageStatus sets field value
func (o *DiscussionmessageResponseCompound) SetEDiscussionmessageStatus(v FieldEDiscussionmessageStatus) {
	o.EDiscussionmessageStatus = v
}

// GetTDiscussionmessageContent returns the TDiscussionmessageContent field value
func (o *DiscussionmessageResponseCompound) GetTDiscussionmessageContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TDiscussionmessageContent
}

// GetTDiscussionmessageContentOk returns a tuple with the TDiscussionmessageContent field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetTDiscussionmessageContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TDiscussionmessageContent, true
}

// SetTDiscussionmessageContent sets field value
func (o *DiscussionmessageResponseCompound) SetTDiscussionmessageContent(v string) {
	o.TDiscussionmessageContent = v
}

// GetSDiscussionmessageCreatorname returns the SDiscussionmessageCreatorname field value
func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageCreatorname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDiscussionmessageCreatorname
}

// GetSDiscussionmessageCreatornameOk returns a tuple with the SDiscussionmessageCreatorname field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageCreatornameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SDiscussionmessageCreatorname, true
}

// SetSDiscussionmessageCreatorname sets field value
func (o *DiscussionmessageResponseCompound) SetSDiscussionmessageCreatorname(v string) {
	o.SDiscussionmessageCreatorname = v
}

// GetSDiscussionmessageActionrequiredname returns the SDiscussionmessageActionrequiredname field value if set, zero value otherwise.
func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageActionrequiredname() string {
	if o == nil || IsNil(o.SDiscussionmessageActionrequiredname) {
		var ret string
		return ret
	}
	return *o.SDiscussionmessageActionrequiredname
}

// GetSDiscussionmessageActionrequirednameOk returns a tuple with the SDiscussionmessageActionrequiredname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetSDiscussionmessageActionrequirednameOk() (*string, bool) {
	if o == nil || IsNil(o.SDiscussionmessageActionrequiredname) {
		return nil, false
	}
	return o.SDiscussionmessageActionrequiredname, true
}

// HasSDiscussionmessageActionrequiredname returns a boolean if a field has been set.
func (o *DiscussionmessageResponseCompound) HasSDiscussionmessageActionrequiredname() bool {
	if o != nil && !IsNil(o.SDiscussionmessageActionrequiredname) {
		return true
	}

	return false
}

// SetSDiscussionmessageActionrequiredname gets a reference to the given string and assigns it to the SDiscussionmessageActionrequiredname field.
func (o *DiscussionmessageResponseCompound) SetSDiscussionmessageActionrequiredname(v string) {
	o.SDiscussionmessageActionrequiredname = &v
}

// GetObjAudit returns the ObjAudit field value
func (o *DiscussionmessageResponseCompound) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *DiscussionmessageResponseCompound) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *DiscussionmessageResponseCompound) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o DiscussionmessageResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscussionmessageResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiDiscussionmessageID"] = o.PkiDiscussionmessageID
	toSerialize["fkiDiscussionID"] = o.FkiDiscussionID
	if !IsNil(o.FkiDiscussionmembershipID) {
		toSerialize["fkiDiscussionmembershipID"] = o.FkiDiscussionmembershipID
	}
	if !IsNil(o.FkiDiscussionmembershipIDActionrequired) {
		toSerialize["fkiDiscussionmembershipIDActionrequired"] = o.FkiDiscussionmembershipIDActionrequired
	}
	toSerialize["eDiscussionmessageStatus"] = o.EDiscussionmessageStatus
	toSerialize["tDiscussionmessageContent"] = o.TDiscussionmessageContent
	toSerialize["sDiscussionmessageCreatorname"] = o.SDiscussionmessageCreatorname
	if !IsNil(o.SDiscussionmessageActionrequiredname) {
		toSerialize["sDiscussionmessageActionrequiredname"] = o.SDiscussionmessageActionrequiredname
	}
	toSerialize["objAudit"] = o.ObjAudit
	return toSerialize, nil
}

func (o *DiscussionmessageResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiDiscussionmessageID",
		"fkiDiscussionID",
		"eDiscussionmessageStatus",
		"tDiscussionmessageContent",
		"sDiscussionmessageCreatorname",
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

	varDiscussionmessageResponseCompound := _DiscussionmessageResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiscussionmessageResponseCompound)

	if err != nil {
		return err
	}

	*o = DiscussionmessageResponseCompound(varDiscussionmessageResponseCompound)

	return err
}

type NullableDiscussionmessageResponseCompound struct {
	value *DiscussionmessageResponseCompound
	isSet bool
}

func (v NullableDiscussionmessageResponseCompound) Get() *DiscussionmessageResponseCompound {
	return v.value
}

func (v *NullableDiscussionmessageResponseCompound) Set(val *DiscussionmessageResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscussionmessageResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscussionmessageResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscussionmessageResponseCompound(val *DiscussionmessageResponseCompound) *NullableDiscussionmessageResponseCompound {
	return &NullableDiscussionmessageResponseCompound{value: val, isSet: true}
}

func (v NullableDiscussionmessageResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscussionmessageResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


