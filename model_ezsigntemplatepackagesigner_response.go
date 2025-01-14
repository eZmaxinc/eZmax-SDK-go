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

// checks if the EzsigntemplatepackagesignerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatepackagesignerResponse{}

// EzsigntemplatepackagesignerResponse A Ezsigntemplatepackagesigner Object
type EzsigntemplatepackagesignerResponse struct {
	// The unique ID of the Ezsigntemplatepackagesigner
	PkiEzsigntemplatepackagesignerID int32 `json:"pkiEzsigntemplatepackagesignerID"`
	// The unique ID of the Ezsigntemplatepackage
	FkiEzsigntemplatepackageID int32 `json:"fkiEzsigntemplatepackageID"`
	// The unique ID of the Ezdoctemplatedocument
	FkiEzdoctemplatedocumentID *int32 `json:"fkiEzdoctemplatedocumentID,omitempty"`
	// The unique ID of the User
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// The unique ID of the Usergroup
	FkiUsergroupID *int32 `json:"fkiUsergroupID,omitempty"`
	// The name of the Ezdoctemplatedocument in the language of the requester
	SEzdoctemplatedocumentNameX *string `json:"sEzdoctemplatedocumentNameX,omitempty" validate:"regexp=^.{0,50}$"`
	// If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain't required to sign the document.
	BEzsigntemplatepackagesignerReceivecopy *bool `json:"bEzsigntemplatepackagesignerReceivecopy,omitempty"`
	EEzsigntemplatepackagesignerMapping *FieldEEzsigntemplatepackagesignerMapping `json:"eEzsigntemplatepackagesignerMapping,omitempty"`
	// The description of the Ezsigntemplatepackagesigner
	SEzsigntemplatepackagesignerDescription string `json:"sEzsigntemplatepackagesignerDescription"`
	// The description of the User in the language of the requester
	SUserName *string `json:"sUserName,omitempty"`
	// The Name of the Usergroup in the language of the requester
	SUsergroupNameX *string `json:"sUsergroupNameX,omitempty" validate:"regexp=^.{0,50}$"`
}

type _EzsigntemplatepackagesignerResponse EzsigntemplatepackagesignerResponse

// NewEzsigntemplatepackagesignerResponse instantiates a new EzsigntemplatepackagesignerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatepackagesignerResponse(pkiEzsigntemplatepackagesignerID int32, fkiEzsigntemplatepackageID int32, sEzsigntemplatepackagesignerDescription string) *EzsigntemplatepackagesignerResponse {
	this := EzsigntemplatepackagesignerResponse{}
	this.PkiEzsigntemplatepackagesignerID = pkiEzsigntemplatepackagesignerID
	this.FkiEzsigntemplatepackageID = fkiEzsigntemplatepackageID
	var eEzsigntemplatepackagesignerMapping FieldEEzsigntemplatepackagesignerMapping = MANUAL
	this.EEzsigntemplatepackagesignerMapping = &eEzsigntemplatepackagesignerMapping
	this.SEzsigntemplatepackagesignerDescription = sEzsigntemplatepackagesignerDescription
	return &this
}

// NewEzsigntemplatepackagesignerResponseWithDefaults instantiates a new EzsigntemplatepackagesignerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatepackagesignerResponseWithDefaults() *EzsigntemplatepackagesignerResponse {
	this := EzsigntemplatepackagesignerResponse{}
	var eEzsigntemplatepackagesignerMapping FieldEEzsigntemplatepackagesignerMapping = MANUAL
	this.EEzsigntemplatepackagesignerMapping = &eEzsigntemplatepackagesignerMapping
	return &this
}

// GetPkiEzsigntemplatepackagesignerID returns the PkiEzsigntemplatepackagesignerID field value
func (o *EzsigntemplatepackagesignerResponse) GetPkiEzsigntemplatepackagesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatepackagesignerID
}

// GetPkiEzsigntemplatepackagesignerIDOk returns a tuple with the PkiEzsigntemplatepackagesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetPkiEzsigntemplatepackagesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatepackagesignerID, true
}

// SetPkiEzsigntemplatepackagesignerID sets field value
func (o *EzsigntemplatepackagesignerResponse) SetPkiEzsigntemplatepackagesignerID(v int32) {
	o.PkiEzsigntemplatepackagesignerID = v
}

// GetFkiEzsigntemplatepackageID returns the FkiEzsigntemplatepackageID field value
func (o *EzsigntemplatepackagesignerResponse) GetFkiEzsigntemplatepackageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplatepackageID
}

// GetFkiEzsigntemplatepackageIDOk returns a tuple with the FkiEzsigntemplatepackageID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetFkiEzsigntemplatepackageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplatepackageID, true
}

// SetFkiEzsigntemplatepackageID sets field value
func (o *EzsigntemplatepackagesignerResponse) SetFkiEzsigntemplatepackageID(v int32) {
	o.FkiEzsigntemplatepackageID = v
}

// GetFkiEzdoctemplatedocumentID returns the FkiEzdoctemplatedocumentID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetFkiEzdoctemplatedocumentID() int32 {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		var ret int32
		return ret
	}
	return *o.FkiEzdoctemplatedocumentID
}

// GetFkiEzdoctemplatedocumentIDOk returns a tuple with the FkiEzdoctemplatedocumentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetFkiEzdoctemplatedocumentIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiEzdoctemplatedocumentID) {
		return nil, false
	}
	return o.FkiEzdoctemplatedocumentID, true
}

// HasFkiEzdoctemplatedocumentID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasFkiEzdoctemplatedocumentID() bool {
	if o != nil && !IsNil(o.FkiEzdoctemplatedocumentID) {
		return true
	}

	return false
}

// SetFkiEzdoctemplatedocumentID gets a reference to the given int32 and assigns it to the FkiEzdoctemplatedocumentID field.
func (o *EzsigntemplatepackagesignerResponse) SetFkiEzdoctemplatedocumentID(v int32) {
	o.FkiEzdoctemplatedocumentID = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetFkiUserID() int32 {
	if o == nil || IsNil(o.FkiUserID) {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUserID) {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasFkiUserID() bool {
	if o != nil && !IsNil(o.FkiUserID) {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsigntemplatepackagesignerResponse) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetFkiUsergroupID() int32 {
	if o == nil || IsNil(o.FkiUsergroupID) {
		var ret int32
		return ret
	}
	return *o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil || IsNil(o.FkiUsergroupID) {
		return nil, false
	}
	return o.FkiUsergroupID, true
}

// HasFkiUsergroupID returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasFkiUsergroupID() bool {
	if o != nil && !IsNil(o.FkiUsergroupID) {
		return true
	}

	return false
}

// SetFkiUsergroupID gets a reference to the given int32 and assigns it to the FkiUsergroupID field.
func (o *EzsigntemplatepackagesignerResponse) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = &v
}

// GetSEzdoctemplatedocumentNameX returns the SEzdoctemplatedocumentNameX field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetSEzdoctemplatedocumentNameX() string {
	if o == nil || IsNil(o.SEzdoctemplatedocumentNameX) {
		var ret string
		return ret
	}
	return *o.SEzdoctemplatedocumentNameX
}

// GetSEzdoctemplatedocumentNameXOk returns a tuple with the SEzdoctemplatedocumentNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetSEzdoctemplatedocumentNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SEzdoctemplatedocumentNameX) {
		return nil, false
	}
	return o.SEzdoctemplatedocumentNameX, true
}

// HasSEzdoctemplatedocumentNameX returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasSEzdoctemplatedocumentNameX() bool {
	if o != nil && !IsNil(o.SEzdoctemplatedocumentNameX) {
		return true
	}

	return false
}

// SetSEzdoctemplatedocumentNameX gets a reference to the given string and assigns it to the SEzdoctemplatedocumentNameX field.
func (o *EzsigntemplatepackagesignerResponse) SetSEzdoctemplatedocumentNameX(v string) {
	o.SEzdoctemplatedocumentNameX = &v
}

// GetBEzsigntemplatepackagesignerReceivecopy returns the BEzsigntemplatepackagesignerReceivecopy field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetBEzsigntemplatepackagesignerReceivecopy() bool {
	if o == nil || IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		var ret bool
		return ret
	}
	return *o.BEzsigntemplatepackagesignerReceivecopy
}

// GetBEzsigntemplatepackagesignerReceivecopyOk returns a tuple with the BEzsigntemplatepackagesignerReceivecopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetBEzsigntemplatepackagesignerReceivecopyOk() (*bool, bool) {
	if o == nil || IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		return nil, false
	}
	return o.BEzsigntemplatepackagesignerReceivecopy, true
}

// HasBEzsigntemplatepackagesignerReceivecopy returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasBEzsigntemplatepackagesignerReceivecopy() bool {
	if o != nil && !IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		return true
	}

	return false
}

// SetBEzsigntemplatepackagesignerReceivecopy gets a reference to the given bool and assigns it to the BEzsigntemplatepackagesignerReceivecopy field.
func (o *EzsigntemplatepackagesignerResponse) SetBEzsigntemplatepackagesignerReceivecopy(v bool) {
	o.BEzsigntemplatepackagesignerReceivecopy = &v
}

// GetEEzsigntemplatepackagesignerMapping returns the EEzsigntemplatepackagesignerMapping field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetEEzsigntemplatepackagesignerMapping() FieldEEzsigntemplatepackagesignerMapping {
	if o == nil || IsNil(o.EEzsigntemplatepackagesignerMapping) {
		var ret FieldEEzsigntemplatepackagesignerMapping
		return ret
	}
	return *o.EEzsigntemplatepackagesignerMapping
}

// GetEEzsigntemplatepackagesignerMappingOk returns a tuple with the EEzsigntemplatepackagesignerMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetEEzsigntemplatepackagesignerMappingOk() (*FieldEEzsigntemplatepackagesignerMapping, bool) {
	if o == nil || IsNil(o.EEzsigntemplatepackagesignerMapping) {
		return nil, false
	}
	return o.EEzsigntemplatepackagesignerMapping, true
}

// HasEEzsigntemplatepackagesignerMapping returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasEEzsigntemplatepackagesignerMapping() bool {
	if o != nil && !IsNil(o.EEzsigntemplatepackagesignerMapping) {
		return true
	}

	return false
}

// SetEEzsigntemplatepackagesignerMapping gets a reference to the given FieldEEzsigntemplatepackagesignerMapping and assigns it to the EEzsigntemplatepackagesignerMapping field.
func (o *EzsigntemplatepackagesignerResponse) SetEEzsigntemplatepackagesignerMapping(v FieldEEzsigntemplatepackagesignerMapping) {
	o.EEzsigntemplatepackagesignerMapping = &v
}

// GetSEzsigntemplatepackagesignerDescription returns the SEzsigntemplatepackagesignerDescription field value
func (o *EzsigntemplatepackagesignerResponse) GetSEzsigntemplatepackagesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatepackagesignerDescription
}

// GetSEzsigntemplatepackagesignerDescriptionOk returns a tuple with the SEzsigntemplatepackagesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetSEzsigntemplatepackagesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatepackagesignerDescription, true
}

// SetSEzsigntemplatepackagesignerDescription sets field value
func (o *EzsigntemplatepackagesignerResponse) SetSEzsigntemplatepackagesignerDescription(v string) {
	o.SEzsigntemplatepackagesignerDescription = v
}

// GetSUserName returns the SUserName field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetSUserName() string {
	if o == nil || IsNil(o.SUserName) {
		var ret string
		return ret
	}
	return *o.SUserName
}

// GetSUserNameOk returns a tuple with the SUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetSUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.SUserName) {
		return nil, false
	}
	return o.SUserName, true
}

// HasSUserName returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasSUserName() bool {
	if o != nil && !IsNil(o.SUserName) {
		return true
	}

	return false
}

// SetSUserName gets a reference to the given string and assigns it to the SUserName field.
func (o *EzsigntemplatepackagesignerResponse) SetSUserName(v string) {
	o.SUserName = &v
}

// GetSUsergroupNameX returns the SUsergroupNameX field value if set, zero value otherwise.
func (o *EzsigntemplatepackagesignerResponse) GetSUsergroupNameX() string {
	if o == nil || IsNil(o.SUsergroupNameX) {
		var ret string
		return ret
	}
	return *o.SUsergroupNameX
}

// GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatepackagesignerResponse) GetSUsergroupNameXOk() (*string, bool) {
	if o == nil || IsNil(o.SUsergroupNameX) {
		return nil, false
	}
	return o.SUsergroupNameX, true
}

// HasSUsergroupNameX returns a boolean if a field has been set.
func (o *EzsigntemplatepackagesignerResponse) HasSUsergroupNameX() bool {
	if o != nil && !IsNil(o.SUsergroupNameX) {
		return true
	}

	return false
}

// SetSUsergroupNameX gets a reference to the given string and assigns it to the SUsergroupNameX field.
func (o *EzsigntemplatepackagesignerResponse) SetSUsergroupNameX(v string) {
	o.SUsergroupNameX = &v
}

func (o EzsigntemplatepackagesignerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatepackagesignerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatepackagesignerID"] = o.PkiEzsigntemplatepackagesignerID
	toSerialize["fkiEzsigntemplatepackageID"] = o.FkiEzsigntemplatepackageID
	if !IsNil(o.FkiEzdoctemplatedocumentID) {
		toSerialize["fkiEzdoctemplatedocumentID"] = o.FkiEzdoctemplatedocumentID
	}
	if !IsNil(o.FkiUserID) {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if !IsNil(o.FkiUsergroupID) {
		toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	}
	if !IsNil(o.SEzdoctemplatedocumentNameX) {
		toSerialize["sEzdoctemplatedocumentNameX"] = o.SEzdoctemplatedocumentNameX
	}
	if !IsNil(o.BEzsigntemplatepackagesignerReceivecopy) {
		toSerialize["bEzsigntemplatepackagesignerReceivecopy"] = o.BEzsigntemplatepackagesignerReceivecopy
	}
	if !IsNil(o.EEzsigntemplatepackagesignerMapping) {
		toSerialize["eEzsigntemplatepackagesignerMapping"] = o.EEzsigntemplatepackagesignerMapping
	}
	toSerialize["sEzsigntemplatepackagesignerDescription"] = o.SEzsigntemplatepackagesignerDescription
	if !IsNil(o.SUserName) {
		toSerialize["sUserName"] = o.SUserName
	}
	if !IsNil(o.SUsergroupNameX) {
		toSerialize["sUsergroupNameX"] = o.SUsergroupNameX
	}
	return toSerialize, nil
}

func (o *EzsigntemplatepackagesignerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatepackagesignerID",
		"fkiEzsigntemplatepackageID",
		"sEzsigntemplatepackagesignerDescription",
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

	varEzsigntemplatepackagesignerResponse := _EzsigntemplatepackagesignerResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatepackagesignerResponse)

	if err != nil {
		return err
	}

	*o = EzsigntemplatepackagesignerResponse(varEzsigntemplatepackagesignerResponse)

	return err
}

type NullableEzsigntemplatepackagesignerResponse struct {
	value *EzsigntemplatepackagesignerResponse
	isSet bool
}

func (v NullableEzsigntemplatepackagesignerResponse) Get() *EzsigntemplatepackagesignerResponse {
	return v.value
}

func (v *NullableEzsigntemplatepackagesignerResponse) Set(val *EzsigntemplatepackagesignerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatepackagesignerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatepackagesignerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatepackagesignerResponse(val *EzsigntemplatepackagesignerResponse) *NullableEzsigntemplatepackagesignerResponse {
	return &NullableEzsigntemplatepackagesignerResponse{value: val, isSet: true}
}

func (v NullableEzsigntemplatepackagesignerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatepackagesignerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


