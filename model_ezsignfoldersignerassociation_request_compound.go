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

// checks if the EzsignfoldersignerassociationRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationRequestCompound{}

// EzsignfoldersignerassociationRequestCompound An Ezsignfoldersignerassociation Object and children to create a complete structure
type EzsignfoldersignerassociationRequestCompound struct {
	// The unique ID of the Ezsignfoldersignerassociation
	PkiEzsignfoldersignerassociationID *int32 `json:"pkiEzsignfoldersignerassociationID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Ezsignsignergroup
	FkiEzsignsignergroupID *int32 `json:"fkiEzsignsignergroupID,omitempty"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain't required to sign the document.
	BEzsignfoldersignerassociationReceivecopy *bool `json:"bEzsignfoldersignerassociationReceivecopy,omitempty"`
	// A custom text message that will be added to the email sent.
	TEzsignfoldersignerassociationMessage *string `json:"tEzsignfoldersignerassociationMessage,omitempty"`
	ObjEzsignsigner *EzsignsignerRequestCompound `json:"objEzsignsigner,omitempty"`
}

type _EzsignfoldersignerassociationRequestCompound EzsignfoldersignerassociationRequestCompound

// NewEzsignfoldersignerassociationRequestCompound instantiates a new EzsignfoldersignerassociationRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationRequestCompound(fkiEzsignfolderID int32) *EzsignfoldersignerassociationRequestCompound {
	this := EzsignfoldersignerassociationRequestCompound{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	return &this
}

// NewEzsignfoldersignerassociationRequestCompoundWithDefaults instantiates a new EzsignfoldersignerassociationRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationRequestCompoundWithDefaults() *EzsignfoldersignerassociationRequestCompound {
	this := EzsignfoldersignerassociationRequestCompound{}
	return &this
}

// GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetPkiEzsignfoldersignerassociationID() int32 {
	if o == nil || IsNil(o.PkiEzsignfoldersignerassociationID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignfoldersignerassociationID
}

// GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignfoldersignerassociationID) {
		return nil, false
	}
	return o.PkiEzsignfoldersignerassociationID, true
}

// HasPkiEzsignfoldersignerassociationID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasPkiEzsignfoldersignerassociationID() bool {
	if o != nil && !IsNil(o.PkiEzsignfoldersignerassociationID) {
		return true
	}

	return false
}

// SetPkiEzsignfoldersignerassociationID gets a reference to the given int32 and assigns it to the PkiEzsignfoldersignerassociationID field.
func (o *EzsignfoldersignerassociationRequestCompound) SetPkiEzsignfoldersignerassociationID(v int32) {
	o.PkiEzsignfoldersignerassociationID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignfoldersignerassociationRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiEzsignsignergroupID returns the FkiEzsignsignergroupID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignsignergroupID() int32 {
	if o == nil || IsNil(o.FkiEzsignsignergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiEzsignsignergroupID
}

// GetFkiEzsignsignergroupIDOk returns a tuple with the FkiEzsignsignergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignsignergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzsignsignergroupID) {
		return nil, false
	}
	return o.FkiEzsignsignergroupID, true
}

// HasFkiEzsignsignergroupID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasFkiEzsignsignergroupID() bool {
	if o != nil && !IsNil(o.FkiEzsignsignergroupID) {
		return true
	}

	return false
}

// SetFkiEzsignsignergroupID gets a reference to the given int32 and assigns it to the FkiEzsignsignergroupID field.
func (o *EzsignfoldersignerassociationRequestCompound) SetFkiEzsignsignergroupID(v int32) {
	o.FkiEzsignsignergroupID = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsignfoldersignerassociationRequestCompound) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetBEzsignfoldersignerassociationReceivecopy() bool {
	if o == nil || IsNil(o.BEzsignfoldersignerassociationReceivecopy) {
		var ret bool
		return ret
	}
	return *o.BEzsignfoldersignerassociationReceivecopy
}

// GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsignfoldersignerassociationReceivecopy) {
		return nil, false
	}
	return o.BEzsignfoldersignerassociationReceivecopy, true
}

// HasBEzsignfoldersignerassociationReceivecopy returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasBEzsignfoldersignerassociationReceivecopy() bool {
	if o != nil && !IsNil(o.BEzsignfoldersignerassociationReceivecopy) {
		return true
	}

	return false
}

// SetBEzsignfoldersignerassociationReceivecopy gets a reference to the given bool and assigns it to the BEzsignfoldersignerassociationReceivecopy field.
func (o *EzsignfoldersignerassociationRequestCompound) SetBEzsignfoldersignerassociationReceivecopy(v bool) {
	o.BEzsignfoldersignerassociationReceivecopy = &v
}

// GetTEzsignfoldersignerassociationMessage returns the TEzsignfoldersignerassociationMessage field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetTEzsignfoldersignerassociationMessage() string {
	if o == nil || IsNil(o.TEzsignfoldersignerassociationMessage) {
		var ret string
		return ret
	}
	return *o.TEzsignfoldersignerassociationMessage
}

// GetTEzsignfoldersignerassociationMessageOk returns a tuple with the TEzsignfoldersignerassociationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetTEzsignfoldersignerassociationMessageOk() (*string, bool) {
	if o == nil || IsNil(o.TEzsignfoldersignerassociationMessage) {
		return nil, false
	}
	return o.TEzsignfoldersignerassociationMessage, true
}

// HasTEzsignfoldersignerassociationMessage returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasTEzsignfoldersignerassociationMessage() bool {
	if o != nil && !IsNil(o.TEzsignfoldersignerassociationMessage) {
		return true
	}

	return false
}

// SetTEzsignfoldersignerassociationMessage gets a reference to the given string and assigns it to the TEzsignfoldersignerassociationMessage field.
func (o *EzsignfoldersignerassociationRequestCompound) SetTEzsignfoldersignerassociationMessage(v string) {
	o.TEzsignfoldersignerassociationMessage = &v
}

// GetObjEzsignsigner returns the ObjEzsignsigner field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsigner() EzsignsignerRequestCompound {
	if o == nil || IsNil(o.ObjEzsignsigner) {
		var ret EzsignsignerRequestCompound
		return ret
	}
	return *o.ObjEzsignsigner
}

// GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsignerOk() (*EzsignsignerRequestCompound, bool) {
	if o == nil || IsNil(o.ObjEzsignsigner) {
		return nil, false
	}
	return o.ObjEzsignsigner, true
}

// HasObjEzsignsigner returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasObjEzsignsigner() bool {
	if o != nil && !IsNil(o.ObjEzsignsigner) {
		return true
	}

	return false
}

// SetObjEzsignsigner gets a reference to the given EzsignsignerRequestCompound and assigns it to the ObjEzsignsigner field.
func (o *EzsignfoldersignerassociationRequestCompound) SetObjEzsignsigner(v EzsignsignerRequestCompound) {
	o.ObjEzsignsigner = &v
}

func (o EzsignfoldersignerassociationRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignfoldersignerassociationID) {
		toSerialize["pkiEzsignfoldersignerassociationID"] = o.PkiEzsignfoldersignerassociationID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiEzsignsignergroupID) {
		toSerialize["fkiEzsignsignergroupID"] = o.FkiEzsignsignergroupID
	}
	toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	if !IsNil(o.BEzsignfoldersignerassociationReceivecopy) {
		toSerialize["bEzsignfoldersignerassociationReceivecopy"] = o.BEzsignfoldersignerassociationReceivecopy
	}
	if !IsNil(o.TEzsignfoldersignerassociationMessage) {
		toSerialize["tEzsignfoldersignerassociationMessage"] = o.TEzsignfoldersignerassociationMessage
	}
	if !IsNil(o.ObjEzsignsigner) {
		toSerialize["objEzsignsigner"] = o.ObjEzsignsigner
	}
	return toSerialize, nil
}

func (o *EzsignfoldersignerassociationRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzsignfolderID",
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

	varEzsignfoldersignerassociationRequestCompound := _EzsignfoldersignerassociationRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldersignerassociationRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsignfoldersignerassociationRequestCompound(varEzsignfoldersignerassociationRequestCompound)

	return err
}

type NullableEzsignfoldersignerassociationRequestCompound struct {
	value *EzsignfoldersignerassociationRequestCompound
	isSet bool
}

func (v NullableEzsignfoldersignerassociationRequestCompound) Get() *EzsignfoldersignerassociationRequestCompound {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) Set(val *EzsignfoldersignerassociationRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationRequestCompound(val *EzsignfoldersignerassociationRequestCompound) *NullableEzsignfoldersignerassociationRequestCompound {
	return &NullableEzsignfoldersignerassociationRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


